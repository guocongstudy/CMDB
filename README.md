# CMDB
小型项目，项目主要实现CMDB信息的录入与查询功能

## 涉及到的技能
1.go http标准库
2.第三方路由库：httprouter
3.go操作mysql

## 项目骨架介绍
### 项目组织的核心思路是
每个业务模块尽量独立，方便后期扩展和迁移成独立的服务。
<br/>cmd： 程序cli工具包
<br/>conf： 程序配置对象
<br/>protocol： 程序监听的协议
<br/>pkg： 业务领域包
    <br/>（&nbsp；）-host 
    <br/>（&nbsp；）-model：业务需要的数据模型
    <br/>（&nbsp；）-interface： 业务接口（领域方法）
    <br/>（&nbsp；）-impl： 业务具体实现
    <br/>（&nbsp；）...
<br/>main 程序入口文件    
