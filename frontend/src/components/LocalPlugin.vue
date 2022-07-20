<template>
    <template v-if="isLoading">loading</template>
    <template v-else>
        <iframe :src="host + props.src" style="width:100%; height:100%"></iframe>
    </template>
</template>
<script lang="ts" setup>
import { Environment } from '~/runtime/runtime.js'
import { useAsyncState } from '@vueuse/core'
import { PropType } from 'vue';

const props = defineProps({
    src: {
        type: String as PropType<string>,
        default: ''
    }
})

const { state, isReady, isLoading } = useAsyncState(
    Environment().then(res => res),
    {
        // @ts-ignore
        buildtype: "",
        platform: "",
        arch: ""
    },
)

// @ts-ignore
const env = computed(() => state.value.buildtype)
const host = computed(() => env.value === 'dev' ? 'http://localhost:34115' : 'wails://wails')
</script>