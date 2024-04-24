<template>
   <main-layout class="fix-black">
     <breadcrumb title="ingress.列表"></breadcrumb>
     <nsSelect @input="(e)=>{changeNs(e)}" url="ingress-create" btn="创建配置"></nsSelect>
     <el-divider></el-divider>
                  <el-table
                      :data="state.List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">
                    <el-table-column
                        prop="status"
                        label="状态"
                        width="120">
                      <template #default="scope">
                        <span  class="Success"  v-if="scope.row.status?.loadBalancer?.ingress?.length>0">active</span>
                        <span  class="Danger" v-else>正在初始化</span>
                      </template>
                    </el-table-column>
                    <el-table-column
                        label="名称"
                        width="200px"
                        >
                      <template #default="scope">
                        <p>{{scope.row.metadata.name}}</p>
                      </template>
                    </el-table-column>
                    <el-table-column label="域名">
                      <template #default="scope">
                        <p><a target="_blank" :href="prefix(scope.row)" >{{ prefix(scope.row) }}</a></p>
                      </template>
                    </el-table-column>
                    <el-table-column label="配置" width="140">
                      <template #default="scope">
                       <p>
                         <template v-if="config(scope.row.metadata.annotations,'nginx.ingress.kubernetes.io/enable-cors')">
                           <el-checkbox  :checked="true" disabled>跨域</el-checkbox>
                         </template>
                         <template v-else>
                           <el-checkbox  :checked="false" disabled>跨域</el-checkbox>
                         </template>
                       </p>
                        <p>
                          <template v-if="config(scope.row.metadata.annotations,'nginx.ingress.kubernetes.io/rewrite-target')">
                            <el-checkbox  :checked="true" disabled>重写</el-checkbox>
                          </template>
                          <template v-else>
                            <el-checkbox  :checked="false" disabled>重写</el-checkbox>
                          </template>
                        </p>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="metadata.namespace"
                        label="命名空间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column
                        prop="metadata.creationTimestamp"
                        label="创建时间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column label="操作" width="300" align="center">
                      <template #default="scope">
                        <el-button
                            type="primary"
                            size="small"
                            @click="()=>doTo('ingress-update',{namespace:scope.row.metadata.namespace,name:scope.row.metadata.name})"  >
                          编辑
                        </el-button>
                        <el-button
                            type="info"
                            size="small"
                            @click="()=>doTo('ingress-detail',{namespace:scope.row.metadata.namespace,name:scope.row.metadata.name})"  >
                          查看详情
                        </el-button>
                        <el-button type="danger" size="small" @click="()=>rmIngress(scope.row.metadata.namespace,scope.row.metadata.name )" >删除</el-button>
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
   </main-layout>
</template>

<script lang="ts" setup>
import {inject, reactive} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import nsSelect from "../../components/common/nsSelect.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import {ingressDel, ingressListByNs} from "../../api/token/ingress/ingress";
import {doTo} from "../../router";
import {copyData, wsCopyData} from "../../helper/helper";
const ws:WebSocket = inject("ws")
let state=reactive({
  List:reactive([]),
  pageTotal:0,
  currentPage:1,
  namespace:"default"
})

ingressList()
async function ingressList(page=1){
  try {
    let tData=await ingressListByNs(state.namespace,page)
    copyData(state, tData)
  }catch (e){
    console.log(e)
  }
}

ws.onmessage = (e)=>{
  if(e.data !== 'ping'){
    let tData=JSON.parse(e.data)
    wsCopyData(state, tData, "ingress")
  }
}

function pageChange(page){
  state.currentPage=page
  ingressList(state.currentPage)
}
function rmIngress(ns,name){
  ElMessageBox.confirm(
      '你确定继续删除ingress操作吗?',
      'Warning',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(async () => {
        try {
          await ingressDel(ns,name)
        }catch (e) {
          ElMessage({
            type: 'info',
            message: '系统异常'+e,
          })
        }
      })
}
function prefix(row){
  if (!row.spec?.rules[0]?.host){
    return  "无域名"
  }
  return row.spec.tls?.length>0?"https://"+row.spec?.rules[0]?.host:"http://"+row.spec?.rules[0]?.host;
}
function config(data,match){
  if(data&&data[match]){
    return true
  }
  return false
}

async function changeNs(ns){
  state.namespace=ns
 await ingressList(ns)
}
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>