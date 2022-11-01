<template>
  <div class="parentContent" ref="parentContent">
    <div id="vditor" class="vditor" ref="vditor"></div>
    <n-drawer v-model:show="showDirList" :width="502" :placement="placement" style="--wails-draggable:drag">
      <n-input-group style="display: flex; justify-content: center; margin: 20px 0;">
        <n-button type="primary" ghost round @click="addDirPath">
          <template #icon>
            <n-icon>
              <folder-icon />
            </n-icon>
          </template>
          添加文件夹
        </n-button>
        <n-button type="primary" ghost round @click="addFilePath">
          <template #icon>
            <n-icon>
              <file-icon />
            </n-icon>
          </template>
          添加文件
        </n-button>
      </n-input-group>
      <div style="padding: 0 20px;">
        <n-collapse accordion @item-header-click="handleItem">
          <n-collapse-item :title="item.path" :name="item.path" v-for="item, index in treeData">
            <n-list hoverable clickable>
              <n-list-item v-for="item in mdFils" @click="getMdContent(item.fpath)">
                <n-thing content-style="margin-left: 10px;">
                  {{ item.fname }}
                </n-thing>
              </n-list-item>
            </n-list>
            <template #header-extra>
              <n-button circle @click="delMdDir(item)">
                <template #icon>
                  <n-icon>
                    <del-icon />
                  </n-icon>
                </template>
              </n-button>
            </template>
          </n-collapse-item>
        </n-collapse>
        <n-divider dashed>
          文件
        </n-divider>
        <n-list hoverable clickable>
          <n-list-item v-for="item in mdFileList" @click="getMdContent(item.path)">
            <n-thing content-style="margin-left: 10px; display: flex; align-items:center; justify-content: left;">
              {{ item.fname }}
              <n-button circle @click="delMdDir(item)" style="margin-left: auto;">
                <template #icon>
                  <n-icon>
                    <del-icon />
                  </n-icon>
                </template>
              </n-button>
            </n-thing>
          </n-list-item>
        </n-list>
      </div>
    </n-drawer>
  </div>

</template>
<script>
import Vditor from "vditor"
import "vditor/dist/index.css"
import { createDiscreteApi } from 'naive-ui'
import mitt from '../utils/event.js'
import { FolderOpen as FolderIcon, DocumentText as FileIcon } from "@vicons/ionicons5";
import { DeleteOutlineOutlined as DelIcon } from '@vicons/material'

// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
  ["message", "dialog"]
);

export default {
  components: {
    FolderIcon,
    FileIcon,
    DelIcon
  },
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
      treeData: [],
      mdFils: [], // 文件夹下的md文档
      mdFileList: [],
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
          name: 'newmd',
          tipPosition: 's',
          tip: '新建文档',
          className: 'right',
          icon: '<svg t="1666974492027" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7368" width="200" height="200"><path d="M658.285714 0l292.571429 292.571429h-73.142857L658.285714 73.142857V0z m292.571429 292.571429v731.428571H73.142857V0h585.142857v292.571429h292.571429z m-73.142857 73.142857H585.142857V73.142857H146.285714v877.714286h731.428572V365.714286zM146.285714 804.571429h146.285715v146.285714h73.142857V804.571429h146.285714v-73.142858H365.714286V585.142857h-73.142857v146.285714H146.285714v73.142858z" p-id="7369"></path></svg>',
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
        if (path !== null && path.trim() !== '') {
          this.app.SaveMdContent(path, content).then((res) => {
            if (res.code != 200) {
              message.error(res.msg)
              localStorage.removeItem("mdPath")
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
    handleItem(data) {
      this.app.GetMdFileList(data.name).then(res => {
        if (res.code == 200) {
          this.mdFils = res.data
        } else {
          message.error(res.msg)
        }
      })
    },
    delMdDir(item) {
      dialog.warning({
        title: '警告（非真实删除）',
        content: '你确定删除: ' + item.path,
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          this.app.DelMdDir(item).then((res) => {
            if (res.code == 200) {
              message.success("删除成功")
              this.treeData = res.data.dir
              this.mdFileList = res.data.file
            } else {
              message.error(result.msg)
            }
          })
        },
        onNegativeClick: () => {

        }
      })
    },
    getMdContent(path) {
      // 判断当前文件是否保存
      if (this.getValue != "" && localStorage.getItem("mdPath") != null) {
        localStorage.setItem("mdPath", path)
        this.app.GetMdContent(path).then((res) => {
          if (res.code == 200) {
            this.setValue(res.data)
          } else {
            message.error(res.msg)
          }
        })
      } else {
        dialog.warning({
          title: '警告',
          content: '当前文档未保存，是否保存',
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: () => {
            this.saveMdContent()
          },
          onNegativeClick: () => {
            localStorage.setItem("mdPath", path)
            this.app.GetMdContent(path).then((res) => {
              if (res.code == 200) {
                this.setValue(res.data)
              } else {
                message.error(res.msg)
              }
            })
          }
        })
      }

    },
    saveMdContent() {
      let _this = this
      let mdPath = ""
      let path = localStorage.getItem("mdPath")
      let content = this.getValue()
      if (path !== null && path.trim() !== '') {
        this.app.SaveMdContent(path, content).then((res) => {
          if (res.code == 200) {
            message.success("保存成功！")
          } else {
            message.error(res.msg)
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
      this.app.GetMdDirList().then(res => {
        if (res.code === 200) {
          this.treeData = res.data.dir
          this.mdFileList = res.data.file
        } else {
          message.error(res.msg)
        }
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
      this.app.OpenMdFolderWindow().then(res => {
        if (res.code == 200) {
          if (res.data == null) {
            message.error("文件夹选择异常！")
            return
          }
          if (res.data == '') { // 取消选择
            return
          }
          this.app.AddMdDirPath({ path: res.data, type: 0 }).then((res) => {
            if (res.code === 200) {
              this.treeData = res.data.dir
              this.mdFileList = res.data.file
            } else {
              message.error(res.msg)
            }
          })
        }
      })
    },
    addFilePath() {
      this.app.AddMdFile().then(res => {
        if (res.code == 200) {
          this.treeData = res.data.dir
          this.mdFileList = res.data.file
        } else {
          message.error(res.msg)
        }
      })
    }
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
