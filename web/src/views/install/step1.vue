<script setup lang="ts">
import {useRouter} from "vue-router";
// 路由
const r = useRouter()

// 下一页
const nextPage = () => {
  r.push({name: "InstallStep2"})
}
</script>

<template>
  <el-row>
    <el-col :span="24">
      <!--标题-->
      <h1 class="ins-title">系统安装引导程序</h1>
      <!--程序介绍-->
      <h2>简介</h2>
      <p>欢迎使用 CqepcAuto
        管理系统，接下来请您花费几分钟对系统初始化。您需要填写一些基本信息，稍后系统会对您填写的信息进行校验，通过后即可进入后台管理界面。</p>
      <p>下面是系统基础信息：</p>
      <p><b>
        版本号：v2.0.0<br/>
        更新时间：2023-07-04 22:15:38<br/>
        作者：Ethan.Wang<br/>
        软件仓库地址： <a href="https://www.baidu.com">github.com</a>
      </b></p>
      <!--协议条约-->
      <h2>免责声明</h2>
      <p>任何用户在使用由 Ethan.Wang 作者（以下简称「本作者」）开发的开源软件 CqepcAuto
        管理系统（以下简称「本程序」）前，请您仔细阅读并透彻理解本声明。</p>
      <p>您可以选择不使用本程序，若您一旦使用本程序，您的使用行为即被视为对本声明全部内容的认可和接受。</p>
      <ol style="word-break: break-all;">
        <li>
          本程序是一款开源免费的自动化任务管理系统，主要用于特定的校园教务系统评课任务。本程序将使用本地数据库记录您的「消息推送Token、个人身份信息、课程信息、系统运行日志」动态功能，确保系统正常运行。
        </li>
        <li>本程序尊重并保护所有用户的个人隐私权，不窃取任何用户计算机中的信息。更不会通过任何形式，如:
          收集用户使用日志、上传用户地理位置信息、通过网络传输存储用户敏感信息等网络传输功能。
        </li>
        <li>您承诺秉着合法、合理的原则使用
          本程序，不利用本程序进行任何违法、侵害他人合法利益等恶意的行为，亦不将本程序运用于任何违反我国法律法规的程序、网站等具备网络通信的终端中。
        </li>
        <li>
          任何单位或个人因下载使用本程序而产生的任何意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失(包括但不限于直接、间接、附带或衍生的损失等)，本作者不承担任何法律责任。
        </li>
        <li>
          用户明确并同意本声明条款列举的全部内容，对使用本程序可能存在的风险和相关后果将完全由用户自行承担，本作者不承担任何法律责任。
        </li>
        <li>任何单位或个人在阅读本免责声明后，应在「Apache 开源许可证」<a
            href="https://github.com/ethanwang9/CqepcAuto/blob/master/LICENSE">https://github.com/ethanwang9/CqepcAuto/blob/master/LICENSE</a>
          所允许的范围内进行合法的发布、传播和使用本程序等行为，若违反本免责声明条款或违反法律法规所造成的法律责任(包括但不限于民事赔偿和刑事责任），由违约者自行承担。
        </li>
        <li>
          如果本声明的任何部分被认为无效或不可执行，则该部分将被解释为反映本作者的初衷，其余部分仍具有完全效力。不可执行的部分声明，并不构成我们放弃执行该声明的权利。
        </li>
        <li>
          本作者有权随时对本声明条款及附件内容进行单方面的变更，并以Github仓库更新、网页公告等方式予以公布，公布后立即自动生效，无需另行单独通知；若您在本声明内容公告变更后继续使用的，表示您已充分阅读、理解并接受修改后的声明内容。
        </li>
      </ol>
    </el-col>
  </el-row>
  <!--按钮-->
  <el-row justify="center">
    <el-button type="danger" @click="nextPage">同意免责声明</el-button>
  </el-row>
</template>

<style scoped lang="scss">
p, ol {
  line-height: 1.6em;
}
</style>