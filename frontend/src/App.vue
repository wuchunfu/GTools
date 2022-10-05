<script setup>
import {h, ref} from "vue";
import {RouterLink} from "vue-router";
import {NIcon} from "naive-ui";
import {
  Code as CodeIcon, Home as HomeIcon,
  LogoMarkdown as MditorIcon
} from '@vicons/ionicons5'

import {lightTheme, darkTheme} from 'naive-ui'

const renderIcon = (icon) => {
  return () => h(NIcon, null, {default: () => h(icon)});
}

// 菜单栏默认折叠
const collapsed = ref(true);

const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: "/",
          }
        },
        {default: () => "主页"}
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
        {default: () => "快捷终端"}
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
        {default: () => "Mditor"}
    ),
    key: "go-back-mditor",
    icon: renderIcon(MditorIcon)
  },
]
</script>

<template>
  <n-config-provider :theme="lightTheme">
    <n-space vertical size="large">
      <n-layout has-sider position="absolute">
        <n-layout-sider bordered collapse-mode="width" :collapsed-width="80" :width="200" :collapsed="collapsed"
                        show-trigger @collapse="collapsed = true" @expand="collapsed = false">
          <n-menu :options="menuOptions" :collapsed-width="64" :collapsed-icon-size="22"/>
          <!-- <n-switch v-model:value="switchTheme" @update:value="changeTheme()" /> -->
        </n-layout-sider>
        <n-layout-content>
          <router-view/>
        </n-layout-content>
      </n-layout>
    </n-space>
  </n-config-provider>
</template>

<style>

</style>
