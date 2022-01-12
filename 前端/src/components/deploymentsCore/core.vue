<template>
  <el-form :inline="true"  :model="form" ref="formRef" :rules="rules">
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        deployments核心数据
        <el-button size="small" style="margin-left:20px" @click="show=!show">
          {{!show?"展开":"收缩"}}
        </el-button>
      </div>
    </template>
    <div v-show="show">
       <div>
          <el-form-item label="副本数"  prop="replicas"  >
            <el-input-number v-model="form.replicas" :min="1" :max="50" />
          </el-form-item>
       </div>
       <div>
         <el-form-item label="新的Pods 已经就绪或者可用(至少持续了MinReadySeconds)"  prop="MinReadySeconds"  >
           <el-input-number v-model="form.MinReadySeconds" :min="0"  />
         </el-form-item>
        </div>
        <div>
          <el-form-item label="保留多少旧的RevisionHistory"  prop="RevisionHistoryLimit"  >
            <el-input-number v-model="form.RevisionHistoryLimit" :min="0" :max="50" />
          </el-form-item>
        </div>
        <div>
          <el-form-item label="Deployment 进展停滞之前，需要等待所给的时长"  prop="ProgressDeadlineSeconds"  >
            <el-input-number v-model="form.ProgressDeadlineSeconds" :min="0"  />
          </el-form-item>
        </div>
      <labelSelector  @input="Input($event,'selector')" ref="selectorRef"></labelSelector>
    </div>
  </el-card>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        升级策略
        <el-button size="small" style="margin-left:20px" @click="show2=!show2">
          {{!show2?"展开":"收缩"}}
        </el-button>
      </div>
    </template>
    <div v-show="show2">
        <el-form-item label="升级策略" prop="strategy.type"   :rules="[{   required: true, message: '升级策略必须填写',},{ validator:     isInArr( ['Recreate','RollingUpdate'],'升级策略只允许Recreate或者RollingUpdate'), trigger: 'change' }]"  >
          <el-select  size="small" v-model="form.strategy.type"  style="width:250px;">
            <el-option label="Recreate" value="Recreate"/>
            <el-option label="RollingUpdate" value="RollingUpdate"/>
          </el-select>
        </el-form-item>
        <div v-if="form.strategy.type=='RollingUpdate'&&form.strategy.rollingUpdate">
          <p>
          <el-form-item label="最大不可用(maxUnavailable)"  prop="strategy.rollingUpdate.maxUnavailable" :rules="[{   required: true, message: '最大不可用必须填写',}]">
            <el-input v-model="form.strategy.rollingUpdate.maxUnavailable" :min="1" :max="50" />
          </el-form-item>
          </p>
          <p>
            <el-form-item label="最大增量(maxSurge)"  prop="strategy.rollingUpdate.maxSurge" :rules="[{   required: true, message: '最大增量必须填写',}]">
              <el-input v-model="form.strategy.rollingUpdate.maxSurge" :min="1" :max="50" />
            </el-form-item>
          </p>
        </div>
    </div>
    </el-card>
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          POD配置
        </div>
      </template>
      <PodTemplate :namespace="namespace" ref="podRef" @input="Input($event,'template')"></PodTemplate>
    </el-card>
  </el-form>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs, watch, nextTick} from 'vue'
import {core} from  "../../core/core.ts"
import labelSelector from "../LabelSelector/labelSelector.vue";
import PodTemplate from "../PodTemplate/PodTemplate.vue";
import md5 from 'js-md5';
export default defineComponent({
  name: 'metadata',
  emits:["input"],
  props:["namespace"],
  components:{
    labelSelector,
    PodTemplate
  },
  setup: (props,{emit}) => {
    let defaultStrategy={
      type:"RollingUpdate",
      rollingUpdate:{
        maxUnavailable:"25%",
        maxSurge:"25%"
      }
    }
    let defaultRollingUpdate={
      maxUnavailable:"25%",
      maxSurge:"25%"
    }
    let state=reactive({
      namespace:props.namespace,
      show:true,
      show2:false,
      form:reactive({
        replicas:1,
        MinReadySeconds:0,
        RevisionHistoryLimit:10,
        ProgressDeadlineSeconds:600,
        strategy:{
          type:"RollingUpdate",
          rollingUpdate:{
            maxUnavailable:"25%",
            maxSurge:"25%"
          }
        }
      }),
      selector:{},
      template:{},
      md5:""
    })
    let selectorRef=ref(null)
    let podRef=ref(null)
    let rules={
      replicas:[
        {
          required: true,
          message: '副本数必须填写',
        }
      ],
    }

    function setData(data){

      try {
        state.form.replicas=data.replicas||1
        state.form.MinReadySeconds=data.MinReadySeconds||0
        state.form.RevisionHistoryLimit=data.RevisionHistoryLimit||10
        state.form.ProgressDeadlineSeconds=data.ProgressDeadlineSeconds||600
        state.form.strategy=data.strategy||defaultStrategy
        selectorRef.value.setData(data.selector)&&podRef.value.setData(data.template)
      }catch (e) {
        console.log(e)
        return false
      }

    }
    function convert(){
     try {
       let result={}
       result.replicas=state.form.replicas
       result.MinReadySeconds=state.form.MinReadySeconds
       result.RevisionHistoryLimit=state.form.RevisionHistoryLimit
       result.ProgressDeadlineSeconds=state.form.ProgressDeadlineSeconds
       if(state.form.strategy.type=='Recreate'){
         //alert(2)
         delete state.form.strategy.rollingUpdate
       }else{
         if(!state.form.strategy.rollingUpdate){
           state.form.strategy.rollingUpdate=defaultRollingUpdate
         }
       }
       result.strategy=state.form.strategy
       result.selector=state.selector
       result.template=state.template
       emit('input',result)

     }catch (e) {
       console.error(e)
     }
    }
    let formRef=ref(null)
    function addPathCfg(){
      selectorRef.value.addPathCfg()
    }
    async function Check(){
      return new Promise(function (resolve){
        formRef.value.validate(async (valid) => {
          if (valid) {
            let selector=await selectorRef.value.Check()
            let pod=await podRef.value.Check()
            resolve(selector&&pod)
          }else{
            await selectorRef.value.Check()
            await podRef.value.Check()
            resolve(false)
          }

        })
      })
    }
    watch(()=>state.form,()=>{
        convert()
    },{deep:true,flush:"post"})
    function Input(data,key) {
        state[key]=data
        convert()
    }

    watch(()=>props.namespace,function () {
      state.namespace=props.namespace
    })

    let isInArr=core.isInArr

    return {formRef,...toRefs(state),addPathCfg,selectorRef,Check,rules,setData,Input,isInArr,podRef}
  }
})
</script>

<style scoped>

</style>
