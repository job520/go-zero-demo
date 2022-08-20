# go-zero-demo

### 开发环境搭建
1. 安装`golang`
[参见：源码安装 golang](https://www.jianshu.com/p/a67f070dbc8f)

2. 开启`go module`
`go env -w GO111MODULE=on`

3. 使用代理镜像
` go env -w GOPROXY=https://goproxy.cn,direct`

4. 安装`goctl`
`go install github.com/zeromicro/go-zero/tools/goctl@latest`

5. 安装`protoc & protoc-gen-go`
`goctl env check -i -f --verbose`

6. 安装`mysql`、`redis`、`etcd`（略）

### 常用命令
1. 生成`model`
`goctl model mysql ddl -src tables.sql -dir . -c`

2. 生成`api`
`goctl api go -api myapi.api -dir .`

3. 生成`rpc`
`goctl rpc protoc myrpc.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.`

4. 启动`api`或`rpc`
`go run myprj.go -f etc/myprj.yaml`

### 示例用法（单体服务/微服务）
1. 目录结构：
```
|—— demo/   // 根目录
****|—— model/
********|—— mtable.sql/
****|—— api/
********|—— mapi.api/
****|—— rpc/
********|—— mrpc.proto/
```

2. 文件内容：
    1. `demo/model/mtable.sql`：
    ```
    CREATE TABLE `mtable` (
          `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
          `name` varchar(255) CHARACTER SET latin1 DEFAULT '' COMMENT '名称',
          PRIMARY KEY (`id`),
          KEY `name` (`name`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
    ```
    2. `demo/api/mapi.api`：
    ```
    type (
        Req {
            Id int64 `path:"id"`
        }
        Reply {
            Id int64 `json:"id"`
            Name string `json:"name"`
        }
    )
    service mapi {
        @handler getMapi
        get /mapi/get/:id (Req) returns (Reply)  // 单体服务
        @handler getMrpc
        get /mrpc/get/:id (Req) returns (Reply)  // 微服务
    }
    ```
    3. `demo/rpc/mrpc.proto`：
    ```
    syntax = "proto3";
    package mrpc;
    option go_package = "./mrpc";
    message Req {
        int64 id = 1;
    }
    message Reply {
        int64 id = 1;
        string name = 2;
    }
    service Mrpc {
        rpc GetById (Req) returns (Reply);
    }
    ```

3. 初始化`go module`：
```
cd demo
go mod init demo
```

4. 生成`model`：
```
cd demo/model
goctl model mysql ddl -src mtable.sql -dir . -c
```

5. 生成`api`：
```
cd demo/api
goctl api go -api mapi.api -dir .
```

6. 生成`rpc`：
```
cd demo/rpc
goctl rpc protoc mrpc.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

7. 代码地址：
[点击访问](https://github.com/job520/go-zero-demo)

8. 启动`rpc`：
```
cd demo/rpc
go run mrpc.go -f etc/mrpc.yaml
```

9. 启动`api`：
```
cd demo/api
go run mapi.go -f etc/mapi.yaml
```

10. 访问测试：
```
	curl localhost:8888/mapi/get/1
	{"id":1,"name":"lee"}

	curl localhost:8888/mrpc/get/1
	{"id":1,"name":"lee"}

	curl localhost:8888/mapi/get/2
	api: 未找到记录

	curl localhost:8888/mrpc/get/2
	rpc-客户端: 未找到记录
```
