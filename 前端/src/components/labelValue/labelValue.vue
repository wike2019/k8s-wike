<template>
  <div  v-if="!render">
      <div class="inline-title-no" v-if="!label">
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
  </div>
  <div v-if="render">
    <div class="inline-title-no-back" v-if="!label">
      <span>lebal数据</span>
    </div>
    <el-table  border :data="form.data" class="smallTable"  size="small" empty-text="无数据">
      <el-table-column prop="key" label="key" width="300" />
      <el-table-column prop="value" label="值" />
    </el-table>
  </div>
</template>

<script lang="ts">
import {defineComponent, inject, reactive, ref, toRefs, watch} from 'vue'
import {arrToMap, MapToArr,IsDirty} from "@/helper/helper";
import {hasValues} from "../../helper/helper";
export default defineComponent({
  name: 'LabelValue',
  emits:["input"],
  props:["value","label"],
  setup(props,{emit}){
    const render = inject("render")
    let state=reactive({
        form:{
          data:[]
        },
       label:props.label,
       md5:"",
      render:render,
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
      if(state.label){
        if(IsDirty(state,{data:props.value},false,2.2)){
          return
        }
      }else{
        if(IsDirty(state,{data:MapToArr(props.value)},false,2)){
          return
        }
      }
      commit()
    },{deep:true,flush:"post"})

    watch(()=>props.value,()=>{
      if(!hasValues(props.value)){
        state.form.data=[]
      }
      if(state.label){
        state.form.data=props.value
      }else{
        state.form.data=MapToArr(props.value)
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
