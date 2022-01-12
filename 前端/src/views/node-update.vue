<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>节点编辑</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-form :inline="true"  :model="form" ref="formRef" >
     <el-card class="box-card">
       <template #header>
         <div class="card-header">
           基本信息
         </div>
       </template>
         <el-form-item label="名称"  prop="name"  >
           <el-input v-model="form.name"  disabled  placeholder="节点名称"></el-input>
         </el-form-item>
       <el-form-item label="名称"  prop="name"  >
         <el-input v-model="form.ip"  disabled  placeholder="节点ip"></el-input>
       </el-form-item>
     </el-card>
       <el-card class="box-card">
           <template #header>
             <div class="card-header">
               系统标签请谨慎修改
             </div>
           </template>
           <KeyValue ref="Label" @input="inputKeyValues" ></KeyValue>
         </el-card>
       <el-card class="box-card">
         <template #header>
           <div class="card-header">
             污点
           </div>
         </template>
         <Taint ref="Taint" @input="inputTaintValues" ></Taint>
       </el-card>
     </el-form>
     <div style="text-align: center;margin-top: 20px">
       <el-button type="primary" @click="postNew()">保存</el-button>
     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {ingressCreate} from "../api/token/ingress/ingress";
import {getSvcList} from "../api/token/svc/svc";
import { ElMessage } from 'element-plus'
import KeyValue from "../components/key_value/KeyValue.vue";
import {secretCreate} from "../api/token/secret/secret";
import Taint from "../components/Taint/Taint.vue";
import {useRoute} from "vue-router";
import {getNodeDetail, saveNode} from "../api/token/node";
export default defineComponent({
  name: 'secret-create',
  components: {
    MainLayout,
    KeyValue,
    Taint
  },
  setup(){
    let state=reactive({
      taintrawInput:[],
      labelrawInput:[],
      form:reactive({
        ip:"",
        name:"",
        orgin_labels: {},
        orgin_taints:[

        ]
      }),
    })
    let formRef=ref(null)
    let Label=ref(null)
    let Taint=ref(null)


    async function getData(){
      const route = useRoute()
      try {
       let tData=await getNodeDetail(route.query.name)
        state.form.ip=tData.data.data.ip
        state.form.name=tData.data.data.name
        let arr=[];
        for ( let i in tData.data.data.OrginLabels){
          arr.push({"key":i,value:tData.data.data.OrginLabels[i],NoRequired:!!tData.data.data.OrginLabels[i]})
        }
        Label.value.setData(arr)
        Taint.value.setData(tData.data.data.OrginTaints)
      }catch (e){
        console.log(e)
      }
    }

    getData()
    function convert(input){
      console.log(input)
      let obj={}
      for (let i=0;i<input.length;i++){
        obj[input[i].key]=input[i].value
      }
      return obj;
    }
    async  function postNew(){

          let flag=await Label.value.Check()
          let flag2=await Taint.value.Check()
          if(!flag||!flag2){
            return
          }
          try {
             state.form.orgin_labels=convert(state.labelrawInput)
             state.form.orgin_taints=state.taintrawInput
            console.log(state)
            let result=await saveNode(state.form)
            if (result.data.code==200){
              ElMessage("资源更新成功")
            }
          }catch (e){
            ElMessage.error("资源格式有误")
            console.log(e)
          }
    }
    function inputKeyValues(input){
      //console.log(inputKeyValues)
      state.labelrawInput=input
    }

    function postNewTLS(){
      formRef.value.validate(async (valid) => {
        if (valid) {
          if(!state.form.tls_key||!state.form.tls_value){
            ElMessage.error("证书文件和证书密钥必须上传")
            return
          }
          try {

            state.form.data={
              "tls.key":state.form.tls_key,
              "tls.crt":state.form.tls_value
            }
            state.form.type="kubernetes.io/tls"
            let result=await  secretCreate(state.form)
            if (result.data.code==200){
              ElMessage("secret资源创建成功")
            }

            console.log(result)
          }catch (e){
            ElMessage.error("资源格式有误")
            console.log(e)
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
    function inputTaintValues(input){
      state.taintrawInput=input
    }

    return {...toRefs(state),inputTaintValues,inputKeyValues,Label,Taint,formRef,postNew}
  }
})
</script>
