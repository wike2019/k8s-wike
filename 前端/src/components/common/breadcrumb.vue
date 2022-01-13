<template>
  <nav class="nav-bar">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item v-for="item in list">{{item}}</el-breadcrumb-item>
    </el-breadcrumb>
  </nav>
</template>
<script lang="ts">
import { computed, defineComponent } from 'vue'
//头部导航
export default defineComponent({
  name: 'breadcrumb',
  props: {
    title: {
      type: String,
      required: true
    }
  },
  setup: (props) => {
    let list=props.title.split(".")
    return {list}
  }
})
</script>

<style scoped>

</style>
