
//给form 表单验证规则,动态填充msg
export function requireRules(msg){
    return {required: true, trigger: 'change', message:msg }
}
//判定是否值在数组中
function isInArr(arr,msg){
    return (rule, value, callback) => {
        if(arr.indexOf(value)==-1){
            callback(new Error(msg))
        }
        callback(null)
    }
}
// 给form 表单验证规则
export function inArrayWithMsg(arr,msg){
    return  { validator:isInArr( arr,msg), trigger: 'change' }
}
// 对象名称转数组
export  function NameToArr(item){
    let arr=[]
    for (let i=0;i<item.length;i++){
        arr.push(item[i].name)
    }
    return arr
}
// 对象名称转数组
export  function NameToArrWithKey(item,key){
    let arr=[]
    for (let i=0;i<item.length;i++){
        arr.push(item[i][key].metadata.name)
    }
    return arr
}
export  function checkYaml(data,state){
    for (let i=0;i<state.notModifyList.length;i++){
        if (state[state.notModifyList[i].key]!=getDataFormKeys(state.notModifyList[i].modify,data)){
            return  state.notModifyList[i].msg
        }
    }
    return ""
}

//递归取得对象值
function  getDataFormKeys(keys,obj){
    let keysList=keys.split('.')

    let needModify= obj
    for (let i=0;i<keysList.length-1;i++){
        needModify=obj[keysList[i]]
    }
    return  needModify[keysList[keysList.length-1]]
}
