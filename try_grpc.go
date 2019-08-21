package main

import (
	"context"
	rawJson "encoding/json"
	"errors"
	"fmt"
	"gpcore/model"
	"gpmember/grpc/client"
	"gpmember/grpc/pb"
	memModel "gpmember/pkg/model"
	"gputils/encoding/json"
)

func NoticePriceConfigUpdate(pcs []model.PriceConfig) error {
	fmt.Println(json.StringifyJson(pcs))

	//var PriceRecord memModel.GetPriceConfigsResp
	var configs []memModel.PriceConfig
	if err := rawJson.Unmarshal([]byte(json.StringifyJson(pcs)), &configs); err != nil {
		fmt.Println(err)
	}

	resp, err := client.GetMenberRpcClient().NoticePriceConfigUpdate(context.Background(), &pb.CoreNotice{
		PriceConfig: json.StringifyJson(pcs),
	})
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New(resp.ErrMsg)
	}
	return nil
}

func GetUserRank() {
	req := &pb.PriceReq{BaseUserId:12751, AppName:"Peso2Go"}
	resp, err := client.GetMenberRpcClient().GetUserRank(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func main() {
	var Product = model.ProductInfo{
		ProductName: "qwe",
	}

	var Price = model.PriceInfo{
		LoanPrincipal: 5000,
		LoanPeriod:    20,
		LoanDays: 23,
	}
	var config = model.PriceConfig{
		ProductInfo: Product,
		PriceInfo:   Price,
	}

	var Produc2 = model.ProductInfo{
		ProductName: "=poi",
	}

	var Price2 = model.PriceInfo{
		LoanPrincipal: 5000,
		LoanPeriod:    20,
		LoanDays: 14,
	}
	var config2 = model.PriceConfig{
		ProductInfo: Produc2,
		PriceInfo:   Price2,
	}
	var array []model.PriceConfig
	array = append(append(array, config), config2)

	//err := NoticePriceConfigUpdate(array)
	//if err != nil {
	//	fmt.Println(err)
	//}
	GetUserRank()
}
