<template>
  <main-layout>
    <nav class="nav-bar">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>配置创建</el-breadcrumb-item>
      </el-breadcrumb>
    </nav>
    <el-form ref="formRef" :model="form"   class="mtb20">
      <el-tabs v-model="mode" @tab-click="Update">
        <el-tab-pane label="可视化展示" name="json">
          <mateData @input="getData($event,'metadata',form)" ref="mateDataRef"  :nameRequired="true"></mateData>
          <KeyValue ref="KeyValueRef" @input="getData($event,'data',form)" ></KeyValue>
        </el-tab-pane>
        <el-tab-pane label="YAML展示" name="yaml">
          <yaml ref="yamlRef"  @input="yamlChange" />
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <div style="text-align: center;margin-top: 20px">
      <el-button type="primary" @click="postNew()">保存</el-button>
    </div>
  </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive, toRefs, nextTick, watch} from 'vue'
import MainLayout from "../../layout/main.vue";

import {ElLoading, ElMessage} from 'element-plus'
import KeyValue from "../../components/key_value/KeyValue.vue";
import {configmapCreate, configmapDetail, configmapUpdate} from "../../api/token/configmap/configmap";
import {useRoute} from "vue-router";
import {getData} from "../../helper/helper.ts"
import md5 from "js-md5";
import yaml from "../../components/Ymal/yaml.vue";
import mateData from "../../components/Metadata/matedata.vue";
export default defineComponent({
  name: 'configmap-create',
  components: {
    MainLayout,
    KeyValue,
    mateData,
    yaml
  },
  setup(){
    let state=reactive({
      nsList:reactive([]),
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
      md5:"",
      apiVersion:'v1',
      Kind:'ConfigMap',
    })
    //公共部分
    let formRef=ref(null)
    let KeyValueRef=ref(null)
    let mateDataRef=ref(null)
    let yamlRef=ref(null)
    function Update(){
      nextTick(()=>{
        yamlRef.value.Update()
      })
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
    function back(mode){
      return function () {
        loading.close()
        nextTick(function () {
          state.mode=mode
          Update()
        })
      }

    }

    function postNew(){

      formRef.value.validate(async (valid) => {
        if (valid) {
          let flag=await KeyValueRef.value.Check()
          if(!flag){
            return
          }

          try {
            let result=await  configmapCreate(state.form)
            if (result.data.code==200){
              ElMessage("配置创建成功")
            }
          }catch (e){
            console.log(e)
          }
        }
      })
    }

    function yamlChange(data){
      try {
        if(data){
          if(data.apiVersion!==state.apiVersion|| data.Kind!=state.Kind){
            state.form.apiVersion=state.apiVersion
            state.form.Kind=state.Kind
            showErr( 'apiVersion和Kind不允许修改',state.mode)
          }

          state.form=data
        }
      }catch (e) {
        showErr("",state.mode)
      }
    }

    watch(()=>state.form,()=>{
      if(state.md5!=md5(JSON.stringify(state.form))){

        yamlRef.value.setData(state.form)
        mateDataRef.value.setData(state.form.metadata)
        KeyValueRef.value.setData(state.form.data)
        state.md5=md5(JSON.stringify(state.form))
      }
    },{deep:true,flush:"post"})
    return {...toRefs(state),postNew,KeyValueRef,formRef,yamlRef,Update,mateDataRef,getData,yamlChange}
  }
})
</script>
