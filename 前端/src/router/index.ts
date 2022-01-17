import {createRouter, createWebHashHistory, RouteRecordRaw, useRouter} from 'vue-router';


//
// import deploymentList from "../views/deployment-list.vue";
// import pod from "../views/pod.vue";
// import podLogs from "../views/pod-logs.vue";
// import podShell from "../views/pod-shell.vue";
//
//
//
//
//
//
//
//
//
//
// import nodeList from "../views/node.vue";
// import nodeUpdate from "../views/node-update.vue"
// import  NodeShell from  "../views/node-shell.vue"
//
//
//
// import  rbacRoleList from  "../views/rbac-role-list.vue"
// import  rbacRoleCreate from  "../views/rbac-role-create.vue"
// import  rbacRoleUpdate from  "../views/rbac-role-update.vue"
//
// import  rbacRoleBindingList from  "../views/rbac-roleBinding-list.vue"
// import  rbacRoleBindingCreate from  "../views/rbac-roleBinding-create.vue"
//
//
//
// import  rbacClusterRoleList from  "../views/rbac-clusterRole-list.vue"
// import  rbacClusterRoleCreate from  "../views/rbac-clusterRole-create.vue"
// import  rbacClusterRoleUpdate from  "../views/rbac-clusterRole-update.vue"
//
//
// import  rbacClusterRoleBindingList from  "../views/rbac-clusterRoleBinding-list.vue"
// import  rbacClusterRoleBindingCreate from  "../views/rbac-clusterRoleBinding-create.vue"
// import rbacUaList from "../views/rbac-ua-list.vue"
//
// ///上面都需要改的
//
 //import deployCreate from "../views/deployments/deploy-create.vue"
//
//
// //修改过后
 import nsList from "../views/namespace/ns-list.vue"
 import login from "../views/common/login.vue";
//
import  saList from "../views/sa/sa-list.vue"
import  saDetail from "../views/sa/sa-detail.vue"
import  saUpdate from "../views/sa/sa-update.vue"
import  saCreate from "../views/sa/sa-create.vue"
//
//
import  secretList from "../views/secret/secret-list.vue"
// import  secretCreate from "../views/secret/secret-create.vue"
// import  secretDetail from "../views/secret/secret-detail.vue"
// import  secretUpdate from "../views/secret/secret-update.vue"
//
//
// import  configmapList from "../views/configmap/configmap-list.vue"
// import  configmapCreate from "../views/configmap/configmap-create.vue"
// import  configmapDetail from "../views/configmap/configmap-detail.vue"
// import  configmapUpdate from "../views/configmap/configmap-update.vue"
//
// import ingressList from "../views/ingress/ingress-list.vue"
// import ingressCreate from "../views/ingress/ingress-create.vue"
// import ingressDetail from "../views/ingress/ingress-detail.vue"
// import ingressUpdate  from "../views/ingress/ingress-update.vue"
//
//
// import svcList from "../views/svc/svc-list.vue"
// import svcCreate from "../views/svc/svc-create.vue"
// import svcDetail from "../views/svc/svc-detail.vue"
// import svcUpdate  from "../views/svc/svc-update.vue"
//secret-update
const routes:RouteRecordRaw[]= [
    {
        path:"/ns-list",
        name: 'ns-list',
        component: nsList
    },
    {
        path: '/',
        name: 'login',
        component: login
    },

    {
        path: '/:pathMatch(.*)',
        redirect: '/'
    },

    {
        path: '/sa-list',
        name: 'sa-list',
        component: saList
    },
    {
        path: '/sa-detail',
        name: 'sa-detail',
        component: saDetail
    },
    {
        path: '/sa-update',
        name: 'sa-update',
        component: saUpdate
    },
    {
        path: '/sa-create',
        name: 'sa-create',
        component: saCreate
    },
    {
        path: '/secret-list',
        name: 'secret-list',
        component: secretList
    },
    //=========需要重构的============
    // {
    //     path: '/deploy-create',
    //     name: 'deploy-create',
    //     component: deployCreate
    // },

    // {
    //     path: '/secret-detail',
    //     name: 'secret-detail',
    //     component: secretDetail
    // },

    // {
    //     path: '/secret-create',
    //     name: 'secret-create',
    //     component: secretCreate
    // },
    // {
    //     path: '/secret-update',
    //     name: 'secret-update',
    //     component: secretUpdate
    // },
    // {
    //     path: '/configmap-update',
    //     name: 'configmap-update',
    //     component: configmapUpdate
    // },
    // {
    //     path: '/configmap-detail',
    //     name: 'configmap-detail',
    //     component: configmapDetail
    // },
    // {
    //     path: '/configmap-create',
    //     name: 'configmap-create',
    //     component: configmapCreate
    // },
    // {
    //     path: '/configmap-list',
    //     name: 'configmap-list',
    //     component: configmapList
    // },
    // {
    //     path: '/svc-list',
    //     name: 'svc-list',
    //     component: svcList
    // },
    // {
    //     path: '/svc-create',
    //     name: 'svc-create',
    //     component: svcCreate
    // },
    // {
    //     path: '/svc-detail',
    //     name: 'svc-detail',
    //     component: svcDetail
    // },
    // {
    //     path: '/svc-update',
    //     name: 'svc-update',
    //     component: svcUpdate
    // },
    // {
    //     path: '/ingress-list',
    //     name: 'ingress-list',
    //     component: ingressList
    // },
    // {
    //     path: '/ingress-create',
    //     name: 'ingress-create',
    //     component: ingressCreate
    // },
    // {
    //     path: '/ingress-detail',
    //     name: 'ingress-detail',
    //     component: ingressDetail
    // },
    // {
    //     path: '/ingress-update',
    //     name: 'ingress-update',
    //     component: ingressUpdate
    // },
    //
    //
    //
    //
    //
    //
    //
    // {
    //     path: '/rbac-clusterRole-create',
    //     name: 'rbac-clusterRole-create',
    //     component: rbacClusterRoleCreate
    // },
    // {
    //     path: '/rbac-roleBinding-list',
    //     name: 'rbac-roleBinding-list',
    //     component: rbacRoleBindingList
    // },
    // {
    //     path: '/rbac-roleBinding-create',
    //     name: 'rbac-roleBinding-create',
    //     component: rbacRoleBindingCreate
    // },
    // {
    //     path: '/rbac-ua-list',
    //     name: 'rbac-ua-list',
    //     component: rbacUaList
    // },
    // // {
    // //     path: '/rbac-sa-list',
    // //     name: 'rbac-sa-list',
    // //     component: rbacSaList
    // // },
    // {
    //     path: '/rbac-clusterRole-update',
    //     name: 'rbac-clusterRole-update',
    //     component: rbacClusterRoleUpdate
    // },
    // {
    //     path: '/rbac-role-update',
    //     name: 'rbac-role-update',
    //     component: rbacRoleUpdate
    // },
    // {
    //     path:'/rbac-clusterRole-list',
    //     name: 'rbac-clusterRole-list',
    //     component: rbacClusterRoleList
    // },
    // {
    //     path:'/rbac-clusterRoleBinding-create',
    //     name: 'rbac-clusterRoleBinding-create',
    //     component: rbacClusterRoleBindingCreate
    // },
    // {
    //     path: '/rbac-clusterRoleBinding-list',
    //     name: 'rbac-clusterRoleBinding-list',
    //     component: rbacClusterRoleBindingList
    // },
    // {
    //     path: '/rbac-roleBinding-create',
    //     name: 'rbac-roleBinding-create',
    //     component: rbacRoleBindingCreate
    // },
    //
    // {
    //     path: '/node-update',
    //     name: 'node-update',
    //     component: nodeUpdate
    // },
    // {
    //     path: '/rbac-role-list',
    //     name: 'rbac-role-list',
    //     component: rbacRoleList
    // },
    // {
    //     path: '/rbac-role-create',
    //     name: 'rbac-role-create',
    //     component: rbacRoleCreate
    // },
    // {
    //     path: '/node-shell',
    //     name: 'node-shell',
    //     component: NodeShell
    // },
    // {
    //     path: '/node-list',
    //     name: 'node-list',
    //     component: nodeList
    // },
    // {
    //     path: '/pod-shell',
    //     name: 'pod-shell',
    //     component: podShell
    // },
    // {
    //     path: '/pod-logs',
    //     name: 'pod-logs',
    //     component: podLogs
    // },
    //
    // {
    //     path: '/deployment-list',
    //     name: 'deployment-list',
    //     component: deploymentList
    // },
    //
    // {
    //     path: '/secret-create',
    //     name: 'secret-create',
    //     component: secretCreate
    // },
    //
    //
    // {
    //     path: '/pod',
    //     name: 'pod',
    //     component: pod
    // },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export function doTo(name:string,query:any={}){
    router.push({name: name,query:query  })
}

export default router;
