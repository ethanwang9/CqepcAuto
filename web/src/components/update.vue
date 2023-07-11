<script lang="ts" setup>

// 是否更新
import {ref} from "vue";

const isUpdate = ref(true)
</script>

<template>
  <el-row class="update">
    <el-col class="update-info">
      <el-icon :size="50" class="hidden-xs-only">
        <i-ep-Refrigerator/>
      </el-icon>
      <div v-if="isUpdate">
        <p class="title">发现新版本</p>
        <p class="desc">发现新版本 v2.0.1, 本次更新 xxxx xxxx xxxx</p>
        <el-button type="danger">立即更新</el-button>
      </div>
      <div v-else>
        <p class="title">您使用的是最新版本</p>
        <p class="desc">上次检查时间: 2023-07-11 09:54:36</p>
      </div>
    </el-col>
    <el-col class="update-footer">
      <p>版本号: v2.0.0</p>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped>
.update {
  min-width: 200px;
  max-width: 500px;
  padding: 20px;
  margin: 0 auto;
  border-radius: 8px;
  color: #ffffff;
  background: #4e64ef;
  background: -webkit-linear-gradient(to right, #4e64ef, #365ce1);
  background: linear-gradient(to right, #4e64ef, #365ce1);

  &-info {
    display: flex;

    & > *:not(:last-child) {
      margin-right: 15px;
    }

    .title {
      margin: 0 0 10px 0;
      font-weight: bold;
      font-size: 20px;
    }

    .desc {
      margin: 0 0 10px 0;
      font-size: 14px;
    }
  }

  &-footer {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    font-size: 12px;

    p {
      margin-bottom: 0;
    }

    & > p:not(:last-child) {
      padding-right: 6px;
    }
  }
}
</style>