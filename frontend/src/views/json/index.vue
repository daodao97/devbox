<template>
  <div data-wails-no-drag class="box">
    <div class="source">
      <el-input v-focus v-model="source" type="textarea" placeholder="past your data" @input="onchange" />
    </div>
    <div class="show">
      <div class="ctrl">
        <el-row :gutter="20" style="width: 100%">
          <el-col :span="8">
            <el-input v-model="filter" placeholder="node path example: a[0]b.c" />
          </el-col>
          <el-col :span="16" class="action">
            <el-icon>
              <CopyDocument />
            </el-icon>
          </el-col>
        </el-row>

      </div>
      <div class="result">
        <div v-if="err" v-html="err"></div>
        <JsonView v-if="json" :data="show" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { JsonView } from '@okiss/vbtf'
//@ts-ignore
import jsonlint from './jsonlint.js'
import { isObject, isArray, get } from 'lodash'
import { CopyDocument } from '@element-plus/icons-vue'

const source = ref('')
const filter = ref('')
const json = ref()
const err = ref()
const show = computed(() => {
  if (filter.value) {
    const val = get(json.value, filter.value)
    if (!isArray(val) && !isObject(val)) {
      return [val]
    }
    return val
  }

  return json.value
})
const onchange = (val: string) => {
  try {
    const value = jsonlint.parse(val)
    if (isObject(value) || isArray(value)) {
      json.value = value
    }
    err.value = ''
  } catch (e) {
    err.value = `${e}`
    json.value = false
  }
}

</script>

<style lang="scss">
.box {
  height: 100%;
  display: flex;

  .source {
    padding: 0 10px;
    height: 100%;
  }

  .show {
    width: 100%;
    height: 100%;
  }

  .ctrl {
    display: flex;
    margin-bottom: 10px;

    .action {
      display: flex;
      align-items: center;
    }
  }

  .result {
    background-color: var(--el-bg-color);
    width: 100%;
    height: 84vh !important;
    overflow: scroll;
  }
}

.el-textarea__inner {
  height: 90vh !important;
}
</style>
