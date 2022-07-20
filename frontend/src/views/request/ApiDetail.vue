<template>
    <div data-wails-no-drag class="box">
        <el-row :gutter="10" class="selection">
            <el-col :span="4">
                <el-select v-model="options.method" placeholder="Select">
                    <el-option v-for="item in methodOptions" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="20">
                <el-input v-model="options.name" placeholder="API name"></el-input>
            </el-col>
        </el-row>
        <el-row :gutter="10" class="selection">
            <el-col :span="20">
                <el-input v-model="options.url" placeholder="API path"></el-input>
            </el-col>
            <el-col :span="4">
                <el-row :gutter="10">
                    <el-col :span="12">
                        <el-button style="width:100%" type="primary" :loading="requesting" @click="send"> Send
                        </el-button>
                    </el-col>
                    <el-col :span="12">
                        <el-button style="width:100%" type="success" @click="save" :disabled="!props.api || props.api?.id === 0"> Save </el-button>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
        <el-tabs v-model="activeTab" type="border-card" class="selection my-tabs" @tab-click="changeTab">
            <el-tab-pane label="Param" name="Params">
                <MonacoEditor ref="params" v-model="options.paramsSource" class="editor" mod="yaml"
                    :options="editorOptions" />
            </el-tab-pane>
            <el-tab-pane label="Header" name="Header">
                <MonacoEditor ref="headers" v-model="options.headerSource" class="editor" mod="yaml" />
            </el-tab-pane>
            <el-tab-pane label="Body" name="Body">
                <VJson ref="body" v-model="options.bodySource" class="editor" />
            </el-tab-pane>
            <el-tab-pane label="Script" name="Script">
                <MonacoEditor ref="body" v-model="options.script" class="editor" mod="javascript" />
            </el-tab-pane>
        </el-tabs>
        <div class="response" v-if="err">{{ err }}</div>
        <div class="response" v-if="response" :key="bodyKey">
            <MonacoEditor v-if="response?.Body" :options="{
                readOnly: true
            }" :model-value="response?.Body" :disabled="true" :mod="bodyType" class="fullEditor" />
        </div>
    </div>

</template>

<script lang="ts" setup>
import type { TabsPaneContext } from 'element-plus'
// @ts-ignore
import { MonacoEditor, VJson } from '@okiss/vbtf'
import { Request } from '~/go/http/Http'
import { http } from '~/go/models'
import type { API } from './model'
import { PropType } from 'vue'
import type { app } from '~/go/models'
import { isString, merge, throttle } from 'lodash'

const emit = defineEmits(['saveApi'])

const methodOptions = [
    { value: "GET", label: "GET" },
    { value: "POST", label: "POST" },
    { value: "PUT", label: "PUT" },
    { value: "DELETE", label: "DELETE" },
]

const props = defineProps({
    api: {
        type: Object as PropType<API>
    }
})

const defaultAPI = {
    id: 0,
    gid: 0,
    method: "GET",
    name: '',
    url: '',
    paramsSource: `## example
# log_date : 20200201
`,
    headerSource: `## example
# Content-Type : application/json
`,
    bodySource: `// {
//     "k": v
// }
`,
    script: `export default {
    preRequestHandle(options) {
        return options
    },
    responseHandle(body) {
        return body
    } 
}
`
}

const options = ref<API>(merge(defaultAPI, { ...props.api }))

const validate = (): boolean => {
    if (!options.value.name) {
        return false
    }
    if (!options.value.url) {
        return false
    }
    return true
}

const requesting = ref(false)

const activeTab = ref('Params')
const changeTab = (tab: TabsPaneContext, event: Event) => {
}


const isDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
const editorOptions = {
    theme: `vs${isDarkMode ? '-dark' : ''}`,
}

const response = ref()

const parserYaml = (source: String) => {
    const tmp = source.replaceAll(/^[#\/\/].*\n?/gm, "")
    const tokens = tmp.split(/\n/).filter(e => e).map(e => {
        const index = e.indexOf(':')
        if (index === -1) {
            return null
        }

        return [e.slice(0, index).trim(), e.slice(index + 1, e.length).trim()]
    }).filter(e => e)

    const obj: Record<string, any> = {}
    if (!tokens) {
        return obj
    }

    (tokens as Array<Array<string>>).forEach((item) => {
        obj[item[0]] = item[1]
    })
    return obj
}

const parseJsonC: Function = (source: string) => {
    const tmp = source.replaceAll(/^[#\/\/].*\n?/gm, "")
    return !tmp ? {} : JSON.parse(tmp)
}

const err = ref()

const save = () => {
    if (validate()) {
        emit('saveApi', options.value)
    }
}

const bodyType = ref('html')
const bodyKey = ref(0)
const send = async () => {
    try {
        err.value = null
        requesting.value = true
        const _options = new http.Options({
            method: options.value.method,
            url: options.value.url,
            query_params: parserYaml(options.value?.paramsSource ?? ''),
            headers: parserYaml(options.value?.headerSource ?? ''),
            body: parseJsonC(options.value?.bodySource ?? '')
        })

        const script = options.value.script?.trim().replace('export default', '')

        let plugin = {
            preRequestHandle(options: http.Options) {
                return options
            },
            responseHandle(body: string) {
                return body
            }
        }

        if (script) {
            plugin = merge(plugin, (new Function(`return ${script}`))())
        }

        const formatJson = (e: string) => JSON.stringify(isString(e) ? JSON.parse(e) : e, null, 2)

        const res = await Request(_options)
        bodyType.value = 'html'
        if (res.Headers['Content-Type'][0].includes('application/json')) {
            bodyType.value = "json"
            res.Body = formatJson(plugin.responseHandle(res.Body))
        }
        response.value = res
        bodyKey.value++
    } catch (e) {
        console.error(e)
        response.value = null
        err.value = e
    } finally {
        requesting.value = false
    }
}

</script>

<style lang="scss" scoped>
.box {
    display: flex;
    flex-direction: column;
    height: 100%;

    .response {
        border: 1px solid var(--el-border-color);
        // height: 100%;
    }
}

.selection {
    margin-bottom: 20px;

    &:last-child {
        margin-bottom: 0;
    }
}

.editor {
    border: 0px !important;
}

.fullEditor {
    overflow-y: scroll;
    // height: 100% !important;
}
</style>