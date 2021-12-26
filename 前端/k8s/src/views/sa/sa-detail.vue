<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>服务账号详情</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="list_item">
       <p>
         账号服务名称: {{item.name}}
       </p>
       <p>
         命名空间: {{item.name_space}}
       </p>
       <p>
       创建时间: {{item.create_time}}
       </p>
       <el-divider></el-divider>
       <h3>标签</h3>
       <div v-if="item.labels">
         <p v-for="(key,value) in item.labels">
          <span>{{key}}</span>:<span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>注释</h3>
       <div v-if="item.annotations">
         <p v-for="(key,value) in item.annotations">
           <span>{{key}}</span>:<span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <el-divider></el-divider>
       <h3>密文列表</h3>
       <div v-if="item.secrets&&item.secrets.length">
          <p v-for="secret in item.secrets">
            密文名称: <a  @click="()=>doTo('secret-detail',{name_space:item.name_space,name:secret.name})">{{secret.name}} 查看密文</a>
          </p>
       </div>
       <div v-else>
         无数据
       </div>
       <el-divider></el-divider>
       <h3>拉取镜像列表</h3>
       <div v-if="item.imagePullSecrets&&item.imagePullSecrets.length">
         <p v-for="imagePullSecret in item.imagePullSecrets">
           镜像密文名称:<a  @click="()=>doTo('secret-detail',{name_space:item.name_space,name:imagePullSecret.name})">{{imagePullSecret.name}} 查看密文</a>
         </p>
       </div>
       <div v-else>
           无数据
       </div>
     </div>
     <el-divider></el-divider>
     <el-button type="primary" @click="()=>doTo('sa-update',{name_space:name_space,name:name})" >
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
import {secretDetail} from "../../api/token/secret";
import {roleCreate} from "../../api/token/rbac";
import {useRoute} from "vue-router";

export default defineComponent({
  name: 'sa-detail',
  components: {MainLayout},
  setup(){
    let state=reactive({
      item:{},
      name:"",
      name_space:""
    })
    const route = useRoute()
    state.name=route.query.name
    state.name_space=route.query.name_space
    async function getData(){
      try {
       let tData=await getSaItem(state.name_space,state.name)
        state.item=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    getData()

    async function showToken(name){
      try {
      let result=await secretDetail(state.name_space,name)
      let token=window.atob(result.data.data.data['token'])
      state.token=token;
      state.dialogVisible=true;
      }catch (e) {
        ElMessage({
          type: 'info',
          message: '系统异常'+e,
        })
      }

    }

    return {...toRefs(state),doTo,showToken}
  }
})
</script>
<style >

</style>