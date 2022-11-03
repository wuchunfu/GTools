<template>
  <div style="padding: 20px">
    <n-button secondary strong @click="cleanCache()">
      <template #icon>
        <n-icon>
          <cleancache-icon />
        </n-icon>
      </template>
      Mac缓存清理
    </n-button>
  </div>
</template>

<script>
import { Trash as CleancacheIcon } from "@vicons/ionicons5";
import { createDiscreteApi } from 'naive-ui'
// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
  ["message", "dialog"]
);
export default {
  components: {
    CleancacheIcon
  },
  data() {
    return {
      app: go.gtools.App
    }
  },
  name: "Setting",
  methods: {
    cleanCache() {
        this.app.CleanWebKitCache().then(res => {
          if (res.code == 200) message.success("缓存已清理")
        })
    }
  }
}
</script>

<style scoped>

</style>
