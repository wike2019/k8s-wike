<template>
    <main class="flex">
      <commonNav />
      <div class="main_detail_content">
        <slot></slot>
      </div>
    </main>
</template>

<script lang="ts">
import { ref, defineComponent } from 'vue'
import commonNav from "@/components/common/commonNav.vue";
export default defineComponent({
  name: 'main-layout',
  components: {commonNav},
})
</script>