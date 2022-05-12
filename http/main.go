package main

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay/apple"
	"github.com/go-pay/gopay/pkg/xlog"
	"sort"
	"strconv"
)

func haha() {
	i := 0
	i++
}
func main() {

	resp, err := apple.VerifyReceipt(context.TODO(), apple.UrlProd, "", "")
	if err != nil {
		xlog.Error(err)
		return
	}
	sort.Slice(resp.LatestReceiptInfo, func(i, j int) bool {
		a, _ := strconv.Atoi(resp.LatestReceiptInfo[i].PurchaseDateTimestamp)
		b, _ := strconv.Atoi(resp.LatestReceiptInfo[j].PurchaseDateTimestamp)
		return a > b
	})
	fmt.Println(resp)
	//
	//// decode signedPayload
	//payload, err := apple.DecodeSignedPayload(signedPayload)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debugf("payload.NotificationType: %s", payload.NotificationType)
	//xlog.Debugf("payload.Subtype: %s", payload.Subtype)
	//xlog.Debugf("payload.NotificationUUID: %s", payload.NotificationUUID)
	//xlog.Debugf("payload.NotificationVersion: %s", payload.NotificationVersion)
	//renewalInfo, err := payload.DecodeRenewalInfo()
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debugf("data.renewalInfo: %+v", renewalInfo)
	//transactionInfo, err := payload.DecodeTransactionInfo()
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//xlog.Debugf("data.transactionInfo: %+v", transactionInfo)

	//client := appstore.New()
	//req := appstore.IAPRequest{
	//	ReceiptData:            receipt,
	//	Password:               "24e2cb3d9d514006b824cdb4c7d21846",
	//	ExcludeOldTransactions: true,
	//}
	//resp := &appstore.IAPResponse{}
	//ctx := context.Background()
	//err := client.Verify(ctx, req, resp)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
