# Go-Note
Learn how to use go.

## 《GO 程序设计语言》

### CH1 入门

[练习1.1](./gopl/ch1/1_1.go)

[练习1.2](./gopl/ch1/1_2.go)

[练习1.3](to do)

[练习1.4](./gopl/ch1/1_4.go)

[练习1.5](./gopl/ch1/1_5.go)

[练习1.6](./gopl/ch1/1_6.go)

[练习1.7](./gopl/ch1/1_7.go)

[练习1.8](./gopl/ch1/1_8.go)

[练习1.9](./gopl/ch1/1_9.go)

### CH2 程序结构

---

## 与C、C++的区别

1. map是基于哈希的，无需。另外和 c++ 不同的是 go 的 map\[key\] 这种下标访问并不会创建该 key 对应的元素，而 c++ map 的下标运算会。go 一般用 if _, ok := map\[key\]; ok == true 的方式得知存不存在元素。
