<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>pod列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
       <h3 class="sum">统计:  总pod数: <span style="color:#409EFF">{{total_pod}}</span>,就绪pod数 <span style="color:#67C23A">{{ok_pod}}</span></h3>
       <el-divider class="clear_both"></el-divider>
       <div v-for="item in nsList" :key="item.name">
         <el-card class="box-card">
           <template #header>
             <div class="card-header">
               命名空间：<span>{{item.name}}</span>
             </div>
           </template>
                  <el-table
                      :data="podList[item.name]"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">
                    <el-table-column
                        prop="status"
                        label="状态"
                        width="80">
                      <template #default="scope">
                        <span  class="Success"  v-if="scope.row.is_ready">active</span>
                        <span  class="Danger" v-else>waiting</span>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="name"
                        label="pod名称"
                        width="350">
                      <template #default="scope">
                        <p>{{scope.row.name}}</p>
                        <p  class="Danger"  v-if="!scope.row.is_ready">{{scope.row.message}}</p>
                      </template>
                    </el-table-column>

                    <el-table-column
                        prop="images"
                        label="镜像名称">
                      <template #default="scope">
                        <p>{{scope.row.images}}</p>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="create_time"
                        label="创建时间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column label="操作" width="240" align="center">
                      <template #default="scope">
                        <el-button size="small" @click="()=>doTo('pod-logs',{name_space:scope.row.name_space,name:scope.row.name})" >查看日志</el-button>
                        <el-button size="small" @click="()=>doTo('pod-shell',{name_space:scope.row.name_space,name:scope.row.name})" >执行shell</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event,item.name)"
                      background
                      :current-page="current_page[item.name]"
                      :page-size="5"
                      :hide-on-single-page="true"
                      layout="prev, pager, next"
                      :total="pageInfo[item.name]">
                  </el-pagination>
                </div>

         </el-card>
       </div>

     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {getPodsByNs} from "../api/token/pod";
import {doTo} from "../router";
export default defineComponent({
  name: 'pod',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      podList:reactive({}),
      pageInfo:reactive({}),
      current_page:reactive({}),
      other:reactive({})
    })

    async function getData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
        for (let i=0;i<state.nsList.length;i++){
          await  getPods(state.nsList[i].name)
        }
      }catch (e){
        console.log(e)
      }
    }

    async function getPods(ns,page=1){
      try {
        let tData=await getPodsByNs(ns,page)
        let T=tData.data.data
        state.podList[T.Ext.ns]=T.Data
        state.pageInfo[T.Ext.ns]=T["Total"]
        state.current_page[T.Ext.ns]=T["Current"]
        state.other[T.Ext.ns]=T.Ext
      }catch (e){
        console.log(e)
      }
    }

    getData()
    let total_pod=computed(()=>{
      let count=0
      for (let i=0;i<state.nsList.length;i++){
       count+=state.other[state.nsList[i].name]?.AllNum
      }
      return count
    })
    let ok_pod=computed(()=>{
      let count=0
      for (let i=0;i<state.nsList.length;i++){
        count+=state.other[state.nsList[i].name]?.ReadyNum
      }
      return count
    })
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)

        if(tData.type=='pod'){
          let T=tData.result;
          state.podList[T.Ext.ns]=T.Data
          state.pageInfo[T.Ext.ns]=T["Total"]
          state.current_page[T.Ext.ns]=T["Current"]
          state.other[T.Ext.ns]=T.Ext
        }
        if(tData.type=='namespace'){
          getData()
        }

      }
    }

    function pageChange(page,ns){
      state.current_page[ns]=page
      getPods(ns,page)
    }
    return {...toRefs(state),total_pod,ok_pod,pageChange,doTo}
  }
})
</script>
