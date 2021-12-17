<template>
    <main class="flex">
      <Aside />
      <div class="main_detail_content">
        <slot></slot>
      </div>
    </main>
</template>

<script lang="ts">
import { ref, defineComponent } from 'vue'
import Aside from "../components/common/aside.vue";
export default defineComponent({
  name: 'main-layout',
  components: {Aside},
})
</script>