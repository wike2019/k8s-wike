  <template>
  <div  v-if="!render">
      <div class="inline-title-no" v-if="!label">
        <span>Label数据设置</span>         <el-button type="primary" size="small"  class="leftBtn" @click="addItem" >添加标签</el-button>
      </div>
      <div class="inline-title-no" v-else>
        <span>{{label}}数据设置</span>         <el-button type="primary" size="small"  class="leftBtn" @click="addItem" >添加{{label}}</el-button>
      </div>
      <el-form :inline="true"  :model="state.form" ref="formRef"  >
        <div v-for="(Mate,index) in state.form.data"  class="flex">
          <template v-if="Mate.name">
            <div   class="el-form-item flex1" :class="{'is-required':!!Mate.required}">
              <label class="el-form-item__label" >
                <span> {{Mate.name}}</span>
              </label>
            </div>
          </template>
          <template v-else >
            <el-form-item class="flex1"  label="key"   :key="'key'+index"  :prop="'data.' + index + '.key'"
                          :rules="{required: true, trigger: 'blur', message: 'key必须填写' }">
              <el-input v-model="Mate.key"  class="top10"  placeholder="填写key"></el-input>
            </el-form-item>
          </template>
            <el-form-item class="flex1"  :label="Mate.name?'':'value'"   :key="'value'+index"  :prop="'data.' + index + '.value'"
                          :rules="Mate.required?{required: true, trigger: 'blur', message: 'value必须填写' }:{}">
              <el-input     class="top10"  v-model="Mate.value"  placeholder="填写value"></el-input>
            </el-form-item>
            <el-form-item label="是否必填"   :key="'required'+index"  :prop="'data.' + index + '.required'" >
              <el-checkbox v-model="Mate.required"></el-checkbox>
            </el-form-item>
            <el-form-item  >
              <el-button  type="danger" size="small" @click="removeItem(index)" >删除此项</el-button>
            </el-form-item>
        </div>
      </el-form>
  </div>
  <div v-if="render">
    <div class="inline-title-no-back" v-if="!label">
      <span>lebal数据</span>
    </div>
    <div v-else>
        <span>{{label}}数据</span>
    </div>
    <el-table  border :data="state.form.data" class="smallTable"  size="small" empty-text="无数据">
      <el-table-column prop="key" label="key" width="300" />
      <el-table-column prop="value" label="值" />
    </el-table>
  </div>
</template>

<script lang="ts" setup>
import {defineComponent, inject, reactive, ref, toRefs, watch} from 'vue'
import {arrToMap, MapToArr,IsDirty} from "../../helper/helper";
import {hasValues} from "../../helper/helper";
const render = inject("render")
const props = defineProps(['label',"value"])
const emit = defineEmits(['input'])
function commit(){
  emit("input",arrToMap(state.form.data))
}
let state=reactive({
  form:{
    data:[]
  },
  md5:"",
  index:[],
  label:props.label,
})
let formRef=ref(null)

watch(()=>state.form,()=>{
  if(!IsDirty(state)) {
    commit()
  }
},{deep:true,flush:"post"})


function addItem(){
  state.form.data.push({key:"",value:"",name:"",required:true})
}
function removeItem(index){
  state.form.data.splice(index,1)
}
async  function Check(){
  return  await formRef.value.validate()
}
function setData(value) {
  if(!hasValues(value)){
    state.form.data=[]
  }
  state.form.data=MapToArr(value)
}
defineExpose({ setData,Check })
</script>

<style>
.top10{margin-top:-15px}
</style>
