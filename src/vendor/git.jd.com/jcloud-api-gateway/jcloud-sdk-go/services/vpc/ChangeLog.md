# 更新历史 #
API版本：0.6.0-beta

|发布时间|版本号|更新|说明|
|---|---|---|---|
|2019-01-17|0.6.0-beta|增加新接口|* 增加AddRoutePropagation<br>* 增加ModifyRoutePropagation<br>* 增加RemoveRoutePropagation<br>* 修改AddRouteTableRules，增加description属性<br>* 修改ModifyRouteTableRules，增加description属性<br>* 修改DescribeRouteTable/DescribeRouteTables返回信息增加路由传播列表（routePropagations）、路由信息增加路由类型、description属性<br>* 查询和修改配额接口支持传播路由|
|2018-12-06|0.5.1|增加新接口|* 增加修改弹性IP接口|
|2018-09-04|0.5.0|增加新接口|* 增加修改私有网络接口<br>* 增加创建、删除、修改子网接口<br>* 增加创建、删除、修改networkAcl接口<br>* 增加查询ACL列表和详情接口<br>* 增加给子网绑定/解绑networkAcl接口<br>* 增加添加、移除、修改networkAcl规则接口<br>* 增加创建、删除、修改路由表接口<br>* 增加查询路由表列表和详情接口<br>* 增加路由表绑定/解绑子网接口<br>* 增加添加、移除、修改路由表规则接口|
|2018-08-20|0.4.1|扩展支持|* Java SDK新增对Android系统的支持|
|2018-08-16|0.4.0|增加新接口|* 增加创建私有网络接口<br>* 增加删除私有网络接口|
|2018-07-26|0.3.0|增加新接口|* 增加查询弹性网卡列表接口<br>* 增加创建网卡接口，只能创建辅助网卡接口<br>* 增加查询弹性网卡信息详情接口<br>* 增加删除弹性网卡接口<br>* 增加修改弹性网卡接口|
|2017-11-31|0.1.0|初始版本|* VPC基本操作接口|
