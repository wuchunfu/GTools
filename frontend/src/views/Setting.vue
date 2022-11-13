<template>
  <div style="height: 100%;">
    <n-card style="height: 100%;">
      <n-tabs type="line" animated default-value="system">
        <n-tab-pane name="system" tab="系统优化">
          <n-card title="软件设置" class="card-radius-10" embedded>
            <n-button secondary strong @click="cleanCache()">
              <template #icon>
                <n-icon color="#E5A241">
                  <cleancache-icon />
                </n-icon>
              </template>
              缓存清理
            </n-button>
          </n-card>
        </n-tab-pane>
        <n-tab-pane name="imgbed" tab="图床设置">
          <n-space vertical>
            <n-card title="Markdown图片存储方式" class="card-radius-10" embedded>
              <n-radio-group v-model:value="data.imgBed.configType" name="imgBedTypeRadiobuttongroup"
                :on-update="changeImgBed(data.imgBed.configType)">
                <n-radio-button v-for="song in imgBedTypes" :key="song.value" :value="song.value" :label="song.label" />
              </n-radio-group>
            </n-card>
            <n-card :title="this.imgBedCardName" embedded class="card-radius-10">
              <n-form ref="formRef" inline :label-width="80" :model="data.localImgPath" size="medium" v-show="data.imgBed.configType == 'localImgPath'">
                <n-form-item label="本地图片存储路径" path="data.localImgPath.path">
                  <n-input v-model:value="data.localImgPath.path" placeholder="本地图片存储路径" :style="{ width: '600px' }"/>
                </n-form-item>
                <n-form-item>
                  <n-button attr-type="button" @click="handleValidateClick">
                    保存
                  </n-button>
                </n-form-item>
              </n-form>
            </n-card>
          </n-space>

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
      imgBedCardName: null,
      imgBedTypes: [
        {
          label: "本地存储",
          value: "localImgPath"
        },
        {
          label: "阿里云OSS",
          value: "alioss"
        },
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
        if (res.code == 200) {
          this.data = res.data
          this.setImgBedCardName()
        }
      })
    },
    changeImgBed(val) {
      this.app.UpdateConfigItem({ "name": "configType", "value": val, "type": "imgBed" }).then(res => {
        if (res.code == 200) {
          if (res.data != null) this.data = res.data
          this.setImgBedCardName()
        } else {
          message.error(res.msg)
        }
      })
    },
    setImgBedCardName() {
      switch (this.data.imgBed.configType) {
        case "localImgPath":
          this.imgBedCardName = "本地存储配置"
          break;
        case "alioss":
          this.imgBedCardName = "阿里云OSS配置"
          break;
        default:
          break;
      }
    }
  }
}
</script>

<style scoped>
.card-radius-10 {
  border-radius: 10px;
}
</style>
