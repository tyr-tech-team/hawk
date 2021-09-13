# hawk

# Lib
## Broker
## TSMM - 排程管理大師管理器

Tyr Schedule Master Manager

### Usage
#### Client
You can new a `Client` with a `Broker`.
```go
import(
  "context"
  "time"

  "github.com/tyr-tech-team/hawk/broker"
  "github.com/tyr-tech-team/hawk/broker/natsstreaming"
  tsmm_client "github.com/tyr-tech-team/hawk/tsmm/client"
)

// New nats streaming Broker
broker := broker.NewBroker(natsstreaming.New())

// New Tsmm Client
client := tsmm_client.NewTsmmClient(broker)
```

In the client, you can send add, update, cancel schedule request like this.  
Because `Tsmm` uses an asynchronous transfer, you will not receive a response immediately.
```go
// Add schedule
client.AddSchedule(context.Background(), &tsmm_client.AddScheduleReq{})

// Update schedule
client.UpdateSchedule(context.Background(), &tsmm_client.UpdateScheduleReq{})

// Cancel schedule
client.CancelSchedule(context.Background(), &tsmm_client.CancelScheduleReq{})
```
#### Server
New a `Server` with the `Broker`.
```go
import(
  "context"
  "time"

  "github.com/tyr-tech-team/hawk/broker"
  "github.com/tyr-tech-team/hawk/broker/natsstreaming"
  tsmm_server "github.com/tyr-tech-team/hawk/tsmm/server"
)

// New nats streaming Broker
broker := broker.NewBroker(natsstreaming.New())

// New Tsmm Server
server := tsmm_server.NewTsmmServer(broker)
```
There is an option you can change the queue group name.
If it is empty or you don't use this option, it will be `tsmm`.
```go
server := tsmm_server.NewTsmmServer(
  broker,
  tsmm_server.SetQueueName("my-queue-name"),
)
```
Use `Server` to subscribe the topic you want to listen or existing add, delete, update reply.
You need to give an closure func to make sure you can run somthing when Tsmm receive response.
> :warning: **Make sure the thread these subscribe function stay in should not be stoped.**
```go
// subscribe add schedule reply
a.tsmmServer.AddScheduleReply(
  func(header map[string]interface{}, res *tsmm_server.AddScheduleRes) error {
    // ...somthing...
    return nil
  },
)

// subscribe update schedule reply
a.tsmmServer.UpdateScheduleReply(
  func(header map[string]interface{}, res *tsmm_server.UpdateScheduleRes) error {
    // ...somthing...
    return nil
  },
)

// subscribe cancel schedule reply
a.tsmmServer.CancelScheduleReply(
  func(header map[string]interface{}, res *tsmm_server.CancelScheduleRes) error {
    // ...somthing...
    return nil
  },
)

// subscribe a customized topic
a.tsmmServer.AddTopic(
  "my-topic",
  func(header map[string]interface{}, res []byte) error {
    // ...somthing...
    return nil
  },
)
```
## config 配置檔
## status 狀態碼
## trace 追蹤模組
## middleware 中間件
### http
### grpc
## env 常數
## service 服務元件模組
### register
## pkg 第三方包
### consul  
