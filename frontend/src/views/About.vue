<template>
  <n-card title="下载">
    <n-button @click="newEvent">start</n-button>
    <n-button @click="eventOff">off</n-button>
    <n-button @click="testEvent">start</n-button>
    <n-divider />
    <n-progress type="line" :percentage="result" :indicator-placement="'inside'" processing />
  </n-card>

</template>
<script setup>
import { ref, watch } from "vue";
// import { EventsOn } from "../../wailsjs/runtime/runtime"
import { EventsEmit, EventsOn, EventsOff } from "/wailsjs/runtime/runtime"
const app = ref(window.go.gtools.App)
const rt = ref(runtime.runtime)
const result = ref(0)

watch(result, (newValue, oldValue) => {
  if (result.value == 100) {
    EventsOff("download")
  }
})

const newEvent = () => {
  EventsEmit("test", {
    test: 123
  })
  EventsOn("download", res => {
    result.value = res
    console.log(res);
  })
}

const eventOff = () => {
  EventsOff("download")
}

const testEvent = () => {
  app.value.TestEvent()
}
</script>