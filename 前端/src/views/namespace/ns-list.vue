<template>
   <main-layout>
     <breadcrumb title="命名空间.列表"></breadcrumb>
     <el-button type="primary" class="right mtb20" @click="dialogFormVisible=true" >创建命名空间</el-button>
     <el-divider class="clear_both"></el-divider>
     <div class="top-list">
       <el-table :data="state.List">
           <el-table-column label="命名空间名称">
             <template #default="scope">
               <span>{{ scope.row.metadata.name }}<em class="Danger mf20">{{ scope.row.metadata.name=='wike-system'?"系统命名空间请不要删除":""}}</em></span>
             </template>
           </el-table-column>
           <el-table-column align="center" label="标签" width="420">
             <template #default="scope">
               <p v-for="(value,key) in scope.row.metadata.labels" :key="scope.row.metadata.name">{{key}}={{value}}</p>
             </template>
           </el-table-column>
           <el-table-column align="center" prop="metadata.creationTimestamp" label="创建时间" width="220"></el-table-column>
           <el-table-column label="操作" width="120" align="center">
           <template #default="scope">
             <el-button size="small" type="danger"   @click="()=>doDelete(scope.row.metadata.name)" >删除</el-button>
           </template>
         </el-table-column>
       </el-table>
       <div class="page_show">
         <el-pagination
             @current-change="fetchData"
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
  <el-dialog v-model="dialogFormVisible" title="创建命名空间">
    <el-form :model="form" ref="formRef">
      <el-form-item label="命名空间名称" prop="metadata.name"    :rules="requireRules('命名空间名称必填')">
        <el-input v-model="form.metadata.name" ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="create">创建</el-button>
        <el-button @click="()=>{formRef.resetFields();dialogFormVisible = false}">取消</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>

import {inject, reactive} from 'vue'
import MainLayout from "@/layout/main.vue";  //主界面
import {createNs, deleteNs, getNsList} from "../../api/token/namespace/ns";
import {requireRules} from "../../helper/rules";
import {ElMessage} from "element-plus";
import breadcrumb from "@/components/common/breadcrumb.vue" //头部面包屑
import {copyData, rmTip, wsCopyData} from "../../helper/helper";

const ws: WebSocket = inject("ws") //通信用的
import {dialogFormVisible, formRef, form} from "./use"
let state = reactive({
  List: [],
  pageTotal:0,
  currentPage:1,
})

async function fetchData(i) {
  try {
    state.currentPage=i
    let tData = await getNsList(state.currentPage)  //获取数据
    copyData(state, tData)
  } catch (e) {
    ElMessage.error(e)
  }
}

fetchData(1)


ws.onmessage = (e) => {
  if (e.data !== 'ping') {
    let tData = JSON.parse(e.data)
    wsCopyData(state, tData, "namespace")
  }
}

function create() {
  formRef.value.validate(async (valid) => { //验证数据
    if (valid) {
      try {
        let result = await createNs(form)  //提交数据
        if (result.data.code == 200) {
          ElMessage("资源创建成功")
          form.metadata.name = ""
          dialogFormVisible.value = false
          formRef.value.resetFields();
        }
      } catch (e) {
        ElMessage.error(e)
      }
    }
  })
}

function doDelete(name) {
  rmTip("命名空间")
      .then(async () => {
        try {
          if (name == "wike-system") {
            ElMessage.error("此命名空间不推荐删除")
            return
          }
          await deleteNs(name)
          ElMessage.success("资源删除成功")
        } catch (e) {
          ElMessage.error(e)
        }
      })
}

</script>
