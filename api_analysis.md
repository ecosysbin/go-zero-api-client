# API 文件解析和 Go HTTP Client 生成

## API 文件结构分析

### 1. 数据类型定义

API 文件定义了以下请求和响应类型：

#### 请求类型 (Request Types)
- `GetUserRequest`: 获取用户信息的请求
- `AddUserRequest`: 添加用户的请求
- `DeleteUserRequest`: 删除用户的请求
- `UpdateUserRequest`: 更新用户的请求

#### 响应类型 (Response Types)
- `GetUserResponse`: 获取用户信息的响应
- `AddUserResponse`: 添加用户的响应
- `DeleteUserResponse`: 删除用户的响应
- `UpdateUserResponse`: 更新用户的响应

### 2. 字段标签分析

#### 标签类型说明：
- `header:"authorization"`: 表示该字段在 HTTP 请求头中传递
- `path:"name"`: 表示该字段在 URL 路径中作为路径参数
- `form:"delete,optional"`: 表示该字段在查询参数中传递，且为可选
- `json:"field_name"`: 表示该字段在 JSON 请求体或响应体中

### 3. 服务端点定义

API 定义了一个名为 `user-api` 的服务，包含 4 个端点：

| 方法 | 路径 | 处理器 | 功能 |
|------|------|--------|------|
| POST | `/v1/user` | AddUserHandler | 创建新用户 |
| DELETE | `/v1/user/:name` | DeleteUserHandler | 删除指定用户 |
| PUT | `/v1/user/:name` | UpdateUserHandler | 更新指定用户 |
| GET | `/v1/user/:name` | GetUserHandler | 获取指定用户信息 |

## 生成的 Go HTTP Client 特性

### 1. 客户端结构

```go
type UserAPIClient struct {
    BaseURL    string
    HTTPClient *http.Client
}
```

### 2. 核心功能

#### 2.1 自动处理不同类型的参数
- **路径参数**: 自动替换 URL 中的 `:name` 占位符
- **请求头**: 自动设置 Authorization 头
- **查询参数**: 自动处理可选的 form 参数
- **JSON 请求体**: 自动序列化请求数据

#### 2.2 错误处理
- 网络错误处理
- HTTP 状态码检查
- JSON 序列化/反序列化错误处理
- 详细的错误信息

#### 2.3 HTTP 方法映射
- `AddUser()` → POST `/v1/user`
- `GetUser()` → GET `/v1/user/:name`
- `UpdateUser()` → PUT `/v1/user/:name`
- `DeleteUser()` → DELETE `/v1/user/:name`

### 3. 使用示例

```go
// 创建客户端
client := NewUserAPIClient("http://localhost:8080")

// 添加用户
addReq := AddUserRequest{
    Authorization: "Bearer token",
    Name: "john_doe",
    Age: "25",
}
resp, err := client.AddUser(addReq)
```

### 4. 设计特点

#### 4.1 类型安全
- 使用强类型的请求和响应结构体
- 编译时检查参数类型

#### 4.2 易于使用
- 简洁的 API 接口
- 自动处理 HTTP 细节
- 清晰的错误信息

#### 4.3 可扩展性
- 可自定义 HTTP 客户端
- 支持不同的 BaseURL
- 易于添加中间件（如重试、日志等）

### 5. 参数处理细节

#### 5.1 路径参数处理
```go
// :name 参数自动替换
url := fmt.Sprintf("%s/v1/user/%s", c.BaseURL, req.Name)
```

#### 5.2 查询参数处理
```go
// 可选参数的条件添加
if req.Delete {
    q.Set("delete", "true")
}
```

#### 5.3 请求头处理
```go
// 自动设置认证头
httpReq.Header.Set("Authorization", req.Authorization)
```

#### 5.4 JSON 请求体处理
```go
// 只序列化 JSON 标签的字段
body := struct {
    Name string `json:"name"`
    Age  string `json:"age"`
}{
    Name: req.Name,
    Age:  req.Age,
}
```

## 总结

生成的 Go HTTP Client 完全基于 API 文件的定义，提供了类型安全、易于使用的接口来调用用户管理 API。客户端自动处理了 HTTP 请求的各个方面，包括路径参数、查询参数、请求头和 JSON 序列化，使开发者可以专注于业务逻辑而不是 HTTP 通信的细节。