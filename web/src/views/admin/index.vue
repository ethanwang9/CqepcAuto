<script lang="ts" setup>
import {onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import usePiniaApp from "@/pinia/modules/app";

// pinia
const piniaApp = usePiniaApp()

// 路由
const router = useRouter()
// 后台路由
const routerAdmin = router.getRoutes().filter(v => v.path === "/admin")[0].children
// 当前路由
const routerCurrent = router.currentRoute.value

// 是否折叠菜单
// false - 展开 | true - 折叠
const isCollapse = ref(false)

// 切换折叠菜单状态
// 屏幕 <768px 的窗口打开侧边栏菜单
const changeCollapse = () => {
  if (windowWidth.value < 768) {
    openSideMenu()
  } else {
    isCollapse.value = !isCollapse.value
  }
}

// 刷新页面
const refresh = () => window.location.reload()

// 打开侧边设置窗口
const isOpenSideSetting = ref(false)
const openSideSetting = () => {
  isOpenSideSetting.value = true
}

// 打开侧边导航窗口
const isOpenSideMenu = ref(false)
const openSideMenu = () => {
  isOpenSideMenu.value = true
}

// 打开侧边栏选中item后自动关闭
const openSideMenuSelectClose = () => {
  setTimeout(() => isOpenSideMenu.value = false, 300)
}

// 监控屏幕大小
const windowWidth = ref(window.innerWidth)
const listenWindowSize = () => {
  // 屏幕大小
  windowWidth.value = window.innerWidth
  // 侧边栏弹窗大小
  if (windowWidth.value < 768) {
    openSideWidth.value = "70%"
  } else {
    openSideWidth.value = "30%"
  }
}

// 侧边栏弹窗大小
const openSideWidth = ref("70%")
if (windowWidth.value < 768) {
  openSideWidth.value = "70%"
} else {
  openSideWidth.value = "30%"
}

// 主题设置数据
const sideSetting = reactive({
  mode: piniaApp.theme,
})

// 主题设置方法
const sideSettingFun = (v: string) => {
  piniaApp.theme = sideSetting.mode
}

// 生命周期 - 挂载
onMounted(() => {
  // 监控屏幕大小
  window.addEventListener('resize', listenWindowSize);
})

// 生命周期 - 结束挂载
onUnmounted(() => {
  // 监控屏幕大小
  window.removeEventListener('resize', listenWindowSize);
})
</script>

<template>
  <el-container class="main">
    <el-aside :width="isCollapse?'64px':'200px'" class="aside hidden-xs-only">
      <el-menu
          :collapse="isCollapse"
          :collapse-transition="false"
          :default-active="routerCurrent.fullPath"
          router
      >
        <el-sub-menu v-for="v in routerAdmin" :key="v.path" :index="v.path">
          <template #title>
            <el-icon>
              <i-ep-Odometer v-show="v.name === 'AdminPanel'"/>
              <i-ep-Histogram v-show="v.name === 'AdminData'"/>
              <i-ep-HelpFilled v-show="v.name === 'AdminWork'"/>
              <i-ep-Promotion v-show="v.name === 'AdminMessage'"/>
              <i-ep-Tools v-show="v.name === 'AdminSystem'"/>
            </el-icon>
            <span>{{ v.meta.title }}</span>
          </template>
          <el-menu-item
              v-for="v2 in v.children"
              :key="v2.path"
              :index="'/admin/'+v.path+'/'+v2.path"
          >
            {{ v2.meta.title }}
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header">
        <el-row align="middle" justify="center" style="height: 60px;">
          <el-col :span="12">
            <el-row align="middle" justify="start">
              <el-col :span="24">
                <el-space :size="20">
                  <!--伸缩菜单按钮-->
                  <div class="hidden-xs-only">
                    <el-icon v-if="isCollapse" @click="changeCollapse">
                      <i-ep-Expand/>
                    </el-icon>
                    <el-icon v-else @click="changeCollapse">
                      <i-ep-Fold/>
                    </el-icon>
                  </div>
                  <div class="hidden-sm-and-up">
                    <el-icon @click="changeCollapse">
                      <i-ep-Expand/>
                    </el-icon>
                  </div>
                  <!--刷新-->
                  <el-icon @click="refresh">
                    <i-ep-RefreshRight/>
                  </el-icon>
                </el-space>
              </el-col>
            </el-row>
          </el-col>
          <el-col :span="12">
            <el-row align="middle" justify="end">
              <el-space :size="20">
                <!--用户信息-->
                <div>
                  <el-dropdown>
                    <span class="header-avatar">
                      <el-avatar
                          :size="35"
                          src="https://th.bing.com/th/id/R.5c80aa95fbd3954894716d1ec12f004c?rik=flmfJ2KO%2fcItUw&riu=http%3a%2f%2fpic.ntimg.cn%2ffile%2f20180425%2f25124298_172519481324_2.jpg&ehk=lCAbTESr6UfvpTHME8gHYXlarjxwHjs8Ny4ODFRWuT4%3d&risl=&pid=ImgRaw&r=0"
                      />
                      <span>于佳怡</span>
                      <el-icon class="el-icon--right">
                        <i-ep-arrow-down/>
                      </el-icon>
                    </span>
                    <template #dropdown>
                      <el-dropdown-menu>
                        <el-dropdown-item>个人资料</el-dropdown-item>
                        <el-dropdown-item divided>退出登录</el-dropdown-item>
                      </el-dropdown-menu>
                    </template>
                  </el-dropdown>
                </div>
                <!--主题设置-->
                <el-icon @click="openSideSetting">
                  <i-ep-Operation/>
                </el-icon>
              </el-space>
            </el-row>
          </el-col>
        </el-row>
      </el-header>
      <el-main class="content">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>

  <!--侧边栏导航窗口-->
  <el-drawer v-model="isOpenSideMenu" class="openSideMenu" direction="ltr" size="70%" title="导航菜单">
    <el-menu
        :default-active="routerCurrent.fullPath"
        router
        unique-opened
        @select="openSideMenuSelectClose"
    >
      <el-sub-menu v-for="v in routerAdmin" :key="v.path" :index="v.path">
        <template #title>
          <el-icon>
            <i-ep-Odometer v-show="v.name === 'AdminPanel'"/>
            <i-ep-Histogram v-show="v.name === 'AdminData'"/>
            <i-ep-HelpFilled v-show="v.name === 'AdminWork'"/>
            <i-ep-Promotion v-show="v.name === 'AdminMessage'"/>
            <i-ep-Tools v-show="v.name === 'AdminSystem'"/>
          </el-icon>
          <span>{{ v.meta.title }}</span>
        </template>
        <el-menu-item
            v-for="v2 in v.children"
            :key="v2.path"
            :index="'/admin/'+v.path+'/'+v2.path"
        >
          {{ v2.meta.title }}
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </el-drawer>

  <!--侧边设置窗口-->
  <el-drawer v-model="isOpenSideSetting" v-model:size="openSideWidth" title="系统主题设置">
    <p>主题颜色</p>
    <el-select v-model="sideSetting.mode" placeholder="主题颜色" @change="sideSettingFun">
      <el-option label="浅色模式" value="light"/>
      <el-option label="深色模式" value="dark"/>
      <el-option label="跟随系统" value="auto"/>
    </el-select>
  </el-drawer>
</template>

<style lang="scss" scoped>
.main {
  height: 100vh;
}

.header {
  width: 100%;
  background-color: var(--el-bg-color);
  padding-right: 0;

  &-avatar {
    display: flex;
    justify-content: center;
    align-items: center;

    & > span {
      margin-left: 6px;
      font-weight: bold;
    }

    &:focus {
      outline: none;
    }
  }

  @media screen and (max-width: 768px) {
    padding: 0;
  }
}

.content {
  background-color: var(--el-bg-color-page);
}

.aside {
  transition: width linear 150ms;
}
</style>

<style lang="scss">
.openSideMenu {
  & > .el-drawer__body {
    padding: 0;
  }
}
</style>