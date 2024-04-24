<template>
  <div  v-if="!render">
    <div class="inline-title-no" >
      <span>端口配置</span>   <el-button type="primary" size="small"  class="leftBtn" @click="addItem" >添加端口</el-button>
    </div>
    <el-form :inline="true"  :model="state.form" ref="formRef"  >
      <div v-for="(port,index) in state.form"  class="flex">
        <el-form-item class="flex1"  label="端口名称"   :key="'name'+index"  :prop="index + '.name'"
                      :rules="{required: true, trigger:'blur', message: 'name必须填写' }">
          <el-input     class="top10"  v-model="port.name"  placeholder="name必须填写"></el-input>
        </el-form-item>
        <el-form-item class="flex1"  label="协议"   :key="'protocol'+index"  :prop="'form.' + index + '.protocol'">
          <el-select class="top10"  v-model="port.protocol"  placeholder="填写protocol">
            <el-option label="TCP" value="TCP"/>
            <el-option label="UDP" value="UDP"/>
            <el-option label="SCTP" value="SCTP"/>
          </el-select>
        </el-form-item>
        <el-form-item class="flex1"  label="映射port"   :key="'port'+index"  :prop="index + '.port'"
                      :rules="{required: true, trigger:'blur', message: 'port必须填写' }">
          <el-input-number min="1" max="65500"    class="top10"  v-model="port.port"  placeholder="port必须填写"></el-input-number>
        </el-form-item>
        <el-form-item class="flex1"  label="容器port"   :key="'targetPort'+index"  :prop="index + '.targetPort'"
                      :rules="{required: true, trigger:'blur', message: 'targetPort必须填写' }">
          <el-input-number  min="1" max="65500"    class="top10"  v-model="port.targetPort"  placeholder="targetPort必须填写"></el-input-number>
        </el-form-item>
        <el-form-item class="flex1" v-if="state.type=='NodePort'"  label="nodePort"   :key="'nodePort'+index"  :prop="index + '.nodePort'">
          <el-input-number   min="30000" max="32767"  class="top10"  v-model="port.nodePort" ></el-input-number>
        </el-form-item>
        <el-form-item  >
          <el-button  type="danger" size="small" @click="removeItem(index)" >删除此项</el-button>
        </el-form-item>
      </div>
    </el-form>
  </div>
  <div v-if="render">
    <h2>端口列表</h2>
    <el-table  border :data="state.form" class="smallTable"  size="small" empty-text="无数据">
      <el-table-column prop="name" label="端口名称"  />
      <el-table-column prop="protocol" label="协议" />
      <el-table-column prop="port" label="映射port" />
      <el-table-column prop="targetPort" label="容器port" />
      <el-table-column prop="nodePort"  v-if="state.type=='nodePort'" label="nodePort" />
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import {computed, inject, reactive, ref, watch} from "vue";
import {arrToMap, CheckData, hasValues, IsDirty, MapToArr} from "../../helper/helper";
import {NameToArr} from "../../helper/rules";
const props = defineProps(["type",])
let formRef=ref(null)
let state=reactive({
   form:[],
   type:props.type,
    md5:""
})
watch(()=>props.type,()=>{
  state.type=props.type
})
const render = inject("render")


async function Check(){
  return  await formRef.value.validate()
}

function setData(data) {
  if (data){
    state.form=data
  }

}
defineExpose({ setData,Check })


function addItem(){
  state.form.push({name:"",protocol:"TCP",port:80,targetPort:80,nodePort:80})
}
function removeItem(index){
  state.form.splice(index,1)
}
const emit = defineEmits(['input'])
function commit(){
  emit("input",state.form)
}
watch(()=>state.form,()=>{
  if(!IsDirty(state)) {
    commit()
  }
},{deep:true,flush:"post"})
</script>

<style scoped>

</style>