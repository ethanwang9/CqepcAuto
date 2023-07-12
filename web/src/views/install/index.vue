<script lang="ts" setup>

</script>

<template>
  <div class="main">
    <el-row class="content" justify="center">
      <el-col :span="24">
        <router-view></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<style lang="scss" scoped>
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
    //background-color: #252525;
    overflow: hidden;
    border-radius: var(--var-border-radius);

    @media screen and (max-width: 768px) {
      padding: 32px 24px;
    }
  }
}

// 浅色模式
html.light {
  .content {
    background-color: #ffffff;
    box-shadow: var(--el-box-shadow-light);
  }
}

// 深色模式
html.dark {
  .content {
    background-color: #252525;
  }
}


</style>