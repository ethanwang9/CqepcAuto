import {Md5} from "ts-md5"

// 获取当前格式化时间
export function dateTime(datetime: number): string {
    let date = new Date(datetime)// 时间戳为10位需*1000，时间戳为13位的话不需乘1000
    let year = date.getFullYear()
    let month = ('0' + (date.getMonth() + 1)).slice(-2)
    let sedate = ('0' + date.getDate()).slice(-2)
    let hour = ('0' + date.getHours()).slice(-2)
    let minute = ('0' + date.getMinutes()).slice(-2)
    let second = ('0' + date.getSeconds()).slice(-2)
    // 返回
    return year + '-' + month + '-' + sedate + ' ' + hour + ':' + minute + ':' + second
}

// 获取当前时间戳
export function timestamp(): number {
    return parseInt(String(new Date().getTime() / 1000))
}

// 签名算法
export function createSign(formDate: any): string {
    // 去除空值
    for (let key in formDate) {
        if (formDate[key] == "") {
            delete formDate[key];
        }
    }
    // 排序
    let formDate2 = Object.keys(formDate).sort();
    let formDate3: any = {};
    for (let i = 0; i < formDate2.length; i++) {
        formDate3[formDate2[i]] = formDate[formDate2[i]];
    }
    formDate = formDate3;
    // 生成序列
    let sign = "";
    for (let index in formDate) {
        sign += index + formDate[index];
    }
    sign += import.meta.env.VITE_SIGN_KEY

    return Md5.hashStr(sign)
}