const {defineConfig} = require('@vue/cli-service')
module.exports = defineConfig({
    transpileDependencies: true,
        chainWebpack: config => {
        config
            .plugin('html')
            .tap(args => {
                args[0].title= 'NFT 管理平台'
                args[0].icon="favicon.ico"
                return args
            })
    },
    pwa: {
        name: 'NFT 管理平台',
        themeColor: '#ff6700',
        msTileColor: '#ff6700',
        display: 'fullscreen',
        scope: "/",
        start_url: "/",
        manifestOptions: {
            background_color: '#efeeee'
        }
    },
})
