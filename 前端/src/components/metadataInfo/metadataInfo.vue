<template>
  <el-card class="box-card" v-if="!render">
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
      <el-form-item v-if="nameRequired" label="名称"  :rules="requireRules(tipTitle+'名称必须填写')" prop="name"  >
        <el-input v-model="form.name"  :disabled="nameDisable"  placeholder="配置名称"></el-input>
      </el-form-item>
      <el-form-item v-if="nameRequired"  label="命名空间" prop="namespace" :rules="inArrayWithMsg(nsCheck,tipTitle+'命名空间不合法')">
        <el-select  v-model="form.namespace" filterable  :disabled="nameDisable">
          <el-option v-for="ns in nsList "
                     :label="ns.name"
                     :value="ns.name"/>
        </el-select>
      </el-form-item>
    </el-form>
      <el-divider></el-divider>
      <labelValue @input="getData($event,'labels',form,true)" :value="form.labels" ref="MyLabelRef"></labelValue>
      <annotations @input="getData($event,'annotations',form,true)"  :value="form.annotations" ref="AnnotationsRef" :type="type"></annotations>
    </div>

  </el-card>
  <el-card class="box-card" v-if="render">
    <el-descriptions
        :column="1"
        title="资源基础信息"
    >
      <el-descriptions-item label="资源类型">{{tipTitle}}</el-descriptions-item>
      <el-descriptions-item label="资源名称">{{form.name}}</el-descriptions-item>
      <el-descriptions-item label="资源所属命名空间">{{form.namespace}}</el-descriptions-item>
      <el-descriptions-item label="资源唯一标识" >{{form.uid}}</el-descriptions-item>
      <el-descriptions-item label="资源创建时间" >{{form.creationTimestamp}}</el-descriptions-item>、
      <el-descriptions-item label="快捷操作" >
        <el-button
            type="primary"
            size="small"
            @click="()=>doTo(edit,rowToQuery(form))"  >
          编辑
        </el-button>
        <el-button type="danger" size="small" @click="()=>rm(form.namespace,form.name )" >删除</el-button>
      </el-descriptions-item>
    </el-descriptions>
    <labelValue :value="form.labels" ref="MyLabelRef"></labelValue>
    <annotations  :value="form.annotations" ref="AnnotationsRef" ></annotations>
  </el-card>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs, watch, computed, inject} from 'vue'
import {getNsList} from "@/api/token/namespace/ns";
import labelValue from "@/components/labelValue/labelValue.vue"
import annotations from "@/components/annotations/annotations.vue";
import { getData, IsDirty} from "@/helper/helper";
import {inArrayWithMsg, NameToArr} from "@/helper/rules";
import {CheckData, rmTip} from "../../helper/helper";
import {requireRules} from "../../helper/rules";
export default defineComponent({
  name: 'metadataInfo',
  emits:["input"],
  props:["nameRequired","nameDisable","type","value","tipTitle"],
  components:{
    labelValue,
    annotations
  },
  setup: (props,{emit}) => {
    const render = inject("render")
    const delFn = inject("delFn")
    const edit = inject("edit")
    let state=reactive({
      nameRequired:!!props.nameRequired,
      nameDisable:!!props.nameDisable,
      tipTitle:props.tipTitle,
      show:true,
      type:props.type,
      nsList:[],
      md5:"",
      defaultData:{
        name:"",
        namespace:"default",
      },
      form:{
        name:"",
        namespace:"default",
      },
      render:render,
      edit:edit
    })
    let MyLabelRef=ref(null)
    let AnnotationsRef=ref(null)
    let checkList=[AnnotationsRef,MyLabelRef]
    let formRef=ref(null)

    watch(()=>props.value,()=>{
      state.form=Object.assign({},state.defaultData,props.value)
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
      await CheckData(checkList,2)
      formRef.value.validate()
    }
    watch(()=>state.form,()=>{
      if(IsDirty(state,props.value,false,1)){
        return
      }
      commit()
    },{deep:true,flush:"post"})
    let nsCheck=computed(()=> {
      return NameToArr(state.nsList)
    })

    function rm(ns,name){
      rmTip(state.tipTitle)
          .then(async () => {
            try {
              await delFn(ns,name)

            }catch (e) {
              console.log(e)
            }
          })
    }
    return {formRef,...toRefs(state),MyLabelRef,Check,AnnotationsRef,getData,inArrayWithMsg,nsCheck,requireRules,rm}
  }
})
</script>

<style scoped>

</style>
