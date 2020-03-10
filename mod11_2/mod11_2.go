package mod11_2

import "errors"

func CalculateWeight(digit uint64) (ans []uint64) {
	sum := 1
	ans = make([]uint64, digit-1)
	for i := digit - 2; i >= 0 && i < digit-1; i-- {
		sum *= 2
		ans[i] = uint64(sum % 11)
	}
	return
}

func CalculateVerificationCode(src, weight []uint64) (ans uint64, err error) {
	ans = 0
	if len(src) != len(weight) {
		err = errors.New("待校验数据长度与权值数据长度不符")
		return
	}
	for i := range src {
		ans += src[i] * weight[i]
	}
	ans = (12 - (ans % 11)) % 11
	return
}
