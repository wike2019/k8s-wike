<template>
  <el-form :inline="true"  :model="state.form" ref="formRef"  >
  <h3>默认后端 <el-button size="small" type="success"  class="mf20" @click="InitDefaultBackend(true)" >添加默认后端</el-button></h3>
    <el-divider></el-divider>
  <div  class="mtb20" v-if="state.form.defaultBackend">
    <el-form-item
        label="选择对应资源"
        :key="'defaultType'"
        :prop="'defaultType'"
    >
      <el-select  v-model="state.defaultType" >
        <el-option label="服务" value="service"/>
        <el-option label="自定义资源" value="resource"/>
      </el-select>
    </el-form-item>
    <el-form-item  >
      <el-button size="small" type="danger" @click="InitDefaultBackend(false)" >删除默认后端</el-button>
    </el-form-item>
    <div v-if="state.defaultType=='service'&&state.form.defaultBackend.service">
      <el-form-item label="服务名"

                    :rules="requireRules('服务名必须填写')"
                    key="defaultBackend-name"  :prop="'defaultBackend.service.name'"
      >
        <el-select  v-model="state.form.defaultBackend.service.name" >
          <el-option v-for="svc in state.svcList"
                     :label="svc.metadata.name"
                     :value="svc.metadata.name"/>
        </el-select>
      </el-form-item>
      <el-form-item
          v-if="checkData(state.form.defaultBackend.service.name).length>0&&!state.port"
          label="端口名称"
          key="defaultBackend-port-name"
          :prop="'defaultBackend.service.port.name'"
          :rules="requireRules('端口名必须填写')"
      >
        <el-select  v-model="state.form.defaultBackend.service.port.name" >
          <el-option :label="item.name" :value="item.name"  v-for="item in checkData(state.form.defaultBackend.service.name)"/>
        </el-select>
      </el-form-item>
      <el-form-item  label="端口"  v-if="(checkData(state.form.defaultBackend.service.name).length==0||state.port)" key="defaultBackend-name-port" :prop="'defaultBackend.service.port.number'">
        <el-input-number min="1" max="65500"     v-model.number="state.form.defaultBackend.service.port.number"  placeholder="填写服务端口"></el-input-number>
      </el-form-item>
      <el-form-item label="自定义port"  >
        <el-checkbox v-model="state.port"></el-checkbox>
      </el-form-item>
    </div>
    <div v-if="state.defaultType=='resource'&&state.form.defaultBackend.resource" >
      <el-form-item
          :rules="requireRules('资源名称必须填写')"
          label="资源名称"
          :key="'defaultBackend.resource.name'"
          :prop="'defaultBackend.resource.name'"

      >
        <el-input   v-model="state.form.defaultBackend.resource.name"  placeholder="资源名称"></el-input>
      </el-form-item>
      <el-form-item
          label="apiGroup"
          :rules="requireRules('apiGroup必须填写')"
          :key="'defaultBackend.resource.apiGroup'"
          :prop="'defaultBackend.resource.apiGroup'"

      >
        <el-input  v-model="state.form.defaultBackend.resource.apiGroup"  placeholder="填写apiGroup"></el-input>
      </el-form-item>
      <el-form-item
          label="kind"
          :rules="requireRules('kind必须填写')"
          :key="'defaultBackend.resource.kind'"
          :prop="'defaultBackend.resource.kind'"

      >
        <el-input  v-model="state.form.defaultBackend.resource.kind"  placeholder="填写kind"></el-input>
      </el-form-item>
    </div>

  </div>
  </el-form>
</template>

<script setup>
import {reactive, watch} from "vue";
 import {inArrayWithMsg, requireRules} from "../../../helper/rules";
 import {getSvcList} from "../../../api/token/svc/svc";
 import {secretAllByNs} from "../../../api/token/secret/secret";
import {arrToMap, hasValues, IsDirty, MapToArr} from "../../../helper/helper";
 const props = defineProps(['namespace'])
const emit = defineEmits(['input'])

 let state=reactive({
   form:{
     defaultBackend:{
     }},
   svcList:[],
   tlsList:[],
   port:false,
   md5:"",
   defaultType:"service"
 })
watch(()=>state.port,()=>{
  state.form.defaultBackend.service.port.number=80
  state.form.defaultBackend.service.port.name=""
})
watch(()=>state.form,()=>{
  if(!IsDirty(state)) {
    let cache=Object.assign({},state.form.defaultBackend)
    if (state.defaultType=="service"){
      if (cache.service){
        if (cache.service.port.name){
          delete cache.service.port.number
        }else {
          delete cache.service.port.name
        }
      }
    }

    emit("input",cache)
  }
},{deep:true,flush:"post"})
  watch(()=>state.defaultType,()=>{
     if (state.defaultType=="service"){
       state.form.defaultBackend={
         service:{
           name:"",
           port:{
             name:"",
             number:80// 2 选一
           }
         }
       }
     }else{
       state.form.defaultBackend={
         resource:{
           name:"",
           apiGroup:"",
           kind:""
         }
       }
     }
  })
 function InitDefaultBackend(flag){
   if(flag){
     if (state.defaultType=="service"){
       state.form.defaultBackend={
         service:{
           name:"",
           port:{
             name:"",
             number:0// 2 选一
           }
         }
       }
     }else{
       state.form.defaultBackend={
         resource:{
           name:"",
           apiGroup:"",
           kind:""
         }
       }
     }
   }else{
     delete state.form.defaultBackend
   }
 }

 async function init(){
   try {
     let tData=await getSvcList(props.namespace)
     let tDataTls=await secretAllByNs(props.namespace)
     state.tlsList=tDataTls.data.data.filter(function (item) {
       return item.type=="TLS凭据"
     })
     state.svcList=tData.data.data
   }catch (e){
     console.log(e)
   }
 }
 init()
 function getsvcnameList(svc){
   if (!svc){
     return []
   }

   let ports=state.svcList.filter((item)=>{
     return item.metadata.name==svc
   })
   if(ports.length){
     return ports[0].spec.ports
   }
   return  []
 }
  function checkData(svc){
    if (!svc){
      return []
    }
   let list= getsvcnameList(svc)
   return list.filter((item)=>{
     return  item.name!=undefined
   })
 }

async  function Check(){
  return  await formRef.value.validate()
}
function setData(value) {
  if (value&&value.service){
    state.defaultType="service"
  }else if  (value&&value.resource){
    state.defaultType="resource"
  }
  state.form.defaultBackend=value
}
defineExpose({ setData,Check })

</script>

<style scoped>

</style>