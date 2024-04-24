<template>
  <main-layout>
    <breadcrumb title="配置.创建"></breadcrumb>
    <el-form ref="formRef" :model="state.form"   class="mtb20">
      <el-button size="small" type="warning" class="btnList" @click="doTo('configmap-list')" >进入列表</el-button >

      <el-tabs v-model="state.mode" @tab-click="Update">
        <el-tab-pane label="可视化展示" name="json">
          <metadataInfo ref="mateDataRef" @input="getData($event,'metadata',state.form,true)"  tipTitle="ConfigMap" type="ConfigMap" :nameRequired="true"></metadataInfo>
          <KeyValue ref="KeyValueRef" :base64="false"  @input="getData($event,'data',state.form,true)" ></KeyValue>
        </el-tab-pane>
        <el-tab-pane label="YAML展示" name="yaml">
          <yaml ref="yamlRef"  />
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div style="text-align: center;margin-top: 20px">
      <el-button type="primary" @click="postNew()">保存</el-button>
    </div>
  </main-layout>
</template>

<script lang="ts" setup>
import { ref, reactive,nextTick, watch, onMounted} from 'vue'
import MainLayout from "../../layout/main.vue";

import  {ElMessage} from 'element-plus'
import KeyValue from "../../components/key_value/KeyValue.vue";
import {configmapCreate} from "../../api/token/configmap/configmap";

import {getData} from "../../helper/helper"

import yaml from "../../components/yaml/yaml.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {requireRules} from "../../helper/rules";
import breadcrumb from "../../components/common/breadcrumb.vue";
import {doTo} from "../../router";

let state=reactive({
  mode:"json",
  form:{
    apiVersion:'v1',
    Kind:'ConfigMap',
    metadata:{
      name:"",
      namespace:""
    },
    data:{}
  },
})
//公共部分
let KeyValueRef=ref(null)
let mateDataRef=ref(null)
let yamlRef=ref(null)
function Update(){
  nextTick(()=>{
    yamlRef.value.Update()
  })
}



async function postNew(){

  let flag=await KeyValueRef.value.Check()&&await mateDataRef.value.Check()
  if(!flag){
    ElMessage.error("数据不合法，有必选项未填")
    return
  }
  try {
    let result=await  configmapCreate(state.form)
    if (result.data.code==200){
      ElMessage("配置创建成功")
    }
  }catch (e){
    ElMessage.error(e)
  }
}


watch(()=>state.form,()=>{
  yamlRef.value.setData(state.form)
},{deep:true,flush:"post"})

</script>
