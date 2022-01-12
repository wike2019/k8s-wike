<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header" style="padding-bottom:10px">
        自定义数据   <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfg" >+</el-button>
      </div>
    </template>
    <el-form :inline="true"  :model="from" ref="formRef"  >
    <div v-for="(item,index) in from.keyValue" style="margin:20px 0">
      <el-form-item label="key"   :key="index"  :prop="'keyValue.' + index + '.key'"
                    :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
        <el-input v-model="item.key"   placeholder="填写key"></el-input>
      </el-form-item>
        <el-form-item label="value"    :key="index"  :prop="'keyValue.' + index + '.value'"
                      :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input   style="width:400px"   v-model="item.value"  placeholder="填写value"></el-input>
        </el-form-item>
        <el-form-item label="影响" :key="index"  :prop="'keyValue.' + index + '.effect'" :rules="{required: true, trigger: 'change', message: '影响必须选择' }">
        <el-select v-model="item.effect"  >
          <el-option label="不调度" value="NoSchedule"/>
          <el-option label="尽量不调度" value="PreferNoSchedule"/>
          <el-option label="不调度+驱逐" value="NoExecute"/>
        </el-select>
      </el-form-item >
        <el-form-item  >
          <el-button  type="primary" @click="rmPathCfg(index)" >-</el-button>
        </el-form-item>
    </div>
    </el-form>
  </el-card>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
export default defineComponent({
  name: 'Mate',
  emits:["input"],
  setup(props,{emit}){
    let state=reactive({
      from:{
        keyValue:reactive([])
      }
    })
    let formRef=ref(null)
    watch(state.from.keyValue,()=>{
      emit("input",state.from.keyValue)
    },{deep:true,flush:"sync"})
    function addPathCfg(){
      state.from.keyValue.push({key:"",value:"",effect:"NoSchedule"})
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
      state.from.keyValue=data
      emit("input",state.from.keyValue)
    }

    return {rmPathCfg,addPathCfg,formRef,...toRefs(state),Check,setData}
  }
})
</script>

<style scoped>

</style>
