<template>
   <main-layout>
     <breadcrumb title="ingress.更新"></breadcrumb>
     <el-form :inline="true"  :model="form" ref="formRef"  >
       <el-button size="small" type="warning" class="btnList" @click="doTo('ingress-list')" >进入列表</el-button >
       <el-tabs v-model="state.mode" @tab-click="Update">
         <el-tab-pane label="可视化展示" name="json">
           <metadataInfo ref="mateDataRef" @input="getData($event,'metadata',state.form,true)"  tipTitle="ingress" type="ingress" :nameRequired="true"></metadataInfo>
         <el-card class="box-card">
           <template #header>
             <div class="card-header">
               规则
             </div>
           </template>
           <ingress-backend ref="backendRef" :namespace="state.namespace"  @input="getData($event,'defaultBackend',state.form.spec,true)"></ingress-backend>
<!--           <el-divider></el-divider>-->
<!--           <h3>添加路由 <el-button size="small" type="success"  class="mf20" @click="addPathCfg" >添加路由</el-button></h3>-->
<!--           <el-divider></el-divider>-->
<!--           <div v-for="(rule, index) in form.spec.rules"  class="mtb20">-->
<!--               <el-form-item-->
<!--                              label="域名"-->
<!--                              :rules="requireRules('域名填写')"-->
<!--                              :key="'spec.rules.' + index + '.host'"-->
<!--                              :prop="'spec.rules.' + index + '.host'"-->
<!--               >-->
<!--                 <el-input v-model="rule.host"   placeholder="填写域名"></el-input>-->
<!--               </el-form-item>-->
<!--               <el-form-item>-->
<!--               <h3>添加路由规则 <el-button size="small" type="success"  class="mf20" @click="addPaths(index)" >添加规则</el-button></h3>-->
<!--               </el-form-item>-->
<!--               <el-form-item  >-->
<!--                 <el-button  size="small" type="danger" @click="rmPathCfg(index,pathindex)" >删除该路由</el-button>-->
<!--               </el-form-item>-->
<!--               <div  v-for="(path,pathindex) in rule.http.paths" class="mtb20" >-->
<!--                  <el-form-item-->
<!--                            label="Path"-->
<!--                            :rules="requireRules('路径必须填写')"-->
<!--                            :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.path'"-->
<!--                            :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.path'"-->
<!--                  >-->
<!--                    <el-input v-model="path.path"   placeholder="Path"></el-input>-->
<!--                  </el-form-item>-->
<!--                <el-form-item-->

<!--                    label="匹配模式"-->
<!--                    :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.pathType'"-->
<!--                    :rules="requireRules('匹配模式必须填写')"-->
<!--                    :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.pathType'"-->
<!--                >-->
<!--                  <el-select v-model="path.pathType">-->
<!--                    <el-option label="前缀匹配" value="Prefix"/>-->
<!--                    <el-option label="精确匹配" value="Exact"/>-->
<!--                  </el-select>-->
<!--                </el-form-item>-->
<!--                <el-form-item-->
<!--                    label="选择对应资源"-->
<!--                    :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.type'"-->
<!--                    :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.type'"-->
<!--                >-->
<!--                <el-select  v-model="type[index][pathindex]" @change="fixdata(index,pathindex)">-->
<!--                  <el-option label="服务" value="service"/>-->
<!--                  <el-option label="自定义资源" value="other"/>-->
<!--                </el-select>-->
<!--                </el-form-item>-->
<!--                 <el-form-item  >-->
<!--                   <el-button  size="small" type="danger" @click="removePaths(index,pathindex)" >删除该规则</el-button>-->
<!--                 </el-form-item>-->
<!--                <div class="mtb20" v-if="type[index][pathindex]=='service'">-->
<!--                 <el-form-item-->
<!--                            :rules="requireRules('服务名必须填写')"-->
<!--                            label="服务名"-->
<!--                            :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.name'"-->
<!--                            :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.name'"-->

<!--                 >-->

<!--                   <el-select v-model="path.backend.service.name">-->
<!--                     <el-option v-for="svc in svcList"-->
<!--                                :label="svc.name"-->
<!--                                :value="svc.name"/>-->
<!--                   </el-select>-->
<!--                 </el-form-item>-->
<!--&lt;!&ndash;                 <em>需要改</em>&ndash;&gt;-->
<!--                 <el-form-item-->
<!--                     label="端口名称"-->
<!--                     :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.port.name'"-->
<!--                     :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.port.name'"-->
<!--                     :rules="inArrayWithMsg(checkData(path.backend.service.name),'你的端口名不正确')"-->
<!--                 >-->
<!--                   <el-select  v-model="path.backend.service.port.name" >-->
<!--                     <el-option :label="item.name" :value="item.name"  v-for="item in getsvcnameList(path.backend.service.name)"/>-->
<!--                   </el-select>-->
<!--                 </el-form-item>-->
<!--                 <el-form-item-->
<!--                        v-if="!path.backend.service.port.name" label="端口"-->
<!--                     :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.port.number'"-->
<!--                     :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.service.port.number'"-->
<!--                 >-->
<!--                   <el-input   v-model.number="path.backend.service.port.number"    placeholder="填写服务端口"></el-input>-->
<!--                 </el-form-item>-->
<!--                </div>-->
<!--                <div class="mtb20"  v-if="type[index][pathindex]=='other'">-->
<!--                  <el-form-item-->
<!--                      :rules="requireRules('资源名称必须填写')"-->
<!--                      label="资源名称"-->
<!--                      :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.name'"-->
<!--                      :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.name'"-->

<!--                  >-->
<!--                    <el-input   v-model="path.backend.resource.name"  placeholder="资源名称"></el-input>-->
<!--                  </el-form-item>-->
<!--                  <el-form-item-->
<!--                      label="apiGroup"-->
<!--                      :rules="requireRules('apiGroup必须填写')"-->
<!--                      :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.apiGroup'"-->
<!--                      :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.apiGroup'"-->

<!--                  >-->
<!--                    <el-input  v-model="path.backend.resource.apiGroup"  placeholder="填写apiGroup"></el-input>-->
<!--                  </el-form-item>-->
<!--                  <el-form-item-->
<!--                      label="kind"-->
<!--                      :rules="requireRules('kind必须填写')"-->
<!--                      :key="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.kind'"-->
<!--                      :prop="'spec.rules.'+index+'.http.paths.'+pathindex+'.backend.resource.kind'"-->

<!--                  >-->
<!--                    <el-input  v-model="path.backend.resource.kind"  placeholder="填写kind"></el-input>-->
<!--                  </el-form-item>-->
<!--                </div>-->
<!--                <div class="dashed"></div>-->
<!--              </div>-->
<!--           </div>-->
<!--           <h3>添加tls证书 <el-button size="small" type="success"  class="mf20" @click="addTls" >添加证书</el-button></h3>-->
<!--           <el-divider></el-divider>-->
<!--           <div v-for="(item, index) in form.spec.tls" class="mtb20" >-->
<!--             <el-form-item-->
<!--                 :rules="requireRules('证书名必须填写')"-->
<!--                 label="证书名"-->
<!--                 :key="'spec.tls.'+index+'.secretName'"-->
<!--                 :prop="'spec.tls.'+index+'.secretName'"-->
<!--             >-->
<!--               <el-select v-model="item.secretName">-->
<!--                 <el-option v-for="tls in tlsList"-->
<!--                            :label="tls.name"-->
<!--                            :value="tls.name"/>-->
<!--               </el-select>-->
<!--             </el-form-item>-->
<!--             <div>-->
<!--               <el-form-item label="添加域名">-->
<!--                 <el-input   v-model="host[index]"  placeholder="例如 www.baidu.com"></el-input>-->
<!--               </el-form-item>-->
<!--               <el-form-item>-->
<!--                 <el-button  size="small" type="primary" @click="addHost(index)" >确认添加域名</el-button>-->
<!--               </el-form-item>-->
<!--               <el-form-item>-->
<!--               <el-button class="mf20"  size="small" type="danger" @click="removeTls(index)" >删除该项</el-button>-->
<!--               </el-form-item>-->
<!--             </div>-->
<!--             <div v-for="(host, index2) in item.hosts">-->
<!--               <div class="mtb20" v-if="host">-->
<!--                 域名: {{host}}    <el-button class="mf20"  size="small" type="danger" @click="removeHost(index,index2)" >删除该域名</el-button>-->
<!--               </div>-->
<!--             </div>-->
<!--             <div class="dashed"></div>-->
<!--           </div>-->
         </el-card>
         </el-tab-pane>
         <el-tab-pane label="YAML展示" name="yaml">
           <yaml ref="yamlRef"  @input="yamlChange" />
         </el-tab-pane>
       </el-tabs>
       <el-divider></el-divider>
       <div style="text-align: center;margin-top: 20px">
         <el-button type="primary" @click="postNew()">保存</el-button>
       </div>
     </el-form>
   </main-layout>
</template>

<script lang="ts" setup>
import {defineComponent, computed, ref, onUnmounted, inject, reactive, toRefs, nextTick, watch, onMounted} from 'vue'
import MainLayout from "../../layout/main.vue";
import {getData} from "../../helper/helper"
import {ingressCreate, ingressItem, ingressUpdate} from "../../api/token/ingress/ingress";
import {getSvcList} from "../../api/token/svc/svc";
import {ElLoading, ElMessage} from 'element-plus'
import md5 from "js-md5";
import {inArrayWithMsg, requireRules} from "../../helper/rules";
import {secretAllByNs} from "../../api/token/secret/secret";
import {getSaItem} from "../../api/token/sa/sa";
import {useRoute} from "vue-router";
import {doTo} from "../../router";
import yaml from "../../components/yaml/yaml.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import IngressBackend from "./components/IngressBackend.vue";
const ws = inject("ws")
let state=reactive({
  mode:"json",
  form:{
    apiVersion:'networking.k8s.io/v1',
    Kind:'Ingress',
    metadata:{
      name:"",
      namespace:""
    },
    spec:{
      ingressClassName:"nginx",
      tls:[],
      rules:[],
      defaultBackend:{}
    }
  },
  name:"",
  namespace:"default",
  md5:"",
})
let formRef=ref(null)
let mateDataRef=ref(null)
let yamlRef=ref(null)
let backendRef=ref(null)
const route = useRoute()
let loading
function Update(){
  nextTick(()=>{
    yamlRef.value.Update()
  })
}
// function showErr(msg,mode){
//   loading=ElLoading.service({
//     lock: true,
//     text: '',
//     spinner:"failed",
//     background: 'rgba(0, 0, 0, 0.7)',
//   })
//   ElMessage({
//     type: 'error',
//     grouping: true,
//     message:msg||"YAML内容有误,请仔细编辑",
//     showClose:true,
//     duration:0,
//     onClose:back(mode)
//   })
// }
// function back(mode){
//   return function () {
//     loading.close()
//     nextTick(function () {
//       state.mode=mode
//       Update()
//     })
//   }
//
// }
// //公共部分结束
// function yamlChange(data){
//   try {
//     if(data){
//       if(data.apiVersion!==state.apiVersion|| data.Kind!=state.Kind){
//         state.form.apiVersion=state.apiVersion
//         state.form.Kind=state.Kind
//         showErr( 'apiVersion和Kind不允许修改',state.mode)
//       }
//       if(data.metadata.namespace!==state.namespace|| data.metadata.name!=state.name){
//         state.form.metadata.namespace=state.namespace
//         state.form.metadata.name=state.name
//         showErr( 'name和namespace不允许修改',state.mode)
//       }
//       data.spec.rules.forEach((item,index)=>{
//         item.http.paths.forEach((path,index2)=>{
//           if(path.backend.service){
//             if(state.type[index]){
//               state.type[index][index2]="service"
//             }else{
//               state.type[index]=[]
//               state.type[index][index2]="service"
//             }
//           }else{
//             if(state.type[index]){
//               state.type[index][index2]="other"
//             }else{
//               state.type[index]=[]
//               state.type[index][index2]="other"
//             }
//           }
//
//         })
//       })
//       state.form=data
//     }
//   }catch (e) {
//     showErr("",state.mode)
//   }
// }
//
watch(()=>state.form,()=>{
    yamlRef.value.setData(state.form)
},{deep:true,flush:"post"})
// async function changeNs(ns){
//   try {
//     let tData=await getSvcList(ns)
//     let tDataTls=await secretAllByNs(ns)
//     state.tlsList=tDataTls.data.data.filter(function (item) {
//       return item.type=="TLS凭据"
//     })
//     state.svcList=tData.data.data
//   }catch (e){
//     console.log(e)
//   }
// }
//
state.name=route.query.name.toString()
state.namespace=route.query.namespace.toString()
async function getDataItem(){
  try {

    let tData=await ingressItem(state.namespace,state.name)
    tData.data.data.spec.rules.forEach((item,index)=>{

    })
    state.form=tData.data.data
    mateDataRef.value.setData(state.form.metadata)
    yamlRef.value.setData(state.form)
    backendRef.value.setData(state.form.spec.defaultBackend)
  }catch (e){
    console.log(e)
  }
}
getDataItem()
// function postNew(){
//
//   formRef.value.validate(async (valid) => {
//     if (valid) {
//       let flag=await mateDataRef.value.Check()
//       if(!flag){
//         return
//       }
//       try {
//         let result=await  ingressUpdate(state.form)
//         if (result.data.code==200){
//           ElMessage("ingress资源创建成功")
//         }
//       }catch (e){
//         console.log(e)
//       }
//     }
//   })
// }
// function addTls(){
//   state.form.spec.tls.push({
//     hosts:[],
//     secretName:""
//   })
// }
// function addPathCfg(index){
//   state.type[state.form.spec.rules.length]=[]
//   state.type[state.form.spec.rules.length][0]="service"
//   state.form.spec.rules.push(
//       {
//         host:"",
//         http:{
//           paths:[{
//             path:"/",
//             pathType:"Prefix",
//             backend:{
//               service:{
//                 name:"",
//                 port:{
//                   name:"",
//                   number:0// 2 选一
//                 }
//               },
//
//             }
//           }]
//         }
//       }
//   )
// }
// function rmPathCfg(ruleIndex){
//   state.form.spec.rules.splice(ruleIndex,1)
// }
// function clearSelect(){
//   state.form.spec.tls=[]
//   state.form.spec.rules=[]
//   state.tlsList=[]
//   state.svcList=[]
// }
// function InitDefaultBackend(flag){
//   if(flag){
//     state.defaultType='service'
//     state.form.spec.defaultBackend={
//       service:{
//         name:"",
//         port:{
//           name:"",
//           number:0// 2 选一
//         }
//       }
//     }
//   }else{
//     delete state.form.spec.defaultBackend
//   }
//
// }
//
//
// function removeHost(index,index2){
//   state.form.spec.tls[index].hosts.splice(index2,1)
// }
// function addHost(index){
//   if (!state.host[index]){
//     return
//   }
//   state.form.spec.tls[index].hosts.push( state.host[index])
//   state.host[index]=[]
// }
// function removeTls(index){
//   state.form.spec.tls.splice(index,1)
// }
// function removePaths(index,index2) {
//   state.form.spec.rules[index].http.paths.splice(index2,1)
// }
// function fixdata(index,index2){
//   if(state.type[index][index2]=='service'){
//     state.form.spec.rules[index].http.paths[index2].backend= {service:{
//         name:"",
//         port:{
//           name:"",
//           number:""// 2 选一
//         }
//       }
//     }
//   }else {
//     state.form.spec.rules[index].http.paths[index2].backend= {resource:{
//         apiGroup:"",
//         kind:"",
//         name:""
//       }
//     }
//   }
//
// }
// function getsvcnameList(svc){
//   let ports=state.svcList.filter((item)=>{
//     return item.name==svc
//   })
//   if(ports.length){
//     return ports[0].port
//   }
//   return  []
// }
// function addPaths(index){
//   state.type[index][state.form.spec.rules[index].http.paths.length]="service"
//   state.form.spec.rules[index].http.paths.push({
//     path:"/",
//     pathType:"Prefix",
//     backend:{
//       service:{
//         name:"",
//         port:{
//           name:"",
//           number:0// 2 选一
//         }
//       },
//     }
//   })
// }
// watch(()=>state.defaultType,()=>{
//   if(state.defaultType=='service'){
//     delete  state.form.spec.defaultBackend.resource
//     if (!state.form.spec.defaultBackend.service){
//       state.form.spec.defaultBackend={service:{
//           name:"",
//           port:{
//             name:"",
//             number:""// 2 选一
//           }
//         }
//       }
//     }
//
//   }else{
//     delete  state.form.spec.defaultBackend.service
//     if (!state.form.spec.defaultBackend.resource){
//       state.form.spec.defaultBackend= {resource:{
//           apiGroup:"",
//           kind:"",
//           name:""
//         }
//       }
//     }
//   }
// })

</script>
