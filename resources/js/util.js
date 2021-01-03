
/** 获取范围内随机数 */
function getRoundRandom(max, min) {
    return Math.ceil(Math.random() * (max - min) + min);
}

/** 获取小于等于x随机数 */
function getRandom(max) {
    return Math.ceil(Math.random() * max);
}
/** 转换成json */
function toJSON(msg) {
    return JSON.stringify(msg);
}
/** 转换成map */
function toMap(str) {
    return JSON.parse(str);
}
/** 获取消息的map */
function getMsgMap(gameId, msgId, data) {
    var map = new Map();
    map.set("gameId", gameId);
    map.set("msgId", msgId);
    map.set("data", data);
    var msg = Object.create(null);
    for (let[k,v] of map) {
        msg[k] = v;
    }
    return msg;
}