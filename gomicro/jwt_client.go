package main

import (
	"context"
	"fmt"
	"github.com/app"
	"github.com/micro/go-micro/v2"
	"github.com/proto"
)

func main(){
	service := micro.NewService(
		micro.Registry(app.EtcdRegistry()),
		)
	service.Init()

	jwt := proto.NewJwtService("micro.service.jwt.test",service.Client())
	rsp,err := jwt.Encode(context.TODO(),&proto.JwtEncodeRequest{
		User:                 &proto.JwtUser{
			Type:                 "user",
			Uuid:                 "",
			Id:                   11,
			Username:             "name",
			Nickname:             "nick",
			LoginTime:            0,
			//XXX_NoUnkeyedLiteral: struct{}{},
			//XXX_unrecognized:     nil,
			//XXX_sizecache:        0,
		},
		//XXX_NoUnkeyedLiteral: struct{}{},
		//XXX_unrecognized:     nil,
		//XXX_sizecache:        0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Token)
}
