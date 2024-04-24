<template>
  <el-alert title="支持k8s关键字提示,非法字段不会影响提交结果但是会验证yaml语法,非法字段系统会自动过滤,提交前会验证数据是否合法。具体错误信息请看可视化界面" type="error"></el-alert>
  <div class="footer-common">
    <el-button size="small" @click="upload" v-if="props.all">导入文件数据</el-button>
    <el-button type="primary" size="small" @click="download">下载Yaml文件</el-button>
    <el-button type="danger" size="small" @click="doTo('yaml-all')">使用多资源编辑器 (运维专用)</el-button>
  </div>
  <div class="yaml-editor">
    <textarea ref="yamlRef"></textarea>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, nextTick, reactive} from 'vue'
import {doTo} from '../../router';
import jsyaml  from "js-yaml"
import {core} from "../../core/core.ts"
import {getAutoComplete} from "../../api/token/common/common";
import {ElMessage} from "element-plus";

let yamlRef=ref(null)
let state=reactive({
  data:""
})
const props = defineProps(['all'])
let editor
onMounted(function () {

  editor=ace.edit(yamlRef.value,{minLines: 0})
  editor.setTheme("ace/theme/nord_dark");
  editor.session.setMode("ace/mode/yaml");
  editor.setAutoScrollEditorIntoView(true);
  editor.setOption("maxLines", 35);
  editor.setOption("minLines", 35);
  editor.setOptions({
    enableBasicAutocompletion: false,
    enableSnippets: true,
    enableLiveAutocompletion: true,
  });
  editor.completers.push({
    getCompletions: async  function(editor, session, pos, prefix, callback) {
      let data=await getAutoComplete()
      if (prefix.length === 0) {
        return callback(null, []);
      } else {
        return callback(null, data.data.data);
      }
    }
  })

  editor.resize(true)
  editor.on('change', data => {
    state.data=editor.getValue()
  });
})



function setData(json){
  if(!json){
    json=""
  }
  nextTick(()=>{
    try {
      let value=jsyaml.dump(json)
      editor.setValue(value);
    }catch (e) {
      ElMessage.error(e)
    }
  })
}
function Update(){
  editor.gotoLine(0, 0);
  editor.focus()
}
function download(){
  core.fileDownLoad(editor.getValue(),"data.yaml")
}

function upload(){
  let fileDom= document.createElement("input")
  fileDom.type="file"
  fileDom.onchange=function (e) {
    if(e.target.files.length>0){
      e.target.result
    }
    const reader=new FileReader()
    reader.readAsText(e.target.files[0],'UTF-8')
    reader.onload=(e)=>{
      editor.setValue(e.target.result);
    }
  }
  fileDom.click();
}
defineExpose({ setData ,Update})
</script>

<style scoped>

</style>
