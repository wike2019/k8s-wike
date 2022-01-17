<template>
  <el-alert title="支持k8s关键字提示,非法字段不会影响提交结果但是会验证yaml语法,非法字段系统会自动过滤,提交前会验证数据是否合法。具体错误信息请看可视化界面" type="error"></el-alert>
  <div class="footer-common">
    <el-button size="small" @click="upload">导入文件数据</el-button>
    <el-button type="primary" size="small" @click="download">下载Yaml文件</el-button>
    <el-button type="danger" size="small" @click="doTo('yaml-all')">使用多资源编辑器 (运维专用)</el-button>
  </div>
  <div class="yaml-editor">
    <textarea ref="yamlRef"></textarea>
  </div>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, nextTick, reactive, onUnmounted} from 'vue'
import {doTo} from '../../router';
import jsyaml  from "js-yaml"
import {core} from "../../core/core.ts"
import md5 from 'js-md5';
import {ElMessage} from "element-plus";
import {getAutoComplete} from "../../api/token/common/common";
import {initFn} from "../../helper/helper";
export default defineComponent({
  name: 'yaml',
  emits:["input","err"],
  setup: (props,{emit}) => {
    let yamlRef=ref(null)
    let state=reactive({
      md5:"",
      data:""
    })
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
      editor.on("blur",function (){
        send()
      })


    })


    function setData(json){
        if(!json){
          json=""
        }
        nextTick(()=>{
          try {

              let value=jsyaml.dump(json)
              state.md5=md5(JSON.stringify(json))
              editor.setValue(value);
              if(state.md5!=md5(JSON.stringify(json))){
                emit("input",json)
              }
          }catch (e) {
            emit("err","数据格式不正确")
          }
        })
    }
    function Update(){
      editor.gotoLine(0, 0);
      editor.focus()
    }
    function download(){
      core.fileDownLoad(editor.getValue(),"text.yaml")
    }
    function send(){
      try {
        let json=jsyaml.load(state.data)
        emit("input",json)
      }catch (e) {
        console.error(e)
        emit("err","数据格式不正确")
      }
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
    return {yamlRef,setData,download,Update,send,upload,doTo,blur}
  }
})
</script>

<style scoped>

</style>
