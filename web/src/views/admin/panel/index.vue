<script lang="ts" setup>
import {reactive} from "vue";

// 版本信息
const versionInfo = reactive([
  {
    key: "当前版本",
    value: "v2.0.0",
  },
  {
    key: "最新版本",
    value: "v2.0.0",
  },
  {
    key: "检查时间",
    value: "2023-07-08 18:25:48",
  },
])

// 今日课表
const todayClass = reactive([
  {
    "time": "2023-05-22 8:30",
    "name": "微信小程序开发",
    "teacher": "孙红雷",
    "room": "A02501",
    "is_pk": true,
  },
  {
    "time": "2023-05-22 10:20",
    "name": "Java技术开发",
    "teacher": "孙红雷",
    "room": "A02506",
    "is_pk": false,
  },
])

// 日志信息
const logInfo = reactive([
  {
    id:"08B07357DE41436DBF6C3BB80C923F83",
    action: "系统-登录日志",
    details:"用户在 ip:xxx.xxx.xxx 归属地: xxx xxx xxx 成功登录",
    time: "2023-07-08 23:20:14",
  },
  {
    id:"D0C6B978DC44425B9A030866C428BC9B",
    action: "系统-定时任务",
    details:"今日课程: 微信小程序开发, 已成功评价成功. 评价详细信息: xxx xxx xxx",
    time: "2023-07-08 23:20:14",
  },{
    id:"D0C6B978DC44425B9A030866C428BC9B",
    action: "系统-定时任务",
    details:"今日课程: 微信小程序开发, 已成功评价成功. 评价详细信息: xxx xxx xxx",
    time: "2023-07-08 23:20:14",
  },{
    id:"D0C6B978DC44425B9A030866C428BC9B",
    action: "系统-定时任务",
    details:"今日课程: 微信小程序开发, 已成功评价成功. 评价详细信息: xxx xxx xxx",
    time: "2023-07-08 23:20:14",
  },{
    id:"D0C6B978DC44425B9A030866C428BC9B",
    action: "系统-定时任务",
    details:"今日课程: 微信小程序开发, 已成功评价成功. 评价详细信息: xxx xxx xxx",
    time: "2023-07-08 23:20:14",
  }
])
</script>

<template>
  <el-row :gutter="15" class="data">
    <el-col :md="6" :sm="12">
      <el-card>
        <el-statistic :value="36500" suffix="天" title="系统运行天数"/>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12">
      <el-card>
        <el-statistic :value="7000000" suffix="次" title="自动任务次数"/>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12" :xs="12">
      <el-card>
        <el-statistic :value="8" suffix="节" title="今日课程数"/>
      </el-card>
    </el-col>
    <el-col :md="6" :sm="12" :xs="12">
      <el-card>
        <el-statistic class="user_type" suffix="超级管理员" title="用户身份" value-style="display: none;"/>
      </el-card>
    </el-col>
  </el-row>
  <el-row :gutter="15">
    <el-col :md="17" class="leftMain">
      <el-card>
        <h2>欢迎回来，王宝钏同学</h2>
        <p>计算机工程系 - 23软件技术04班(专本贯通)</p>
      </el-card>
      <el-card>
        <el-table v-show="todayClass.length != 0" :data="todayClass" border>
          <el-table-column label="上课时间" min-width="150" prop="time"/>
          <el-table-column label="课程名称" min-width="200" prop="name"/>
          <el-table-column label="教师" prop="teacher"/>
          <el-table-column label="教室" prop="room"/>
          <el-table-column label="自动任务" min-width="95" prop="is_pk">
            <template #default="scope">
              <span v-if="scope.row.is_pk">已评课</span>
              <span v-else>未评课</span>
            </template>
          </el-table-column>
        </el-table>
        <el-result
            v-show="todayClass.length == 0"
            icon="info"
            sub-title="给自己放个假，休息一下吧！"
            title="今日无课程"
        />
      </el-card>
      <el-card>
        <el-table :data="logInfo" border>
          <el-table-column label="ID"  prop="id"/>
          <el-table-column label="日志类型"  prop="action"/>
          <el-table-column label="日志信息" prop="details"/>
          <el-table-column label="记录时间" prop="time"/>
        </el-table>
      </el-card>
    </el-col>
    <el-col :md="7" class="rightSideCard">
      <el-card>
        <template #header>
          <div style="display: flex;justify-content: space-between;align-items:center;">
            <span>版本信息</span>
            <el-button class="button" text>查看更多</el-button>
          </div>
        </template>
        <el-table :data="versionInfo" :show-header="false">
          <el-table-column prop="key"/>
          <el-table-column prop="value"/>
        </el-table>
      </el-card>
      <el-card class="rightSideCard-info" header="实时监控">
        <div class="rightSideCard-info-box">
          <el-progress :indeterminate="true" :percentage="80" type="dashboard">
            <template #default="{ percentage }">
              <span class="percentage-value">{{ percentage }}%</span>
              <span class="percentage-label">处理器</span>
            </template>
          </el-progress>
          <el-progress :indeterminate="true" :percentage="80" type="dashboard">
            <template #default="{ percentage }">
              <span class="percentage-value">{{ percentage }}%</span>
              <span class="percentage-label">内存</span>
            </template>
          </el-progress>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped>
.data {
  & > * {
    margin-bottom: 15px;
  }
}

.leftMain {
  & > * {
    margin-bottom: 15px;
  }
}

.rightSideCard {
  & > * {
    margin-bottom: 15px;
  }

  &-info {
    &-box {
      display: flex;
      justify-content: space-around;
      align-items: center;
      flex-wrap: wrap;

      & > *:not(:last-child) {
        margin-bottom: 10px;
      }
    }

    .percentage-value {
      display: block;
      margin-top: 10px;
      font-size: 28px;
    }

    .percentage-label {
      display: block;
      margin-top: 10px;
      font-size: 12px;
    }
  }
}
</style>

<style lang="scss">
//覆盖样式
.data .user_type .el-statistic__suffix {
  margin-left: 0;
}
</style>