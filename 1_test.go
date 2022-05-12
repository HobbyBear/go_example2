package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

//func init() {
//	os.Args[1] = "2"
//}

const (
	mutexLocked           = 1 << iota // mutex is locked 001
	mutexWoken                        // 010
	mutexStarving                     // 100
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

const (
	a = iota
	b
	c = iota
)

func Test(t *testing.T) {
	fmt.Println(mutexLocked | mutexStarving)
}

func Test2(t *testing.T) {
	//service, _ := androidpublisher.NewService(context.Background())
	//androidpublisher.NewPurchasesSubscriptionsService(service).Get().Do()
	//c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	//c.SetLog(os.Stdout) // Set log to terminal stdout
	//
	//accessToken, err := c.VerifyWebhookSignature()

	buyType := ""
	t2(&buyType)
	fmt.Println(buyType)

}

func t2(t *string) {
	*t = "haha"
}

func getAppleSinglePayLoad(body string) string {
	res := struct {
		SignedPayload string `json:"signedPayload"`
	}{SignedPayload: ""}
	err := json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Fatal(err)
	}
	return res.SignedPayload
}
