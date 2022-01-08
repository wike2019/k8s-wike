<template>

    <div v-if="type=='ingress'">
      <div>
        <span class="fast_action">快捷操作</span>
        <el-switch style="margin-right:20px"  v-model="enable_cors"  inline-prompt active-text="是否跨域" />
        <el-switch  v-model="enable_rewriter"  inline-prompt active-text="是否重写" />
      </div>
      <el-divider></el-divider>
    </div>
    <MyLabel @input="input" ref="MateRef"></MyLabel>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch,nextTick,onMounted} from 'vue'
import MyLabel from "../Label/label.vue"
export default defineComponent({
  components:{MyLabel},
  name: 'Annotations',
  emits:["input"],
  props:['type'],
  setup(props,{emit}){
    let MateRef=ref(null)
    let state=reactive({
      enable_cors:false,
      enable_rewriter:false,
      type:props.type||""
    })
    let data=[
      {
        key:"nginx.ingress.kubernetes.io/enable-cors",
        value:"true",
        name:"启用跨域",
        required:true
      },
      {
        key:"nginx.ingress.kubernetes.io/cors-allow-origin",
        value:"*",
        name:"允许来源",
        required:true
      },
      {
        key:"nginx.ingress.kubernetes.io/cors-allow-methods",
        value:'GET, PUT, POST, DELETE, PATCH, OPTIONS',
        name:"请求方法"
      },
      {
        key:"nginx.ingress.kubernetes.io/cors-allow-headers",
        value:'Content-Type,AccessToken,X-CSRF-Token, Authorization, Token',
        name:"允许头"
      },
      {
        key:"nginx.ingress.kubernetes.io/cors-expose-headers",
        value:' ',
        name:"暴露头(JS可访问)"
      },
      {
        key:"nginx.ingress.kubernetes.io/cors-allow-credentials",
        value:false,
        name:"允许凭据",
        required:true
      },
    ]
    let cache=[]
    function input(data){
      cache=data
      emit("input",data)
    }
    function setData(data) {
      MateRef.value.setData(data)
    }
    async function Check(){
      return  MateRef.value.Check()
    }
    function addPathCfg(){
      MateRef.value.addPathCfg()
    }
    watch(()=>state.enable_cors,function (){
       let oldCache=cache
       clear_cors(oldCache)
       if(state.enable_cors){
       add_cors(oldCache)
       }
        MateRef.value.setData(oldCache)
        emit("input",oldCache)
    },{deep:true,flush:"post"})
    watch(()=>state.enable_rewriter,function (){
      let oldCache=cache
      clear_rewriter(oldCache)
      if(state.enable_rewriter){
        add_rewriter(oldCache)
      }
        MateRef.value.setData(oldCache)
        emit("input",oldCache)
    },{deep:true,flush:"post"})
    function add_rewriter(cache){

       let data= {
          key:"nginx.ingress.kubernetes.io/rewrite-target",
          value:"/$1",
          name:"重写规则",
          required:true
        }
      cache.push(data)
    }
    function clear_rewriter(cache){
      for (let i=0;i<cache.length;i++){
        if(cache[i].key=="nginx.ingress.kubernetes.io/rewrite-target"){
          console.log(cache[i].key)
          cache.splice(i,1)
        }
      }
    }
    function clear_cors(cache){
      for (let i=0;i<cache.length;i++){
        for (let j=0;j<data.length;j++){
          if(cache[i].key==data[j].key){
            console.log(cache[i].key)
            cache.splice(i,1)
          }

        }
      }
    }
    function add_cors(cache){
      for (let i=0;i<data.length;i++){
        cache.push(data[i])
      }
    }
    onMounted(()=>{
      if (props.type=="ingress"){
        MateRef.value.setData([
          {key:"kubernetes.io/ingress.class",value:"nginx",name:"ingress类型"}
        ])
      }
    })

    return {MateRef,Check,input,addPathCfg,...toRefs(state),setData}
  }
})
</script>

<style scoped>

</style>
