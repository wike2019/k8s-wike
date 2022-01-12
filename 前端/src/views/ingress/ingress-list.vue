<template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>ingress列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <div class="top-list">
         <div class="name_space_choose">
           <span>命名空间:</span>
           <el-select placeholder="选择命名空间" @change="changeNs" filterable v-model="namespace">
             <el-option v-for="ns in nsList "
                        :label="ns.name"
                        :value="ns.name"/>
           </el-select>
           <el-button style="margin-left:40px" type="primary" @click="doTo('ingress-create')">创建Ingress</el-button>
         </div>
         <el-divider></el-divider>
                  <el-table
                      :data="ingressList"
                      border
                      empty-text="暂无数据"
                      style="width: 100%">
                    <el-table-column
                        prop="status"
                        label="状态"
                        width="120">
                      <template #default="scope">
                        <span  class="Success"  v-if="scope.row.is_ready">active</span>
                        <span  class="Danger" v-else>正在初始化</span>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="name"
                        label="名称"
                        >
                      <template #default="scope">
                        <p>{{scope.row.name}}</p>
                      </template>
                    </el-table-column>
                    <el-table-column label="域名" width="200">
                      <template #default="scope">
                        <p><a target="_blank" :href="prefix(scope.row)" >{{ prefix(scope.row) }}</a></p>
                      </template>
                    </el-table-column>
                    <el-table-column label="配置" width="140">
                      <template #default="scope">
                       <p>
                         <template v-if="config(scope.row.annotations,'nginx.ingress.kubernetes.io/enable-cors')">
                           <el-checkbox  :checked="true" disabled>跨域</el-checkbox>
                         </template>
                         <template v-else>
                           <el-checkbox  :checked="false" disabled>跨域</el-checkbox>
                         </template>
                       </p>
                        <p>
                          <template v-if="config(scope.row.annotations,'nginx.ingress.kubernetes.io/rewrite-target')">
                            <el-checkbox  :checked="true" disabled>重写</el-checkbox>
                          </template>
                          <template v-else>
                            <el-checkbox  :checked="false" disabled>重写</el-checkbox>
                          </template>
                        </p>
                      </template>
                    </el-table-column>
                    <el-table-column
                        prop="name_space"
                        label="命名空间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column
                        prop="create_time"
                        label="创建时间"
                        width="180"
                    >
                    </el-table-column>
                    <el-table-column label="操作" width="300" align="center">
                      <template #default="scope">
                        <el-button
                            type="primary"
                            size="small"
                            @click="()=>doTo('ingress-update',{namespace:scope.row.namespace,name:scope.row.name})"  >
                          编辑
                        </el-button>
                        <el-button
                            type="info"
                            size="small"
                            @click="()=>doTo('ingress-detail',{namespace:scope.row.namespace,name:scope.row.name})"  >
                          查看详情
                        </el-button>
                        <el-button type="danger" size="small" @click="()=>rmIngress(scope.row.namespace,scope.row.name )" >删除</el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                <div class="page_show">
                  <el-pagination
                      @current-change="pageChange($event)"
                      background
                      :current-page="current_page"
                      :page-size="10"
                      :hide-on-single-page="true"
                      layout="prev, pager, next"
                      :total="pageInfo">
                  </el-pagination>
                </div>
       </div>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../../layout/main.vue";
import {getNsList} from "../../api/token/namespace/ns";
import {ingressDel, ingressListByNs} from "../../api/token/ingress/ingress";
import {secretListByNs} from "../../api/token/secret/secret";
import {doTo} from "../../router";
export default defineComponent({
  name: 'ingress-list',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nsList:[],
      ingressList:[],
      pageInfo:0,
      current_page:0,
      namespace:"default"
    })

    async function fetchData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
        for (let i=0;i<state.nsList.length;i++){
          await  ingressList(state.namespace)
        }
      }catch (e){
        console.log(e)
      }
    }

    async function ingressList(ns,page=1){
      try {
        let tData=await ingressListByNs(ns,page)
        let T=tData.data.data
        state.ingressList=T.Data
        state.pageInfo=T["Total"]
        state.current_page=T["Current"]
      }catch (e){
        console.log(e)
      }
    }

    fetchData()
    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)

        if(tData.type=='ingress'&&T.Ext.ns==state.namespace){
          let T=tData.result;
          state.ingressList=T.Data
          state.pageInfo=T["Total"]
          state.current_page=T["Current"]
        }
      }
    }

    function pageChange(page,ns){
      state.current_page=page
      ingressList(ns,state.current_page)
    }
    function rmIngress(ns,name){
      ElMessageBox.confirm(
          '你确定继续删除ingress操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
      )
          .then(async () => {
              try {
                await ingressDel(ns,name)
              }catch (e) {
                ElMessage({
                  type: 'info',
                  message: '系统异常'+e,
                })
              }
          })
    }
    function prefix(row){
      if (!row.host){
        return  "无域名"
      }
      return row.isHttps?"https://"+row.host:"http://"+row.host;
    }
    function config(data,match){
      if(data&&data[match]){
        return true
      }
      return false
    }

    async function changeNs(){
      try {
        let tData=await ingressListByNs(state.namespace )
        let T=tData.data.data
        console.log(T)
        state.ingressList=T.Data
        state.pageInfo=T["Total"]
        state.current_page=T["Current"]
      }catch (e){
        console.log(e)
      }

    }
    return {...toRefs(state),pageChange,rmIngress,prefix,config,changeNs,doTo}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>