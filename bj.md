# go基础知识

# 学习资料
- 相关参考文档  
https://www.liwenzhou.com/archives/

- 学习视频  
https://www.bilibili.com/video/BV16E411H7og?p=14  
尚硅谷韩顺平：  
https://www.bilibili.com/video/BV1kt411C7fK?from=search&seid=12793183262124070993  



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

## beego
- github  
start 25.5k  
https://github.com/astaxie/beego  

## iris  
- github  
start 19.6K  
https://github.com/kataras/iris  




