<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务账号.详情"></breadcrumb>
     <metadataInfo tipTitle="ServiceAccount"  :value="form.metadata" ref="metadataInfoRef"  ></metadataInfo>
     <el-divider></el-divider>
     <secrets tipTitle="secrets"  :key="'secrets'" :namespace="form.metadata?.namespace"   secretsType="secrets【ServiceAccount关联密文】"  :value="form.secrets"     ref="secretsRef" ></secrets>
     <secrets tipTitle="imagePullSecrets" :namespace="form.metadata?.namespace"   :key="'imagePullSecrets'" secretsType="imagePullSecrets【拉取私有镜像凭证密文】"   :value="form.imagePullSecrets" ref="imagePullSecretsRef"></secrets>
     <eventList kind="ServiceAccount" :uid="form.metadata?.uid" :namespace="form.metadata?.namespace" :name="form.metadata?.name"></eventList>

   </main-layout>
</template>

<script lang="ts">
import {defineComponent, provide, reactive, toRefs} from 'vue'
import MainLayout from "../../layout/main.vue";
import {doTo} from "../../router";
import {getSaItem,SaDel} from "../../api/token/sa/sa";
import {secretDetail} from "../../api/token/secret/secret";
import {useRoute} from "vue-router";
import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
import secrets from "./components/secrets.vue";
import eventList from "../../components/event/eventList.vue";
import breadcrumb from "../../components/common/breadcrumb.vue";

export default defineComponent({
  name: 'sa-detail',
  components: {MainLayout,metadataInfo,secrets,eventList,breadcrumb},
  setup(){
    let state=reactive({
      form:{},
      name:"",
      namespace:""
    })
    provide("render",true)
    provide("edit",'sa-update')
    provide("delFn",SaDel)
    const route = useRoute()
    state.name=route.query.name
    state.namespace=route.query.namespace
    async function fetchData(){
      try {
       let tData=await getSaItem(state.namespace,state.name)
        state.form=tData.data.data
      }catch (e){
        console.log(e)
      }
    }
    fetchData()

    return {...toRefs(state),doTo}
  }
})
</script>
<style >

</style>