# 随着自己淡出游戏行业，接触到更多互联网行业的东西，发现好多好多需要优化的地方，暂时这个框架停止更新，有时间会重新更新，可能会用更多的新学到的知识去优化。

# Single Server Framework for Game

借鉴gonet1.0，使用mysql，redis实现的基于tcp长连接的单服设计的游戏服务器。


服务器架构
ResourceSrv
LoginSrv
PortalSrv
GameSrv

A. ResourceSrv基本流程
1. ResourceSrv切换API
    Method: POST
    Url: resource/v1/switch
    Params: channel_id version_id   

    描述: 客户端获取指定ResourceSrv入口信息
    配置表: resource_info_data.csv

2. ResourceSrv详情API
	Method: POST
	Url: resource/v1/detail
	Params: channel_id platform_type
	
	描述: 获取LoginSrv信息,CDN资源服务器地址,当前服务器支持的兼容版本,以及当前平台下热更资源信息
	配置表: channel_info_data.csv version_info_data.csv

3. PortalSrv列表API
	Method: POST
	Url: resource/v1/portals
	Params: 空
	
	描述: 获取PortalSrv列表
	配置表: portal_info_data.csv
	注: 此处可以配置指向统一组GameSrv的多个PortalSrv或者指向PortalSrv负载均衡Addr,实现单服设计
	    也可以配置指向不同组GameSrv的多个PortalSrv,实现多服/滚服设计

