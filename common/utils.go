package common

import (
	"myproject/ginrpc"
	"myproject/rpc"
	"math/big"
	"reflect"
	"regexp"
	"runtime"

	"github.com/mitchellh/copystructure"
)

// DeepCopy deep copy dst must be a pointer
func DeepCopy(dst interface{}) (src interface{}, err error) {
	return copystructure.Copy(dst)
}

// BigRatDiv18BitToRat big Rat div 10^18 for big.Rat
func BigRatDiv18BitToRat(number *big.Rat) *big.Rat {
	tu := big.NewRat(0, 1).SetInt(TokenUnit)
	z := big.NewRat(0, 1).Inv(tu)
	return number.Mul(number, z)
}

// BigIntDiv18BitToRat big int div 10^18 for big.Rat
func BigIntDiv18BitToRat(number *big.Int) *big.Rat {
	return big.NewRat(1, 1).SetFrac(number, TokenUnit)
}

// BigIntDivStrWithPrec remain prec digits for big int div
func BigIntDivStrWithPrec(number, unit *big.Int, prec int) string {
	return big.NewRat(1, 1).SetFrac(number, unit).FloatString(prec)
}

// BigIntAdd  bigInt Add
func BigIntAdd(x, y *big.Int) *big.Int {
	return x.Add(x, y)
}

// BigIntSub  bigInt sub
func BigIntSub(x, y *big.Int) *big.Int {
	return x.Sub(x, y)
}

// BigIntMul  bigInt mul
func BigIntMul(x, y *big.Int) *big.Int {
	return x.Mul(x, y)
}

// BigIntDiv  bigInt div
func BigIntDiv(x, y *big.Int) *big.Int {
	return x.Div(x, y)
}

// BigIntDivToRat  bigInt div convert result to rat
func BigIntDivToRat(x, y *big.Int) *big.Rat {
	return big.NewRat(0, 1).SetFrac(x, y)
}

// BigIntDivToFloat32  bigInt div convert result to float32
func BigIntDivToFloat32(x, y *big.Int) (f float32, exact bool) {
	return big.NewRat(0, 1).SetFrac(x, y).Float32()
}

// BigIntDivToFloat64  bigInt div convert result to float64
func BigIntDivToFloat64(x, y *big.Int) (f float64, exact bool) {
	return big.NewRat(0, 1).SetFrac(x, y).Float64()
}

// BigIntToFloat32WithTokenUnit  bigInt convert to float32 div token unit
func BigIntToFloat32WithTokenUnit(x, unit *big.Int) (f float32) {
	result, _ := big.NewRat(0, 1).SetFrac(x, unit).Float32()
	return result
}

// BigIntToFloat64WithTokenUnit  bigInt convert to float64 div token unit
func BigIntToFloat64WithTokenUnit(x, unit *big.Int) (f float64) {
	result, _ := big.NewRat(0, 1).SetFrac(x, unit).Float64()
	return result
}

// BigIntToFloat32WithDefaultTokenUnit  bigInt convert to float32 div default tokenUnit
func BigIntToFloat32WithDefaultTokenUnit(x *big.Int) (f float32) {
	result, _ := big.NewRat(0, 1).SetFrac(x, TokenUnit).Float32()
	return result
}

// BigIntToFloat64WithDefaultTokenUnit  bigInt convert to float64 div default tokenUnit
func BigIntToFloat64WithDefaultTokenUnit(x *big.Int) (f float64) {
	result, _ := big.NewRat(0, 1).SetFrac(x, TokenUnit).Float64()
	return result
}

// CheckEmail check email
func CheckEmail(email string) (b bool) {
	return CheckEmailWithRegex(email, DefaultEmailRegex)
}

// CheckMobile check mobile
func CheckMobile(mobile string) (b bool) {
	return CheckMobileWithRegex(mobile, DefaultMobileRegex)
}

// CheckEmailWithRegex check email with regex
func CheckEmailWithRegex(email string, regex string) (b bool) {
	if email == "" || regex == "" {
		return false
	}
	if m, _ := regexp.MatchString(regex, email); !m {
		return false
	}
	return true
}

// CheckMobileWithRegex check mobile with regex
func CheckMobileWithRegex(mobil string, regex string) (b bool) {
	if mobil == "" || regex == "" {
		return false
	}
	if m, _ := regexp.MatchString(regex, mobil); !m {
		return false
	}
	return true
}

// StringSliceEqual string slice check equal
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Convert2GinAPI convert rpc.API to ginrpc.API
func Convert2GinAPI(api *rpc.API) *ginrpc.API {
	return &ginrpc.API{
		Namespace: api.Namespace,
		Version:   api.Version,
		Service:   api.Service,
		Public:    api.Public,
	}
}

// GetBeginTime
// compare with OnlineTime, if gt OnlineTime, use NewTime; else use OnlineTime
// Args: newTime int64(ex: time.Now().Unix()  1539771236)
func GetBeginTime(newTime int64) int64 {
	if newTime > Dex4DOnlineTime {
		return newTime
	} else {
		return Dex4DOnlineTime
	}
}

// GetFuncName get name of a method
func GetFuncName(f func()) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
