<script setup lang="ts">
import {ref} from "vue";

// 是否折叠菜单
const isCollapse = ref(false)

// 切换折叠菜单状态
const changeCollapse = () => isCollapse.value = !isCollapse.value

// 刷新页面
const refresh = () => window.location.reload()

// 打开侧边栏
const openSide = () => {

}
</script>

<template>
  <el-container class="main">
    <!--在<768px的窗口打开侧边栏菜单-->
    <el-aside class="aside hidden-xs-only" :width="isCollapse?'64px':'200px'">
      <el-menu
          default-active="2"
          :collapse="isCollapse"
          :collapse-transition="false"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon>
              <i-ep-Odometer/>
            </el-icon>
            <span>仪表盘</span>
          </template>
          <el-menu-item index="1-1">主面板</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="2">
          <template #title>
            <el-icon>
              <i-ep-Histogram/>
            </el-icon>
            <span>数据分析</span>
          </template>
          <el-menu-item index="2-1">今日课程</el-menu-item>
          <el-menu-item index="2-2">评课进度</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="3">
          <template #title>
            <el-icon>
              <i-ep-HelpFilled/>
            </el-icon>
            <span>自动化任务</span>
          </template>
          <el-menu-item index="3-1">评课相关配置</el-menu-item>
          <el-menu-item index="3-2">自动评课日志</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="4">
          <template #title>
            <el-icon>
              <i-ep-Promotion/>
            </el-icon>
            <span>消息推送</span>
          </template>
          <el-menu-item index="4-1">消息推送配置</el-menu-item>
          <el-menu-item index="4-2">消息推送日志</el-menu-item>
        </el-sub-menu>
        <el-sub-menu index="5">
          <template #title>
            <el-icon>
              <i-ep-Tools/>
            </el-icon>
            <span>系统管理</span>
          </template>
          <el-menu-item index="5-1">个人中心</el-menu-item>
          <el-menu-item index="5-2">用户管理</el-menu-item>
          <el-menu-item index="5-3">系统日志</el-menu-item>
          <el-menu-item index="5-4">关于系统</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header">
        <el-row justify="center" align="middle" style="height: 60px;">
          <el-col :span="12">
            <el-row justify="start" align="middle">
              <el-col :span="24">
                <el-space :size="20">
                  <!--伸缩菜单栏-->
                  <div>
                    <el-icon v-if="isCollapse" @click="changeCollapse">
                      <i-ep-Expand/>
                    </el-icon>
                    <el-icon v-else @click="changeCollapse">
                      <i-ep-Fold/>
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
            <el-row justify="end" align="middle">
              <el-space :size="20">
                <!--用户信息-->
                <!--TODO 完善用户信息显示，添加头像和姓名-->
                <div>
                  <p>userinfo</p>
                </div>
                <!--主题设置-->
                <!--完成侧边栏的打开，添加主题设置-->
                <el-icon @click="openSide">
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
      <el-footer class="footer">
        <p>Copyright © 2023 Ethan.Wang CqepcAuto自动化任务管理系统.</p>
      </el-footer>
    </el-container>
  </el-container>
</template>

<style lang="scss" scoped>
.main {
  height: 100vh;
}

.header {
  width: 100%;
  background-color: var(--el-bg-color);
  padding-right: 0;
}

.content {
  background-color: var(--el-bg-color-page);
}

.aside {
  transition: width 0.15s;
}

.footer {
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--el-bg-color);
  font-size: 14px;
  letter-spacing: 0.1em;
}
</style>