# Go-Note
Learn how to use go.

## 《GO 程序设计语言》

### CH1 入门

#### 1.1

[练习1.1](./gopl/ch1/1_1.go)

#### 1.2

[练习1.2](./gopl/ch1/1_2.go)

#### 1.3

[练习1.3](to do)

#### 1.4

[练习1.4](./gopl/ch1/1_4.go)

#### 1.5

[练习1.5](./gopl/ch1/1_5.go)

#### 1.6

[练习1.6](./gopl/ch1/1_6.go)

#### 1.7

[练习1.7](./gopl/ch1/1_7.go)

#### 1.8

[练习1.8](./gopl/ch1/1_8.go)

#### 1.9

[练习1.9](./gopl/ch1/1_9.go)

#### 1.10

[练习1.10](./gopl/ch1/1_10.go)，注意要修改代码，而且在同一个进程内连续访问，才能看到这种缓存的效果。

![IMG](./image/gopl/1_10.png)

#### 1.11

[练习1.11](./gopl/ch1/1_11.go)，会等待一段时间后超时错误，我这里是关上 vpn 然后测试 google。

![IMG](./image/gopl/1_11.png)

#### 1.12

[练习1.12](./gopl/ch1/1_12.go)，命令行输入 go run 1_12.go web & 后台运行服务器，然后浏览器 localhost:8000 访问即可。

![IMG](./image/gopl/1_12.png)

至于为什么我不用 localhost 是因为我用的是 wsl2，windows 下访问 wsl2 的子系统网络需要对应的 ip (可以在 ip addr | grep eth0 获取)，另外对应的 go 源程序也要把 localhost 改成对应的要监听的网卡 ip。

### CH2 程序结构

---

## 与C、C++的区别

1. map是基于哈希的，无需。另外和 c++ 不同的是 go 的 map\[key\] 这种下标访问并不会创建该 key 对应的元素，而 c++ map 的下标运算会。go 一般用 if _, ok := map\[key\]; ok == true 的方式得知存不存在元素。
