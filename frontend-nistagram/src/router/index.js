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
import DislikedPosts from '../views/DislikedPosts.vue'
import SavedPosts from '../views/SavedPosts.vue'
import VerificationRequest from '../views/VerificationRequest.vue'
import ProfilePrivacy from '../views/ProfilePrivacy.vue'
import BlockedUsers from '../views/BlockedUsers.vue'
import MutedUsers from '../views/MutedUsers.vue'
import CloseFriends from '../views/CloseFriends.vue'
import Notification from '../views/Notification.vue'
import AdministratorHome from '../views/AdministratorHome.vue'
import VerificationRequests from '../views/VerificationRequests.vue'
import RemoveProfiles from '../views/RemoveProfiles.vue'
import InappropriateContent from '../views/InappropriateContent.vue'
import NotificationPost from '../views/ViewPostFromNotification.vue'
import AgentVerification from '../views/AgentVerification.vue'

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
  },
  {
    path: '/disliked-posts/',
    name: 'DislikedPosts',
    component: DislikedPosts
  },
  {
    path: '/saved-posts/',
    name: 'SavedPosts',
    component: SavedPosts
  },
  {
    path: '/verify',
    name: 'VerificationRequest',
    component: VerificationRequest
  },
  {
    path: '/updateProfilePrivacy',
    name: 'ProfilePrivacy',
    component: ProfilePrivacy
  },
  {
    path: '/blocked-users',
    name: 'BlockedUsers',
    component: BlockedUsers
  },
  {
    path: '/muted-users',
    name: 'MutedUsers',
    component: MutedUsers
  },
  {
    path: '/close-friends',
    name: 'CloseFriends',
    component: CloseFriends
  },
  {
    path: '/notifications',
    name: 'Notification',
    component: Notification
  },
  {
    path: '/administratorHome',
    name: 'AdministratorHome',
    component: AdministratorHome
  },
  {
    path: '/verificationRequests',
    name: 'VerificationRequests',
    component: VerificationRequests
  },
  {
    path: '/removeProfiles',
    name: 'RemoveProfiles',
    component: RemoveProfiles
  },
  {
    path: '/inappropriateContent',
    name: 'InappropriateContent',
    component: InappropriateContent
  },
  {
    path: '/notification-post/:post/:contentType',
    name: 'NotificationPost',
    component: NotificationPost
  },
  {
    path: '/agentVerification',
    name: 'AgentVerification',
    component: AgentVerification
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
