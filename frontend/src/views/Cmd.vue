<template>
    <div>
        <n-button type="success" @click="startTomcat">21313123</n-button>
        <n-button @click="selectFile">选择文件</n-button>
    </div>
</template>

<script>
import { createDiscreteApi } from 'naive-ui'

// 脱离上下文的 API 引入消息提示框
const { message } = createDiscreteApi(
    ["message"]
);
export default {
    data() {
        return {
            app: window.go.gtools.App
        }
    },
    methods: {
        startTomcat() {
            this.app.StartTomcat().then((res) => {
                console.log(res);
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