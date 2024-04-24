<template>
    <div v-if="type=='ingress'">
      <div class="mtb20">
        <span class="fast_action">快捷操作</span>
        <el-switch style="margin-right:20px"  v-model="state.enable_cors"  active-text="是否跨域" />
        <el-switch  v-model="state.enable_rewriter" active-text="是否重写" />
      </div>
      <el-divider></el-divider>
    </div>
    <div class="inline-title-no-back" v-if="render">
      <span>注解数据</span>
    </div>
    <labelValue @input="($event)=>{getData($event,'annotations',state.form,true,emit,'annotations');update($event)}"  ref="MyLabelRef" label="Annotations"></labelValue>
</template>

<script lang="ts" setup>
import {defineComponent, inject, reactive, ref, toRefs, watch} from 'vue'
import labelValue from "@/components/labelValue/labelValue.vue"
import {arrToMap, getData, MapToArr,IsDirty} from "../../helper/helper";
import {hasValues} from "../../helper/helper";
const render = inject("render")
const props = defineProps(['nameRequired',"nameDisable","tipTitle","type","value"])
const emit = defineEmits(['input'])

let MyLabelRef=ref(null)
let state=reactive({
  enable_cors:false,
  enable_rewriter:false,
  type:props.type||"",
  form:{
    annotations: {}
  },
  defaultData:{
    annotations:{}
  },
  render:render
})
let data=
  {
    "nginx.ingress.kubernetes.io/enable-cors":"true",
    "nginx.ingress.kubernetes.io/cors-allow-origin":"*",
    "nginx.ingress.kubernetes.io/cors-allow-methods":'GET, PUT, POST, DELETE, PATCH, OPTIONS',
    "nginx.ingress.kubernetes.io/cors-allow-headers":'Content-Type,AccessToken,X-CSRF-Token, Authorization, Token',
    "nginx.ingress.kubernetes.io/cors-expose-headers":"",
    "nginx.ingress.kubernetes.io/cors-allow-credentials":false,
  }




async function Check(){
  return await MyLabelRef.value.Check()
}

watch(()=>state.enable_cors,function (){
  clear_cors()
  if(state.enable_cors){
    add_cors()
  }
},{deep:true,flush:"post"})
watch(()=>state.enable_rewriter,function (){
  clear_rewriter()
  if(state.enable_rewriter){
    add_rewriter()
  }
},{deep:true,flush:"post"})
function AddTitle(map){
    if (map["nginx.ingress.kubernetes.io/rewrite-target"]){
      state.enable_rewriter=true
    }else {
      state.enable_rewriter=false
    }
  if (map["nginx.ingress.kubernetes.io/enable-cors"]){
    state.enable_cors=true
  }else {
    state.enable_cors=false
  }
}
function add_rewriter(){
  state.form.annotations["nginx.ingress.kubernetes.io/rewrite-target"]="/$1"
  MyLabelRef.value.setData(state.form.annotations)
}
function clear_rewriter(){
  delete  state.form.annotations["nginx.ingress.kubernetes.io/rewrite-target"]
  MyLabelRef.value.setData(state.form.annotations)
}
function clear_cors(){
  for (let i in data){
    delete  state.form.annotations[i]
    MyLabelRef.value.setData(state.form.annotations)
  }
}
function add_cors(){
  for (let i in data){
    state.form.annotations[i]=data[i]
  }
  MyLabelRef.value.setData(state.form.annotations)
}

function setData(value) {
  if(!hasValues(value)){
    state.form.annotations={}
  }
  state.form.annotations=value
  AddTitle(value)
  MyLabelRef.value.setData(value)
}
function update(value){
  AddTitle(value)
}
defineExpose({ setData,Check })
</script>

<style scoped>

</style>
