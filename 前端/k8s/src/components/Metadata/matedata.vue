<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        基本配置(metadata数据)
        <el-button size="small" style="margin-left:20px" @click="show=!show">
          {{!show?"展开":"收缩"}}
        </el-button>
      </div>
    </template>
    <div v-show="show">
    <el-form  :inline="true"  :model="form" ref="formRef" :rules="rules">
      <el-form-item v-if="nameRequired" label="名称"  prop="name"  >
        <el-input v-model="form.name"  :disabled="nameDisable"  placeholder="配置名称"></el-input>
      </el-form-item>
      <el-form-item v-if="nameRequired"  label="命名空间" prop="name_space">
        <el-select  v-model="form.namespace"  :disabled="nameDisable">
          <el-option v-for="ns in nsList "
                     :label="ns.name"
                     :value="ns.name"/>
        </el-select>
      </el-form-item>
    </el-form>

    <div class="inline-title">
     <span>Label数据设置</span>         <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfg" >+</el-button>
    </div>
      <MyLabel @input="getData($event,'label',data)" ref="MateRef"></MyLabel>
      <el-divider></el-divider>
      <div class="inline-title">
        <span>注解数据设置</span>         <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfgAnnotations" >+</el-button>
      </div>
      <Annotations @input="getData($event,'annotations',data)" ref="Annotations"></Annotations>
      <el-divider></el-divider>
    </div>

  </el-card>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs, watch} from 'vue'
import {getNsList} from "../../api/token/namespace/ns";
import MyLabel from "../Label/label.vue"
import Annotations from "../Annotations/Annotations.vue";
import md5 from 'js-md5';

import {arrToMap, data2arr, getData} from "../../helper/helper";
export default defineComponent({
  name: 'metadata',
  emits:["input"],
  props:["nameRequired","nameDisable"],
  components:{
    MyLabel,
    Annotations
  },
  setup: (props,{emit}) => {
    let state=reactive({
      nameRequired:!!props.nameRequired,
      nameDisable:!!props.nameDisable,
      show:true,
      nsList:reactive([]),
      form:{
        name:"",
        namespace:"default",
      },
      data:{
        labels:[],
        annotations:[],
      },
    })
    let MateRef=ref(null)
    let Annotations=ref(null)
    let rules={
      name:[
        {
          required: true,
          message: '配置名称必须填写',
        }
      ],
      namespace:[
        {
          required: true,
          message: '命名空间必须填写',
          trigger: 'blur',
        }
      ],
    }
    async function fetchData(){
      try {
        let tData=await getNsList()
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }

    function setData(data){
      try {
        if(data&&data.name){
          state.form.name=data.name||""
        }
        if(data&&data.namespace){
          state.form.namespace=data.namespace||"default"
        }
        if(data&&data.labels){
          data2arr(data.labels,state.data.labels,MateRef.value)
        }
        if(data&&data.annotations){
          data2arr(data.annotations,state.data.annotations,Annotations.value)
        }
        commit()
        return true
      }catch (e) {
        console.log(e)
        return false
      }

    }
    function commit(){

      let result={}
      result.name=state.form.name
      result.namespace=state.form.namespace

      result.labels= arrToMap(state.data.labels)
      result.annotations= arrToMap(state.data.annotations)
      emit('input',result)

    }
    fetchData()
    let formRef=ref(null)
    function addPathCfg(){
      MateRef.value.addPathCfg()
    }
    async function Check(){
      return new Promise(function (resolve){
        formRef.value.validate(async (valid) => {
          if (valid) {
            let flag=await MateRef.value.Check()
            let flag2=await Annotations.value.Check()
            resolve(flag&&flag2)
          }else{
            await MateRef.value.Check()
            await Annotations.value.Check()
            resolve(false)
          }

        })
      })
    }
    watch(()=>state.form,()=>{
      commit()
    },{deep:true,flush:"post"})
    watch(()=>state.data,()=>{
      commit()
    },{deep:true,flush:"post"})
    function addPathCfgAnnotations() {
      Annotations.value.addPathCfg()
    }
    return {formRef,...toRefs(state),addPathCfg,MateRef,Check,rules,setData,addPathCfgAnnotations,Annotations,getData}
  }
})
</script>

<style scoped>

</style>
