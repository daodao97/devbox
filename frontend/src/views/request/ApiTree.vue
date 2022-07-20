<template>
    <div class="mybox" data-wails-no-drag>
        
        <div class="action">
            <div>
                <Workspace></Workspace>
            </div>
            <div class="ctrl">
                <div class="add-group" @click="addGroup">
                    <el-icon>
                        <Plus />
                    </el-icon>
                </div>
                <div class="add-group">
                    <el-icon>
                        <Folder />
                    </el-icon>
                </div>
            </div>
        </div>

        <el-input v-model="filterText" placeholder="Filter keyword" class="filter" />

        <el-tree ref="treeRef" node-key="id" :key="treeKey" :default-expanded-keys="expandKeys"
            :expand-on-click-node="false" :data="props.tree" :props="defaultProps" :icon="''"
            :filter-node-method="filterNode">
            <template #default="{ node, data }">
                <div class="tree-node" v-if="data.workspace_id">
                    <div class="name-ctrl">
                        <div class="text">{{ data.name }}</div>
                        <el-input class="input" size="small" v-model="data.name" @change="() => saveGroup(data)">
                        </el-input>
                    </div>
                    <div class="icon" @click="() => addApi(data)">
                        <el-icon>
                            <Plus />
                        </el-icon>
                    </div>
                </div>
                <div v-if="data.gid" :class="['tree-node', data.gid && (data.id === activeApiId) ? 'active-api' : '']"
                    @click="() => clickAPI(data)">
                    <div class="name">
                        {{ data.id === activeApi?.id ? (activeApi?.name || data.name) : data.name }}
                    </div>
                    <div :class="['method', 'method_' + data.method.toLowerCase()]">
                        {{ data.id === activeApi?.id ? (activeApi?.method || data.method) : data.method }}
                    </div>
                </div>
            </template>
        </el-tree>
    </div>
</template>

<script lang="ts" setup>
import { ElMessage, ElTree } from 'element-plus'
import { Plus, Folder, Search } from '@element-plus/icons-vue';
import { PropType } from 'vue';
import { app } from '~/go/models'
import { CreateApi, CreateGroup, UpdateApi, UpdateGroup } from '~/go/app/App';
import Workspace from './workspace.vue'

const props = defineProps({
    tree: {
        type: Object as PropType<Array<app.ApiGroup>>,
        default: () => []
    },
    workspaceId: {
        type: Number,
        default: 1
    },
    activeGroupId: {
        type: Array<string | number>,
        default: () => [1]
    },
    activeApiId: {
        type: Number,
        default: 1
    },
    activeApi: {
        type: Object as PropType<app.ApiList>,
    }
})

interface TreeNodeData {
    [key: string]: any;
}

const filterText = ref('')
const treeRef = ref<InstanceType<typeof ElTree>>()

watch(filterText, (val) => {
    treeRef.value!.filter(val)
})

const filterNode = (value: string, data: TreeNodeData) => {
    if (!value) return true
    return data.name.includes(value)
}

const emit = defineEmits(['change', 'saveGroup'])

const clickAPI = (data: app.ApiList) => {
    if (data.id > 0) {
        emit('change', data)
    }
}

const defaultProps = {
    children: 'children',
    label: 'label'
}
const saveGroup = (data: app.ApiGroup) => {
    const action = data.id ? UpdateGroup : CreateGroup
    action(data).then(res => {
        emit('saveGroup', data)
    }).catch(e => {
        ElMessage.error(e)
    })
}
const addGroup = async () => {
    console.log(props.tree)
    const group = new app.ApiGroup({
        id: 0,
        name: "NewGroup",
        workspace_id: props.workspaceId,
        config: '{}'
    })
    await saveGroup(group)
    props.tree.push(group)
    treeKey.value++
}


const expandKeys = ref(props.activeGroupId)
const treeKey = ref(0)
const addApi = (data: app.ApiGroup) => {
    const api = new app.ApiList({
        gid: data.id,
        method: "GET",
        name: 'NewRequest',
        options: '{}'
    })
    data.children = data.children || []
    CreateApi(api).then(id => {
        api.id = id as number
        data.children.push(api)
        expandKeys.value.push(data.id)
        treeKey.value++
    }).catch(e => {
        ElMessage.error(e)
    })
}
</script>
<style lang="scss" scope>
.mybox {
    height: 100%;
    // border: 1px solid var(--el-border-color);
}

.tree-node {
    display: flex;
    width: 100%;
    font-size: 13px;
    align-items: center;
    justify-content: space-between;

    .icon {
        padding: 10px;

        &:hover {
            color: var(--el-color-primary);
        }
    }

    .method {
        font-size: 11px;
        margin: 5px;
        padding: 2px;
    }

    .method_get {
        background-color: #409EFF;
    }

    .method_post {
        background-color: #67C23A;
    }

    .method_delete {
        background-color: #F56C6C;
    }

    .method_option {
        background-color: #CDD0D6;
    }

    .method_head {
        background-color: #CDD0D6;
    }
}

.action {
    display: flex;
    font-size: 12px;
    align-items: center;
}

.filter {
    margin-bottom: 10px;
}

.ctrl {
    display: flex;
    justify-content: flex-end;
    width: 100%;

    .add-group {
        width: 40px;
        height: 40px;
        text-align: center;
        line-height: 40px;
    }
}

.name-ctrl {
    .name {
        font-size: 14px;
        padding: 0 5px;
    }

    .input {
        display: block;
    }

    &:hover .text {
        display: none;
    }

    .input {
        display: none;
    }

    &:hover .input {
        display: block;
    }
}
</style>