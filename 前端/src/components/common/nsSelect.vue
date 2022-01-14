<template>
  <div class="name_space_choose">
    <span>命名空间:</span>
    <el-select placeholder="选择命名空间" @change="changeNs" filterable v-model="namespace">
      <el-option v-for="ns in nsList "
                 :label="ns.name"
                 :value="ns.name"/>
    </el-select>
    <el-button class="mf40" type="primary" @click="doTo(url)">{{btn}}</el-button>
  </div>
</template>

<script lang="ts">
import {defineComponent, reactive, toRefs, watch} from 'vue'
import {doTo} from '../../router';
import {getNsList} from "../../api/token/namespace/ns";
export default defineComponent({
  name: 'ns-select',
  props:["url","btn"],
  emits:["input"],
  setup: (props,{emit}) => {
    let state=reactive({
      nsList: [],
      url:props.url,
      btn:props.btn,
      namespace:"default"
    })
    watch(()=>state.namespace,()=>{
      emit('input',state.namespace)
    })
    async function fetchData(){
      try {
        let tData=await getNsList()
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    fetchData()
    return { doTo,...toRefs(state)}
  }
})
</script>

<style scoped>

</style>
