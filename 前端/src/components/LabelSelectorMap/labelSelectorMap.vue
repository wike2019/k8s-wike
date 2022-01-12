<template>

  <div class="inline-title">
    <span>标签选择器</span>         <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addPathCfg" >添加标签</el-button>
  </div>
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div v-for="(Mate,index) in form.MateList" style="margin:20px 0">

        <el-form-item label="key"   :key="index"  :prop="'MateList.' + index + '.key'"
                      :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input v-model="Mate.key"   placeholder="填写key"></el-input>
        </el-form-item>
      <el-form-item :label="Mate.name?'':'value'"   :key="index"  :prop="'MateList.' + index + '.value'"
                    :rules="{required: true, trigger: 'blur', message: 'value必须填写' }">
        <el-input  style="width:600px"     v-model="Mate.value"  placeholder="填写value"></el-input>
      </el-form-item>
      <el-form-item  >
        <el-button  type="danger" size="small" @click="rmPathCfg(index)" >删除此项</el-button>
      </el-form-item>
    </div>
  </el-form>
  <el-divider></el-divider>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch, watchEffect} from 'vue'
import {core} from "../../core/core";
import {data2arr, MapToArr} from "../../helper/helper.ts"
export default defineComponent({
  name: 'Mate',
  emits:["input"],
  setup(props,{emit}){
    let state=reactive({
      form:reactive({
        MateList:[]
      }),
    })
    let formRef=ref(null)
    function commit(){
      let arr={}
      state.form.MateList.forEach(function (item){
        arr[item.key]=item.value.toString()||""
      })
      emit("input",arr)
    }
    watch(()=>state.form.MateList,()=>{
      commit()
    },{deep:true,flush:"post"})
    function addPathCfg(){
      state.form.MateList.push({key:"",value:"",required:true})
    }

    function addValue(index) {
      console.log( state.MatchExpressions.MateList[index].value)
      state.MatchExpressions.MateList[index].value.push("")
    }
    function rmValue(index,valueIndex) {
      state.MatchExpressions.MateList[index].value.splice(valueIndex,1)
      if(state.MatchExpressions.MateList[index].value.length==0){
        state.MatchExpressions.MateList.splice(index,1)
      }
    }
    function rmPathCfg(index){
      state.form.MateList.splice(index,1)
    }

    function Check(){
      return new Promise(function (resolve){
        formRef.value.validate(async (valid) => {
          if (valid) {
            resolve(true)
          }else{
            resolve(false)
          }
        })
      })

    }
    function setData(data){
      //alert(JSON.stringify(data))
      try {

          state.form.MateList=MapToArr(data )
          commit()

        return true
      }catch (e) {
        console.log(e)
        return false
      }

    }
    let isInArr=core.isInArr
    return {isInArr,addValue,rmValue,rmPathCfg,addPathCfg,...toRefs(state),Check,setData,formRef}
  }
})
</script>

<style scoped>

</style>
