<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>ingress详情</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="list_item" v-if="item.metadata" >
       <p>
         ingress名称: {{item.metadata.name}}
       </p>
       <p>
         命名空间: {{item.metadata.namespace}}
       </p>
       <p>
         创建时间: {{item.metadata.creationTimestamp}}
       </p>
       <el-divider></el-divider>
       <h3>标签</h3>
       <div v-if="item.metadata.labels">
         <p v-for="(value,key) in item.metadata.labels">
          <strong>{{key}}</strong><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>注释</h3>
       <div v-if="item.metadata.annotations">
         <p v-for="(value,key) in item.metadata.annotations">
           <strong>{{key}}</strong><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <el-divider></el-divider>
       <h3>默认后台</h3>
       <div v-if="item.spec.defaultBackend">
         <template v-if="item.spec.defaultBackend.service">
          <h4>service资源</h4>
          <p><strong>service名称</strong><em>:</em><span>{{item.spec.defaultBackend.service.name}}</span></p>
          <p><strong>端口</strong><em>:</em><span >{{item.spec.defaultBackend.service.port.name?item.spec.defaultBackend.service.port.name:item.spec.defaultBackend.service.port.number}}</span></p>
         </template>
         <template v-if="item.spec.defaultBackend.resource">
           <h4>自定义资源</h4>
           <p><strong>apiGroup</strong><em>:</em><span>{{item.spec.defaultBackend.resource.apiGroup}}</span></p>
           <p><strong>kind</strong><em>:</em><span>{{item.spec.defaultBackend.resource.kind}}</span></p>
           <p><strong>name</strong><em>:</em><span>{{item.spec.defaultBackend.resource.name}}</span></p>
         </template>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>路由规则</h3>
       <div v-if="item.spec.rules.length">
         <template v-for="rule in item.spec.rules">

           <p><strong>host</strong><em>:</em><span>{{rule.host}}</span></p>
           <div class="dashed"></div>
           <div v-for="path in rule.http.paths">
             <p><strong>path</strong><em>:</em><span>{{path.path}}</span></p>
             <p><strong>pathType</strong><em>:</em><span>{{path.pathType}}</span></p>
             <div v-if="path.backend">
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
             <div class="dashed"></div>
           </div>
         </template>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>证书</h3>
       <div v-if="item.spec.tls&&item.spec.tls.length">
         <template v-for="value in item.spec.tls">
           <h4>证书</h4>
           <p>
             <strong>密文名称</strong>
             <em>:</em>
             <span>
             <a  @click="()=>doTo('secret-detail',{namespace:namespace,name:value.secretName})">{{value.secretName}} 查看密文</a>
             </span>
           </p>
           <h4>域名</h4>
           <p v-for="host in value.hosts">
             <span>{{host}}</span>
           </p>
           <div class="dashed"></div>
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

<script lang="ts">
import {defineComponent, inject, reactive, ref, toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {doTo} from "../../router";
import {getSaItem, getSaList, SACreate, SaDel} from "../../api/token/sa/sa";
import {secretDetail} from "../../api/token/secret/secret";
import {roleCreate} from "../../api/token/rbac";
import {useRoute} from "vue-router";
import {ingressItem} from "../../api/token/ingress/ingress";

export default defineComponent({
  name: 'ingress-detail',
  components: {MainLayout},
  setup(){
    let state=reactive({
      item:{},
      name:"",
      namespace:""
    })
    const route = useRoute()
    state.name=route.query.name
    state.namespace=route.query.namespace
    async function getData(){
      try {
       let tData=await ingressItem(state.namespace,state.name)
        state.item=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    getData()

    return {...toRefs(state),doTo}
  }
})
</script>
<style >

</style>