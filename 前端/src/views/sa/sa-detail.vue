<template>
   <main-layout class="fix-black">
     <breadcrumb title="服务账号.详情"></breadcrumb>
     <metadataInfo tipTitle="ServiceAccount"   ref="metadataInfoRef"  ></metadataInfo>
     <el-divider></el-divider>
     <secrets tipTitle="secrets"  :key="'secrets'" :namespace="state.form.metadata?.namespace"   secretsType="secrets【ServiceAccount关联密文】"  :value="state.form.secrets"     ref="secretsRef" ></secrets>
     <secrets tipTitle="imagePullSecrets" :namespace="state.form.metadata?.namespace"   :key="'imagePullSecrets'" secretsType="imagePullSecrets【拉取私有镜像凭证密文】"   :value="state.form.imagePullSecrets" ref="imagePullSecretsRef"></secrets>
     <eventList kind="ServiceAccount" :uid="state.form.metadata?.uid" :namespace="state.form.metadata?.namespace" :name="state.form.metadata?.name"></eventList>
     <el-button class="mtb20" type="primary" @click="()=>doTo('sa-update',{namespace:state.namespace,name:state.name})" >
       编辑
     </el-button>
   </main-layout>
</template>

<script lang="ts" setup>
import {defineComponent, provide, reactive, ref, toRefs} from 'vue'
    import MainLayout from "../../layout/main.vue";
    import {doTo} from "../../router";
    import {getSaItem,SaDel} from "../../api/token/sa/sa";
    import {secretDetail} from "../../api/token/secret/secret";
    import {useRoute} from "vue-router";
    import metadataInfo from "../../components/metadataInfo/metadataInfo.vue";
    import secrets from "./components/secrets.vue";
    import eventList from "../../components/event/eventList.vue";
    import breadcrumb from "../../components/common/breadcrumb.vue";
    let state=reactive({
      form:{
        metadata:{},
        imagePullSecrets:{},
        secrets:{}
      },
      name:"",
      namespace:""
    })
    provide("render",true)
    let metadataInfoRef=ref(null)
    let imagePullSecretsRef=ref(null)
    let secretsRef=ref(null)
    const route = useRoute()
    state.name=route.query.name.toString()
    state.namespace=route.query.namespace.toString()
    async function fetchData(){
      try {
        let tData=await getSaItem(state.namespace,state.name)
        state.form=tData.data.data
        metadataInfoRef.value.setData( state.form.metadata)
        imagePullSecretsRef.value.setData(state.form.imagePullSecrets)
        secretsRef.value.setData(state.form.secrets)
      }catch (e){
        console.log(e)
      }
    }
fetchData()
</script>
<style >

</style>