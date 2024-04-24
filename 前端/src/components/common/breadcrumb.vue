<template>
  <nav class="nav-bar">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="item in list">{{item}}</el-breadcrumb-item>
    </el-breadcrumb>
  </nav>
</template>
<script lang="ts" setup>
const props = defineProps(['title'])
let list=props.title.split(".")
</script>