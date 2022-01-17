

export function requireRules(msg){
    return {required: true, trigger: 'change', message:msg }
}
function isInArr(arr,msg){
    return (rule, value, callback) => {
        if(arr.indexOf(value)==-1){
            callback(new Error(msg))
        }
        callback(null)
    }
}
export function inArrayWithMsg(arr,msg){
    return  { validator:isInArr( arr,msg), trigger: 'change' }
}
export  function NameToArr(item){
    let arr=[]
    for (let i=0;i<item.length;i++){
        arr.push(item[i].name)
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
function  getDataFormKeys(keys,obj){
    let keysList=keys.split('.')

    let needModify= obj
    for (let i=0;i<keysList.length-1;i++){
        needModify=obj[keysList[i]]
    }
    return  needModify[keysList[keysList.length-1]]
}
