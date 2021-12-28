package service

import (
	api "backend/api/user/v1"
	"backend/app/user/internal/biz"
	"backend/app/user/internal/conf"
	dt "backend/app/user/internal/data"
	"backend/app/user/internal/mock_biz"
	"backend/pkg/jwt"
	"context"
	"github.com/golang/mock/gomock"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var privateKeyPem = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAwba12rXE3NQTBukFF+uk+Z+Pwh55/onxM6mO7YMaRkV388TB\nZkAKIveeoS3luXeYeL7k5ZOlgcKz0jrGBUny6rvohdWSlWpc5LuP1PXxROm5of3I\n5gtJjJASrEqKLK5d6Ru+aCbnn1b+AgA+kmOyUKYsRmnGaDg5TilmoyWB/3w6989u\niRmACbHqGPVKEl83yfFV46ZSbUNGhNmYZck2txQRA06YcP6ecggVR5L2wCrzLpNz\ndoC3bQzytzWK1Fn8jLPe3hTR3gQ//INg7yABwDvihbdZOgaKnI8E80xyPJoHcWyf\nGO86wte/U8dgo8hBRBaKEnGlm95Hsw/ZcxtKNQIDAQABAoIBABRNpy/eP1z56Wif\nAcapDyiOvc2VzimMobhNfEqOpDFbVKA7Lh4edjGGDJ1OJzbSPyvgrjMVz5ITKy/M\nszaYsppBybRFV1DLziK3OfMTOA+GA8vjwqvB4RqXey2Nvn/CYtts6f8WnM5JmuPw\nzJ4hTu4/DILw0TfZNMBpfHV7F+4EE2MsCsi96LGeJmVsE9A53Gp7RaOkYCp6R864\n2hmZrAZzNaqU4aD1rzd5S3wqw4pmSc2B57a8/H38AF1U9YW+iR1pXtq4OItDDXdg\n26ek3wbkuGC2UPJaoyos0l7MaBBl0YFFP13Q7N3d+Zcr+gEaq0z2wCYMAtjDsWI2\nZufPSZ0CgYEA99ppe25P1g1tQ6vuHiR73lGB4aawNi+F+qY5CObLvJ76yX+L4IAN\n6Tk5CVCfDlh+bDBajMbluR6eg5aeVIW6hPC3WyILcCTY9flCjbqWWS4liQ8JJLMS\nEm4A7UpiLd5KQRxsXS8LDoDURlrY+yfqRxP9jJBNkTzl3/egPTXsa1sCgYEAyBS8\nYLhyU6aq10OayO6Ipn4TGOUxTOt4aa9FaE9B149ZtoJce2wjuPFh7MpoPCRfaRLh\nT5x13W+xsQGbsJjvVva30W/Rp6i2Fq1pL50FJH+cgeJcmVoPux+vbYEVPkN7fqkg\nSPSAhxcC2/P9mHb1s9aAVSweTo5bMxMjPfDZZa8CgYEAhVCq2iR8tuMj+XlaLEZt\nhiiLVwek0pB/XVHZbctOnRdaR9XeNBRM5zzLTBJca4f4AFOF8SDu4cLxelAiu83u\nhKFBzrgiNODs/mljff517mQe9njq7x2Ow/D9eKVA5/EgOaODOiAar2NmSq2E9psC\nrda308qunkeGUhDM1P/TOe8CgYAGmbiFMFCFNfBY3aATlNrpMyuKHLV9ph74zZFq\nmYLAi7gX70EByVV8Wmoyl5LMuR50puzL5Yt13KNuBXGPZ9wtcEIsJJY0A7rOELZx\nnap3w8Xz+vW3EWOHdsogwKtkvHEsgoPQJFDBJB8yBmCNUQ9V+XOOW8A8MzILA0yc\nVH+3fQKBgQDZL/xXUTSNIPtaLBDG/Ezg3MeWr5A8gddx03zrpyMlxoNgF3ohzeik\nH58x9xpmZZcN3Z+YMYUBoEgwBVkgDG98Hk29xnb8RR/9NxE7wN21r83g0Bf1PUv5\n1edwiVh+fEtjwV6AbL399FDCGi8LpWRIexNudj4mdiZFH+EVbo9tsw==\n-----END RSA PRIVATE KEY-----\n"
var publicKeyPem = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwba12rXE3NQTBukFF+uk\n+Z+Pwh55/onxM6mO7YMaRkV388TBZkAKIveeoS3luXeYeL7k5ZOlgcKz0jrGBUny\n6rvohdWSlWpc5LuP1PXxROm5of3I5gtJjJASrEqKLK5d6Ru+aCbnn1b+AgA+kmOy\nUKYsRmnGaDg5TilmoyWB/3w6989uiRmACbHqGPVKEl83yfFV46ZSbUNGhNmYZck2\ntxQRA06YcP6ecggVR5L2wCrzLpNzdoC3bQzytzWK1Fn8jLPe3hTR3gQ//INg7yAB\nwDvihbdZOgaKnI8E80xyPJoHcWyfGO86wte/U8dgo8hBRBaKEnGlm95Hsw/ZcxtK\nNQIDAQAB\n-----END PUBLIC KEY-----\n"
var config = &conf.Config{
	JwtRsaPrivateKeyPem: privateKeyPem,
	JwtRsaPublicKeyPem:  publicKeyPem,
	FaunaDBSecret:       "fnAEbfitSAACVKRgPF0ZYX-Q3zZiIE3jQpr_9km0",
}

var logger = log.Default()

func TestUserService_Login(t *testing.T) {

	d, _ := dt.NewData(config, logger)
	userRepo := dt.NewUserRepo(d, logger)
	twitterRepo := dt.NewTwitterOauth1Repo(config)
	userUseCase := biz.NewUserUseCase(userRepo, twitterRepo, config, logger)
	us := &UserService{
		uc:     userUseCase,
		logger: logger,
		conf:   config,
	}
	addresOne := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	addresTwo := "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f"
	signatureOne := "0xff0a604b4400dbc23d2a8ed7a728c552246cd59bcd6a795a7e212622142e9b814f1da8e8af26e03205131b323cb1076486755abb1fbed5f852879257cb4e60c01b"
	originMessage := "Hello World!"
	tests := []struct {
		name    string
		args    api.LoginPayLoad
		wantErr string
	}{
		{"1", api.LoginPayLoad{
			Address:       &addresOne,
			Signature:     &signatureOne,
			OriginMessage: &originMessage,
		}, ""},
		{"2", api.LoginPayLoad{
			Address:       &addresTwo,
			Signature:     &signatureOne,
			OriginMessage: &originMessage,
		}, "request.verify.error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.Login(context.TODO(), &tt.args)
			if tt.wantErr == "" {
				assert.NotEmpty(t, got)
				assert.Empty(t, err)
			} else {
				assert.Equalf(t, tt.wantErr, err.Error(), "Login(%v)", tt.args)
			}
		})
	}
}

func TestUserService_GetProfile(t *testing.T) {
	d, _ := dt.NewData(config, logger)
	userRepo := dt.NewUserRepo(d, logger)
	twitterRepo := dt.NewTwitterOauth1Repo(config)
	userUseCase := biz.NewUserUseCase(userRepo, twitterRepo, config, logger)
	us := &UserService{
		uc:     userUseCase,
		logger: logger,
		conf:   config,
	}

	address := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	claims := jwt.NewMapClaims(address)
	exp := int64(86400)
	jwtStr, err := jwt.EncodeJwt(claims, privateKeyPem, exp)
	if err != nil {
		t.Failed()
	}
	assert.NotEmpty(t, jwtStr)
	tests := []struct {
		name string
		args context.Context
	}{
		{"1", context.WithValue(context.Background(), "authorization", jwtStr)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.GetProfile(tt.args)
			assert.Empty(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func TestUserService_PutProfile(t *testing.T) {
	d, _ := dt.NewData(config, logger)
	userRepo := dt.NewUserRepo(d, logger)
	twitterRepo := dt.NewTwitterOauth1Repo(config)
	userUseCase := biz.NewUserUseCase(userRepo, twitterRepo, config, logger)
	us := &UserService{
		uc:     userUseCase,
		logger: logger,
		conf:   config,
	}

	address := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	claims := jwt.NewMapClaims(address)
	exp := int64(86400)
	jwtStr, err := jwt.EncodeJwt(claims, privateKeyPem, exp)
	if err != nil {
		t.Failed()
	}

	ctx := context.WithValue(context.Background(), "authorization", jwtStr)

	tests := []struct {
		name string
		args api.PutProfilePayLoad
	}{
		{"1", api.PutProfilePayLoad{
			Name: "haha",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := us.PutProfile(ctx, &tt.args)
			assert.Equal(t, got.Address, address)
			assert.Equal(t, got.Name, "haha")
			assert.Empty(t, err)
		})
	}
}

func TestUserService_ConnectTwitter(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockTwitterRepo := mock_biz.NewMockTwitterRepo(mockCtl)

	d, _ := dt.NewData(config, logger)
	userRepo := dt.NewUserRepo(d, logger)
	userUseCase := biz.NewUserUseCase(userRepo, mockTwitterRepo, config, logger)
	us := &UserService{
		uc:     userUseCase,
		logger: logger,
		conf:   config,
	}

	address := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	claims := jwt.NewMapClaims(address)
	exp := int64(86400)
	jwtStr, err := jwt.EncodeJwt(claims, privateKeyPem, exp)
	if err != nil {
		t.Failed()
	}

	ctx := context.WithValue(context.Background(), "authorization", jwtStr)

	tests := []struct {
		name    string
		twitter string
		tweets  []string
		args    api.ConnectTwitterPayLoad
		success bool
	}{
		{"1", "a", []string{"I am verifying my identity as 1 on peopleland"}, api.ConnectTwitterPayLoad{
			Twitter: "a",
		}, true},
		{"2", "b", []string{"I am verifying my identity as 1 on peopleland"}, api.ConnectTwitterPayLoad{
			Twitter: "b",
		}, false},
		{"3", "c", []string{"aaaaI am verifying my identity as 3 on peoplelandbbbb"}, api.ConnectTwitterPayLoad{
			Twitter: "c",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = userRepo.UpdateUserByAddress(ctx, address, map[string]interface{}{"name": tt.name})
			mockTwitterRepo.EXPECT().GetTwitterUserTimeline(tt.twitter).Return(tt.tweets)

			got, err := us.ConnectTwitter(ctx, &tt.args)
			if tt.success {
				assert.Equal(t, got.Address, address)
				assert.Equal(t, got.Twitter, tt.twitter)
				assert.Empty(t, err)
			} else {
				assert.NotEmpty(t, err)
			}

		})
	}
}
