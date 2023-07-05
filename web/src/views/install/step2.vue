<script setup lang="ts">
import {reactive, ref} from "vue";
import ProcessLabel from "@/components/process_label.vue";
import {useRouter} from "vue-router";

// 路由
const r = useRouter()

// 任务进度
const active = ref(4)

// 表单数据
const form = reactive({
  auto: {
    sid: "",
    pwd: "",
    name: "",
  },
  message: {
    type: "pushplus",
    token: "",
    openid: "",
    code: "",
  },
  sys: {
    username: "",
    password: "",
    verify: "",
  },
})

// 验证码
const code = reactive({
  s: "",
  m: "",
})

// 程序安装label
const labels = reactive([
  {
    icon: "success",
    value: "安装成功",
  }, {
    icon: "error",
    value: "安装失败",
  }, {
    icon: "process",
    value: "处理多余接口中...",
  }
])

// 返回上一步
const backStepUp = () => {
  if (active.value > 0) {
    active.value--
  }
}

// 安装按钮时候显示
const btnInstallShow = ref(true)

// 安装按钮点击事件
const btnInstallClick = () => {
  btnInstallShow.value = false
}

// 前往登录页面
const goToPageLogin = () => {
  r.push({name: "Login"})
}
</script>

<template>
  <el-row class="main">
    <el-col :span="24">
      <el-steps :active="active" finish-status="success" simple class="hidden-xs-only">
        <el-step title="评课配置"/>
        <el-step title="消息配置"/>
        <el-step title="系统配置"/>
        <el-step title="安装程序"/>
      </el-steps>
      <el-row style="margin-top: 20px;">
        <!--评课配置-->
        <el-col :span="24" v-show="active === 0">
          <h1 class="hidden-sm-and-up title">评课配置</h1>
          <el-form :model="form.auto" label-position="top">
            <el-form-item label="学号">
              <el-input v-model="form.auto.sid" clearable/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input
                  v-model="form.auto.pwd"
                  type="text"
                  show-password
                  clearable
              />
            </el-form-item>
            <el-form-item label="真实姓名">
              <el-input v-model="form.auto.name" clearable/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" class="btn">提交</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <!--消息配置-->
        <el-col :span="24" v-show="active === 1">
          <h1 class="hidden-sm-and-up title">消息配置</h1>
          <el-form :model="form.message" label-position="top">
            <el-form-item label="消息推送渠道">
              <el-select v-model="form.message.type" disabled style="width: 100%;">
                <el-option value="pushplus">推送加</el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="Token">
              <el-input v-model="form.message.token" clearable style="width: 70%"/>
              <el-button type="info" plain style="width: 30%;">获取Token</el-button>
            </el-form-item>
            <el-form-item label="验证码">
              <el-input v-model="form.message.code" clearable/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" class="btn">提交</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <!--系统配置-->
        <el-col :span="24" v-show="active === 2">
          <h1 class="hidden-sm-and-up title">系统配置</h1>
          <el-form :model="form.sys" label-position="top">
            <el-form-item label="用户名">
              <el-input type="text" v-model="form.sys.username" clearable/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="text" v-model="form.sys.password" clearable show-password/>
            </el-form-item>
            <el-form-item label="验证码">
              <el-input type="text" v-model="form.sys.verify" clearable show-password/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" class="btn">提交</el-button>
            </el-form-item>
          </el-form>
        </el-col>
        <!--安装程序-->
        <el-col :span="24" v-show="active === 3 || active === 4">
          <div v-show="active == 3">
            <h1 class="hidden-sm-and-up title">安装程序</h1>
            <ProcessLabel :icon="v.icon" v-for="(v,i) in labels" :key="i">{{ v.value }}</ProcessLabel>
            <el-button type="primary" class="btn" v-show="btnInstallShow" @click="btnInstallClick">安装程序</el-button>
          </div>
          <el-result
              icon="success"
              title="大功告成"
              sub-title="系统已成功安装，快去登录看看吧"
              v-show="active == 4"
          >
            <template #extra>
              <el-button type="primary" @click="goToPageLogin">进入登录页面</el-button>
            </template>
          </el-result>
        </el-col>
      </el-row>
      <el-row class="back" v-show="active > 0 && active != 4">
        <p @click="backStepUp">返回上一步</p>
      </el-row>
    </el-col>
  </el-row>
</template>

<style scoped lang="scss">
.main {
  min-width: 680px;
  padding: 0;

  @media screen and (max-width: 768px) {
    min-width: calc(100vw - 48px);
    height: calc(100vh - 64px);
  }

  .btn {
    width: 100%;
  }

  .title {
    text-align: center;
    margin-bottom: 40px;
  }

  .back {
    padding-top: 20px;
    font-size: 14px;
    color: gray;
  }
}
</style>