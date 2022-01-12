<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>密文详情</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="list_item">
       <p>
         账号服务名称: {{secretData.name}}
       </p>
       <p>
         命名空间: {{secretData.namespace}}
       </p>
       <p>
         创建时间: {{secretData.create_time}}
       </p>
       <el-divider></el-divider>
       <h3>标签</h3>
       <div v-if="secretData.labels">
         <p v-for="(value,key) in secretData.labels">
           <strong>{{key}}</strong><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>注释</h3>
       <div v-if="secretData.annotations">
         <p v-for="(value,key) in secretData.annotations">
           <strong>{{key}}</strong><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <el-divider></el-divider>
       <div v-if="secretData.type=='TLS凭据'" >
           <p v-for="(data,item) in secretData.tls_data">
             <span class="key_format"  style="padding-left:15px">{{item}}:</span>
             <span class="value_format">{{data}}</span>
             <el-divider></el-divider>
           </p>

       </div>

       <h2 class="sub_title">Data</h2>
       <el-table
           :data="data"
           border
           empty-text="暂无数据"
           style="width: 100%">

         <el-table-column label="密文key" width="200" align="center">
           <template #default="scope">
              {{ scope.row.key }}
           </template>
         </el-table-column>
         <el-table-column label="密文value(长按显示解密后的值)"   >
           <template #default="scope">
             <div @mousedown="()=>scope.row.show=true" @mouseup="()=>scope.row.show=false">
               <p v-if="!scope.row.show">
                 {{ scope.row.value }}
               </p>
               <p v-else>
                 {{ scope.row.raw }}
               </p>
             </div>
           </template>
         </el-table-column>
       </el-table>
       <el-divider></el-divider>
       <el-button class="mtb20" type="primary" @click="()=>doTo('secret-update',{namespace:namespace,name:name})" >
         编辑
       </el-button>
     </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {ingressDel, ingressListByNs} from "../../api/token/ingress/ingress";
import {secretDel, secretDetail, secretListByNs} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import {doTo} from "../../router";
export default defineComponent({
  name: 'secret-detail',
  components: {MainLayout},
  setup(){

    let state=reactive({
      namespace:"",
      name:"",
      secretData:reactive({data:{},string_data:{}}),
    })
    let data=computed(()=>{
      let arr=[];
      for ( let i in state.secretData.data){
        arr.push({"show":false,"key":i,value:state.secretData.data[i],raw:decodeURIComponent(atob(state.secretData.data[i]))})
      }
      return arr
    })
    let string_data=computed(()=>{
      let arr=[];
      for ( let i in state.secretData.string_data){
        arr.push({"show":false,"key":i,value:state.secretData.data[i],raw:decodeURIComponent(atob(state.secretData.data[i]))})
      }
      return arr
    })
    async function getData(){
      const route = useRoute()
      try {
        state.namespace=route.query.namespace
        state.name=route.query.name
       let tData=await secretDetail(state.namespace,state.name)
        state.secretData=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    getData()
    return {...toRefs(state),data,string_data,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>