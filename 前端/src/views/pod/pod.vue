<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>pod列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
       <div class="name_space_choose">
         <span>命名空间:</span>
         <el-select placeholder="选择命名空间" @change="changeNs" filterable v-model="namespace">
           <el-option v-for="ns in state.nsList "
                      :label="ns.metadata.name"
                      :value="ns.metadata.name"/>
         </el-select>
         <el-button style="margin-left:40px" type="primary" @click="doTo('pod-create')">创建Pod</el-button>
       </div>
       <el-divider class="clear_both"></el-divider>
         <el-card class="box-card">
           <template #header>
             <div class="card-header">
               命名空间：<span>{{namespace}}</span>
             </div>
           </template>
                  <el-table
                      :data="state.podList"
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
                        <p>{{scope.row.Pod.metadata.name}}</p>
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
                        prop="Pod.metadata.creationTimestamp"
                        label="创建时间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column label="操作" width="320" align="center">
                      <template #default="scope">
                        <el-button size="small" @click="()=>doTo('pod-logs',{name_space:scope.row.name_space,name:scope.row.name})" >查看日志</el-button>
                        <el-button size="small" @click="()=>doTo('pod-shell',{name_space:scope.row.name_space,name:scope.row.name})" >执行shell</el-button>
                        <el-button size="small" type="danger" @click="()=>doTo('pod-shell',{name_space:scope.row.name_space,name:scope.row.name})" >删除</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange"
                      background
                      :current-page="current_page"
                      :page-size="5"
                      :hide-on-single-page="true"
                      layout="prev, pager, next"
                      :total="state.total">
                  </el-pagination>
                </div>

         </el-card>
       </div>
   </main-layout>
</template>

<script lang="ts" setup>
import {computed, inject, reactive, ref} from 'vue'
import MainLayout from "../../layout/main.vue";
import {getNsALL} from "../../api/token/namespace/ns";
import {getPodsByNs} from "../../api/token/pod";
import {doTo} from "../../router";
const ws:WebSocket  = inject("ws")
let state=reactive({"nsList":[],total:0,podList:[]})
let namespace=ref("")
let current_page=ref(1)
async function getData(){
  try {
    let tData=await getNsALL()
    state.nsList=tData.data
    if (state.nsList.length>0){
      namespace.value="kube-system"
     await getPods(namespace.value,1)
    }
  }catch (e){
    console.log(e)
  }
}

async function getPods(ns,page=1){
  try {
    let tData=await getPodsByNs(ns,page)
    let T=tData.data.data
    console.log(T)
    state.podList=T.Data
    state.total=T.Total
  }catch (e){
    console.log(e)
  }
}

getData()

ws.onmessage = (e)=>{
  if(e.data !== 'ping'){
    let tData=JSON.parse(e.data)

    if(tData.type=='pod'){
      let T=tData.result;
      if (T.Ext.ns==namespace.value){
        state.podList=T.Data
        state.total=T.Total
        current_page.value=1
      }

    }
    if(tData.type=='namespace'){
      getData()
    }

  }
}
function changeNs() {
  current_page.value=1
  getPods(namespace.value,1)
}
function pageChange(page){
  current_page.value=page
  getPods(namespace.value,page)
}
</script>
