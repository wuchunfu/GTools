<template>
  <div style="height: 100%;">
    <n-card :bordered="false">
      <n-tabs type="line" animated default-value="system">
        <n-tab-pane name="system" tab="软件设置">
          <n-space vertical>
            <n-card title="软件优化" class="card-radius-10" embedded>
              <n-button secondary strong @click="cleanCache()">
                <template #icon>
                  <n-icon color="#E5A241">
                    <cleancache-icon />
                  </n-icon>
                </template>
                缓存清理
              </n-button>
            </n-card>
            <n-card title="Markdown图片存储方式" class="card-radius-10" embedded>
              <n-radio-group v-model:value="data.imgbed.configType" name="imgBedTypeRadiobuttongroup"
              :on-update="changeImgBed(data.imgbed.configType)">
                <n-radio-button v-for="song in imgBedTypes" :key="song.value" :value="song.value" :label="song.label" />
              </n-radio-group>
            </n-card>
            <n-card title="翻译方式" class="card-radius-10" embedded>
              <n-radio-group v-model:value="data.trans.transType" name="transTypeRadiobuttongroup"
                :on-update="changeTrans(data.trans.transType)">
                <n-radio-button v-for="song in transTypes" :key="song.value" :value="song.value" :label="song.label" />
              </n-radio-group>
            </n-card>
          </n-space>
        </n-tab-pane>
        <n-tab-pane name="imgbed" tab="图床配置">
          <n-space vertical>
            <n-card title="本地存储配置" embedded class="card-radius-10">
              <n-form inline :label-width="80" :model="data.limgpath" size="medium">
                <n-form-item label="本地图片存储路径" path="data.limgpath.path">
                  <n-input v-model:value="data.limgpath.path" placeholder="本地图片存储路径" :style="{ width: '600px' }" />
                </n-form-item>
                <n-form-item>
                  <n-button attr-type="button" @click="updateConfigByType('limgpath', data.limgpath)" type="success">
                    更新
                  </n-button>
                </n-form-item>
              </n-form>
            </n-card>
            <n-card title="阿里云OSS配置" embedded class="card-radius-10">
              <n-alert title="注意事项" type="info" v-show="data.imgbed.configType == 'alioss'" :bordered="false">
                Region、Bucket、Object内容请勿添加任何符号
              </n-alert>
              <n-form inline :label-width="80" style="margin-top: 30px;" :model="data.alioss" size="medium">
                <n-grid cols="2 400:2 800:3 1000:4">
                  <n-grid-item>
                    <n-form-item label="存储地域(Region)">
                      <n-input v-model:value="data.alioss.point" placeholder="存储地域名称" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="访问密钥id">
                      <n-input v-model:value="data.alioss.accessKeyId" placeholder="访问密钥id" :style="{ width: '250px' }"
                        type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="访问密钥key">
                      <n-input v-model:value="data.alioss.accessKeySecret" placeholder="访问密钥key"
                        :style="{ width: '250px' }" type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="存储空间(Bucket)">
                      <n-input v-model:value="data.alioss.bucketName" placeholder="存储空间名称"
                        :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="对象/文件(Object)">
                      <n-input v-model:value="data.alioss.projectDir" placeholder="对象或文件" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item>
                      <n-button attr-type="button" @click="updateConfigByType('alioss', data.alioss)" type="success">
                        更新
                      </n-button>
                    </n-form-item>
                  </n-grid-item>
                </n-grid>
              </n-form>
            </n-card>
          </n-space>
        </n-tab-pane>
        <n-tab-pane name="ocr" tab="图文配置">
          <n-space vertical>
            <n-card title="百度OCR" embedded class="card-radius-10">
              <n-form inline :label-width="80" :model="data.bdocr" size="medium">
                <n-grid cols="2 400:2 800:3 1000:4">
                  <n-grid-item>
                    <n-form-item label="ID(clientId)">
                      <n-input v-model:value="data.bdocr.clientId" placeholder="存储地域名称" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="密钥(clientSecret)">
                      <n-input v-model:value="data.bdocr.clientSecret" placeholder="访问密钥id" :style="{ width: '250px' }"
                        type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item>
                      <n-button attr-type="button" @click="updateConfigByType('bdocr', data.bdocr)" type="success">
                        更新
                      </n-button>
                    </n-form-item>
                  </n-grid-item>
                </n-grid>
              </n-form>
            </n-card>
            <n-card title="百度翻译(翻译开放平台)" embedded class="card-radius-10">
              <n-alert title="注意事项" type="info" :bordered="false">
                百度翻译免费额度下并发量为1，请勿在短时间内快速、重复地进行翻译。<br>
                appid 和 secret 务必准确，其他三项根据实际需求配置。
              </n-alert>
              <n-form inline :label-width="80" :model="data.bdtrans" size="medium" style="margin-top: 15px">
                <n-grid cols="2 400:2 800:3 1000:4">
                  <n-grid-item>
                    <n-form-item label="appid">
                      <n-input v-model:value="data.bdtrans.appid" placeholder="appid" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="secret">
                      <n-input v-model:value="data.bdtrans.secret" placeholder="secret" :style="{ width: '250px' }"
                        type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="salt(盐默认gtools)">
                      <n-input v-model:value="data.bdtrans.salt" placeholder="salt" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="from(默认auto自动识别)">
                      <n-input v-model:value="data.bdtrans.from" placeholder="from" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="to(默认翻译为中文)">
                      <n-input v-model:value="data.bdtrans.to" placeholder="to" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item>
                      <n-button attr-type="button" @click="updateConfigByType('bdtrans', data.bdtrans)" type="success">
                        更新
                      </n-button>
                    </n-form-item>
                  </n-grid-item>
                </n-grid>
              </n-form>
            </n-card>
            <n-card title="腾讯云翻译" embedded class="card-radius-10">
              <n-alert title="注意事项" type="info" :bordered="false">
                腾讯云翻译免费额度下默认接口请求频率限制：5次/秒，文本翻译的每月免费额度为5百万字符，免费额度每月1日0点重置。<br>
              </n-alert>
              待对接
              <!-- <n-form inline :label-width="80" :model="data.bdtrans" size="medium" style="margin-top: 15px">
                <n-grid cols="2 400:2 800:3 1000:4">
                  <n-grid-item>
                    <n-form-item label="appid">
                      <n-input v-model:value="data.bdtrans.appid" placeholder="appid" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="secret">
                      <n-input v-model:value="data.bdtrans.secret" placeholder="secret" :style="{ width: '250px' }"
                        type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="salt(盐默认gtools)">
                      <n-input v-model:value="data.bdtrans.salt" placeholder="salt" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="from(默认auto自动识别)">
                      <n-input v-model:value="data.bdtrans.from" placeholder="from" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="to(默认翻译为中文)">
                      <n-input v-model:value="data.bdtrans.to" placeholder="to" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item>
                      <n-button attr-type="button" @click="updateConfigByType('bdtrans', data.bdtrans)" type="success">
                        更新
                      </n-button>
                    </n-form-item>
                  </n-grid-item>
                </n-grid>
              </n-form> -->
            </n-card>
            <n-card title="阿里云机器翻译" embedded class="card-radius-10">
              <n-alert title="注意事项" type="info" :bordered="false">
                阿里云机器翻译免费额度下QPS为50，通用版机器翻译每月100万字符免费额度，单次请求字符不要太大(控制在1000字符以内)，否则请使用文本翻译。<br>
              </n-alert>
              待对接
              <!-- <n-form inline :label-width="80" :model="data.bdtrans" size="medium" style="margin-top: 15px">
                <n-grid cols="2 400:2 800:3 1000:4">
                  <n-grid-item>
                    <n-form-item label="appid">
                      <n-input v-model:value="data.bdtrans.appid" placeholder="appid" :style="{ width: '250px' }" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="secret">
                      <n-input v-model:value="data.bdtrans.secret" placeholder="secret" :style="{ width: '250px' }"
                        type="password" :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="salt(盐默认gtools)">
                      <n-input v-model:value="data.bdtrans.salt" placeholder="salt" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="from(默认auto自动识别)">
                      <n-input v-model:value="data.bdtrans.from" placeholder="from" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item label="to(默认翻译为中文)">
                      <n-input v-model:value="data.bdtrans.to" placeholder="to" :style="{ width: '250px' }"
                        :clearable="true" />
                    </n-form-item>
                  </n-grid-item>
                  <n-grid-item>
                    <n-form-item>
                      <n-button attr-type="button" @click="updateConfigByType('bdtrans', data.bdtrans)" type="success">
                        更新
                      </n-button>
                    </n-form-item>
                  </n-grid-item>
                </n-grid>
              </n-form> -->
            </n-card>
          </n-space>
        </n-tab-pane>
        <n-tab-pane name="other" tab="其他配置">
          我爱波多野结衣
        </n-tab-pane>
        <n-tab-pane name="instructions" tab="使用说明">
          我爱波多野结衣
        </n-tab-pane>
        <n-tab-pane name="feedback" tab="问题反馈">
          我爱波多野结衣
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script async setup>
import { Trash as CleancacheIcon } from "@vicons/ionicons5";
import { createDiscreteApi } from 'naive-ui'
import { ref, onBeforeMount } from "vue";
// 生命周期放在最前面å
onBeforeMount(() => {
  getConfigMap()
})

// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
  ["message", "dialog"]
);
const imgBedCardName = ref(null)
const imgBedTypes = ref([
  {
    label: "本地存储",
    value: "limgpath"
  },
  {
    label: "阿里云OSS",
    value: "alioss"
  },
])
const transTypes = ref([
  {
    label: "百度翻译",
    value: "bdtrans"
  },
  {
    label: "阿里翻译",
    value: "alitrans"
  },
  {
    label: "腾讯翻译",
    value: "trans"
  },
])
const data = ref({
  // 默认值必须有，否则会重复触发更新方法
  imgbed: {
    configType: null
  },
  trans: {
    transType: null
  }
})
const app = ref(window.go.gtools.App)

const cleanCache = () => {
  app.value.CleanWebKitCache().then(res => {
    if (res.code == 200) message.success("缓存已清理")
  })
}
const getConfigMap = () => {
  app.value.GetConfigOnounted().then(res => {
    if (res.code == 200) {
      data.value = res.data
      setImgBedCardName()
    }
  })
}

const changeImgBed = (val) => {
  if(val == null) return
  app.value.UpdateConfigItem({ "name": "configType", "value": val, "type": "imgbed" }).then(res => {
    if (res.code == 200) {
      if (res.data != null) data.value = res.data
      setImgBedCardName()
    } else {
      message.error(res.msg)
    }
  })
}
const changeTrans = (val) => {
  if(val == null) return
  app.value.UpdateConfigItem({ "name": "transType", "value": val, "type": "trans" }).then(res => {
    if (res.code == 200) {
      if (res.data != null) data.value = res.data
      setImgBedCardName()
    } else {
      message.error(res.msg)
    }
  })
}

const updateConfigByType = (type, value) => {
  app.value.UpdateConfigByType({ "type": type, "value": value }).then(res => {
    if (res.code == 200) {
      if (res.data != null) data.value = res.data
      // this.setImgBedCardName()
      message.success("配置已更新")
    } else {
      message.error(res.msg)
    }
  })
}
const setImgBedCardName = () => {
  switch (data.value.imgbed.configType) {
    case "limgpath":
      imgBedCardName.value = "本地存储配置"
      break;
    case "alioss":
      imgBedCardName.value = "阿里云OSS配置"
      break;
    default:
      break;
  }
}
</script>

<style scoped>
.card-radius-10 {
  border-radius: 10px;
}
</style>
