<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务.详情"></breadcrumb>
     <div class="list_item">
       <metadataInfo  ref="mateDataRef" tipTitle="Services" :nameDisable="true" :nameRequired="true"></metadataInfo>
       <el-divider></el-divider>
       <div>
         <port ref="portRef" :type="state.form.spec.type"></port>
         <labelValue ref="selectorRef" label="Selector设置" ></labelValue>
       </div>
       <el-divider></el-divider>
       <el-card class="box-card">
         <el-descriptions
             :column="1"
             title="其他数据"
         >
           <el-descriptions-item label="集群ip">{{state.form.spec.clusterIP}}</el-descriptions-item>
           <el-descriptions-item label="服务类型">{{state.form.spec.type}}</el-descriptions-item>
           <el-descriptions-item label="clusterIPs列表">{{state.form.spec.clusterIPs}}</el-descriptions-item>
           <el-descriptions-item label="节点亲和性" >{{state.form.spec.sessionAffinity}}</el-descriptions-item>
           <el-descriptions-item label="服务发现策略" >{{state.form.spec.internalTrafficPolicy}}</el-descriptions-item>
         </el-descriptions>
       </el-card>
       <el-button type="primary" @click="()=>doTo('svc-update',{namespace:state.namespace,name:state.name})" >
         编辑
       </el-button>
     </div>
   </main-layout>
</template>

<script lang="ts" setup>
import {defineComponent, inject, provide, reactive, ref, toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {doTo} from "../../router";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {useRoute} from "vue-router";
import {getSvcItem} from "../../api/token/svc/svc";
import Port from "../../components/Port/port.vue";
import LabelValue from "../../components/labelValue/labelValue.vue";
let state=reactive({
  form:{
    metadata:{},
    spec:{
      ports:[],
      selector:{}
    }
  },
  name:"",
  namespace:""
})
const route = useRoute()
let mateDataRef=ref(null)
let portRef=ref(null)
let selectorRef=ref(null)
state.name=route.query.name.toString()
state.namespace=route.query.namespace.toString()
async function getData(){
  try {
    let tData=await getSvcItem(state.namespace,state.name)
    state.form=tData.data.data
    mateDataRef.value.setData( state.form.metadata)
    portRef.value.setData(state.form.spec.ports)
    selectorRef.value.setData(state.form.spec.selector)
  }catch (e){
    console.log(e)
  }
}
getData()

provide("render",true)
</script>
<style >

</style>