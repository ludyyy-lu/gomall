# 🛒 GoMall - 基于 Go 的电商后端系统

GoMall 是一个使用 Go 开发的电商后端项目，涵盖了用户系统、商品管理、购物车、订单处理、库存控制、消息队列、Redis 缓存等完整模块，适合作为中大型后端项目的实战练习或简历项目。

---

## 🚀 项目亮点

* ✅ **用户认证与权限控制**：使用 JWT 实现用户登录鉴权，封装中间件自动提取用户信息。
* ✅ **商品管理系统**：支持上架/下架、库存字段、秒杀标记、多分类（多对多表）。
* ✅ **购物车功能**：增删改查接口齐全，支持多商品购物车操作。
* ✅ **订单系统**：支持下单、支付、发货、订单状态变更等完整流程。
* ✅ **库存控制机制**：

  * 乐观锁控制商品库存，防止并发超卖
  * Redis + Lua 脚本保证高并发下的原子性（适用于秒杀场景）
* ✅ **消息队列集成**：

  * RabbitMQ 异步发送订单创建消息
  * 消费者服务独立监听处理消息（支持拓展发邮件/通知）
* ✅ **Redis 延迟任务**：

  * 使用 Redis ZSet 管理订单超时关闭
  * 后台协程定时扫描并处理超时订单
* ✅ **Docker 一键部署**：通过 Docker Compose 编排 MySQL、Redis、RabbitMQ、应用服务，实现本地一键启动

---

## 🧱 技术栈

| 分类    | 技术                     |
| ----- | ---------------------- |
| 开发语言  | Go 1.20+               |
| Web框架 | Gin                    |
| ORM   | GORM                   |
| 数据库   | MySQL                  |
| 缓存    | Redis                  |
| 消息队列  | RabbitMQ               |
| 配置管理  | godotenv               |
| 容器化   | Docker, Docker Compose |

---

## 📦 模块设计

* `/models`：数据库模型定义
* `/controllers`：路由控制器（商品、购物车、订单等）
* `/routers`：接口路由注册
* `/config`：数据库、Redis、RabbitMQ 初始化
* `/utils`：通用工具函数（JWT、订单消费者、库存脚本等）

---

## 📜 简历项目描述模板（中文）

> **GoMall 商城系统 | 后端开发（个人项目）**
> 使用 Go + Gin + GORM 构建电商后端系统，支持用户认证、商品管理、购物车、订单处理等完整业务流程。采用 Redis + Lua 保证秒杀库存原子扣减，使用 RabbitMQ 解耦异步通知，结合 Docker Compose 实现一键部署。具备完整事务控制、库存并发安全与任务超时处理能力。

---

## 📄 简历项目描述模板（英文）

> **GoMall E-commerce Backend | Backend Developer (Solo Project)**
> Developed a full-featured e-commerce backend using Go, with user authentication, product management, cart and order processing. Implemented concurrency-safe stock deduction with optimistic locking and Redis Lua scripting. Integrated RabbitMQ for asynchronous order notifications. Used Docker Compose for local multi-service orchestration and deployment.

---

## 🛠 如何运行

```bash
# 克隆项目
$ git clone https://github.com/yourname/gomall.git
$ cd gomall

# 创建 .env 文件，配置数据库/Redis/RabbitMQ 连接信息
$ cp .env.example .env

# 启动服务（需安装 Docker & docker-compose）
$ docker-compose up -d

# 运行主服务（API 网关）
$ go run main.go

# 可选：单独运行订单消费者
$ go run main/consumer.go
```

---

## ✅ TODO（可拓展方向）

* [ ] 接入支付接口（支付宝、微信）
* [ ] 商品全文搜索（如 Elasticsearch）
* [ ] 后台管理面板（支持权限控制）
* [ ] 日志与监控（Prometheus + Grafana）
* [ ] 前后端分离部署 + CDN 加速

---


本项目基本完结，感谢GPT感谢我自己。
readme中记录我在编写项目时遇到的问题以及认为值得留意的部分。
适合小白阅读学习。

# gomall
电商后端系统

## 总结复盘

✅ 项目完成进度（截至目前）
| 功能模块          | 描述                                    |
| ------------- | ------------------------------------- |
| **用户系统**      | 注册 / 登录 / JWT 鉴权                      |
| **商品管理**      | 支持上架/下架/分类绑定                          |
| **购物车系统**     | 增删改查购物车商品                             |
| **订单系统**      | 创建订单、发货、超时取消、历史记录                     |
| **库存控制**      | ✅ 乐观锁实现 + ✅ Redis + Lua 双重方案（兼容普通和秒杀） |
| **Redis 应用**  | Session缓存、订单超时、库存缓存                   |
| **消息队列**      | RabbitMQ 实现订单异步通知（消费者独立部署）            |
| **Docker 集成** | MySQL、Redis、RabbitMQ 全部容器化部署          |
| **代码结构清晰**    | Controller-Service-Model 分层，支持依赖注入    |
| **可测试性强**     | 可单独运行消费者、支持接口 Postman 调试              |
| **预留扩展**      | ✅ 支付功能、✅ 秒杀专用下单接口、✅ Kafka 替换 MQ       |

✅ 项目亮点（面试能讲）
1. 并发安全库存处理（Redis + Lua + 乐观锁）；

2. 消息队列异步解耦（RabbitMQ消费者独立部署）；

3. 订单自动取消机制（Redis SortedSet 模拟延迟队列）；

4. 模块化清晰依赖注入（Go项目结构优雅、可维护）；

5. 一键部署（docker-compose）；

6. 支持横向扩展的设计（你甚至可以扩展为多服务、微服务）；

7. 具备秒杀系统的雏形能力！（Redis预热 + 高并发接口）

✅ 下一步建议（可选）
| 方向       | 可做的事情                             |
| -------- | --------------------------------- |
| 💳 支付    | 模拟支付接口，或集成支付宝/微信的沙箱环境             |
| 📊 监控    | 接入 Prometheus + Grafana，监控订单、库存变化 |
| 🔄 Kafka | 把 RabbitMQ 替换成 Kafka，进一步提升吞吐      |
| 🌏 部署    | 放到 Linux 云服务器上，完整上线一次系统           |
| 🧪 测试    | 用 Postman + Swagger 文档写 API 测试计划  |

* * *
## 面试相关
### 项目讲解脚本
✅ 一、项目讲解脚本（用于面试时讲项目）
适合口头表达，我会帮你写成适合你背诵/自由发挥的风格：
结构建议：
1. 一句话概括项目
2. 项目背景与目标
3. 整体架构设计
4. 关键技术难点 & 你的解决方案
5. 你具体负责的模块
6. 当前项目亮点
7. 未来可拓展的方向

🗂 项目讲解脚本结构
项目背景（为什么做）
核心功能（做了什么）
技术选型（怎么做）
重点难点（解决了什么问题）
架构设计（模块划分 & 技术栈）
亮点细节（可展示代码或场景）
项目收获（你成长了什么）
可拓展点（下一步会怎么做）


1️⃣ 项目背景（Why）
“这个项目是我在学习 Go 后端开发时主导完成的一个完整实践项目，模拟的是一个线上商城系统，涵盖用户登录、商品管理、购物车、订单处理、库存控制等完整电商核心流程。我希望通过这个项目了解真实业务的开发流程，掌握主流后端技术栈，提升自己在高并发场景下的设计与优化能力。”

2️⃣ 核心功能（What）
“整个商城系统包含了从用户注册登录、商品浏览、加入购物车、创建订单、支付发货、订单超时处理等完整电商流程。
我还引入了 Redis 预加载库存、使用 Lua 脚本实现原子扣库存机制，以及 RabbitMQ 作为消息队列进行异步处理订单通知等功能。
系统支持秒杀类高并发场景，并通过乐观锁机制防止库存超卖问题。”
🧠（补充说明）你也可以在这一步用一句话强调：“这个项目的目标，是从 0 到 1 练习电商后端核心逻辑，包括高并发控制、消息异步、缓存设计等。”

3️⃣ 技术选型（How）
“我使用 Go 语言作为主要开发语言，Gin 框架构建 HTTP 接口，Gorm 作为 ORM 工具操作 MySQL。
Redis 负责缓存与库存控制，RabbitMQ 用于订单相关的异步消息通知。系统配置采用 .env 管理，服务容器化部署采用 Docker Compose。”
你也可以补一句：
“所有依赖服务如 MySQL、Redis、RabbitMQ 都通过 Docker Compose 一键启动，方便本地开发与部署。”

4️⃣ 重点难点（Pain Point）
这一部分是加分项，非常重要，尤其适合回答面试官问的“你在这个项目里最大的挑战是什么？”
“我在项目中遇到的最大挑战，是如何在高并发场景下防止库存超卖。为了解决这个问题，我使用了两种方式：
一种是数据库层的乐观锁（通过版本号控制）；
另一种是在 Redis 中提前预加载库存，并使用 Lua 脚本实现原子扣减，确保在并发秒杀时库存扣减不出错。
另一个难点是消息异步处理。为了解耦订单逻辑，我使用 RabbitMQ 将订单创建事件发送到独立的消费者服务中，在消费者中进行通知处理、日志记录等。这也让我更深入理解了消息队列在真实项目中的使用价值。”

5️⃣ 架构设计（Architecture）
“这个项目采用了典型的 MVC 架构，逻辑清晰、层次分明：
Model 层使用 GORM 定义数据模型，并负责数据库操作；
Controller 层处理请求参数和业务逻辑；
Router 层进行路由注册和中间件挂载。
另外，为了提升系统的扩展性，我将 Redis、RabbitMQ 等依赖组件封装在 config 包中集中初始化；并为 Redis、MQ 单独开了 utils 工具包，用于封装通用功能，比如：库存预加载、订单消费者、订单超时处理。”
你也可以补充一句技术图景：
“整体架构通过 Docker Compose 编排多个服务，方便在本地进行端到端测试与部署。”

6️⃣ 项目亮点（Highlights）
可以重点讲你做得“超出预期”的部分，这个真的会加分：
“这个项目中，我不仅完成了常规的电商流程，还实现了一些更偏真实场景的设计，比如：
* 使用乐观锁与 Redis + Lua 双重手段防止超卖，支持秒杀场景
* 使用 RabbitMQ 解耦核心业务，异步处理订单通知等操作
* 独立部署订单消费者服务，增强系统可扩展性与稳定性
* 使用 Redis 有序集合实现订单延迟关闭机制，避免僵尸订单
* 项目整体支持容器化部署，快速在本地还原生产环境”

7️⃣ 项目收获（Growth）
这部分是你的自我提升，也可以稍微带点 Z 世代的语气（真诚但不自吹）：
“通过这个项目，我系统掌握了 Web 后端项目从零开发的完整流程，也真正体会到 Redis 和消息队列在真实业务中的使用场景。
在踩过坑、调过 bug 的过程中，我的代码规范、调试能力和系统架构理解都得到了很大提升。”
“我也更加意识到，项目不仅要能跑，还要具备可维护性、可扩展性。这也是我尝试拆分配置、抽象组件、容器化部署的原因。”

8️⃣ 可拓展性（Scalability）
“如果有更多时间，我希望将这个项目进一步完善，比如：
加入用户权限系统（后台管理 / 普通用户）
接入支付接口模拟（如支付宝沙箱）
日志与监控（接入 ELK 或 Prometheus）
商品搜索功能（结合全文检索）
多语言与多币种支持（针对国际用户）”

🎯 项目讲解提纲卡片《GoMall 商城系统》
🪪 一句话介绍（Who am I）
    “我做了一个基于 Go 的电商后端系统，涵盖了用户登录、商品管理、购物车、下单、库存、消息队列等完整电商业务流程。”

🧱 项目核心模块（What did I build）
用户系统：JWT 认证登录，用户鉴权中间件
商品系统：上架、下架、库存字段，支持多分类
购物车系统：商品加入、修改、删除、查看
订单系统：下单、支付、发货、取消、超时关闭
库存控制：乐观锁 + Redis + Lua 双保险防超卖
订单通知：RabbitMQ 实现消息异步发送
延迟取消：Redis ZSet 实现订单自动过期关单
容器部署：Docker Compose 启动 MySQL、Redis、RabbitMQ

🧠 技术细节亮点（How I solved the problem）
✅ 并发库存问题：
乐观锁方式更新 stock 和 version 字段
Redis + Lua 原子操作用于秒杀商品限购
✅ 消息解耦：
RabbitMQ 生产者发送订单创建消息
消费者服务独立监听并处理，如发送通知、日志等
✅ 自动关单：
Redis 有序集合 ZSet 存储超时订单 ID
Goroutine 定时扫描，关闭未支付订单
✅ 中间件与封装：
环境变量配置统一加载
Redis/MQ 初始化模块化
控制器依赖注入（注入 DB、Redis、MQ）

🔭 可扩展方向（What's next）
接入支付接口（支付宝、微信）
后台管理系统（商品审核、用户管理）
分布式架构演进（拆分服务、Kafka 替代 MQ）
加入 Prometheus + Grafana 做监控告警
前后端分离部署 + CDN 加速

🧑‍💻 一句话收尾
“这个项目虽然是练手项目，但我尽可能贴近真实业务来设计，特别注重可扩展性和真实可运行。很多方案也都可以在实际工作中复用。”

### 项目总结文档
✅ 二、项目总结文档（README 或简历描述）
适合写在 GitHub / 简历 / 项目文档中，写作风格偏向正式专业，突出技术点和个人贡献。
内容结构：
1. 项目名称 & 简介
2. 技术栈
3. 功能模块介绍
4. 架构图（如果有）
5. 项目亮点（比如：Redis 秒杀、RabbitMQ 异步处理、乐观锁防超卖、Docker Compose）
6. 启动方式（包含 .env、docker-compose up 等）
7. 如果有：测试截图 / API 示例
8. 你在项目中扮演的角色 & 技术贡献点


#### 📜 简历项目描述模板（中文）

> **GoMall 商城系统 | 后端开发（个人项目）**
> 使用 Go + Gin + GORM 构建电商后端系统，支持用户认证、商品管理、购物车、订单处理等完整业务流程。采用 Redis + Lua 保证秒杀库存原子扣减，使用 RabbitMQ 解耦异步通知，结合 Docker Compose 实现一键部署。具备完整事务控制、库存并发安全与任务超时处理能力。



#### 📄 简历项目描述模板（英文）

> **GoMall E-commerce Backend | Backend Developer (Solo Project)**
> Developed a full-featured e-commerce backend using Go, with user authentication, product management, cart and order processing. Implemented concurrency-safe stock deduction with optimistic locking and Redis Lua scripting. Integrated RabbitMQ for asynchronous order notifications. Used Docker Compose for local multi-service orchestration and deployment.



* * *
## 给自己看
认真看一看鉴权
认真看一看分页
认真看一看模拟支付-开启事务-事务回滚
认真看一看订单超时取消（redis+goroutine）
redis和rabbitMQ的相关部分多看一下
lua+redis秒杀

* * *
## 用户模块
注册、登陆（jwt鉴权）
Me接口：获取当前登录用户信息
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
**后期可扩展：库存、SKU、图片上传、秒杀、ElasticSearch搜索、MQ异步上架通知...**

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

**可以加的一些进阶功能（等基础完成后）**
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

## docker 部署 这一部分写乱了，docker和rabbitMQ的内容混在一起了
#### 一些docker配置的注意事项
.env文件
```
# mysql
DB_USER=gomalluser
DB_PASS=gomallpass
# 用于宿主机访问
DB_HOST_EXTERNAL=localhost
# 用于容器内部通信
DB_HOST=mysql
DB_PORT=3307
DB_NAME=gomall
```
docker-compose.yml文件
```
  mysql:
    image: mysql:8.0
    container_name: gomall-mysql
    environment:
      MYSQL_ROOT_PASSWORD: ${DB_PASS} # root用户的密码
      MYSQL_DATABASE: ${DB_NAME}
      MYSQL_USER: ${DB_USER}  # 新建一个普通用户
      MYSQL_PASSWORD: ${DB_PASS} # 普通用户的密码
    ports:
      - "${DB_PORT}:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    restart: always  # 自动重启
```
ports:
     - "${DB_PORT}:3306"
类似于这种部分，是3307映射到3306，其中3307是宿主机端口号，而3306是docker内部的端口号。

#### 外部访问 Docker 容器中的服务时，为什么用 localhost 而不是容器内的配置
这个问题源于
```
如何验证代码运行正常？
✅ 验证主服务 gomall-app
访问 http://localhost:8080：
查看商品接口是否能访问。
注册登录是否成功。
下单是否能收到响应。
```
为什么是localhost而不是我们docker里面设置的其他的host？

gpt回答：
为什么你访问的是 http://localhost:8080/orders
因为你用的是 Postman 或浏览器在 主机（Windows）上访问接口，而你的商城后端（Gin服务）绑定在：
```
r.Run(":8080")
```
这表示它监听的是容器的 8080 端口，并通过 Docker 映射到了你本地的 localhost:8080。

那 RabbitMQ 怎么是用 rabbitmq:5672 呢？
这是因为你的 consumer.go 是运行在 Docker 容器内部的，在 Docker Compose 的同一个网络中：
容器之间可以通过服务名互相通信。
你的 RabbitMQ 服务的名称是 rabbitmq（看 docker-compose.yml 中的 key）。
所以 consumer.go 中连接 RabbitMQ 要用：
```
RABBITMQ_URL=amqp://gomall:gomall123@rabbitmq:5672/
```
而不是 localhost:5672，因为 **localhost 代表的是容器自己**，而不是宿主机。

| 场景                   | 正确地址                                     | 错误写法                       |
| -------------------- | ---------------------------------------- | -------------------------- |
| 浏览器访问商城接口            | `http://localhost:8080`                  | `http://rabbitmq:8080` ❌   |
| consumer 访问 RabbitMQ | `amqp://gomall:gomall123@rabbitmq:5672/` | `amqp://localhost:5672/` ❌ |

豆包回答：
这里的核心问题是 “外部访问 Docker 容器中的服务时，为什么用 localhost 而不是容器内的配置”，这涉及到 Docker 的网络映射机制，我们一步步拆解：
**一、先明确几个关键概念**
宿主机（Host）：你正在使用的电脑（比如你的 Windows 或 Mac），是 Docker 容器的 “宿主”。
Docker 容器（Container）：一个独立的 “小虚拟环境”，里面运行着你的应用（比如这里的 gomall-consumer、RabbitMQ 等）。
端口映射（Port Mapping）：Docker 允许将容器内的端口 “暴露” 给宿主机，这样宿主机或外部设备才能访问容器内的服务。
**二、为什么用 localhost:8080 访问？**
假设你的 订单服务（处理 /orders 请求的服务） 是运行在 Docker 容器中的，且启动时配置了端口映射，比如：
```
# docker-compose.yml 中的订单服务配置
order-service:
  image: your-order-image
  ports:
    - "8080:8080"  # 关键：将容器内的 8080 端口映射到宿主机的 8080 端口
```
这时：
容器内部的服务监听的是容器自己的 8080 端口（容器内的网络是独立的，和宿主机不直接互通）。
但通过 ports: "8080:8080" 配置，Docker 会在宿主机上 “开放” 8080 端口，并将所有发送到宿主机 8080 端口的请求转发到容器内的 8080 端口。
因此，当你在宿主机上用工具（如 Postman）发送请求时：
你实际操作的是宿主机的网络，所以必须用宿主机的地址（localhost 或宿主机的 IP，比如 192.168.1.100）。
端口用宿主机上映射的端口（这里是 8080），而不是容器内的端口（虽然这里两者相同，但逻辑上必须用宿主机的端口）。
**三、为什么不用 Docker 内部的配置（比如容器名）？**
Docker 容器之间可以通过 容器名 或 服务名 互相访问（比如你的 gomall-consumer 连接 RabbitMQ 时用 rabbitmq 作为地址），但这仅限于 容器之间的内部通信。
例如：
你的 gomall-consumer 在容器内连接 RabbitMQ 时，用的是 rabbitmq:5672（rabbitmq 是容器名，容器内部可以直接通过这个名字找到对应的容器）。
但如果你在 宿主机 上（比如用 Postman 或浏览器）访问，宿主机不认识 rabbitmq 这个名字（这是 Docker 内部的 “别名”），必须用宿主机的地址（localhost）+ 映射的端口。
**四、总结：localhost vs 容器名的区别**
|场景	| 访问方式	| 原因 |
| -------------------- | ---------------------------------------- | -------------------------- |
|宿主机访问容器内服务 |	http://localhost:映射的端口	| 宿主机只能通过自己的地址 + 映射的端口访问容器 |
|容器内访问其他容器服务	| http://容器名:容器内端口	|Docker 内部有 “DNS 解析”，能通过容器名找到对应容器 |
|容器内访问宿主机服务	| http://host.docker.internal:端口	| 这是 Docker 提供的特殊地址，用于容器访问宿主机 |
#### RabbitMQ 功能测试
启动 RabbitMQ 容器（gomall-rabbitmq）。
启动消费者容器（gomall-consumer）。
```
docker-compose up --build consumer
```
启动成功之后，
登录 RabbitMQ 管理界面（默认 http://localhost:15672，账号密码 从.env中得到），找到 order.created 队列，手动发送一条 JSON 格式的测试消息。
在 RabbitMQ 管理界面中，向order.created队列手动发送 JSON 格式测试消息，可按以下步骤操作：
登录 RabbitMQ 管理界面后，在左侧菜单中点击 “Queues”，找到 “order.created” 队列并点击进入该队列的详情页面。
在队列详情页面，向下滚动页面，找到 “Publish message” 部分。
在 “Headers” 区域，添加一个键值对，键为 “contentType”，值为 “application/json”，以此指定消息内容为 JSON 格式。
在 “Payload” 输入框中，填写你要发送的 JSON 格式消息内容。注意键值需要和建表时表示一致。例如：
```
{
  "id": 1001,
  "user_id": 123,
  "total_price": 99.99
}
```
确认填写无误后，点击 “Publish message” 按钮，即可将 JSON 格式的测试消息发送到order.created队列中。

发送成功之后，终端显示如下：
```
gomall-consumer  | 2025/07/14 06:12:57 📥 收到新订单：ID=1001, 用户ID=123, 总价=99.99
```
说明消息队列功能正常。

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
1. 系统高并发
什么是高并发？
系统同时处理大量请求的能力。比如某电商大促时，成千上万用户同时下单，系统需要同时响应这些请求。
为什么难？
资源有限（CPU、内存、数据库连接等），请求过多可能导致响应慢、系统崩溃。
举例
你打开一个网站，看到页面秒开，不管多少人同时访问，这就是高并发处理得好。
2. 解耦
什么是解耦？
将系统中各个模块/组件之间的依赖降低，使它们独立工作、相互影响小。
为什么要解耦？
方便维护、扩展，某个模块改动不会影响其他模块，提高系统稳定性。
举例
订单系统和库存系统本来直接连着，如果改为订单系统发消息通知库存系统扣库存，就是解耦了。两个系统之间不直接调用，彼此独立。
3. 异步处理
什么是异步？
任务发出后不等它完成，立即返回，任务在后台慢慢执行。
为什么用异步？
避免用户等待，提高系统响应速度。
举例
用户下单后，立即返回“订单已提交”，实际扣库存、发通知这些操作异步完成。
4. 削峰填谷
什么意思？
把瞬间大量涌入的请求（峰值）变平缓，分散到一段时间内处理，避免系统瞬间超载。
怎么实现？
用队列、消息中间件缓冲请求，慢慢消费。
举例
双十一秒杀时，用户抢购请求瞬间爆发，系统用消息队列缓冲排队，逐个处理，系统就不会崩。

| 概念   | 意义/作用             |
| ---- | ----------------- |
| 高并发  | 同时处理大量请求，保持系统稳定响应 |
| 解耦   | 各部分独立，方便维护和扩展     |
| 异步处理 | 不阻塞用户请求，后台处理耗时操作  |
| 削峰填谷 | 平滑流量高峰，防止系统超载     |

#### 订单超时取消分别可以使用 cron 定时器（推荐）或者使用 MQ 结合延迟队列（阿里云/RabbitMQ/Redis Stream），两者分别具体是什么？有什么不同？
1. MQ 是什么？解决什么问题？
简单说：
异步处理：下单后马上响应用户，后续发短信/发邮件放入队列慢慢处理。
解耦系统：订单系统不需要关心短信系统是否挂了。
削峰限流：高峰期先把任务写入队列，慢慢消费。
1.  会用至少一种消息队列
推荐你学 Redis 或 RabbitMQ，因为文档多、好部署、入门快。
🌟 建议从 Redis 的延迟队列 开始学，因为你已经在用 Redis！
1.  理解消息的三要素（核心）
消息生产者：谁发消息（订单系统）
消息队列：暂时存储消息（Redis/RabbitMQ）
消息消费者：谁处理消息（发送邮件、更新状态等）
1.  能写一个实际的小项目用上 MQ
用户下单后，发一条消息进入队列
消费者从队列中取出消息 → 执行“30分钟后取消订单”
1.  知道基础概念即可
| 名词                | 你可以现在先跳过 |
| ----------------- | -------- |
| 消息确认（ACK）         |          |
| 死信队列（DLX）         |          |
| 分区、Broker、高可用副本   |          |
| Kafka Offset 提交机制 |          |

#### 依赖注入是什么？
1. 背景：什么是“依赖”？
在后端项目中，一个组件通常依赖其他组件才能运行：
比如你写一个服务 OrderService，它依赖于：
DB（数据库）、Redis（缓存）、Logger（日志）
这些就是它的“依赖”对象。
2. 传统写法（不使用依赖注入）
你可能会这样写：
```
var db = config.DB
var rdb = config.RDB

func CreateOrder() {
    db.Create(...)
    rdb.Set(...)
}
```
上面用了全局变量，虽然简单，但**耦合**严重：
不好测试（无法注入 mock 数据库）
组件不灵活（只能用全局的 Redis 和 DB）

*注：耦合指的是一个模块/函数/组件 对另一个模块依赖的紧密程度。通俗点说：高耦合：两个模块绑死了，一个改动，另一个就出问题，像“连体婴”。低耦合：模块之间松耦合，各自独立，就像插头插座，想换就换。*

 *情况1：高耦合*
```
func Register() {
	db := config.DB // 全局变量
	db.Create(&User{...})
}
```
*问题：函数里面写死了 config.DB，你无法给它换数据库、做测试模拟。*
*结果：要改数据库，就得改这个函数的内部逻辑。*

*情况2：低耦合（使用依赖注入）*
```
func Register(db *gorm.DB) {
	db.Create(&User{...})
}
```
*Register 不关心你给什么 DB，只管用。你可以轻松传入 mock、测试库、MySQL、SQLite...*

1. 依赖注入：用参数或结构体注入依赖
```
type OrderService struct {
    DB  *gorm.DB
    RDB *redis.Client
}

func (s *OrderService) CreateOrder() {
    s.DB.Create(...)
    s.RDB.Set(...)
}
```
然后在 main.go 或路由注册时注入：
```
service := OrderService{
    DB:  db,
    RDB: rdb,
}
```
好处：
更容易测试
更灵活（每个模块可以配置自己的依赖）
解耦合（模块间更清晰）

总结：依赖注入的核心就是不要直接在函数里“拿”依赖，而是通过参数传进来，保持灵活性和可控性。
#### rabbitMQ是什么？
RabbitMQ 是一个消息队列中间件，用于在系统之间传递消息。
场景：假如你有这样的系统流程
用户下单
系统扣库存
系统发邮件
系统推送消息通知
如果这些操作都写在一个函数里，万一邮件接口超时、库存不够，会影响整体流程。

引入 RabbitMQ → 解耦 + 异步处理
你可以：
用户下单成功
往 RabbitMQ 发送一个消息（比如：OrderCreated）
其他服务（库存服务 / 邮件服务）去订阅这个消息
各自独立处理各自的逻辑
这样做的好处：
解耦：订单服务不管库存服务有没有成功
异步：发送消息是立即完成的，不会阻塞主流程
可靠：RabbitMQ 保证消息不会丢

RabbitMQ 的核心概念
| 概念          | 作用说明                |
| ----------- | ------------------- |
| Producer    | 生产者，发送消息            |
| Consumer    | 消费者，接收消息            |
| Queue       | 消息队列，临时存放消息         |
| Exchange    | 路由器，决定消息进哪个队列       |
| Routing Key | 路由键，和 Exchange 配合使用 |


#### RabbitMQ 消息传递流程图
```
[你的代码：PublishOrderCreated]
        │
        │  json.Marshal(order)
        ▼
 ┌────────────────────────────┐
 │     Exchange（交换机）      │      <—— 通常是 ""（默认交换机）
 └──────────┬─────────────────┘
            │
    Routing Key: "order.created"
            │
            ▼
     ┌──────────────────┐
     │   Queue（队列）    │
     │  "order.created"  │ <—— 你 QueueDeclare 的队列
     └──────────────────┘
            │
         消费者监听
            ▼
 [处理逻辑：发邮件 / 发短信 / 通知物流系统等]
 ```
1. Exchange（交换机）
RabbitMQ 中的消息都 先送到交换机。
再由交换机根据规则把消息投递到对应的队列。
我们上面用的 "" 是 RabbitMQ 的默认交换机（直接根据 routing key 找到对应队列）。
1. Routing Key（路由键）
相当于“快递地址”
默认交换机会把 routing key == 队列名 的消息直接送到那个队列。
1. Queue（队列）
存放消息的容器，消息在这里等待消费者来取。
一个队列可以有多个消费者来消费。
1. 消费者
用来监听某个队列，一旦有消息就触发回调函数执行你的业务逻辑。

现在这个 PublishOrderCreated 函数的意思是：
订单创建成功后 ➜ 把订单消息发给 RabbitMQ ➜ 进入 order.created 队列 ➜ 等着其他消费者去接收这个订单数据。

🧠 面试准备建议
🎯 项目讲解模板（你要背下）
“我做了一个基于 Go 的电商后端系统，用了 Gin + GORM + Redis + RabbitMQ，包含用户、商品、购物车、订单四大模块。订单部分做了防止超卖（乐观锁）和超时取消机制（Redis + 定时任务）。消息队列部分通过 RabbitMQ 发送订单创建通知，已完成消费者模块的初步逻辑，未来可拓展发短信/邮件提醒。系统采用依赖注入方式管理 Redis、数据库和 MQ，提高了可维护性。”
💡 面试常问：
| 问题              | 你可以怎么答                                                |
| --------------- | ----------------------------------------------------- |
| 怎么防止超卖？         | 使用乐观锁，给商品加版本号 version，通过 WHERE 条件判断库存 + version 是否一致。 |
| 怎么取消超时订单？       | Redis ZSet 保存过期时间，定时任务扫描并取消订单，回滚库存。                   |
| 你为什么选 RabbitMQ？ | 学习成本低、适合业务解耦、已有管理 UI。Kafka 可选，未来扩展可切换。                |


消费者函数StartOrderConsumer如果要使用goroutine的注意事项
✅ 优点：
✅ 异步监听，不会阻塞主线程
✅ 适合你想继续做其他事（比如继续初始化 HTTP 服务）
✅ 程序结构更清晰，main 不会被卡住
⚠️ 注意：
你要保证 main 不会马上退出，否则 goroutine 没机会执行（这就是为啥你在 main 里面要留住程序运行，比如用 gin.Run()）

## Redis库存+秒杀限购（用 Lua 脚本 原子操作）
我们已经完成的部分
| 功能项             | 状态 | 说明                            |
| --------------- | -- | ----------------------------- |
| Redis 已接入       | ✅  | 已经封装成 `InitRedis()` 并注入控制器    |
| RabbitMQ 异步消息队列 | ✅  | 在订单创建时发送消息，在消费者中异步处理          |
| Redis 超时订单处理    | ✅  | 使用 `ZSet + 轮询 goroutine` 定期处理 |

我们需要在哪部分引入 Redis+Lua 秒杀控制？
你目前的订单创建逻辑里是在这里做库存检查和扣减的：
```
// controller/order.go 的 CreateOrder 方法中
result := tx.Model(&models.Product{}).
	Where("id = ? AND version = ? AND stock >= ?", product.ID, product.Version, item.Quantity).
	Updates(map[string]interface{}{
		"stock":   gorm.Expr("stock - ?", item.Quantity),
		"version": gorm.Expr("version + 1"),
	})
```
这里用的是 ✅ 乐观锁 —— 适合普通购物流程，但不适合秒杀。

✅ 秒杀版本应该如何改？
1️⃣ 预加载库存到 Redis（初始化的时候做一次）

    目的
    在系统启动或秒杀开始前，把商品库存预先加载进 Redis，后续所有库存判断、扣减都只走 Redis，避开慢吞吞的数据库。
    我们设计一个 Redis key 规则：seckill:stock:<product_id>
    比如商品 ID 是 101，那 Redis 里就存：seckill:stock:101 = 100

    数据准备（假设）
    你数据库里 products 表中，有一部分商品字段可以标记为 is_seckill = true，表示这是秒杀商品。

2️⃣ 使用 Lua 脚本实现原子扣减逻辑
    
    为什么用 Lua？
    我们不能直接用普通 Redis 操作：
    stock, _ := rdb.Get(ctx, key).Int()
    if stock > 0 {
        rdb.Decr(ctx, key)
    }
    因为这是两个 Redis 命令，非原子操作，高并发下还是会出现库存变成负数。
    所以：我们必须把库存判断 + 扣减，封装在一段 Lua 脚本里执行，Redis 会保证整个 Lua 脚本原子执行 ✅

### 一些疑问 或者 补充问题
为什么需要 Redis + Lua 做秒杀库存控制？
1. 高并发场景下，MySQL 性能瓶颈明显：
* 秒杀时千万用户并发下单，频繁操作 MySQL 的 stock-- 会导致锁争用、写入瓶颈。
2. Redis 是单线程 + 内存级别操作，非常快：
* 用 Redis 来“挡住”第一波请求，能大幅减轻数据库压力。
3. Lua 脚本可以在 Redis 中一次性执行多个命令，保证原子性：
* Redis 本身不支持事务回滚，而 Lua 脚本是官方推荐的原子操作方式。