import Vue from 'vue'
import iView from 'iview'
import conf from '@/config/config'
Vue.prototype.conf= conf
import axios from '@/util/preRequests'
Vue.prototype.axios= axios
import util from '@/util/util'
Vue.prototype.util= util

import VueScroller from 'vue-scroller'

import 'iview/dist/styles/iview.css'    // 使用 CSS
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import NewsList from '@/components/NewsList'
import WaterFall from '@/components/WaterFall'

import Admin from '@/components/Admin'


Vue.use(Router)
Vue.use(iView)
Vue.use(VueScroller)

export default new Router({
    routes: [
        {
            path: '/hello',
            name: 'HelloWorld',
            component: HelloWorld
        },
        {
            path: '/',
            name: 'NewsList',
            component: NewsList,
            meta: {
                title: "快新闻"
            }
        },
        {
            path: '/WaterFall',
            name: 'WaterFall',
            component: WaterFall,
            meta: {
                title: "快新闻"
            }
        },
        {
            path: '/Admin',
            name: 'Admin',
            component: Admin,
            meta: {
                title: "快新闻后台管理系统"
            }
        },
    ]
})
