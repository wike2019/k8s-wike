<template>
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div class="inline-title">
      <span>设置{{title}}</span>         <el-button type="success" size="small" style="margin-right:50px;float: right" @click="addPathCfg" >添加{{title}}</el-button>
    </div>
    <div v-for="(Mate,index) in form.MateList" style="margin:20px 0">
      <el-form-item label="值"   :key="'key'+index"  :prop="'MateList.' + index + '.key'"
                    :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
        <el-select v-model="Mate.key">
          <el-option v-for="item in list"
                     :label="item"
                     :value="item"/>
        </el-select>
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
  props:['title','list'],
  setup(props,{emit}){
    let state=reactive({
      form:reactive({
        MateList:[]
      }),
      title:props.title,
      list:props.list
    })


    let formRef=ref(null)
    watch(()=>state.form.MateList,()=>{
      let arr=[];
      for (let i=0;i<state.form.MateList.length;i++){
        arr.push(state.form.MateList[i].key)
      }
      emit("input",arr)
    },{deep:true,flush:"post"})
    function addPathCfg(){
      console.error( state.form.MateList)
      state.form.MateList.push({key:""})
      console.error( state.form.MateList)
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
      let mapData=[]
      for (let i=0;i<data.length;i++){
        mapData.push({key:data[i]})
      }
      state.form.MateList=mapData
    }
    return {rmPathCfg,addPathCfg,...toRefs(state),Check,setData,formRef}
  }
})
</script>

<style scoped>

</style>
