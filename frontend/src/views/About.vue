<template>
    <!-- <button @click="invokeHtmlMethod">调用html种方法</button> -->
    <div style="height: 100%; overflow: hidden; -webkit-overflow-scrolling: touch;">
      <iframe  name = "iframeMap" id="iframeMapViewComponent"  v-bind:src="getPageUrl"
               width="100%" height="100%"
               frameborder="0" scrolling="yes" ref="iframeDom"
      ></iframe>
    </div>
</template>
<script>
export default {
    data() {
      return {
        getPageUrl: '/Users/pixiao/WorkSpace/Tomcat/webapps/docs/index.html'
      }
    },
    created() {
      // 初始化时为window绑定一个方法
      window['vueDefinedMyProp'] = (receiveParams) => {
        this.receiveParamsFromHtml(receiveParams)
      }
    },
    methods: {
      receiveParamsFromHtml(res) {
        console.log(res)
      },
      invokeHtmlMethod() {
        window.frames['iframeMap'].lodaTable()
    }
      },
}
</script>