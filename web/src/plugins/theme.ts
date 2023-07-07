import {type App, watchEffect} from "vue";
import {useColorMode} from '@vueuse/core'
import usePiniaApp from "@/pinia/modules/app";

export default {
    install: function (app: App) {
        // 获取用户缓存
        const piniaApp = usePiniaApp()

        // 获取主题管理方案
        const mode = useColorMode({
            emitAuto: true,
        })

        // 监听用户主题设置
        watchEffect(() => {
            switch (piniaApp.theme) {
                case "auto":
                    mode.value = "auto"
                    break
                case "light":
                    mode.value = "light"
                    break
                case "dark":
                    mode.value = "dark"
                    break
            }
        })
    }
}