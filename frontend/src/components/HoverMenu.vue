<template>
  <div class="my-card">
    <div class="card-body">
      <slot />
    </div>
    <div
      v-if="props.menus.length > 0"
      class="card-action"
    >
      <div class="card-action-text">
        <el-icon><MoreFilled /></el-icon>
      </div>
      <div class="card-action-menus hide">
        <div
          v-for="(item, index) in props.menus"
          :key="index"
          class="contextmenu_item"
          @click="() => item.onclick()"
        >{{ item.text }}</div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { MoreFilled } from '@element-plus/icons-vue'
import { PropType } from 'vue'

interface menu {
  text: string,
  onclick: Function
}
const props = defineProps({
  menus: {
    type: Array as PropType<Array<menu>>,
    default: () => []
  }
})
</script>
<style lang="scss" scoped>
.my-card {
  position: relative;
  background: var(--el-bg-color);
  border-radius: var(--el-card-border-radius);
  padding: 20px;
  .card-body {
    height: 100%;
  }
}

.card-action {
  display: block;
  position: absolute;
  right: 16px;
  bottom: 10px;
  cursor:pointer;
  overflow: visible;
}

.card-action-text:hover + .hide {
  display: block;
}

.card-action-menus {
  margin: auto;
  position: absolute;
  color: var(--el-text-color-regular);
  background: var(--el-bg-color);
  border-radius: var(--el-card-border-radius);
  border: 1px solid rgba(0, 0, 0, 0.15);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.175);
  white-space: nowrap;
  z-index: 9999;
}

.card-action-menus:hover {
  display: block;
}

.contextmenu_item {
  padding: 5px;
  display: block;
  line-height: 20px;
  text-align: center;
}
.contextmenu_item:not(:last-child) {
  border-bottom: 1px solid rgba(0, 0, 0, 0.1);
}

.contextmenu_item:hover {
  cursor: pointer;
  background: #66b1ff;
  border-color: #66b1ff;
  color: #fff;
}

.hide {
  display: none;
}
</style>
