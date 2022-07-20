<template>
  <el-container class="app-main">
    <div data-wails-drag @dblclick="WindowToggleMaximise">
      <div class="header">
        <div class="header-left">
          <div class="app-ctrl"></div>
          <Breadcrumb></Breadcrumb>
        </div>
        <div class="header-right">
          <div class="user">
            <el-popover placement="bottom" trigger="hover" content="this is content, this is content, this is content">
              <template #reference>
                <el-icon>
                  <UserFilled></UserFilled>
                </el-icon>
              </template>
            </el-popover>
          </div>
        </div>
      </div>
    </div>

    <div class="my-main">
      <router-view />
    </div>
  </el-container>
</template>

<script setup lang="ts">
import { WindowToggleMaximise } from '~/runtime/runtime'
import Breadcrumb from './Breadcrumb.vue'
import { UserFilled } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'
const route = useRoute()
const padding = computed(() => {
  const isFull = !!route?.meta?.fullPage
  return isFull ? '0' : '20px'
})
console.log(111, padding)
</script>

<style lang="scss">
$app-background-color: var(--app-bg);
$app-sidebar-color: var(--app-sidebar-bg);
$app-sidebar-with: 68px;

html,
body,
#app {
  height: 100%;
  width: 100%;
}

.app-main {
  padding: 0px;
  margin: 0px;
  height: 100%;
  flex-direction: column !important;

  .header {
    width: 100%;
    height: 28px;
    align-items: center;
    display: flex;
    justify-content: space-between;
    background-color: $app-background-color;

    .header-left {
      display: flex;

      .app-ctrl {
        width: 70px;
        height: 100%;
      }
    }

    .header-right {
      display: flex;
      width: 100px;
      justify-content: flex-end;

      .user {
        padding: 0 20px;
      }
    }
  }

  .my-main {
    width: 100%;
    height: 100%;
    display: block;
    flex: 1;
    flex-basis: auto;
    overflow: auto;
    box-sizing: border-box;
    padding: v-bind(padding);
    // background-color: $app-background-color;
  }
}

.table {
  display: flex;
  flex-direction: column;

  .table-line {
    width: 100%;
    display: flex;
  }

  .table-cell {
    padding: 5px;
    display: flex;
  }
}
</style>
