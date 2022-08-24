# go-zero-demo

### 示例用法（在api中调用go-zero-rpc）
1. 目录结构：
```
|—— demo/   // 根目录
****|—— api/
********|—— test.api/
****|—— rpc/
********|—— test.proto/
```

2. 文件内容：
    1. `demo/api/test.api`：
    ```
        type (
            Req {
                Id string `path:"id"`
            }
            Resp {
                Id   string `json:"id"`
                Name string `json:"name"`
            }
        )
        service test {
            @handler testGet
            get /test/:id (Req) returns (Resp)
        }
    ```
    2. `demo/rpc/test.proto`：
    ```
        syntax = "proto3";
        package test;
        option go_package = "./test";
        service Test {
            rpc test(Req) returns(Resp);
        }
        message Req {
            uint32 id = 1;
        }
        message Resp {
            uint32 id = 1;
            string name = 2;
        }
    ```

3. 初始化`go module`：
```
cd demo
go mod init demo
```

4. 生成`api`：
```
cd demo/api
goctl api go -api test.api -dir .
```

5. 生成`rpc`：
```
cd demo/rpc
goctl rpc protoc test.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

6. 代码地址：
[点击访问](https://github.com/job520/go-zero-demo)

7. 启动`rpc`：
```
cd demo/rpc
go run test.go
```

8. 启动`api`：
```
cd demo/api
go run test.go
```

9. 访问测试：
`localhost:8888/test/1`
