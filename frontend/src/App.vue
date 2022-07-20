<template>
  <el-config-provider>
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </el-config-provider>
</template>

<script setup lang="ts">
onMounted(() => {
  document.body.removeChild(document.getElementById('preloader') as Node)
})
provide('vsPath', import.meta.env.PROD ? location.pathname + 'assets/monaco-editor/vs' : 'node_modules/monaco-editor/min/vs')
</script>

<style lang="scss">
html,
body,
#app,
.el-container {
  /*设置内部填充为0，几个布局元素之间没有间距*/
  padding: 0px;
  /*外部间距也是如此设置*/
  margin: 0px;
  /*统一设置高度为100%*/
  height: 100%;
}
</style>
