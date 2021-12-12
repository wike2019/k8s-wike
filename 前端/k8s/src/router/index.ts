import {createRouter, createWebHashHistory, RouteRecordRaw, useRouter} from 'vue-router';

import login from "../views/login/login.vue";  //登录

const routes:RouteRecordRaw[]= [
    {
        path: '/',
        name: 'login',
        component: login
    },
    {
        path: '/:pathMatch(.*)',
        redirect: '/'
    }

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export function go(name:string,query:any={}){
    router.push({name: name,query:query  })
}

export default router;
