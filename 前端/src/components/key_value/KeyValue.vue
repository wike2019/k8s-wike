<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header" style="padding-bottom:10px">
        自定义数据 <el-button  size="small" style="margin-left:20px" @click="importFile" >从文件导入</el-button>   <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfg" >+</el-button>
      </div>
    </template>
    <el-form :inline="true"  :model="state.from" ref="formRef"  >
    <div v-for="(item,index) in state.from.keyValue" style="margin:20px 0">
      <el-form-item label="key"   :key="index"  :prop="'keyValue.' + index + '.key'"
                    :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
        <el-input v-model="item.key"   placeholder="填写key"></el-input>
      </el-form-item>

        <el-form-item label="value"    :key="index"  :prop="'keyValue.' + index + '.value'"
                      :rules="!item.NoRequired?{}:{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input  type="textarea"  :autosize="true" style="width:600px"   v-model="item.value"  placeholder="填写value"></el-input>
        </el-form-item>
      <el-form-item label="是否必填" :key="index"  :prop="'keyValue.' + index + '.NoRequired'" style="margin-right:20px" >
        <el-checkbox v-model="item.NoRequired"></el-checkbox>
      </el-form-item>
        <el-form-item  >
          <el-button  type="danger" @click="rmPathCfg(index)" >-</el-button>
        </el-form-item>
    </div>
    </el-form>
  </el-card>
</template>

<script lang="ts" setup>
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {arrToMap, MapToArr} from "../../helper/helper";
let state=reactive({
  from:{
    keyValue:reactive([])
  }
})
const props = defineProps(['base64'])

let formRef=ref(null)
const emit = defineEmits(['input'])
watch(()=>state.from.keyValue,()=>{
  emit("input",arrToMap(state.from.keyValue))
},{deep:true,flush:"post"})
function addPathCfg(){
  state.from.keyValue.push({key:"",value:"",NoRequired:true})
}
function rmPathCfg(index){
  state.from.keyValue.splice(index,1)
}
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
function setData(data){
  state.from.keyValue= MapToArr(data,props.base64)
}
function importFile(){
  let fileDom= document.createElement("input")
  fileDom.accept=".txt"
  fileDom.type="file"
  fileDom.multiple=true
  fileDom.onchange=function (e) {
    if(e.target.files.length>0){
      for(let i=0;i<e.target.files.length;i++){
        const reader=new FileReader()
        reader.filename=e.target.files[i].name
        reader.readAsText(e.target.files[i],'UTF-8')
        reader.onload=(e)=>{
          if (state.from.keyValue.length===1 && state.from.keyValue.key===""){
            state.from.keyValue[0]={key:e.target.filename,value:e.target.result}
          }else{
            for (let k=0;k<state.from.keyValue.length;k++){
              if(state.from.keyValue[k].key==e.target.filename){
                state.from.keyValue[k].value=e.target.result
                return
              }
            }
            state.from.keyValue.push({key:e.target.filename,value: e.target.result,NoRequired:false})
          }
        }

      }
    }


  }
  fileDom.click();
}
defineExpose({ setData,Check })

</script>

<style scoped>

</style>
