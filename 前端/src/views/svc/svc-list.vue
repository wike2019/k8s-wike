<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>svc列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                 <div class="name_space_choose">
                   <span>命名空间:</span>
                   <el-select placeholder="选择命名空间" @change="changeNs" filterable v-model="namespace">
                     <el-option v-for="ns in nsList "
                                :label="ns.name"
                                :value="ns.name"/>
                   </el-select>
                   <el-button style="margin-left:40px" type="primary" @click="doTo('svc-create')">创建Svc</el-button>
                 </div>
                  <el-divider></el-divider>
                  <el-table
                      :data="List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">

               <el-table-column label="svc名称" >
                 <template #default="scope">
                   <p >{{ scope.row.name }}</p>
                 </template>
               </el-table-column>
               <el-table-column label="命名空间" width="180" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.namespace }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="创建时间" width="200" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.create_time }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="操作" width="300" align="center">
                 <template #default="scope">
                   <el-button
                       type="primary"
                       size="small"
                       @click="()=>doTo('svc-update',{namespace:scope.row.namespace,name:scope.row.name})"  >
                     编辑
                   </el-button>
                     <el-button
                         type="info"
                         size="small"
                         @click="()=>doTo('svc-detail',{namespace:scope.row.namespace,name:scope.row.name})"  >
                       查看详情
                     </el-button>
                     <el-button type="danger" size="small" @click="()=>rm(scope.row.namespace,scope.row.name )" >删除</el-button>
                 </template>
               </el-table-column>
    </el-table>
       <div class="page_show">
         <el-pagination
             @current-change="pageChange"
             background
             :current-page="current_page"
             :page-size="10"
             :hide-on-single-page="true"
             layout="prev, pager, next"
             :total="pageInfo">
         </el-pagination>
       </div>
     </div>
   </main-layout>

</template>

<script lang="ts">
import {defineComponent, inject, onMounted, reactive, ref, toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {doTo} from "../../router";
import {getSaList, SACreate, SaDel} from "../../api/token/sa/sa";
import {secretDetail} from "../../api/token/secret/secret";
import {roleCreate} from "../../api/token/rbac";
import {getSvcListByPage, removeSvc} from "../../api/token/svc/svc";

export default defineComponent({
  name: 'svc-list',
  components: {MainLayout},
  setup(){
    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      List:reactive([]),
      namespace:"default",
      pageInfo:0,
      current_page:1,
    })
    let formRef=ref(null)
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
          let tData=await getSvcListByPage(state.namespace,state.current_page)

           let T=tData.data.data
            state.List=T.Data
            state.pageInfo=T.Total
            state.current_page=T.Current
            console.log(tData)
        }catch (e){
          console.log(e)
        }

    }
    getData()
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        let T=tData.result;
        console.log(T)
        if(tData.type=='service'&&tData.ns==state.namespace){
          state.List=T.Data
          state.pageInfo=T.Total
          state.current_page=T.Current
        }
        if(tData.type=='namespace'){
          getData()
        }

      }
    }

    function rm(ns,name){
      ElMessageBox.confirm(
          '你确定继续删除这个service操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                await removeSvc(ns,name)
              }catch (e) {
                console.log(e)
              }
          })
    }

    function pageChange(page){
      state.current_page=page
      changeNs()
    }

    return {...toRefs(state),rm,changeNs,doTo,pageChange}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>