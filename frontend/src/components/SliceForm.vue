<template>
  <el-row>
    <el-col>
      <div
        v-for="(item, index) in localValue"
        :key="index"
        class="items"
      >
        <el-input
          v-model="localValue[index]"
          @change="onchange"
        >
          <template #append>
            <el-icon
              style="color: red"
              @click="removeOne(index)"
            ><Remove /></el-icon>
          </template>
        </el-input>
      </div>
    </el-col>
    <el-col>
      <el-icon
        style="color: green"
        @click="addOne"
      ><CirclePlus /></el-icon>
    </el-col>
  </el-row>
</template>
<script lang="ts" setup>
import { PropType, ref } from 'vue'
import { CirclePlus, Remove } from '@element-plus/icons-vue'

const props = defineProps({
  modelValue: {
    type: Array as PropType<Array<string>>,
    default: () => []
  },
  clearable: {
    type: Boolean,
    default: false
  }
})

const emits = defineEmits(['update:modelValue'])

const localValue = ref(props.modelValue)

const onchange = () => {
  const value = localValue.value.filter(e => e)
  console.log(value)
  emits('update:modelValue', value)
}

const addOne = () => {
  localValue.value.push('')
}
const removeOne = (index: number) => {
  localValue.value.splice(index, 1)
}

</script>

<style lang="scss" scoped>
.items {
  margin: 10px 0;
  &:last-child {
    margin-bottom: 0;
  }
}

</style>
