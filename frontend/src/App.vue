<script>
import { h, ref } from "vue";
import { RouterLink } from "vue-router";
import { NIcon, NConfigProvider } from "naive-ui";
import {
  Code as CodeIcon, Home as HomeIcon,
  LogoMarkdown as MditorIcon,
  SettingsSharp as SettingIcon
} from '@vicons/ionicons5'

import { darkTheme } from 'naive-ui'

const renderIcon = (icon) => {
  return () => h(NIcon, null, { default: () => h(icon) });
}

export default {
  data() {
    return {
      collapsed: false,
      switchTheme: false,
      myTheme: null,
      menuOptions: [
        {
          label: () => h(
            RouterLink,
            {
              to: {
                path: "/",
              }
            },
            { default: () => "主页" }
          ),
          key: "go-back-home",
          icon: renderIcon(HomeIcon)
        },
        {
          label: () => h(
            RouterLink,
            {
              to: {
                name: 'cmd',
                path: "/cmd",
              }
            },
            { default: () => "快捷指令" }
          ),
          key: "go-back-cmd",
          icon: renderIcon(CodeIcon)
        },
        {
          label: () => h(
            RouterLink,
            {
              to: {
                name: 'mditor',
                path: "/mditor",
              }
            },
            { default: () => "Mditor" }
          ),
          key: "go-back-mditor",
          icon: renderIcon(MditorIcon)
        },
        {
          label: () => h(
            RouterLink,
            {
              to: {
                name: 'setting',
                path: "/setting",
              }
            },
            { default: () => "设置" }
          ),
          key: "go-back-setting",
          icon: renderIcon(SettingIcon)
        },
      ],
      railStyle: ({
        focused,
        checked
      }) => {
        const style = {};
        if (checked) {
          style.background = "#4B9D5F";
          if (focused) {
            style.boxShadow = "0 0 0 2px #d0305040";
          }
        } else {
          style.background = "#000000";
          if (focused) {
            style.boxShadow = "0 0 0 2px #2080f040";
          }
        }
        return style;
      }
    }
  },
  methods: {
    changeTheme() {
      console.log(this.switchTheme);
      if (this.switchTheme) {
        this.myTheme = darkTheme
      } else {
        this.myTheme = null
      }
    }
  }
}

</script>

<template>
  <n-config-provider :theme="myTheme">
    <n-space vertical size="large">
      <n-layout has-sider position="absolute">
          <n-layout-sider bordered collapse-mode="width" :collapsed-width="80" :width="150" :collapsed="collapsed"
            show-trigger @collapse="collapsed = true" @expand="collapsed = false">
            <n-menu :options="menuOptions" :collapsed-width="64" :collapsed-icon-size="22" style="margin-top: 40px;" />
            <div class="switchBtnPar">
              <n-divider />
              <n-switch :rail-style="railStyle" v-model:value="switchTheme" @update:value="changeTheme()"
                class="switchBtn">
                <template #checked>
                  亮
                </template>
                <template #unchecked>
                  暗
                </template>
              </n-switch>
            </div>
          </n-layout-sider>
          <n-layout-content>
            <router-view/>
          </n-layout-content>
        </n-layout>
    </n-space>
  </n-config-provider>
</template>

<style>
.switchBtnPar {
  position: relative;
}

.switchBtn {
  position: absolute;
  left: 50%;
  transform: translate(-50%);
}
</style>
