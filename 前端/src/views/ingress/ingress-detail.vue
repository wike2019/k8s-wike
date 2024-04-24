<template>
   <main-layout class="fix-black">
     <breadcrumb title="ingress.详情"></breadcrumb>
     <div class="list_item">
       <metadataInfo  ref="mateDataRef" tipTitle="ingress" :nameDisable="true" :nameRequired="true"></metadataInfo>
       <el-divider></el-divider>
       <h3>路由规则</h3>
       <div v-if="state.form.spec?.rules?.length">
         <template v-for="rule in state.form.spec.rules">
           <p><strong>host</strong><em>:</em><span>{{rule.host}}</span></p>
           <div class="dashed"></div>
           <div v-for="path in rule.http.paths">
             <p><strong>path</strong><em>:</em><span>{{path.path}}</span></p>
             <p><strong>pathType</strong><em>:</em><span>{{path.pathType}}</span></p>
             <div v-if="path.backend" style="margin-left:30px">
               <template v-if="path.backend.service">
                 <h4>service资源</h4>
                 <p><strong>service名称</strong><em>:</em><span>{{path.backend.service.name}}</span></p>
                 <p><strong>端口</strong><em>:</em><span >{{path.backend.service.port.name?path.backend.service.port.name:path.backend.service.port.number}}</span></p>
               </template>
               <template v-if="path.backend.resource">
                 <h4>自定义资源</h4>
                 <p><strong>apiGroup</strong><em>:</em><span>{{path.backend.resource.apiGroup}}</span></p>
                 <p><strong>kind</strong><em>:</em><span>{{path.backend.resource.kind}}</span></p>
                 <p><strong>name</strong><em>:</em><span>{{path.backend.resource.name}}</span></p>
               </template>
             </div>
             <el-divider></el-divider>
           </div>
         </template>
       </div>
       <h3>证书</h3>
       <div v-if="state.form.spec?.tls&&state.form.spec.tls.length">
         <template v-for="value in state.form.spec.tls">
           <h4>证书</h4>
           <p>
             <strong>密文名称</strong>
             <em>:</em>
             <span>
             <a  @click="()=>doTo('secret-detail',{namespace:state.namespace,name:value.secretName})">{{value.secretName}} 查看密文</a>
             </span>
           </p>
           <h4>域名</h4>
           <p v-for="host in value.hosts">
             <span>{{host}}</span>
           </p>
           <el-divider></el-divider>
         </template>
       </div>
       <div v-else>
         无数据
       </div>
     </div>
     <el-divider></el-divider>
     <el-button type="primary" @click="()=>doTo('ingress-update',{namespace:namespace,name:name})" >
       编辑
     </el-button>
   </main-layout>
</template>

<script lang="ts" setup>
import { computed, ref,  reactive,provide} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import {useRoute} from "vue-router";
import {configmapDetail} from "../../api/token/configmap/configmap";
import {doTo} from "../../router";
import {ingressItem} from "../../api/token/ingress/ingress";
let mateDataRef=ref(null)
const route = useRoute()
let state=reactive({
  namespace:"",name:"",
  form:reactive({
    metadata:{}
  }),
})
state.name=route.query.name.toString()
state.namespace=route.query.namespace.toString()

async function getData(){

  try {
    let tData=await ingressItem(state.namespace,state.name)
    state.form=tData.data.data
    mateDataRef.value.setData( state.form.metadata)
  }catch (e){
    console.log(e)
  }
}
getData()
provide("render",true)
</script>
<style >

</style>