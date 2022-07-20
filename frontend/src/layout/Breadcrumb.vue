<template>
    <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: item.path }" v-for="item in route.matched">
         <Render v-if="item.path === route.path" :content="route?.query.title ?? (item?.meta?.icon ?? item.name)"></Render>
         <Render v-else :content="item?.meta?.icon ?? item.name"></Render>
        </el-breadcrumb-item>
    </el-breadcrumb>
</template>
<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { Render } from '@/components';
const route = computed(() => useRoute()) 
</script>