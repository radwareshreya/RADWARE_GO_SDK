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
	Index           string `json:"Index"`
	IpAddr          string `json:"IpAddr"`
	Weight          int    `json:"Weight"`
	MaxConns        int    `json:"MaxConns"`
	TimeOut         int    `json:"TimeOut"`
	PingInterval    int    `json:"PingInterval"`
	FailRetry       int    `json:"FailRetry"`
	SuccRetry       int    `json:"SuccRetry"`
	Name            string `json:"Name"`
	State           int    `json:"State"`
	DeleteStatus    int	`json:"DeleteStatus"`
	NxtRportIdx     int    `json:"NxtRportIdx"`
	NxtBuddyIdx     int    `json:"NxtBuddyIdx"`
	LLBType         int    `json:"LLBType"`
	Copy            string `json:"Copy"`
	PortsIngress    string `json:"PortsIngress"`
	PortsEgress     string `json:"PortsEgress"`
	AddPortsIngress int    `json:"AddPortsIngress"`
	RemPortsIngress int    `json:"RemPortsIngress"`
	AddPortsEgress  int    `json:"AddPortsEgress"`
	RemPortsEgress  int    `json:"RemPortsEgress"`
	Type            int    `json:"Type"`
	AddUrl          int    `json:"AddUrl"`
	RemUrl          int    `json:"RemUrl"`
	Cookie          int    `json:"Cookie"`
	ExcludeStr      int    `json:"ExcludeStr"`
	Submac          int    `json:"Submac"`
	Idsport         int    `json:"Idsport"`
	IpVer           int    `json:"IpVer"`
	Ipv6Addr        string `json:"Ipv6Addr"`
	VlanIngress     int    `json:"VlanIngress"`
	VlanEgress      int    `json:"VlanEgress"`
	EgressIf        int    `json:"EgressIf"`
	SecType         int    `json:"SecType"`
	IngressIf       int    `json:"IngressIf"`
	SecDeviceFlag   int    `json:"SecDeviceFlag"`
	Ingport         int    `json:"Ingport"`
}
// ServerGroupItem -
type ServerGroupItem struct {
	Index              string `json:"Index"`
	AddServer          string `json:"AddServer"`
	RemoveServer       string `json:"RemoveServer"`
	HealthCheckUrl     string `json:"HealthCheckUrl"`
	Name               string `json:"Name"`
	HealthCheckLayer   int    `json:"HealthCheckLayer"`
	Metric             int    `json:"Metric"`
	BackupServer       string `json:"BackupServer"`
	BackupGroup        string `json:"BackupGroup"`
	RealThreshold      int    `json:"RealThreshold"`
	VipHealthCheck     int    `json:"VipHealthCheck"`
	IdsState           int    `json:"IdsState"`
	IdsPort            int    `json:"IdsPort"`
	DeleteStatus       int    `json:"DeleteStatus"`
	IdsFlood           int    `json:"IdsFlood"`
	MinmissHash        int    `json:"MinmissHash"`
	PhashMask          string `json:"PhashMask"`
	Rmetric            int    `json:"Rmetric"`
	HealthCheckFormula string `json:"HealthCheckFormula"`
	OperatorAccess     int    `json:"OperatorAccess"`
	Wlm                int    `json:"Wlm"`
	RadiusAuthenString string `json:"RadiusAuthenString"`
	SecBackupGroup     string `json:"SecBackupGroup"`
	Slowstart          int    `json:"Slowstart"`
	MinThreshold       int    `json:"MinThreshold"`
	MaxThreshold       int    `json:"MaxThreshold"`
	IpVer              int    `json:"IpVer"`
	Backup             string `json:"Backup"`
	BackupType         int    `json:"BackupType"`
	HealthID           string `json:"HealthID"`
	PhashPrefixLength  int    `json:"PhashPrefixLength"`
	Type               int    `json:"Type"`
	Copy               string `json:"Copy"`
	IdsChain           int    `json:"IdsChain"`
	SecType            int    `json:"SecType"`
	SecDeviceFlag      int    `json:"SecDeviceFlag"`
	MaxConEx           int    `json:"MaxConEx"`
}

type VirtualServer struct {
	Items []VirtualServerItem `json:"SlbNewCfgEnhVirtServerTable,omitempty"`
}
 
// VirtualServerItem -
type VirtualServerItem struct {
	Index                    string `json:"Index"`
	VirtServerIpAddress      string `json:"VirtServerIpAddress"`
	VirtServerState          int    `json:"VirtServerState"`
	VirtServerLayer3Only     int    `json:"VirtServerLayer3Only"`
	VirtServerDname          string `json:"VirtServerDname"`
	VirtServerBwmContract    int    `json:"VirtServerBwmContract"`
	VirtServerDelete         int    `json:"VirtServerDelete"`
	VirtServerWeight         int    `json:"VirtServerWeight"`
	VirtServerAvail          int    `json:"VirtServerAvail"`
	VirtServerRule           string `json:"VirtServerRule"`
	VirtServerAddRule        int    `json:"VirtServerAddRule"`
	VirtServerRemoveRule     int    `json:"VirtServerRemoveRule"`
	VirtServerIpVer          int    `json:"VirtServerIpVer"`
	VirtServerIpv6Addr       string `json:"VirtServerIpv6Addr"`
	VirtServerFreeServiceIdx int    `json:"VirtServerFreeServiceIdx"`
	VirtServerVname          string `json:"VirtServerVname"`
	VirtServerCReset         int    `json:"VirtServerCReset"`
	VirtServerSrcNetwork     string `json:"VirtServerSrcNetwork"`
	VirtServerNat            string `json:"VirtServerNat"`
	VirtServerNat6           string `json:"VirtServerNat6"`
	VirtServerIsDnsSecVip    int    `json:"VirtServerIsDnsSecVip"`
	VirtServerAvailPersist   int    `json:"VirtServerAvailPersist"`
	VirtServerWanlink        string `json:"VirtServerWanlink"`
	VirtServerRtSrcMac       int    `json:"VirtServerRtSrcMac"`
	VirtServerCreationType   int    `json:"irtServerCreationType"`
}
type VirtualService struct {
	Items []VirtualServiceItem `json:"SlbNewCfgEnhVirtServicesTable,omitempty"`
}
 
// VirtualServiceItem -
type VirtualServiceItem struct {
	ServIndex     string `json:"ServIndex"`
	Index         int    `json:"Index"`
	VirtPort      int    `json:"VirtPort"`
	RealPort      int    `json:"RealPort"`
	UDPBalance    int    `json:"UDPBalance"`
	BwmContract   int    `json:"BwmContract"`
	DirServerRtn  int    `json:"DirServerRtn"`
	RtspUrlParse  int    `json:"RtspUrlParse"`
	DBind         int    `json:"DBind"`
	FtpParsing    int    `json:"FtpParsing"`
	RemapUDPFrags int    `json:"RemapUDPFrags"`
	DnsSlb        int    `json:"DnsSlb"`
	ResponseCount int    `json:"ResponseCount"`
	PBind         int    `json:"PBind"`
	Coffset       int    `json:"Coffset"`
	Clength       int    `json:"Clength"`
	UriCookie     int    `json:"UriCookie"`
	CookieMode    int    `json:"CookieMode"`
	HttpSlb       int    `json:"HttpSlb"`
	HttpSlbOption int    `json:"HttpSlbOption"`
	HttpSlb2      int    `json:"HttpSlb2"`
	DeleteStatus  int    `json:"DeleteStatus"`
	Apm           int    `json:"Apm"`
	NonHTTP       int    `json:"NonHTTP"`
	IpRep         int    `json:"IpRep"`
	CdnProxy      int    `json:"CdnProxy"`
	Status        int    `json:"Status"`
	RtSrcTnl      int    `json:"RtSrcTnl"`
	Sideband      string `json:"Sideband"`
}

// SslPolicyItem -
type SslPolicyItem struct {
	NameIdIndex    string `json:"NameIdIndex"`
	Name           string `json:"Name"`
	AdminStatus    int    `json:"AdminStatus"`
	Fessl          int    `json:"Fessl"`
	Bessl          int    `json:"Bessl"`
	FESslv3Version int    `json:"FESslv3Version"`
}

// Http2PolicyItem -
type Http2PolicyItem struct {
	NameIdIndex       string `json:"NameIdIndex"`
	Name              string `json:"Name"`
	AdminStatus       int    `json:"AdminStatus"`
	Streams           int    `json:"Streams"`
	Idle              int    `json:"Idle"`
	EnaInsert         int    `json:"EnaInsert"`
	Header            string `json:"Header"`
	EnaServerPush     int    `json:"EnaServerPush"`
	HpackSize         string `json:"HpackSize"`
	DeleteStatus      int    `json:"DeleteStatus"`
	BackendStatus     int    `json:"BackendStatus"`
	BackendStreams    int    `json:"BackendStreams"`
	BackendHpackSize  string `json:"BackendHpackSize"`
	BackendServerPush int    `json:"BackendServerPush"`
}

