# gomall
电商后端系统
* * *
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

* * *


# 疑问
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