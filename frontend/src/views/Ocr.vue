<template>
    <n-button type="primary" dashed @click="Ocr">
        Primary
    </n-button>
    <n-input v-model:value="ocrStr" type="textarea" placeholder="基本的 Textarea" />
</template>
<script setup>
import { ref } from 'vue'
import { createDiscreteApi } from 'naive-ui'
const { message, dialog } = createDiscreteApi(
    ["message", "dialog"]
);
const ocrStr = ref("")

const Ocr = () => {
    window.go.gtools.App.BaiduOCR(0).then(res => {
        if (res.code == 200) {
            ocrStr.value = res.data
        } else {
            message.error(res.msg)
        }

    })
}
</script>
<style>

</style>