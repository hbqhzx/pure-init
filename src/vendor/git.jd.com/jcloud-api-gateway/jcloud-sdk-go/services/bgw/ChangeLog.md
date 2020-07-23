# 更新历史 #
API版本：1.1.0-beta

|发布时间|版本号|更新|说明|
|---|---|---|---|
|2019-01-17|1.1.0-beta|增加新接口|* 增加CreateVpcAttachment<br>* 增加DeleteVpcAttachment<br>* 增加DescribeVpcAttachment和DescribeVpcAttachments接口<br>* 增加ModifyVpcAttachment<br>* 修改DescribeBgws和DescribeBgw接口，返回 Bgw上vpc接口Id列表<br>* 修改CreateBgwRoutes接口，增加、修改bgwRouteType、origin、description等信息<br>* 修改ModifyBgwRoutes接口，增加、修改bgwRouteType、description等信息<br>* 修改DescribeBgwRoutes接口，传播路由只展示优选路由，增加查询过滤条件，返回静态、动态、未生效路由数及新增routes属性等<br>* 查询和修改配额接口支持vpcAttachment、bgwRouter(传播)|
|2018-11-26|1.0.0|增加新接口|* 增加边界网关的创建、修改、删除、查询列表、查询详情接口<br>* 增加边界网关路由的创建、修改、删除、查询列表接口<br>* 增加私有虚拟连接的创建、修改、删除、查询列表、查询详情接口<br>* 增加连接的创建、修改、删除、查询列表、查询详情接口<br>* 增加合作伙伴查询连接资源列表接口<br>* 增加合作伙伴修改连接状态接口<br>* 增加查询Location列表、详情接口<br>* 增加查询合作伙伴列表、详情接口<br>* 增加查询配额接口|
