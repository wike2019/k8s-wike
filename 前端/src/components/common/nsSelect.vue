<template>
  <div class="name_space_choose">
    <span>命名空间:</span>
    <el-select placeholder="选择命名空间"  filterable v-model="state.namespace">
      <el-option v-for="ns in state.nsList "
                 :label="ns.metadata.name"
                 :value="ns.metadata.name"/>
    </el-select>
    <el-button class="mf40" type="primary" @click="doTo(url)">{{btn}}</el-button>
  </div>
</template>

<script lang="ts" setup>
import { reactive, watch} from 'vue'
import {doTo} from '../../router';
import {getNsALL} from "../../api/token/namespace/ns";
const props = defineProps(['url',"btn"])
let state=reactive({
  nsList: [],
  url:props.url,
  btn:props.btn,
  namespace:"default"
})
const emit = defineEmits(['input'])

watch(()=>state.namespace,()=>{
  emit('input',state.namespace)
})
async function fetchData(){
  try {
    let tData=await getNsALL()
    state.nsList=tData.data
  }catch (e){
    console.log(e)
  }
}
fetchData()
</script>

<style scoped>

</style>
