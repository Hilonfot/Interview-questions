# Grpc相关

### 相较于restful的优势
- grpc使用了二进制编码，所以它比 JSON/HTTP 更快
- 清晰的接口规范
- 对流式传输的支持

### 数据交互方式
- 二进制protobuf 

### 限流（通过流模式传输时，发送方数据量过大，会发生什么？）

### protobuf和json的对比
- protobuf性能比JSON快5倍
- 在对大小和性能不敏感时用JSON，敏感时用Protobuf 
### grpc负载均衡的实现
- 服务发现（go-micro 使用selector 随机和轮询）