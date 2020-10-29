package status

import "fmt"

// ServiceCode - 服務代號
type ServiceCode int64

func (s ServiceCode) String() string {
	return fmt.Sprintf("%03d", s)
}

//
const (
	ServiceNONE ServiceCode = iota
	ServiceAuth
	ServiceCard
	ServiceEventLog
	ServiceItem
	ServiceMember
	ServiceNFCReader
	ServiceOrder
	ServiceStorage
	ServiceTransaction
	ServiceUser
	ServiceWebSite
	ServiceInitService
	ServiceBrand
)
