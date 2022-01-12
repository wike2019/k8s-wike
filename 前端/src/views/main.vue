<template>
 <main-layout>
   <div>测试</div>
 </main-layout>
</template>

<script lang="ts">
import {defineComponent } from 'vue'
import MainLayout from "../layout/main.vue";
export default defineComponent({
  name: 'view-main',
  components: {MainLayout},
})
</script>

<style scoped>

</style>
