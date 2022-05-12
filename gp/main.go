package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/androidpublisher/v3"
	"net/http"
	"time"
)

const (
	GpJsonKey = "{\n  \"type\": \"service_account\",\n  \"project_id\": \"pc-api-8448201651253025094-228\",\n  \"private_key_id\": \"4782739b72f7b1ce4b6770bc9a3e571bf1052c0c\",\n  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDGx6ZY2DMfznCt\\n4M2VCkVvSMVsiDENBKGByvoWAO86taRc6ExUWdxDMqsnOVFILIG/o4AGbeveWBLJ\\ne6mcJUIQuqZjzaGkSEXtqWSUNt2FWr5d7YanHx5qMnZF0vImIWlkkfiAzt3KYLSI\\nklCkoPoExfOtBjL1e5F5y6Yk6ZYIhiFqSlUoz/0nLYW81Ao0MFdB0I5z+yJMgV1O\\n3jOFKeWmrwadYn1rMC+7C2TAFMDXxospXRN8qsyxfbqLvbLj4n9XKeEaRQUIVApQ\\nZdLK1pk2CjR5ozfi3loX2gp668Q+RON2nZ+vw9DOg2jQwSVkaOZ86jd4qLY0aYFc\\nOPIUJZZ5AgMBAAECggEAHXiXQJ/ZQJN2YEkuobSyo7Op1ZQBlbkEiv5BgsajsVQh\\n6MZmHrFt9yaJepjEZMt4dyWyAwIEEOSE8Itdk7FwcJODvNNNblJqz7goGdNil3Vj\\nOQDQ0HFMLJmLsIF8dkvI2OjNa6d/+ZPJ4F+1Bht5xgBlRtde8cdew6x+VwWswVmC\\nScHUfyhFM3ov7Iudb55aqceRjMj5WVZQWKAzcQ49AA/VcSD/tbd9MWPNvx9l5huT\\nBBZWxH1nd7iszVp/m4qZ80sSBuc8jitsNF7IeNZQbnl+1Jm42nPhaGHxf+9yY0Dp\\n/NVPm064p4JwdtAhTYJUliY7ZIZ/LOVPgd5wiE28AwKBgQD72XT0BSd2B9woTXxP\\nFL1mHNv2jUFWroeJ+UYtNhj/L0juhsHOaWglynzCJzXklQme1C3II1mMgpTrhtKf\\nzoM+sgyr0aHOp7b1Odzj+XdFTZTsiHD3gclkyXFg48zg8RGB2cc/suk31jwBwIxi\\njYDT8E24S+HcJgRPBrtjafos2wKBgQDKDktk1ASFmdIjYvIIQrGMTEWq2TxJyMxs\\nx857Y0u0hJOfGq/I37xlVyOY3ShyYqhyxuOGPSKDKsGJtn0afGVxu2qN5wCFa7Fs\\n7yWEUgRU4BGMwgHusvmj38gvtmwWzfah0Ej8Vymc7dEURks7Yt8BsBDevdvsjfsP\\nf6I/fFvAOwKBgQCdIHbPJ8gO46h/nJhJ3m6sHqU9lbOGGS2G73bRDkdaLm4aIwkX\\ndrv7l/FuUHhXPITcCxTlIOjluTKgH2WZPoazvvIY4JlLDWggZxFhSmsXWefoCtgw\\nE1FLCEMsZ/OTA9QgZ3bgTwciRw75KOBtPEt77xkOOomXiEB37YNL6vqisQKBgAp/\\ndaFWdqmuVUGraWHOfUUJzXDT+pLm0OAx/WNhxFQrszPIAiwQm8w0G9BcZTaUM6ZK\\n+RVHZdykwwtxpMxUfDzNfPrEbr1/2+cKAcv/KuRyA9dpWQpa2CeQf+kitDj3GIJN\\nvQ0HRk01+e/wuwXp3CZNTIMWIDAJLGsguYBM6MNzAoGARFHS8f1PCoBx4MSHkKza\\nzBnrywQ9N//kvLFq30ntXOL9XE0uNbFZL4tSmeZJzLWdxNeAKZhKT3recasD4w68\\nwlBQfeoEQEu9a1KLoWcUbedKVUyE2XAlU5Ns20+eiaZadv1dPG0NOrDwI3+lEv03\\nVnNf5JwDqR3o83rLn1cqYlo=\\n-----END PRIVATE KEY-----\\n\",\n  \"client_email\": \"in-app-purchase-verification@pc-api-8448201651253025094-228.iam.gserviceaccount.com\",\n  \"client_id\": \"107349581794938168427\",\n  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n  \"token_uri\": \"https://oauth2.googleapis.com/token\",\n  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/in-app-purchase-verification%40pc-api-8448201651253025094-228.iam.gserviceaccount.com\"\n}"
)

func main() {
	orderId := "jkndabkplldapefilcobgbpb.AO-J1OzvbVJMjUI4V3hPfG7S7-cwKFMA38ykfqr3HCQPkIL17ajj_x_FMUHrkDcRiJwsKvhU6V4VfQryk8HaqQRZujoFXOUKsw"

	service, _ := New([]byte(GpJsonKey))

	resp, err := service.Purchases.Subscriptions.Get("com.linkbox.app", "linkbox_premium_monthly_test_35", orderId).Do()
	fmt.Println(resp, err)

}

func New(jsonKey []byte) (*androidpublisher.Service, error) {
	c := &http.Client{Timeout: 10 * time.Second}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, c)

	conf, err := google.JWTConfigFromJSON(jsonKey, androidpublisher.AndroidpublisherScope)
	if err != nil {
		return nil, err
	}

	val := conf.Client(ctx).Transport.(*oauth2.Transport)
	_, err = val.Source.Token()
	if err != nil {
		return nil, err
	}

	service, err := androidpublisher.New(conf.Client(ctx))
	if err != nil {
		return nil, err
	}
	return service, nil
}
