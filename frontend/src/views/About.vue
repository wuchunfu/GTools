<template>
<n-button @click="fanyi">123</n-button>
<n-input
      v-model:value="invalue"
      type="textarea"
      placeholder="基本的 Textarea"
    />
    <n-input
      v-model:value="trans"
      type="textarea"
      placeholder="基本的 Textarea"
    />
</template>
<script setup>
import { ref } from "vue";
const app = ref(window.go.gtools.App)
const invalue = ref(null)
const trans = ref(null)

const fanyi = () => {
  app.value.BaiDuTrans(invalue.value).then(res => {
    if(res.code == 200) {
      trans.value = res.data
    } else {
      console.log(res.msg);
    }
    
  })
}
</script>