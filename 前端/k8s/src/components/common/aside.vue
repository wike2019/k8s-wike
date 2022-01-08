<template>
  <div class="aside">
    <el-menu>
      <el-menu-item index="1" @click="doTo('namespace')">
        <template #title><span>命名空间</span></template>
      </el-menu-item>
      <el-submenu index="2">
        <template #title>
          <i class="fa fa-briefcase" aria-hidden="true"></i>
          <span>服务账号</span>
        </template>
        <el-menu-item-group title="服务账号相关">
          <el-menu-item @click="doTo('sa-create')" index="2-1"><span>SA创建</span></el-menu-item>
          <el-menu-item @click="doTo('sa-list')"  index="2-2"><span>SA列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="3">
        <template #title>
          <i class="fa fa-briefcase" aria-hidden="true"></i>
          <span>密文</span>
        </template>
        <el-menu-item-group title="密文相关">
          <el-menu-item @click="doTo('secret-create')" index="3-1"><span>Secret创建</span></el-menu-item>
          <el-menu-item @click="doTo('secret-list')"  index="3-2"><span>Secret列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="4">
        <template #title>
          <i class="fa fa-briefcase" aria-hidden="true"></i>
          <span>配置</span>
        </template>
        <el-menu-item-group title="configmap相关">
          <el-menu-item @click="doTo('configmap-create')" index="4-1"><span>ConfigMap创建</span></el-menu-item>
          <el-menu-item @click="doTo('configmap-list')"  index="4-2"><span>ConfigMap列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-menu-item index="1" @click="doTo('pod')">
        <i class="fa fa-file-powerpoint-o" aria-hidden="true"></i>
        <template #title><span>POD</span></template>
      </el-menu-item>
      <el-menu-item index="2" @click="doTo('svc')">
        <i class="fa fa-server" aria-hidden="true"></i>
        <template #title><span>服务</span></template>
      </el-menu-item>
      <el-submenu index="4">
        <template #title>
          <i class="fa fa-briefcase" aria-hidden="true"></i>
          <span>工作负载</span>
        </template>
        <el-menu-item-group title="Deployment相关">
          <el-menu-item @click="doTo('deploy-create')" index="4-1"><i class="fa fa-wrench" aria-hidden="true"></i><span>创建</span></el-menu-item>
          <el-menu-item @click="doTo('deployment-list')"  index="4-2"><i class="fa fa-th-list" aria-hidden="true"></i><span>列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="5">
        <template #title>
          <i class="fa fa-question-circle" aria-hidden="true"></i>
          <span>nginx-ingress</span>
        </template>
        <el-menu-item-group title="Ingress相关">
          <el-menu-item index="5-1"  @click="doTo('ingress-create')"><i class="fa fa-wrench" aria-hidden="true"></i><span>创建</span></el-menu-item>
          <el-menu-item index="5-2"  @click="doTo('ingress-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="6">
        <template #title>
          <i class="fa fa-question-circle" aria-hidden="true"></i>
          <span>资源</span>
        </template>
        <el-menu-item-group title="secret相关">
          <el-menu-item index="6-1"  @click="doTo('secret-create')"><i class="fa fa-wrench" aria-hidden="true"></i><span>创建</span></el-menu-item>
          <el-menu-item index="6-2"  @click="doTo('secret-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>列表</span></el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="configmap相关">
          <el-menu-item index="6-3"  @click="doTo('configmap-create')"><i class="fa fa-wrench" aria-hidden="true"></i><span>创建</span></el-menu-item>
          <el-menu-item index="6-4"  @click="doTo('configmap-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="7">
        <template #title>
          <i class="fa fa-hand-o-right" aria-hidden="true"></i>
          <span>节点管理</span>
        </template>
        <el-menu-item-group title="节点相关">
          <el-menu-item index="7-1"  @click="doTo('node-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>列表</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
      <el-submenu index="8">
        <template #title>
          <i class="fa fa-hand-o-right" aria-hidden="true"></i>
          <span>权限</span>
        </template>
        <el-menu-item-group title="RBAC相关-Role">
          <el-menu-item index="8-6" @click="doTo('rbac-ua-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>用户列表</span></el-menu-item>
          <el-menu-item index="8-1" @click="doTo('rbac-role-create')"><i class="fa fa-wrench" aria-hidden="true"></i><span>角色创建</span></el-menu-item>
          <el-menu-item index="8-2" @click="doTo('rbac-role-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>角色列表</span></el-menu-item>
          <el-menu-item index="8-3" @click="doTo('rbac-roleBinding-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>角色绑定列表</span></el-menu-item>
          <el-menu-item index="8-4" @click="doTo('rbac-roleBinding-create')"><i class="fa fa-th-list" aria-hidden="true"></i><span>创建角色绑定</span></el-menu-item>
          <el-menu-item index="8-6" @click="doTo('rbac-sa-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>SA列表</span></el-menu-item>
          <el-menu-item index="8-7" @click="doTo('rbac-clusterRole-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>集群角色列表</span></el-menu-item>
          <el-menu-item index="8-8" @click="doTo('rbac-clusterRole-create')"><i class="fa fa-th-list" aria-hidden="true"></i><span>创建集群角色</span></el-menu-item>

          <el-menu-item index="8-9" @click="doTo('rbac-clusterRoleBinding-list')"><i class="fa fa-th-list" aria-hidden="true"></i><span>集群角色绑定列表</span></el-menu-item>

          <el-menu-item index="8-10" @click="doTo('rbac-clusterRoleBinding-create')"><i class="fa fa-th-list" aria-hidden="true"></i><span>创建集群角色绑定</span></el-menu-item>
        </el-menu-item-group>
      </el-submenu>
    </el-menu>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from 'vue'
import {doTo} from '../../router';
export default defineComponent({
  name: 'aside',
  props: {
    msg: {
      type: String,
      required: true
    }
  },
  setup: () => {
    const count = ref(0)
    return { count,doTo }
  }
})
</script>

<style scoped>

</style>
