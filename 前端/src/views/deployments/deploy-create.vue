<template>
 <main-layout>
   <nav class="nav-bar">
     <el-breadcrumb separator="/">
       <el-breadcrumb-item>deployment创建</el-breadcrumb-item>
     </el-breadcrumb>
   </nav>
   <el-tabs v-model="mode" @tab-click="Update">
     <el-tab-pane label="可视化展示" name="json">
       <mateData @input="mateDataInput" ref="mateDataRef" :nameRequired="true"></mateData>
       <core :namespace="form.metadata.namespace"  @input="coreInput" ref="coreRef"></core>
     </el-tab-pane>
     <el-tab-pane label="YAML展示" name="yaml">
       <yaml ref="yamlRef" @err="err" @input="yamlChange" />
     </el-tab-pane>
   </el-tabs>
   <el-divider></el-divider>
   <div style="text-align: center;margin-top: 20px">
     <el-button type="primary" @click="postNew()">保存</el-button>
   </div>
 </main-layout>
</template>

<script lang="ts">
import {defineComponent, nextTick, onMounted, reactive, ref, toRefs, watch} from 'vue'
import MainLayout from "../../layout/main.vue";
import yaml from "../../components/Ymal/yaml.vue";
import mateData from "../../components/Metadata/matedata.vue";
import core from "../../components/deploymentsCore/core.vue";
import md5 from 'js-md5';
import {ElMessage} from "element-plus";
import { ElLoading } from 'element-plus'
export default defineComponent({
  name: 'deploy-create',
  components: {MainLayout,yaml,mateData,core},
  setup(){
    let state=reactive({
      mode:"json",
      form:{
        apiVersion:'apps/v1',
        Kind:'Deployment',
        metadata:{
          name:"",
          namespace:"default"
        },
        spec:{

        }
      },
      md5:"",
      history:{},
      apiVersion:'apps/v1',
      Kind:'Deployment',
    })
    let loading
    let yamlRef=ref(null)
    let mateDataRef=ref(null)
    let coreRef=ref(null)
    function back(mode){
       return function () {
         loading.close()
         nextTick(function () {
           state.mode=mode
           Update()
         })
       }

    }
    function err() {
      showErr("",state.mode)
    }
    function showErr(msg,mode){
      loading=ElLoading.service({
        lock: true,
        text: '',
        spinner:"failed",
        background: 'rgba(0, 0, 0, 0.7)',
      })
      ElMessage({
        type: 'error',
        grouping: true,
        message:msg||"YAML内容有误,请仔细编辑",
        showClose:true,
        duration:0,
        onClose:back(mode)
      })
    }
    function toCore() {
      let core={
        replicas:state.form.spec.replicas,
        selector:state.form.spec.selector,
        MinReadySeconds:state.form.spec.MinReadySeconds,
        RevisionHistoryLimit:state.form.spec.RevisionHistoryLimit,
        ProgressDeadlineSeconds:state.form.spec.ProgressDeadlineSeconds,
        strategy:state.form.spec.strategy,
        template:state.form.spec.template,
      }
      return core
    }
    function yamlChange(data){
      try {
        if(data){

          if(data.apiVersion!==state.apiVersion|| data.Kind!=state.Kind){
            state.form.apiVersion=state.apiVersion
            state.form.Kind=state.Kind
            showErr( 'apiVersion和Kind不允许修改')
            return
          }
          console.log(data.spec?.template?.spec?.volumes,"yaml1111")
          //alert("修改了")
          state.form=data

        }
      }catch (e) {
        showErr("",state.mode)

      }

    }

    function coreInput(data){
          state.form.spec.selector=data.selector
          state.form.spec.template=data.template
          state.form.spec.replicas=data.replicas
          state.form.spec.strategy=data.strategy
          state.form.spec.MinReadySeconds=data.MinReadySeconds
          state.form.spec.RevisionHistoryLimit=data.RevisionHistoryLimit
          state.form.spec.ProgressDeadlineSeconds=data.ProgressDeadlineSeconds
          console.log("coreInput","执行此时")
    }
    function mateDataInput(data){
      state.form.metadata=data
    }
    watch(()=>state.form,()=>{

        if(state.md5!=md5(JSON.stringify(state.form))){
          state.md5=md5(JSON.stringify(state.form))
        // alert(JSON.stringify(state.form),"数据")
          yamlRef.value.setData(state.form)
          mateDataRef.value.setData(state.form.metadata)
          coreRef.value.setData(toCore())
          console.log('执行了几次')
        }

    },{deep:true,flush:"post"})
    function Update(){
      nextTick(()=>{
        yamlRef.value.Update()
      })
    }
   async function postNew(){
     let flag=await mateDataRef.value.Check()
     let flag2=await coreRef.value.Check()
     if(!(flag&&flag2)){
       showErr("你的有未填写的表单",state.mode)
       return
     }
     alert("成功")
    }
    onMounted(function () {

      mateDataRef.value.setData(state.form.metadata)

      coreRef.value.setData(toCore())

    })
    return {yamlRef,yamlChange,...toRefs(state),mateDataInput,Update,postNew,mateDataRef,err,coreRef,coreInput}
  }
})
</script>

<style scoped>

</style>
