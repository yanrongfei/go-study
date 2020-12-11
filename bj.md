# go语言基础知识

# 学习资料
## 相关参考文档  
- 李文周的blog  
https://www.liwenzhou.com/archives/

- go sdk中文的文档  
https://studygolang.com/pkgdoc


## 学习视频  
- https://www.bilibili.com/video/BV16E411H7og?p=14  
- 尚硅谷韩顺平：  
https://www.bilibili.com/video/BV1kt411C7fK?from=search&seid=12793183262124070993  

# 基础知识

## 位运算
### 二进制
- 二进制计算都是以补码的方式进行计算  
- 正数的 原码，反码，补码 都是本身
- 负数的反码：符号位不变其他位取反，负数的补码：反码+1
### & | ^ 实际的作用




## printf
```
%d 数字类型  
%02d 两位数字
%s 字符串类型 
%x 返回16进制的数据
%T 返回数据类型
%v 返回数据的value
```

# go mod
## 开启mod  
- 修改配置文件  
```  
vim ~/.bash_profile
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```
- 相关命令  
```
go mod init # 初始化go.mod  
go mod tidy  # 更新依赖文件  
go mod download  # 下载依赖文件  
go mod vendor  # 将依赖转移至本地的vendor文件  
go mod edit  # 手动修改依赖文件  
go mod graph  # 打印依赖图  
go mod verify  # 校验依赖  
go list -m -u all #检查更新
go get -u need-upgrade-package #升级需要升级的包
go get -u  #升级所有的
```


# web框架
- 各种主流web框架的对比  https://zhuanlan.zhihu.com/p/86519607?utm_source=wechat_timeline

## gin
- github  
start 44K  
https://github.com/gin-gonic/gin   
周边建设  
https://github.com/hhstore/blog/issues/132

## beego
- github  
start 25.5k  
https://github.com/astaxie/beego  

## iris  
- github  
start 19.6K  
https://github.com/kataras/iris  




