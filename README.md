# go-kit-demo
1. [serverA](https://github.com/suyuanhxx/go-kit-demo/tree/master/serverA) go-kit demo [官方demo](https://gokit.io/examples/stringsvc.html)
2. [serverB](https://github.com/suyuanhxx/go-kit-demo/tree/master/serverB) go grpc jsonRpc demo
3. [serverC](https://github.com/suyuanhxx/go-kit-demo/tree/master/serverC) go-kit demo 使用go-kit简易微服务demo ==》实现rpc功能，暂时不能自动代理...
4. 使用指南（服务发现，负载均衡）
    1. cd .../serverA
    2. `go run *.go -listen=:8001`
    3. 在新建shell窗口运行：`go run *.go -listen=:8002 -proxy=127.0.0.1:8001`
    4. 新建shell窗口运行： `curl -d"{\"s\":\"abcd\"}" 127.0.0.1:8001/uppercase`
