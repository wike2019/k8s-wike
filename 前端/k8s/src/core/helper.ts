export function str2utf8(str) {
    let encoder = new TextEncoder("utf8");
    return encoder.encode(str);
}
export function fixedEncodeURIComponent(str) {
    return encodeURIComponent(str).replace(/[!'()*]/g, function(c) {
        return '%' + c.charCodeAt(0).toString(16);
    });
}
export function toUtf8(str) {
    var out, i, len, c;
    out = "";
    len = str.length;
    for(i = 0; i < len; i++) {
        c = str.charCodeAt(i);
        if ((c >= 0x0001) && (c <= 0x007F)) {
            out += str.charAt(i);
        } else if (c > 0x07FF) {
            out += String.fromCharCode(0xE0 | ((c >> 12) & 0x0F));
            out += String.fromCharCode(0x80 | ((c >>  6) & 0x3F));
            out += String.fromCharCode(0x80 | ((c >>  0) & 0x3F));
        } else {
            out += String.fromCharCode(0xC0 | ((c >>  6) & 0x1F));
            out += String.fromCharCode(0x80 | ((c >>  0) & 0x3F));
        }
    }
    return out;
}
export function buf2str(buffer) {
    let encodedString = String.fromCodePoint.apply(null, new Uint8Array(buffer));
    encodedString=escape(encodedString)
    console.log(decodeURIComponent(encodedString))
        return    decodeURIComponent(encodedString);//没有这一步中文会乱码
}
export function fileDownLoad(data,name){
    let pom = document.createElement('a');
    pom.setAttribute('href', 'data:text/plain;charset=UTF-8,' + encodeURIComponent(data));
    pom.setAttribute('download', name);
    pom.style.display = 'none';
    if (document.createEvent) {
        const event = document.createEvent('MouseEvents');
        event.initEvent('click', true, true);
        pom.dispatchEvent(event);
    } else {
        pom.click();
    }
}
export    function isInArr(arr,msg){
    return (rule, value, callback) => {
        if(arr.indexOf(value)==-1){
            callback(new Error(msg))
        }
        callback()
    }
}

export   function parseData(str){
    //类似 echo -c 'echo The app is running! && sleep 3600'
    //单引号和双引号 里面的内容 不做 split
    if(str==""){
        return []
    }
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
    console.log(ret,"^^^^^^^^")
    return ret
}