   <template>
   <main-layout class="fix-black">
     <nav class="nav-bar">
          <el-breadcrumb separator="/">
             <el-breadcrumb-item>Role创建</el-breadcrumb-item>
          </el-breadcrumb>
     </nav>
     <el-form :inline="true" rel="formRef" :model="form" ref="formRef" :rules="rules">
     <el-card class="box-card">
       <template #header>
         <div class="card-header">
           基本配置
         </div>
       </template>
       <el-form-item label="角色绑定名称"  prop="name"  >
         <el-input v-model="form.name"   placeholder="角色绑定名称"></el-input>
       </el-form-item>
       <el-form-item label="命名空间" prop="name_space">
         <el-select  v-model="form.name_space" @change="updateRole">
           <el-option v-for="ns in nsList "
                      :label="ns.name"
                      :value="ns.name"/>
         </el-select>
       </el-form-item>
       <el-form-item label="选择角色" prop="role" >
         <el-select  v-model="form.role">
           <el-option v-for="role in roleList "
                      :label="role.name"
                      :value="role.name"/>
         </el-select>
       </el-form-item>
       <el-form-item >
           <el-button style="margin-left:20px"  type="primary" @click="post">创建角色绑定</el-button>
       </el-form-item>
     </el-card>
       <el-card class="box-card">
           <template #header>
             <div class="card-header" style="padding-bottom:10px">
               选择用户     <el-button type="primary" size="small" style="float: right;margin-right:50px" @click="addSubject" >+</el-button>
             </div>
           </template>
           <div   v-for="(sub,index) in form.subjects" style="margin-top: 20px">
             <el-form-item  label="用户类型" :key="index"  :prop="'subjects.' + index + '.kind'"
                           :rules="{required: true, trigger: 'blur', message: '用户类型必须填写' }">
               <el-select
                   v-model="sub.kind"
                   placeholder="用户类型"  >
                 <el-option
                     label="User"
                     value="User" >
                 </el-option>
                 <el-option
                     label="ServiceAccount"
                     value="ServiceAccount" >
                 </el-option>
                 <el-option
                     label="Group"
                     value="Group" >
                 </el-option>
               </el-select>
             </el-form-item>
             <el-form-item label="用户账号"   :prop="'subjects.' + index + '.name'"  :rules="{required: true, trigger: 'blur', message: '账号必须选择' }"   v-show="sub.kind=='ServiceAccount'">

               <el-select v-model="sub.name"   size="small"  style="margin:0 8px;width: 180px">
                 <el-option :key='sa.name' v-for="sa in SaList "
                            :label="sa.name"
                            :value="sa.name">
                 </el-option>
               </el-select>
             </el-form-item>
             <el-form-item style="vertical-align: middle" label="用户名称或组名"  :rules="{required: true, trigger: 'blur', message: '用户名称必须填写' }" v-show="sub.kind!='ServiceAccount'"   :prop="'subjects.' + index + '.name'" >
               <el-input  style="margin:0 8px;width: 180px"  size="small" placeholder="输入用户名" v-model="sub.name" ></el-input>
             </el-form-item>
             <el-button type="primary" size="small" style="margin-top:-20px"  @click="rmSubject(index)">
              -</el-button>
           </div>
       </el-card>
     </el-form>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, computed, ref, onUnmounted, inject, reactive,toRefs} from 'vue'
import { ElMessageBox, ElMessage } from 'element-plus'
import MainLayout from "../layout/main.vue";
import {getNsList} from "../api/token/namespace/ns";
import {ingressDel, ingressListByNs} from "../api/token/ingress/ingress";
import {secretDel, secretListByNs} from "../api/token/secret/secret";
import {doTo} from "../router";
import {
  addUserToBinding,
  getRoleAllList,
  getRoleList,
  postRoleBindings,
  roleBindingCreate,
  roleCreate
} from "../api/token/rbac";
import Resources from "../components/Resources/Resources.vue";
import {getSaList} from "../api/token/sa/sa";

export default defineComponent({
  name: 'role-create',
  components: {MainLayout,Resources},
  setup(){
    const apiGroup='rbac.authorization.k8s.io'
    let state=reactive({
      nsList:reactive([]),
      SaList:reactive([]),
      form:reactive({
        name:"",
        name_space:"default",
        role:"",
        subjects:reactive([  //前端所使用的rule
          {apiGroup,kind:'User',name:''}
        ]),
      }),
      roleList:reactive([]),
    })
    let rules={
      name:[
        {
          required: true,
          message: '角色绑定名称',
        }
      ],
      name_space:[
        {
          required: true,
          message: '命名空间必须填写',
          trigger: 'blur',
        }
      ],
      role:[
        {
          required: true,
          message: '角色必须选择',
          trigger: 'change',
        }
      ],
    }
    let ResourcesRef=ref(null)
    let formRef=ref(null)
    async function getData(){
      try {
       let tData=await getNsList()
        state.nsList=tData.data.data
        await updateRole()
      }catch (e){
        console.log(e)
      }
    }
    async function updateRole(){

      try {
        let tData= await getRoleAllList(state.form.name_space)
        state.roleList=tData.data.data
        let tData2=await getSaList(state.form.name_space)
        state.SaList=tData2.data.data
        state.form.subjects=[  //前端所使用的rule
          {apiGroup,kind:'User',name:''}
        ]
      }catch (e){
        console.log(e)
      }
    }
    getData()
    function post() {
      formRef.value.validate(async (valid) => {
        if (valid) {
          let subjects=[]
          state.form.subjects.forEach(function (item) {
            if(item.kind=="ServiceAccount"){
              item.apiGroup=""
            }
            subjects.push(item)
          })
          let postData={metadata:{name:state.form.name,namespace:state.form.name_space},roleRef:{apiGroup,kind:"Role",name:state.form.role},subjects:subjects}

          try {
            let result= await  postRoleBindings(postData)
            if (result.data.code==200){
              ElMessage("角色资源创建成功")
            }

          }catch (e){
            console.log(e)
          }

        }
      })
    }
    function addSubject(){
      state.form.subjects.unshift({apiGroup,kind:'User',name:''})
    }

    function rmSubject(index){
      state.form.subjects.splice(index,1)
    }

    return {...toRefs(state),rules,doTo,ResourcesRef,post,formRef,updateRole,addSubject,rmSubject}
  }
})
</script>
<style >
.fix-black .el-checkbox__input.is-disabled.is-checked .el-checkbox__inner {
  background-color:#555;
}
</style>