package cloudflare_test

import (
	"testing"

	"github.com/chenjl-ops/go-lib/cloudflare"
)

const (
	ApiKey       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	ZoneId       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	ZoneRecordId = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

func TestDnsZone_GetZoneInfo(t *testing.T) {
	dnsZone := &cloudflare.DnsZone{
		Key: cloudflare.Key{
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
	dnsZone := &cloudflare.DnsZone{
		Key: cloudflare.Key{
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
	dnsRecord := &cloudflare.DnsRecord{
		Zone: cloudflare.DnsZone{ZoneId: ZoneId, Key: cloudflare.Key{ApiKey: ApiKey}},
		Id:   ZoneRecordId,
	}
	recordDetail, err := dnsRecord.GETRecordDetails()
	if err != nil {
		t.Error(err)
	}
	t.Log(recordDetail)
}
