<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>pod列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
                  <el-table
                      :data="nodeList"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">
                    <el-table-column
                        prop="status"
                        label="状态"
                        width="80">
                      <template #default="scope">
                        <span  class="Success"  v-if="scope.row.is_ready=='Ready'">active</span>
                        <span  class="Danger" v-else>waiting</span>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="name"
                        label="节点名称"
                        width="250">
                      <template #default="scope">
                        <p>{{scope.row.name}}</p>
                        <el-tag size="mini" type="info" v-for="label in scope.row.labels">{{label}}</el-tag>
                      </template>
                    </el-table-column>
                    <el-table-column
                        label="污点列表"
                        width="250">
                      <template #default="scope">
                        <div v-if="scope.row.taints.length">
                          <el-tag size="mini" type="danger" v-for="item in scope.row.taints">{{item}}</el-tag>
                        </div>
                        <div v-else>无污点</div>
                      </template>
                    </el-table-column>
                    <el-table-column
                        label="节点信息">
                      <template #default="scope">
                        <p>{{scope.row.ip}}</p>
                        <p>{{scope.row.host_name}}</p>
                      </template>
                    </el-table-column>
                    <el-table-column
                        label="资源指标">
                      <template #default="scope">
                        <p>CPU指标: <span style="margin-left:10px">{{ Math.round(scope.row.Usage.Cpu*100) }}%  / {{scope.row.Capacity.Cpu }}核</span></p>
                        <p>内存指标: <span style="margin-left:10px"> {{ Math.round(scope.row.Usage.Memory*100) }}% / {{ Math.round(scope.row.Capacity.Memory/1024/1024/1024) }}G</span></p>
                        <p>POD指标: <span style="margin-left:10px">{{ scope.row.Usage.Pods }} / {{ scope.row.Capacity.Pods }}</span></p>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="create_time"
                        label="创建时间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column label="操作" width="280" align="center">
                      <template #default="scope">
                        <div style="margin-bottom:10px">
                          <el-button size="small"  @click="()=>setPassword(scope.row.ip)" >设置shell密码 (<span :id="'node_'+scope.row.ip">{{config(scope.row.ip)}}</span>)</el-button>
                        </div>
                        <el-button size="small" @click="()=>shell(scope.row.ip,scope.row.name)" >执行shell</el-button>
                        <el-button size="small" @click="()=>doTo('node-update',{name:scope.row.name})" >编辑<i class="el-icon-s-platform el-icon--right"></i></el-button>
                      </template>
                    </el-table-column>
                  </el-table>
     </div>
   </main-layout>
  <el-dialog :append-to-body="true" v-model="dialogTableVisible" title="设置节点shell登录密码">
    <el-form
        :model="ruleForm"
        :rules="rules"
        ref="formRef"
        label-width="120px"
    >
      <el-form-item label="节点ip" prop="ip">
        <el-input v-model="ruleForm.ip" disabled ></el-input>
      </el-form-item>
      <el-form-item label="登录用户名" prop="name">
        <el-input v-model="ruleForm.name"></el-input>
      </el-form-item>
      <el-form-item label="登录密码" prop="password">
        <el-input v-model="ruleForm.password" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button style="float:right;margin-right:20px" type="primary" @click="submitForm()"
        >确定</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {getPodsByNs} from "../api/token/pod";
import {doTo} from "../router";
import {getNodes} from "../api/token/node";
import {ElMessage} from "element-plus";
export default defineComponent({
  name: 'node',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nodeList:reactive([]),
      dialogTableVisible:false,
      ruleForm:reactive({
        ip:"",
        password:"",
        name:""
      })
    })
    let formRef=ref(null)
    let rules = {
      name: [
        {
          required: true,
          message: '用户名必须输入',
          trigger: 'blur',
        },],
      password: [
        {
          required: true,
          message: '密码必须输入',
          trigger: 'blur',
        },]
    }

    async function getData(){
      try {
       let tData=await getNodes()
        state.nodeList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }

    getData()

    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        if(tData.type=='node'){
          state.nodeList=tData.result
        }
      }
    }
    function  config(ip){
      if(localStorage.getItem(ip)){
        return "已设置"
      }else{
        return "未设置"
      }
    }
    function setPassword(ip){

      state.dialogTableVisible=true
      state.ruleForm.ip=ip
      console.log(state.ruleForm.ip)
    }
    function submitForm(){
      formRef.value.validate((valid) => {
        if (valid) {
          localStorage.setItem(state.ruleForm.ip,JSON.stringify(state.ruleForm))
          document.getElementById("node_"+state.ruleForm.ip).innerHTML=config( state.ruleForm.ip)
          state.ruleForm={
            ip:"",
            name:"",
            password:""
          }

          state.dialogTableVisible=false
          ElMessage("数据保存成功")
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
    function shell(ip,name){
      if(!localStorage.getItem(ip)){
        ElMessage("节点shell密码没有设置,请先设置")
        setPassword(ip)
        return
      }
      doTo('node-shell',{"name":name})
    }
    return {...toRefs(state),rules,doTo,config,setPassword,submitForm,formRef,shell}
  }
})
</script>
