package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	privateKeyPem := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAwba12rXE3NQTBukFF+uk+Z+Pwh55/onxM6mO7YMaRkV388TB\nZkAKIveeoS3luXeYeL7k5ZOlgcKz0jrGBUny6rvohdWSlWpc5LuP1PXxROm5of3I\n5gtJjJASrEqKLK5d6Ru+aCbnn1b+AgA+kmOyUKYsRmnGaDg5TilmoyWB/3w6989u\niRmACbHqGPVKEl83yfFV46ZSbUNGhNmYZck2txQRA06YcP6ecggVR5L2wCrzLpNz\ndoC3bQzytzWK1Fn8jLPe3hTR3gQ//INg7yABwDvihbdZOgaKnI8E80xyPJoHcWyf\nGO86wte/U8dgo8hBRBaKEnGlm95Hsw/ZcxtKNQIDAQABAoIBABRNpy/eP1z56Wif\nAcapDyiOvc2VzimMobhNfEqOpDFbVKA7Lh4edjGGDJ1OJzbSPyvgrjMVz5ITKy/M\nszaYsppBybRFV1DLziK3OfMTOA+GA8vjwqvB4RqXey2Nvn/CYtts6f8WnM5JmuPw\nzJ4hTu4/DILw0TfZNMBpfHV7F+4EE2MsCsi96LGeJmVsE9A53Gp7RaOkYCp6R864\n2hmZrAZzNaqU4aD1rzd5S3wqw4pmSc2B57a8/H38AF1U9YW+iR1pXtq4OItDDXdg\n26ek3wbkuGC2UPJaoyos0l7MaBBl0YFFP13Q7N3d+Zcr+gEaq0z2wCYMAtjDsWI2\nZufPSZ0CgYEA99ppe25P1g1tQ6vuHiR73lGB4aawNi+F+qY5CObLvJ76yX+L4IAN\n6Tk5CVCfDlh+bDBajMbluR6eg5aeVIW6hPC3WyILcCTY9flCjbqWWS4liQ8JJLMS\nEm4A7UpiLd5KQRxsXS8LDoDURlrY+yfqRxP9jJBNkTzl3/egPTXsa1sCgYEAyBS8\nYLhyU6aq10OayO6Ipn4TGOUxTOt4aa9FaE9B149ZtoJce2wjuPFh7MpoPCRfaRLh\nT5x13W+xsQGbsJjvVva30W/Rp6i2Fq1pL50FJH+cgeJcmVoPux+vbYEVPkN7fqkg\nSPSAhxcC2/P9mHb1s9aAVSweTo5bMxMjPfDZZa8CgYEAhVCq2iR8tuMj+XlaLEZt\nhiiLVwek0pB/XVHZbctOnRdaR9XeNBRM5zzLTBJca4f4AFOF8SDu4cLxelAiu83u\nhKFBzrgiNODs/mljff517mQe9njq7x2Ow/D9eKVA5/EgOaODOiAar2NmSq2E9psC\nrda308qunkeGUhDM1P/TOe8CgYAGmbiFMFCFNfBY3aATlNrpMyuKHLV9ph74zZFq\nmYLAi7gX70EByVV8Wmoyl5LMuR50puzL5Yt13KNuBXGPZ9wtcEIsJJY0A7rOELZx\nnap3w8Xz+vW3EWOHdsogwKtkvHEsgoPQJFDBJB8yBmCNUQ9V+XOOW8A8MzILA0yc\nVH+3fQKBgQDZL/xXUTSNIPtaLBDG/Ezg3MeWr5A8gddx03zrpyMlxoNgF3ohzeik\nH58x9xpmZZcN3Z+YMYUBoEgwBVkgDG98Hk29xnb8RR/9NxE7wN21r83g0Bf1PUv5\n1edwiVh+fEtjwV6AbL399FDCGi8LpWRIexNudj4mdiZFH+EVbo9tsw==\n-----END RSA PRIVATE KEY-----\n"

	viper.Set("PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM", privateKeyPem)

	lp := map[string]string{
		"address":        "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f",
		"signature":      "0xb765e5e88f4b2b3efd7eb8bf08bd83f6521a4282b0fb54401eeb08301603308575a2da7cdf133ce329a475cf79d6cd324038eeaea8d9877ca798619315443c9d1b",
		"origin_message": "Hello World!",
	}
	data, err := json.Marshal(lp)
	if err != nil {
		t.Failed()
	}
	request := events.APIGatewayProxyRequest{HTTPMethod: "POST", Body: string(data)}
	response, err := handler(request)
	if err != nil {
		t.Failed()
	}
	assert.Equal(t, response.Headers["Content-Type"], "text/json")
	assert.Equal(t, response.StatusCode, 200)
	var body map[string]interface{}

	err = json.Unmarshal([]byte(response.Body), &body)
	if err != nil {
		t.Failed()
	}
	assert.NotEmpty(t, body["data"])
}

func Test_process(t *testing.T) {
	privateKeyPem := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAwba12rXE3NQTBukFF+uk+Z+Pwh55/onxM6mO7YMaRkV388TB\nZkAKIveeoS3luXeYeL7k5ZOlgcKz0jrGBUny6rvohdWSlWpc5LuP1PXxROm5of3I\n5gtJjJASrEqKLK5d6Ru+aCbnn1b+AgA+kmOyUKYsRmnGaDg5TilmoyWB/3w6989u\niRmACbHqGPVKEl83yfFV46ZSbUNGhNmYZck2txQRA06YcP6ecggVR5L2wCrzLpNz\ndoC3bQzytzWK1Fn8jLPe3hTR3gQ//INg7yABwDvihbdZOgaKnI8E80xyPJoHcWyf\nGO86wte/U8dgo8hBRBaKEnGlm95Hsw/ZcxtKNQIDAQABAoIBABRNpy/eP1z56Wif\nAcapDyiOvc2VzimMobhNfEqOpDFbVKA7Lh4edjGGDJ1OJzbSPyvgrjMVz5ITKy/M\nszaYsppBybRFV1DLziK3OfMTOA+GA8vjwqvB4RqXey2Nvn/CYtts6f8WnM5JmuPw\nzJ4hTu4/DILw0TfZNMBpfHV7F+4EE2MsCsi96LGeJmVsE9A53Gp7RaOkYCp6R864\n2hmZrAZzNaqU4aD1rzd5S3wqw4pmSc2B57a8/H38AF1U9YW+iR1pXtq4OItDDXdg\n26ek3wbkuGC2UPJaoyos0l7MaBBl0YFFP13Q7N3d+Zcr+gEaq0z2wCYMAtjDsWI2\nZufPSZ0CgYEA99ppe25P1g1tQ6vuHiR73lGB4aawNi+F+qY5CObLvJ76yX+L4IAN\n6Tk5CVCfDlh+bDBajMbluR6eg5aeVIW6hPC3WyILcCTY9flCjbqWWS4liQ8JJLMS\nEm4A7UpiLd5KQRxsXS8LDoDURlrY+yfqRxP9jJBNkTzl3/egPTXsa1sCgYEAyBS8\nYLhyU6aq10OayO6Ipn4TGOUxTOt4aa9FaE9B149ZtoJce2wjuPFh7MpoPCRfaRLh\nT5x13W+xsQGbsJjvVva30W/Rp6i2Fq1pL50FJH+cgeJcmVoPux+vbYEVPkN7fqkg\nSPSAhxcC2/P9mHb1s9aAVSweTo5bMxMjPfDZZa8CgYEAhVCq2iR8tuMj+XlaLEZt\nhiiLVwek0pB/XVHZbctOnRdaR9XeNBRM5zzLTBJca4f4AFOF8SDu4cLxelAiu83u\nhKFBzrgiNODs/mljff517mQe9njq7x2Ow/D9eKVA5/EgOaODOiAar2NmSq2E9psC\nrda308qunkeGUhDM1P/TOe8CgYAGmbiFMFCFNfBY3aATlNrpMyuKHLV9ph74zZFq\nmYLAi7gX70EByVV8Wmoyl5LMuR50puzL5Yt13KNuBXGPZ9wtcEIsJJY0A7rOELZx\nnap3w8Xz+vW3EWOHdsogwKtkvHEsgoPQJFDBJB8yBmCNUQ9V+XOOW8A8MzILA0yc\nVH+3fQKBgQDZL/xXUTSNIPtaLBDG/Ezg3MeWr5A8gddx03zrpyMlxoNgF3ohzeik\nH58x9xpmZZcN3Z+YMYUBoEgwBVkgDG98Hk29xnb8RR/9NxE7wN21r83g0Bf1PUv5\n1edwiVh+fEtjwV6AbL399FDCGi8LpWRIexNudj4mdiZFH+EVbo9tsw==\n-----END RSA PRIVATE KEY-----\n"

	viper.Set("PEOPLELAND_JWT_RSA_PRIVATE_KEY_PEM", privateKeyPem)

	p1 := LoginPayload{Address: "0x40fcc42c5a25945c02b19204d082a67591d30cf6",
		Signature:     "0xff0a604b4400dbc23d2a8ed7a728c552246cd59bcd6a795a7e212622142e9b814f1da8e8af26e03205131b323cb1076486755abb1fbed5f852879257cb4e60c01b",
		OriginMessage: "Hello World!"}

	errorP2 := LoginPayload{Address: "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f",
		Signature:     "0xff0a604b4400dbc23d2a8ed7a728c552246cd59bcd6a795a7e212622142e9b814f1da8e8af26e03205131b323cb1076486755abb1fbed5f852879257cb4e60c01b",
		OriginMessage: "Hello World!"}

	data1, err1 := process(&p1)
	assert.Empty(t, err1)
	assert.NotEmpty(t, data1)

	data2, err2 := process(&errorP2)
	assert.Empty(t, data2)
	assert.NotEmpty(t, err2)
}
