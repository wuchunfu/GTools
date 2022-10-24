<template>
  <div class="parentContent" ref="parentContent">
    <div id="vditor" class="vditor" ref="vditor"></div>
    <n-drawer v-model:show="showDirList" :width="502" :placement="placement">
      <n-input-group style="margin-top: 30px; margin-bottom: 10px;">
        <n-input :style="{ width: '100%'}" placeholder="添加文件夹" v-model:value="dirPath" clearable />
        <n-button type="primary" @click="addDirPath">
          添加
        </n-button>
      </n-input-group>
      <div class="dark">
        <el-tree :data="treeData" @node-click="handleNodeClick" :default-expand-all="true" />
      </div>
    </n-drawer>
  </div>

</template>
<script>
import Vditor from "vditor"
import "vditor/dist/index.css"
import { createDiscreteApi } from 'naive-ui'
import mitt from '../utils/event.js'

// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
  ["message", "dialog"]
);

export default {
  props: ['path'],
  setup(props) {
    // const route = useRoute()
    // localStorage.setItem('mdPath', route.query.path)
  },
  data() {
    return {
      app: window.go.gtools.App,
      collapsed: true,
      contentEditor: "",
      dirPath: null,
      treeData: [],
      showDirList: false,
      placement: 'left',  // 抽屉展示位置
      mdPath: null,
      themeType: true,  // true 亮 false 暗
      editorHeight: 0,
    }
  },
  //mounted
  mounted() {
    let _this = this
    this.editorHeight = this.$refs.parentContent.clientHeight
    let themeType = localStorage.getItem("theme")
    let initTheme = 'classic'
    let contentTheme = 'light'
    if (themeType == 0) {
      initTheme = 'dark'
      contentTheme = 'dark'
    }
    this.contentEditor = new Vditor('vditor', {
      height: this.editorHeight, // 默认使用窗口的高度
      width: '100%',
      toolbar: [
        {
          name: 'menu',
          tipPosition: 's',
          tip: '菜单展开',
          className: 'right',
          icon: '<svg t="1665030513743" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10508" width="200" height="200"><path d="M104 64h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H104c-22.091 0-40-17.909-40-40s17.909-40 40-40z m1 408h462c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m0 408h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m666.627-567.931l177.304 177.304c12.497 12.496 12.497 32.758 0 45.254L771.627 711.931c-12.496 12.497-32.758 12.497-45.254 0A32 32 0 0 1 717 689.304V334.696c0-17.673 14.327-32 32-32a32 32 0 0 1 22.627 9.373z" p-id="10509"></path></svg>',
          click() {
            _this.showDirList = !_this.showDirList
          },
        },
        {
          name: 'deldir',
          tipPosition: 's',
          tip: '目录删除',
          className: 'right',
          icon: '<svg t="1665242133951" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="13011" width="200" height="200"><path d="M354.133333 418.133333c-17.066667 0-34.133333 12.8-34.133333 34.133334v341.333333c0 17.066667 12.8 34.133333 34.133333 34.133333s29.866667-17.066667 29.866667-34.133333v-341.333333c0-21.333333-12.8-34.133333-29.866667-34.133334zM512 418.133333c-17.066667 0-34.133333 12.8-34.133333 34.133334v341.333333c0 17.066667 12.8 34.133333 34.133333 34.133333s34.133333-12.8 34.133333-34.133333v-341.333333c0-21.333333-17.066667-34.133333-34.133333-34.133334zM640 452.266667v341.333333c0 17.066667 12.8 34.133333 34.133333 34.133333s34.133333-12.8 34.133334-34.133333v-341.333333c0-17.066667-12.8-34.133333-34.133334-34.133334s-34.133333 12.8-34.133333 34.133334z" p-id="13012"></path><path d="M938.666667 128h-213.333334v-21.333333C725.333333 46.933333 678.4 0 618.666667 0h-213.333334C345.6 0 298.666667 46.933333 298.666667 106.666667V128H85.333333c-25.6 0-42.666667 17.066667-42.666666 42.666667s17.066667 42.666667 42.666666 42.666666h42.666667v704C128 977.066667 174.933333 1024 234.666667 1024h554.666666c59.733333 0 106.666667-46.933333 106.666667-106.666667V213.333333h42.666667c25.6 0 42.666667-17.066667 42.666666-42.666666s-17.066667-42.666667-42.666666-42.666667zM384 106.666667c0-12.8 8.533333-21.333333 21.333333-21.333334h213.333334c12.8 0 21.333333 8.533333 21.333333 21.333334V128H384v-21.333333z m426.666667 810.666666c0 12.8-8.533333 21.333333-21.333334 21.333334h-554.666666c-12.8 0-21.333333-8.533333-21.333334-21.333334V213.333333h597.333334v704z" p-id="13013"></path></svg>',
          click() {
            _this.delMdDir()
          },
        },
        {
          name: 'newmd',
          tipPosition: 's',
          tip: '新建文档',
          className: 'right',
          icon: '<svg t="1666324094016" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5807" width="200" height="200"><path d="M608 384h-320c-19.2 0-32 12.8-32 32s12.8 32 32 32h320c19.2 0 32-12.8 32-32S627.2 384 608 384zM608 576h-320c-19.2 0-32 12.8-32 32s12.8 32 32 32h320c19.2 0 32-12.8 32-32S627.2 576 608 576z" p-id="5808"></path><path d="M800 0h-512C236.8 0 192 44.8 192 96V128h-32C108.8 128 64 172.8 64 224v704c0 51.2 44.8 96 96 96h576c51.2 0 96-44.8 96-96V896h32c51.2 0 96-44.8 96-96v-640C960 70.4 889.6 0 800 0zM768 928c0 19.2-12.8 32-32 32h-576c-19.2 0-32-12.8-32-32v-704c0-19.2 12.8-32 32-32h576c19.2 0 32 12.8 32 32v704z m128-128c0 19.2-12.8 32-32 32H832V224c0-51.2-44.8-96-96-96H256v-32c0-19.2 12.8-32 32-32h512c51.2 0 96 44.8 96 96v640z" p-id="5809"></path></svg>',
          click() {
            _this.setValue("")
            localStorage.removeItem("mdPath")
          },
        },
        {
          name: 'save',
          tipPosition: 's',
          tip: '保存',
          className: 'right',
          icon: '<svg t="1665237693513" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4108" width="200" height="200"><path d="M938.666667 1024H85.333333a85.333333 85.333333 0 0 1-85.333333-85.333333V85.333333a85.333333 85.333333 0 0 1 85.333333-85.333333h640v0.256c12.16-0.768 24.106667 3.626667 32.853334 12.074667L1010.346667 264.533333c8.917333 9.002667 13.44 21.504 12.373333 34.133334H1024v640a85.333333 85.333333 0 0 1-85.333333 85.333333zM298.666667 938.666667h426.666666v-298.666667H298.666667v298.666667z m341.333333-853.333334H384v256h256V85.333333z m298.666667 233.685334l-213.333334-213.333334V341.333333a85.333333 85.333333 0 0 1-85.333333 85.333334H384a85.333333 85.333333 0 0 1-85.333333-85.333334V85.333333H85.333333v853.333334h128v-298.666667a85.333333 85.333333 0 0 1 85.333334-85.333333h426.666666a85.333333 85.333333 0 0 1 85.333334 85.333333v298.666667h128V319.018667z" p-id="4109"></path></svg>',
          click() {
            _this.saveMdContent()
          },
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
        'table',
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
        'undo',
        'redo',
        '|',
        'edit-mode',
        'fullscreen',
        'outline',
        // {
        //   name: 'changetheme',
        //   tipPosition: 's',
        //   tip: '亮暗模式',
        //   className: 'right',
        //   icon: '<svg t="1665281516966" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6580" width="200" height="200"><path d="M512 0h-0.7C298.7 0.4 126.2 174.8 128 387.4c1 112.1 50 212.8 127.5 282.4 51 45.8 80.5 110.8 80.5 179.3v57c0 65.1 52.8 117.9 117.9 117.9h116.2c65.1 0 117.9-52.8 117.9-117.9v-57c0-68.6 29.6-133.6 80.6-179.5C846.8 599.3 896 497.4 896 384 896 171.9 724.1 0 512 0z m213.9 622C661.1 680.3 624 763 624 849.1V896c0 35.3-28.7 64-64 64h-96c-35.3 0-64-28.7-64-64v-46.9c0-86.1-37.1-168.9-101.7-227-32.8-29.4-58.7-64.6-77.1-104.4-19-41.2-28.8-85.2-29.2-130.9-0.4-43.3 7.8-85.4 24.3-125.1 16-38.4 39-72.9 68.4-102.7 29.4-29.7 63.7-53.1 101.9-69.5 39.5-16.9 81.5-25.5 124.8-25.6h0.6c43.2 0 85.1 8.5 124.5 25.1 38.1 16.1 72.3 39.2 101.7 68.6 29.4 29.4 52.5 63.6 68.6 101.7C823.6 298.9 832 340.8 832 384c0 46.2-9.6 90.8-28.6 132.5C785 556.8 759 592.3 725.9 622z" p-id="6581"></path><path d="M516.7 624l19.4-173.3H460L627.3 208l-19.4 173.3H684L516.7 624zM576 912H448c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h128c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z m0-64H448c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h128c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z m16-64H432c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h160c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z" p-id="6582"></path></svg>',
        //   click() {
        //     _this.themeType = !_this.themeType
        //     if (_this.themeType) {
        //       _this.setLightTheme()
        //     } else {
        //       _this.setDarkTheme()
        //     }
        //   },
        // },
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
      mode: 'wysiwyg',
      outline: {
        enable: false,
        position: 'left'
      },
      preview: {
        hljs: {
          enable: true,
          style: 'vim',
          lineNumber: true
        },
        theme: {
          current: contentTheme
        }
      },
      resize: {
        enable: true,
        after(h) {
          console.log(h);
        }
      },
      fullscreen: {
        index: 150
      },
      theme: initTheme,
      cache: { // 缓存
        enable: true,
      },
      after: () => {

      },
      input: () => {

      },
      esc: () => {

      },
      blur: (content) => {
        let path = localStorage.getItem("mdPath")
        if (path !== null && path.trim() !== '' && !path.startsWith("menu")) {
          this.app.SaveMdContent(path, content).then((res) => {
            if (res.code != 200) {
              message.error(result.msg)
            }
          })
        }
      },
      upload: {
        handler() {
          _this.uploadClipboard()
        }
      },
      counter: {
        enable: true,
      }
    })
    this.getDirList()
    mitt.on("theme", (val) => {
      if (val == 0) {
        _this.setDarkTheme()
      } else {
        _this.setLightTheme()
      }
    })
  },
  methods: {
    handleNodeClick(data) {
      this.mdPath = data.key
      localStorage.setItem("mdPath", this.mdPath)
      if (this.mdPath.startsWith("menu")) {
        this.setValue("")
      } else {
        this.getMdContent(data.key)
      }
    },
    delMdDir() {
      let path = localStorage.getItem("mdPath")
      if (path.startsWith("menu")) {
        dialog.warning({
          title: '警告',
          content: '你确定删除目录: ' + this.mdPath.substring(5),
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: () => {
            this.app.DelMdDir(this.mdPath.substring(5)).then((res) => {
              let result = JSON.parse(res)
              if (result.code == 200) {
                message.success("删除成功")
                this.getDirList()
                localStorage.removeItem("mdPath")
              } else {
                message.error(result.msg)
              }
            })
          },
          onNegativeClick: () => {

          }
        })
      }
    },
    getMdContent(path) {
      if (!path.startsWith("menu")) {
        this.app.GetMdContent(path).then((res) => {
          if (res.code == 200) {
            this.setValue(res.data)
          } else {
            message.error(res.msg)
          }
        })
      }
    },
    saveMdContent() {
      let _this = this
      let mdPath = ""
      let path = localStorage.getItem("mdPath")
      let content = this.getValue()
      if (path !== null && path.trim() != '' && !path.startsWith("menu")) {
        this.app.SaveMdContent(path, content).then((res) => {
          let result = JSON.parse(res)
          if (result.code == 200) {
            message.success("保存成功！")
          } else {
            message.error(result.msg)
          }
        })
      } else {
        this.app.OpenMdSaveFileWindow().then(res => {
          if (res.code != 200) {
            message.error(res.msg)
            return
          }
          mdPath = res.data
          if (res.data == '') {
            message.info("已取消保存")
            return
          }
          let content = _this.getValue()
          console.log(content);
          _this.app.NewMd(mdPath, content).then(res => {
            if (res.code == 200) {
              message.success("保存成功")
              localStorage.setItem("mdPath", mdPath)
            } else {
              message.error("保存失败")
            }
          })
        })
      }
    },
    getDirList() {
      let _this = this
      this.app.GetDirList().then((res) => {
        if (res.code == 200) {
          let data = JSON.parse(res.data)
          _this.treeData = []
          for (let item in data) {
            let newItem = {
              label: item,
              key: "menu-" + item,
              children: data[item].length == 0 ? [] : _this.getChildren(data[item])
            }
            _this.treeData.push(newItem)
          }
        } else {
          message.error(res.msg)
        }
      })
    },
    getChildren(list) {
      let data = []
      for (let item of list) {
        let newItem = {
          label: item.fname,
          key: item.fpath,
        }
        data.push(newItem)
      }
      return data
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
    setDarkTheme() {
      return this.contentEditor.setTheme('dark', 'dark');
    },
    setLightTheme() {
      return this.contentEditor.setTheme('classic', 'light');
    },
    insertValue(str, render) {
      return this.contentEditor.insertValue(str, render);
    },
    toPreview() {
      var evt = document.createEvent('Event');
      evt.initEvent('click', true, true);
      this.contentEditor.vditor.toolbar.elements.preview.firstElementChild.dispatchEvent(evt);
    },
    uploadClipboard() {
      let _this = this
      this.app.UploadScreenshot().then((res) => {
        if (res.code == 200) {
          _this.insertValue(res.data)
        } else {
          message.error(res.msg)
        }
      })
    },
    addDirPath() {
      if (this.dirPath !== null && this.dirPath.trim() !== '') {
        this.app.AddDirPath(this.dirPath).then((res) => {
          let result = JSON.parse(res)
          if (result.code === 200) {
            this.getDirList()
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

.parentContent {
  width: 100%;
  height: 100%;
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
}
</style>
