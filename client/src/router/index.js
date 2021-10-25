import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Auth from '../views/Auth.vue'
import Todo from '../views/Todo.vue'
import store from '../store'

const requireAuth = (to, from, next) => {
    if (!store.getters['auth/isAuth'] && store.state.ready) {
        next({ name: 'Auth' })
    } else {
        next()
    }
}

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        beforeEnter: requireAuth,
    },
    {
        path: '/auth',
        name: 'Auth',
        component: Auth,
    },
    {
        path: '/list/',
        component: Home,
        beforeEnter: requireAuth,
        children: [
            {
                path: ':listId',
                name: 'todo',
                component: Todo,
            },
        ],
    },
    // {
    //     path: '/list/:listId',
    //     name: 'todo',
    //     component: Todo,
    //     beforeEnter: requireAuth,
    // },
    {
        path: '/about',
        name: 'About',
        beforeEnter: requireAuth,
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
    },
    {
        path: '/:catchAll(.*)',
        component: () => import('../views/NotFound.vue'),
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router
