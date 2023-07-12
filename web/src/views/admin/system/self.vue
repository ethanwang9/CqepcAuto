<script lang="ts" setup>
import {onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";

// 路由
const router = useRouter()
// 当前路由
const currentRoute = router.currentRoute.value.matched
currentRoute.splice(0, 1)

// 表单数据
const data = reactive({
  avatar: "http://pic.ntimg.cn/file/20180425/25124298_172519481324_2.jpg",
  name: "汪溶溶",
  sid: "20239042",
  class: "计算机工程系",
  className: "网络安全",
  classCode: "2318012723123",
  // ==========
  newPass: "",
  ConfPass: "",
  // ==========
  qq: "99FB49A3EA26464BB75D2EDD73331EB3",
  wx: "123",
})

// 活动下标
const active = ref('user')

// 导航条位置
const sidePosition = ref("top")

// 监控屏幕大小
const listenWindowSize = function () {
  // 屏幕大小
  let w = window.innerWidth
  if (w > 768) {
    sidePosition.value = "left"
  } else {
    sidePosition.value = "top"
  }
}
listenWindowSize()

// 生命周期 - 挂载
onMounted(() => {
  // 监控屏幕大小
  window.addEventListener('resize', listenWindowSize);
})

// 生命周期 - 结束挂载
onUnmounted(() => {
  // 监控屏幕大小
  window.removeEventListener('resize', listenWindowSize);
})
</script>

<template>
  <el-breadcrumb>
    <el-breadcrumb-item v-for="(v,i) in currentRoute" :key="i">{{ v.meta.title }}</el-breadcrumb-item>
  </el-breadcrumb>
  <br>
  <el-card>
    <el-tabs v-model="active" :tab-position="sidePosition">
      <el-tab-pane label="个人信息" name="user">
        <el-form :model="data" label-width="80px">
          <el-form-item>
            <el-avatar :size="128" :src="data.avatar" style="margin: 0 auto;"/>
          </el-form-item>
          <el-form-item label="头像">
            <el-input v-model="data.avatar"/>
          </el-form-item>
          <el-form-item label="姓名">
            <el-input v-model="data.name" disabled/>
          </el-form-item>
          <el-form-item label="学号">
            <el-input v-model="data.sid" disabled/>
          </el-form-item>
          <el-form-item label="二级学院">
            <el-input v-model="data.class" disabled/>
          </el-form-item>
          <el-form-item label="专业">
            <el-input v-model="data.className" disabled/>
          </el-form-item>
          <el-form-item label="班级号">
            <el-input v-model="data.classCode" disabled/>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="面板信息" name="panel">
        <el-form :model="data" label-width="80px">
          <el-form-item label="新密码">
            <el-input v-model="data.newPass" clearable show-password/>
          </el-form-item>
          <el-form-item label="确认密码">
            <el-input v-model="data.ConfPass" clearable show-password/>
          </el-form-item>
          <el-form-item>
            <el-button type="primary">提交</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="绑定账号" name="bind">
        <el-row>
          <el-col class="bind">
            <div class="bind-info">
              <i class="iconfont bind-icon" title="微信">&#xe641;</i>
              <div v-if="data.wx.length != 0">
                <span>已绑定</span>
                <el-button type="danger">解绑</el-button>
              </div>
              <div v-else>
                <span>未绑定</span>
                <el-button type="primary">绑定</el-button>
              </div>
            </div>
            <div v-show="data.wx.length != 0" class="bind-input">
              <el-input v-model="data.wx" disabled/>
            </div>
          </el-col>
          <el-col class="bind">
            <div class="bind-info">
              <i class="iconfont bind-icon" title="QQ">&#xe642;</i>
              <div v-if="data.qq.length != 0">
                <span>已绑定</span>
                <el-button type="danger">解绑</el-button>
              </div>
              <div v-else>
                <span>未绑定</span>
                <el-button type="primary">绑定</el-button>
              </div>
            </div>
            <div v-show="data.qq.length != 0" class="bind-input">
              <el-input v-model="data.qq" disabled/>
            </div>
          </el-col>
        </el-row>
      </el-tab-pane>
    </el-tabs>
  </el-card>
</template>

<style lang="scss" scoped>
.bind {
  display: flex;
  flex-direction: column;
  margin: 0 0 15px 10px;

  &-icon {
    padding-right: 10px;
    font-size: 30px;
  }

  &-info {
    display: flex;

    & > div > span {
      padding: 0 15px 0 50px;
    }
  }

  &-input {
    width: 350px;
    padding-top: 15px;


  }

  @media screen and (max-width: 768px) {
    &-input {
      width: 100%;
    }
  }
}
</style>