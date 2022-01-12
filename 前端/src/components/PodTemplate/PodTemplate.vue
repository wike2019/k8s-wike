<template>
  <mateData @input="getData($event,'metadata',form)" ref="mateDataRef" :nameRequired="false"></mateData>
  <el-divider></el-divider>
  <h3>基础属性</h3>
  <el-divider></el-divider>
  <volumes :namespace="namespace" @input="getData($event,'spec.volumes',form)" ref="volumesRef"></volumes>
  <containers :namespace="namespace" tip="init" @input="getData($event,'spec.initContainers',form)" ref="initContainersRef"></containers>
</template>

<script lang="ts">
import {ref, defineComponent, reactive, toRefs, watch} from 'vue'
import mateData from "../Metadata/matedata.vue";
import volumes from "../Volumes/volumes.vue";
import containers from "../Container/containers.vue";
import {getData} from "../../helper/helper.ts"
export default defineComponent({
  name: 'PodTemplate',
  components:{
    mateData,
    volumes,
    containers
  },
  props:["namespace"],
  emits:["input"],
  setup: (props,{emit}) => {
    let mateDataRef=ref(null)
    let volumesRef=ref(null)
    let initContainersRef=ref(null)
    let state=reactive({
      namespace:props.namespace,
      form:{
        metadata:{

        },
        spec:{
          volumes:[],
          initContainers:[]
        }
      },
    })


    function commit() {
       //整合好数据 提交给父组件
       let result={}
       result.metadata={}
       result.metadata.annotations=state.form.metadata.annotations||{}
       result.metadata.labels=state.form.metadata.labels||{}
       result.spec={}
       result.spec.volumes=state.form.spec.volumes||[]
       result.spec.initContainers=state.form.spec.initContainers||[]
       emit("input",result)
    }
    function setData(data){
      if(data&&data.metadata){
        mateDataRef.value.setData(data.metadata)
      }
      if(data&&data.spec&&data.spec.volumes){
        volumesRef.value.setData(data.spec.volumes)
      }
      if(data&&data.spec&&data.spec.initContainers){
        initContainersRef.value.setData(data.spec.initContainers)
      }
      commit()
    }
    watch(()=>props.namespace,()=>{
      state.namespace=props.namespace
    })

    watch(()=>state.form,()=>{
      commit()
    },{deep:true,flush:"post"})


    function Check(){
      return new Promise(async  function (resolve){
          let mateData= await  mateDataRef.value.Check()
          let volumesData= await  volumesRef.value.Check()
          let initContainersData= await  volumesRef.value.Check()
          resolve(mateData&&volumesData&&initContainersData)
      })

    }
    return { mateDataRef ,volumesRef,initContainersRef,...toRefs(state),setData,getData,Check}
  }
})
</script>

<style scoped>

</style>
