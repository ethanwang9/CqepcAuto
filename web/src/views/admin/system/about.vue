<script lang="ts" setup>
import {useRouter} from "vue-router";

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
    <p>xxx</p>
  </el-card>
</template>

<style lang="scss" scoped>

</style>