<template>
  <el-card class="box-card" v-if="!render">
    <template #header>
      <div class="card-header">
        <span class="title">基本配置(metadata数据)</span>
        <el-button size="small" class="mf20" @click="state.show=!state.show">
          {{!state.show?"展开":"收缩"}}
        </el-button>
      </div>
    </template>
    <div v-show="state.show">
    <el-form  :inline="true"  :model="state.form" ref="formRef">
      <el-form-item v-if="nameRequired" label="名称"  :rules="requireRules(tipTitle+'名称必须填写')" prop="name"  >
        <el-input v-model="state.form.name"  :disabled="nameDisable"  placeholder="配置名称"></el-input>
      </el-form-item>
      <el-form-item v-if="nameRequired"  label="命名空间" prop="namespace">
        <el-select  v-model="state.form.namespace" filterable  :disabled="nameDisable">
          <el-option v-for="ns in state.nsList "
                     :label="ns.metadata.name"
                     :value="ns.metadata.name"/>
        </el-select>
      </el-form-item>
    </el-form>
      <el-divider></el-divider>
      <labelValue @input="getData($event,'labels',state.form,true,emit)" ref="MyLabelRef"></labelValue>
      <annotations @input="getData($event,'annotations',state.form,true,emit)"   ref="AnnotationsRef" :type="type"></annotations>
    </div>

  </el-card>
  <el-card class="box-card" v-if="render">
    <el-descriptions
        :column="1"
        title="资源基础信息"
    >
      <el-descriptions-item label="资源类型">{{tipTitle}}</el-descriptions-item>
      <el-descriptions-item label="资源名称">{{state.form.name}}</el-descriptions-item>
      <el-descriptions-item label="资源所属命名空间">{{state.form.namespace}}</el-descriptions-item>
      <el-descriptions-item label="资源唯一标识" >{{state.form.uid}}</el-descriptions-item>
      <el-descriptions-item label="资源创建时间" >{{state.form.creationTimestamp}}</el-descriptions-item>
    </el-descriptions>
    <labelValue :value="state.form.labels" ref="MyLabelRef"></labelValue>
    <annotations  :value="state.form.annotations" ref="AnnotationsRef" ></annotations>
  </el-card>
</template>

<script lang="ts" setup>
import {ref, defineComponent, onMounted, reactive, toRefs, watch, computed, inject} from 'vue'
import labelValue from "@/components/labelValue/labelValue.vue"
import annotations from "@/components/annotations/annotations.vue";
import { getData, IsDirty} from "../../helper/helper";
import {inArrayWithMsg, NameToArr} from "../../helper/rules";
import {arrToMap, CheckData, rmTip} from "../../helper/helper";
import {requireRules} from "../../helper/rules";
import {getNsALL} from "../../api/token/namespace/ns";
const render = inject("render")
const props = defineProps(['nameRequired',"nameDisable","tipTitle","type","value"])
const emit = defineEmits(['input'])

let state=reactive({
  nameRequired:!!props.nameRequired,
  nameDisable:!!props.nameDisable,
  tipTitle:props.tipTitle,
  show:true,
  type:props.type,
  nsList:[],
  md5:"",
  defaultData:{
    name:"",
    namespace:"default",
  },
  form:{
    name:"",
    namespace:"default",
    annotations:{},
    labels:{}
  },
  render:render,
})
let MyLabelRef=ref(null)
let AnnotationsRef=ref(null)
let checkList=[AnnotationsRef,MyLabelRef]
let formRef=ref(null)

watch(()=>props.value,()=>{
  state.form=Object.assign({},state.defaultData,props.value)
})
watch(()=>state.form.name,()=>{
  emit("input",state.form)
})
watch(()=>state.form.namespace,()=>{
  emit("input",state.form)
})

async function fetchData(){
  try {
    let tData=await getNsALL()
    state.nsList=tData.data
  }catch (e){
    console.log(e)
  }
}
fetchData()

async function Check(){
  console.log( await CheckData(checkList))
  return await CheckData(checkList)&&formRef.value.validate()
}

let nsCheck=computed(()=> {
  return NameToArr(state.nsList)
})

function setData(metadata) {
  if (metadata){

    state.form=metadata
    MyLabelRef.value.setData(metadata.labels)
    AnnotationsRef.value.setData(metadata.annotations)
  }

}
defineExpose({ setData,Check })

</script>

<style scoped>

</style>
