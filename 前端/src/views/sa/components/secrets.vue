<template>
           <el-card class="box-card" v-if="!render">
             <template #header>
               <div class="card-header mtb20" >
                 <span class="title">添加密文-{{secretsType}}</span>
                 <el-button type="primary" size="small" class="leftBtn"  @click="addItem" >添加密文</el-button>
               </div>
             </template>
             <el-form   :model="form" ref="formRef"  >
             <div v-for="(item,index) in form.secrets" :key="secretsType+'secrets'+index" class="mtb20">
                 <el-form-item
                     label="密文名称"
                     :key="secretsType+'secrets'+index+'name'"
                     :prop="'secrets.'+index+'.name'"
                     :rules="[requireRules(tipTitle+'密文名称必须填写'),inArrayWithMsg(secretCheck,tipTitle+'密文名称不合法')]"
                 >
                   <el-select  size="small" v-model="item.name"  filterable>
                     <el-option :key="key" v-for="(item,key) in secretList" :label="item.name" :value="item.name"/>
                   </el-select>
                 </el-form-item>
                 <el-form-item  >
                   <el-button size="small" type="danger" @click="removeItem(index)" >删除该项items</el-button>
                 </el-form-item>
             </div>
             </el-form>
           </el-card>
          <div  v-if="render&&form.secrets.length">
            <el-descriptions
                class="smallTable"

                :column="1"
                :title="secretsType"
            >
              <el-descriptions-item v-for="item in form.secrets" :label="tipTitle+':密文名称'">
                {{item.name}}
                <a  @click="()=>doTo('secret-detail',{namespace:namespace,name:item.name})">查看密文详情</a>
              </el-descriptions-item>
            </el-descriptions>
            <el-divider></el-divider>
          </div>

</template>


<script lang="ts">
import {computed, defineComponent, inject, reactive, ref, toRefs, watch} from 'vue'
import {IsDirty} from "@/helper/helper";
import {requireRules, NameToArr, inArrayWithMsg} from "../../../helper/rules";
import {secretAllByNs} from "../../../api/token/secret/secret";
import {doTo} from "../../../router";
export default defineComponent({
  name: 'LabelValue',
  emits:["input"],
  props:["value","namespace","secretsType",'tipTitle'],
  setup(props,{emit}){
    const render = inject("render")
    let state=reactive({
      secretList:[],
      tipTitle:props.tipTitle,
      form:{
        secrets:[]
      },
      defaultData:{
        secrets:[]
      },
      md5:"",
      namespace:"default",
      secretsType:props.secretsType,
      render:render
    })
    let formRef=ref(null)
    function commit(){
        emit("input",state.form.secrets)
    }
    async function  fetchData(){
      let secret=await  secretAllByNs(props.namespace)
      state.secretList=secret.data.data.filter(item=>{
         if(state.tipTitle=='secrets'){
           if (item.type=="服务账号令牌"){
             return true
           }else{
             return  false
           }
         }
        if(state.tipTitle=='imagePullSecrets'){
          if (item.type=="docker配置"||item.type=="docker配置(JSON)"){
            return true
          }else{
            return  false
          }
        }
      })
    }
    watch(()=>props.namespace,async ()=>{
      state.namespace=props.namespace
      await fetchData()
    })
    fetchData()
    watch(()=>state.form,()=>{
      if(IsDirty(state,{secrets:props.value})){
        return
      }
      commit()
    },{deep:true,flush:"post"})

    watch(()=>props.value,()=>{
       state.form=Object.assign({},state.defaultData,{secrets:props.value})
    },{deep:true,flush:"post"})

    function Check(){
      formRef.value.validate()
    }
    function removeItem(index) {
      state.form.secrets.splice(index,1)
    }
    function addItem() {
      if( state.form.secrets){
        state.form.secrets.push({name:""})
      }else{
        state.form.secrets=[{name:""}]
      }

    }
    let secretCheck=computed(()=> {
      return NameToArr(state.secretList)
    })
    return {removeItem,addItem,...toRefs(state),Check,formRef,requireRules,inArrayWithMsg,secretCheck,doTo}
  }
})
</script>

<style scoped>

</style>