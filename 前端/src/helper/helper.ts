import {ElMessageBox} from "element-plus";

export  function getData(data,keys,obj){
    let arr=keys.split('.')
    for (let i =0;i<arr.length-1;i++){
        obj=obj[arr[i]]
    }
    obj[arr[arr.length-1]]=data
}

export function data2arr(newData,oldData,obj){
    for (let i in newData){ //修改老的
        let find=false

        oldData.forEach(function (modify) {
            if(modify.key==i){
                find=true
                modify.value=newData[i].toString()
            }
        })

    }
    for (let key in newData){
        let find=true
        oldData.forEach(function (item) {
            if (item.key==key){
                find=false
            }
        })
        if(find==true){
            oldData.push({key:key.toString(),value:newData[key].toString(),name:"",required:false})
        }
    }

    oldData.forEach(function (item,i) {
        let find=false
        for (let i in newData){ //删除没有的
            if (item.key==i){
                find=true
            }
        }
        if(find==false){
            oldData.splice(i,1)
        }
    })
    obj.setData(oldData)
}
export   function parseData(str){
    //类似 echo -c 'echo The app is running! && sleep 3600'
    //单引号和双引号 里面的内容 不做 split

    if(!str){
        return []
    }
    try {
        str=str.replace(/^\s+|\s+$/gm,'')  //实现trim
        let pattern =/[\"|'](.*?)[\"|']/gi;
        let mList=str.match(pattern)
        if(!mList || mList.length===0){
            str=str.replace(/\s+/g,' ')
            return str.split(' ')
        }
        for(let i=0;i<mList.length;i++){
            str=str.replace(mList[i],'_#_'+i) // _#_1  _#_2  _#_3
        }
        str=str.replace(/\s+/g,' ')  //多个空格要变成一个 空格
        let ret=str.split(" ") //数组
        let p=/^_#_(\d+)$/
        ret.forEach((item,index)=>{
            if(p.test(item)){
                //要替换掉开头和结尾的 单引号或双引号
                ret[index]=mList[RegExp.$1].replace(/^[\"|']|[\"|']$/gm,'')
            }else{
                ret[index]=ret[index].replace(/^[\"|']|[\"|']$/gm,'')
            }
        })
        return ret
    }catch (e) {
        return []
    }


}
export function UnParseData(arr){
    let str=""
    for (let i=0;i<arr.length;i++){
        if(/\s+/g.test(arr[i])){
            str+=" '"+arr[i]+"'"
        }else{
            str+=" "+arr[i]
        }
    }
    return str
}

export function arrToMap(input){
    let obj={}
    for (let i=0;i<input.length;i++){
        obj[input[i].key]=input[i].value
    }
    return obj;
}
export function MapToArr(input){
    let arr=[]
    console.warn(input)
    for (let i in input){
        arr.push({key:i,value:input[i],NoRequired:!!input[i]})
    }
    return arr;
}
export function  rowToQuery(row){
    return {namespace:row.namespace,name:row.name}
}

export function copyData(state,tData){
    let T=tData.data.data
    state.List=T.Data
    state.pageInfo=T.Total
    state.current_page=T.Current
}
export function wsCopyData(state,tData,type){
    let T=tData.result;
    if(tData.type==type&&tData.ns==state.namespace){
        state.List=T.Data
        state.pageInfo=T.Total
        state.current_page=T.Current
    }
}
export  function rmTip(name){
    return ElMessageBox.confirm(
        '你确定继续删除这个'+name+'账户操作吗?',
        'Warning',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
}