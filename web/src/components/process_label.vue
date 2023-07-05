<script setup lang="ts">
import {computed} from "vue";

const props = defineProps({
  icon: {
    type: String,
    required: true,
    validator(value) {
      return ['success', 'error', 'process'].includes(value)
    }
  },
})

const propsType = computed(() => {
  if (props.icon === 'error') {
    return 'danger'
  } else if (props.icon === 'process') {
    return ''
  } else if (props.icon === 'success') {
    return 'success'
  }
})
</script>

<template>
  <el-row class="app_process_label">
    <el-col :span="24">
      <el-tag :type="propsType" effect="dark">
        <div class="icon">
          <i-ep-SuccessFilled v-show="props.icon == 'success'"/>
          <span v-show="props.icon == 'success'">成功</span>
          <i-ep-WarnTriangleFilled v-show="props.icon == 'error'"/>
          <span v-show="props.icon == 'error'">失败</span>
          <i-ep-MoreFilled v-show="props.icon == 'process'"/>
          <span v-show="props.icon == 'process'">处理中</span>
        </div>
      </el-tag>
      <span class="text"><slot name="default"></slot></span>
    </el-col>
  </el-row>
</template>

<style scoped lang="scss">
.app_process_label {
  margin: 10px 0;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  overflow: hidden;

  .icon {
    display: flex;
    justify-content: center;
    align-items: center;

    & > span {
      margin-left: 4px;
    }
  }

  .text {
    margin-left: 8px;
  }
}
</style>