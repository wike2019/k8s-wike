<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header" style="padding-bottom:10px">
        事件数据
      </div>
    </template>
    <div v-if="eventList.length">
      {{eventList}}
    </div>
    <div v-else class="no-content">
    </div>
  </el-card>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {arrToMap, MapToArr} from "../../helper/helper";
import {getEventList} from "../../api/token/event/event";
export default defineComponent({
  name: 'event-list',
  props:['namespace','name','kind','uid'],
  setup(props,{emit}){
    let state=reactive({
      eventList:[]
    })
    //ns,name,kind,uid
    async function fetchData (){
      let tData=await getEventList(props.namespace,props.name,props.kind,props.uid)
      state.eventList=tData.data.data
    }
    watch(()=>props.uid,()=>{
      if(props.namespace&&props.name&&props.kind&&props.uid){
        fetchData()
      }
    })

    return {...toRefs(state),}
  }
})
</script>

<style scoped>

</style>
