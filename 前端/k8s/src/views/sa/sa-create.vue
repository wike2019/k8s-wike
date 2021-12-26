<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>服务账号创建</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-form :inline="true"  :model="form" ref="formRef"  >
     <el-tabs v-model="mode" @tab-click="Update">
       <el-tab-pane label="可视化展示" name="json">
         <mateData @input="getData($event,'metadata',form)" ref="mateDataRef"  :nameRequired="true"></mateData>
         <el-card class="box-card">
           <template #header>
             <div class="card-header mtb20" >
               密文
               <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addSecrets" >添加密文</el-button>
             </div>
           </template>
           <div v-for="(item,index) in form.secrets" class="mtb20">
               <el-form-item
                   label="密文名称"
                   :key="'secrets'+index+'name'"
                   :prop="'secrets.'+index+'.name'"
                   :rules="requireRules('密文名称必须填写')"
               >
                 <el-select  size="small" v-model="item.name" class="selectLong" filterable>
                   <el-option :key="item.name" v-for="item in secret" :label="item.name" :value="item.name"/>
                 </el-select>
               </el-form-item>
               <el-form-item  >
                 <el-button size="small" type="danger" @click="removeSecretItems(index,'secrets')" >删除该项items</el-button>
               </el-form-item>
           </div>
         </el-card>
         <el-card class="box-card">
           <template #header>
             <div class="card-header mtb20" >
               镜像密文
               <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addImagePullSecrets" >添加镜像密文</el-button>
             </div>
           </template>
           <div v-for="(item,index) in form.imagePullSecrets" class="mtb20">
             <el-form-item
                 label="密文名称"
                 :key="'imagePullSecrets'+index+'name'"
                 :prop="'imagePullSecrets.'+index+'.name'"
                 :rules="requireRules('密文名称必须填写')"
             >
               <el-select  size="small" v-model="item.name" class="selectLong" filterable>
                 <el-option :key="item.name" v-for="item in secret" :label="item.name" :value="item.name"/>
               </el-select>
             </el-form-item>
             <el-form-item  >
               <el-button size="small" type="danger" @click="removeSecretItems(index,'imagePullSecrets')" >删除该项items</el-button>
             </el-form-item>
           </div>
         </el-card>
       </el-tab-pane>
       <el-tab-pane label="YAML展示" name="yaml">
         <yaml ref="yamlRef"  @input="yamlChange" />
       </el-tab-pane>
     </el-tabs>
     <el-divider></el-divider>
     <div style="text-align: center;margin-top: 20px">
       <el-button type="primary" @click="post()">保存</el-button>
     </div>
     </el-form>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, inject, nextTick, onMounted, reactive, ref, toRefs, watch} from 'vue'
import {ElMessageBox, ElMessage, ElLoading} from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {doTo} from "../../router";
import {getSaItem, getSaList, SACreate, SaDel, SAUpdate} from "../../api/token/sa/sa";
import {secretAllByNs, secretDetail} from "../../api/token/secret";
import {roleCreate} from "../../api/token/rbac";
import {useRoute} from "vue-router";
import yaml from "../../components/Ymal/yaml.vue";
import mateData from "../../components/Metadata/matedata.vue";
import {requireRules,inArrayWithMsg} from "../../helper/rules.ts"
import md5 from 'js-md5';
import {getData} from "../../helper/helper.ts"
import {pvcAllByNs} from "../../api/token/pvc";
import {configmapAllByNs} from "../../api/token/configmap";
export default defineComponent({
  name: 'sa-detail',
  components: {MainLayout,yaml,mateData},
  setup(){
    let state=reactive({
      item:{},
      name:"",
      name_space:"",
      mode:"json",
      form:{
        apiVersion:'v1',
        Kind:'ServiceAccount',
        metadata:{
          name:"",
          namespace:""
        },
        secrets:[],
        imagePullSecrets:[]
      },
      md5:"",
      history:{},
      secret:[],
      apiVersion:'v1',
      Kind:'ServiceAccount',
    })
    const route = useRoute()
    let loading
    state.name=route.query.name
    state.name_space=route.query.name_space
    let yamlRef=ref(null)
    let mateDataRef=ref(null)
    const formRef=ref(null)
    function Update(){
      nextTick(()=>{
        yamlRef.value.Update()
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
        state.md5=md5(JSON.stringify(state.form))
        yamlRef.value.setData(state.form)
        mateDataRef.value.setData(state.form.metadata)
      }
    },{deep:true,flush:"post"})
    function addSecrets() {
        if( state.form.secrets){
          state.form.secrets.push({name:""})
        }else{
          state.form.secrets=[{name:""}]
        }

    }
    function addImagePullSecrets() {
      if( state.form.imagePullSecrets){
        state.form.imagePullSecrets.push({name:""})
      }else{
        state.form.imagePullSecrets=[{name:""}]
      }

    }
    watch(()=>state.form.metadata.namespace,()=>{
      fetchData(state.form.metadata.namespace)
    })
    async function fetchData(ns){
      try {
        let secret=await  secretAllByNs(ns)
        state.secret=secret.data.data
        state.form.secrets=[]
        state.form.imagePullSecrets=[]
      }catch (e){
        console.log(e)
      }

    }
    function removeSecretItems(index,name) {
      state.form[name].splice(index,1)
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
    function post(){
      ElMessageBox.confirm(
          '你确定继续操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
                  formRef.value.validate(async (valid) => {
                    if (valid) {
                      try {
                        let result= await  SACreate(state.form)
                        if (result.data.code==200){
                          ElMessage("SA资源创建成功")
                          doTo('sa-detail',{name_space:state.form.metadata.namespace,name:state.form.metadata.name})
                        }
                      }catch (e){
                        console.log(e)
                      }
                    }
                  })
          })
    }
    onMounted(function () {

      yamlRef.value.setData(state.form)
      mateDataRef.value.setData(state.form.metadata)

    })
    return {...toRefs(state),doTo,Update,yamlRef,mateDataRef,yamlChange,getData,formRef,requireRules,addSecrets,addImagePullSecrets,removeSecretItems,post}
  }
})
</script>
<style >

</style>