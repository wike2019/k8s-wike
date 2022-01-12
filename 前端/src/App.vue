<template>
  <header class="header_all">
    WIKE K8S 可视化管理系统-基于vue3
    <a href="https://github.com/wike2019/k8s-wike">项目github地址</a>
    <span>项目作者 wike 联系qq 200569525</span>

    <a v-if="isLogin" @click="quit" class="quit">退出登陆</a>
  </header>
  <router-view></router-view>
  <footer></footer>
</template>

<script lang="ts">
import {defineComponent, onUnmounted, provide, ref} from 'vue'
import {doTo} from "./router";
import {core} from "./core/core";
import {API, WSAPI} from "./config/api";
export default defineComponent({
  name: 'App',
  setup(){

    //let rootPath="ws://127.0.0.1:7772"
    let rootPath=WSAPI
    let ws= new WebSocket(rootPath+"/v1/ws");
    let isLogin=ref(false);
    provide("ws",ws)
    provide("rootPath",rootPath)
    core.bus.on('login',function (){
      isLogin.value=true;
    })
    if(localStorage.getItem("token")){
      isLogin.value=true;
    }
    function quit(){
      localStorage.removeItem("token")
      isLogin.value=false;
      doTo("login")
    }
    onUnmounted(()=>{
      ws.close()
      //ws=null
    })
    return {isLogin,quit}
  }
})
</script>

<style>

</style>
