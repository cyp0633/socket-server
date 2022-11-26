# Socket Server

The server part of Computer Networking course project, a chat software based on socket programming.

计算机网络课程项目的服务端部分，基于 socket 编程的聊天软件。

## Usage

服务器监听 65432/tcp 端口，你也可以使用 Telnet 连接。

要获得在线客户端列表，可以发送 `HELO`。正常情况下返回 `CLIENTS <client1_ip> <client2_ip> ...`。

要发送消息，可以使用 `SEND <dest_ip> MSG <message>`。正常情况下返回 `OK`。

要获取用户列表，可以使用 `USER` 命令。正常情况下返回：

```plain
<user_num> USERS
<ip1>
<ip2>
...
```

要接收消息，可以使用 `PULL`。正常情况下返回：

```plain
LEN <length>
FROM <from_ip1> CONTENT <content1>
FROM <from_ip2> CONTENT <content2>
FROM <from_ip2> CONTENT <content2>
...
```

此外，还可以通过 UDP 广播获得服务端地址。服务器监听 65433/udp 端口，你可以发送 `PROBE`，服务器会返回 `HERE` 到 65433/udp 端口。

## Known Issues

好像对 IPv6 支持还有点问题。
