<template>
  <header class="header_all">
    WIKE K8S 可视化管理系统-基于vue3
    <em style="margin-left:20px;font-size:14px;">本项目部分代码来自<a style="color:#67C23A;margin-left:5px;font-weight: bold" href="https://www.jtthink.com/">https://www.jtthink.com(学习k8s请进入)</a> </em>
    <span>项目作者 wike 联系qq 200569525</span>
    <a v-if="isLogin" @click="quit" class="quit">退出登陆</a>
  </header>
  <router-view></router-view>
  <footer></footer>
</template>

<script lang="ts">
import {defineComponent, onUnmounted, provide, reactive,toRefs} from 'vue'
import {go} from "./router";
import {core} from "./core/core";
import {API, WSAPI} from "./config/api";
export default defineComponent({
  name: 'App',
  setup(){
    let rootPath=WSAPI
    let ws= new WebSocket(rootPath+"/v1/ws");
    let state=reactive({
      isLogin:false
    })
    provide("ws",ws) //提供ws 客户端给子组件
    provide("rootPath",rootPath)

    core.bus.on('login',function (){
      //监听登录变化
      state.isLogin=true;
    })
    if(localStorage.getItem("token")){
      //判断是否登录
      state.isLogin=true;
    }
    function quit(){
      //退出登录
      localStorage.removeItem("token")
      state.isLogin=false;
      go("login")
    }
    onUnmounted(()=>{
      ws.close()
    })
    return {...toRefs(state),quit}
  }
})
</script>

<style>

</style>
