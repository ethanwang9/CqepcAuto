import {defineStore} from "pinia";
import {ref} from "vue";

// 用户信息
const usePiniaApp = defineStore("app", () => {
    const theme = ref("light")
    return {
        theme
    }
}, {
    persist: {
        key: "app",
        paths: ['theme'],
    },
})

export default usePiniaApp