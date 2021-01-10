# go 语言编程之网络编程

## Socket 编程

以前socket编程步骤：

    1. 新建socket
    2. 绑定socket
    3. listen 监听端口
    4. accept 接收连接conn
    5. read 读取数据包，或者write 发送包
    
go语言提供一个net.Dial() 获得conn，使用read 和 write 接收数据和发送数据。

## RPC

go 自带的net/rpc，同步call，异步Go  