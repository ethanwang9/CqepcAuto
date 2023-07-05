<script setup lang="ts">

</script>

<template>
  <div class="main">
    <el-row justify="center" class="content">
      <el-col :span="24">
        <router-view></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 5% 0;
  overflow: hidden;

  @media screen and (max-width: 768px) {
    padding: 0;
  }

  .content {
    max-width: 820px;
    padding: 32px;
    box-sizing: border-box;
    background-color: #252525;
    overflow: hidden;
    border-radius: var(--var-border-radius);

    @media screen and (max-width: 768px) {
      padding: 32px 24px;
    }
  }
}
</style>