package status

import "fmt"

// ServiceCode - 服務代號
type ServiceCode int64

func (s ServiceCode) String() string {
	return fmt.Sprintf("%03d", s)
}

//
const (
	ServiceNONE          ServiceCode = iota // 0
	ServiceNormal                           // 1
	ServiceAuth                             // 2
	ServiceCard                             // 3
	ServiceEventLog                         // 4
	ServiceItem                             // 5
	ServiceBrand                            // 6
	ServiceMember                           // 7
	ServiceNFCReader                        // 8
	ServiceStorage                          // 9
	ServiceTransaction                      // 10
	ServiceUser                             // 11
	ServiceWebSite                          // 12
	ServiceInitService                      // 13
	ServiceSellOrder                        // 14
	ServiceBuyOrder                         // 15
	ServiceSMS                              // 16
	ServiceIDCard                           // 17
	ServiceRBAC                             // 18
	ServiceBgmGateway                       // 19
	ServiceOfficeGateway                    // 20
	ServiceBlog                             // 21
	ServiceCategory                         // 22
	ServiceCatalog                          // 23
	ServiceQuickFilter                      // 24
	ServiceECPay                            // 25
	ServiceZeroCard                         // 26
	ServiceCart                             // 27
	ServicePromo                            // 28
	ServiceReceipt                          // 29
	ServiceWebKit                           // 30
)
