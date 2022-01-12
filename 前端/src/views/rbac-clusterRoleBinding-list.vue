<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>ClusterRoleBinding列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                  <el-divider></el-divider>
                  <el-table
                      :data="List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">


               <el-table-column label="角色绑定名" >
                 <template #default="scope">
                   <p >{{ scope.row.metadata.name }}</p>
                   <p>
                     <el-tag style="margin:5px 10px"  @close="rmUserFromBinding(scope.row.metadata.name,sub.kind,sub.name,sub.namespace)"  closable :type="(sub.kind=='User')?'info':(sub.kind=='ServiceAccount')? 'danger':'warning'" v-for="sub in scope.row.subjects">
                       <span class="role"> 角色:</span> {{ sub.name }}
                       <span class="role"> 绑定名:</span> {{ scope.row.metadata.name }}
                       <template v-if="sub.namespace">
                         <span class="role"> 命名空间:</span> {{ sub.namespace }}
                       </template>
                       <span class="role"> 类型:</span> {{ sub.kind }}
                     </el-tag>
                   </p>

                 </template>
               </el-table-column>
                <el-table-column label="增加角色" width="520" >
                  <template #default="scope">
                    <i  v-show="addConfig.index==-1||addConfig.index!==scope.$index" class=" el-icon-circle-plus" @click="()=>showAdd(scope.$index,scope.row.metadata.name)" ><span style="padding:0 5px">增加</span></i>
                    <div v-show="addConfig.index===scope.$index" >
                      <el-form :inline="true"  :model="addConfig" ref="formRef" :rules="rules">
                        <el-form-item  prop="kind"  label=" "  label-width="10px">
                          <el-select size="small"   @change="clear"  v-model="addConfig.kind" style="width:100px;height:30px!important;">
                            <el-option label="User" value="User"/>
                            <el-option label="ServiceAccount" value="ServiceAccount"/>
                            <el-option label="Group"  value="Group" >
                            </el-option>
                          </el-select>
                        </el-form-item>
                        <el-form-item size="small"  label=" "  label-width="10px" prop="ns" v-show="addConfig.kind=='ServiceAccount'">
                          <el-select  style="width:100px;height:30px!important;" placeholder="选择命名空间" @change="changeNs" v-model="addConfig.ns">
                            <el-option v-for="ns in nsList "
                                       :label="ns.name"
                                       :value="ns.name"/>
                          </el-select>
                        </el-form-item>
                        <el-form-item label=" "  label-width="10px" prop="name" v-show="addConfig.kind=='ServiceAccount'">
                          <el-select  size="small" v-model="addConfig.name"  v>
                            <el-option v-for="sa in SaList "
                                       :label="sa.name"
                                       :value="sa.name"/>
                          </el-select>
                        </el-form-item>
                        <el-form-item v-show="addConfig.kind!='ServiceAccount'"   label=" "  label-width="10px"  prop="name" >
                          <el-input  size="small" placeholder="输入用户名"   style="width:200px;height:30px!important;" v-model="addConfig.name" ></el-input>
                        </el-form-item>
                      <div style="margin-top:10px">
                            <i style="margin:0 8px"  class="ii el-icon-s-claim" @click="appendUserToBinding" > 保存</i>
                            <i style="margin:0 8px"   class="ii el-icon-remove" @click="()=>{addConfig.index=-1;addConfig.kind='User';addConfig.name=''}"> 取消</i>
                      </div>
                      </el-form>
                    </div>
                  </template>
               </el-table-column>
               <el-table-column label="创建时间" width="160" align="center">
                 <template #default="scope">
                   <span>{{

                       timeFormat(scope.row.metadata.creationTimestamp)
                     }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="操作" width="100" align="center">
                 <template #default="scope">
                   <el-button type="danger" @click="()=>rm(scope.row.name )" icon="el-icon-delete" circle></el-button>
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
  addClusterRoleBindings,
  addUserToBinding, ClusterRoleBindingList, deleteClusterRoleBinding,
  deleteUserFromBinding,
  getRoleBindingList,
  getRoleList,
  roleBindingDel,
  roleDel
} from "../api/token/rbac";
import {getSaList} from "../api/token/sa/sa";
import {core} from "../core/core";

export default defineComponent({
  name: 'secret-list',
  components: {MainLayout},
  setup(){
    const apiGroup='rbac.authorization.k8s.io'
    const ws = inject("ws")
    let state=reactive({
      List:reactive([]),
      nsList:reactive([]),
      SaList:reactive([]),
      pageInfo:0,
      current_page:1,
      addConfig:reactive({
        index:-1,
        kind:'User',
        name:'',//用户名
        bindingName:'', //binding名称
        ns:'default' //命名空间
      })
    })

    const rules= {
      kind:[
        {
          required: true,
          message: '角色类型必须选择',
        }
      ],
      ns:[
        {
          required: true,
          message: '命名空间必须填写',
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
        let tData2=await getList()
        state.nsList=tData2.data.data
        let tData=await ClusterRoleBindingList(state.current_page)
        let T=tData.data.data
        state.List=T.Data
        state.pageInfo=T["Total"]
        state.current_page=T["Current"]
      }catch (e){
        console.log(e)
      }
    }
    getData()
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        let T=tData.result;
        if(tData.type=='clusterRoleBinding'){
          state.List=T.Data
          state.pageInfo=T["Total"]
          state.current_page=T["Current"]
        }
      }
    }

    function pageChange(page,ns){
      state.current_page=page
      getData()
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
    function rmUserFromBinding(name,kind,username,namespace){
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
              let result=await  deleteClusterRoleBinding(name,{kind,name:username,apiGroup,namespace:namespace})
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

    function showAdd(index,bindingName){
      state.addConfig.index=index
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
            let data=null
            if (kind!=="ServiceAccount"){
              data={kind, name, apiGroup}
            }else{
              data={namespace:ns,kind, name, apiGroup}
            }
            let result = await addClusterRoleBindings( bindingName, data)
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
      changeNs()

    }

      async function changeNs(){
        try {
          let tData2=await getSaList(state.addConfig.ns)
          state.SaList=tData2.data.data
        }catch (e){
          console.log(e)
        }
    }
    let timeFormat=core.timeFormat
    return {...toRefs(state),pageChange,changeNs,rm,doTo,rmUserFromBinding,showAdd,appendUserToBinding,clear,rules,formRef, timeFormat}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>