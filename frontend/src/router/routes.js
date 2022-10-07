import Home from '../components/Home.vue'
// 务必要使用此种方法引用一个组件，否则wails打包后不能正常显示组件，我也不清楚怎么肥事儿
const routes = [
    {
        path: '/',
        component: Home
    },
    {
        name: 'cmd',
        path: '/cmd',
        component: () => import('../components/Cmd.vue')
    },
    {
        name: 'mditor',
        path: '/mditor',
        component: () => import('../components/Mditor.vue')
    },
    {
        name: 'setting',
        path: '/setting',
        component: () => import('../components/Setting.vue')
    }
]

export default routes
