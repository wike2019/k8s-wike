<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>Role列表</el-breadcrumb-item>
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


               <el-table-column label="角色名" width="280">
                 <template #default="scope">
                   <p @click="()=>doTo('rbac-role-update',{name_space:scope.row.name_space,name:scope.row.name})">{{ scope.row.name }}</p>
                 </template>
               </el-table-column>
               <el-table-column label="命名空间" width="180" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.name_space }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="创建时间" width="200" align="center">
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
import {getRoleList, roleDel} from "../api/token/rbac";

export default defineComponent({
  name: 'secret-list',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      List:reactive([]),
      pageInfo:0,
      current_page:1,
      name_space:"default"
    })

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
          let tData=await getRoleList(state.name_space,state.current_page)
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
        if(tData.type=='role'&&tData.ns==state.name_space){

          state.List=T.Data
          state.pageInfo=T["Total"]
          state.current_page=T["Current"]
        }
        if(tData.type=='namespace'){
          getData()
        }

      }
    }

    function pageChange(page,ns){
      state.current_page=page
      changeNs()
    }
    function rm(ns,name){
      ElMessageBox.confirm(
          '你确定继续删除这个角色操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                await roleDel(ns,name)
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }

    return {...toRefs(state),pageChange,rm,changeNs,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>