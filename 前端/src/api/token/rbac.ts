
import request from "./read";

export  function  getRoleList(ns,page){
    console.log(ns)
    return request.get("/v1/roles",{params:{ns,current:page}})
}
export  function  getClusterRoles(page){
    return request.get("/v1/clusterRoles",{params:{current:page}})
}
export  function getClusterRoleDetail(name){
    return request.get("/v1/GetClusterRoles",{params:{name}})
}
export  function clusterRolesDel(name){
    return request.delete("/v1/clusterRoles",{params:{name}})
}

export  function  getRoleAllList(ns){
    console.log(ns)
    return request.get("/v1/rolesAll",{params:{ns}})
}

export  function  getClusterRoleAllList(){

    return request.get("/v1/clusterRolesAll")
}
export  function  roleCreate(data){
    console.log(data)
    return request.post("/v1/roles",data)
}
export  function  clusterRoleCreate(data){
    console.log(data)
    return request.post("/v1/clusterRoles",data)
}

export  function roleDel(ns,name){
    return request.delete("/v1/roles",{params:{ns,name}})
}
export  function getRoleBindingList(ns,page){
    return request.get("/v1/roleBindings",{params:{ns,current:page}})
}

export  function  roleBindingCreate(data){
    console.log(data)
    return request.post("/v1/roleBindings",data)
}


export  function  roleBindingDel(ns,name){
    return request.delete("/v1/roleBindings",{params:{ns,name}})
}
export function deleteUserFromBinding(ns,name,data){
   return  request.put("/v1/roleBindings",data,{params:{ns,name,type:1}})
}
export function addUserToBinding(ns,name,data){
    return  request.put("/v1/roleBindings",data,{params:{ns,name}})
}
export function postRoleBindings(data) {
    return request.post("/v1/roleBindings", data)

}

export  function getRoleDetail(ns,name){
    return request.get("/v1/GetRoles",{params:{ns,name}})
}
export function roleUpdate(ns,name,data){
    return  request.put("/v1/roles",data,{params:{ns,name}})
}

export function clusterRoleUpdate(name,data){
    return  request.put("/v1/clusterRoles",data,{params:{name}})
}

export function ClusterRoleBindingList(page){
    return  request.get("/v1/clusterRoleBindings",{params:{current:page}})
}


export function addClusterRoleBindings(name,data){
    return  request.put("/v1/clusterRoleBindings",data,{params:{name}})
}
export  function  ClusterRoleBindingDel(name){
    return request.delete("/v1/clusterRoleBindings",{params:{name}})
}

export function deleteClusterRoleBinding(name,data){
    return  request.put("/v1/clusterRoleBindings",data,{params:{name,type:1}})
}

export function createClusterRoleBinding(data){
    return  request.post("/v1/clusterRoleBindings",data)
}