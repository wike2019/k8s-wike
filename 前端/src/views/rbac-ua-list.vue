<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>用户列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-card class="box-card">
         <template #header>
           <div class="card-header">
             快捷操作
           </div>
         </template>
         <el-form  :inline="true" :rules="rules" :model="form" ref="formRef" >
           <el-form-item  label="用户名" prop="name">
             <el-input   v-model="form.name" placeholder="必填,至少两个字符,同名用户会被覆盖"></el-input>
           </el-form-item>
           <el-form-item label="组" prop="group">
             <el-input   v-model="form.group" placeholder="选填"></el-input>
           </el-form-item>
           <el-form-item  >
             <el-button  type="primary" @click="post">提交</el-button>
           </el-form-item>

         </el-form>
     </el-card>
     <el-divider></el-divider>
     <el-form   >
       <el-form-item label="后端的APIServer地址,根据需要填写:"  >
            <el-input  style="width:260px"  v-model="api" ></el-input>
            <el-button style="margin-left:30px;margin-top:-5px" @click="remember">记忆APIServer地址</el-button>
       </el-form-item>
     </el-form>
     <el-divider></el-divider>

     <div class="top-list">

                  <el-table
                      :data="List"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">

               <el-table-column label="用户名称" >
                 <template #default="scope">
                   <p >{{ scope.row.name }} </p>
                 </template>
               </el-table-column>
                <el-table-column label="创建时间" width="200" align="center">
                  <template #default="scope">
                    <span>{{ scope.row.create_time }}</span>
                  </template>
                </el-table-column>
               <el-table-column label="操作" width="200" align="center">
                 <template #default="scope">
                   <el-button type="danger" size="small" @click="()=>rm(scope.row.name )" >删除</el-button>
                   <el-button size="small" @click="()=>download(scope.row.name )" >下载配置文件</el-button>
                 </template>
               </el-table-column>
    </el-table>
     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, inject, reactive, ref, toRefs,nextTick} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {doTo} from "../router";
import {getSaList, SACreate, SaDel} from "../api/token/sa/sa";
import {secretDetail} from "../api/token/secret/secret";
import {roleCreate} from "../api/token/rbac";
import {core} from "../core/core.ts"
import {genConfigFile, getUaList, UaCreate, UaDel} from "../api/token/ua";

export default defineComponent({
  name: 'ua-list',
  components: {MainLayout},
  setup(){

    let state=reactive({
      List:reactive([]),
      api:localStorage.getItem('api')||"https://127.0.0.1:6443",
      form:reactive({
        name:"",
        group:"",
      }),
    })
    let formRef=ref(null)
    async function getData(){
      try {
        let tData=await getUaList()
        if (tData.data.code==200){
          state.List=tData.data.data
        }

       // state.List=tData.data
      }catch (e){
        console.log(e)
      }
    }
    getData()
    function rm(name){
      ElMessageBox.confirm(
          '你确定继续删除这个UA账户操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                let result=  await UaDel(name)
                if (result.data.code==200){
                  ElMessage("UA资源删除成功")
                }
                await getData()
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }

    function post(){
      formRef.value.validate(async (valid) => {
        if (valid) {
          try {
            let result= await   UaCreate({cn:state.form.name,o:state.form.group})
            if (result.data.code==200){
              ElMessage("UA资源创建成功")
              await getData()
              await nextTick(() => {
                state.form.name = ""
                state.form.group = ""
              })

            }

          }catch (e){
            console.log(e)
            ElMessage({
              type: 'info',
              message: '系统异常'+e,
            })
          }

        }
      })
    }
    let rules={
      name:[
        {
          required: true,
          message: '用户名必须填写',
          trigger: 'blur',
        }
      ],
    }
    async function download(user){
        try {

          let tData=await genConfigFile(user,state.api)
          if (tData.data.code==200){
            core.fileDownLoad(tData.data.data,"config")
          }

        }catch (e){
          console.log(e)
          ElMessage({
            type: 'info',
            message: '系统异常'+e,
          })
        }
    }
    function remember(){
      localStorage.setItem("api",state.api)
    }
    return {...toRefs(state),rm,doTo,post,rules,formRef,download,remember}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>