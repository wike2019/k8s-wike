<template>
  <div class="yaml-editor">
    <textarea ref="yamlRef"></textarea>
  </div>
  <div class="footer-common">
    <el-button type="primary" @click="download">下载Yaml文件</el-button>
  </div>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, nextTick, reactive} from 'vue'
import {doTo} from '../../router';
import jsyaml  from "js-yaml"
import {core} from "../../core/core.ts"
import md5 from 'js-md5';
import {ElMessage} from "element-plus";
import {autoComplete} from "../../api/other";
export default defineComponent({
  name: 'yaml',
  props: {
    msg: {
      type: String,
      required: true
    }
  },
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
          console.log(prefix)
          let data=await autoComplete()
          console.log(data.data)
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
        nextTick(()=>{
          try {
              if(state.md5!=md5(JSON.stringify(json||""))){
                emit("input",json)
              }
              let value=jsyaml.dump(json)
              state.md5=md5(JSON.stringify(json||""))
              editor.setValue(value);

          }catch (e) {
            console.log(e)
          }
        })


    }
    function Update(){
      let session = editor.getSession();
      let count = session.getLength();
      editor.gotoLine(count, session.getLine(count - 1).length);
      editor.focus()
    }
    function check(){

    }
    function download(){
      core.fileDownLoad(editor.getValue(),"value.yaml")
    }
    function send(){
      try {
        let json=jsyaml.load(state.data)
        emit("input",json)
      }catch (e) {
        emit("err","")
      }
    }

    return {yamlRef,setData,download,Update,send}
  }
})
</script>

<style scoped>

</style>
