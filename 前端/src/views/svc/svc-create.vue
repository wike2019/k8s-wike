<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>svc创建</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-form :inline="true"  :model="form" ref="formRef"  >
     <el-tabs v-model="mode" @tab-click="Update">
       <el-tab-pane label="可视化展示" name="json">
         <mateData @input="getData($event,'metadata',form)" ref="mateDataRef"  :nameRequired="true"></mateData>
         <el-card class="box-card">
           <template #header>
             <div class="card-header mtb20" >
               端口列表
               <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addPortsItems" >添加端口</el-button>
             </div>
           </template>
           <div v-for="(item,index) in form.spec.ports" class="mtb20">
               <el-form-item
                   label="名称"
                   :key="'spec.ports.'+index+'name'"
                   :prop="'spec.ports.'+index+'.name'"
               >
                 <el-input   v-model="item.name"  ></el-input>
               </el-form-item>
               <el-form-item
                   label="svc对外暴露端口"
                   :key="'spec.ports.'+index+'port'"
                   :prop="'spec.ports.'+index+'.port'"
               >
                 <el-input   v-model.number="item.port"  ></el-input>
               </el-form-item>
               <el-form-item
                   label="目标端口"
                   :key="'spec.ports.'+index+'targetPort'"
                   :prop="'spec.ports.'+index+'.targetPort'"
               >
                 <el-input   v-model.number="item.targetPort"  ></el-input>
               </el-form-item>
             <el-form-item
                 label="协议名称"
                 :prop="'spec.ports.'+index+'.protocol'"
                 :key="'spec.ports.'+index+'protocol'"
             >
               <el-select  size="small" v-model="item.protocol" >
                 <el-option key="TCP" label="TCP" value="TCP"/>
                 <el-option key="UDP" label="UDP" value="UDP"/>
                 <el-option key="SCTP" label="SCTP" value="SCTP"/>
               </el-select>
             </el-form-item>
               <el-form-item
                   label="nodePort"
                   :key="'spec.ports.'+index+'nodePort'"
                   :prop="'spec.ports.'+index+'.nodePort'"
               >
                 <el-input   v-model.number="item.nodePort"  ></el-input>
               </el-form-item>
               <el-form-item  >
                 <el-button size="small" type="danger" @click="removePortsItems(index)" >删除该项items</el-button>
               </el-form-item>
           </div>
           <labelSelectorMap :MatchLabels="true"  @input="Input($event,'selector')" ref="selectorRef"></labelSelectorMap>
           <div>
             <el-form-item
                 label="设置clusterIP"
                 :key="'spec.clusterIP'"
                 :prop="'spec.clusterIP'"
             >
               <el-input   v-model="form.spec.clusterIP"  ></el-input>

             </el-form-item>
             <em class="Danger" style="margin-left: 10px;">如果不需要设置为none,一般用于自定义Endpoints</em>
           </div>
           <h2 class="advanced">高级应用
             <el-button size="small" style="margin-left:20px" @click="setAdvanced(0)">
               {{!advanced[0]?"展开":"收缩"}}
             </el-button>
           </h2>
           <div v-show="advanced[0]">
             <el-form-item
                 label="设置loadBalancerIP"
                 :key="'spec.loadBalancerIP'"
                 :prop="'spec.loadBalancerIP'"
             >
               <el-input   v-model="form.spec.loadBalancerIP"  ></el-input>
             </el-form-item>
             <el-form-item
                 label="设置externalName"
                 :key="'spec.externalName'"
                 :prop="'spec.externalName'"
             >
               <el-input   v-model="form.spec.externalName"  ></el-input>
             </el-form-item>
             <el-form-item
                 label="设置ServiceType"
                 :key="'spec.type'"
                 :prop="'spec.type'"
             >
               <el-select  size="small" v-model="form.spec.type" >
                 <el-option key="ClusterIP" label="ClusterIP" value="ClusterIP"/>
                 <el-option key="NodePort" label="NodePort" value="NodePort"/>
                 <el-option key="LoadBalancer" label="LoadBalancer" value="LoadBalancer"/>
                 <el-option key="ExternalName" label="ExternalName" value="ExternalName"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="设置externalTrafficPolicy"
                 :key="'spec.type'"
                 :prop="'spec.type'"
             >
               <el-select  size="small" v-model="form.spec.externalTrafficPolicy" >
                 <el-option key="Local" label="Local" value="Local"/>
                 <el-option key="Cluster" label="Cluster" value="Cluster"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="服务亲和性"
                 :key="'spec.type'"
                 :prop="'spec.type'"
             >
               <el-select  size="small" v-model="form.spec.sessionAffinity" >
                 <el-option key="ClientIP" label="ClientIP" value="ClientIP"/>
                 <el-option key="None" label="None" value="None"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="healthCheckNodePort"
                 :key="'spec.healthCheckNodePort'"
                 :prop="'spec.healthCheckNodePort'"
             >
               <el-input   v-model.number="form.spec.healthCheckNodePort"  ></el-input>
             </el-form-item>
             <el-form-item
                 label="publishNotReadyAddresses"
                 :key="'spec.publishNotReadyAddresses'"
                 :prop="'spec.publishNotReadyAddresses'"
             >
               <el-select  size="small" v-model="form.spec.publishNotReadyAddresses" >
                 <el-option key="true" label="true" :value="true"/>
                 <el-option key="false" label="false" :value="false"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="sessionAffinityConfig"
                 :key="'spec.sessionAffinityConfig.clientIP.timeoutSeconds'"
                 :prop="'spec.sessionAffinityConfig.clientIP.timeoutSeconds'"
             >
               <el-input   v-model.number="form.spec.sessionAffinityConfig.clientIP.timeoutSeconds"  ></el-input>
             </el-form-item>
             <el-form-item
                 label="ipFamilyPolicy"
                 :key="'spec.ipFamilyPolicy'"
                 :prop="'spec.ipFamilyPolicy'"
             >
               <el-select  size="small" v-model="form.spec.ipFamilyPolicy" >
                 <el-option key="SingleStack" label="SingleStack" value="SingleStack"/>
                 <el-option key="PreferDualStack" label="PreferDualStack" value="PreferDualStack"/>
                 <el-option key="RequireDualStack" label="RequireDualStack" value="RequireDualStack"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="allocateLoadBalancerNodePorts"
                 :key="'spec.allocateLoadBalancerNodePorts'"
                 :prop="'spec.allocateLoadBalancerNodePorts'"
             >
               <el-select  size="small" v-model="form.spec.allocateLoadBalancerNodePorts" >
                 <el-option key="true" label="true" :value="true"/>
                 <el-option key="false" label="false" :value="false"/>
               </el-select>
             </el-form-item>
             <el-form-item
                 label="publishNotReadyAddresses"
                 :key="'spec.publishNotReadyAddresses'"
                 :prop="'spec.publishNotReadyAddresses'"
             >
               <el-select  size="small" v-model="form.spec.internalTrafficPolicy" >
                 <el-option key="Cluster" label="Cluster" value="Cluster"/>
                 <el-option key="Local" label="Local" value="Local"/>
               </el-select>
             </el-form-item>
             <el-divider></el-divider>
             <arrData title="clusterIPs" @input="Input($event,'clusterIPs')" ref="clusterIPsRef"></arrData>
             <el-divider></el-divider>
             <arrData title="externalIPs" @input="Input($event,'externalIPs')" ref="externalIPsRef"></arrData>
             <el-divider></el-divider>
             <arrData title="loadBalancerSourceRanges" @input="Input($event,'loadBalancerSourceRanges')" ref="loadBalancerSourceRangesRef"></arrData>
              <el-divider></el-divider>
             <arrDataLimt title="ipFamilies" @input="Input($event,'ipFamilies')" :list="['IPv4','IPv6']" ref="ipFamiliesRef"></arrDataLimt>
             <el-divider></el-divider>

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
import {secretAllByNs, secretDetail} from "../../api/token/secret/secret";
import {roleCreate} from "../../api/token/rbac";
import {useRoute} from "vue-router";
import yaml from "../../components/Ymal/yaml.vue";
import mateData from "../../components/Metadata/matedata.vue";
import {requireRules,inArrayWithMsg} from "../../helper/rules.ts"
import md5 from 'js-md5';
import {getData} from "../../helper/helper.ts"
import {pvcAllByNs} from "../../api/token/pvc";
import {configmapAllByNs} from "../../api/token/configmap/configmap";
import labelSelectorMap from "../../components/labelSelectorMap/labelSelectorMap.vue";
import arrData from "../../components/ArrData/arrData.vue";
import arrDataLimt from "../../components/ArrDataLimt/arrDataLimt.vue";
export default defineComponent({
  name: 'sa-detail',
  components: {MainLayout,yaml,mateData,labelSelectorMap,arrData,arrDataLimt},
  setup(){
    let state=reactive({
      advanced:[],
      item:{},
      name:"",
      namespace:"",
      mode:"json",
      form:{
        apiVersion:'v1',
        Kind:'Service',
        metadata:{
          name:"",
          namespace:""
        },
        spec: {
          ports:[
            {
              name:"",
              port:0,
              targetPort:"",
              protocol:"TCP",
              nodePort:0
            }
          ],
          selector:{
            "":""
          },
          clusterIP:"",
          clusterIPs:[],
          type:"ClusterIP",
          externalIPs:[],
          sessionAffinity:"None",
          loadBalancerIP:"",
          loadBalancerSourceRanges:[],
          externalName:"",
          externalTrafficPolicy:"Local",
          healthCheckNodePort:0,
          publishNotReadyAddresses:false,
          sessionAffinityConfig:{
            clientIP:{
              timeoutSeconds:1800
            }
          },
          ipFamilies:"",
          ipFamilyPolicy:"",
          allocateLoadBalancerNodePorts:false,
          internalTrafficPolicy:""
        }
      },
      md5:"",
      apiVersion:'v1',
      Kind:'Service',
    })
    let loading
    let yamlRef=ref(null)
    let mateDataRef=ref(null)
    let selectorRef=ref(null)
    let clusterIPsRef=ref(null)
    let externalIPsRef=ref(null)
    let loadBalancerSourceRangesRef=ref(null)
    let ipFamiliesRef=ref(null)
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
        selectorRef.value.setData(state.form.spec.selector)
        clusterIPsRef.value.setData(state.form.spec.clusterIPs)
        externalIPsRef.value.setData(state.form.spec.externalIPs)
        loadBalancerSourceRangesRef.value.setData(state.form.spec.loadBalancerSourceRanges)
        ipFamiliesRef.value.setData(state.form.spec.ipFamilies)
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

      }catch (e){
        console.log(e)
      }

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
      selectorRef.value.setData(state.form.spec.selector)
      clusterIPsRef.value.setData(state.form.spec.clusterIPs)
      externalIPsRef.value.setData(state.form.spec.externalIPs)
      loadBalancerSourceRangesRef.value.setData(state.form.spec.loadBalancerSourceRanges)
      ipFamiliesRef.value.setData(state.form.spec.ipFamilies)
    })
    function removePortsItems(index){
      state.form.spec.ports.splice(index,1)
    }
    function addPortsItems(){
      state.form.spec.ports.push({
        name:"",
        port:0,
        targetPort:"",
        protocol:"TCP",
        nodePort:0
      })
    }
    function Input(data,key) {
      state.form.spec[key]=data
    }
    function setAdvanced(i){
      let temp=state.advanced[i]
      for (let k=0;k<state.advanced.length;k++){
        state.advanced[k]=false
      }
      state.advanced[i]=!temp
    }

    return {...toRefs(state),doTo,Update,Input,ipFamiliesRef,loadBalancerSourceRangesRef,externalIPsRef,clusterIPsRef,yamlRef,mateDataRef,selectorRef,yamlChange,getData,formRef,addPortsItems,requireRules,removePortsItems,post,setAdvanced}
  }
})
</script>
<style >

</style>