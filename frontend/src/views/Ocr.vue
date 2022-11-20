<template>
    <n-card :bordered="false">
        <n-alert title="注意事项" type="info" :bordered="false">
            目前只能识别系统剪贴板中的图片，请先截图或者复制图片再进行文字识别
        </n-alert>
        <n-space style="margin-top: 20px">
            <n-button type="primary" dashed strong @click="ocr(0)">
                文字识别-标准版
            </n-button>
            <n-button type="warning" dashed strong @click="ocr(1)">
                文字识别-高精度
            </n-button>
            <n-button type="success" strong secondary @click="copy" class="copy" :data-clipboard-text="ocrStr">
                <template #icon>
                    <n-icon>
                        <copy-icon />
                    </n-icon>
                </template>
                内容复制
            </n-button>
        </n-space>
        <n-divider />
        <n-spin :show="loading">
            <n-input v-model:value="ocrStr" type="textarea" placeholder="文字识别内容" :autosize="{
                minRows: 3,
                maxRows: 100
            }" />
        </n-spin>

    </n-card>

</template>
<script setup>
import { ref } from 'vue'
import { createDiscreteApi } from 'naive-ui'
import Clipboard from 'clipboard'
import { Copy as CopyIcon } from "@vicons/ionicons5";
const { message, dialog } = createDiscreteApi(
    ["message", "dialog"]
);
const ocrStr = ref("")
const loading = ref(false)
const app = ref(window.go.gtools.App)

const ocr = (ocrType) => {
    loading.value = true
    app.value.BaiduOCR(ocrType).then(res => {
        if (res.code == 200) {
            ocrStr.value = res.data
        } else {
            message.error(res.msg)
        }
        loading.value = false
    })
}

const copy = () => {
    if(ocrStr.value == "") {
        return
    }
    let clipboard = new Clipboard('.copy')
    clipboard.on('success', (e) => {
        message.success('复制成功')
        clipboard.destroy()
    })
}
</script>