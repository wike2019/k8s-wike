<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>pod控制台</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
         <div class="name_space_choose name_space_show">
           <span>命名空间:{{name_space}}</span>
           <span>pod名称:{{name}}</span>
           <span>容器名称:</span>
           <el-select placeholder="选择执行控制太的容器" @change="changeContainer" v-model="container">
             <el-option v-for="item in containerList "
                        :label="item.name"
                        :value="item.name"/>
           </el-select>
         </div>
         <el-divider></el-divider>
         <el-form :inline="true"   >
           <el-form-item label="快捷输入"   >
             <el-input v-model="data"  @keyup.enter.native="post" style="width:550px"   placeholder="命令输入"></el-input>
           </el-form-item>
           <el-form-item>
             <el-button size="small" type="primary" @click="post()">确定</el-button>
           </el-form-item>
         </el-form>
         <div class="logs" id="shell">
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
import { FitAddon } from 'xterm-addon-fit';
import {useRoute} from "vue-router";
import { Terminal } from "xterm";
import "xterm/css/xterm.css";
import {buf2str, str2utf8} from "../core/helper";
export default defineComponent({
  name: 'pod',
  components: {MainLayout},
  setup(){
    let rootPath=inject("rootPath")
    let ws
    let term
    let wsInited
    let state=reactive({
      data:"",
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
      wsInited=false
      document.getElementById('shell').innerHTML=""
      ws=new WebSocket(rootPath+"/v1/ws/pods/shell?ns="+state.name_space+"&podname="+state.name+"&cname="+state.container);
      ws.binaryType = "arraybuffer";
      term=new Terminal({
        rendererType: "canvas", //渲染类型
        rows: 35, //行数
        convertEol: true, //启用时，光标将设置为下一行的开头
        disableStdin: false, //是否应禁用输入。
        cursorStyle: "underline", //光标样式
        cursorBlink: true, //光标闪烁
        lineHeight: 1.8,
        theme: {
          foreground: "#e6f1fe", //字体
          background: "#454545", //背景色
          cursor: "help", //设置光标
        }
      });
      let div=document.createElement("div")
      document.getElementById('shell').append(div)
      const fitAddon = new FitAddon();
      term.loadAddon(fitAddon);
      // 创建terminal实例
      term.open(div);

      fitAddon.fit();
      term.prompt = () => {
        term.writeln("WIKE K8S 可视化管理系统.作者qq:200569525");
        term.writeln("正在初始化终端");
      };
      term.prompt();

      term.onData((key)=> {
         if(wsInited){
           ws.send(str2utf8(key))
         }
      });
      ws.onmessage = (e)=>{
        wsInited=true
        if (e.data instanceof ArrayBuffer) {
          term.write(buf2str(e.data));
        } else {
          term.write(buf2str(e.data));
        }
      }
    }
    onUnmounted(()=>{
      if (ws){
        ws.send("exit\n")
        ws.close();
      }
    })
    function post(){
      if(wsInited){
        ws.send(str2utf8(state.data+"\n"))
        state.data=""
      }
    }

    return {...toRefs(state),changeContainer,post}
  }
})
</script>
