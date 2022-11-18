import Home from '../views/Home.vue'
// 务必要使用此种方法引用一个组件，否则wails打包后不能正常显示组件，我也不清楚怎么肥事儿
const routes = [
    {
        path: '/',
        component: Home
    },
    {
        name: 'cmd',
        path: '/cmd',
        component: () => import('../views/Cmd.vue')
    },
    {
        name: 'mditor',
        path: '/mditor',
        component: () => import('../views/Mditor.vue')
    },
    {
        name: 'wditor',
        path: '/wditor',
        component: () => import('../views/WangEditor.vue')
    },
    {
        name: 'setting',
        path: '/setting',
        component: () => import('../views/Setting.vue')
    },
    {
        name: 'todo',
        path: '/todo',
        component: () => import('../views/Todo.vue')
    },
    {
        name: 'about',
        path: '/about',
        component: () => import('../views/About.vue')
    },
    {
        name: 'ocr',
        path: '/ocr',
        component: () => import('../views/Ocr.vue')
    },
]

export default routes
