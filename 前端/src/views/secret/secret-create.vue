<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>secret创建</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
       <el-tabs v-model="mode" @tab-click="Update">
         <el-tab-pane label="可视化展示" name="json">
          <mateData @input="getData($event,'metadata',form)" ref="mateDataRef"  :nameRequired="true"></mateData>
          <el-tabs v-model="activeName" @tab-click="changeType" >
           <el-tab-pane label="自定义类型" name="Opaque">
             <el-form ref="formRef" :model="tls"   class="mtb20">
             <KeyValue ref="KeyValue" @input="getData($event,'stringData',form)"></KeyValue>
             </el-form>
           </el-tab-pane>
           <el-tab-pane label="TLS类型" name="tls">
             <div v-if="form.type=='kubernetes.io/tls'">
               <el-alert title="为了统一密文输入必须符合可视化规范，以便后台处理" type="success" />
               <el-form ref="formRef" :model="tls"   class="mtb20">
               <el-form-item :prop="'tls_key'"    :rules="requireRules('私钥内容必须填写')" >
                   <h2 class="sub_title" >私钥内容 tls.key</h2>
                   <el-input v-model="tls.tls_key" type="textarea" :rows="8" style="width:80%"  placeholder="-----BEGIN RSA PRIVATE KEY-----开头的内容"></el-input>
                   <el-button style="margin-left:15px;vertical-align: top" type="primary" @click="()=>fileLoad('tls_key')">文件上传</el-button>
               </el-form-item>
               <el-divider></el-divider>
               <h2 class="sub_title"  >证书内容 tls.crt</h2>
               <el-form-item  :prop="'tls_value'"   :rules="requireRules('证书内容必须填写')" >
                 <el-input v-model="tls.tls_value" type="textarea" :rows="8"  style="width:80%" placeholder="-----BEGIN CERTIFICATE-----开头的内容"></el-input>
                 <el-button style="margin-left:15px;vertical-align: top" type="primary" @click="()=>fileLoad('tls_value')">文件上传</el-button>
               </el-form-item>
               </el-form>
             </div>
           </el-tab-pane>
            <el-tab-pane label="docker镜像拉取密文" name="docker">
              <div v-if="form.type=='kubernetes.io/dockercfg'">
                <el-alert title="为了统一docker密文输入必须符合可视化规范，以便后台处理" type="success" />
                <el-form ref="formRef" :model="docker"  class="mtb20" >
                <el-form-item label="server" label-width="100px" :prop="'server'"  :rules="requireRules('server必须填写')" >
                  <el-input v-model="docker.server"   placeholder="用户名" autocomplete="off" ></el-input>
                </el-form-item>
                <el-form-item label="用户名"  label-width="100px" :prop="'username'"  :rules="requireRules('用户名必须填写')" >
                  <el-input v-model="docker.username"   placeholder="用户名" autocomplete="off" ></el-input>
                </el-form-item>
                <el-form-item label="密码"  label-width="100px" :prop="'password'" :rules="requireRules('密码必须填写')" >
                  <el-input v-model="docker.password"   type="password"  placeholder="密码" autocomplete="off" ></el-input>
                </el-form-item>
                <el-form-item label="邮箱"   label-width="100px" :prop="'email'" :rules="requireRules('邮箱必须填写')" >
                  <el-input v-model="docker.email"   placeholder="邮箱" autocomplete="off" ></el-input>
                </el-form-item>
                </el-form>
              </div>
            </el-tab-pane>
            <el-divider></el-divider>
            <el-alert title="只读属性,可视化不能修改,为了兼容yaml操作," type="warning" />
            <h2 class="sub_title">Data数据</h2>
            <div  class="list_item">
              <p v-for=" (value,key) in form.data">
                key：{{key}} base64后的值 {{value}}
              </p>
            </div>
       </el-tabs>
         </el-tab-pane>
           <el-tab-pane label="YAML展示" name="yaml">
             <yaml ref="yamlRef"  @input="yamlChange" />
           </el-tab-pane>
       </el-tabs>
       <div style="text-align: center;margin-top: 20px">
         <el-button type="primary" @click="postNew()">保存</el-button>
       </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive, toRefs, nextTick, watch, onMounted} from 'vue'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {ElLoading, ElMessage} from 'element-plus'
import KeyValue from "../../components/key_value/KeyValue.vue";
import {secretCreate} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import {getData} from "../../helper/helper.ts"
import md5 from "js-md5";
import {arrToMap,data2arr} from "../../helper/helper";
import yaml from "../../components/Ymal/yaml.vue";
import mateData from "../../components/Metadata/matedata.vue";
import {requireRules} from "../../helper/rules";
export default defineComponent({
  name: 'secret-create',
  components: {
    MainLayout,
    KeyValue,
    yaml,
    mateData
  },
  setup(){
    let state=reactive({
      nsList:reactive([]),
      name:"",
      name_space:"",
      mode:"json",
      form:{
        apiVersion:'v1',
        Kind:'Secret',
        metadata:{
          name:"",
          namespace:""
        },
        type:"Opaque",
        stringData:{},
        data:{}
      },

      md5:"",
      apiVersion:'v1',
      Kind:'Secret',
      activeName:"Opaque",
      tls:{
        tls_key:"",
        tls_value:"",
      },
      docker:{
        username:"",
        server:"",
        password:"",
        email:""
      }
    })
    const route = useRoute()
    let loading
    state.name=route.query.name
    state.name_space=route.query.name_space
    let yamlRef=ref(null)
    let mateDataRef=ref(null)
    const formRef=ref(null)
    let KeyValue=ref(null)
    function Update(){
      nextTick(()=>{
        yamlRef.value.Update()
      })
    }
    function changeType(tab){
      state.tls.tls_key=""
      state.tls.tls_value=""
      state.docker={
            username:"",
            server:"",
            password:"",
            email:""
      }
      state.form.stringData={}
      if(tab.props.name=="tls"){
        state.form.type="kubernetes.io/tls"
      }else if(tab.props.name=="docker"){
        state.form.type="kubernetes.io/dockercfg"
      }else{
        state.form.type="Opaque"
      }
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

    async function fetchData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }

    fetchData()
    watch(()=>state.form,()=>{
      if(state.md5!=md5(JSON.stringify(state.form))){
        state.md5=md5(JSON.stringify(state.form))
        yamlRef.value.setData(state.form)
        mateDataRef.value.setData(state.form.metadata)
        KeyValue.value.setData(state.form.stringData)
        if( state.form.type=="kubernetes.io/tls"){
          state.tls.tls_key=state.form.stringData["tls.key"]
          state.tls.tls_value=state.form.stringData["tls.crt"]
        }
        if( state.form.type=="kubernetes.io/dockercfg"){
          state.docker=state.form.stringData
        }
        if(state.form.type=="kubernetes.io/tls"&&state.activeName!='tls'){
          state.activeName='tls'
        }
        if( state.form.type=="kubernetes.io/dockercfg"&&state.activeName!='docker'){
          state.activeName="docker"
        }
        if(state.form.type=="Opaque"&&state.activeName!='Opaque'){
          state.activeName="Opaque"
        }
      }
    },{deep:true,flush:"post"})
    function postNew(){

      formRef.value.validate(async (valid) => {
        if (valid) {
          let flag=await mateDataRef.value.Check()
          if(!flag){
            return
          }
          let flag2=await KeyValue.value.Check()
          if(!flag2){
            return
          }
            let result=await  secretCreate(state.form)
            if (result.data.code==200){
              ElMessage("secret资源创建成功")
            }
          }
      })
    }
    function inputKeyValues(input){
      state.form.rawInput=input
    }
    function fileLoad(key){
       let fileDom= document.createElement("input")
           fileDom.accept=".txt,.pem,.key,.pkey"
           fileDom.type="file"
           fileDom.onchange=function (e) {
             const reader=new FileReader()
             reader.readAsText(e.target.files[0],'UTF-8')
             reader.onload=(e)=>{
               state.tls[key]=e.target.result
             }
           }
           fileDom.click();
    }
    watch(()=> state.tls.tls_key,()=>{
      if(state.form.type=="kubernetes.io/tls"){
        state.form.stringData["tls.key"]=state.tls.tls_key
      }
    },{deep:true,flush:"post"})
    watch(()=> state.tls.tls_value,()=>{
      if(state.form.type=="kubernetes.io/tls"){
        state.form.stringData["tls.crt"]=state.tls.tls_value
      }
    },{deep:true,flush:"post"})
    watch(()=> state.docker,()=>{
      if( state.form.type=="kubernetes.io/dockercfg"){
        state.form.stringData=state.docker
      }
    },{deep:true,flush:"post"})


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
    onMounted(function () {

      yamlRef.value.setData(state.form)
      mateDataRef.value.setData(state.form.metadata)
      KeyValue.value.setData(state.form.stringData)

    })
    return {...toRefs(state),postNew,KeyValue,inputKeyValues,formRef,fileLoad,mateDataRef,yamlChange,mateDataRef,KeyValue,yamlRef,getData,Update,changeType,requireRules,data2arr}
  }
})
</script>
