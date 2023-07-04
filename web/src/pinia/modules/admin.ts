import {defineStore} from "pinia";

const usePiniaAdmin = defineStore("admin", () => {

}, {
    persist: {
        key: "admin",
        paths: [],
    },
})

export default usePiniaAdmin