<template>
    <el-popover class="menu-item" placement="bottom-start" :width="400" trigger="hover">
        <template #reference>
            <div>
                <span>Workspace</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
        </template>
        <div style="width:400px">
            <el-button link size="small">
                <el-icon @click="handleAdd">
                    <Plus />
                </el-icon>
            </el-button>
            <div class="table">
                <el-table :data="wsList" :show-header="false">
                    <el-table-column width="150" property="name" label="date" />
                    <el-table-column align="right">
                        <template #default="scope">
                            <el-button link size="small" @click="handleEdit(scope.$index, scope.row)">
                                <el-icon>
                                    <EditPen />
                                </el-icon>
                            </el-button>
                            <el-button link size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">
                                <el-icon>
                                    <Remove />
                                </el-icon>
                            </el-button>
                            <el-button link :type="activeWsId === scope.row.id ? 'success' : 'info'" :icon="Check"
                                circle @click="setCurrentWs(scope.row)" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </div>
    </el-popover>
    <el-dialog v-model="dialog" width="80%">
        <VForm :key="formKey" v-model="workspace" v-bind="editForm" @submit="saveWs"></VForm>
    </el-dialog>
</template>
<script lang="ts" setup>
import { app } from '~/go/models'
import { ArrowDown, EditPen, Remove, Plus, Check } from '@element-plus/icons-vue'
import { CreateWorkspace, UpdateWorkspace, WorkspaceList } from '~/go/app/App'
import { VForm } from '@okiss/vbtf'
import { ElMessage } from 'element-plus'
import { useAppStore } from '@/store/app'

function load() {
    WorkspaceList().then(res => {
        wsList.value = res as Array<app.Workspace>
    })
}

load()

const store = useAppStore()
const dialog = ref(false)
interface ws {
    id: number,
    name: string,
    env: string
}

const defaultEnv = `{
  "local": {},
  "pre": {},
  "prod": {}
}`

const editForm = {
    formItems: [
        {
            field: 'id',
            type: 'hidden'
        },
        {
            field: 'name',
            rules: 'required'
        },
        {
            field: 'env',
            type: 'json',
            value: defaultEnv,
            rules: 'required'
        }
    ]
}

const activeWsId = computed(() => store.activeWsId)

const workspace = ref<ws>()
const wsList = ref<Array<app.Workspace>>()
const formKey = ref(0)
const handleAdd = () => {
    dialog.value = true
    workspace.value = {
        id: 0,
        name: 'NewWorkspace',
        env: defaultEnv
    }
}

const handleEdit = (index: number, row: app.Workspace) => {
    formKey.value++
    dialog.value = true
    const conf = JSON.parse(row.config)
    workspace.value = {
        id: row.id,
        name: row.name,
        env: conf.env
    }
}
const handleDelete = (index: number, row: app.Workspace) => { }

const saveWs = (form: Record<string, any>) => {
    const conf = {
        env: form.env
    }
    const ws: app.Workspace = new app.Workspace({ id: form.id || 0, name: form.name })
    ws.config = JSON.stringify(conf)

    const action = ws.id ? UpdateWorkspace : CreateWorkspace

    action(ws).then(res => {
        dialog.value = false
        workspace.value = undefined
        load()
    }).catch(e => {
        ElMessage.error(e)
    })

}
const setCurrentWs = (row: app.Workspace) => {
    store.SetActiveWs(row.id)
}

</script>