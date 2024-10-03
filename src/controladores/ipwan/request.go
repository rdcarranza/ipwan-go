package ipwan

type Request struct {
	WanIpV4          string `xml:"WanIPAddress"`
	PrimaryIPv4Dns   string `xml:"PrimaryDns"`
	SecondaryIPv4Dns string `xml:"SecondaryDns"`
	WanIpV6          string `xml:"WanIPv6Address"`
	PrimaryIPv6Dns   string `xml:"PrimaryIPv6Dns"`
	SecondaryIPv6Dns string `xml:"SecondaryIPv6Dns"`
}
