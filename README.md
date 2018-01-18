# 模拟自定义协议通讯工具 Peenet


![golang](https://ss0.bdstatic.com/94oJfD_bAAcT8t7mm9GUKT-xh_/timg?image&quality=100&size=b4000_4000&sec=1507515058&di=e1760b51b338fa571de9fb72545b1040&src=http://static.open-open.com/news/uploadImg/20151214/20151214234004_732.jpg "Don")



## golang language
## 支持多协议通讯


名词解释
    
    deviceId 设备终端码 唯一标识
    userId      用户id 唯一标识
        
    



内存存储详情
    
    user:id  用户自增id
    user:detail:<userId> 存储登录用户的详情 username|时间戳
    
    device:detail:<deviceId> 设备详情  协议类型|userId
    
    user:count 用户在线情况
    
    ws:work:device:<wsworkId> 存储websocket sessionId和设备的对应关系
    
    

`