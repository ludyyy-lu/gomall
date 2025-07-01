# gomall
电商后端系统
* * *
## 给自己看
认真看一看鉴权
认真看一看分页
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