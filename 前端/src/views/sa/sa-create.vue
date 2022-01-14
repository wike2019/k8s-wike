<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务账号.创建"></breadcrumb>
     <el-form :inline="true"  :model="form" ref="formRef"  >
     <el-button size="small" type="warning" class="btnList" @click="doTo('sa-list')" >进入列表</el-button >
     <el-tabs v-model="mode" @tab-click="Update">
       <el-tab-pane label="可视化展示" name="json">
             <metadataInfo @input="getData($event,'metadata',form)" :value="form.metadata" ref="metadataInfoRef"  :nameRequired="true"></metadataInfo>
             <secrets  :key="'secrets'" @input="getData($event,'secrets',form)" secretsType="secrets"  :value="form.secrets" :namespace="form.metadata?.namespace"    ref="secretsRef" ></secrets>
             <imagePullSecrets :key="'imagePullSecrets'" @input="getData($event,'imagePullSecrets',form)"  secretsType="imagePullSecrets"  :namespace="form.metadata?.namespace"  :value="form.imagePullSecrets" ref="imagePullSecretsRef"></imagePullSecrets>
       </el-tab-pane>
       <el-tab-pane label="YAML展示" name="yaml">
         <yaml ref="yamlRef"  @input="yamlChange" />
       </el-tab-pane>
     </el-tabs>
     <el-divider></el-divider>
     <div  class="submit-box">
       <el-button  type="info" @click="post()">提交</el-button>
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
import {requireRules, checkYaml} from "@/helper/rules.ts"
import {CheckData, createTip, getData, IsDirty, IsReady, showErr} from "@/helper/helper.ts"
import breadcrumb from "@/components/common/breadcrumb.vue";
import metadataInfo from "@/components/metadataInfo/metadataInfo.vue";
import secrets from "./components/secrets.vue";
import imagePullSecrets from "./components/imagePullSecrets.vue";
export default defineComponent({
  name: 'sa-create',
  components: {MainLayout,yaml,metadataInfo,breadcrumb,secrets,imagePullSecrets},
  setup: function () {
    let state = reactive({
      mode: "json",
      form: {
        apiVersion: 'v1',
        Kind: 'ServiceAccount',
        metadata: {
          name: "",
          namespace: "default",
          labels: {},
          annotations: {},
        },
        secrets: [],
        imagePullSecrets: []
      },
      md5: "",
      apiVersion: 'v1',
      Kind: 'ServiceAccount',
    })
    let yamlRef = ref(null)
    let metadataInfoRef = ref(null)
    let secretsRef = ref(null)
    let imagePullSecretsRef = ref(null)
    let checkList=[imagePullSecretsRef,secretsRef,metadataInfoRef]
    function Update() {
      nextTick(() => {
        yamlRef.value.Update()
      })
    }

    function yamlChange(data) {
      if (data) {
        let msg = checkYaml(data, state)
        if (msg) {
          showErr(msg, back)
          return
        }
        state.form = data
      }
    }

    function back() {
      state.mode = "yaml"
      Update()
    }

    watch(() => state.form, () => {
      if (IsDirty(state, 4)) {
        return
      }
      yamlRef.value.setData(state.form)
    }, {deep: true, flush: "post"})


    function post() {
      createTip()
          .then(async () => {
              await CheckData(checkList,1)
              IsReady(SACreate,state)
          })
    }

    onMounted(function () {
      yamlRef.value.setData(state.form)
    })
    return {
      ...toRefs(state),
      doTo,
      Update,
      secretsRef,
      imagePullSecretsRef,
      yamlRef,
      metadataInfoRef,
      yamlChange,
      getData,
      requireRules,
      post,
    }
  }
})
</script>
<style >

</style>