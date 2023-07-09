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
  type: "pushplus",
  token: "E34A34C938E34736923D700208646235",
})
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-form :model="data" label-width="160px">
      <el-form-item label="推送类型">
        <el-select v-model="data.type">
          <el-option value="pushplus">推送加</el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="消息 Token">
        <el-input v-model="data.token"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary">保存</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<style lang="scss" scoped>

</style>