<template>
    <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
        <n-calendar style="margin: 20px;" v-model:value="value" #="{ year, month, date }"
            @update:value="handleUpdateValue">
            {{ year }}-{{ month }}-{{ date }}
        </n-calendar>
    </n-config-provider>

</template>
  
<script lang="ts">
import { defineComponent, ref } from 'vue'
import { createDiscreteApi, zhCN, dateZhCN, NConfigProvider } from 'naive-ui'
import { isYesterday, addDays } from 'date-fns/esm'
// 脱离上下文的 API 引入消息提示框
// const { message, dialog } = createDiscreteApi(
//     ["message", "dialog"]
// );

export default defineComponent({
    components: {
      NConfigProvider
    },
    setup() {
        return {
            zhCN,
            dateZhCN,
            value: ref(Date.now().valueOf()),
            handleUpdateValue(
                _: number,
                { year, month, date }: { year: number; month: number; date: number }
            ) {
                // message.success(`${year}-${month}-${date}`)
            },
            isDateDisabled(timestamp: number) {
                if (isYesterday(timestamp)) {
                    return true
                }
                return false
            }
        }
    }
})
</script>