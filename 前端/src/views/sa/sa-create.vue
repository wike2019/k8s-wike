<template>
   <main-layout>
     <breadcrumb title="服务账号.创建"></breadcrumb>
     <el-form :inline="true"  :model="form" ref="formRef"  >
     <el-button size="small" type="warning" class="btnList" @click="doTo('sa-list')" >进入列表</el-button >
     <el-tabs v-model="mode" @tab-click="common.Update">
       <el-tab-pane label="可视化展示" name="json">
             <metadataInfo tipTitle="ServiceAccount"   @input="getData($event,'metadata',form)" :value="form.metadata" ref="metadataInfoRef"  :nameRequired="true"></metadataInfo>
             <el-tabs   v-model="spec" class="inline-box">
               <el-tab-pane label="ServiceAccount关联密文" name="secrets">
                 <secrets tipTitle="secrets"  :key="'secrets'" @input="getData($event,'secrets',form,true)" secretsType="secrets【ServiceAccount关联密文】"  :value="form.secrets" :namespace="form.metadata?.namespace"    ref="secretsRef" ></secrets>
               </el-tab-pane>
               <el-tab-pane label="拉取私有镜像凭证密文" name="imagePullSecrets">
                 <secrets tipTitle="imagePullSecrets" :key="'imagePullSecrets'" @input="getData($event,'imagePullSecrets',form,true)"  secretsType="imagePullSecrets【拉取私有镜像凭证密文】"  :namespace="form.metadata?.namespace"  :value="form.imagePullSecrets" ref="imagePullSecretsRef"></secrets>
               </el-tab-pane>
             </el-tabs>
       </el-tab-pane>
       <el-tab-pane :label="'YAML展示(限定资源'+Kind+')'" name="yaml">
         <yaml ref="yamlRef" @err="common.showErr"  @input="common.yamlChange" />
       </el-tab-pane>
     </el-tabs>
     <el-divider></el-divider>
     <div v-show="hasErr">
         <div class="inline-title">
           <span>错误信息提示</span>
         </div>
         <div id="errorReport">

         </div>
     </div>
     <div  class="submit-box">
       <el-button  type="info" @click="post(false)">校验数据提交(推荐)</el-button>
       <el-button  type="danger"  @click="post(true)" >不校验数据提交</el-button >
     </div>
     </el-form>
   </main-layout>
</template>

<script lang="ts">
import {defineComponent, nextTick, onMounted, reactive, ref, toRefs, watch} from 'vue'

import MainLayout from "@/layout/main.vue";
import {doTo} from "@/router";
import { SACreate} from "@/api/token/sa/sa";
import yaml from "@/components/yaml/yaml.vue";
import {requireRules} from "@/helper/rules.ts"
import {CheckData, createTip, getData, IsDirty, IsReady} from "@/helper/helper.ts"
import breadcrumb from "@/components/common/breadcrumb.vue";
import metadataInfo from "@/components/metadataInfo/metadataInfo.vue";
import secrets from "./components/secrets.vue";
import {initFn} from "../../helper/helper";
export default defineComponent({
  name: 'sa-create',
  components: {MainLayout,yaml,metadataInfo,breadcrumb,secrets},
  setup: function () {
    let state = reactive({
      mode: "json",
      spec:'secrets',
      form: {
        apiVersion: 'v1',
        Kind: 'ServiceAccount',
        metadata: {
          name: "",
          namespace: "default",
        },
      },
      md5: "",
      apiVersion: 'v1',
      Kind: 'ServiceAccount',
      hasErr:false,
      notModifyList:[
        {
          key:"apiVersion",
          msg:"apiVersion不允许修改",
          modify:"apiVersion"
        },
        {
          key:"Kind",
          msg:"Kind不允许修改",
          modify:"Kind"
        },
      ]
    })
    let yamlRef = ref(null)
    let metadataInfoRef = ref(null)
    let secretsRef = ref(null)
    let imagePullSecretsRef = ref(null)
    let checkList=[imagePullSecretsRef,secretsRef,metadataInfoRef]


    watch(() => state.form, () => {
      if (IsDirty(state,null,true)) {
        return
      }
      yamlRef.value.setData(state.form)
    }, {deep: true, flush: "post"})


    function post(flag) {
      createTip()
          .then(async () => {
              await CheckData(checkList,1)
              IsReady(SACreate,state,"ServiceAccount资源创建成功",flag)
          })
    }
    let common=initFn(yamlRef,nextTick,state)
    onMounted(function () {
      yamlRef.value.setData(state.form)
    })
    return {
      ...toRefs(state),
      doTo,
      secretsRef,
      imagePullSecretsRef,
      yamlRef,
      metadataInfoRef,
      common,
      getData,
      requireRules,
      post,
    }
  }
})
</script>
<style >

</style>