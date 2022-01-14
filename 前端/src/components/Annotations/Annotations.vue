<template>
    <div class="inline-title">
      <span>注解数据设置</span>         <el-button type="primary" size="small" class="leftBtn" @click="addItem" >添加注释</el-button>
    </div>
    <div v-if="type=='ingress'">
      <div class="mtb20">
        <span class="fast_action">快捷操作</span>
        <el-switch style="margin-right:20px"  v-model="enable_cors"  active-text="是否跨域" />
        <el-switch  v-model="enable_rewriter" active-text="是否重写" />
      </div>
      <el-divider></el-divider>
    </div>
    <labelValue @input="getData($event,'annotations',form)" :value="form.annotations" ref="MyLabelRef" :label="true"></labelValue>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import labelValue from "@/components/labelValue/labelValue.vue"
import {arrToMap, getData, MapToArr,IsDirty} from "@/helper/helper";
export default defineComponent({
  components:{labelValue},
  name: 'annotations',
  emits:["input"],
  props:['type',"value"],
  setup(props,{emit}){
    let MyLabelRef=ref(null)
    let state=reactive({
      enable_cors:false,
      enable_rewriter:false,
      type:props.type||"",
      md5:"",
      form:{
        annotations:[]
      }
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

    async function Check(){
      return  MyLabelRef.value.Check()
    }
    function addItem(){
      MyLabelRef.value.addItem()
    }
    watch(()=>state.enable_cors,function (){
       clear_cors()
       if(state.enable_cors){
       add_cors()
       }
    },{deep:true,flush:"post"})
    watch(()=>state.enable_rewriter,function (){
      clear_rewriter()
      if(state.enable_rewriter){
        add_rewriter()
      }
    },{deep:true,flush:"post"})
    function AddTitle(list){
      for (let i=0;i<list.length;i++){
        for (let j=0;j<data.length;j++){
           if(list[i].key==data[j].key){
             list[i].name=data[j].name
           }
        }
        if(list[i].key=="nginx.ingress.kubernetes.io/rewrite-target"){
          list[i].name="重写规则"
        }
      }
      return list
    }
    function add_rewriter(){

       let data= {
          key:"nginx.ingress.kubernetes.io/rewrite-target",
          value:"/$1",
          name:"重写规则",
          required:true
        }
       state.form.annotations.push(data)
    }
    function clear_rewriter(){
      for (let i=0;i< state.form.annotations.length;i++){
        if( state.form.annotations[i].key=="nginx.ingress.kubernetes.io/rewrite-target"){
          state.form.annotations.splice(i,1)
        }
      }
    }
    function clear_cors(){
      for (let i=0;i< state.form.annotations.length;i++){
        for (let j=0;j<data.length;j++){
          if( state.form.annotations[i].key==data[j].key){
            state.form.annotations.splice(i,1)
          }
        }
      }
    }
    function add_cors(){
      for (let i=0;i<data.length;i++){
        state.form.annotations.push(data[i])
      }
    }
    function commit(){
      emit("input",arrToMap(state.form.annotations))
    }
    watch(()=>props.value,()=>{
      if(props.value){
        state.form.annotations=AddTitle(MapToArr(props.value))
      }else{
        state.md5=""
        state.form.annotations=[]
      }

    })
    watch(()=>state.form,()=>{
      if(IsDirty(state,4)){
        return
      }
     commit()
    },{deep:true,flush:"post"})
    return {MyLabelRef,Check,getData,addItem,...toRefs(state)}
  }
})
</script>

<style scoped>

</style>
