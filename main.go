package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type ID struct {
	Area             string
	Year             uint64
	Month            uint64
	Day              uint64
	SequenceCode     uint64
	Gender           uint64
	EachNum          [17]uint64
	VerificationCode uint64
	Correct          bool
}

var srcID = flag.String("id", "", "输入18位身份证号码")

func string2uint64(str string) (ret uint64) {
	if str == "X" || str == "x" {
		ret = 10
		return
	}
	ret, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func getNumber() (personalID ID, err error) {
	flag.Parse()
	var idNum string
	if *srcID != "" {
		idNum = *srcID
	} else {
		fmt.Println("请输入18位身份证号码")
		fmt.Scanf("%v", &idNum)
	}
	if len(idNum) != 18 {
		err = errors.New("身份证号码位数不正确，请重试")
		return
	}
	personalID.Year = string2uint64(idNum[6:10])
	personalID.Month = string2uint64(idNum[10:12])
	personalID.Day = string2uint64(idNum[12:14])
	personalID.SequenceCode = string2uint64(idNum[14:17])
	personalID.Gender = string2uint64(idNum[16:17]) % 2
	personalID.VerificationCode = string2uint64(idNum[17:18])
	for i := 0; i < 17; i++ {
		personalID.EachNum[i] = string2uint64(idNum[i : i+1])
	}
	areaFile, err := os.Open("./conf/area_code.json")
	defer areaFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	areaJson := json.NewDecoder(areaFile)
	area := map[uint64]string{}
	err = areaJson.Decode(&area)
	if err != nil {
		fmt.Println(err)
	}
	personalID.Area = area[personalID.EachNum[0]*100000+personalID.EachNum[1]*10000]
	if personalID.Area == "" {
		err = errors.New("不是合法的身份证号码，请重试")
	}
	return
}

func checkNumber(in [17]uint64) (s uint64, err error) {
	year := in[6]*1000 + in[7]*100 + in[8]*10 + in[9]
	mon := in[10]*10 + in[11]
	day := in[12]*10 + in[13]
	breakEnable := false
	switch {
	case day == 0:
		breakEnable = true
	case mon == 1 || mon == 3 || mon == 5 || mon == 7 || mon == 8 || mon == 10 || mon == 12:
		if day > 31 {
			breakEnable = true
		}
	case mon == 4 || mon == 6 || mon == 9 || mon == 11:
		if day > 30 {
			breakEnable = true
		}
	case mon == 2:
		if year%4 == 0 {
			if year%100 == 0 && year%400 != 0 {
				if day > 28 {
					breakEnable = true
				}
			} else {
				if day > 29 {
					breakEnable = true
				}
			}
		} else {
			if day > 28 {
				breakEnable = true
			}
		}
	default:
		breakEnable = true
	}
	if breakEnable == true {
		err = errors.New("日期输入有问题，你确定这是地球日期？")
		return
	}
	weight := [17]uint64{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	s = 0
	for i := 0; i < 17; i++ {
		s += in[i] * weight[i]
	}
	s = (12 - (s % 11)) % 11
	return
}

func output(personalID ID) {
	if personalID.Correct != true {
		fmt.Println("此身份证号码未能通过校验")
		fmt.Println("正确的校验码为", personalID.VerificationCode)
		return
	}
	fmt.Printf("出生日期：%v年%v月%v日 ", personalID.Year, personalID.Month, personalID.Day)
	if personalID.Gender == 1 {
		fmt.Println("性别：男")
	} else {
		fmt.Println("性别：女")
	}
	fmt.Println("所属地区：", personalID.Area)
	fmt.Println("此身份证号码通过了校验")
}

func main() {
	personalID, err := getNumber()
	if err != nil {
		fmt.Println(err)
		return
	}
	ans, err := checkNumber(personalID.EachNum)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ans == personalID.VerificationCode {
		personalID.Correct = true
	} else {
		personalID.VerificationCode = ans
	}
	output(personalID)
}
