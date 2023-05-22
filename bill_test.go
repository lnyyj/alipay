package alipay_test

import (
	"github.com/smartwalle/alipay/v3"
	"testing"
)

func TestClient_BillDownloadURLQuery(t *testing.T) {
	t.Log("========== BillDownloadURLQuery ==========")
	var p = alipay.BillDownloadURLQuery{}
	p.BillType = "trade"
	p.BillDate = "2019-01-01"
	rsp, err := client.BillDownloadURLQuery(p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.Failed() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_BillBalanceQuery(t *testing.T) {
	t.Log("========== BillBalanceQuery ==========")
	var p = alipay.BillBalanceQuery{}
	rsp, err := client.BillBalanceQuery(p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.Failed() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_BillAccountLogQuery(t *testing.T) {
	t.Log("========== BillAccountLogQuery ==========")
	var p = alipay.BillAccountLogQuery{}
	rsp, err := client.BillAccountLogQuery(p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.Failed() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
