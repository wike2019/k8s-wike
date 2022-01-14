<template>
           <el-card class="box-card">
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
                     :rules="[requireRules('密文名称必须填写'),inArrayWithMsg(secretCheck,'密文名称不合法')]"
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
</template>


<script lang="ts">
import {computed, defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {IsDirty} from "@/helper/helper";
import {requireRules, NameToArr, inArrayWithMsg} from "../../../helper/rules";
import {secretAllByNs} from "../../../api/token/secret/secret";
export default defineComponent({
  name: 'LabelValue',
  emits:["input"],
  props:["value","namespace","secretsType"],
  setup(props,{emit}){
    let state=reactive({
      secretList:[],
      form:{
        secrets:[]
      },
      label:props.label,
      md5:"",
      secretsType:props.secretsType
    })
    let formRef=ref(null)
    function commit(){
        emit("input",state.form.secrets)
    }
    async function  fetchData(){
      let secret=await  secretAllByNs(props.namespace)
      state.secretList=secret.data.data
      state.form.secrets=[]
    }
    watch(()=>props.namespace,async ()=>{
      await fetchData(props.namespace)
    })
    fetchData()
    watch(()=>state.form,()=>{
      if(IsDirty(state,4)){
        return
      }
      commit()
    },{deep:true,flush:"post"})

    watch(()=>props.value,()=>{

        if(props.value){
          state.form.secrets=props.value
        }else{
          state.md5=""
          state.form.secrets=[]
        }
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
    return {removeItem,addItem,...toRefs(state),Check,formRef,requireRules,inArrayWithMsg,secretCheck,}
  }
})
</script>

<style scoped>

</style>