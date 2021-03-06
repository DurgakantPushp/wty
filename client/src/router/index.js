import Vue from 'vue'
import Router from 'vue-router'

import auth from '../auth'

import Register from '@/components/Register'
import Login from '@/components/Login'

import BlockchainStatus from '@/components/BlockchainStatus'
import Home from '@/components/Home'
import GuestHome from '@/components/GuestHome'
import EntityHome from '@/components/EntityHome'
import TeamHome from '@/components/TeamHome'
import BusinessHome from '@/components/BusinessHome'
import AdminHome from '@/components/AdminHome'

import TechUsed from '@/components/TechUsed'
import Blockchain from '@/components/Blockchain'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/home',
      name: 'Guest Home',
      component: GuestHome
    },
    {
      path: '/users/signup',
      name: 'sign up',
      component: Register
    },
    {
      path: '/users/login',
      name: 'login',
      component: Login
    },
    {
      path: '/users/home',
      name: 'normal user home',
      component: Home,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'user')
    },
    {
      path: '/users/entity/home',
      name: 'entity home',
      component: EntityHome,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'entity')
    },
    {
      path: '/users/wty-team/home',
      name: 'wty-team home',
      component: TeamHome,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'team')
    },
    {
      path: '/users/business/home',
      name: 'business home',
      component: BusinessHome,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'business')
    },
    {
      path: '/admin/home',
      name: 'admin home',
      component: AdminHome,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'admin')
    },
    {
      path: '/admin/tech-used',
      name: 'technology used',
      component: TechUsed,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'admin')
    },
    {
      path: '/admin/blockchain',
      name: 'current blockchain',
      component: Blockchain,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'admin')
    },
    {
      path: '/admin/blockchain/status',
      name: 'blockchain status',
      component: BlockchainStatus,
      beforeEnter: (to, from, next) => authHook(to, from, next, 'admin')
    }
  ],
  redirect: {
    '*': '/home'
  }
})

function authHook (to, from, next, role) {
  if (!auth.isAuthenticated(role)) {
    next(false)
    console.log('not authorized to access this page')
    this.$router.replace('/home')
  } else {
    next()
    console.log('next called on', role)
  }
}
