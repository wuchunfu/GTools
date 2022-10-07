<template>
  <n-space vertical size="large">
    <n-layout has-sider>
      <n-layout-sider content-style="padding: 24px;" :bordered="true" collapse-mode="width" :collapsed="collapsed"
        @collapse="collapsed = true" @expand="collapsed = false" :collapsed-width="0" :width="250">
        <n-input-group>
          <n-input :style="{ width: '70%' }" placeholder="添加文件夹" v-model:value="dirPath" />
          <n-button type="primary" @click="addDirPath">
            添加
          </n-button>
        </n-input-group>
        <div class="treeDiv">
          <n-tree block-line :data="treeData" :selectable="false" :node-props="nodeProps" />
        </div>
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
import {GetMdContent, UploadImgByPicgo, AddDirPath, GetDirList} from "../../wailsjs/go/main/App"
import { createDiscreteApi, NButton } from 'naive-ui'
import { h } from "vue";

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
      collapsed: false,
      contentEditor: "",
      dirPath: null,
      treeData: [
        {
          label: 'test',
          key: 'test'
        }
      ]
    }
  },
  //mounted
  mounted() {
    let _this = this
    this.contentEditor = new Vditor('vditor', {
      height: '700px', // 默认就是全屏模式，但是此项不能省略，否则会出现大纲无法导航的问题
      width: '100%',
      toolbar: [
        {
          name: 'sponsor',
          tipPosition: 's',
          tip: '菜单展开',
          className: 'right',
          icon: '<svg t="1665030513743" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10508" width="200" height="200"><path d="M104 64h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H104c-22.091 0-40-17.909-40-40s17.909-40 40-40z m1 408h462c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m0 408h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m666.627-567.931l177.304 177.304c12.497 12.496 12.497 32.758 0 45.254L771.627 711.931c-12.496 12.497-32.758 12.497-45.254 0A32 32 0 0 1 717 689.304V334.696c0-17.673 14.327-32 32-32a32 32 0 0 1 22.627 9.373z" p-id="10509"></path></svg>',
          click() { _this.collapsed = !_this.collapsed },
        },
        '|',
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
      cache: { // 缓存
        enable: false,
      },
      after: () => {

      },
      input: () => {

      },
      esc: () => {

      },
      upload: {
        handler() {
          _this.uploadClipboardToOSS()
        }
      }
    })
    this.getDirList()
  },
  methods: {
    nodeProps: ({ option }) => {
      return {
        onClick() {
          message.info("[LClick] " + option.label);
        },
        onContextmenu(e) { // 右击
          message.info("[RClick] " + option.label);
        }
      };
    },
    getDirList() {
      GetDirList().then((res) => {
        console.log(res)
      })
    },
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
        // _this.contentEditor.insertValue("![url]("+res+")\n", true)
        message.info(res)
      })
    },
    addDirPath() {
      if (this.dirPath !== null && this.dirPath.trim() !== '') {
        console.log(
          this.dirPath
        );
        AddDirPath(this.dirPath).then((res) => {
          let result = JSON.parse(res)
          if (result.code === 200) {
            message.success(result.data)
          } else {
            message.error(result.msg)
          }
        })
      }
    },
  },

}
</script>
<style>
.treeDiv {
  margin-top: 10px;
}
</style>
