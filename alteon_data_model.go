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
