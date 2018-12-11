package common

import (
	"math/big"
	"testing"
)

func Test_DeepCopy(t *testing.T) {
	mapA := map[string]interface{}{
		"A": 123,
		"d": 44,
		"B": map[string]interface{}{
			"F": "X123",
		},
	}

	var mapB interface{}
	t.Logf("mapSrc %#v", mapA)
	mapB, err := DeepCopy(mapA)
	if err != nil {
		t.Error(err)
	}
	mapA["A"] = 34
	mapA["d"] = 66
	mapBB, ok := mapB.(map[string]interface{})
	if !ok {
		t.Error("infer error")
	}
	mapBBA, ok := mapBB["A"].(int)
	if !ok {
		t.Error("infer error")
	}
	if mapA["A"] == mapBBA {
		t.Errorf("mapA[A] %v shouldn`t be equal mapB[A] %v", mapA["A"], mapBBA)
	}
	mapBBd, _ := mapBB["d"].(int)
	if mapA["d"] == mapBBd {
		t.Errorf("mapA[d] %v shouldn`t be equal mapB[d] %v", mapA["d"], mapBBd)
	}
}

func Test_RemainNDigitsForBigInt(t *testing.T) {
	a := big.NewInt(10)
	a.Mul(a, TokenUnit)
	b := BigIntDivToRat(a, TokenUnit)
	if b.Cmp(big.NewRat(10, 1)) != 0 {
		t.Error("BigIntDivToRat error")
	}

	c := BigIntDivStrWithPrec(a, TokenUnit, 3)
	if c != "10.000" {
		t.Error("BigIntDivStrWithPrec error")
	}

	f := BigIntDiv18BitToRat(a)
	if f.Cmp(big.NewRat(10, 1)) != 0 {
		t.Error("BigIntDivToRat error")
	}
}

func Test_BigIntAdd(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("100000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}

	z := BigIntAdd(x, y)
	exp, ok := big.NewRat(1, 1).SetString("200000000000000000")
	if !ok {
		t.Error("new Rat error")
	}
	if exp.Cmp(big.NewRat(1, 1).SetInt(z)) != 0 {
		t.Errorf("x %v, y %v, add= %v", x, y, z)
	}
}

func Test_BigIntSub(t *testing.T) {
	x, ok := big.NewInt(0).SetString("200000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("100000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	exp, ok := big.NewRat(1, 1).SetString("100000000000000000")
	if !ok {
		t.Error("new Rat error")
	}
	z := BigIntSub(x, y)
	if exp.Cmp(big.NewRat(1, 1).SetInt(z)) != 0 {
		t.Errorf("x %v, y %v, sub= %v", x, y, z)
	}
}

func Test_BigIntMul(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("2", 10)
	if !ok {
		t.Error("new Int error")
	}
	exp, ok := big.NewRat(1, 1).SetString("200000000000000000")
	if !ok {
		t.Error("new Rat error")
	}
	z := BigIntMul(x, y)
	if exp.Cmp(big.NewRat(1, 1).SetInt(z)) != 0 {
		t.Errorf("x %v, y %v, mul= %v", x, y, z)
	}
}

func Test_BigIntDiv(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("10000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	exp, ok := big.NewRat(1, 1).SetString("10")
	if !ok {
		t.Error("new Rat error")
	}
	z := BigIntDiv(x, y)
	if exp.Cmp(big.NewRat(1, 1).SetInt(z)) != 0 {
		t.Errorf("x %v, y %v, div= %v", x, y, z)
	}
}

func Test_BigIntDivToRat(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("10000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	exp, ok := big.NewRat(1, 1).SetString("10")
	if !ok {
		t.Error("new Rat error")
	}
	z := BigIntDivToRat(x, y)
	if exp.Cmp(z) != 0 {
		t.Errorf("x %v, y %v, div= %v", x, y, z)
	}
}

func Test_BigIntDivToFloat32(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("10000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	exp, ok := big.NewFloat(1.0).SetString("10")
	if !ok {
		t.Error("new Float error")
	}

	z, _ := BigIntDivToFloat32(x, y)
	if exp.Cmp(big.NewFloat(float64(z))) != 0 {
		t.Errorf("x %v, y %v, div= %v", x, y, z)
	}
}

func Test_BigIntDivToFloat64(t *testing.T) {
	x, ok := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	y, ok := big.NewInt(0).SetString("10000000000000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}
	z, _ := BigIntDivToFloat64(x, y)
	exp, ok := big.NewFloat(1.0).SetString("10")
	if !ok {
		t.Error("new Float error")
	}
	if exp.Cmp(big.NewFloat(z)) != 0 {
		t.Errorf("x %v, y %v, div= %v", x, y, z)
	}
}

func Test_BigIntToFloatWithTokenUnit(t *testing.T) {
	x, _ := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	unit := big.NewInt(1e18)
	y := BigIntToFloat32WithTokenUnit(x, unit)
	t.Logf("x %v, unit %v, convert to float32 with div token unit %v", x, unit, y)
}

func Test_BigIntToFloat64WithTokenUnit(t *testing.T) {
	x, _ := big.NewInt(0).SetString("100000000000000000000000000000", 10)
	a := big.NewInt(1e18)
	b := big.NewInt(1e10)
	b.Mul(b, a)
	unit := b
	y := BigIntToFloat64WithTokenUnit(x, unit)
	exp, ok := big.NewFloat(1.0).SetString("10")
	if !ok {
		t.Error("new Float error")
	}
	if exp.Cmp(big.NewFloat(y)) != 0 {
		t.Errorf("x %v, unit %v, convert to float64 with div token unit %v", x, unit, y)
	}
}

func Test_BigIntToFloat32WithDefaultTokenUnit(t *testing.T) {
	x, ok := big.NewInt(0).SetString("10000000000000000000", 10)
	if !ok {
		t.Error("new Int error")
	}

	y := BigIntToFloat32WithDefaultTokenUnit(x)
	exp, ok := big.NewFloat(1.0).SetString("10")
	if !ok {
		t.Error("new Float error")
	}
	if exp.Cmp(big.NewFloat(float64(y))) != 0 {
		t.Errorf("x %v, convert to float32 with div token unit %v", x, y)
	}
}

func Test_BigIntToFloat64WithDefaultTokenUnit(t *testing.T) {
	x, _ := big.NewInt(0).SetString("10000000000000000000", 10)

	y := BigIntToFloat64WithDefaultTokenUnit(x)
	exp, ok := big.NewFloat(1.0).SetString("10")
	if !ok {
		t.Error("new Float error")
	}
	if exp.Cmp(big.NewFloat(y)) != 0 {
		t.Logf("x %v, convert to float64 with div token unit %v", x, y)
	}
}

func Test_CheckEmail(t *testing.T) {
	email := "1234567@qq.com"
	ok := CheckEmail(email)
	if ok {
		t.Logf("%v is valid email", email)
	} else {
		t.Errorf("%v is not valid email", email)
	}
	emailBad := "1234567qq.com"
	ok = CheckEmail(emailBad)
	if ok {
		t.Errorf("%v is valid email", emailBad)
	} else {
		t.Logf("%v is not valid email", emailBad)
	}
}

func Test_CheckMobile(t *testing.T) {
	mobile := "18310422106"
	ok := CheckMobile(mobile)
	if ok {
		t.Logf("%v is valid mobile", mobile)
	} else {
		t.Errorf("%v is not valid mobile", mobile)
	}
	mobileBad := "183104221061"
	ok = CheckMobile(mobileBad)
	if ok {
		t.Errorf("%v is valid mobile", mobileBad)
	} else {
		t.Logf("%v is not valid mobile", mobileBad)
	}
}

func Test_CheckEmailWithRegex(t *testing.T) {
	email := "1234567@qq.com"
	regex := DefaultEmailRegex
	ok := CheckEmailWithRegex(email, regex)
	if ok {
		t.Logf("%v is valid email", email)
	} else {
		t.Errorf("%v is not valid email", email)
	}
	regexBad := ""
	ok = CheckEmailWithRegex(email, regexBad)
	if ok {
		t.Errorf("%v is valid email", email)
	} else {
		t.Logf("%v is not valid email", email)
	}
}

func Test_CheckMobileWithRegex(t *testing.T) {
	mobile := "18310422106"
	regex := DefaultMobileRegex

	ok := CheckMobileWithRegex(mobile, regex)
	if ok {
		t.Logf("%v is mobile", mobile)
	} else {
		t.Errorf("%v is mobile", mobile)
	}
	regexBad := ""
	ok = CheckMobileWithRegex(mobile, regexBad)
	if ok {
		t.Errorf("%v is mobile", mobile)
	} else {
		t.Logf("%v is mobile", mobile)
	}
}

func Test_StringSliceEqual(t *testing.T) {
	a := []string{"123"}
	b := []string{"123"}
	c := []string{"1234"}

	ok := StringSliceEqual(a, b)
	if !ok {
		t.Errorf("%v==%v ? %v", a, b, ok)
	}
	ok = StringSliceEqual(a, c)
	if ok {
		t.Errorf("%v==%v ? %v", a, c, ok)
	}

	ok = StringSliceEqual(b, c)
	if ok {
		t.Errorf("%v==%v ? %v", b, c, ok)
	}
}
