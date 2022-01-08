<template>

  <div class="inline-title">
    <span><b class="Info">selector:</b> MatchLabels数据设置</span>         <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfg" >+</el-button>
  </div>
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div v-for="(Mate,index) in form.MateList" style="margin:20px 0">

        <el-form-item label="key"   :key="index"  :prop="'MateList.' + index + '.key'"
                      :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
          <el-input v-model="Mate.key"   placeholder="填写key"></el-input>
        </el-form-item>
      <el-form-item :label="Mate.name?'':'value'"   :key="index"  :prop="'MateList.' + index + '.value'"
                    :rules="Mate.required?{required: true, trigger: 'blur', message: 'value必须填写' }:{}">
        <el-input  style="width:600px"   type="textarea"  :autosize="true"   v-model="Mate.value"  placeholder="填写value"></el-input>
      </el-form-item>
      <el-form-item label="是否必填"   :key="index"  :prop="'MateList.' + index + '.required'" style="margin-right:20px">
        <el-checkbox v-model="Mate.required"></el-checkbox>
      </el-form-item>
      <el-form-item  >
        <el-button  type="primary" @click="rmPathCfg(index)" >-</el-button>
      </el-form-item>
    </div>
  </el-form>
  <el-divider></el-divider>
  <div class="inline-title">
    <span><b class="Info">selector:</b> MatchExpressions数据设置 <em class="Danger" style="margin-left:10px">(注意:可以设置多个value)</em></span>         <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addPathCfgMatchExpressions" >+</el-button>
  </div>
  <el-form :inline="true"  :model="MatchExpressions" ref="MatchExpressionsRef"  >
    <div v-for="(Mate,index) in MatchExpressions.MateList" style="margin:20px 0">

      <el-form-item label="key"   :key="index"  :prop="'MateList.' + index + '.key'"
                    :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
        <el-input v-model="Mate.key"   placeholder="填写key"></el-input>
      </el-form-item>
      <el-form-item label=" "  :rules="[{required: true, trigger: 'change', message: '匹配规则必须填写' },{ validator: isInArr(['In','NotIn','Exists','DoesNotExist'],'匹配规则不合法'), trigger: 'change' }]"  label-width="10px" :key="index"  :prop="'MateList.' + index + '.operator'">
        <el-select  size="small" v-model="Mate.operator" style="margin:0 8px;width: 180px">
          <el-option label="In" value="In"/>
          <el-option label="NotIn" value="NotIn"/>
          <el-option label="Exists" value="Exists"/>
          <el-option label="DoesNotExist" value="DoesNotExist"/>
        </el-select>
      </el-form-item>
      <el-form-item  >
        <el-button  type="primary" @click="rmPathCfgMatchExpressions(index)" >-</el-button>
        <el-button  type="primary" @click="addValue(index)">添加value</el-button>
      </el-form-item>
      <div  v-for="(values,valueIndex)  in Mate.value">
        <el-form-item :rules="{required: true, trigger: 'blur', message: 'value必须填写' }" label="value"   :key="index+'_'+valueIndex"  :prop="'MateList.' + index + '.value.'+valueIndex" >
          <el-input  style="width:400px"      v-model="Mate.value[valueIndex]"  placeholder="填写value"></el-input>
          <el-button  style="margin-left:20px" type="primary" @click="rmValue(index,valueIndex)" >删除value</el-button>
        </el-form-item>
      </div>
      <el-divider></el-divider>
    </div>
  </el-form>
  <el-divider></el-divider>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch, watchEffect} from 'vue'
import {core} from "../../core/core";
export default defineComponent({
  name: 'Mate',
  emits:["input"],
  setup(props,{emit}){
    let state=reactive({
      form:reactive({
        MateList:[]
      }),
      MatchExpressions:reactive({
        MateList:[]
      })
    })


    let formRef=ref(null)
    let MatchExpressionsRef=ref(null)
    function commit(){
      let result={}
      let arr={}
      state.form.MateList.forEach(function (item){
        arr[item.key]=item.value.toString()||""
      })
      result.matchLabels=arr||{}
      let arr2=[]
      state.MatchExpressions.MateList.forEach(function (item){
        item.value.forEach(function (values,i) {
          item.value[i]=item.value[i].toString()
        })
        arr2.push({
          key:item.key.toString(),
          values:item.value,
          operator:item.operator
        })
      })
      result.matchExpressions=arr2||[]
      emit("input",result)
    }
    watch(()=>state.form.MateList,()=>{
      commit()
    },{deep:true,flush:"post"})
    watch(()=>state.MatchExpressions.MateList,()=>{
      commit()
    },{deep:true,flush:"post"})
    function addPathCfg(){
      state.form.MateList.push({key:"",value:"",required:true})
    }
    function addPathCfgMatchExpressions() {
      state.MatchExpressions.MateList.push({key:"",value:[""],operator:"in",required:true})
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
    function rmPathCfgMatchExpressions(index){
      state.MatchExpressions.MateList.splice(index,1)
    }
    function Check(){
      return new Promise(function (resolve){
        formRef.value.validate(async (valid) => {
          if (valid) {
            MatchExpressionsRef.value.validate(async (valid) => {
              if (valid) {
                resolve(true)
              }else{
                resolve(false)
              }
            })
          }else{
            MatchExpressionsRef.value.validate(async (valid) => {
                resolve(false)
            })
          }
        })
      })

    }
    function setData(data){
      //alert(JSON.stringify(data))
      try {
        if(data){
          data2arr(data.matchLabels,state.form.MateList)

          if( data.matchExpressions){
            let arrList=[];
            data.matchExpressions.forEach(function (item){
              arrList.push({
                key:item.key.toString(),
                value:item.values,
                operator:item.operator
              })
            })
            state.MatchExpressions.MateList=arrList
          }else{
            state.MatchExpressions.MateList=[]
          }


        }else{
          state.form.MateList=[]
          state.MatchExpressions.MateList=[];
        }
        commit()
        return true
      }catch (e) {
        console.log(e)
        return false
      }

    }
    let isInArr=core.isInArr
    return {isInArr,addValue,rmValue,rmPathCfg,addPathCfg,...toRefs(state),Check,MatchExpressionsRef,setData,formRef,addPathCfgMatchExpressions,rmPathCfgMatchExpressions}
  }
})
</script>

<style scoped>

</style>
