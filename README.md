Go 语言包管理器 Trove
=
> 这是一个Go语言包管理器，希望你能发现其中的宝藏。

## 官网
https://trove.daohang.dev
## 开发语言
既然是 Go 的包管理工具，自然用 Go 语言来开发
## 开源协议
基于 [木兰许可协议](http://license.coscl.org.cn/MulanPSL) 发布开源
## 开发团队
暂不公开

更新日志

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