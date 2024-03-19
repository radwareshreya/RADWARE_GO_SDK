package RadwareGoSDK

// CliCommand-
type CliCommand struct {
    Items []CliCommandItem `json:"agAlteonCliCommand,omitempty"`
}
 
// CliCommandItem -
type CliCommandItem struct {
    CliCommand string `json:"agAlteonCliCommand"`
}

type RealServer struct {
	Items []RealServerItem `json:"SlbNewCfgEnhRealServerTable,omitempty"`
}

// RealServerItem -
type RealServerItem struct {
	Index        string `json:"Index"`
	IpAddr       string `json:"IpAddr"`
	Weight       int    `json:"Weight"`
	MaxConns     int    `json:"MaxConns"`
	TimeOut      int    `json:"TimeOut"`
	PingInterval int    `json:"PingInterval"`
	FailRetry    int    `json:"FailRetry"`
	SuccRetry    int    `json:"SuccRetry"`
	Name         string `json:"Name"`
	State        int    `json:"State"`
}

// ServerGroupItem -
type ServerGroupItem struct {
	Index          string `json:"Index"`
	AddServer      string `json:"AddServer"`
	RemoveServer   string `json:"RemoveServer"`
	HealthCheckUrl string `json:"HealthCheckUrl"`
	Name           string `json:"Name"`
	HealthCheckLayer int `json:"HealthCheckLayer"`
}

type VirtualServer struct {
	Items []VirtualServerItem `json:"SlbNewCfgEnhVirtServerTable,omitempty"`
}
 
// VirtualServerItem -
type VirtualServerItem struct {
	Index        string `json:"Index"`
	VirtServerIpAddress       string `json:"VirtServerIpAddress"`
	VirtServerState        int    `json:"VirtServerState"`
}

type VirtualService struct {
	Items []VirtualServiceItem `json:"SlbNewCfgEnhVirtServicesTable,omitempty"`
}
 
// VirtualServiceItem -
type VirtualServiceItem struct {
	ServIndex        string `json:"ServIndex"`
	Index        int `json:"Index"`
	VirtPort       int `json:"VirtPort"`
	RealPort        int    `json:"RealPort"`
}
