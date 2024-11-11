package cloudflare

import "github.com/chenjl-ops/go-lib/requests"

const URL = "https://api.cloudflare.com"

func (k *Key) getKey() string {
	return k.ApiKey
}

func (k *Key) getAuthor() string {
	return "Bearer " + k.getKey()
}

func (k *Key) getHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": k.getAuthor(),
	}
}

func (d *DnsZone) GetZoneInfo() (map[string]interface{}, error) {
	headers := d.Key.getHeaders()

	data := make(map[string]interface{})
	err := requests.Request(URL+"/client/v4/zones/"+d.ZoneId, "GET", headers, nil, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (d *DnsZone) GetZoneRecords() (map[string]interface{}, error) {
	headers := d.Key.getHeaders()
	data := make(map[string]interface{})

	err := requests.Request(URL+"/client/v4/zones/"+d.ZoneId+"/dns_records", "GET", headers, nil, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

/*
CreateZoneRecord
docs: https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-create-dns-record
*/
func (d *DnsRecord) CreateZoneRecord(requestData map[string]interface{}) (map[string]interface{}, error) {
	headers := d.Zone.Key.getHeaders()
	data := make(map[string]interface{})

	err := requests.Request(URL+"/client/v4/zones/"+d.Zone.ZoneId+"/dns_records", "POST", headers, requestData, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (d *DnsRecord) GETRecordDetails() (map[string]interface{}, error) {
	headers := d.Zone.Key.getHeaders()
	data := make(map[string]interface{})

	err := requests.Request(URL+"/client/v4/zones/"+d.Zone.ZoneId+"/dns_records/"+d.Id, "GET", headers, nil, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (d *DnsRecord) UpdateZoneRecord(requestData map[string]interface{}, isOverwrite bool) (map[string]interface{}, error) {
	headers := d.Zone.Key.getHeaders()
	data := make(map[string]interface{})

	var err error
	if isOverwrite {
		err = requests.Request(URL+"/client/v4/zones/"+d.Zone.ZoneId+"/dns_records/"+d.Id, "PUT", headers, requestData, &data)
	} else {
		err = requests.Request(URL+"/client/v4/zones/"+d.Zone.ZoneId+"/dns_records/"+d.Id, "PATCH", headers, requestData, &data)
	}
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (d *DnsRecord) DeleteZoneRecord() (map[string]interface{}, error) {
	headers := d.Zone.Key.getHeaders()
	data := make(map[string]interface{})

	err := requests.Request(URL+"/client/v4/zones/"+d.Zone.ZoneId+"/dns_records/"+d.Id, "DELETE", headers, nil, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
