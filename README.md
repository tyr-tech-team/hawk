# hawk





# Lib
## TSMM - 排程管理大師管理器

Tyr Schedule Master Manager

### Usage

```go=
import(
  "context"
  "time"

  "github.com/tyr-tech-team/hawk/broker"
  "github.com/tyr-tech-team/hawk/broker/natsstreaming"
  tsmm_client "github.com/tyr-tech-team/hawk/tsmm/client"
)

// New nats streaming Broker
natsStreamingHost := "localhost:5000"
natsStreamingClusterID := "test-cluster"

natsStreaming := natsstreaming.New(
    natsstreaming.SetURL(natsStreamingHost),
    natsstreaming.SetStanClusterID(natsStreamingClusterID),
)
broker := broker.NewBroker(natsStreaming)

// New Tsmm Client
client := tsmm_client.NewTsmmClient()

// AddSchedule

// Add header
header := make(map[string]interface{}, 0)
header["itemIDList"] = []string{"1234"}
// Add schedule request
addReq := &tsmm_client.AddScheduleReq{
  Header: header,
  ScheduleList: []\*tsmm_client.AddScheduleDetail{
    {
      Topic:      "update-item",
      Webhook:    "",
      Time:       "@every 1h",
      StartTime:  time.Time{},
      StopTime:   time.Time{},
      TotalTimes: 1,
      TimesType:  1,
      Action:     "update",
      Name:       "name",
      Data:       []byte{"1234"},
      Memo:       "This is the memo",
    }
  }
}

// Add schedule
client.AddSchedule(context.Background(), addReq)

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
