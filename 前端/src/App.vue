<template>
    <div id="panel">
    </div>
    <div v-if="!init" class="body" @click="init=true">
      <div   class="outWrap">
        <div class="font"><span>WIKE</span> </div>
        <div class="wrap">
          <div class="wrapOver">
            <div class="circle1"></div>
          </div>
          <div class="wrapOver">
            <div class="circle2"></div>
          </div>
          <div class="bottomSliver"></div>
        </div>
        <div class="back">k8S</div>
      </div>
      <div class="animate__rotateIn animate__animated enter" >点击体验</div>
    </div>

      <div v-if="init">
          <header class="header_all">
            WIKE K8S 可视化管理系统-基于vue3
            <a href="https://github.com/wike2019/k8s-wike">项目github地址</a>
            <span>项目作者 wike 联系qq 200569525</span>
            <a v-if="isLogin" @click="quit" class="quit">退出登陆</a>
          </header>
          <router-view></router-view>
          <footer></footer>
      </div>
</template>

<script lang="ts">
import {defineComponent, onUnmounted,onMounted, provide, ref} from 'vue'
import {doTo} from "./router";
import {core} from "./core/core";
import {WSAPI} from "./config/api";
export default defineComponent({
  name: 'App',
  setup(){
    let ws= new WebSocket(WSAPI+"/v1/ws");
    let isLogin=ref(false);
    let init=ref(false)
    provide("ws",ws)
    provide("rootPath",WSAPI)
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
    })
    onMounted(()=>{

    })

    return {isLogin,quit,init}
  }
})
</script>

<style scoped   lang="less">

.body {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  overflow: hidden;
  background: black;
}
.outWrap{
  position: fixed;
  top: 30%;
  left: 12%;
  display: flex;
  animation: moveFiexd 4s forwards
}
.font{
  height: 177px;
  color: white;
  font-size: 104px;
  line-height: 177px;
  vertical-align: middle;
  opacity: 0;
  margin-right:20px;
  animation: transparency 4s forwards
}
.font span{padding-top:20px}
.back{
  height: 150px;
  color: white;
  font-size: 131px;
  line-height: 177px;
  opacity: 0;
  margin-left:20px;
  animation: transparency 4s forwards
}
.wrap{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.wrapOver {
  position: relative;
  width: 100px;
  height: 50px;
  overflow: hidden
}

.circle1 {
  position: absolute;
  box-sizing: border-box;
  width: 100px;
  height: 100px;
  border: 15px solid white;
  border-radius: 50%;
  transform: translateY(100%);
  filter:blur(2px);
  animation: myMoveTop 2s forwards
}
.circle2 {
  position: absolute;
  box-sizing: border-box;
  width: 50px;
  height: 50px;
  border: 15px solid white;
  border-radius: 50%;
  transform: translateY(-150%);
  filter:blur(2px);
  animation: myMoveBottom 2s forwards
}
.bottomSliver{
  width: 0px;
  height: 5px;
  background:#00a1d6 ;
  margin-top: 5px;
  animation: myMoveSliver 4s forwards
}
@keyframes  transparency {
  50% { opacity: 0;}
  100% { opacity: 1;}
}
@keyframes myMoveSliver {
  50% { width: 0px;}
  100% { width: 50px;}
}
@keyframes myMoveBottom {
  90% {filter:blur(2px); }
  100% { transform: translateY(-50%);filter:blur(0px);}
}

@keyframes myMoveTop {
  90% {filter:blur(2px); }
  100% { transform: none;filter:blur(0px); }
}
@keyframes moveFiexd {
  50% {left:15%; }
  100%  {left:32%;  }
}

.enter{color:#fff;font-size:40px;text-align: center;position: absolute;bottom:150px;width:97%;animation-delay:4s;cursor: pointer!important;}

</style>

