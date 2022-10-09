<template>
  <div id="vditor" class="vditor"></div>
  <n-drawer v-model:show="showDirList" :width="502" :placement="placement">
    <n-input-group style="margin-top: 30px; margin-bottom: 10px;">
      <n-input :style="{ width: '100%'}" placeholder="添加文件夹" v-model:value="dirPath" clearable />
      <n-button type="primary" @click="addDirPath">
        添加
      </n-button>
    </n-input-group>
    <div>
      <el-tree :data="treeData" @node-click="handleNodeClick" :default-expand-all="true" />
    </div>
  </n-drawer>
</template>
<script>
import Vditor from "vditor"
import "vditor/dist/index.css"
import { UploadImgByPicgo, AddDirPath, GetDirList, GetMdContent, SaveMdContent, DelMdDir } from "../../wailsjs/go/main/App"
import { createDiscreteApi, NButton } from 'naive-ui'
import { h } from "vue";

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
      collapsed: true,
      contentEditor: "",
      dirPath: null,
      treeData: [],
      showDirList: false,
      placement: 'left',  // 抽屉展示位置
      mdPath: null,
      themeType: true,  // true 亮 false 暗
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
          name: 'changetheme',
          tipPosition: 's',
          tip: '亮暗模式',
          className: 'right',
          icon: '<svg t="1665281516966" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6580" width="200" height="200"><path d="M512 0h-0.7C298.7 0.4 126.2 174.8 128 387.4c1 112.1 50 212.8 127.5 282.4 51 45.8 80.5 110.8 80.5 179.3v57c0 65.1 52.8 117.9 117.9 117.9h116.2c65.1 0 117.9-52.8 117.9-117.9v-57c0-68.6 29.6-133.6 80.6-179.5C846.8 599.3 896 497.4 896 384 896 171.9 724.1 0 512 0z m213.9 622C661.1 680.3 624 763 624 849.1V896c0 35.3-28.7 64-64 64h-96c-35.3 0-64-28.7-64-64v-46.9c0-86.1-37.1-168.9-101.7-227-32.8-29.4-58.7-64.6-77.1-104.4-19-41.2-28.8-85.2-29.2-130.9-0.4-43.3 7.8-85.4 24.3-125.1 16-38.4 39-72.9 68.4-102.7 29.4-29.7 63.7-53.1 101.9-69.5 39.5-16.9 81.5-25.5 124.8-25.6h0.6c43.2 0 85.1 8.5 124.5 25.1 38.1 16.1 72.3 39.2 101.7 68.6 29.4 29.4 52.5 63.6 68.6 101.7C823.6 298.9 832 340.8 832 384c0 46.2-9.6 90.8-28.6 132.5C785 556.8 759 592.3 725.9 622z" p-id="6581"></path><path d="M516.7 624l19.4-173.3H460L627.3 208l-19.4 173.3H684L516.7 624zM576 912H448c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h128c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z m0-64H448c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h128c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z m16-64H432c-4.4 0-8.4-1.8-11.3-4.7-2.9-2.9-4.7-6.9-4.7-11.3 0-8.8 7.2-16 16-16h160c4.4 0 8.4 1.8 11.3 4.7 2.9 2.9 4.7 6.9 4.7 11.3 0 8.8-7.2 16-16 16z" p-id="6582"></path></svg>',
          click() {
            _this.themeType = !_this.themeType
            if (_this.themeType) {
              _this.setLightTheme()
            } else {
              _this.setDarkTheme()
            }
          },
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
      theme: '',
      fullscreen: {
        index: 150
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
      blur: (content) => {
        let path = _this.mdPath
        if (path !== null && path.trim() !== '' && !path.startsWith("menu")) {
          SaveMdContent(path, content).then((res) => {
            let result = JSON.parse(res)
            if (result.code != 200) {
              message.error(result.msg)
            }
          })
        }
      },
      upload: {
        handler() {
          _this.uploadClipboardToOSS()
        }
      },
      counter: {
        enable: true,
      }
    })
    this.getDirList()
  },
  methods: {
    handleNodeClick(data) {
      this.mdPath = data.key
      this.getMdContent(data.key)
    },
    delMdDir() {
      if (this.mdPath.startsWith("menu")) {
        dialog.warning({
          title: '警告',
          content: '你确定删除目录: ' + this.mdPath.substring(5),
          positiveText: '确定',
          negativeText: '取消',
          onPositiveClick: () => {
            DelMdDir(this.mdPath.substring(5)).then((res) => {
              let result = JSON.parse(res)
              if (result.code == 200) {
                message.success("删除成功")
                this.getDirList()
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
        GetMdContent(path).then((res) => {
          let result = JSON.parse(res)
          if (result.code == 200) {
            this.setValue(result.data)
          } else if (result.code == 500) {
            message.error(result.msg)
          } else { }
        })
      }
    },
    saveMdContent() {
      let _this = this
      let path = _this.mdPath
      let content = _this.getValue()
      if (path !== null && path.trim() !== '' && !path.startsWith("menu")) {
        SaveMdContent(path, content).then((res) => {
          let result = JSON.parse(res)
          if (result.code == 200) {
            message.success("保存成功！")
          } else {
            message.error(result.msg)
          }
        })
      }
    },
    getDirList() {
      let _this = this
      GetDirList().then((res) => {
        let result = JSON.parse(res)
        if (result.code == 200) {
          let data = JSON.parse(result.data)
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
          message.error(result.msg)
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
    uploadClipboardToOSS() {
      let _this = this
      UploadImgByPicgo().then((res) => {
        message.info(res)
      })
    },
    addDirPath() {
      if (this.dirPath !== null && this.dirPath.trim() !== '') {
        AddDirPath(this.dirPath).then((res) => {
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
</style>
