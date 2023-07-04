// name: 接口定义
// author: Ethan.Wang
// desc:

// 全局接口返回类型
export type PromiseRes<T> = Promise<RequestRes<T>>

export interface RequestRes<T = {}> {
    code: number,
    message: string,
    data: T,
}
