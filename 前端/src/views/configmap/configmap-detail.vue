<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>密文详情</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
       <div class="list_item">
         <p>
           配置名称: {{configmapData.name}}
         </p>
         <p>
           命名空间: {{configmapData.namespace}}
         </p>
         <p>
           创建时间: {{configmapData.create_time}}
         </p>

         <el-divider></el-divider>
       <h3>标签</h3>
       <div v-if="configmapData.labels">
         <p v-for="(value,key) in configmapData.labels">
           <span>{{key}}</span><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <h3>注释</h3>
       <div v-if="configmapData.annotations">
         <p v-for="(value,key)  in configmapData.annotations">
           <span>{{key}}</span><em>:</em><span>{{value}}</span>
         </p>
       </div>
       <div v-else>
         无数据
       </div>
       <el-divider></el-divider>
       <el-table
           :data="data"
           border
           empty-text="暂无数据"
           style="width: 100%">

         <el-table-column label="配置key" width="200" align="center">
           <template #default="scope">
              {{ scope.row.key }}
           </template>
         </el-table-column>
         <el-table-column label="配置value"   align="center">
           <template #default="scope">
             <div>
               <p v-if="!scope.row.show">
                 {{ scope.row.value }}
               </p>
             </div>
           </template>
         </el-table-column>
       </el-table>
         <el-button class="mtb20" type="primary" @click="()=>doTo('configmap-update',{namespace:configmapData.namespace,name:configmapData.name})" >
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
import {configmapDetail} from "../../api/token/configmap/configmap";
import {doTo} from "../../router";
export default defineComponent({
  name: 'configmap-detail',
  components: {MainLayout},
  setup(){

    let state=reactive({
      configmapData:reactive({}),
    })
    let data=computed(()=>{
      let arr=[];
      for ( let i in state.configmapData.data){
        arr.push({"key":i,value:state.configmapData.data[i]})
      }
      return arr
    })
    async function getData(){
      const route = useRoute()
      try {
       let tData=await configmapDetail(route.query.namespace,route.query.name)
        state.configmapData=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    getData()
    return {...toRefs(state),data,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>