# gomall
电商后端系统
* * *
## 给自己看
认真看一看鉴权
认真看一看分页
认真看一看模拟支付-开启事务-事务回滚
认真看一看订单超时取消（redis+goroutine）
* * *
## 商品模块
商品模块功能目标（基础版）
| 步骤  | 功能点        | 接口示例                   | 是否登录 |
| --- | ---------- | ---------------------- | ---- |
| 1️⃣ | 创建商品       | POST `/products`       | ✅ 是  |
| 2️⃣ | 获取商品列表（分页） | GET `/products`        | ❌ 否  |
| 3️⃣ | 获取某个商品详情   | GET `/products/:id`    | ❌ 否  |
| 4️⃣ | 更新商品信息     | PUT `/products/:id`    | ✅ 是  |
| 5️⃣ | 删除商品       | DELETE `/products/:id` | ✅ 是  |

Gin + GORM
图片字段、价格字段处理
分页 + 过滤查询（关键词、分类、价格区间等，后面加）
鉴权（只允许自己创建的商品被自己编辑）
数据建模（One-to-Many）
后期可扩展：库存、SKU、图片上传、秒杀、ElasticSearch搜索、MQ异步上架通知...
商品分类模块
基础功能：
创建分类
获取分类列表
更新分类
删除分类（软删除优先）
关联商品（一个商品属于一个或多个分类）
models/category.go 定义分类模型
controllers/category.go 处理分类相关接口
路由注册写 routers/router.go 或单独的 routers/category.go
### 疑问
#### .env 是做什么用的

#### routers写成这样对的嘛?
func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
}

func LoginRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
}

不对。RegisterRoutes里的Register不是“用户注册”的意思，而是“注册所有路由到GIn的Engine上”的意思。
* 就像注册事件、注册插件一样，这里的“register”指的是“注册（挂载）路由”。

#### REST client中如何获取token创建商品？
错误：
```http
 > {% 
   const res = JSON.parse(responseBody);
   await client.global.set("token", res.token);
 %}
 ```
正确：
```http
# @name loginAdmin
POST http://localhost:8080/login
Content-Type: application/json

{
  "email": "riki@example.com",
  "password": "12345678"
}

@token = {{loginAdmin.response.body.$.token}}
POST http://localhost:8080/products
Authorization: Bearer {{token}}
Content-Type: application/json

{
    "name": "iPhone 15",
    "description": "苹果新款",
    "price": 9999,
    "stock": 100,
    "image_url": "https://example.com/img.png"
}
```
* * *
## 购物车模块
购物车模块的核心需求
🎯 用户视角：
✅ 添加商品到购物车（数量可选）
✅ 查看购物车（分页、展示总价）
✅ 修改购物车中商品的数量
✅ 删除购物车中的商品
✅ 清空购物车
✅ 购物车项中不能有下架或库存为 0 的商品

CartItem
购物车的每一条记录，表示某个用户向购物车中添加了某个商品，并且指定数量。

🔗 建立了用户与商品的多对多关系（中间带属性）
📦 使用 Quantity 字段代表购买数量
🧼 购物车数据独立，不污染商品或用户表

接口设计规范（RESTful 风格）
| 接口功能     | 方法     | 路径            | 权限要求 | 说明                          |
| -------- | ------ | ------------- | ---- | --------------------------- |
| 添加商品到购物车 | POST   | `/cart`       | 登录用户 | 参数：`product_id`, `quantity` |
| 获取购物车列表  | GET    | `/cart`       | 登录用户 | 可分页、可统计总价                   |
| 修改购物车项数量 | PUT    | `/cart/:id`   | 登录用户 | 只能改自己购物车项                   |
| 删除某个购物车项 | DELETE | `/cart/:id`   | 登录用户 | 删除某一项                       |
| 清空购物车    | DELETE | `/cart/clear` | 登录用户 | 删除所有购物车项                    |

一些业务规则（防呆机制）
| 规则                 | 说明                |
| ------------------ | ----------------- |
| 商品必须存在且未下架         | 否则不能加入购物车         |
| 购买数量不能为 0 或负数      | 添加和修改时需校验         |
| 一个用户只能有一条相同商品的购物车项 | 添加时如果已存在就更新数量     |
| 用户只能操作自己的购物车项      | 修改、删除时校验 `UserID` |

可以加的一些进阶功能（等基础完成后）
 商品价格变化记录（加入购物车时记录快照）
 Redis 加速购物车读写
 限购逻辑（一个人不能买太多）
 用户登录后合并匿名购物车
 多选 + 结算 API（连接订单模块）

 ### 疑问
#### 为什么商品和购物车的更新操作一个PUT一个是PATCH？
| 方法      | 中文名  | 用途     | 语义       | 举例        |
| ------- | ---- | ------ | -------- | --------- |
| `PUT`   | 全量更新 | 更新整个资源 | 替换整个对象   | 更新商品的所有字段 |
| `PATCH` | 局部更新 | 更新部分字段 | 只改动提交的字段 | 购物车里只修改数量 |

✅ PUT /products/:id
对应：更新一个商品的信息（比如价格、名称、库存等）
产品设计角度：后台管理系统编辑商品，一般是填完整个表单再保存
语义：全量更新，旧的数据会被你新传入的数据「整体替换」
所以我们用的是 PUT
✅ PATCH /cart/:id
对应：修改购物车某个商品的数量
用户行为：只是单独改一下数量，不会动其它字段
语义：局部更新，只修改 quantity 字段
所以我们用的是 PATCH
🔥 那到底什么时候用 PUT，什么时候用 PATCH？
✅ 用 PUT 的典型情况：
更新用户资料（传入整个 profile）
更新商品（后台提交整个商品表单）
✅ 用 PATCH 的典型情况：
修改状态字段（启用/禁用）
修改部分字段（购物车数量、设置一个开关等）
修改密码（只传密码字段）

电商系统 HTTP 方法规范清单（RESTful）
🧍 用户模块
| 功能       | 路径          | 方法     | 说明            |
| -------- | ----------- | ------ | ------------- |
| 用户注册     | `/register` | `POST` | 创建用户账号        |
| 用户登录     | `/login`    | `POST` | 获取 JWT Token  |
| 获取当前用户信息 | `/me`       | `GET`  | JWT 鉴权后返回当前用户 |

🛍️ 商品模块
| 功能       | 路径                     | 方法       | 说明              |
| -------- | ---------------------- | -------- | --------------- |
| 商品列表     | `/products`            | `GET`    | 支持分页/搜索/筛选      |
| 创建商品     | `/products`            | `POST`   | 管理员使用           |
| 获取商品详情   | `/products/:id`        | `GET`    | 任何人可查看          |
| 更新商品（全量） | `/products/:id`        | `PUT`    | 替换商品信息（如后台表单提交） |
| 修改上下架状态  | `/products/:id/status` | `PATCH`  | 只修改上架状态         |
| 删除商品     | `/products/:id`        | `DELETE` | 逻辑删除或硬删除        |

🧩 分类模块
| 功能     | 路径                | 方法       | 说明           |
| ------ | ----------------- | -------- | ------------ |
| 创建分类   | `/categories`     | `POST`   | 创建新分类        |
| 获取分类列表 | `/categories`     | `GET`    | 商品分类导航用      |
| 更新分类   | `/categories/:id` | `PUT`    | 更新分类名称或父级分类  |
| 删除分类   | `/categories/:id` | `DELETE` | 删除分类（建议逻辑删除） |

🛒 购物车模块
| 功能      | 路径          | 方法       | 说明        |
| ------- | ----------- | -------- | --------- |
| 添加到购物车  | `/cart`     | `POST`   | 加入购物车     |
| 获取购物车列表 | `/cart`     | `GET`    | 查看购物车详情   |
| 更新购物项数量 | `/cart/:id` | `PATCH`  | 只更新数量     |
| 删除购物项   | `/cart/:id` | `DELETE` | 从购物车中移除商品 |

📦 订单模块（进阶）
| 功能         | 路径                   | 方法      | 说明              |
| ---------- | -------------------- | ------- | --------------- |
| 创建订单       | `/orders`            | `POST`  | 下单，生成订单         |
| 获取订单列表     | `/orders`            | `GET`   | 查看用户历史订单        |
| 获取订单详情     | `/orders/:id`        | `GET`   | 订单详情页           |
| 修改订单状态（后台） | `/orders/:id/status` | `PATCH` | 改为“已支付/已发货/已收货” |
| 取消订单       | `/orders/:id/cancel` | `PATCH` | 用户取消订单          |

💬 评论模块（可选）
| 功能       | 路径                       | 方法     | 说明       |
| -------- | ------------------------ | ------ | -------- |
| 添加评论     | `/products/:id/comments` | `POST` | 订单完成后可评论 |
| 获取商品评论列表 | `/products/:id/comments` | `GET`  | 商品详情页展示  |

📢 通知/日志系统（可选）
| 功能     | 路径        | 方法    | 说明          |
| ------ | --------- | ----- | ----------- |
| 获取日志列表 | `/logs`   | `GET` | 管理员查看操作日志   |
| 添加操作日志 | 后台逻辑中自动调用 | -     | 用户每次操作都记录日志 |

* * * 
## 订单模块
我们的目标
实现一个支持下单 → 查看订单列表 → 查看订单详情 → 支付 / 取消订单的完整流程。后续还能拓展发货、售后、评价、退款等流程。

| 接口功能     | 方法    | 路径                   | 是否登录 |
| -------- | ----- | -------------------- | ---- |
| 创建订单     | POST  | `/orders`            | ✅    |
| 获取当前用户订单 | GET   | `/orders`            | ✅    |
| 获取订单详情   | GET   | `/orders/:id`        | ✅    |
| 支付订单     | PATCH | `/orders/:id/pay`    | ✅    |
| 取消订单     | PATCH | `/orders/:id/cancel` | ✅    |

查询订单的两种模式
我们要支持两个查询接口：
1. 🔐 当前登录用户查看自己的订单列表（分页 + 详情预加载）
GET /orders
需要 JWT 鉴权
支持分页参数 ?page=1&size=10
预加载订单项（OrderItems）和商品信息
2. 🔍 查询单个订单详情
GET /orders/:id
同样需要鉴权
校验该订单是否属于当前用户

模拟支付场景设计
我们不集成真实支付网关（支付宝、Stripe），但要实现核心逻辑：
✅ 用户点击“支付订单”：
校验该订单是否属于当前用户。
校验订单状态是否是“未支付（pending）”。
校验商品库存是否足够。
扣除每个商品的库存。
修改订单状态为“已支付（paid）”。

**订单状态流转** 的核心部分：
✅ 接下来实现：
🚚 1. 商家发货（从 paid → shipped）
📦 2. 用户确认收货（从 shipped → delivered）
这两个动作其实本质上就是：
订单状态的有条件变更；
但每个变更都是一个明确动作，不是随便能改的。
状态流转规则回顾：
| 当前状态      | 操作     | 新状态         |
| --------- | ------ | ----------- |
| `paid`    | 商家发货   | `shipped`   |
| `shipped` | 用户确认收货 | `delivered` |
| 任意        | 用户取消订单 | `cancelled` |
| 任意        | 系统超时关闭 | `timeout`   |

**订单超时取消**
设计思路（Redis 延迟队列）
👇 创建订单时：
将订单信息放进一个 Redis zset（有序集合），score 是 “过期时间戳”（现在+10分钟）
后台起一个 goroutine 定时（比如每 1 分钟）轮询 Redis，取出超时未支付的订单，然后取消掉
```
expireAt := time.Now().Add(10 * time.Minute).Unix()
config.RDB.ZAdd(ctx, "order:timeout", redis.Z{
	Score:  float64(expireAt),
	Member: order.ID,
})
```
ZAdd 把 order.ID 插入到一个名为 order:timeout 的 Redis 有序集合中。
使用 expireAt 时间作为 Score，未来你可以用 ZRangeByScore 查出过期订单。
用在订单自动取消功能中（定时任务扫描 Redis）。

#### 防止超卖的核心思路（面试重点）
1. 数据库层面锁
乐观锁（Optimistic Locking）：每次更新库存时带版本号或时间戳，更新前检查版本一致性，不一致则失败重试。
悲观锁（Pessimistic Locking）：直接在数据库上对库存行加锁，串行处理。

2. 库存预扣减（Redis 锁/缓存）
先在 Redis 里扣减库存，保证高性能响应，再异步同步到数据库。
通过 Redis 原子操作，避免超卖。

3. 消息队列削峰
用户请求先进入消息队列，后台顺序消费库存，避免数据库直接高并发。

**建议实践路线**
第一步：乐观锁改造库存字段
* 给 Product 表加一个 Version 字段（int），每次更新库存时检查版本，失败则重试。简单且不影响架构。
第二步：Redis库存预扣减
* 在用户下单流程中，先调用 Redis lua 脚本原子减少库存，只有成功才写数据库订单。失败直接返回库存不足。
第三步：消息队列（RabbitMQ/Kafka/NSQ）异步处理订单
* 将下单请求写入消息队列，后台逐条处理库存和订单。

建议实践路线
第一步：乐观锁改造库存字段
给 Product 表加一个 Version 字段（int），每次更新库存时检查版本，失败则重试。简单且不影响架构。

第二步：Redis库存预扣减
在用户下单流程中，先调用 Redis lua 脚本原子减少库存，只有成功才写数据库订单。
失败直接返回库存不足。

第三步：消息队列（RabbitMQ/Kafka/NSQ）异步处理订单
将下单请求写入消息队列，后台逐条处理库存和订单。

**乐观锁的基本原理**
乐观锁（Optimistic Locking） 假设在数据更新时“不会发生冲突”，每次读取数据时带上一个版本号（或者时间戳），更新时检查这个版本有没有变化：
✨ 工作流程：
查询商品信息，同时读取 version 字段（或者 updated_at 时间戳）；
提交更新时，带上旧的 version；
在更新库存时，使用 SQL 的 WHERE id = ? AND version = ? 限定；
如果更新成功（返回行数为1），说明没有人同时修改库存；
如果失败（返回行数为0），说明发生了并发修改 → 回滚/重试。

* 还可以扩展的内容
* 支持部分商品下单
* 支持事务控制（失败就回滚库存）
* 支持并发下单控制库存（防止超卖）

### 疑问
#### 自动迁移的一些注意事项
| 注意事项                     | 说明                                                |
| ------------------------ | ------------------------------------------------- |
| AutoMigrate **不会删除字段或表** | 它只会“添加”字段，**不会移除旧字段**，所以不适合复杂的迁移场景。               |
| 生产环境建议手动迁移               | 在生产环境中，一般会使用专业的迁移工具（如 `golang-migrate`）手动控制迁移版本。  |
| 外键关联需显式声明                | GORM 只在某些数据库中自动添加外键，MySQL 有时需要你手动指定 `constraint`。 |

#### 软删除是什么？
GORM 的软删除机制其实是给表里自动加了一个 deleted_at 字段：
正常 .Delete() 并不会真删除数据，而是把 deleted_at 填上时间戳
查询时默认会加一个 WHERE deleted_at IS NULL 的过滤条件
如果你想连“已删除”的也查出来，加上 .Unscoped() 就好
```
// 查包括已软删的
db.Unscoped().Find(&orders)
// 真正删除（硬删除）
db.Unscoped().Delete(&order)
```

```
userID := c.GetUint("user_id")
```
这行代码从 Gin 的 Context 中获取当前用户的 ID。这个值通常来自于 JWT 中间件中设置的：
```
c.Set("user_id", userID) // 在登录鉴权时设置
```

#### 为什么是GetOrderStats而不是GetOrderStatus？
status 是单个状态，而 stats 是统计信息（statistics）
| 路径                | 表达意思        | 期望返回                                      |
| ----------------- | ----------- | ----------------------------------------- |
| `/orders/status`  | 状态？哪个订单的状态？ | ❓不清晰、不具体                                  |
| `/orders/stats` ✅ | 全部订单的统计信息   | {"paid": 3, "pending": 2, "cancelled": 1} |
| 路径                     | 用法              | 说明          |
| ---------------------- | --------------- | ----------- |
| `/orders/stats` ✅      | 获取当前用户所有订单的状态统计 | 是一个聚合接口     |
| `/orders/:id/status` ✅ | 获取某个订单的状态       | 是一个精确查询     |
| `/orders/status` ❌     | 意义模糊            | 不符合 REST 规范 |

#### 关于:id贪婪匹配的一些注意事项
为什么不能
```
order := auth.Group("/orders")
{
	order.POST("", controllers.CreateOrder)
	order.GET("", controllers.GetOrders)
	order.GET("/:id", controllers.GetOrderDetail)
	order.POST("/:id/pay", controllers.PayOrder)
	order.GET("/stats", controllers.GetOrderStats) // 👈 这个会被误解读
}
```
而必须是
```
order := auth.Group("/orders")
{
	order.POST("", controllers.CreateOrder)
	order.GET("/stats", controllers.GetOrderStats)      // 👈 先写这个
	order.GET("", controllers.GetOrders)
	order.GET("/:id", controllers.GetOrderDetail)
	order.POST("/:id/pay", controllers.PayOrder)
}
```
这和 Gin 的路由匹配机制有关，尤其是 :id 这种路径参数的贪婪匹配（greedy matching）行为。
原因：:id 是贪婪匹配，你期望执行的是 GetOrderStats，但实际执行的却是 GetOrderDetail。
Gin 会从上往下查找哪个路由“能匹配得上”，它看到：/orders/stats，然后看看 /orders/:id 能不能匹配？ Gin 判断：
:id 是一个变量，占位符。
/orders/stats ➜ /orders/:id ➜ :id = "stats"
于是 Gin 就把这个请求交给了 GetOrderDetail 处理。
但 "stats" 显然不是一个合法订单 ID，查询数据库失败，自然返回“订单不存在”。

**本质问题**：动态参数 :id 把静态路径 /stats 吃掉了
Gin 的路径匹配优先级：
静态路径（如 /orders/stats）
带参数路径（如 /orders/:id）
通配符路径（如 /orders/*path）
但是如果 /orders/:id 写在 /orders/stats 上面，Gin 会优先匹配成功第一个符合的路径，而不会继续尝试后面的。

| 匹配类型  | 示例路径            | 匹配优先级 | 说明     |
| ----- | --------------- | ----- | ------ |
| 静态路径  | `/orders/stats` | ✅ 最高  | 必须完全一致 |
| 参数路径  | `/orders/:id`   | 中     | 匹配任意单段 |
| 通配符路径 | `/orders/*path` | 最低    | 匹配多段路径 |

#### 为什么订单取消是post？
这里的「取消订单」是对订单状态的修改操作，严格来说它是一个 **“写”动作，属于改变服务器资源状态** 的请求。HTTP 语义中：
GET：拿数据，不能改东西。想象成“看”东西。
POST：改东西或者做动作，比如“取消订单”、“支付订单”。是“操作”。
PUT/PATCH：直接改资源的具体内容，比如“改订单里的地址”。
DELETE：删东西。

取消订单是“执行一个操作”，用POST表达“我要执行取消”这个动作最合适。
取消订单虽然是修改状态，但并不是用 PUT/PATCH 的典型场景（比如直接替换订单资源），而是触发一个“操作”（action），让订单状态从“待支付”变为“已取消”，而且这种操作一般没有幂等性（连续调用可能会报错），POST 更符合语义。
（其实还不太明白）

#### 库存的状态是怎么改变的？
订单取消释放库存
怎么占用库存
订单支付之后的库存又怎么变

#### 什么是系统高并发、解耦、异步处理、削峰填谷？

#### 订单超时取消分别可以使用 cron 定时器（推荐）或者使用 MQ 结合延迟队列（阿里云/RabbitMQ/Redis Stream），两者分别具体是什么？有什么不同？
1. MQ 是什么？解决什么问题？
简单说：
异步处理：下单后马上响应用户，后续发短信/发邮件放入队列慢慢处理。
解耦系统：订单系统不需要关心短信系统是否挂了。
削峰限流：高峰期先把任务写入队列，慢慢消费。
2.  会用至少一种消息队列
推荐你学 Redis 或 RabbitMQ，因为文档多、好部署、入门快。
🌟 建议从 Redis 的延迟队列 开始学，因为你已经在用 Redis！
3.  理解消息的三要素（核心）
消息生产者：谁发消息（订单系统）
消息队列：暂时存储消息（Redis/RabbitMQ）
消息消费者：谁处理消息（发送邮件、更新状态等）
4.  能写一个实际的小项目用上 MQ
用户下单后，发一条消息进入队列
消费者从队列中取出消息 → 执行“30分钟后取消订单”
5.  知道基础概念即可
| 名词                | 你可以现在先跳过 |
| ----------------- | -------- |
| 消息确认（ACK）         |          |
| 死信队列（DLX）         |          |
| 分区、Broker、高可用副本   |          |
| Kafka Offset 提交机制 |          |
