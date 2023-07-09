<script lang="ts" setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";
// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)

// 数据
const data = reactive([
  {
    id: "05031077",
    name: "微信小程序开发",
    join: 5,
    total: 5,
    pk: "100%",
  },
  {
    id: "05031123",
    name: "Java程序开发",
    join: 4,
    total: 5,
    pk: "80%",
  },
])
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-table :data="data" border>
      <el-table-column label="课程ID" prop="id"/>
      <el-table-column label="课程名称" prop="name"/>
      <el-table-column label="已上课" prop="join">
        <template #default="data">
          <span>{{ data.row.join }} 次</span>
        </template>
      </el-table-column>
      <el-table-column label="已评价" prop="total">
        <template #default="data">
          <span>{{ data.row.total }} 次</span>
        </template>
      </el-table-column>
      <el-table-column label="评课率" prop="pk"/>
    </el-table>
  </el-card>
</template>

<style lang="scss" scoped>

</style>