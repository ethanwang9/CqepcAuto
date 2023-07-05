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