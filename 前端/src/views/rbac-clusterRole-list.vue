<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>ClusterRole列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                  <el-table
                      :data="List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">


               <el-table-column label="集群角色名">
                 <template #default="scope">
                   <p @click="()=>doTo('rbac-clusterRole-update',{name:scope.row.name})">{{ scope.row.name }}</p>
                 </template>
               </el-table-column>
               <el-table-column label="创建时间" width="200" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.create_time }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="操作" width="100" align="center">
                 <template #default="scope">
                   <el-button type="danger" @click="()=>rm(scope.row.name )" >删除</el-button>
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
import {ClusterRoleBindingDel, clusterRolesDel, getClusterRoles, getRoleList, roleDel} from "../api/token/rbac";

export default defineComponent({
  name: 'secret-list',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      List:reactive([]),
      pageInfo:0,
      current_page:1,
    })

    async function getData(){
      try {
        let tData=await getClusterRoles(state.current_page)
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
        if(tData.type=='__ClusterRole__'){
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
          '你确定继续删除这个集群角色绑定操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                await ClusterRoleBindingDel(ns,name)
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }

    return {...toRefs(state),pageChange,rm,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>