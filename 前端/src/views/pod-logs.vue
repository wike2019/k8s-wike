<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>pod日志</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
         <div class="name_space_choose name_space_show">
           <span>命名空间:{{name_space}}</span>
           <span>pod名称:{{name}}</span>
           <span>容器名称:</span>
           <el-select placeholder="选择查看日志的容器" @change="changeContainer" v-model="container">
             <el-option v-for="item in containerList "
                        :label="item.name"
                        :value="item.name"/>
           </el-select>
         </div>
         <el-divider></el-divider>
         <div class="logs" id="msg_end">
           <div>{{logs}}</div>
         </div>

     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive, toRefs, nextTick} from 'vue'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {getPodsByNs, getPodsContainers} from "../api/token/pod";
import {secretDetail} from "../api/token/secret/secret";
import {useRoute} from "vue-router";
export default defineComponent({
  name: 'pod',
  components: {MainLayout},
  setup(){
    let rootPath=inject("rootPath")
    let ws
    let state=reactive({
      containerList:reactive([]),
      logs:"",
      container:"",
      name_space:null,
      name:null
    })
    const route = useRoute()
    state.name_space=route.query.name_space
    state.name=route.query.name
    async function getData(){
      try {

        let tData=await getPodsContainers(state.name_space,state.name)
        state.containerList=tData.data.data
        state.container=state.containerList[0].name
        changeContainer()
      }catch (e){
        console.log(e)
      }
    }
    getData()
    function changeContainer(){

      if (ws){
        ws.close();
      }
      ws=new WebSocket(rootPath+"/v1/ws/pods/logs?ns="+state.name_space+"&podname="+state.name+"&cname="+state.container);
      ws.onmessage = (e)=>{
        console.log(e.data)
        state.logs+=e.data

        nextTick(function () {
          let div=document.getElementById('msg_end')
          console.log(div.scrollHeight)
          console.log(   div.scrollTop)
          div.scrollTop = div.scrollHeight
          console.log(   div.scrollTop)
        })

      }
    }
    onUnmounted(()=>{
      if (ws){
        ws.close();
      }
    })

    return {...toRefs(state),changeContainer}
  }
})
</script>
