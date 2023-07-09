<script lang="ts" setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)

// 表单数据
const data = reactive({
  u: "20230323",
  p: "123342",
  o: "oI2q25YdQuGQdRaDQWc5-o",
})
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-form :model="data" label-width="160px">
      <el-form-item label="学号">
        <el-input v-model="data.u"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="data.p"/>
      </el-form-item>
      <el-form-item label="微信小程序OpenID">
        <el-input v-model="data.o"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary">保存</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<style lang="scss" scoped>

</style>