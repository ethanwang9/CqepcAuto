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
    uid: "865ADC2EAFCF4276A12BC86A756870D9",
    sid: "20230948",
    name: "孔淳美",
    class: "软件智能工程系",
    class_name: "计算机网络安全",
    type: "超级管理员",
    is_stop: false,
    time: "2023-07-11 10:37:07",
  },
  {
    uid: "700AD6253BB24C99B5A3416AF0FAE612",
    sid: "20230942",
    name: "丁淑君",
    class: "软件智能工程系",
    class_name: "计算机网络安全",
    type: "用户",
    is_stop: true,
    time: "2023-07-11 22:52:24",
  },
])

// 账号激活
const accountActive = () => {
  alert("账号激活，停用或启用账号")
}
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-table :data="data">
      <el-table-column label="用户ID" prop="uid"></el-table-column>
      <el-table-column label="学号" prop="sid"></el-table-column>
      <el-table-column label="姓名" prop="name"></el-table-column>
      <el-table-column label="二级学院" prop="class"></el-table-column>
      <el-table-column label="专业名称" prop="class_name"></el-table-column>
      <el-table-column label="账号类型" prop="type"></el-table-column>
      <el-table-column label="是否停用" prop="is_stop">
        <template #default="data">
          <!--TODO 添加停用弃用事件-->
          <el-button v-if="data.row.is_stop" text type="danger" @click="accountActive">停用</el-button>
          <el-button v-else text type="success" @click="accountActive">正常</el-button>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="time"></el-table-column>
      <!--      <el-table-column label="操作" width="120">-->
      <!--        <template #default="data">-->
      <!--          <el-button-group>-->
      <!--            <el-button>查看详情</el-button>-->
      <!--          </el-button-group>-->
      <!--        </template>-->
      <!--      </el-table-column>-->
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