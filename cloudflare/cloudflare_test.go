package cloudflare

import "testing"

const (
	ApiKey       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	ZoneId       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	ZoneRecordId = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

func TestDnsZone_GetZoneInfo(t *testing.T) {
	dnsZone := &DnsZone{
		Key: Key{
			ApiKey: ApiKey,
		},
		ZoneId: ZoneId,
	}
	zoneInfo, err := dnsZone.GetZoneInfo()
	if err != nil {
		t.Error(err)
	}
	t.Log(zoneInfo)

}

func TestDnsRecord_GetZoneRecords(t *testing.T) {
	dnsZone := &DnsZone{
		Key: Key{
			ApiKey: ApiKey,
		},
		ZoneId: ZoneId,
	}

	zoneRecords, err := dnsZone.GetZoneRecords()
	if err != nil {
		t.Error(err)
	}
	t.Log(zoneRecords)
}

func TestDnsRecord_GETRecordDetails(t *testing.T) {
	dnsRecord := &DnsRecord{
		Zone: DnsZone{ZoneId: ZoneId, Key: Key{ApiKey: ApiKey}},
		Id:   ZoneRecordId,
	}
	recordDetail, err := dnsRecord.GETRecordDetails()
	if err != nil {
		t.Error(err)
	}
	t.Log(recordDetail)
}
