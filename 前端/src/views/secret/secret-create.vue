<template>
  <main-layout>
    <breadcrumb title="密文.创建"></breadcrumb>
    <el-button size="small" type="warning" class="btnList" @click="doTo('secret-list')" >进入列表</el-button >
    <el-tabs v-model="state.mode" @tab-click="Update">
      <el-tab-pane label="可视化展示" name="json">
        <metadataInfo ref="mateDataRef" @input="getData($event,'metadata',state.form,true)"  tipTitle="Secret" type="secret" :nameRequired="true"></metadataInfo>
        <el-tabs v-model="state.form.type"  >
          <el-tab-pane label="自定义类型" name="Opaque" >
            <KeyValue ref="KeyValueRef" :base64="true" @input="getData($event,'stringData',state.form,true)" v-if="state.form.type=='Opaque'"></KeyValue>
          </el-tab-pane>
          <el-tab-pane label="TLS类型" name="kubernetes.io/tls" >
            <Tls ref="tlsRef" @input="getData($event,'stringData',state.form,true)" v-if="state.form.type=='kubernetes.io/tls'"></Tls>
          </el-tab-pane>
          <el-tab-pane label="docker镜像拉取密文" name="kubernetes.io/dockercfg" >
            <docker ref="dockerRef" @input="getData($event,'data',state.form,true)" v-if="state.form.type=='kubernetes.io/dockercfg'"></docker>
          </el-tab-pane>
          <el-divider></el-divider>
          <el-alert title="只读属性,可视化不能修改,为了兼容yaml操作," type="info" />
          <h2 class="sub_title">Data数据</h2>
          <div  class="list_item">
            <div v-for=" (value,key) in state.form.data">
              <p><em>key:</em>{{key}}</p>
              <p><em>base64后的值:</em> {{value}}</p>
              <el-divider></el-divider>
            </div>
          </div>
        </el-tabs>
      </el-tab-pane>
      <el-tab-pane label="YAML展示" name="yaml">
        <yaml ref="yamlRef"  />
      </el-tab-pane>
    </el-tabs>
    <div style="text-align: center;margin-top: 20px">
      <el-button type="primary" @click="postNew()">保存</el-button>
    </div>
  </main-layout>
</template>

<script lang="ts" setup>
import {
  ref,
  reactive,
  nextTick,
  watch,
  provide,
} from 'vue'
import MainLayout from "../../layout/main.vue";
import {getNsALL} from "../../api/token/namespace/ns";
import { ElMessage} from 'element-plus'
import KeyValue from "../../components/key_value/KeyValue.vue";
import {secretCreate, secretDetail, secretUpdate} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import {getData} from "../../helper/helper"
import yaml from "../../components/yaml/yaml.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {requireRules} from "../../helper/rules";
import breadcrumb from "../../components/common/breadcrumb.vue";
import Docker from "./components/docker.vue";
import Tls from "./components/tls.vue";
import {doTo} from "../../router";

provide("render",false)
let state=reactive({
  nsList:reactive([]),
  mode:"json",
  form:{
    apiVersion:'v1',
    Kind:'Secret',
    metadata:{
      name:"",
      namespace:"",
      labels:{},
      annotations:{},
    },
    type:"Opaque",
    stringData:{},
    data:{}
  },
  md5:"",
})
const route = useRoute()

let yamlRef=ref(null)
let mateDataRef=ref(null)
let tlsRef=ref(null)
let KeyValueRef=ref(null)
let dockerRef=ref(null)

function Update(){
  nextTick(()=>{
    yamlRef.value.Update()
  })
}




async function fetchData(){
  try {
    let tData=await getNsALL()
    state.nsList=tData.data.data
  }catch (e){
    console.log(e)
  }
}

fetchData()
//
watch(()=>state.form,()=>{
  yamlRef.value.setData(state.form)
},{deep:true,flush:"post"})

async  function postNew(){

  let flag=await mateDataRef.value.Check()

  if (state.form.type=="Opaque"){
     flag=flag&&await KeyValueRef.value.Check()
  }
  if (state.form.type=="kubernetes.io/tls"){
    flag=flag&&await tlsRef.value.Check()
  }

  if (state.form.type=="kubernetes.io/dockercfg"){
    flag=flag&&await dockerRef.value.Check()
  }
  if(!flag){
    ElMessage.error("数据不合法，有必选项未填")
    return
  }
  try {
    let result=await  secretCreate(state.form)
    if (result.data.code==200){
      ElMessage("secret资源创建成功")
    }
  }catch (e){
    ElMessage.error(e)
  }
}


</script>
