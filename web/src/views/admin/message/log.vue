<script lang="ts" setup>
import {reactive} from "vue";
import {useRouter} from "vue-router";

// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)

// 表单数据
const data = reactive([
  {
    id: "E57BED04A3B842F0A3AAFD13265691CF",
    type: "系统-消息推送日志",
    detail: "[定时任务] 系统已自动评课: 微信小程序 成功, 评课详细信息: xxx xxx xxx",
    time: "2023-07-09 22:06:52",
  },
])
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-table :data="data">
      <el-table-column label="ID" prop="id"></el-table-column>
      <el-table-column label="日志类型" prop="type"></el-table-column>
      <el-table-column label="日志信息" prop="detail"></el-table-column>
      <el-table-column label="记录时间" prop="time"></el-table-column>
    </el-table>
    <div class="pagination">
      <!--TODO 小屏幕只显示 prev, pager, next-->
      <el-pagination
          :hide-on-single-page="true"
          :page-sizes="[10, 20, 30, 50]"
          :total="100"
          background
          current-page="1"
          layout="total, sizes, prev, pager, next, jumper"
          page-size="10"
      />
    </div>
  </el-card>
</template>

<style lang="scss" scoped>
.pagination {
  width: 100%;
  padding-top: 15px;
  display: flex;
  justify-content: center;
}
</style>