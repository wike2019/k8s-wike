<template>
  <main-layout class="fix-black">
    <breadcrumb title="服务.创建"></breadcrumb>
    <el-form :inline="true"  :model="state.form" ref="formRef"  >
      <el-button size="small" type="warning" class="btnList" @click="doTo('sa-list')" >进入列表</el-button >
      <el-tabs v-model="state.mode" @tab-click="Update">
        <el-tab-pane label="可视化展示" name="json">
          <metadataInfo tipTitle="Service"   @input="getData($event,'metadata',state.form,true)" ref="metadataInfoRef"  :nameRequired="true"></metadataInfo>
          <port ref="portRef" :type="state.form.spec.type"   @input="getData($event,'ports',state.form.spec,true)"></port>
          <labelValue ref="selectorRef"  label="Selector设置"  @input="getData($event,'selector',state.form.spec,true)"></labelValue>
          <el-divider></el-divider>
          <h2>其他配置</h2>
          <el-form-item class="flex1"  label="clusterIP">
            <el-select v-model="state.form.spec.clusterIP">
              <el-option label="无头服务" value="None"></el-option>
              <el-option label="集群服务" value=""></el-option>
            </el-select>
          </el-form-item>
          <el-form-item class="flex1"  label="服务类型">
            <el-select v-model="state.form.spec.type">
              <el-option label="ClusterIP" value="ClusterIP"></el-option>
              <el-option label="NodePort" value="NodePort"></el-option>
              <el-option label="LoadBalancer" value="LoadBalancer"></el-option>
              <el-option label="ExternalName" value="ExternalName"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item class="flex1"  label="节点亲和性">
            <el-select v-model="state.form.spec.sessionAffinity">
              <el-option label="ClientIP" value="ClientIP"></el-option>
              <el-option label="None" value="None"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item class="flex1"  label="服务发现策略">
            <el-select v-model="state.form.spec.internalTrafficPolicy">
              <el-option label="Cluster" value="Cluster"></el-option>
              <el-option label="Local" value="Local"></el-option>
            </el-select>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane :label="'YAML展示'" name="yaml">
          <yaml ref="yamlRef"   />
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
import {defineComponent, inject, nextTick, provide, reactive, ref, toRefs, watch} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {doTo} from "../../router";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {useRoute} from "vue-router";
import {createSvc, getSvcItem, UpdateSvc} from "../../api/token/svc/svc";
import Port from "../../components/Port/port.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import yaml from "../../components/yaml/yaml.vue";
import {getData} from "../../helper/helper"

import LabelValue from "../../components/labelValue/labelValue.vue";
import {CheckData, createTip} from "../../helper/helper";
let state=reactive({
  form:{
    metadata:{},
    spec:{
      ports:[],
      selector:{},
      clusterIP:"",
      type:"ClusterIP",
      sessionAffinity:"None",
      internalTrafficPolicy:"Cluster"
    }
  },
  name:"",
  namespace:"",
  mode:"json"
})
const route = useRoute()
let metadataInfoRef=ref(null)
let portRef=ref(null)
let selectorRef=ref(null)
let yamlRef = ref(null)

watch(() => state.form, () => {
  yamlRef.value.setData(state.form)
}, {deep: true, flush: "post"})

function Update(){
  nextTick(()=>{
    yamlRef.value.Update()
  })
}
let formRef=ref(null)
let checkList=[portRef,metadataInfoRef,selectorRef]

function post() {
  createTip()
      .then(async () => {
        if(await CheckData(checkList)&& await formRef.value.validate()){
          try {
            if (state.form.spec.clusterIP!='None'){
              delete state.form.spec.clusterIP
            }
            if (state.form.spec.type!="NodePort"){
              for (let i in state.form.spec.ports){
                delete this.state.form.spec.ports[i].nodePort
              }
            }
            let result=await createSvc(state.form)
            if (result.data.code==200){
              ElMessage("创建svc成功")
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