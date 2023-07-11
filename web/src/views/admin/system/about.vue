<script lang="ts" setup>
import {useRouter} from "vue-router";
import Update from "@/components/update.vue";
// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <h2 class="title">CqepcAuto 自动化任务管理系统</h2>
    <p class="desc">欢迎使用 CqepcAuto 自动化任务管理系统，本系统简称 CqepcAuto 是一款开源免费的自动化任务管理系统，主要用于特定的校园教务系统评课任务。
      软件设计的初衷是为了解放双手，让复杂、易忘、不必要的流程精简化，化繁为简极至轻量。</p>
    <Update class="update"/>
  </el-card>
</template>

<style lang="scss" scoped>
.title {
  text-align: center;
}

.desc {
  text-indent: 2em;
  line-height: 1.5em;
}

.update {
  margin-top: 10%;
}
</style>