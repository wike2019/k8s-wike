
import request from "../read";
import {doTo} from "../../../router";
//带分页列表
export  function  getSaList(ns,page){
    return request.get("/v1/sa/page",{params:{ns,page}})
}
//删除
export  function  SaDel(ns,name){
  return request.delete("/v1/sa",{params:{ns,name}}).then(()=>{
      doTo("sa-list")
  })
}
//创建
export  function  SACreate(data){
    return request.post("/v1/sa",data)
}
//更新
export  function  SAUpdate(data){
    return request.put("/v1/sa",data)
}
//所有列表
export  function  getSaListAll(ns){
    return request.get("/v1/sa",{params:{ns}})
}
//获取详细信息
export  function  getSaItem(ns,name){
    return request.get("/v1/sa/item",{params:{ns,name}})
}