
var oss = require('ali-oss');
var fs = require('fs');
var client = oss({
    accessKeyId: 'LTAI5tGKYSobyanQcarHqrnc',
    accessKeySecret: '6mS9IlDkXP6WxLz6UMOnfTE7HQxsiW',
    bucket: 'wikecloud',
    region: 'oss-cn-hangzhou'
});




async function putOSS (src,dist) {
    try {
        let result = await client.put(dist, src);
        console.log(result);
    } catch (e) {
        console.log(e);
    }
}
function addFileToOSSSync(src,dist){
    var docs = fs.readdirSync(src);
    docs.forEach(function(doc){
        var _src = src + '/' + doc,
            _dist = dist + '/' + doc;
        var st = fs.statSync( _src);
        // 判断是否为文件
        if( st.isFile()&&doc!=='.DS_Store' ){
            putOSS(_src,_dist);
            // console.log(_src+'是文件',_dist)
        }
        // 如果是目录则递归调用自身
        else if( st.isDirectory() ){
            // console.log(_src+'是文件夹')
            addFileToOSSSync( _src,_dist);
        }
    })
}
addFileToOSSSync('./dist',"/test")


let sourceFile = "./dist/index.html"
let destPath ="../api/html/index.html"

let readStream = fs.createReadStream(sourceFile);
let writeStream = fs.createWriteStream(destPath);
readStream.pipe(writeStream);
console.log("移动完成")