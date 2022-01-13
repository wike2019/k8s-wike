<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务账号.列表"></breadcrumb>
     <div class="top-list">
      <nsSelect @input="(e)=>{changeNs(e)}" url="sa-create" btn="创建ServiceAccount"></nsSelect>
      <el-divider></el-divider>
       <el-table
              :data="List"
              border
              empty-text="暂无数据"
              style="width: 100%">

       <el-table-column label="服务账号名称" >
         <template #default="scope">
           <p  @click="()=>doTo('sa-detail',rowToQuery(scope.row))"  >
             {{ scope.row.name }}
             <a style="margin-left:25px" @click.stop="showToken(scope.row.secrets[0].name)">
               查看token
             </a>
           </p>
         </template>
       </el-table-column>
       <el-table-column label="命名空间" width="180" align="center">
         <template #default="scope">
           <span>{{ scope.row.namespace }}</span>
         </template>
       </el-table-column>
       <el-table-column label="创建时间" width="200" align="center">
         <template #default="scope">
           <span>{{ scope.row.createTime }}</span>
         </template>
       </el-table-column>
       <el-table-column label="操作" width="300" align="center">
         <template #default="scope">
           <el-button
               type="primary"
               size="small"
               @click="()=>doTo('sa-update',rowToQuery(scope.row))"  >
             编辑
           </el-button>
             <el-button
                 type="info"
                 size="small"
                 @click="()=>doTo('sa-detail',rowToQuery(scope.row))"  >
               查看详情
             </el-button>
             <el-button type="danger" size="small" @click="()=>rm(scope.row.namespace,scope.row.name )" >删除</el-button>
         </template>
       </el-table-column>
    </el-table>
       <div class="page_show">
         <el-pagination
             @current-change="pageChange"
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

  <el-dialog
      v-model="dialogVisible"
      v-if="dialogVisible"
      title="token"
  >
    <div style="line-height:1.5">{{token}}</div>
    <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">关闭</el-button>
        </span>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import {defineComponent, inject, reactive, ref, toRefs} from 'vue'
import MainLayout from "@/layout/main.vue";
import {doTo} from "@/router";
import {getSaList, SaDel} from "@/api/token/sa/sa";
import {secretDetail} from "@/api/token/secret/secret";
import nsSelect from "@/components/common/nsSelect.vue";
import breadcrumb from "@/components/common/breadcrumb.vue";
import {rowToQuery,copyData,wsCopyData,rmTip} from "@/helper/helper";

export default defineComponent({
  name: 'sa-list',
  components: {MainLayout,nsSelect,breadcrumb},
  setup(){
    const ws = inject("ws")
    let state=reactive({
      List:[],
      namespace:"default",
      dialogVisible:false,
      token:'',
      pageInfo:0,
      current_page:1,
    })
    let formRef=ref(null)

    async function changeNs(ns){
      state.namespace=ns
      try {
          let tData=await getSaList(state.namespace,state.current_page)
          copyData(state,tData)
        }catch (e){
          console.log(e)
        }
    }
    changeNs(state.namespace)
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        wsCopyData(state,tData,"sa")
      }
    }

    function rm(ns,name){
        rmTip('ServiceAccount')
          .then(async () => {
              try {
                await SaDel(ns,name)
              }catch (e) {
                console.log(e)
              }
          })
    }
    async function showToken(name){
      try {
      let result=await secretDetail(state.namespace,name)
          let token=window.atob(result.data.data.data['token'])
          state.token=token;
          state.dialogVisible=true;
      }catch (e) {
        console.log(e)
      }

    }
    function pageChange(page){
      state.current_page=page
      changeNs(state.namespace)
    }

    return {...toRefs(state),rm,changeNs,doTo,showToken,pageChange,rowToQuery}
  }
})
</script>
