# 第四天[2022-03-09]

实现一个restful服务

restful实际上是一种由接口规范，主要是把接口要操作的东西抽象为一种资源，然后对其操作。比如这里，目的是操作专辑album，那专辑就是一种资源。

一般规范就是:

1. 获取专辑列表: GET /albums
1. 获取一张专辑的详情：GET /albums/:id 
1. 添加一张专辑：POST /albums  body: album
1. 删除一张专辑：DELETE /albums/:id
1. 修改一张专辑：PUT /albums/:id  body: album
