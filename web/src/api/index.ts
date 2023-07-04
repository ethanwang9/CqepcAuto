import axios from "axios";
import {createSign, timestamp} from "@/utils/tool";


const request = axios.create({
    baseURL: import.meta.env.VITE_SERVER_PATH,
})


// 请求拦截 - 发送前
request.interceptors.request.use(function (config) {
    // 添加签名
    if (config.method === "get") {
        let {params} = config
        params["timestamp"] = timestamp()
        params["sign"] = createSign(params)
        config.params = params
    } else {
        let {data} = config
        data["timestamp"] = timestamp()
        data["sign"] = createSign(data)
        config.data = data
    }

    // 添加认证
    config.headers!["Authorization"] = "Bearer " + "token"
    config.headers!["Content-Type"] = "application/x-www-form-urlencoded"

    return config;
}, function (error) {
    return Promise.reject(error);
})

// 请求拦截- 接收后
request.interceptors.response.use(function (response) {
    // 判断返回接口状态码
    const code = response.data.code || undefined

    if (code != 200) {
        return Promise.reject(response)
    }

    return Promise.resolve(response.data);
}, function (error) {
    return Promise.reject(error);
})

export default request