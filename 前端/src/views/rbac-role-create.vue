<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>Role创建</el-breadcrumb-item>
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
         <el-input v-model="form.name"   placeholder="角色名称"></el-input>
       </el-form-item>
       <el-form-item label="命名空间" prop="name_space">
         <el-select  v-model="form.name_space">
           <el-option v-for="ns in nsList "
                      :label="ns.name"
                      :value="ns.name"/>
         </el-select>
       </el-form-item>
       <el-form-item >
           <el-button style="margin-left:20px"  type="primary" @click="post">创建角色</el-button>
       </el-form-item>
     </el-card>
       <Resources ref="ResourcesRef"></Resources>
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
import {getRoleList, roleCreate} from "../api/token/rbac";
import Resources from "../components/Resources/Resources.vue";

export default defineComponent({
  name: 'role-create',
  components: {MainLayout,Resources},
  setup(){

    let state=reactive({
      nsList:reactive([]),
      form:reactive({
        name:"",
        name_space:"default",
      }),
    })
    let rules={
      name:[
        {
          required: true,
          message: '角色名必须填写',
        }
      ],
      name_space:[
        {
          required: true,
          message: '命名空间必须填写',
          trigger: 'blur',
        }
      ],
    }
    let ResourcesRef=ref(null)
    let formRef=ref(null)
    async function getData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
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
            let result= await  roleCreate(postData)
            if (result.data.code==200){
              ElMessage("角色资源创建成功")
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