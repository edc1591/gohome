package hap

import (
	"reflect"
	"testing"
)

func TestMDNS(t *testing.T) {
	mdns := NewMDNSService("My MDNS Service", "1234", 5010)
	expect := []string{
		"pv=1.0",
		"id=1234",
		"c#=1",
		"s#=1",
		"sf=1",
		"ff=0",
		"md=My MDNS Service",
		"ci=1",
	}
	if x := mdns.txtRecords(); reflect.DeepEqual(x, expect) == false {
		t.Fatal(expect)
	}
}
