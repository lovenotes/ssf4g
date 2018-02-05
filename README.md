# Single Server Framework for Game

借鉴gonet1.0，使用mysql，redis实现的基于tcp长连接的单服设计的游戏服务器。


服务器架构
ResourceSrv
LoginSrv
PortalSrv
GameSrv

A. ResourceSrv基本流程
1. ResourceSrv的resource_info_data.csv提供所有ResourceSrv信息
客户端通过channel_id和version_id获取正确的ResourceSrv入口信息

2. channel_info_data.csv,version_info_data.csv提供LoginSrv信息
客户端通过channel_id获取LoginSrv信息,CDN资源服务器地址,以及当前服务器支持的兼容版本

3. portal_info_data.csv提供PortalSrv信息

