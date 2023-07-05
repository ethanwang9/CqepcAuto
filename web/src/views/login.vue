<script setup lang="ts">
import {reactive} from "vue";
import QRCode from 'qrcode.vue'

// 登录信息类型
interface loginType {
  // 登录方式
  type: "account" | "wx" | "qq",
  // 登录标识符
  token: string,
  // 用户名
  username: string,
  // 密码
  password: string,
}

// 登录信息
const login = reactive<loginType>({
  type: "wx",
  token: "",
  username: "",
  password: "",
})

// 改变登录方式
const changeLoginType = (e: MouseEvent) => {
  const target = e.target as HTMLElement;
  login.type = target.dataset.type as loginType['type']
}
</script>

<template>
  <el-row justify="center" class="main">
    <el-col :span="24" class="login">
      <!--标题-->
      <div class="login-title">
        <h2>CA后台管理系统</h2>
        <p>自动化 · 可视化 · 简单化</p>
      </div>
      <!--登录类型-->
      <div class="login-type">
        <!--账号密码登录-->
        <div v-show="login.type === 'account'">
          <el-form label-position="top">
            <el-form-item label="账号">
              <el-input v-model="login.username" clearable/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="login.password" clearable show-password/>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" style="width: 100%;">登录</el-button>
            </el-form-item>
          </el-form>
        </div>
        <!--微信登录-->
        <div class="login-qrcode" v-show="login.type === 'wx'">
          <QRCode :size="150" value="EDD2CBDBFB5D4B0D93C9E0132D6D57B5"/>
          <p>使用微信扫一扫登录</p>
        </div>
        <!--QQ登录-->
        <div class="login-qq" v-show="login.type === 'qq'">
          <i-ep-RefreshRight class="login-qq-animation"/>
          <p>等待登录中...</p>
        </div>
      </div>
      <!--登录方式-->
      <div class="login-btn">
        <i class="iconfont" @click="changeLoginType" data-type="account" title="账号登录">&#xe646;</i>
        <i class="iconfont" @click="changeLoginType" data-type="wx" title="微信登录">&#xe641;</i>
        <i class="iconfont" @click="changeLoginType" data-type="qq" title="QQ登录">&#xe642;</i>
      </div>
    </el-col>
  </el-row>
  <el-row class="footer hidden-xs-only" justify="center" align="middle">
    <el-col :span="24">copyright &copy; 2023 Ethan.Wang CqepcAuto自动化任务管理系统</el-col>
  </el-row>
</template>

<style scoped lang="scss">
.main {
  height: 100vh;
  align-items: center;
  background-image: url("@/assets/images/bg/1.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center;
}

.login {
  max-width: 400px;
  background-color: var(--el-bg-color-page);
  border-radius: var(--var-border-radius);
  overflow: hidden;
  padding: 20px;
  box-sizing: border-box;

  &-title {
    text-align: center;
    letter-spacing: 0.1em;

    & > h2 {
      margin-bottom: 8px;
    }

    & > p {
      font-size: 14px;
      margin: 0;
    }
  }

  &-type {
    padding-top: 30px;
  }

  &-qrcode {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    letter-spacing: 2px;

    & > p {
      font-size: 14px;
      margin: 10px 0 0 0;
    }
  }

  &-qq {
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bold;
    text-align: center;
    padding: 20px 0;

    & > p {
      margin-left: 6px;

    }

    &-animation {
      animation: infiniteRotate 1s linear infinite;
      @keyframes infiniteRotate {
        0% {
          transform: rotate(0deg);
        }
        100% {
          transform: rotate(361deg);
        }
      }
    }
  }

  &-btn {
    display: flex;
    justify-content: center;
    align-items: center;
    padding-top: 30px;

    & > i {
      font-size: 35px;
    }

    & > i:not(:last-child) {
      padding-right: 40px;
    }
  }
}

.footer {
  width: 100%;
  height: 50px;
  letter-spacing: 1px;
  font-size: 12px;
  position: absolute;
  bottom: 0;
  color: #fff;
  text-align: center;
}
</style>