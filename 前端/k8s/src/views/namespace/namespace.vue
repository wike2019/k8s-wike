<template>
   <main-layout>
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>命名空间列表</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-button type="primary" class="right mtb20" @click="dialogFormVisible=true" >创建命名空间</el-button>
     <el-divider class="clear_both"></el-divider>
     <div class="top-list">
       <el-table :data="nsList">
           <el-table-column prop="name" label="命名空间名称"></el-table-column>
           <el-table-column align="center" prop="create_time" label="创建时间" width="220"></el-table-column>
           <el-table-column label="操作" width="120" align="center">
           <template #default="scope">
             <el-button size="small" type="danger"   @click="()=>doDelete(scope.row.name)" >删除</el-button>
           </template>
         </el-table-column>
       </el-table>
     </div>
   </main-layout>
  <el-dialog v-model="dialogFormVisible" title="创建命名空间">
    <el-form :model="form" ref="formRef">
      <el-form-item label="命名空间名称" prop="name"    :rules="requireRules('命名空间名称必填')">
        <el-input v-model="form.name" ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="create">创建</el-button>
        <el-button @click="dialogFormVisible = false">取消</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts">
import {defineComponent,computed ,ref,onUnmounted, inject,reactive,toRefs } from 'vue'
import MainLayout from "../../layout/main.vue";
import {createNs, deleteNs, getNsList} from "../../api/token/namespace/ns";
import {requireRules} from "../../helper/rules";
import {ingressCreate} from "../../api/token/ingress";
import {ElMessage, ElMessageBox} from "element-plus";
import {configmapDel} from "../../api/token/configmap";
export default defineComponent({
  name: 'namespace',
  components: {MainLayout},
  setup(){

    const ws = inject("ws")
    let state=reactive({
      nsList:reactive([]),
      dialogFormVisible:false,
      form:{
        name:""
      }
    })
    const formRef=ref(null)
    async function getData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    getData()


    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        if(tData.type=='namespace'){
          state.nsList=tData.result
        }
      }
    }
    function create(){
      formRef.value.validate(async (valid) => {
        if(valid){
          try {
            let result= await  createNs({name:state.form.name})
            if (result.data.code==200){
              ElMessage("资源创建成功")
              state.form.name=""
              state.dialogFormVisible=false
            }

          }catch (e){
            ElMessage.error("资源格式有误:"+e)
          }
        }
      })
    }
    function doDelete(name){
      ElMessageBox.confirm(
          '你确定继续删除操作吗?',
          'Warning',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
          }
          )
          .then(async () => {
            try {
              await deleteNs(name)
            }catch (e) {
              ElMessage.error('系统异常:'+e)
            }
          })
    }
    return {...toRefs(state),requireRules,create,doDelete,formRef}
  }
})
</script>
