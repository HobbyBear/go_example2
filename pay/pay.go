package main

import (
	"context"
	"fmt"
	"github.com/plutov/paypal/v4"
	"log"
)

func main() {
	// todo 地址修改
	clientId := "AZQ-hoHICLDgHcfKgeJQdxtjGYDG91eFuotEBgWoLtODInvcFA29kMdD6p18TWC30D-tQNhIST-1XBkR"
	sercreteId := "ECw7YNdOBPsSEZgoEFzBiM7IcbrihAE3LnYElGFF33kAqUA-21XIWdlag5KpWnbzkA-Y0FEr2s-sUCbV"
	c, err := paypal.NewClient(clientId, sercreteId, paypal.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.GetAccessToken(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	webhook, err := c.GetWebhook(context.TODO(), "8KJ93346K61163620")
	fmt.Println(webhook)

	plan, err := c.GetSubscriptionDetails(context.TODO(), "I-XCAPUB5863CR")

	//plan.BillingCycles[0].PricingScheme.FixedPrice.Value = "2500"
	//err = c.UpdateSubscriptionPlanPricing(context.TODO(), "P-3VE1235668573242PMJ53X3A", []paypal.PricingSchemeUpdate{
	//	{
	//		BillingCycleSequence: plan.BillingCycles[0].Sequence,
	//		PricingScheme:        plan.BillingCycles[0].PricingScheme,
	//	},
	//})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(plan)
	//plan.BillingCycles[0].Frequency.IntervalUnit = paypal.IntervalUnitDay
	//plan.BillingCycles[0].Frequency.IntervalCount = 1
	//
	//resp, err := c.CreateSubscriptionPlan(context.TODO(), *plan)
	//fmt.Println("每天扣费一次", resp.ID)
	//
	//plan.BillingCycles[0].Frequency.IntervalUnit = paypal.IntervalUnitWeek
	//plan.BillingCycles[0].Frequency.IntervalCount = 1
	//
	//resp, err = c.CreateSubscriptionPlan(context.TODO(), *plan)
	//fmt.Println("每周扣费一次", resp.ID)
	//
	//plan.BillingCycles[0].Frequency.IntervalUnit = paypal.IntervalUnitMonth
	//plan.BillingCycles[0].Frequency.IntervalCount = 1
	//
	//resp, err = c.CreateSubscriptionPlan(context.TODO(), *plan)
	//fmt.Println("每月扣费一次", resp.ID)
	//
	//plan.BillingCycles[0].Frequency.IntervalUnit = paypal.IntervalUnitYear
	//plan.BillingCycles[0].Frequency.IntervalCount = 1
	//
	//resp, err = c.CreateSubscriptionPlan(context.TODO(), *plan)
	//fmt.Println("每年扣费一次", resp.ID)

	//p, err := c.GetSubscriptionDetails(context.TODO(), "I-9XPB4TVKRKW6")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//p.PlanID = "P-49R7951905229025GMJ47KXQ"
	//err = c.UpdateSubscription(context.TODO(), paypal.Subscription{
	//	SubscriptionDetailResp: *p,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//data, _ := json.Marshal(p)
	//
	//fmt.Println(time.Time(*p.StartTime).Unix())
	//fmt.Println(time.Time(*p.StartTime).AddDate(0, 3, 0))
	//fmt.Println(time.Time(*p.StartTime))
	//log.Println(string(data))

	//response, err := c.GetSubscriptionPlan(context.TODO(), "P-7PG29513FD0578609MJ4PAKI")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(response)
	//c.GetAccessToken(context.Background())
	//resp, err := c.CreateSubscription(context.Background(), paypal.SubscriptionBase{PlanID: ""})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)

	//token := uuid.New()
	//returnUrl := fmt.Sprintf("%s?res={status:%d,data:{token:%s}}", "https://www.baidu.com", 1, token)
	//cancelUrl := fmt.Sprintf("%s?res={status:%d,data:{token:%s}}", "https://www.baidu.com", 0, token)
	//
	//ru, _ := url.Parse(returnUrl)
	//cu, _ := url.Parse(cancelUrl)
	//subBase := paypal.SubscriptionBase{
	//	PlanID: "P-6BY76772Y64953117MJ5IR2Q",
	//}
	//q := ru.Query()
	//ru.RawQuery = q.Encode()
	//fmt.Println(ru.String())
	//subBase.ApplicationContext = &paypal.ApplicationContext{
	//	ReturnURL: ru.String(),
	//	CancelURL: cu.String(),
	//}
	//resp, err := c.CreateSubscription(context.Background(), subBase)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(resp)
}

func createTestP(plan *paypal.Product) {
	// todo 地址修改
	clientId := "AZQ-hoHICLDgHcfKgeJQdxtjGYDG91eFuotEBgWoLtODInvcFA29kMdD6p18TWC30D-tQNhIST-1XBkR"
	sercreteId := "ECw7YNdOBPsSEZgoEFzBiM7IcbrihAE3LnYElGFF33kAqUA-21XIWdlag5KpWnbzkA-Y0FEr2s-sUCbV"
	c, err := paypal.NewClient(clientId, sercreteId, paypal.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.GetAccessToken(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	response, err := c.CreateProduct(context.TODO(), *plan)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}

func createTestPlan(plan *paypal.SubscriptionPlan) {
	// todo 地址修改
	clientId := "AZQ-hoHICLDgHcfKgeJQdxtjGYDG91eFuotEBgWoLtODInvcFA29kMdD6p18TWC30D-tQNhIST-1XBkR"
	sercreteId := "ECw7YNdOBPsSEZgoEFzBiM7IcbrihAE3LnYElGFF33kAqUA-21XIWdlag5KpWnbzkA-Y0FEr2s-sUCbV"
	c, err := paypal.NewClient(clientId, sercreteId, paypal.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.GetAccessToken(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	response, err := c.CreateSubscriptionPlan(context.TODO(), *plan)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
