<template>
  <div shadow="never" class="my-card" data-wails-no-drag>
    <slot />
  </div>
</template>

<style lang="scss" scoped>
.my-card {
    height: 100%;
    width: 100%;
}
.list {
    margin: 10px 0;
}
.upload-btn {
    width: 100%;
}
</style>
