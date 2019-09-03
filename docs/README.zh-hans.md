# Trove
[English](../README.md)

Trove 是 Go 语言的包依赖管理器。使用最新 `v1.12.9` SDK 编译，最低需求版本未测试。

Trove 因属于实验性项目，没有资金去建立单独的包管理服务，目前采用Git仓库式管理，测试兼容GitHub仓库，理论上支持所有GIT代码仓库（未测试），省却单独服务器的开销；

Trove 可随意引入私有仓库代码包（需要拥有仓库克隆权限），随意引用私有仓库代码无需额外配置与搭建私有服务器。

Trove 基于 Git ,所以请确保您的运行环境已安装Git，并可全局使用git命令
## 官网
https://trove.daohang.dev
## 开发语言
既然是 Go 的包管理工具，自然用 Go 语言来开发
## 开源协议
基于 [木兰许可协议](http://license.coscl.org.cn/MulanPSL) 发布开源

## 使用方式
#### 初始化
    trove init
#### 引用依赖
    trove require https://github.com/XXXXXX
        默认为 commit 版本控制 等同于如下
        trove require commit@https://github.com/XXXXXX
        [commitId] 可选 commitId 版本，默认为最新版
        
    可设置为tag标签控制版本
    trove require tag@https://github.com/XXXXXX [tag]
        [tag] 可选tag版本，默认为最新版
#### 移除依赖
    trove remove [packageName]
    [packageName] 包名
#### 查看依赖列表
    trove --list
    可查看项目直接依赖列表
    trove --list--all
    可查看项目所有依赖列表，包括直接依赖和间接依赖
#### 安装所有依赖包
    trove install [packageName]
    手动编辑 trove.json 增加依赖时，可使用该命令安装未下载的依赖包，或在初始化项目时下载所有的依赖包
    [packageName] 可选安装单独依赖包
#### 更新依赖包
    trove update [packageName]
    更新所有依赖包到版本限定范围内的最新版本
    [packageName] 可选仅更新指定包
#### Trove 版本
    trove -V,--version
    查看 Trove 版本号及最后修订时间
#### 帮助
    trove -h,--help
    查看 Trove 帮助信息
## 开发团队
>排名不分先后

[Yanlong-li](https://github.com/yanlong-li)

## 更新日志
    2019年9月3日
    增加 init 命令进行项目的初始化
    2019年9月3日
    优化版本控制结构
    修复更新单个包时未切换到指定版本问题

    2019年9月2日
    新增git版本控制支持Tag tag@https://github.com/XXXXX/XXXX 默认为commit commit@https://github.com/XXXX
    优化git@ssh引入包
    
    2019年9月2日
    强化命令行
    新增require命令引入git包
    新增install安装不存在的包
    新增update更新包命令
    新增list显示所有引入的包
    新增-v,--version输出版本号
    新增-h,--help 输出帮助提示
    新增 remove 移除包