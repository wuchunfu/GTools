<template>
  <div style="height: 100%;">
    <n-card style="height: 100%;">
      <n-tabs type="line" animated default-value="system">
        <n-tab-pane name="system" tab="系统优化">
          <n-card title="软件设置" style="border-radius: 10px;">
            <n-button secondary strong @click="cleanCache()">
              <template #icon>
                <n-icon>
                  <cleancache-icon />
                </n-icon>
              </template>
              Mac缓存清理
            </n-button>
          </n-card>
        </n-tab-pane>
        <n-tab-pane name="imgbed" tab="图床设置">
          <n-card title="Markdown图片存储方式" style="border-radius: 10px;">
            <n-radio-group v-model:value="data.imgBed.configType" name="imgBedTypeRadiobuttongroup"
              :on-update="changeImgBed(data.imgBed.configType)">
              <n-radio-button v-for="song in imgBedTypes" :key="song.value" :value="song.value" :label="song.label" />
            </n-radio-group>
          </n-card>
        </n-tab-pane>
        <n-tab-pane name="jay chou" tab="其他设置">
          我爱波多野结衣
        </n-tab-pane>
      </n-tabs>
    </n-card>
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
      app: go.gtools.App,
      imgBedTypes: [
        {
          label: "本地存储",
          value: "localImgPath"
        },
        {
          label: "阿里云OSS",
          value: "alioss"
        },
        // {
        //   label: "腾讯云OSS",
        //   value: ""
        // },
        // {
        //   label: "sm.ms",
        //   value: ""
        // },
      ],
      data: {}
    }
  },
  name: "Setting",
  mounted() {
    this.getConfigMap()
  },
  methods: {
    cleanCache() {
      this.app.CleanWebKitCache().then(res => {
        if (res.code == 200) message.success("缓存已清理")
      })
    },
    getConfigMap() {
      this.app.GetConfigOnounted().then(res => {
        if (res.code == 200) this.data = res.data
      })
    },
    changeImgBed(val) {
      this.app.UpdateConfigItem({"name": "configType", "value": val, "type": "imgBed"}).then(res => {
        if(res.code == 200) {
          if(res.data != null) this.data = res.data
        } else {
          message.error(res.msg)
        }
      })
    },
  }
}
</script>

<style scoped>

</style>
