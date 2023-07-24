# CqepcAuto 文档

> 最新版本号: v1.0.6
>
> 更新日期: 2023-07-24 15:09:36

本软件是重航自动评课系统，解决学生忘记评课、评课困难等问题开发的一个脚本系统管理软件

# 项目地址

```http
https://github.com/ethanwang9/CqepcAuto
```

# docker 编译

```shell
# 编译
docker build -f Dockerfile -t ethan/cqepc_auto:1.0.6 .
```

# docker 运行

```shell
# 运行
sudo docker run -d \
--name CqepcAuto \
-p 10000:10000 \
-v ~/docker/CqepcAuto/log:/app/log \
-v ~/docker/CqepcAuto/db:/app/db \
--restart=always \
ethan/cqepc_auto:1.0.6
```



# 安装系统

## 1. 获取小程序 Openid（可选）

使用**抓包软件**对重航微信自动评课小程序进行抓包，请求地址包含`/login`或`/loginOpenid`的请求，查看返回数据`openid`

```json
{
    "code": 200,
    "msg": "SUCCESS",
    "data": {
        "token": "4666E59B02EB4CC32DA967xxxxxxxxxx",
        "id": "1c1e5a5269704c38axxxxxxxxxxxxxxx",
        "username": "xxxxxxxx",
        "nickname": "段**",
        "has_student": 1,
        "has_cadre": 0,
        "openid": "oI2q25Ydxxxxxxxxxx5-o",  <---  这个就是我们需要的小程序 Openid
        "class_name": "2xxxxxxx班",
        "class_code": "2xxxxxxxx"
    }
}
```

## 2. 获取钉钉群机器人

1. 打开钉钉（电脑版），创建群
2. 进入创建的群聊，添加**自定义机器人**按照下图设置

![image 1](https://compeition-excute.oss-cn-beijing.aliyuncs.com/poss/438c9e1bbfdadc1e8337f813972038f7.png)

## 3. 登录模式

|      | 账号密码登录                                                 | 小程序OPENID登录                               |
| ---- | ------------------------------------------------------------ | ---------------------------------------------- |
| 优点 | 使用账号和密码登录                                           | 通过小程序OPENID登录，登录方式比较接近原生模式 |
| 缺点 | 1. 自动生成小程序OPENID，下一次手机登录小程序需要手动输入账号密码<br />2. 无法接收微信小程序原生的评课消息通知 | 有几率获取到微信小程序评课消息通知             |

由于微信小程序的限制，必须要点击消息授权才能接收消息通知，所以原生消息通知可能无法获取，但是该评课系统采用**钉钉**消息通知，无需担心

# 修复程序

当程序出现问题时，请运行程序目录下的`fix/fix.exe`修复程序出现的问题

程序出现闪退，请查看`log/server_error.log`滚动到文件末，查看具体错误原因

# 常见问题

**程序在终端显示异常**

- 请更换终端运行即可，推荐：Windows 终端，该问题仅在 Windows 出现

# 更新日志

Version 1.0.6

- 修复依赖中的漏洞

Version 1.0.5

- 去除强制校验程序正版

Version 1.0.4

- 该版本为最终版本，后续将维护系统稳定，不维护新功能和当前功能
- 不维护原因：毕业咯！