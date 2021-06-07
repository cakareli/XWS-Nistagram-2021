import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '@/views/guest/Home.vue'
import Registration from "@/views/Registration.vue"

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        component: Home,
        name : 'Home'   
    }, 

    {
        path: '/registration',
        component: Registration,
        name: 'Registration',
    }
      
];

const router = new VueRouter({
    routes
})

export default router;