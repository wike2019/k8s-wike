<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>工作负载</el-breadcrumb-item>
             <el-breadcrumb-item>deployment列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div>
       <div v-for="item in nsList" :key="item.name" class="top">
         <el-card class="box-card">
           <template #header>
             <div class="card-header">
               命名空间：<span>{{item.name}}</span>
             </div>
           </template>
           <el-table
               :data="depList[item.name]"
               empty-text="暂无数据"
               border
               style="width: 100%">
             <el-table-column
                 prop="name"
                 label="deployment名称"
                 width="350">
               <template #default="scope">
                 <p>{{scope.row.name}}</p>
                 <p  class="Danger"  v-if="!scope.row.is_complete">{{scope.row.message}}</p>
               </template>
             </el-table-column>
             <el-table-column
               prop="status"
               label="状态"
               width="180">
               <template #default="scope">
                 <span  class="Success"  v-if="scope.row.is_complete">active</span>
                 <span  class="Danger" v-else>waiting</span>
               </template>
           </el-table-column>
             <el-table-column
                 prop="name_space"
                 label="命名空间"
                 width="180">
             </el-table-column>
             <el-table-column
                 prop="images"
                 label="镜像名称">
               <template #default="scope">
                 <p>{{scope.row.images}}</p>
                 <p class="replicas"><span>{{scope.row.replicas[0]}}</span>/<span class="Success">{{scope.row.replicas[1]}}</span>/<span class="Danger">{{scope.row.replicas[2]}}</span></p>
               </template>
             </el-table-column>
             <el-table-column
                 prop="create_time"
                 label="创建时间"
                 width="180"
             >
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
import {defineComponent, ref, onUnmounted, reactive, toRefs, inject} from 'vue'
import MainLayout from "../layout/main.vue";
import {getList} from "../api/token/deployment";
import {getNsList} from "../api/token/namespace/ns";
export default defineComponent({
  name: 'dep',
  components: {MainLayout},
  props: {
    ws: {
      type: Object,
    }
  },
  setup(props){
    const ws = inject("ws")

    let state=reactive({
      nsList:reactive([]),
      depList:reactive({}),
      pageInfo:reactive({}),
      current_page:reactive({}),
      other:reactive({})
    })
    async function getData(){
      try {
        let tData=await getNsList()
        state.nsList=tData.data.data
        for (let i=0;i<state.nsList.length;i++){
          try {
            let tData=await  getList(state.nsList[i].name,1)
            let T=tData.data.data
            state.depList[T.Ext.ns]=T.Data
            state.pageInfo[T.Ext.ns]=T["Total"]
            state.current_page[T.Ext.ns]=T["Current"]
            state.other[T.Ext.ns]=T.Ext
          }catch (e){
            console.log(e)
          }
        }
      }catch (e){
        console.log(e)
      }
    }
    getData()

    ws.onmessage = (e)=>{

      if(e.data !== 'ping'){
         let tData=JSON.parse(e.data)
         if(tData.type=='dep'){
           let T=tData.result;
           state.depList[T.Ext.ns]=T.Data
           state.pageInfo[T.Ext.ns]=T["Total"]
           state.current_page[T.Ext.ns]=T["Current"]
           state.other[T.Ext.ns]=T.Ext
         }
          if(tData.type=='namespace'){
            getData()
          }
      }
    }
    async function pageChange  (page,ns){
      state.current_page[ns]=page
      try {
        let tData=await  getList(ns,page)
        let T=tData.data.data
        state.depList[T.Ext.ns]=T.Data
        state.pageInfo[T.Ext.ns]=T["Total"]
        state.current_page[T.Ext.ns]=T["Current"]
        state.other[T.Ext.ns]=T.Ext
      }catch (e){
        console.log(e)
      }
    }
    return {...toRefs(state),pageChange}
  }
})
</script>

