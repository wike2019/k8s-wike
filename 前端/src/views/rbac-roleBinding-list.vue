<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>RoleBinding列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                 <div class="name_space_choose">
                   <span>命名空间:</span>
                   <el-select placeholder="选择命名空间" @change="changeNs" v-model="name_space">
                     <el-option v-for="ns in nsList "
                                :label="ns.name"
                                :value="ns.name"/>
                   </el-select>
                 </div>
                  <el-divider></el-divider>
                  <el-table
                      :data="List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">


               <el-table-column label="角色绑定名" >
                 <template #default="scope">
                   <p @click="()=>doTo('secret-detail',{name_space:scope.row.name_space,name:scope.row.name})">{{ scope.row.name }}</p>
                   <p>
                     <el-tag style="margin:5px 10px"  @close="rmUserFromBinding(scope.row.name_space,scope.row.name,sub.kind,sub.name)"  closable :type="(sub.kind=='User')?'info':(sub.kind=='ServiceAccount')? 'danger':'warning'" v-for="sub in scope.row.subject">
                       <span class="role">角色:</span>{{ sub.name }}
                       <span class="role">绑定名:</span>{{ scope.row.name }}
                       <span class="role">类型:</span>{{ sub.kind }}
                     </el-tag>
                   </p>

                 </template>
               </el-table-column>
                <el-table-column label="增加角色" width="380" >
                  <template #default="scope">
                    <i  v-show="addConfig.index==-1||addConfig.index!==scope.$index" class=" el-icon-circle-plus" @click="()=>showAdd(scope.$index,scope.row.name_space,scope.row.name)" ><span style="padding:0 5px">增加</span></i>
                    <div v-show="addConfig.index===scope.$index" >
                      <el-form :inline="true"  :model="addConfig" ref="formRef" :rules="rules">
                        <el-form-item  prop="kind"  label=" "  label-width="10px">
                          <el-select size="small"   @change="clear" v-model="addConfig.kind" style="width:100px;height:30px!important;">
                            <el-option label="User" value="User"/>
                            <el-option label="ServiceAccount" value="ServiceAccount"/>
                            <el-option label="Group"  value="Group" >
                            </el-option>
                          </el-select>
                        </el-form-item>
                        <el-form-item label=" "  label-width="10px" prop="name" v-show="addConfig.kind=='ServiceAccount'">
                            <el-select  size="small" v-model="addConfig.name" style="margin:0 8px;width: 180px">
                              <el-option v-for="sa in SaList "
                                         :label="sa.name"
                                         :value="sa.name"/>
                            </el-select>
                        </el-form-item>
                        <el-form-item  label=" "  label-width="10px" v-show="addConfig.kind!='ServiceAccount'"  prop="name" >
                          <el-input  style="margin:0 8px;width: 180px"  size="small" placeholder="输入用户名" v-model="addConfig.name" ></el-input>
                        </el-form-item>
                      <div style="margin-top:10px">
                            <i style="margin:0 8px"  class="ii el-icon-s-claim" @click="appendUserToBinding" > 保存</i>
                            <i style="margin:0 8px"   class="ii el-icon-remove" @click="()=>{addConfig.index=-1;addConfig.kind='User';addConfig.name=''}"> 取消</i>
                      </div>
                      </el-form>
                    </div>
                  </template>
               </el-table-column>
               <el-table-column label="命名空间" width="120" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.name_space }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="创建时间" width="160" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.create_time }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="操作" width="100" align="center">
                 <template #default="scope">
                   <el-button type="danger" @click="()=>rm(scope.row.name_space,scope.row.name )" icon="el-icon-delete" circle></el-button>
                 </template>
               </el-table-column>
    </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event)"
                      background
                      :current-page="current_page"
                      :page-size="5"
                      :hide-on-single-page="true"
                      layout="prev, pager, next"
                      :total="pageInfo">
                  </el-pagination>
                </div>

     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {ingressDel, ingressListByNs} from "../api/token/ingress/ingress";
import {secretDel, secretListByNs} from "../api/token/secret/secret";
import {doTo} from "../router";
import {
  addUserToBinding,
  deleteUserFromBinding,
  getRoleBindingList,
  getRoleList,
  roleBindingDel,
  roleDel
} from "../api/token/rbac";
import {getSaList} from "../api/token/sa/sa";

export default defineComponent({
  name: 'secret-list',
  components: {MainLayout},
  setup(){
    const apiGroup='rbac.authorization.k8s.io'
    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      List:reactive([]),
      pageInfo:0,
      current_page:1,
      name_space:"default",
      SaList:reactive([]),
      addConfig:reactive({
        index:-1,
        kind:'User',
        name:'',//用户名
        bindingName:'', //binding名称
        ns:'' //命名空间
      })
    })
    const rules= {
      kind:[
        {
          required: true,
          message: '角色类型必须选择',
        }
      ],
      name:[
        {
          required: true,
          message: '角色名称必须填写',
        }
      ],
    }
    let  formRef=ref(null)
    async function getData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
        await  changeNs();
      }catch (e){
        console.log(e)
      }
    }
    async function changeNs(){
      try {
          let tData=await getRoleBindingList(state.name_space,state.current_page)
          let T=tData.data.data
          state.List=T.Data
          state.pageInfo=T["Total"]
          state.current_page=T["Current"]
          let tData2=await getSaList(state.name_space)
          state.SaList=tData2.data.data
          state.addConfig.name=""
          state.addConfig.kind="User"
          state.addConfig.index=-1;
        }catch (e){
          console.log(e)
        }

    }
    getData()
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        let T=tData.result;
        if(tData.type=='roleBinding'&&tData.ns==state.name_space){

          state.List=T.Data
          state.pageInfo=T["Total"]
          state.current_page=T["Current"]
        }
        if(tData.type=='namespace'){
          getData()
        }
        if(tData.type=='sa'&&tData.ns==state.name_space){
          state.SaList=T.Data
        }
      }
    }

    function pageChange(page,ns){
      state.current_page=page
      changeNs()
    }
    function rm(ns,name){
      ElMessageBox.confirm(
          '你确定继续删除这个角色绑定操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                let result=await roleBindingDel(ns,name)
                if (result.data.code==200){
                  ElMessage("资源删除成功")
                }
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }
    function rmUserFromBinding(ns,name,kind,username){
      ElMessageBox.confirm(
          '你确定继续删除这个角色绑定操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
            try {
              let result=await  deleteUserFromBinding(ns,name,{kind,name:username,apiGroup})
              if (result.data.code==200){
                ElMessage("资源删除成功")
              }
            }catch (e) {
              ElMessage({
                type: 'info',
                message: '系统异常'+e,
              })
            }
          })
    }

    function showAdd(index,ns,bindingName){
      state.addConfig.index=index
      state.addConfig.ns=ns
      state.addConfig.bindingName=bindingName
    }
    async function appendUserToBinding(){
      formRef.value.validate(async (valid) => {
        if (state.addConfig.name == "") {
          ElMessage.error("用户名称不能为空")
          return
        }
        if (valid) {


          const {ns, bindingName, kind, name} = state.addConfig
          //构建出一个subject
          try {
            let result = await addUserToBinding(ns, bindingName, {kind, name, apiGroup})
            if (result.data.code == 200) {
              ElMessage("资源创建成功")
              state.addConfig.index = -1
              state.addConfig.name = ''
            }
          } catch (e) {
            ElMessage({
              type: 'info',
              message: '系统异常' + e,
            })
          }
        }
       })
    }
    function clear(){
      state.addConfig.name=''
    }
    return {...toRefs(state),pageChange,rm,changeNs,doTo,rmUserFromBinding,showAdd,appendUserToBinding,clear,rules,formRef}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>