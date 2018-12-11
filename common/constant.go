package common

import (
	"math"
	"math/big"
)

// !!! NOTICE, when first online, need set this const
// ex: time.Now().Unix()  1539771236
const (
	Dex4DOnlineTime = 1539154800
)

// Default Regex string
const (
	// DefaultEmailRegex  = "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
	DefaultEmailRegex  = "^([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
	DefaultMobileRegex = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
)

// DefaultLegalName default legal name
const (
	DefaultLegalName = LegalNameUSD

	LegalNameUSD = "USD"
	LegalNameCNY = "CNY"
)

const (
	DefaultLegalNameLower = LegalNameUSDLower

	LegalNameUSDLower = "usd"
	LegalNameCNYLower = "cny"
)

var CurrencyMap map[string]string = map[string]string{
	LegalNameCNYLower: LegalNameCNY,
	LegalNameUSDLower: LegalNameUSD,
}

// Keep float bit
const (
	KeepOneFloatBit = 1
	KeepTwoFloatBit = 2
	KeepSixFloatBit = 6

	// default legal bit to keep
	KeepLegalBit       = KeepTwoFloatBit
	KeepTokenAmountBit = KeepSixFloatBit
)

const tokenUnitLen = 18

// TokenUnit token bits should divide for converting to real value, care for the tokenUnitLen  <= 18
var TokenUnit = big.NewInt(int64(math.Pow10(tokenUnitLen)))
var TokenUnitF = big.NewFloat(math.Pow10(tokenUnitLen))

// User for market
const (
	OneDaySeconds = int64(86400)
)
