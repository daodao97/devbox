<template>
    <iframe ref="iframe" :src="src" style="width:100%;height:100%"></iframe>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router';
const route = useRoute()
const src = route.query.src as string
if (!src) {
    useRouter().go(-1)
}
const iframe = ref()

watchEffect(() => {
    if (iframe.value) {
        iframe.value.onload = iframe.value.onreadystatechange = (e) => {
            const w = iframe.value.contentWindow
            console.log(w.document.body)
        }
    }

})

</script>
<style>
::v-dep(.el-main) {
    padding: var(--el-main-padding);
}
</style>