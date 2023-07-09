<script lang="ts" setup>
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";

// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)

// 表单数据
const data = reactive({
  type: "pushplus",
  token: "E34A34C938E34736923D700208646235",
})

// 活动下标
const active = ref('user')
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <!--TODO 小屏幕标签栏显示在上方-->
    <el-tabs v-model="active" tab-position="left">
      <el-tab-pane name="user" label="个人信息">头像 | 姓名、学号、二级学院、专业、班级号</el-tab-pane>
      <el-tab-pane name="panel" label="面板信息">修改 账号、密码</el-tab-pane>
      <el-tab-pane name="bind" label="绑定账号">QQ、微信</el-tab-pane>
    </el-tabs>
  </el-card>
</template>

<style lang="scss" scoped>

</style>