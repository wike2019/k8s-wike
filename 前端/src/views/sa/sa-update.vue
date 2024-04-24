<template>
   <main-layout >
     <breadcrumb title="服务账号.更新"></breadcrumb>
     <el-form :inline="true"  :model="form" ref="formRef"  >
       <el-button size="small" type="warning" class="btnList" @click="doTo('sa-list')" >进入列表</el-button >
       <el-tabs v-model="state.mode" @tab-click="Update">
         <el-tab-pane label="可视化展示" name="json">
           <metadataInfo tipTitle="ServiceAccount" :nameDisable="true"   @input="getData($event,'metadata',state.form,true)"  ref="metadataInfoRef"  :nameRequired="true"></metadataInfo>
           <secrets tipTitle="imagePullSecrets" :key="'imagePullSecrets'" @input="getData($event,'imagePullSecrets',state.form,true)"  secretsType="imagePullSecrets【拉取私有镜像凭证密文】"  :namespace="state.form.metadata?.namespace"  :value="state.form.imagePullSecrets" ref="imagePullSecretsRef"></secrets>
         </el-tab-pane>
         <el-tab-pane :label="'YAML展示'" name="yaml">
           <yaml ref="yamlRef"  />
         </el-tab-pane>
       </el-tabs>
       <el-divider></el-divider>
       <div  class="submit-box">
         <el-button  type="info" @click="post()">更新</el-button>
       </div>
     </el-form>
   </main-layout>
</template>

<script lang="ts" setup>
import {nextTick, reactive, ref, watch} from 'vue'
import MainLayout from "../../layout/main.vue";

import {doTo} from "../../router";
import {getSaItem,  SAUpdate} from "../../api/token/sa/sa";
import {useRoute} from "vue-router";
import yaml from "../../components/yaml/yaml.vue";
import {requireRules} from "../../helper/rules"
import {CheckData, createTip, getData, } from "../../helper/helper"
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import secrets from "./components/secrets.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import {form} from "../namespace/use";
import {ElMessage} from "element-plus";
let state=reactive({
  mode:"json",
  spec:"secrets",
  form:{
    apiVersion: 'v1',
    Kind: 'ServiceAccount',
    metadata:{
      name:"",
      namespace:""
    },
    imagePullSecrets:{}
  },
  md5:"",
})
const route = useRoute()
let yamlRef = ref(null)
let metadataInfoRef = ref(null)
let imagePullSecretsRef = ref(null)
let checkList=[imagePullSecretsRef,metadataInfoRef]
async function fetchData(){
  try {
    let tData=await getSaItem(route.query.namespace,route.query.name)
    state.form.metadata.name=route.query.name.toString()
    state.form.metadata.namespace=route.query.namespace.toString()
    state.form=Object.assign({},state.form,tData.data.data)
    metadataInfoRef.value.setData(state.form.metadata)
    imagePullSecretsRef.value.setData(state.form.imagePullSecrets)
  }catch (e){
    console.log(e)
  }
}
watch(() => state.form, () => {
  yamlRef.value.setData(state.form)
}, {deep: true, flush: "post"})

function Update(){
  nextTick(()=>{
    yamlRef.value.Update()
  })
}

fetchData()



function post() {
  createTip()
      .then(async () => {
        if(await CheckData(checkList)){
          try {
            let result=await  SAUpdate(state.form)
            if (result.data.code==200){
              ElMessage("修改SA成功")
            }
          }catch (e){
            ElMessage.error(e)
          }
        }
      })
}

</script>
<style >

</style>