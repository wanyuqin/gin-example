# gin框架

## 简介

Gin 是用 Go 语言编写的轻量级 Web 框架，具有快速、高性能、灵活和易于使用的特点。它提供了路由、中间件、错误处理等功能，适用于构建各种类型的 Web 应用程序和 API。

> github：[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
>

> 官方文档 [https://gin-gonic.com/zh-cn/docs/](https://gin-gonic.com/zh-cn/docs/)
>

## 快速入门

### 要求

- GO1.16及以上版本

### 下载并安装

```bash
go get -u github.com/gin-gonic/gin
```

### 导入

```go
import "github.com/gin-gonic/gin"
```

### 快速开始

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个server
	server := gin.Default()
	// 定义一个路由
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	// 启动服务
	server.Run(":8082")
}

```

## 路由

在我们日常的项目中，我们会为我们的web服务对外提供很多接口，同时对每个接口都会有对应的分类来方便我们后期的管理。gin在这方面也为我们很贴心的提供了对应的方法。同时gin也很好的为我们提供了Restful风格API支持。接下来我们就来看看如何对其进行使用。

### 路由分组

我们可以使用gin提供的group方法对路由进行分组

```go
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
```

```go
	server := gin.Default()
	
	v1 := server.Group("/v1")
	
	v2 := server.Group("/v2")
	
	userGroup:= v1.Group("/users")

	
	
```

在这里我们三次使用了group方法进行路由分组，但是和v1、v2不同的是，userGroup是我们基于v1来创建的，那么创建出来的user分组就会继承v1的path，整体创建出来的效果就是

```go
localhost:8080/v1
localhost:8080/v2
localhost:8080/v1/users
```

### Restful API

RESTful API（Representational State Transfer API）是一种基于REST架构风格的应用程序接口。REST是一种轻量级的、无状态的通信协议，广泛应用于Web服务。RESTful API通过HTTP协议进行通信，利用HTTP方法（如GET、POST、PUT、DELETE等）来执行不同的操作

```go
userGroup := v1.Group("/users")
	{ // 获取用户
		userGroup.GET("/:id", func(context *gin.Context) {
			
		})
		// 创建用户
		userGroup.POST("/", func(context *gin.Context) {

		})
		// 删除用户
		userGroup.DELETE("/:id", func(context *gin.Context) {

		})
		// 更新用户
		userGroup.PUT("/:id", func(context *gin.Context) {

		})
	}
```

我们可以直接使用gin内置的GET、POST、PUT、DELETE方法来表示不同的HTTP动作，这样我们就可以很方便的定义RESTful风格的API，当然如果你不需要也可以只使用任意的方法去定义你的路由

## 参数传递

在日常开发的过程中，我们用的比较多参数传递方式有，Param，Json，接下来我就这两类方式来进行示例

### Query

gin为我们提供了多种方式去解析query和form表单的数据

在Go语言的Gin框架中，`ShouldBindQuery`和`Query`是用于处理HTTP请求中查询参数的两种方法。它们有不同的使用场景和功能。以下是它们的区别和用法：

### **ShouldBindQuery**

`ShouldBindQuery`方法用于将查询参数绑定到结构体中。它会根据查询参数的名称和结构体字段的标签将参数值填充到结构体字段中。这种方法适用于需要将多个查询参数绑定到一个结构体的情况。

用法示例：

```go

type User struct {
	Name string `form:"name"`
	Age   uint64 `form:"age"`
}

server.POST("/struct", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBindQuery(&user); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
		return

	})
```

在这个例子中，`ShouldBindQuery`会将查询参数`name`和`age`绑定到`UserQuery`结构体的字段中。如果参数绑定成功，可以通过结构体字段访问这些参数。

### **Query**

`Query`方法用于获取单个查询参数的值，返回值类型是字符串。如果查询参数不存在，返回一个空字符串。它适用于需要获取单个查询参数的情况。

用法示例：

```go
	server := gin.Default()
	server.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
```

在这个例子中，`Query`方法分别获取查询参数`name`和`age`的值。由于返回值是字符串，如果需要将其转换为其他类型（如整数），需要手动进行转换。

### 区别总结

1. **绑定方式**：
    - `ShouldBindQuery`：用于将查询参数绑定到结构体中，适合处理多个参数，并且参数的类型可以是多种数据类型（如字符串、整数等）。
    - `Query`：用于获取单个查询参数的值，返回值为字符串，需要手动进行类型转换。
2. **适用场景**：
    - `ShouldBindQuery`：适用于参数较多且需要绑定到结构体的情况。
    - `Query`：适用于简单的单个参数获取的情况。
3. **错误处理**：
    - `ShouldBindQuery`：如果绑定失败，会返回一个错误，可以进行错误处理。
    - `Query`：如果查询参数不存在，返回空字符串，不会返回错误。

通过这两种方法，Gin框架提供了灵活的查询参数处理方式，可以根据具体需求选择合适的方法。

### **ShouldBindJson**

在Gin框架中，可以使用`ShouldBindJSON`方法将请求体中的JSON数据绑定到结构体中。这个方法非常适合处理客户端发送的JSON格式的请求数据。以下是详细介绍和示例：

`ShouldBindJSON`方法用于将请求体中的JSON数据解析并绑定到结构体中。它会根据JSON字段的名称和结构体字段的标签进行匹配，并将JSON数据填充到结构体的对应字段中。

### 用法示例

假设我们有一个处理用户信息的API，客户端发送的请求体是JSON格式的数据：

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30
}

```

我们可以定义一个结构体来表示这个用户信息，然后使用`ShouldBindJSON`将请求体中的数据绑定到这个结构体中。

```go
type User struct {
	Name  string `form:"name" json:"name"`
	Age   uint64 `form:"age" json:"age"`
	Email string `json:"email" `
}

	server.POST("/users", func(c *gin.Context) {
		u := User{}
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, u)

	})

```

### 详细说明

1. **定义结构体**：首先，我们定义了一个`User`结构体，其中包含三个字段：`Name`、`Email`和`Age`。每个字段都使用了`json`标签来指定对应的JSON字段名称。
2. **处理请求**：在处理函数中，我们创建了一个`User`结构体实例，并使用`ShouldBindJSON`方法将请求体中的JSON数据绑定到这个实例。如果绑定过程中出现错误（例如JSON格式不正确或缺少必要字段），我们返回一个400错误响应，并在响应体中包含错误信息。
3. **返回响应**：如果数据绑定成功，我们返回一个200成功响应，并在响应体中包含一个消息和绑定后的用户数据。

### 错误处理

`ShouldBindJSON`方法会返回一个错误对象，如果JSON解析失败（例如请求体不是有效的JSON格式）或者数据绑定失败（例如JSON字段类型不匹配），可以通过检查错误对象来处理这些情况。

### **BindJson**

`BindJson`  的使用效果和`ShouldBindJSON` 基本是一致的，都是为了绑定提交的json数据，唯一不同的是当出现了错误的时候，`BindJson` 会自动返回状态为400的响应

```go
type User struct {
	Name  string `form:"name" json:"name"`
	Age   uint64 `form:"age" json:"age"`
	Email string `json:"email" `
}
server.POST("/users", func(c *gin.Context) {
		u := User{}
		if err := c.BindJSON(&u); err != nil {
			return
		}

		c.JSON(http.StatusOK, u)

	})
```

## 渲染

在Gin框架中，`LoadHTMLFiles`和 `LoadHTMLGlob` 都是用于加载HTML模板的方法。这两个方法有不同的使用场景和功能。以下是它们的区别和用法：

### LoadHTMLFiles

`LoadHTMLFiles` 方法用于加载一个或多个具体路径的HTML文件。每个文件的路径需要单独指定，适用于需要精确控制加载哪些模板文件的情况。

用法示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // 加载具体的模板文件
    router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

    router.GET("/page1", func(c *gin.Context) {
        c.HTML(http.StatusOK, "template1.html", gin.H{
            "title": "Page 1",
        })
    })

    router.GET("/page2", func(c *gin.Context) {
        c.HTML(http.StatusOK, "template2.html", gin.H{
            "title": "Page 2",
        })
    })

    router.Run(":8080")
}

```

在这个例子中，我们明确指定了要加载的模板文件 `template1.html` 和 `template2.html`。

### LoadHTMLGlob

`LoadHTMLGlob` 方法用于加载符合指定模式的HTML文件。它使用通配符匹配文件路径，适用于需要一次性加载某个目录下的所有模板文件的情况。

用法示例：

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    // 加载指定模式的模板文件
    router.LoadHTMLGlob("templates/*")

    router.GET("/page1", func(c *gin.Context) {
        c.HTML(http.StatusOK, "template1.html", gin.H{
            "title": "Page 1",
        })
    })

    router.GET("/page2", func(c *gin.Context) {
        c.HTML(http.StatusOK, "template2.html", gin.H{
            "title": "Page 2",
        })
    })

    router.Run(":8080")
}

```

在这个例子中，我们使用了 `LoadHTMLGlob("templates/*")` 来加载 `templates` 目录下的所有HTML文件。

### 区别总结

1. **文件加载方式**：
    - `LoadHTMLFiles`：逐个指定文件路径，适用于需要加载特定模板文件的情况。
    - `LoadHTMLGlob`：使用通配符模式加载，适用于需要一次性加载某个目录下所有模板文件的情况。
2. **适用场景**：
    - `LoadHTMLFiles`：当需要精确控制加载哪些模板文件时，适合使用这个方法。
    - `LoadHTMLGlob`：当需要加载一个目录下的所有模板文件时，使用这个方法更加方便。

通过这两个方法，Gin框架提供了灵活的模板文件加载机制，可以根据具体需求选择合适的方法。

## 中间件

## 会话控制

## 其他