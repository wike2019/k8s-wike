

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
    if(data.apiVersion!==state.apiVersion|| data.Kind!=state.Kind){
        state.form.apiVersion=state.apiVersion
        state.form.Kind=state.Kind
        return 'apiVersion和Kind不允许修改'
    }
    return ""
}