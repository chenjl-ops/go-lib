package cloudflare

type Key struct {
	ApiKey string `json:"api_key"`
}

type DnsZone struct {
	Key    Key    `json:"key"`
	ZoneId string `json:"zone_id"`
}

type DnsRecord struct {
	Zone DnsZone `json:"zone"`
	Id   string  `json:"id"`
}
