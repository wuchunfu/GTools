<template>
  <div class="parentContent" ref="parentContent">
    <n-layout has-sider sider-placement="right">
      <n-layout-content content-style="overflow: hidden">
        <div id="vditor" class="vditor" ref="vditor"></div>
      </n-layout-content>
      <n-layout-sider collapse-mode="transform" :collapsed-width="0" :width="300" :native-scrollbar="false"
        content-style="" show-trigger="bar" :default-collapsed="true">
        <n-space style="margin-top: 20px; margin-left: 30px">
          <n-button type="success" dashed @click="contentTrans()">
            立即翻译
          </n-button>
          <n-button type="success" class="copy" @click="translationCopy()" :data-clipboard-text="translation">
            复制译文
          </n-button>
        </n-space>
        <n-divider />
        <n-card :bordered="false" title="原文">
          <n-input placeholder="原文内容..." type="textarea" size="small" :autosize="{
            minRows: 5,
            maxRows: 10
          }" v-model:value="original" />
        </n-card>
        <n-card :bordered="false" title="译文">
          <n-spin :show="!transOk">
            <n-input placeholder="译文内容..." type="textarea" size="small" :autosize="{
              minRows: 5,
              maxRows: 10
            }" v-model:value="translation" />
          </n-spin>

        </n-card>
      </n-layout-sider>
    </n-layout>
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
<script setup>
import { ref, onMounted } from "vue";
import Vditor from "vditor"
import "vditor/dist/index.css"
import { createDiscreteApi } from 'naive-ui'
import mitt from '../utils/event.js'
import { FolderOpen as FolderIcon, DocumentText as FileIcon } from "@vicons/ionicons5";
import { DeleteOutlineOutlined as DelIcon } from '@vicons/material'
import Clipboard from 'clipboard'


onMounted(() => {
  editorHeight.value = document.documentElement.clientHeight;
  let themeType = localStorage.getItem("theme")
  let initTheme = 'classic'
  let contentTheme = 'light'
  if (themeType == 0) {
    initTheme = 'dark'
    contentTheme = 'dark'
  }
  contentEditor.value = new Vditor('vditor', {
    height: editorHeight.value, // 默认使用窗口的高度
    width: '100%',
    // icon: "material",
    cache: {
      enable: true
    },
    toolbar: [
      {
        name: 'menu',
        tipPosition: 's',
        tip: '菜单展开',
        className: 'right',
        icon: '<svg t="1665030513743" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10508" width="200" height="200"><path d="M104 64h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H104c-22.091 0-40-17.909-40-40s17.909-40 40-40z m1 408h462c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m0 408h815c22.091 0 40 17.909 40 40s-17.909 40-40 40H105c-22.091 0-40-17.909-40-40s17.909-40 40-40z m666.627-567.931l177.304 177.304c12.497 12.496 12.497 32.758 0 45.254L771.627 711.931c-12.496 12.497-32.758 12.497-45.254 0A32 32 0 0 1 717 689.304V334.696c0-17.673 14.327-32 32-32a32 32 0 0 1 22.627 9.373z" p-id="10509"></path></svg>',
        click() {
          showDirList.value = !showDirList.value
        },
      },
      {
        name: 'newmd',
        tipPosition: 's',
        tip: '新建文档',
        className: 'right',
        icon: '<svg t="1666974492027" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7368" width="200" height="200"><path d="M658.285714 0l292.571429 292.571429h-73.142857L658.285714 73.142857V0z m292.571429 292.571429v731.428571H73.142857V0h585.142857v292.571429h292.571429z m-73.142857 73.142857H585.142857V73.142857H146.285714v877.714286h731.428572V365.714286zM146.285714 804.571429h146.285715v146.285714h73.142857V804.571429h146.285714v-73.142858H365.714286V585.142857h-73.142857v146.285714H146.285714v73.142858z" p-id="7369"></path></svg>',
        click() {
          setValue("")
          clearCache()
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
          saveMdContent()
        },
      },
      {
        name: 'exphtml',
        tipPosition: 's',
        tip: '导出HTML',
        className: 'right',
        icon: '<svg t="1669430938977" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3499" width="200" height="200"><path d="M921.6 1024H102.4c-56.32 0-102.4-46.08-102.4-102.4V102.4C0 46.08 46.08 0 102.4 0h256c30.72 0 51.2 20.48 51.2 51.2s-20.48 51.2-51.2 51.2H153.6c-30.72 0-51.2 20.48-51.2 51.2v716.8c0 30.72 20.48 51.2 51.2 51.2h716.8c30.72 0 51.2-20.48 51.2-51.2v-204.8c0-30.72 20.48-51.2 51.2-51.2s51.2 20.48 51.2 51.2v256c0 56.32-46.08 102.4-102.4 102.4z" p-id="3500"></path><path d="M552.96 399.36L916.48 35.84c20.48-20.48 51.2-20.48 71.68 0 20.48 20.48 20.48 51.2 0 71.68l-363.52 363.52c-20.48 20.48-51.2 20.48-71.68 0-20.48-15.36-20.48-51.2 0-71.68zM768 0h204.8c30.72 0 51.2 20.48 51.2 51.2s-20.48 51.2-51.2 51.2h-204.8c-30.72 0-51.2-20.48-51.2-51.2s20.48-51.2 51.2-51.2z m204.8 0c30.72 0 51.2 20.48 51.2 51.2v204.8c0 30.72-20.48 51.2-51.2 51.2s-51.2-20.48-51.2-51.2V51.2c0-30.72 20.48-51.2 51.2-51.2z" p-id="3501"></path></svg>',
        click() {
          exportHTML()
        },
      },
      '|',
      // 'emoji',
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
          'preview',
          // 'content-theme',
          // 'code-theme',
          // 'info',
          // 'help',
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
      },
      math: {
        engine: 'KaTeX',
      },
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
    select(md) {
      original.value = md
      translation.value = ""
    },
    input: () => {

    },
    esc: () => {

    },
    // cdn: location.origin,
    blur: (content) => {
      let path = localStorage.getItem("mdPath")
      if (path !== null && path.trim() !== '') {
        app.value.SaveMdContent(path, content).then((res) => {
          if (res.code != 200) {
            message.error(res.msg)
            localStorage.removeItem("mdPath")
          }
        })
      }
    },
    upload: {
      handler() {
        uploadClipboard()
      }
    },
    counter: {
      enable: true,
    }
  })
  getDirList()
  mitt.on("theme", (val) => {
    if (val == 0) {
      setDarkTheme()
    } else {
      setLightTheme()
    }
  })
})

// 脱离上下文的 API 引入消息提示框
const { message, dialog } = createDiscreteApi(
  ["message", "dialog"]
);
const app = ref(window.go.gtools.App)
const collapsed = ref(true)
const contentEditor = ref("")
const treeData = ref([])
const mdFils = ref([]) // 文件夹下的md文档
const mdFileList = ref([])
const showDirList = ref(false)
const placement = ref('left') // 抽屉展示位置
const mdPath = ref(null)
const themeType = ref(true)  // true 亮 false 暗
const editorHeight = ref(0)
const translation = ref("")
const original = ref("")
const transOk = ref(true)

const handleItem = (data) => {
  app.value.GetMdFileList(data.name).then(res => {
    if (res.code == 200) {
      mdFils.value = res.data
    } else {
      message.error(res.msg)
    }
  })
}
const delMdDir = (item) => {
  dialog.warning({
    title: '警告（非真实删除）',
    content: '你确定删除: ' + item.path,
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      app.value.DelMdDir(item).then((res) => {
        if (res.code == 200) {
          message.success("删除成功")
          treeData.value = res.data.dir
          mdFileList.value = res.data.file
        } else {
          message.error(result.msg)
        }
      })
    },
    onNegativeClick: () => {

    }
  })
}
const getMdContent = (path) => {
  // 判断当前文件是否保存
  if (localStorage.getItem("mdPath") != null) {
    localStorage.setItem("mdPath", path)
    getContent(path)
  } else { // 未保存
    if (getValue().trim() != '') { // 有内容
      dialog.warning({
        title: '警告',
        content: '当前文档未保存，是否保存',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          saveMdContent()
        },
        onNegativeClick: () => {
          localStorage.setItem("mdPath", path)
          getContent(path)
        }
      })
    } else { // 无内容
      localStorage.setItem("mdPath", path)
      getContent(path)
    }
  }
}
const getContent = (path) => {
  app.value.GetMdContent(path).then((res) => {
    if (res.code == 200) {
      setValue(res.data)
    } else {
      message.error(res.msg)
    }
  })
}
const saveMdContent = () => {
  let mdPath = ""
  let path = localStorage.getItem("mdPath")
  let content = getValue()
  if (path !== null && path.trim() !== '') {
    app.value.SaveMdContent(path, content).then((res) => {
      if (res.code == 200) {
        message.success("保存成功！")
      } else {
        message.error(res.msg)
      }
    })
  } else {
    app.value.OpenMdSaveFileWindow().then(res => {
      if (res.code != 200) {
        message.error(res.msg)
        return
      }
      mdPath = res.data
      if (res.data == '') {
        message.info("已取消保存")
        return
      }
      let content = getValue()
      app.value.NewMd(mdPath, content).then(res => {
        if (res.code == 200) {
          message.success("保存成功")
          localStorage.setItem("mdPath", mdPath)
          treeData.value = res.data.dir
          mdFileList.value = res.data.file
        } else {
          message.error("保存失败")
        }
      })
    })
  }
}
const getDirList = () => {
  app.value.GetMdDirList().then(res => {
    if (res.code === 200) {
      treeData.value = res.data.dir
      mdFileList.value = res.data.file
    } else {
      message.error(res.msg)
    }
  })
}
const getValue = () => {
  return contentEditor.value.getValue();     //获取 Markdown 内容
}
const getHTML = () => {
  return contentEditor.value.getHTML();      //获取 HTML 内容
}
const setValue = (value) => {
  return contentEditor.value.setValue(value);     //设置 Markdown 内容
}
const clearCache = () => {
  return contentEditor.value.clearCache(); // 清理编辑器缓存
}
const disabled = () => {
  return contentEditor.value.disabled();     //设置 只读
}
const setDarkTheme = () => {
  return contentEditor.value.setTheme('dark', 'dark');
}
const setLightTheme = () => {
  return contentEditor.value.setTheme('classic', 'light');
}
const insertValue = (str, render) => {
  return contentEditor.value.insertValue(str, render);
}
const toPreview = () => {
  var evt = document.createEvent('Event');
  evt.initEvent('click', true, true);
  contentEditor.value.vditor.toolbar.elements.preview.firstElementChild.dispatchEvent(evt);
}
const uploadClipboard = () => {
  app.value.UploadScreenshot().then((res) => {
    if (res.code == 200) {
      insertValue(res.data)
    } else {
      message.error(res.msg)
    }
  })
}
function exportHTML() {
  app.value.OpenHtmlSaveWindow().then(res => {
    if (res.code == 200) {
      if (res.data == null) {
        message.error("文件选择异常！")
        return
      }
      if (res.data == '') { // 取消选择
        return
      }
      let content = getHTML()
      app.value.ExportHTML(res.data, content).then((res) => {
        if (res.code != 200) {
          message.error(res.msg)
        } else {
          message.success("导出成功")
        }
      })
    }
  })
}
const addDirPath = () => {
  app.value.OpenMdFolderWindow().then(res => {
    if (res.code == 200) {
      if (res.data == null) {
        message.error("文件夹选择异常！")
        return
      }
      if (res.data == '') { // 取消选择
        return
      }
      app.value.AddMdDirPath({ path: res.data, type: 0 }).then((res) => {
        if (res.code === 200) {
          treeData.value = res.data.dir
          mdFileList.value = res.data.file
        } else {
          message.error(res.msg)
        }
      })
    }
  })
}
const addFilePath = () => {
  app.value.AddMdFile().then(res => {
    if (res.code == 200) {
      treeData.value = res.data.dir
      mdFileList.value = res.data.file
    } else {
      message.error(res.msg)
    }
  })
}
const contentTrans = () => {
  if (original.value == "") {
    return
  }
  if (!transOk.value) {
    message.error("点击过于频繁")
    return
  }
  transOk.value = false
  app.value.BaiDuTrans(original.value).then(res => {
    if (res.code == 200) {
      translation.value = res.data
    } else {
      message.error(res.msg)
    }
    transOk.value = true
  })
}
const translationCopy = () => {
  if (translation.value == "") {
    return
  }
  let clipboard = new Clipboard('.copy')
  clipboard.on('success', (e) => {
    message.success('复制成功')
    clipboard.destroy()
  })
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
