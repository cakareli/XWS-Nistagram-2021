import Vue from 'vue'
import VueRouter from 'vue-router'

import GuestHome from '@/views/guest/GuestHome.vue'

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        component: GuestHome,
        name : 'GuestHome'
    }
];

const router = new VueRouter({
    routes
})

export default router;