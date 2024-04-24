<template>
   <main-layout>
     <breadcrumb title="密文.更新"></breadcrumb>
     <el-button size="small" type="warning" class="btnList" @click="doTo('secret-list')" >进入列表</el-button >
     <el-tabs v-model="state.mode" @tab-click="Update">
           <el-tab-pane label="可视化展示" name="json">
           <metadataInfo ref="mateDataRef" @input="getData($event,'metadata',state.form,true)"  tipTitle="Secret" type="secret" :nameDisable="true" :nameRequired="true"></metadataInfo>
           <el-alert title="同名key会覆盖Data,下面的数据同key数据，且type不允许切换" type="error" />
          <el-tabs v-model="state.form.type"  >
           <el-tab-pane label="自定义类型" name="Opaque" disabled>
             <KeyValue ref="KeyValueRef" :base64="true"  @input="getData($event,'stringData',state.form,true)" v-if="state.form.type=='Opaque'"></KeyValue>
           </el-tab-pane>
            <el-tab-pane label="服务账号令牌" name="kubernetes.io/service-account-token" disabled>
              <el-alert title="服务账号令牌数据不可以修改" type="warning" />
            </el-tab-pane>
           <el-tab-pane label="TLS类型" name="kubernetes.io/tls" disabled>
              <Tls ref="tlsRef" @input="getData($event,'stringData',state.form,true)" v-if="state.form.type=='kubernetes.io/tls'"></Tls>
           </el-tab-pane>
            <el-tab-pane label="docker镜像拉取密文" name="kubernetes.io/dockercfg" disabled>
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
  provide
} from 'vue'
import MainLayout from "../../layout/main.vue";
import {getNsALL} from "../../api/token/namespace/ns";
import { ElMessage} from 'element-plus'
import KeyValue from "../../components/key_value/KeyValue.vue";
import { secretDetail, secretUpdate} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import {getData} from "../../helper/helper"

import yaml from "../../components/yaml/yaml.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {requireRules} from "../../helper/rules";
import breadcrumb from "../../components/common/breadcrumb.vue";
import Docker from "./components/docker.vue";
import Tls from "./components/tls.vue";
import {doTo} from "../../router";

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
provide("render",false)

state.form.metadata.name=route.query.name.toString()
state.form.metadata.namespace=route.query.namespace.toString()
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

async function getDataItem(){
  try {
    let tData=await secretDetail(state.form.metadata.namespace,state.form.metadata.name)
    let secret=tData.data.data
    state.form=secret.Secret
    mateDataRef.value.setData(secret.Secret.metadata)
    yamlRef.value.setData(state.form)
    if (state.form.type=="Opaque"){
      nextTick(()=>{
        KeyValueRef.value.setData(state.form.data)
      })

    }
    if (state.form.type=="kubernetes.io/tls"){
      nextTick(()=>{
        tlsRef.value.setData(state.form.data)
      })

    }
    if (state.form.type=="kubernetes.io/dockercfg"){
      nextTick(()=>{
        dockerRef.value.setData(state.form.data)
      })

    }
  }catch (e){
    console.log(e)
  }
}
getDataItem()


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
        let result=await  secretUpdate(state.form)
        if (result.data.code==200){
          ElMessage("secret资源修改成功")
        }
    }catch (e){
      ElMessage.error(e)
    }
}




</script>
