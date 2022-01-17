import {ElLoading, ElMessage, ElMessageBox} from "element-plus";
import md5 from "js-md5";
import {nextTick} from "vue";
import {checkYaml} from "./rules";

export  function getData(data,keys,obj,doDelete){
    let arr=keys.split('.')
    for (let i =0;i<arr.length-1;i++){
        obj=obj[arr[i]]
    }

    if( hasValues(data)){
        obj[arr[arr.length-1]]=data
    }else{
        if (doDelete){
            delete obj[arr[arr.length-1]]
        }

    }
}

export  function hasValues(data){
    return data&&Object.values(data).length
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
    state.pageTotal=T.Total
    state.currentPage=T.Current
}
export function wsCopyData(state,tData,type){
    let T=tData.result;
    if(tData.type==type&&tData.ns==state.namespace){
        state.List=T.Data
        state.pageTotal=T.Total
        state.currentPage=T.Current
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

export  function createTip(){
    return ElMessageBox.confirm(
        '你确定要继续操作吗?',
        'Warning',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
}
function  convert(data){
    let temp
    try {
        temp= JSON.stringify(JSON.parse(JSON.stringify(data)))
    }catch (e) {
        console.log(e)
        temp=""
    }

   return temp
}
export function IsDirty(state,props,root,flag){

    if (root){
        if (state.md5==md5(convert(state.form))){
            return true
        }
    }

    if (!root){
        // if (flag==2.2){
        //     console.error(flag,props,state.form,state.md5,md5(convert(state.form)),md5(convert(props)),root)
        //     console.warn(state.md5==md5(convert(state.form)),state.md5==md5(convert(props)),"*****")
        // }
        if (state.md5==md5(convert(state.form))&&state.md5==md5(convert(props))){
           // console.log(flag,"不执行更新")
            return true
        }
    }
    // if (flag==2.2){
    //     console.log(flag,"执行更新")
    // }

    state.md5=md5(convert(state.form))
    return  false
}
let  loading


export  function CheckData(list,num){
    return new Promise(async function (resolve) {
        for (let i = 0; i < list.length; i++) {
            try {
                await list[i].value.Check()
                resolve(true)
            } catch (e) {
                console.log(e)
            }
        }
        resolve(true)
    })

}
export async  function  IsReady(CreateFn,state,msg,flag){
    if(flag){
        let result = await CreateFn(state.form)
        if (result.data.code == 200) {
            ElMessage(msg)
        }
        return
    }
    document.getElementById("errorReport").innerHTML=""
    setTimeout(async ()=>{

        state.hasErr=false
        let list=document.querySelectorAll('.el-form-item__error')
        if (!list.length) {
            try {
                let result = await CreateFn(state.form)
                if (result.data.code == 200) {
                  ElMessage(msg)
                }
            } catch (e) {
                console.log(e)
            }
        }else{
            for (let i=0;i<list.length;i++){
                let node=list[i].cloneNode(true)
                document.getElementById("errorReport").appendChild(node)
            }
            ElMessage.error("数据不合法")
            state.hasErr=true
        }
    },100)
}
export  function showError(msg){
    ElMessage.error(msg)
}

export  function uuid(len, radix) {
    let chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    let uuid = [], i;
    radix = radix || chars.length;

    if (len) {
        // Compact form
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random()*radix];
    } else {
        // rfc4122, version 4 form
        var r;

        // rfc4122 requires these characters
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';

        // Fill in random data.  At i==19 set the high bits of clock sequence as
        // per rfc4122, sec. 4.1.5
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random()*16;
                uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }

    return uuid.join('');
}

export  function initFn(yamlRef,nextTick,state){
    function Update() {
        if(state.mode!="json"){
            yamlRef.value.Update()
        }else{
            yamlRef.value.send()
        }
    }

    function yamlChange(data) {
        if (data) {
            let msg = checkYaml(data, state)
            console.warn(msg)
            if (msg) {
                showErr(msg)
                return
            }
            state.form = data
        }else{
            yamlRef.value.setData(state.form)
        }
    }

    function back() {
        state.mode = "yaml"
        Update()
    }
    function showErr(msg){
        loading=ElLoading.service({
            lock: true,
            text: '',
            spinner:"failed",
            background: 'rgba(0, 0, 0, 0.7)',
        })
        ElMessage({
            type: 'error',
            grouping: true,
            message:msg||"YAML内容有误,请仔细编辑",
            showClose:true,
            duration:0,
            onClose:function () {
                loading.close()
                back()
            }
        })
    }
    return {yamlChange,Update,back,showErr}
}
export function deleteOtherInfo(state){
    delete  state.form.metadata.managedFields
}