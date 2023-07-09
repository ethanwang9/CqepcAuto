import {createRouter, createWebHistory} from "vue-router";
import type {RouteRecordRaw} from 'vue-router'


const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "Index",
        meta: {
            title: "首页",
        },
        redirect: {name: "Login"},
        children: [
            {
                path: "login",
                name: "Login",
                meta: {
                    title: "登录系统",
                },
                component: () => import("@/views/login.vue")
            },
            {
                path: "404",
                name: "404",
                meta: {
                    title: "页面不存在",
                },
                component: () => import("@/views/404.vue"),
            },
        ]
    },
    {
        path: "/install",
        name: "Install",
        meta: {
            title: "安装系统",
        },
        component: () => import("@/views/install/index.vue"),
        redirect: {name: "InstallStep1"},
        children: [
            {
                path: "step1",
                name: "InstallStep1",
                meta: {
                    title: "欢迎 | 安装系统",
                },
                component: () => import("@/views/install/step1.vue")
            },
            {
                path: "step2",
                name: "InstallStep2",
                meta: {
                    title: "系统配置 | 安装系统",
                },
                component: () => import("@/views/install/step2.vue")
            },
        ],
    },
    {
        path: "/admin",
        name: "Admin",
        meta: {
            title: "CqepcAuto自动化任务管理系统",
        },
        component: () => import("@/views/admin/index.vue"),
        redirect: {name: "AdminPanelIndex"},
        children: [
            {
                path: "panel",
                name: "AdminPanel",
                meta: {
                    title: "仪表盘",
                },
                children:[
                    {
                        path: "index",
                        name: "AdminPanelIndex",
                        meta: {
                            title: "主面板",
                        },
                        component: () => import("@/views/admin/panel/index.vue")
                    },
                ]
            },
            {
                path: "data",
                name: "AdminData",
                meta: {
                    title: "数据分析",
                },
                children:[
                    {
                        path: "history",
                        name: "AdminDataHistory",
                        meta: {
                            title: "评课进度",
                        },
                        component: () => import("@/views/admin/data/history.vue")
                    },
                    {
                        path: "lesson",
                        name: "AdminDataLesson",
                        meta: {
                            title: "今日课程",
                        },
                        component: () => import("@/views/admin/data/lesson.vue")
                    },
                ]
            },
            {
                path: "work",
                name: "AdminWork",
                meta: {
                    title: "自动化任务",
                },
                children:[
                    {
                        path: "config",
                        name: "AdminWorkConfig",
                        meta: {
                            title: "课程相关配置",
                        },
                        component: () => import("@/views/admin/work/config.vue")
                    },
                    {
                        path: "log",
                        name: "AdminWorkLog",
                        meta: {
                            title: "自动评课日志",
                        },
                        component: () => import("@/views/admin/work/log.vue")
                    },
                ],
            },
            {
                path: "message",
                name: "AdminMessage",
                meta: {
                    title: "消息推送",
                },
                children:[
                    {
                        path: "config",
                        name: "AdminMessageConfig",
                        meta: {
                            title: "消息推送配置",
                        },
                        component: () => import("@/views/admin/message/config.vue")
                    },
                    {
                        path: "log",
                        name: "AdminMessageLog",
                        meta: {
                            title: "消息推送日志",
                        },
                        component: () => import("@/views/admin/message/log.vue")
                    },
                ],
            },
            {
                path: "system",
                name: "AdminSystem",
                meta: {
                    title: "系统管理",
                },
                children:[
                    {
                        path: "self",
                        name: "AdminSystemSelf",
                        meta: {
                            title: "个人中心",
                        },
                        component: () => import("@/views/admin/system/self.vue")
                    },
                    {
                        path: "user",
                        name: "AdminSystemUser",
                        meta: {
                            title: "用户管理",
                        },
                        component: () => import("@/views/admin/system/user.vue")
                    },
                    {
                        path: "log",
                        name: "AdminSystemLog",
                        meta: {
                            title: "系统日志",
                        },
                        component: () => import("@/views/admin/system/log.vue")
                    },
                    {
                        path: "about",
                        name: "AdminSystemAbout",
                        meta: {
                            title: "关于系统",
                        },
                        component: () => import("@/views/admin/system/about.vue")
                    },
                ],
            },
        ],
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    // 动态设置网页标题
    window.document.title = to.meta.title as string || 'CqepcAuto'

    // 404
    if (!router.getRoutes().find((v => v.path === to.path))) {
        next({name: '404'})
    }

    // 请求正常，返回正确页面
    next()
})

router.afterEach((to, from, failure) => {

})


export default router