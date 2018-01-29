// import sha256 from 'js-sha256'
import conf from '@/config/config'

/**
 * 生成签名 sign
 * @param obj
 * @returns {Object}
 */
// function getSign(obj) {
//     var timestamp = Date.parse(new Date());
//     // var timestamp = 1499675521446;
//     obj.timestamp = timestamp;
//     obj.appid = conf.appid;
//
//     // 获取所有的key
//     var keys = Object.keys(obj).map(item=>{return item.toLocaleLowerCase()}).sort();
//
//     // console.log(obj)
//     // 排序后的对象
//     var obj_tmp = new Object();
//     for (var i in keys) {
//         for (var j in obj) {
//             // console.log(obj, j)
//             if (j.toLocaleLowerCase() == keys[i]) {
//                 obj_tmp[j] = obj[j];
//             }
//         }
//     }
//
//     // 生成sign
//     var tmp = [];
//     for (var i in obj_tmp) {
//         tmp.push(i + "=" + obj_tmp[i]);
//     }
//     // baad57fb839190b90eb3f10a21b9fe9509cc24d6d9e502747694710129d067f2
//     // 原: appid=1&iccid=898606182222832823&timestamp=1499675521446dla20op8ui0eujflajsdf
//     // 现: appid=1&iccid=898606182222832823&sign=5b4d262efc4712b3ead4dd6e0115ecb2d3f0e9614cb53e7d8e6cf2c10940a005&timestamp=1499675521446dla20op8ui0eujflajsdf
//     // console.log(tmp.join("&") + conf.secret)
//     obj_tmp.sign = sha256.sha256(tmp.join("&") + conf.secret);
//
//     return obj_tmp;
// }

function checkResponse(obj) {
    this.$Message.error(obj)
}

export default {
    // getSign,
    checkResponse
}