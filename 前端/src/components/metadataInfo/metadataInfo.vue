<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        <span class="title">基本配置(metadata数据)</span>
        <el-button size="small" class="mf20" @click="show=!show">
          {{!show?"展开":"收缩"}}
        </el-button>
      </div>
    </template>
    <div v-show="show">
    <el-form  :inline="true"  :model="form" ref="formRef">
      <el-form-item v-if="nameRequired" label="名称"  :rules="requireRules('名称必须填写')" prop="name"  >
        <el-input v-model="form.name"  :disabled="nameDisable"  placeholder="配置名称"></el-input>
      </el-form-item>
      <el-form-item v-if="nameRequired"  label="命名空间" prop="namespace" :rules="inArrayWithMsg(nsCheck,'命名空间不合法')">
        <el-select  v-model="form.namespace" filterable  :disabled="nameDisable">
          <el-option v-for="ns in nsList "
                     :label="ns.name"
                     :value="ns.name"/>
        </el-select>
      </el-form-item>
    </el-form>
      <el-divider></el-divider>
      <labelValue @input="getData($event,'labels',form)" :value="form.labels" ref="MyLabelRef"></labelValue>
      <annotations @input="getData($event,'annotations',form)"  :value="form.annotations" ref="AnnotationsRef" :type="type"></annotations>
    </div>

  </el-card>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs, watch, computed} from 'vue'
import {getNsList} from "@/api/token/namespace/ns";
import labelValue from "@/components/labelValue/labelValue.vue"
import annotations from "@/components/annotations/annotations.vue";
import { getData, IsDirty} from "@/helper/helper";
import {inArrayWithMsg, NameToArr} from "@/helper/rules";
import {CheckData} from "../../helper/helper";
import {requireRules} from "../../helper/rules";
export default defineComponent({
  name: 'metadataInfo',
  emits:["input"],
  props:["nameRequired","nameDisable","type","value"],
  components:{
    labelValue,
    annotations
  },
  setup: (props,{emit}) => {
    let state=reactive({
      nameRequired:!!props.nameRequired,
      nameDisable:!!props.nameDisable,
      show:true,
      type:props.type,
      nsList:[],
      md5:"",
      defaultData:{
        name:"",
        namespace:"default",
        labels:{},
        annotations:{},
      },
      form:{
        name:"",
        namespace:"default",
        labels:{},
        annotations:{},
      },
    })
    let MyLabelRef=ref(null)
    let AnnotationsRef=ref(null)
    let checkList=[AnnotationsRef,MyLabelRef]
    let formRef=ref(null)

    watch(()=>props.value,()=>{
      if (props.value){
        state.form=props.value
      }else{
        state.md5=""
        state.form=state.defaultData
      }
    })
    function commit(){
      emit('input',state.form)
    }
    async function fetchData(){
      try {
        let tData=await getNsList()
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    fetchData()

    async function Check(){
      formRef.value.validate()
    }
    watch(()=>state.form,()=>{
      if(IsDirty(state,2)){
        return
      }
      commit()
    },{deep:true,flush:"post"})
    let nsCheck=computed(()=> {
      return NameToArr(state.nsList)
    })

    return {formRef,...toRefs(state),MyLabelRef,Check,AnnotationsRef,getData,inArrayWithMsg,nsCheck,requireRules}
  }
})
</script>

<style scoped>

</style>
