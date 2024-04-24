<template>
   <main-layout class="fix-black">
     <breadcrumb title="密文.列表"></breadcrumb>
     <div class="top-list">
       <nsSelect @input="(e)=>{changeNs(e)}" url="secret-create" btn="创建Secret"></nsSelect>
       <el-divider></el-divider>
       <el-table
           :data="state.List"
           border
           empty-text="暂无数据"
           style="width: 100%">
         <el-table-column label="名称" >
           <template #default="scope">
             <p @click="()=>doTo('secret-detail',{namespace:scope.row.Secret.metadata.namespace,name:scope.row.Secret.metadata.name})">{{ scope.row.Secret.metadata.name }}</p>
           </template>
         </el-table-column>
         <el-table-column label="类型" width="180" align="center">
           <template #default="scope">
             <span>{{ scope.row.type||"未知类型" }}</span>
           </template>
         </el-table-column>
         <el-table-column label="命名空间" width="180" align="center">
           <template #default="scope">
             <span>{{ scope.row.Secret.metadata.namespace }}</span>
           </template>
         </el-table-column>
         <el-table-column label="创建时间" width="200" align="center">
           <template #default="scope">
             <span>{{ scope.row.Secret.metadata.creationTimestamp }}</span>
           </template>
         </el-table-column>
         <el-table-column label="操作" width="300" align="center">
           <template #default="scope">
             <el-button
                 type="primary"
                 size="small"
                 @click="()=>doTo('secret-update',{namespace:scope.row.Secret.metadata.namespace,name:scope.row.Secret.metadata.name})"  >
               编辑
             </el-button>
             <el-button
                 type="info"
                 size="small"
                 @click="()=>doTo('secret-detail',{namespace:scope.row.Secret.metadata.namespace,name:scope.row.Secret.metadata.name})"  >
               查看详情
             </el-button>
             <el-button type="danger" size="small" @click="()=>rm(scope.row.Secret.metadata.namespace,scope.row.Secret.metadata.name ,scope.row.type )" >删除</el-button>
           </template>
         </el-table-column>
    </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event)"
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
import {inject, reactive} from 'vue'
import { ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import { getSecretList, secretDel,} from "../../api/token/secret/secret";
import {doTo} from "../../router";
import {copyData, rmTip, wsCopyData} from "../../helper/helper";
import nsSelect from "../../components/common/nsSelect.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";

const ws: WebSocket = inject("ws")
let state = reactive({
  List: [],
  pageTotal:0,
  currentPage:1,
  namespace: "default"
})

changeNs(state.namespace)

async function changeNs(ns) {
  state.namespace = ns
  try {
    let tData = await getSecretList(state.namespace, state.currentPage)
    copyData(state, tData)
  } catch (e) {
    ElMessage.error(e)
  }
}

ws.onmessage = (e) => {
  if (e.data !== 'ping') {
    let tData = JSON.parse(e.data)
    wsCopyData(state, tData, "secret")
  }
}


function pageChange(page) {
  state.currentPage = page
  changeNs(state.namespace)
}

function rm(ns, name, type) {
  rmTip('Secret')
      .then(async () => {
        if (type == "服务账号令牌") {
          ElMessage.error('不允许删除服务账号令牌密文')
          return
        }
        try {
          await secretDel(ns, name)
          ElMessage.success("资源删除成功")
        } catch (e) {
          ElMessage.error(e)
        }
      })
}
</script>
