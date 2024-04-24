<template>
   <main-layout class="fix-black">
     <breadcrumb title="密文.详情"></breadcrumb>
     <div class="list_item">
       <metadataInfo  ref="mateDataRef" tipTitle="Secret" :nameDisable="true" :nameRequired="true"></metadataInfo>
       <el-divider></el-divider>
       <div v-if="state.secretData.type=='TLS凭据'" >
           <p v-for="(data,item) in state.secretData.tls_data">
             <span class="key_format"  style="padding-left:15px">{{item}}:</span>
             <span class="value_format">{{data}}</span>
           </p>
           <el-divider></el-divider>
       </div>
       <h2 class="sub_title">配置数据Data</h2>
       <el-table
           :data="data"
           border
           empty-text="暂无数据"
           style="width: 100%">

         <el-table-column label="密文key" width="200" align="center">
           <template #default="scope">
              {{ scope.row.key }}
           </template>
         </el-table-column>
         <el-table-column label="密文value(长按显示解密后的值)"  align="left"  >
           <template #default="scope">
             <div @mousedown="()=>scope.row.show=true" @mouseup="()=>scope.row.show=false" >
               <div v-if="!scope.row.show" class="break-all">{{ scope.row.value }}</div>
               <div v-else class="break-all" >{{ scope.row.raw }}</div>
             </div>
           </template>
         </el-table-column>
       </el-table>
       <el-divider></el-divider>
       <el-button class="mtb20" type="primary" @click="()=>doTo('secret-update',{namespace:state.namespace,name:state.name})" >
         编辑
       </el-button>
     </div>
   </main-layout>
</template>

<script lang="ts" setup>
import {computed, provide, reactive, ref} from 'vue'
import breadcrumb from "../../components/common/breadcrumb.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import MainLayout from "../../layout/main.vue";
import { secretDetail} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import {doTo} from "../../router";

import {ElMessage} from "element-plus";


let mateDataRef=ref(null)

let state=reactive({
  namespace:"",
  name:"",
  secretData:{Secret:{metadata:{},data:{}}},
})
let data=computed(()=>{
  let arr=[];
  for ( let i in state.secretData?.Secret.data){
    arr.push({"show":false,"key":i,value:state.secretData.Secret.data[i],raw:decodeURIComponent(atob(state.secretData.Secret.data[i]))})
  }
  return arr
})
provide("render",true)
async function getData(){
  const route = useRoute()
  try {
    state.namespace=route.query.namespace.toString()
    state.name=route.query.name.toString()
    let tData=await secretDetail(state.namespace,state.name)
    state.secretData=tData.data.data
    mateDataRef.value.setData(state.secretData.Secret.metadata)
  }catch (e){
    ElMessage.error(e)
  }
}
getData()
</script>
