<template>
  <el-table
    :data="showListData()"
    stripe
    style="width: 100%"
    :height="400"
  >
    <el-table-column
      v-for="(item, index) in props.headers"
      :key="index + '-' + item.field"
      :label="item.label"
      :width="item.width"
    >
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <span>{{ scope.row[item.field] }}</span>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script setup lang="ts">
import { ref, defineProps, PropType, computed } from 'vue'
import ReconnectingWebSocket from 'reconnecting-websocket'
type List = Array<Record<string, any>>

const props = defineProps({
  headers: {
    type: Array as PropType<List>,
    default: () => []
  },
  ws: {
    type: String,
    default: ''
  },
  dataType: {
    type: String,
    default: 'push'
  },
  messageHandle: {
    type: Function,
    default: (e: any) => e
  },
  filterWords: {
    type: String,
    default: ''
  },
  max: {
    type: Number,
    default: 200
  }
})

const options = {
  connectionTimeout: 1000,
  maxRetries: 10
}

const showListData = () => {
  console.log(111, props.filterWords)
  if (props.filterWords === '') {
    return list.value
  }
  return list.value.filter(item => {
    return JSON.stringify(item).toLowerCase().indexOf(props.filterWords.toLowerCase()) > -1
  })
}

const list = ref<List>([])
const ws = new ReconnectingWebSocket(props.ws, [], options)
ws.onopen = () => {}
ws.onmessage = (msg) => {
  const m = props.messageHandle(JSON.parse(msg.data))
  switch (props.dataType) {
    case 'push':
      if (list.value.length >= 200) {
        list.value = list.value.splice(0, 1)
      }
      list.value.push(m)
      break
    case 'overwrite':
      list.value = m
      break
  }
}
ws.onerror = (err) => {
}
</script>
