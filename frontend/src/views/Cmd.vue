<template>
    <div>
        <n-button type="success" @click="startTomcat">21313123</n-button>
    </div>
</template>

<script>
import { createDiscreteApi } from 'naive-ui'
import { StartTomcat }from '../../wailsjs/go/main/App'

// 脱离上下文的 API 引入消息提示框
const { message } = createDiscreteApi(
    ["message"]
);
export default {
    methods: {
        startTomcat() {
            StartTomcat().then((res) => {
                console.log(res);
            })
        }
    }
}
</script>