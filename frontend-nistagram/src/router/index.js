import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Registration from '../views/Registration.vue'
import Login from '../views/Login.vue'
import Update from '../views/Update.vue'
import Account from '../views/Account.vue'
import NewPost from '../views/NewPost.vue'
import Search from '../views/Search.vue'
import UserProfile from '../views/UserProfile.vue'
import LikedPosts from '../views/LikedPosts.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/registration',
    name: 'Registration',
    component: Registration
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/update',
    name: 'Update',
    component: Update
  },
  {
    path: '/account',
    name: 'Account',
    component: Account
  },
  {
    path: '/new-post',
    name: 'new-post',
    component: NewPost
  },
  {
    path: '/search',
    name: 'Search',
    component: Search
  },
  {
    path: '/user-profile/:username',
    name: 'UserProfile',
    component: UserProfile
  },
  {
    path: '/liked-posts/',
    name: 'LikedPosts',
    component: LikedPosts
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
