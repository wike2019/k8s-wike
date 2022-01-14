<template>
  <div class="inline-title" v-if="!label">
    <span>Label数据设置</span>         <el-button type="primary" size="small"  class="leftBtn" @click="addItem" >添加标签</el-button>
  </div>
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div v-for="(Mate,index) in form.data"  class="mtb10 flex">
      <template v-if="Mate.name">
        <div class="el-form-item" :class="{'is-required':!!Mate.required}">
          <label class="el-form-item__label" >
            <span> {{Mate.name}}</span>
          </label>
        </div>
      </template>
      <template v-else >
        <el-form-item label="key"   :key="'key'+index"  :prop="'data.' + index + '.key'"
                      :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input v-model="Mate.key"   placeholder="填写key"></el-input>
        </el-form-item>
      </template>
        <el-form-item class="flex1"  :label="Mate.name?'':'value'"   :key="'value'+index"  :prop="'data.' + index + '.value'"
                      :rules="Mate.required?{required: true, trigger: 'blur', message: 'value必须填写' }:{}">
          <el-input      v-model="Mate.value"  placeholder="填写value"></el-input>
        </el-form-item>
        <el-form-item label="是否必填"   :key="'required'+index"  :prop="'data.' + index + '.required'" >
          <el-checkbox v-model="Mate.required"></el-checkbox>
        </el-form-item>
        <el-form-item class="m40"  >
          <el-button  type="danger" size="small" @click="removeItem(index)" >删除此项</el-button>
        </el-form-item>
    </div>
  </el-form>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {arrToMap, MapToArr,IsDirty} from "@/helper/helper";
export default defineComponent({
  name: 'LabelValue',
  emits:["input"],
  props:["value","label"],
  setup(props,{emit}){
    let state=reactive({
        form:{
          data:[]
        },
       label:props.label,
       md5:""
      })
    let formRef=ref(null)
    function commit(){
      if(state.label){
        emit("input",state.form.data)
      }else{
        emit("input",arrToMap(state.form.data))
      }

    }
    watch(()=>state.form,()=>{
      if(IsDirty(state,3)){
        return
      }
      commit()
    },{deep:true,flush:"post"})

    watch(()=>props.value,()=>{
      if(state.label){
        if(props.value){
          state.form.data=props.value
        }else{
          state.md5=""
          state.form.data={}
        }
      }else{
        if(props.value){
          state.form.data=MapToArr(props.value)

        }else{
          state.md5=""
          state.form.data=[]
        }
      }
    },{deep:true,flush:"post"})

    function addItem(){
      state.form.data.push({key:"",value:"",name:"",required:true})
    }
    function removeItem(index){
      state.form.data.splice(index,1)
    }
    function Check(){
      formRef.value.validate()
    }

    return {removeItem,addItem,...toRefs(state),Check,formRef}
  }
})
</script>

<style scoped>

</style>
