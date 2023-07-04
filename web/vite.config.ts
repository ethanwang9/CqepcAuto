import {resolve} from "path"
import {type ConfigEnv, type UserConfigExport, loadEnv} from "vite"
import vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import IconsResolver from 'unplugin-icons/resolver'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

export default (configEnv: ConfigEnv): UserConfigExport => {
    const env = loadEnv(configEnv.mode, process.cwd())
    return {
        resolve: {
            alias: {
                "@": resolve(__dirname, "./src")
            }
        },
        plugins: [
            vue(),
            AutoImport({
                imports: ['vue'],
                resolvers: [
                    ElementPlusResolver(),
                    IconsResolver({
                        prefix: 'Icon',
                    }),
                ],
            }),
            Components({
                resolvers: [
                    IconsResolver({
                        enabledCollections: ['ep'],
                    }),
                    ElementPlusResolver(),
                ],
            }),
            Icons({
                autoInstall: true,
            }),
        ],
        server: {
            host: true,
            port: parseInt(env.VITE_WEB_PORT),
            proxy: {
                [env.VITE_SERVER_PATH]: {
                    target: `${env.VITE_SERVER_ADDR}:${env.VITE_SERVER_PORT}/`,
                    changeOrigin: true,
                    rewrite: (path) => path.replace(new RegExp('^' + env.VITE_SERVER_PATH), ''),
                },
            },
        }
    }
}