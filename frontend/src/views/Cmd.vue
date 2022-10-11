<template>
    <div>
        <n-button type="success" @click="getWd">21313123</n-button>
    </div>
</template>

<script>
import { WindowSetLightTheme } from "../../wailsjs/runtime/runtime"
import { createDiscreteApi } from 'naive-ui'

// 脱离上下文的 API 引入消息提示框
const { message } = createDiscreteApi(
    ["message"]
);
export default {
    methods: {
        getWd() {
        }
    }
}
</script>