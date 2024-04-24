<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务.列表"></breadcrumb>
     <div class="top-list">
       <nsSelect @input="changeNs" url="svc-create" btn="创建配置"></nsSelect>
                  <el-divider></el-divider>
                  <el-table
                      :data="state.List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">

               <el-table-column label="svc名称" >
                 <template #default="scope">
                   <p >{{ scope.row.metadata.name }}</p>
                 </template>
               </el-table-column>
               <el-table-column label="命名空间" width="180" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.metadata.namespace }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="创建时间" width="200" align="center">
                 <template #default="scope">
                   <span>{{ scope.row.metadata.creationTimestamp }}</span>
                 </template>
               </el-table-column>
               <el-table-column label="操作" width="300" align="center">
                 <template #default="scope">
                   <el-button
                       type="primary"
                       size="small"
                       @click="()=>doTo('svc-update',{namespace:scope.row.metadata.namespace,name:scope.row.metadata.name})"  >
                     编辑
                   </el-button>
                     <el-button
                         type="info"
                         size="small"
                         @click="()=>doTo('svc-detail',{namespace:scope.row.metadata.namespace,name:scope.row.metadata.name})"  >
                       查看详情
                     </el-button>
                     <el-button type="danger" size="small" @click="()=>rm(scope.row.metadata.namespace,scope.row.metadata.name )" >删除</el-button>
                 </template>
               </el-table-column>
    </el-table>
       <div class="page_show">
         <el-pagination
             @current-change="pageChange"
             background
             :current-page="state.currentPage"
             :page-size="10"
             :hide-on-single-page="true"
             layout="prev, pager, next"
             :total="state.pageTotal">
         </el-pagination>
       </div>
     </div>
   </main-layout>

</template>

<script lang="ts" setup>
import {defineComponent, inject, onMounted, reactive, ref, toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {doTo} from "../../router";
import nsSelect from "../../components/common/nsSelect.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";

import {getSvcListByPage, removeSvc} from "../../api/token/svc/svc";
import {copyData, wsCopyData} from "../../helper/helper";

const ws: WebSocket= inject("ws")
let state=reactive({
  nsList:reactive([]),
  List:reactive([]),
  pageTotal:0,
  currentPage:1,
  namespace:"default"
})
let formRef=ref(null)
async function getData(){
  try {
    await  changeNs(state.namespace);
  }catch (e){
    console.log(e)
  }
}
async function changeNs(ns){
    state.namespace= ns
    try {
      let tData=await getSvcListByPage(state.namespace,state.currentPage)
      copyData(state, tData)
    }catch (e){
      console.log(e)
    }

}
getData()
ws.onmessage = (e)=>{
  if(e.data !== 'ping'){
    let tData=JSON.parse(e.data)
    wsCopyData(state, tData, "service")
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
</script>
