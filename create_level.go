package main

import (
	"gpmember/pkg/dao/mysql"
	"gpmember/pkg/model"
)

func main() {
	db := mysql.NewMemberDB()

	item0 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(0), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:10000 ,UpperLoanPeriod:14, UpperPrincipal:3000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item0)
	im0 := model.MpLevel{ProductName:"Peso2Go", ConfigID:1, Level:0}
	db.CreateMpLevel(&im0)

	item1 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(1), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:4 ,UpperLoanPeriod:14, UpperPrincipal:3500.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item1)
	im1 := model.MpLevel{ProductName:"Peso2Go", ConfigID:2, Level:1}
	db.CreateMpLevel(&im1)

	item2 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(2), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:4000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item2)
	im2 := model.MpLevel{ProductName:"Peso2Go", ConfigID:3, Level:2}
	db.CreateMpLevel(&im2)

	item3 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(3), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:4500.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item3)
	im3 := model.MpLevel{ProductName:"Peso2Go", ConfigID:4, Level:3}
	db.CreateMpLevel(&im3)

	item4 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(4), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:5000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item4)
	im4 := model.MpLevel{ProductName:"Peso2Go", ConfigID:5, Level:4}
	db.CreateMpLevel(&im4)

	item5 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(5), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:6000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item5)
	im5 := model.MpLevel{ProductName:"Peso2Go", ConfigID:6, Level:5}
	db.CreateMpLevel(&im5)

	item6 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(6), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:3, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:7000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item6)
	im6 := model.MpLevel{ProductName:"Peso2Go", ConfigID:7, Level:6}
	db.CreateMpLevel(&im6)

	item7 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(7), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:-1, DeOverdueDays:7 ,UpperLoanPeriod:14, UpperPrincipal:8000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item7)
	im7 := model.MpLevel{ProductName:"Peso2Go", ConfigID:8, Level:7}
	db.CreateMpLevel(&im7)

	item8 := model.MpLevelConfig{ProductName: "Peso2Go", Level: uint(8), InUse: true, Version:"Peso2Go_2019_0805_v1", UpOverdueDays:0, DeOverdueDays:0 ,UpperLoanPeriod:14, UpperPrincipal:8000.0, LowerLoanPeriod:7, LowerPrincipal:2000.0,}
	db.CreateMpLevelConfig(&item8)
	im8 := model.MpLevel{ProductName:"Peso2Go", ConfigID:9, Level:8}
	db.CreateMpLevel(&im8)



}