<template>
  <n-space vertical size="large">
    <n-layout has-sider>
      <n-layout-sider content-style="padding: 24px;" :bordered="true" collapse-mode="width"
                      :collapsed="collapsed" show-trigger @collapse="collapsed = true" @expand="collapsed = false"
                      :collapsed-width="0" :width="250">
        <n-input-group>
          <n-input :style="{ width: '70%' }" placeholder="添加文件夹" v-model:value="dirPath"/>
          <n-button type="primary" @click="addDirPath">
            添加
          </n-button>
        </n-input-group>
      </n-layout-sider>
      <n-layout>
        <n-layout-content>
          <div id="vditor" class="vditor"></div>
        </n-layout-content>
      </n-layout>
    </n-layout>
  </n-space>
</template>
<script>
import Vditor from "vditor"
import "vditor/dist/index.css"
import { GetMdContent, UploadImgByPicgo } from "../../wailsjs/go/main/App"
import { createDiscreteApi } from 'naive-ui'

// 脱离上下文的 API 引入消息提示框
const { message } = createDiscreteApi(
    ["message"]
);

export default {
  props: ['path'],
  setup(props) {
    // const route = useRoute()
    // localStorage.setItem('mdPath', route.query.path)
  },
  data() {
    return {
      collapsed: true,
      contentEditor: "",
      dirPath: null,
      toolbar: [
        'emoji',
        'headings',
        'bold',
        'italic',
        'strike',
        'link',
        '|',
        'list',
        'ordered-list',
        'check',
        'outdent',
        'indent',
        '|',
        'quote',
        'line',
        'code',
        'inline-code',
        'insert-before',
        'insert-after',
        '|',
        'upload',
        'record',
        'table',
        '|',
        'undo',
        'redo',
        '|',
        'edit-mode',
        'fullscreen',
        'outline',
        {
          hotkey: '⇧⌘S',
          name: 'sponsor',
          tipPosition: 's',
          tip: '成为赞助者',
          className: 'right',
          icon: '<svg t="1663730736397" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9510" width="200" height="200"><path d="M511.250625 414.911719a46.545031 46.545031 0 0 1 46.545031 46.545031l0.162908 283.575604a46.545031 46.545031 0 0 1-93.090063 0l-0.162907-283.575604a46.545031 46.545031 0 0 1 46.545031-46.545031z m-50.012636-136.53985a50.035909 50.035909 0 0 1 100.071817 0l0.18618 1.512714a50.035909 50.035909 0 0 1-100.071817 0zM511.995345 1024a508.178653 508.178653 0 0 1-293.233697-93.299515 46.405396 46.405396 0 0 1-34.210598-44.683231l-0.418906-4.305415a46.405396 46.405396 0 0 1 80.592722-31.557531 420.534359 420.534359 0 1 0-132.653339-160.161453l-7.540295 7.540295a46.545031 46.545031 0 0 1 29.020827 43.077426l0.442177 4.328688a46.428669 46.428669 0 0 1-91.088626 12.776611A511.995345 511.995345 0 1 1 511.995345 1024z" p-id="9511"></path></svg>',
          click() { console.log(213); },
        },
        {
          name: 'more',
          toolbar: [
            'export',
            'preview',
            'content-theme',
            'code-theme',
            'info',
            'help',
          ],
        }],
    }
  },
  //mounted
  mounted() {
    let _this = this
    this.contentEditor = new Vditor('vditor', {
      height: '700px', // 默认就是全屏模式，但是此项不能省略，否则会出现大纲无法导航的问题
      width: '100%',
      toolbar: this.toolbar,
      toolbarConfig: {
        pin: true,
        hide: false
      },
      outline: {
        enable: true,
        position: 'left'
      },
      preview: {
        hljs: {
          enable: true,
          style: 'vim',
          lineNumber: true
        },
        theme: {
          current: 'ant-design'
        }
      },
      resize: {
        enable: true,
        after(h) {
          console.log(h);
        }
      },
      cache: {
        enable: true,
      },
      after: () => {

      },
      input: () => {

      },
      esc: () => {
        console.log("esc");
        localStorage.removeItem('mdPath')
      },
      upload: {
        handler(){
          _this.uploadClipboardToOSS()
        }
      }
    })
  },
  methods: {
    getValue() {
      return this.contentEditor.getValue();     //获取 Markdown 内容
    },
    getHTML() {
      return this.contentEditor.getHTML();      //获取 HTML 内容
    },
    setValue(value) {
      return this.contentEditor.setValue(value);     //设置 Markdown 内容
    },
    disabled() {
      return this.contentEditor.disabled();     //设置 只读
    },
    insertValue(str, render) {
      return this.contentEditor.insertValue(str, render)
    },
    toPreview() {
      var evt = document.createEvent('Event');
      evt.initEvent('click', true, true);
      this.contentEditor.vditor.toolbar.elements.preview.firstElementChild.dispatchEvent(evt);
    },
    // getMdContent() {
    //   GetMdContent(this.mdPath).then((res) => {
    //     let rest = JSON.parse(res)
    //     this.contentEditor.setValue(rest.data)
    //   })
    // },
    uploadClipboardToOSS() {
      let _this = this
      UploadImgByPicgo().then((res) => {
        _this.contentEditor.insertValue("![url]("+res+")\n", true)
      })
    },
    addDirPath() {
      message.info(this.dirPath)
    }
  },

}
</script>
