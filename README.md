# hawk
`hawk` is a toolkit that provides service building, startup and other functions. Here are the main features.
  * [broker](##Broker)
  * [TSMM](##TSMM)
  * [config](##config)
  * [status](##status)
  * [trace](##trace)
  * [middleware](##middleware)
  * [register](##register)
  * [consul](##consul)

## Getting Started
### Installation
Run the following command under your project
```shell
go get -u github.com/tyr-tech-team/hawk
```

## Features
### TSMM
###### `Tyr Schedule Master Manager!`
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
> :warning: **Make sure the thread these subscribe function stay in should not be stopped.**
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
### broker

### config

### status

### trace

### middleware

### env

### service

### register

### consul