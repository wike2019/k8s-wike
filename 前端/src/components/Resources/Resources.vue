<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header" style="padding-bottom:10px">
        权限配置 <el-input style="width:600px;margin-left:30px" size="small" v-model="filterText" placeholder="关键字过滤" />
        <el-button size="small" style="margin-left:30px" @click="flush" type="danger">强制获取API资源列表(不走缓存)</el-button>
      </div>
      <el-divider></el-divider>
      <el-tree  :default-checked-keys="defaultKey" ref="treeRef" node-key="id" :data="data"  :props="{label:'label',children:'children'}" show-checkbox  :filter-node-method="filterNode" >
      </el-tree>
    </template>
  </el-card>
</template>

<script lang="ts">
import {ref, defineComponent, reactive, computed, toRefs, watch} from 'vue'
import {doTo} from '../../router';
import {getNsList} from "../../api/token/namespace/ns";
import {getResources} from "../../api/token/common/common";
export default defineComponent({
  props:['defaultKey'],
  name: 'Resources',
  setup: (props) => {
    let state=reactive({
      ResourcesList:reactive([]),
      filterText:"",
    })
    let defaultKey=ref([])
    async function flush(){
      try {
        let tData=await getResources()
        localStorage.setItem("ResourcesList",JSON.stringify(tData.data.data))
        state.ResourcesList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    async function getData(){
      if (localStorage.getItem("ResourcesList")){
        state.ResourcesList=JSON.parse(localStorage.getItem("ResourcesList"))
        return
      }
      await flush()
    }
    watch(()=>state.filterText,(val)=>{
      treeRef.value.filter(val)

    },{deep:true,flush:"sync"})
    let treeRef=ref(null)
    getData()
    let data=computed(()=>{
      let result=[];
      state.ResourcesList.forEach(function (GroupData) {
        let child=[]
        GroupData.Resources.forEach(function (val){
           let lastData=[]
          val.Verbs.forEach(function (v) {
             lastData.push({label:v,id:GroupData.Group  +"@"+val.Name+"@"+v})
          })
          child.push({label:val.Name,children:lastData})
        })
        result.push({label:GroupData.Group,children:child})
      })
      return result
    })

    function getCheckedKeys() {
      let result=[]
      let data=treeRef.value.getCheckedKeys()
      data.forEach(function (item){
        if(item){
          let arr=item.split("@")
          if (arr[0]=="core"){
            arr[0]=""
          }
          result.push({apiGroups:[arr[0]],resources:[arr[1]],verbs:[arr[2]]})
        }

      })
      return result
    }
    function filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    }
    function setData(arr){
      defaultKey.value=arr
    }
    return { doTo,...toRefs(state) ,data,treeRef,getCheckedKeys,filterNode,flush,defaultKey,setData}
  }
})
</script>

<style scoped>

</style>
