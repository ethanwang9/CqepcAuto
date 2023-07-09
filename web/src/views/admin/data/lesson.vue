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
    rid: "2021092203040503104605010093",
    cid: "05031077",
    name: "微信小程序开发",
    teacher: "方贺祥",
    classroom: "A02501",
    time: "2023-07-09 21:19:28",
    pk: true,
  },
  {
    rid: "2021092203040503104605010099",
    cid: "05031123",
    name: "Java程序开发",
    teacher: "张建林",
    classroom: "A02501",
    time: "2023-07-09 21:19:29",
    pk: false,
  },
])
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-alert title="补评只能补评当天的课程，一般在开课后一个小时之内成功率较高，其他时间成功率不超过45%且该功能非常危险，有较高概率会被请去喝茶！" type="warning" />
    <br>
    <el-table :data="data" border>
      <el-table-column label="教室ID" prop="rid"/>
      <el-table-column label="课程ID" prop="cid"/>
      <el-table-column label="课程名称" prop="name"/>
      <el-table-column label="老师" prop="teacher"/>
      <el-table-column label="教室" prop="classroom"/>
      <el-table-column label="上课时间" prop="time"/>
      <el-table-column label="是否评课" prop="pk">
        <template #default="data">
          {{ data.row.pk ? '已评课' : '未评课'}}
        </template>
      </el-table-column>
      <el-table-column label="工具" prop="help">
        <template #default="data">
          <el-button>补评</el-button>
        </template>
      </el-table-column>
    </el-table>
    <p class="time">更新时间: 2023-07-09 21:30:13</p>
  </el-card>
</template>

<style lang="scss" scoped>
.time {
  width: 100%;
  text-align: right;
  font-size: 12px;
  margin-bottom: 0;
}
</style>