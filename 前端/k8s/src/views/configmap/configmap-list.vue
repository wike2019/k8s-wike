<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>配置列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                 <div class="name_space_choose">
                   <span>命名空间:</span>
                   <el-select placeholder="选择命名空间" @change="changeNs" v-model="namespace">
                     <el-option v-for="ns in nsList "
                                :label="ns.name"
                                :value="ns.name"/>
                   </el-select>
                   <el-button style="margin-left:40px" type="primary" @click="doTo('configmap-create')">创建Configmap</el-button>
                 </div>
                  <el-divider></el-divider>
                  <el-table
                      :data="configMapList"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">


               <el-table-column label="名称" >
                 <template #default="scope">
                   <p @click="()=>doTo('configmap-detail',{namespace:scope.row.namespace,name:scope.row.name})">{{ scope.row.name }}</p>
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
               <el-table-column label="操作" width="200" align="center">
                 <template #default="scope">
                   <el-button
                       type="info"
                       size="small"
                       @click="()=>doTo('configmap-detail',{namespace:scope.row.namespace,name:scope.row.name})"  >
                     查看详情
                   </el-button>
                   <el-button type="danger" size="small" @click="()=>rmSecret(scope.row.namespace,scope.row.name )" >
                     删除
                   </el-button>
                 </template>
               </el-table-column>
    </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event)"
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
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {doTo} from "../../router";
import {configmapDel, configmapListByNs} from "../../api/token/configmap/configmap";

export default defineComponent({
  name: 'configmap-list',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      configMapList:reactive([]),
      pageInfo:0,
      current_page:1,
      namespace:"default"
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
          let tData=await configmapListByNs(state.namespace,state.current_page)
          let T=tData.data.data
          state.configMapList=T.Data
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
        if(tData.type=='cm'&&tData.ns==state.name_space){

          state.configMapList=T.Data
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
    function rmSecret(ns,name){
      ElMessageBox.confirm(
          '你确定继续删除configmap操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                await configmapDel(ns,name)
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }

    return {...toRefs(state),pageChange,rmSecret,changeNs,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>