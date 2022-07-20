<template>
    <el-row :gutter="20" class="full-height">
        <el-col :span="6">
            <ApiTree :tree="w?.group || []" @change="onchange" @save-group="saveGroup" :active-api="active" :workspace-id="activeWsId"></ApiTree>
        </el-col>
        <el-col :span="18">
            <ApiDetail :key="key" :api="api" @save-api="saveApi"></ApiDetail>
        </el-col>
    </el-row>
</template>
<script lang="ts" setup>
import ApiDetail from './ApiDetail.vue'
import ApiTree from './ApiTree.vue'
import { app } from '~/go/models'
import { LoadWorkspace, UpdateApi } from '~/go/app/App.js'
import { ElMessage } from 'element-plus'
import type { API } from './model'
import { useAppStore } from '@/store/app'

const store = useAppStore()
const activeWsId = computed(() => store.activeWsId)

watch(activeWsId, () => {
    load()
})

const w = ref<app.Workspace>()

function load() {
    LoadWorkspace(activeWsId.value).then(res => {
        w.value = res as app.Workspace
    }).catch(err => {
        ElMessage.error(err)
    })
}

load()

const api = ref<API>()
const active = ref<app.ApiList>()
const key = ref(0)

function apiList2API(a: app.ApiList): API {
    const detail: Record<string, any> = JSON.parse(a.options)
    return {
        id: a.id || 0,
        gid: a.gid || 0,
        method: a.method,
        name: a.name,
        url: detail?.url || '',
        paramsSource: detail?.paramsSource || `## example
# log_date : 20200201
`,
        headerSource: detail?.headerSource || `## example
# Content-Type : application/json
`,
        bodySource: detail?.bodySource || `// {
//     "k": v
// }
`
    }
}

function API2ApiList(data: API): app.ApiList {
    const _opt = {
        url: data.url,
        paramsSource: data.paramsSource,
        headerSource: data.headerSource,
        bodySource: data.bodySource,
        script: data.script,
    }

    return new app.ApiList({
        id: data.id,
        gid: data.gid,
        name: data.name,
        method: data.method,
        options: JSON.stringify(_opt)
    })
}

const onchange = (data: app.ApiList) => {
    active.value = data
    api.value = apiList2API(data)
    key.value++
}

const saveApi = (data: API) => {
    const _data = API2ApiList(data)
    UpdateApi(_data).then(res => {
        res ? ElMessage.success('saved') : ElMessage.error('saved error')
        active.value = _data
    }).catch((e) => {
        ElMessage.error(e as string)
    })

}

const saveGroup = () => {
    load()
}
</script>
<style lang="scss" scope>
.full-height {
    height: 100%;
}
.ws {
    font-size: 13px;
    padding: 10px 5px;
}
</style>