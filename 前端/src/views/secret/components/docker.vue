<template>
  <div >
    <el-alert title="为了统一docker密文输入必须符合可视化规范，以便后台处理" type="success" />
    <el-form ref="formRef" :model="state.form"  class="mtb20" >
      <el-form-item label="server" label-width="100px" :prop="'server'"  :rules="requireRules('server必须填写')" >
        <el-input v-model="state.form.server"   placeholder="用户名" autocomplete="off" ></el-input>
      </el-form-item>
      <el-form-item label="用户名"  label-width="100px" :prop="'username'"  :rules="requireRules('用户名必须填写')" >
        <el-input v-model="state.form.username"   placeholder="用户名" autocomplete="off" ></el-input>
      </el-form-item>
      <el-form-item label="密码"  label-width="100px" :prop="'password'" :rules="requireRules('密码必须填写')" >
        <el-input v-model="state.form.password"   type="password"  placeholder="密码" autocomplete="off" ></el-input>
      </el-form-item>
      <el-form-item label="邮箱"   label-width="100px" :prop="'email'" :rules="requireRules('邮箱必须填写')" >
        <el-input v-model="state.form.email"   placeholder="邮箱" autocomplete="off" ></el-input>
      </el-form-item>
    </el-form>
  </div>

</template>

<script setup>
import {reactive, ref, watch} from "vue";
import {requireRules} from "../../../helper/rules";
import {IsDirty} from "../../../helper/helper";

let state=reactive({
  form:{
    server:"",
    username:"",
    password:"",
    email:""
  }
})
let formRef=ref(null)
function Check(){
  return new Promise(function (resolve){
    formRef.value.validate(async (valid) => {
      if (valid) {
        resolve(true)
      }
      resolve(false)
    })
  })
}
defineExpose({ setData,Check })
const emit = defineEmits(['input'])
watch(()=>state.form,()=>{
  if(!IsDirty(state)) {
    let data= {
      [state.form.server]:{
        "auth":btoa(state.form.username+":"+state.form.password),
        "email":state.form.email
     }
    }
    emit("input",{".dockercfg":btoa(JSON.stringify(data))})
  }
},{deep:true,flush:"post"})
function setData(data){
 let info= JSON.parse(atob(data['.dockercfg']))
  for (let key in info){
    state.form.server=key
    state.form.email=info[key].email
    let user=atob(info[key].auth)
    let infoData=user.split(":")
    state.form.username=infoData[0]
    state.form.password=infoData[1]
  }
}
</script>

<style scoped>

</style>