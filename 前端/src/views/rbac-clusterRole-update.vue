<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>clusterRole更新</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-form :inline="true" rel="formRef" :model="form" ref="formRef" :rules="rules">
     <el-card class="box-card">
       <template #header>
         <div class="card-header">
           基本配置
         </div>
       </template>
       <el-form-item label="角色名称"  prop="name"  >
         <el-input v-model="form.name"  disabled  placeholder="角色名称"></el-input>
       </el-form-item>
       <el-form-item >
           <el-button style="margin-left:20px"  type="primary" @click="post">修改内容</el-button>
       </el-form-item>
     </el-card>
       <Resources ref="ResourcesRef" ></Resources>
     </el-form>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {ingressDel, ingressListByNs} from "../api/token/ingress/ingress";
import {secretDel, secretListByNs} from "../api/token/secret/secret";
import {doTo} from "../router";
import {
  clusterRoleUpdate,
  getClusterRoleDetail,
  getRoleDetail,
  getRoleList,
  roleCreate,
  roleUpdate
} from "../api/token/rbac";
import Resources from "../components/Resources/Resources.vue";
import {useRoute} from "vue-router";
import {configmapDetail} from "../api/token/configmap/configmap";

export default defineComponent({
  name: 'clusterRole-create',
  components: {MainLayout,Resources},
  setup(){

    let state=reactive({
      nsList:reactive([]),
      form:reactive({
        name:"",
      }),
      defaultKey:[],
    })
    let route = useRoute()
    state.form.name=route.query.name
    let rules={
      name:[
        {
          required: true,
          message: '角色名必须填写',
        }
      ],
    }
    let ResourcesRef=ref(null)
    let formRef=ref(null)
    async function getData(){
      try {
        let tData2=await getClusterRoleDetail(state.form.name)
        let roleInfo=tData2.data.data
        console.log(roleInfo)
        roleInfo.rules.forEach((item)=>{
          item.apiGroups.forEach(function (apiGroup) {
            if (apiGroup==""){
              apiGroup="core"
            }
            item.resources.forEach(function (resource) {
              item.verbs.forEach(function (verb) {
                state.defaultKey.push(apiGroup+"@"+resource+"@"+verb)
              })
            })
          })
        })
        ResourcesRef.value.setData(state.defaultKey)
      }catch (e){
        console.log(e)
      }
    }

    getData()
    function post() {
      formRef.value.validate(async (valid) => {
        if (valid) {
          let data=ResourcesRef.value.getCheckedKeys()
          let postData={metadata:{name:state.form.name,namespace:state.form.name_space},rules:data}
          console.log(postData)
          try {
            let result= await  clusterRoleUpdate(state.form.name,postData)
            if (result.data.code==200){
              ElMessage("集群角色资源修改成功")
            }

          }catch (e){
            console.log(e)
          }

        }
      })
    }
    return {...toRefs(state),rules,doTo,ResourcesRef,post,formRef}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>