<template>
    <div style="padding: 20px">
        <n-form ref="formRef" inline :label-width="80" :model="cmdForm" size="medium">
            <n-grid x-gap="12" :cols="3">
                <n-gi>
                    <n-form-item label="指令名称" path="cmd.name">
                        <n-input v-model:value="cmdForm.name" placeholder="输入名称"  />
                    </n-form-item>
                    <n-form-item label="启动指令" path="cmd.start">
                        <n-input v-model:value="cmdForm.start" type="textarea" placeholder="启动" :rows="1" />
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="程序端口" path="cmd.port">
                        <n-input v-model:value="cmdForm.port" placeholder="输入端口号" />
                    </n-form-item>


                    <n-form-item label="停止指令" path="cmd.stop">
                        <n-input v-model:value="cmdForm.stop" type="textarea" placeholder="停止" :rows="1" />
                    </n-form-item>
                </n-gi>
                <n-gi>
                    <n-form-item label="执行方式" path="cmd.type">
                        <n-select v-model:value="cmdForm.type" :options="types" 
                            placeholder="请选择" />
                    </n-form-item>
                    <n-form-item>
                        <n-button attr-type="button" @click="addCmdItem">
                            添加指令
                        </n-button>
                    </n-form-item>
                </n-gi>
            </n-grid>



        </n-form>
        <n-divider />
        <div>
            <n-grid x-gap="12" :cols="3">
                <n-gi v-for="item in data">
                    <div style="margin-top: 10px;">
                        <n-card :title="item.name" style="border-radius: 10px;">
                            <div style="display: flex; height: 40px; line-height: 40px;">
                                端口：{{ item.port }}
                                <div style="margin-left: auto">
                                    <n-button circle style="margin-right: 10px;">
                                        <template #icon>
                                            <n-icon>
                                                <edit-icon />
                                            </n-icon>
                                        </template>
                                    </n-button>
                                    <n-button circle  @click="cmdHandler(item)">
                                        <template #icon>
                                            <n-icon v-if="item.state == 0">
                                                <start-icon />
                                            </n-icon>
                                            <n-icon v-else>
                                                <stop-icon />
                                            </n-icon>
                                        </template>
                                    </n-button>
                                </div>
                            </div>
                        </n-card>
                    </div>
                </n-gi>
            </n-grid>
        </div>
    </div>

</template>

<script>
import { createDiscreteApi } from 'naive-ui'
import {
    CaretForward as StartIcon,
    Pause as StopIcon, BuildSharp as EditIcon
} from '@vicons/ionicons5'

// 脱离上下文的 API 引入消息提示框
const { message } = createDiscreteApi(
    ["message"]
);
export default {
    components: {
        StartIcon,
        StopIcon,
        EditIcon
    },
    data() {
        return {
            app: window.go.gtools.App,
            formRef: null,
            types: [
                {
                    label: "sh",
                    value: "sh"
                },
                {
                    label: "无",
                    value: ""
                },
            ],
            cmdForm: {
                name: null,
                type: null,
                start: null,
                stop: null,
                port: null,
            },
            data: []
        }
    },
    mounted() {
        this.getCmdItemList()
    },
    methods: {
        getCmdItemList() {
            this.app.GetCmdItemList().then(res => {
                if (res.code == 200) {
                    this.data = res.data
                }
            })
        },
        addCmdItem() {
            this.app.AddCmdItem(this.cmdForm).then(res => {
                if (res.code == 200) {
                    this.data = res.data
                } else {
                    message.error(res.msg)
                }
            })
        },
        cmdHandler(item) {
            this.app.CmdHandler(item).then((res) => {
                if (res.code == 200) {
                    this.data = res.data
                } else {
                    message.error(res.msg)
                }
            })
        },
        selectFile() {
            this.app.OpenSelectFileWindow().then((res) => {
                console.log(res);
            })

        }
    }
}
</script>