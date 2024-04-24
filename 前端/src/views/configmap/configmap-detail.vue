<template>
   <main-layout class="fix-black">
     <breadcrumb title="配置.详情"></breadcrumb>
     <div class="list_item">
       <metadataInfo  ref="mateDataRef" tipTitle="ConfigMap" :nameDisable="true" :nameRequired="true"></metadataInfo>
       <el-divider></el-divider>
       <el-table
           :data="data"
           border
           empty-text="暂无数据"
           style="width: 100%">

         <el-table-column label="配置key" width="200" align="center">
           <template #default="scope">
              {{ scope.row.key }}
           </template>
         </el-table-column>
         <el-table-column label="配置value"   align="left">
           <template #default="scope">
               <div class="break-all"  v-if="!scope.row.show" style="word-wrap: break-word;">{{ scope.row.value }}
               </div>
           </template>
         </el-table-column>
       </el-table>
       <el-divider></el-divider>
         <el-button class="mtb20" type="primary" @click="()=>doTo('configmap-update',{namespace:state.namespace,name:state.name})" >
           编辑
         </el-button>
       </div>
   </main-layout>
</template>

<script lang="ts" setup>
import { computed, ref,  reactive,provide} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {useRoute} from "vue-router";
import {configmapDetail} from "../../api/token/configmap/configmap";
import {doTo} from "../../router";
let mateDataRef=ref(null)
const route = useRoute()
let state=reactive({
  namespace:"",name:"",
  form:reactive({data:{},metadata:{}}),
})
state.name=route.query.name.toString()
state.namespace=route.query.namespace.toString()
let data=computed(()=>{
  let arr=[];
  for ( let i in state.form.data){
    arr.push({"key":i,value:state.form.data[i]})
  }
  return arr
})
async function getData(){

  try {
    let tData=await configmapDetail(state.namespace,state.name)

    state.form=tData.data.data
    mateDataRef.value.setData( state.form.metadata)
  }catch (e){
    console.log(e)
  }
}
getData()
provide("render",true)

</script>
