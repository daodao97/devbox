<template>
    <LocalPlugin v-if="name" :src="`/plugin/${name}/index.html`" ></LocalPlugin>
</template>

<script lang="ts" setup>
import LocalPlugin from '@/components/LocalPlugin.vue';
import { useRoute, useRouter } from 'vue-router';
const route = useRoute()
const name = route.query.name
if (!name) {
    useRouter().go(-1)
}
</script>