import {useDark, useToggle} from "@vueuse/core";
import {reactive, computed, App, watch, onMounted} from "vue";
import usePiniaApp from "@/pinia/modules/app";

export default {
    install: function (app: App) {
        // 初始化Pinia
        const piniaApp = usePiniaApp()

        // 主题选项
        const theme = reactive({
            // 系统主题
            env: "light",
            // 环境主题
            mode: computed(() => {
                return useDark().value ? 'dark' : 'light'
            }),
            // 选择主题模式
            // auto - 自动切换 | light - 亮色模式 | dark - 暗色模式
            select: computed(() => {return piniaApp.theme}),
            // 切换主题
            changeTheme: useToggle(useDark()),
        });

        // 监听系统环境主题
        const darkThemeQuery = window.matchMedia('(prefers-color-scheme: dark)')
        const lightThemeQuery = window.matchMedia('(prefers-color-scheme: light)')
        theme.env = darkThemeQuery.matches ? 'dark' : 'light'
        darkThemeQuery.onchange = (e) => {
            if (e.matches) {
                theme.env = "dark"
            }
        }
        lightThemeQuery.onchange = (e) => {
            if (e.matches) {
                theme.env = 'light'
            }
        }

        // 自动切换主题
        watch(theme, () => {
            if (theme.select === "auto" && theme.mode !== theme.env) {
                theme.changeTheme()
            } else if (theme.select !== "auto" && theme.select !== theme.mode) {
                theme.changeTheme()
            }
        })
    }
}