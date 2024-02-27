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
