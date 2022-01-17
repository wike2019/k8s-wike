<template>
   <main-layout class="fix-black">
     <breadcrumb title="密文.列表"></breadcrumb>
     <div class="top-list">
       <nsSelect @input="(e)=>{changeNs(e)}" url="secret-create" btn="创建Secret"></nsSelect>
       <el-divider></el-divider>
       <el-table
           :data="List"
           border
           empty-text="暂无数据"
           style="width: 100%">
         <el-table-column label="名称" >
           <template #default="scope">
             <p @click="()=>doTo('secret-detail',{namespace:scope.row.namespace,name:scope.row.name})">{{ scope.row.name }}</p>
           </template>
         </el-table-column>
         <el-table-column label="类型" width="180" align="center">
           <template #default="scope">
             <span>{{ scope.row.type||"未知类型" }}</span>
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
         <el-table-column label="操作" width="300" align="center">
           <template #default="scope">
             <el-button
                 type="primary"
                 size="small"
                 @click="()=>doTo('secret-update',{namespace:scope.row.namespace,name:scope.row.name})"  >
               编辑
             </el-button>
             <el-button
                 type="info"
                 size="small"
                 @click="()=>doTo('secret-detail',{namespace:scope.row.namespace,name:scope.row.name})"  >
               查看详情
             </el-button>
             <el-button type="danger" size="small" @click="()=>rm(scope.row.namespace,scope.row.name, scope.row.type )" >删除</el-button>
           </template>
         </el-table-column>
    </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event)"
                      background
                      :current-page="currentPage"
                      :page-size="10"
                      :hide-on-single-page="true"
                      layout="prev, pager, next"
                      :total="pageTotal">
                  </el-pagination>
                </div>

     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent,inject, reactive,toRefs} from 'vue'
import { ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getSecretist, secretDel, secretListByNs} from "../../api/token/secret/secret";
import {doTo} from "../../router";
import {copyData, rmTip, wsCopyData} from "../../helper/helper";
import nsSelect from "../../components/common/nsSelect.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";

export default defineComponent({
  name: 'secret-list',
  components: {
    MainLayout, breadcrumb, nsSelect},
    setup() {

      const ws = inject("ws")
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
          let tData = await getSecretist(state.namespace, state.currentPage)
          copyData(state, tData)
        } catch (e) {
          console.log(e)
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
                ElMessage({
                  type: 'info',
                  message: '不允许删除服务账号令牌密文   '
                })
                return
              }
              try {
                await secretDel(ns, name)
              } catch (e) {
                console.log(e)
              }
            })
      }

      return {...toRefs(state), pageChange, rm, changeNs, doTo}

  }
})
</script>
