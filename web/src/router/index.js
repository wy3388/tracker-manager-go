import {createRouter, createWebHistory} from "vue-router"

const routes = [
    {
        path: '/',
        redirect: '/root/tracker',
    },
    {
        path: '/root',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
        children: [
            {
                path: 'tracker',
                name: 'Tracker',
                component: () => import('@/views/Tracker.vue')
            },
            {
                path: 'source',
                name: 'Source',
                component: () => import('@/views/Source.vue')
            },
            {
                path: 'sync',
                name: 'Sync',
                component: () => import('@/views/Sync.vue')
            },
            {
                path: 'client',
                name: 'Client',
                component: () => import('@/views/Client.vue')
            }
        ]
    },
    {
        path: '/:catchAll(.*)',
        redirect: '/root/tracker'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})
export default router