package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/proto"
	"time"
)

type jwtMicroService struct {
	key []byte
}

// 用户登录加密
func (s jwtMicroService) Encode(ctx context.Context,req *proto.JwtEncodeRequest,rsp *proto.JwtEncodeResponse) error {
	// 把struct marshal成为json string，再unmarshal到 map[string]interface{}
	j, err := json.Marshal(req.User)
	if err != nil {
		return err
	}
	m := jwt.MapClaims{}
	err = json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)

	str, err := token.SignedString(s.key)
	if err != nil {
		return err
	}

	rsp.Token = str
	return nil
}

func (s jwtMicroService) Decode(ctx context.Context,req *proto.JwtDecodeRequest,rsp *proto.JwtDecodeResponse) error {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return s.key, nil
	})
	if err != nil {
		return err
	}
	err = mapstructure.WeakDecode(token.Claims.(jwt.MapClaims), &rsp.User)
	if err != nil {
		return err
	}
	return nil
}

func main(){
	// etcd registry
	registry := app.EtcdRegistry()

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("micro.service.jwt.test"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*10),
		)

	service.Init()

	serviceHandler := jwtMicroService{key:[]byte("adfasasfwe")}
	err := proto.RegisterJwtHandler(service.Server(),&serviceHandler)
	if err != nil {
		panic(err)
	}
	err = service.Run()
	if err != nil {
		panic(err)
	}
}


