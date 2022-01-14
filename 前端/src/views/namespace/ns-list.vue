<template>
   <main-layout>
     <breadcrumb title="命名空间.列表"></breadcrumb>
     <el-button type="primary" class="right mtb20" @click="dialogFormVisible=true" >创建命名空间</el-button>
     <el-divider class="clear_both"></el-divider>
     <div class="top-list">
       <el-table :data="nsList">
           <el-table-column label="命名空间名称">
             <template #default="scope">
               <span>{{ scope.row.name }}<em class="Danger mf20">{{ scope.row.name=='wike-system'?"系统命名空间请不要删除":""}}</em></span>
             </template>

           </el-table-column>
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
import MainLayout from "@/layout/main.vue";
import {createNs, deleteNs, getNsList} from "@/api/token/namespace/ns";
import {requireRules} from "../../helper/rules";
import {ElMessage, ElMessageBox} from "element-plus";
import breadcrumb from  "@/components/common/breadcrumb.vue"
import {rmTip} from "@/helper/helper";
export default defineComponent({
  name: 'namespace',
  components: {
    MainLayout,
    breadcrumb
  },
  setup(){
    const ws = inject("ws") //通信用的
    let state=reactive({
      nsList:[],//ns数据
      dialogFormVisible:false, //弹窗
      form:{
        name:""  //表单
      }
    })
    const formRef=ref(null) //refs引用
    async function fetchData(){
      try {
       let tData=await getNsList()  //获取数据
        state.nsList=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    fetchData()


    ws.onmessage = (e)=>{
      if(e.data !== 'ping'){
        let tData=JSON.parse(e.data)
        if(tData.type=='namespace'){
          state.nsList=tData.result
        }
      }
    }
    function create(){
      formRef.value.validate(async (valid) => { //验证数据
        if(valid){
          try {
            let result= await  createNs({name:state.form.name})  //提交数据
            if (result.data.code==200){
              ElMessage("资源创建成功")
              state.form.name=""
              state.dialogFormVisible=false
            }
          }catch (e){

          }
        }
      })
    }
    function doDelete(name){
      rmTip("命名空间")
          .then(async () => {
            try {
              if (name=="wike-system"){
                ElMessage.error("此命名空间不推荐删除")
                return
              }
              await deleteNs(name)
            }catch (e) {
            }
          })
    }
    return {...toRefs(state),requireRules,create,doDelete,formRef}
  }
})
</script>
