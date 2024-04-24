<template>
    <main class="flex">
      <commonNav />
      <div class="main_detail_content">
        <slot></slot>
      </div>
    </main>
</template>

<script lang="ts" setup>
import commonNav from "@/components/common/commonNav.vue";

</script>