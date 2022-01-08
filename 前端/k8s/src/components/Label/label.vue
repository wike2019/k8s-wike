<template>
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div v-for="(Mate,index) in form.MateList" style="margin:20px 0">
      <template v-if="Mate.name">
        <span class="name_mate">
          <em v-if="Mate.required" class="Danger">*</em>
          <span>{{Mate.name}}</span>
        </span>
      </template>
      <template v-else >
        <el-form-item label="key"   :key="'key'+index"  :prop="'MateList.' + index + '.key'"
                      :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input v-model="Mate.key"   placeholder="填写key"></el-input>
        </el-form-item>
      </template>
        <el-form-item :label="Mate.name?'':'value'"   :key="'value'+index"  :prop="'MateList.' + index + '.value'"
                      :rules="Mate.required?{required: true, trigger: 'blur', message: 'value必须填写' }:{}">
          <el-input  style="width:600px"   type="textarea"  :autosize="true"   v-model="Mate.value"  placeholder="填写value"></el-input>
        </el-form-item>
        <el-form-item label="是否必填"   :key="'required'+index"  :prop="'MateList.' + index + '.required'" style="margin-right:20px">
          <el-checkbox v-model="Mate.required"></el-checkbox>
        </el-form-item>
        <el-form-item  >
          <el-button  type="danger" size="small" @click="rmPathCfg(index)" >-</el-button>
        </el-form-item>
    </div>
  </el-form>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch, watchEffect} from 'vue'
export default defineComponent({
  name: 'Mate',
  emits:["input"],
  setup(props,{emit}){
    let state=reactive({
        form:reactive({
          MateList:[]
        })
      })


    let formRef=ref(null)
    watch(()=>state.form.MateList,()=>{
      emit("input",state.form.MateList)
    },{deep:true,flush:"post"})
    function addPathCfg(){
      state.form.MateList.push({key:"",value:"",name:"",required:true})
    }
    function rmPathCfg(index){
      state.form.MateList.splice(index,1)
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

      state.form.MateList=data
      emit("input",state.form.MateList)
    }
    return {rmPathCfg,addPathCfg,...toRefs(state),Check,setData,formRef}
  }
})
</script>

<style scoped>

</style>
