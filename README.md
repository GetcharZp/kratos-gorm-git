# Kratos gorm git

> 使用kratos+gorm搭建的Git代码托管平台
> 
> kratos 参考文档：https://go-kratos.dev/
> 
> gorm 参考文档：https://gorm.io/zh_CN/docs/

## 同类产品

1. [Github](https://github.com/) 全球最大的代码托管平台
2. [Gitee](https://gitee.com/) 国内最大的代码托管平台
3. [Gitea](https://gitea.io/) 轻量级的开源的代码托管平台

## 安装
1. 安装 cli 工具
```shell
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest 
```

2. 初始项目
```shell 
kratos new kratos-gorm-git
```

3. 运行
```shell
kratos run
```

## 系统模块

- [ ] 仓库管理
- [ ] 用户管理
- [ ] 组织管理

## 快速体验GIT远程仓库

1. 初始化空的存储仓库
```shell
git init --bare /root/git-test/hello.git 
```

2.生成ssh密钥
```
ssh-keygen -t rsa -C "get@qq.com"
```

3. 将客户端生成的公钥复制到服务器端
```shell
# 公钥文件地址
cat ~/.ssh/id_rsa.pub
# 服务端的配置文件路径
vi ~/.ssh/authorized_keys
```

4. 操作远程仓库
```shell
# clone
git clone root@119.27.164.148:/root/git-test/hello.git
# 添加远程仓库
git remote add origin root@119.27.164.148:/root/git-test/hello.git
# 推送本地代码到远程仓库
git push -u origin master
```