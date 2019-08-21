package main

import (
	"context"
	"fmt"
	"giapproval/model"
	"github.com/olivere/elastic"
	"github.com/rs/xid"
	"giutils/logger"
	"giutils/util"
	"math/rand"
	"time"
)

func main() {
	c, _ := elastic.NewSimpleClient(elastic.SetURL("http://172.16.5.122:9200"))

	esIndex := "approval_dev"
	esType := "duplicate_feature_data"

	startAt := time.Now()
	logger.InitLoggerBak(logger.LvlDebug, nil)
	on := fmt.Sprintf("dasczcz%v", startAt.UnixNano())
	currFd := model.DuplicateFeatureData{
		OrderNo:               on,
		ApplyTime:             &startAt,
		ApplyPhone:            "18y78y72131",
		FullName:              "acbj cnzjk",
		ImmediateContactPhone: "1231741841",
		OtherContactPhone:     "1874812471",
		IdType:                "PanCard",
		IdCardNum:             "1548781231afas",
		AddrCardType:          "Addr",
		AddrCardNum:           "123123213151",
		DeviceMac:             "17cbxzbuyh124",
	}
	currFd.AssignOrderHash()
	// case 2: success - all rule
	// .Refresh("wait_for")
	ir, err := c.Index().Index(esIndex).Type(esType).BodyJson(currFd).Do(context.Background())
	fmt.Println(ir, "es-err=", err)

	for i := 0; i < 2000; i++ {
		bulkRequest := c.Bulk()

		logger.Debug("doing......", "i", i)
		m := model.DuplicateFeatureData{
			OrderNo:               "111_" + util.GetMsgId(),
			ApplyTime:             &startAt,
			ApplyPhone:            currFd.ImmediateContactPhone,
			ImmediateContactPhone: currFd.ImmediateContactPhone,
			IdType:                "cznxjcbhja",
			IdCardNum:             currFd.IdCardNum,
			AddrCardType:          currFd.AddrCardType,
			AddrCardNum:           currFd.AddrCardNum,
			DeviceMac:             currFd.DeviceMac,
			FullName:              currFd.FullName,
		}
		m.AssignOrderHash()
		bulkRequest.Add(elastic.NewBulkIndexRequest().Index(esIndex).Type(esType).Doc(m))

		forStart := time.Now()
		rand.Seed(forStart.UnixNano())
		at := rand.Int63n(5)
		for z := 0; z < 5000; z++ {
			nd := &model.DuplicateFeatureData{
				OrderNo:               "111_" + util.GetMsgId(),
				ApplyTime:             &startAt,
				ApplyPhone:            xid.New().String(),
				ImmediateContactPhone: xid.New().String(),
				IdType:                "cznxjcbhja",
				IdCardNum:             xid.New().String(),
				AddrCardType:          fmt.Sprintf("%v", at),
				AddrCardNum:           xid.New().String(),
				DeviceMac:             xid.New().String(),
				FullName:              xid.New().String(),
			}
			bulkRequest.Add(elastic.NewBulkIndexRequest().Index(esIndex).Type(esType).Doc(nd))
		}
		fmt.Println("for use", time.Since(forStart))
		insertStart := time.Now()
		resp, err := bulkRequest.Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Println("insert use", time.Since(insertStart), "count", len(resp.Indexed()))
	}

	fmt.Println("finish all insert use time:", time.Since(startAt))

	//time.Sleep(1 * time.Hour)
}