<template>
  <div>
    <el-alert title="为了统一密文输入必须符合可视化规范，以便后台处理" type="success" />
    <el-form ref="formRef" :model="state.form"   class="mtb20">
      <el-form-item :prop="'tls.key'"    :rules="requireRules('私钥内容必须填写')" >
        <h2 class="sub_title" >私钥内容 tls.key</h2>
        <el-input v-model="state.form['tls.key']" type="textarea" :rows="8" style="width:80%"  placeholder="-----BEGIN RSA PRIVATE KEY-----开头的内容"></el-input>
        <el-button style="margin-left:15px;vertical-align: top" type="primary" @click="()=>fileLoad('tls.key')">文件上传</el-button>
      </el-form-item>
      <el-divider></el-divider>
      <el-form-item  :prop="'tls.crt'"   :rules="requireRules('证书内容必须填写')" >
        <h2 class="sub_title"  >证书内容 tls.crt</h2>
        <el-input v-model="state.form['tls.crt']" type="textarea" :rows="8"  style="width:80%" placeholder="-----BEGIN CERTIFICATE-----开头的内容"></el-input>
        <el-button style="margin-left:15px;vertical-align: top" type="primary" @click="()=>fileLoad('tls.crt')">文件上传</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import {reactive, ref, watch} from "vue";
import {requireRules} from "../../../helper/rules";
import {arrToMap, IsDirty} from "../../../helper/helper";
const emit = defineEmits(['input'])

let state=reactive({
  form:{
      "tls.crt":"",
      "tls.key":""
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



watch(()=>state.form,()=>{
  if(!IsDirty(state)) {
    emit("input",state.form)
  }
},{deep:true,flush:"post"})

function fileLoad(key){
  let fileDom= document.createElement("input")
  fileDom.accept=".txt,.pem,.key,.pkey,.crt"
  fileDom.type="file"
  fileDom.onchange=function (e) {
    const reader=new FileReader()
    reader.readAsText(e.target.files[0],'UTF-8')
    reader.onload=(e)=>{
      state.form[key]=e.target.result
    }
  }
  fileDom.click();
}
function setData(data){
  state.form={
    "tls.crt":atob(data['tls.crt']),
    "tls.key":atob(data['tls.key'])
  }
}
defineExpose({ setData,Check })
</script>

<style scoped>

</style>