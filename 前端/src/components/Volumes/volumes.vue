<template>
  <el-card class="box-card">
    <template #header>
      <div class="card-header">
        数据卷配置
        <el-button size="small" style="margin-left:20px" @click="show=!show">
          {{!show?"展开":"收缩"}}
        </el-button>
        <el-button type="success" size="small" style="float: right;margin-right:50px" @click="addVolumes" >添加数据卷</el-button>
      </div>
    </template>
  {{form.volumes}}
    <div v-show="show">
  <el-form :inline="true"  :model="form" ref="formRef"  >
    <div v-for="(volume,index) in form.volumes" class="mtb20">
      <!--数据名称-->
      <el-form-item
          label="数据卷名称"
          :rules="requireRules('数据卷名称必须填写')"
          :key="'volumes_'+index+'_name'"
          :prop="'volumes.'+index+'.name'"
      >
        <el-input v-model="volume.name" placeholder="数据卷名称"  />
      </el-form-item>
      <!--数据卷类型，根据不同的类型显示不同数据 辅助字段-->
      <el-form-item
          label="数据卷类型"
          :key="'volumes_'+index+'_type'"
          :prop="'volumes.'+index+'.type'"
      >
        <el-select  size="small"
                    v-model="volume.type" class="select" @change="initType($event,index)">
          <el-option label="hostPath" value="hostPath"/>
          <el-option label="emptyDir" value="emptyDir"/>
          <el-option label="secret" value="secret"/>
          <el-option label="persistentVolumeClaim" value="persistentVolumeClaim"/>
          <el-option label="downwardAPI" value="downwardAPI"/>
          <el-option label="configMap" value="configMap"/>
        </el-select>
      </el-form-item>


      <el-form-item  >
        <el-button size="small" type="danger" @click="removeVolumes(index)" >删除数据卷</el-button>
      </el-form-item>


      <div v-if="volume.type=='hostPath'">
        <el-form-item
            label="主机目录类型"
            :rules="inArrayWithMsg(['','DirectoryOrCreate','Directory','FileOrCreate','File','configMap','Socket','CharDevice','BlockDevice'],'type类型不正确')"
            :key="'host_'+index+'_type'"
            :prop="'volumes.'+index+'.hostPath.type'"
        >
            <el-select  size="small" v-model="volume.hostPath.type" class="select">
              <el-option label="任意" value=""/>
              <el-option label="DirectoryOrCreate" value="DirectoryOrCreate"/>
              <el-option label="Directory" value="Directory"/>
              <el-option label="FileOrCreate" value="FileOrCreate"/>
              <el-option label="File" value="File"/>
              <el-option label="configMap" value="configMap"/>
              <el-option label="Socket" value="Socket"/>
              <el-option label="CharDevice" value="CharDevice"/>
              <el-option label="BlockDevice" value="BlockDevice"/>
            </el-select>
        </el-form-item>

        <el-form-item
            label="主机目录路径"
            :key="'host_'+index+'_path'"
            :prop="'volumes.'+index+'.hostPath.path'"
            :rules="requireRules('主机目录必须填写')"
        >
          <el-input v-model="volume.hostPath.path"  placeholder="主机目录路径" />
        </el-form-item>
      </div>

      <div v-if="volume.type=='configMap'">
        <el-form-item
            label="配置名称"
            :key="'configMap_'+index+'_name'"
            :prop="'volumes.'+index+'.configMap.name'"
            :rules="requireRules('配置名称必须填写')"
        >

          <el-select  size="small" v-model="volume.configMap.name" class="select">
            <el-option :key="item.name"  v-for="item in configMap" :label="item.name" :value="item.name"/>
          </el-select>
        </el-form-item>

        <el-form-item
            label="文件默认权限"
            :key="'configMap_'+index+'_defaultMode'"
            :prop="'volumes.'+index+'.configMap.defaultMode'"
        >
          <el-input v-model="volume.configMap.defaultMode" placeholder="文件默认权限"/>
        </el-form-item>

        <el-form-item
            label="是否可选"
            :key="'configMap_'+index+'_optional'"
            :prop="'volumes.'+index+'.configMap.optional'"
        >
          <el-select  size="small" v-model="volume.configMap.optional"  class="select">
            <el-option label="false"  :value="false"/>
            <el-option label="true"   :value="true"/>
          </el-select>
        </el-form-item>
        <el-form-item  >
          <el-button size="small" type="success" @click="addConfigMapItems(index)" >添加子items</el-button>
        </el-form-item>
        <div class="dashed"></div>
        <div v-for="(configMap,index2) in volume.configMap.items"  class="mtb20">
          <el-form-item
              label="配置key"
              :rules="requireRules('配置key必须填写')"
              :key="'configMap_'+index+'_items_'+index2+'_key'"
              :prop="'volumes.'+index+'.configMap.items.'+index2+'.key'"
          >
            <el-input v-model="configMap.key" placeholder="配置key" />

          </el-form-item>

          <el-form-item
              label="配置路径"
              :rules="requireRules('配置路径必须填写')"
              :key="'configMap_'+index+'_items_'+index2+'_path'"
              :prop="'volumes.'+index+'.configMap.items.'+index2+'.path'"
          >
            <el-input v-model="configMap.path"  placeholder="配置路径" />
          </el-form-item>
          <el-form-item
              label="配置模式"
              :key="'configMap_'+index+'_items_'+index2+'_mode'"
              :prop="'volumes.'+index+'.configMap.items.'+index2+'.mode'"   >
            <el-input v-model="configMap.mode" placeholder="配置模式"  />
          </el-form-item>
          <el-form-item  >
            <el-button size="small" type="danger" @click="removeConfigMapItems(index,index2)" >删除该项items</el-button>
          </el-form-item>
        </div>

      </div>

      <div v-if="volume.type=='secret'">
        <el-form-item
            label="密文名称"
            :key="'secret_'+index+'_secretName'"
            :prop="'volumes.'+index+'.secret.secretName'"
            :rules="requireRules('密文名称必须填写')"
        >
          <el-select  size="small" v-model="volume.secret.secretName" class="select">
            <el-option :key="item.name" v-for="item in secret" :label="item.name" :value="item.name"/>
          </el-select>
        </el-form-item>
        <el-form-item
            label="文件默认权限"
            :key="'secret_'+index+'_defaultMode'"
            :prop="'volumes.'+index+'.secret.defaultMode'"
        >
          <el-input v-model="volume.secret.defaultMode" placeholder="文件默认权限"/>
        </el-form-item>

        <el-form-item
            label="是否可选"
            :key="'secret_'+index+'_optional'"
            :prop="'volumes.'+index+'.secret.optional'"
        >
          <el-select  size="small" v-model="volume.secret.optional"  class="select">
            <el-option label="false"  :value="false"/>
            <el-option label="true"   :value="true"/>
          </el-select>
        </el-form-item>

        <el-form-item  >
          <el-button size="small" type="success" @click="addSecretItems(index)" >添加子items</el-button>
        </el-form-item>
        <div class="dashed"></div>
        <div v-for="(secretItem,index2) in volume.secret.items"  class="mtb20">
          <el-form-item
              label="密文key"
              :rules="requireRules('密文key必须填写')"
              :key="'secret_'+index+'_items_'+index2+'_key'"
              :prop="'volumes.'+index+'.secret.items.'+index2+'.key'"
          >

            <el-input v-model="secretItem.key" placeholder="密文key"/>

          </el-form-item>
          <el-form-item
              label="密文路径"
              :rules="requireRules('密文路径必须填写')"
              :key="'secret_'+index+'_items_'+index2+'_path'"
              :prop="'volumes.'+index+'.secret.items.'+index2+'.path'"
          >

            <el-input v-model="secretItem.path" placeholder="密文路径"/>

          </el-form-item>
          <el-form-item
              label="密文模式"
              :key="'secret_'+index+'_items_'+index2+'_mode'"
              :prop="'volumes.'+index+'.secret.items.'+index2+'.mode'"
          >

            <el-input v-model="secretItem.mode" placeholder="密文模式" />

          </el-form-item>

          <el-form-item  >
            <el-button size="small" type="danger" @click="removeSecretItems(index,index2)" >删除该项items</el-button>
          </el-form-item>
        </div>
      </div>



      <div v-if="volume.type=='persistentVolumeClaim'">
        <el-form-item
            label="pvc名称"
            :key="'persistentVolumeClaim_'+index+'_name'"
            :prop="'volumes.'+index+'.persistentVolumeClaim.claimName'"
            :rules="requireRules('pvc名称必须填写')"
        >
          <el-select  size="small" v-model="volume.persistentVolumeClaim.claimName" class="select">
            <el-option :key="item.metadata.name"  v-for="item in persistentVolumeClaim" :label="item.metadata.name" :value="item.metadata.name"/>
          </el-select>
        </el-form-item>
        <el-form-item label="是否只读"
                      :key="'persistentVolumeClaim_'+index+'readOnly'"
                      :prop="'volumes.'+index+'.persistentVolumeClaim.readOnly'"
        >

          <el-select  size="small" v-model="volume.persistentVolumeClaim.readOnly"  class="select">
            <el-option label="false"  :value="false" />
            <el-option label="true"   :value="true"  />
          </el-select>

        </el-form-item>
      </div>

      <div v-if="volume.type=='downwardAPI'">
        <el-form-item
            label="文件默认权限"
            :key="'downwardAPI_'+index+'_defaultMode'"
            :prop="'volumes.'+index+'.downwardAPI.defaultMode'"
        >
          <el-input v-model="volume.downwardAPI.defaultMode" placeholder="文件默认权限"/>
        </el-form-item>
        <el-form-item  >
          <el-button size="small" type="success" @click="addDownwardAPIItems(index)" >添加子项</el-button>
        </el-form-item>
        <div class="dashed"></div>
        <div v-for="(downwardAPIItem,index2) in volume.downwardAPI.items" style="margin:20px 0">
          <el-form-item
              label="路径"
              :rules="requireRules('路径必须填写')"
              :key="'downwardAPI_'+index+'_items_'+index2+'_path'"
              :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.path'"
          >

            <el-input v-model="downwardAPIItem.path" placeholder="路径"/>

          </el-form-item>

          <el-form-item  >
            <el-button size="small" type="danger" @click="removeDownwardAPIItems(index,index2)" >删除该项items</el-button>
          </el-form-item>

          <div class="dashed"></div>
          <h2 class="mtb20">子项设置元数据对象</h2>
          <div v-if="downwardAPIItem.fieldRef">
          <el-form-item
              label="元数据对象-fieldRef.apiVersion"
              :key="'downwardAPI_'+index+'_items_'+index2+'.fieldRef.apiVersion'"
              :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.fieldRef.apiVersion'"
          >
            <el-input v-model="downwardAPIItem.fieldRef.apiVersion" placeholder="元数据对象-fieldRef.apiVersion" class="longWidth"/>
          </el-form-item>
          <el-form-item
              label="元数据对象-fieldRef.FieldPath"
              :key="'downwardAPI_'+index+'_items_'+index2+'.fieldRef.FieldPath'"
              :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.fieldRef.FieldPath'"
          >
            <el-input v-model="downwardAPIItem.fieldRef.FieldPath" placeholder="元数据对象-fieldRef.FieldPath" class="longWidth" />
          </el-form-item>
          </div>
          <el-divider></el-divider>
          <h2  class="mtb20">子项设置资源对象</h2>
          <div v-if="downwardAPIItem.resourceFieldRef">
            <el-form-item
                label="资源对象-resourceFieldRef.containerName"
                :key="'downwardAPI_'+index+'_items_'+index2+'.resourceFieldRef.containerName'"
                :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.resourceFieldRef.containerName'"
            >
              <el-input v-model="downwardAPIItem.resourceFieldRef.containerName" placeholder="资源对象-resourceFieldRef.containerName" class="longWidth" />
            </el-form-item>
            <el-form-item
                label="资源对象-resourceFieldRef.resource"
                :key="'downwardAPI_'+index+'_items_'+index2+'.resourceFieldRef.resource'"
                :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.resourceFieldRef.resource'"
            >
              <el-input v-model="downwardAPIItem.resourceFieldRef.resource" placeholder="资源对象-resourceFieldRef.resource" class="longWidth"/>
            </el-form-item>
            <el-form-item
                label="资源对象单位"
                :key="'downwardAPI_'+index+'_items_'+index2+'.resourceFieldRef.divisor'"
                :prop="'volumes.'+index+'.downwardAPI.items.'+index2+'.resourceFieldRef.divisor'"
            >
              <el-input v-model="downwardAPIItem.resourceFieldRef.divisor" placeholder="资源对象-resourceFieldRef.divisor"/>
            </el-form-item>
          </div>

        </div>
      </div>
      <el-divider></el-divider>
    </div>
  </el-form>
    </div>
  </el-card>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {requireRules,inArrayWithMsg} from "../../helper/rules.ts"
import {pvcAllByNs} from "../../api/token/pvc";
import {secretAllByNs} from "../../api/token/secret/secret";
import {configmapAllByNs} from "../../api/token/configmap/configmap";
export default defineComponent({
  name: 'volumes',
  emits:["input"],
  props:["namespace"],
  setup: (props,{emit}) => {
    const formRef=ref(null)
    let state=reactive({
      show:false,
      secret:[],
      configMap:[],
      persistentVolumeClaim:[],
      form:{
        volumes:[]
      }
    })
    fetchData(props.namespace)
    function initType(config,index) {
      let data={}
      data.name=state.form.volumes[index].name
      if("hostPath"==config){
        data.type='hostPath'
        data.hostPath={"path":"",type:"Directory"}
      }
      if("emptyDir"==config){
        data.type='emptyDir'
        data.emptyDir={}

      }
      if("secret"==config){
        data.type='secret'
        data.secret={"secretName":"","defaultMode":"0644",optional:false,items:[]}
      }
      if("persistentVolumeClaim"==config){
        data.type='persistentVolumeClaim'
        data.persistentVolumeClaim={"claimName":"",readOnly:false}
      }
      if("downwardAPI"==config){
        data.type='downwardAPI'
        data.downwardAPI={"defaultMode":"0644",items:[]}
      }
      if("configMap"==config){
        data.type='configMap'
        data.configMap={"name":"","defaultMode":"0644",optional:false,items:[]}
      }
      state.form.volumes.splice(index,1,data)

    }
    //添加数据卷
    function addVolumes(){
      state.show=true
      //默认 hostPath 类型
      state.form.volumes.push({"name":"","type":"hostPath","hostPath":{"path":"","type":"Directory"}})
    }
    //删除数据卷
    function removeVolumes(index) {
      state.form.volumes.splice(index,1)
    }
    //删除ConfigMap的Items
    function removeConfigMapItems(index,index2) {
      state.form.volumes[index].configMap.items.splice(index2,1)

    }
    //添加ConfigMap的Items
    function addConfigMapItems(index) {
      if( !state.form.volumes[index].configMap.items){
        state.form.volumes[index].configMap.items=[];
      }
      state.form.volumes[index].configMap.items.push({mode:"0644",path:"",key:""})
    }
    //添加Secret的Items
    function addSecretItems(index) {
      if( !state.form.volumes[index].secret.items){
        state.form.volumes[index].secret.items=[];
      }
      state.form.volumes[index].secret.items.push({mode:"0644",path:"",key:""})
    }
    //删除Secret的Items
    function removeSecretItems(index,index2) {
      state.form.volumes[index].secret.items.splice(index2,1)

    }
    //添加DownwardAPI的Items
    function addDownwardAPIItems(index) {
      if( !state.form.volumes[index].downwardAPI.items){
        state.form.volumes[index].downwardAPI.items=[];
      }
      state.form.volumes[index]['downwardAPI'].items.push({mode:"0644",path:"",fieldRef:{apiVersion:"",FieldPath:""},resourceFieldRef:{containerName:"",resource:"",divisor:1}})
    }
    //删除DownwardAPI的Items
    function removeDownwardAPIItems(index,index2) {
      state.form.volumes[index]['downwardAPI'].items.splice(index2,1)

    }
    function commit() {
      let volumes=[]
      state.form.volumes.forEach((item)=>{
        if(item.type=='hostPath'){
          volumes.push({"hostPath":item.hostPath,"name":item.name})
        }
        if(item.type=='emptyDir'){
          volumes.push({"emptyDir":item.emptyDir,"name":item.name})
        }
        if(item.type=='secret'){
          let secret={}
          secret.secretName=item.secret.secretName
          secret.defaultMode=item.secret.defaultMode
          secret.optional=item.secret.optional
          if(item.secret.items&&item.secret.items.length){
            secret.items=item.secret.items
          }
         volumes.push({"secret":secret,"name":item.name})
        }
        if(item.type=='persistentVolumeClaim'){
          volumes.push({"persistentVolumeClaim":item.persistentVolumeClaim,"name":item.name})
        }
        if("downwardAPI"==item.type){
          let downwardAPI={}
          downwardAPI.defaultMode=item.downwardAPI.defaultMode
          if(item.downwardAPI.items&&item.downwardAPI.items.length){
            downwardAPI.items=[]
            for (let i=0;i<item.downwardAPI.items.length;i++){
              let temp={}
              if(item.downwardAPI.items[i].fieldRef&&(item.downwardAPI.items[i].fieldRef.apiVersion!=""|| item.downwardAPI.items[i].fieldRef.FieldPath!="")){
                temp.fieldRef=item.downwardAPI.items[i].fieldRef
              }
              if(item.downwardAPI.items[i].resourceFieldRef &&(item.downwardAPI.items[i].resourceFieldRef?.containerName!=""|| item.downwardAPI.items[i].resourceFieldRef?.resource!="")){
                temp.resourceFieldRef=item.downwardAPI.items[i].resourceFieldRef
              }
              temp.mode=item.downwardAPI.items[i].mode
              temp.path=item.downwardAPI.items[i].path
              downwardAPI.items.push(temp)
            }
          }

          volumes.push({"downwardAPI":downwardAPI,"name":item.name})
        }
        if("configMap"==item.type){
          let configMap={}
          configMap.name=item.configMap.name
          configMap.defaultMode=item.configMap.defaultMode
          configMap.optional=item.configMap.optional
          if(item.configMap.items&&item.configMap.items.length){
            configMap.items=item.configMap.items
          }
          volumes.push({"configMap":configMap,"name":item.name})
        }
      })
      console.error(JSON.stringify(volumes),"8888888")
      emit("input",volumes)
    }
    watch(()=>props.namespace,()=>{
      fetchData(props.namespace)
    })
    //数据改变了触发提交
    watch(()=>state.form,()=>{
      commit()
    },{deep:true,flush:"post"})


    function setData(volumes){
      let arr=[]
      volumes.forEach(item=>{
        if(item.hostPath){
          arr.push({"hostPath":item.hostPath,"name":item.name,type:'hostPath'})
        }
        if(item.emptyDir){
          arr.push({"emptyDir":item.emptyDir,"name":item.name,type:'emptyDir'})
        }
        if(item.secret){
              let secret={}
              secret.secretName=item.secret.secretName
              secret.defaultMode=item.secret.defaultMode
              secret.optional=item.secret.optional
              if(item.secret.items){
                secret.items=item.secret.items
              }
          arr.push({"secret":secret,"name":item.name,type:'secret'})
        }
        if(item.persistentVolumeClaim){
          arr.push({"persistentVolumeClaim":item.persistentVolumeClaim,"name":item.name,type:'persistentVolumeClaim'})
        }
        if(item.downwardAPI){

          let downwardAPI={}
          if(item.downwardAPI.items){
            downwardAPI.items=item.downwardAPI.items
            for (let i=0;i<downwardAPI.items.length;i++){
              if(!downwardAPI.items[i].fieldRef){
                downwardAPI.items[i].fieldRef={apiVersion:"",FieldPath:""}
              }
              if(!downwardAPI.items[i].resourceFieldRef){
                downwardAPI.items[i].resourceFieldRef={containerName:"",resource:"",divisor:1}
              }
            }
          }
          downwardAPI.defaultMode=item.downwardAPI.defaultMode

          arr.push({"downwardAPI":downwardAPI,"name":item.name,type:'downwardAPI'})
        }
        if(item.configMap){
          let configMap={}
          configMap.secretName=item.configMap.name
          configMap.defaultMode=item.configMap.defaultMode
          configMap.optional=item.configMap.optional
          if(item.configMap.items){
            configMap.items=item.configMap.items
          }
          arr.push({"configMap":configMap,"name":item.name,type:'configMap'})
        }
      })
      state.form.volumes=arr
      commit()
      console.log("数据设置成功",arr)
    }
    function Check(){
      return new Promise(function (resolve){
        formRef.value.validate(async (valid) => {
          if(valid){
            resolve(true)
          }else{
            resolve(false)
          }
        })
     })
   }
   function clearInput(){
     let arr=[]
     state.form.volumes.forEach((item,index)=>{

       if(item.type=='secret'){
         arr.push({type:"secret","secret":{"secretName":"","defaultMode":"0644",optional:false,items:[]},"name":item.name})
         return;
       }
       if(item.type=='persistentVolumeClaim'){
         arr.push({type:"persistentVolumeClaim","persistentVolumeClaim":{"claimName":"",readOnly:false},"name":item.name})
         return;
       }
       if("configMap"==item.type){
         arr.push({type:"configMap","configMap":{"name":"","defaultMode":"0644",optional:false,items:[]},"name":item.name})
         return;
       }

       arr.push(item)
     })
     state.form.volumes=arr
   }
   async function fetchData(ns){
   // alert(ns)
     try {
       let pvc=await  pvcAllByNs(ns)
       let secret=await  secretAllByNs(ns)
       let configmap=await  configmapAllByNs(ns)
       //console.log(pvc.data.data)
       state.persistentVolumeClaim=pvc.data.data
       state.secret=secret.data.data
       state.configMap=configmap.data.data
       clearInput()
     }catch (e){
       console.log(e)
     }

   }

    return {
      ...toRefs(state),
      Check,formRef,requireRules,initType,addVolumes,removeVolumes,inArrayWithMsg,removeConfigMapItems,addConfigMapItems,addSecretItems,removeSecretItems,addDownwardAPIItems,removeDownwardAPIItems,setData}
  }
})
</script>

<style scoped>

</style>
