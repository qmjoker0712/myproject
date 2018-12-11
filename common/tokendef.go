package common

import (
	"sync"
)

/*
	Ugly codes, high deadlock risk, when cycle used the same map in one procedure call
	Should be reconstruction!!

	Todo: remove maps upper and lower maps, and provide method call instead of used maps
*/

// TOKEN symbol lowerLetter unique
const (
	TokenSymbolLowerALL = "all"
)

//=========================================================
var symbolIdToNameLowerMapLock sync.RWMutex
var symbolIdToNameLowerMap map[int]string = map[int]string{}

func GetSymbolIdToNameLowerMap(id int) string {
	// symbolIdToNameLowerMapLock.RLock()
	// defer symbolIdToNameLowerMapLock.RUnlock()

	return symbolIdToNameLowerMap[id]
}

func SymbolIdToNameLowerMap() map[int]string {
	// symbolIdToNameLowerMapLock.RLock()
	// defer symbolIdToNameLowerMapLock.RUnlock()
	return symbolIdToNameLowerMap
}

//=========================================================
var symbolIdToNameUpperMapLock sync.RWMutex
var symbolIdToNameUpperMap map[int]string = map[int]string{}

func GetSymbolIdToNameUpperMap(id int) string {
	// symbolIdToNameUpperMapLock.RLock()
	// defer symbolIdToNameUpperMapLock.RUnlock()

	return symbolIdToNameUpperMap[id]
}

//=========================================================
var symbolNameToIdMapLock sync.RWMutex
var symbolNameToIdMap map[string]int = map[string]int{TokenSymbolLowerALL: 0}

func GetSymbolNameToIdMap(name string) int {
	// symbolNameToIdMapLock.RLock()
	// defer symbolNameToIdMapLock.RUnlock()

	return symbolNameToIdMap[name]
}

func ExistSymbolNameToIdMap(name string) bool {
	// symbolNameToIdMapLock.RLock()
	// defer symbolNameToIdMapLock.RUnlock()

	_, ret := symbolNameToIdMap[name]
	return ret
}

//=========================================================
var symbolNameToNameLowerMapLock sync.RWMutex
var symbolNameToNameLowerMap map[string]string = map[string]string{}

func GetSymbolNameToNameLowerMap(name string) string {
	// symbolNameToNameLowerMapLock.RLock()
	// defer symbolNameToNameLowerMapLock.RUnlock()

	return symbolNameToNameLowerMap[name]
}

func SymbolNameToNameLowerMap() map[string]string {
	// symbolNameToNameLowerMapLock.RLock()
	// defer symbolNameToNameLowerMapLock.RUnlock()

	return symbolNameToNameLowerMap
}

//=========================================================
var symbolNameToNameUpperMapLock sync.RWMutex
var symbolNameToNameUpperMap map[string]string = map[string]string{}

func GetSymbolNameToNameUpperMap(name string) string {
	// symbolNameToNameUpperMapLock.RLock()
	// defer symbolNameToNameUpperMapLock.RUnlock()

	return symbolNameToNameUpperMap[name]
}

func SymbolNameToNameUpperMap() map[string]string {
	return symbolNameToNameUpperMap
}

//=========================================================
// SymbolInitZeroMap init SymbolMap zero
var symbolInitZeroMapLock sync.RWMutex
var symbolInitZeroMap map[string]float64 = map[string]float64{}

func GetSymbolInitZeroMap(name string) float64 {
	// symbolInitZeroMapLock.RLock()
	// defer symbolInitZeroMapLock.RUnlock()

	return symbolInitZeroMap[name]
}

func SymbolInitZeroMap() map[string]float64 {
	// symbolInitZeroMapLock.RLock()
	// defer symbolInitZeroMapLock.RUnlock()

	return symbolInitZeroMap
}

//=========================================================
// AllActiveTokenSymbolLower all active tokenSymbol
var allActiveTokenSymbolLowerLock sync.RWMutex
var allActiveTokenSymbolLower = []string{}

func AllActiveTokenSymbolLower() []string {
	// allActiveTokenSymbolLowerLock.RLock()
	// defer allActiveTokenSymbolLowerLock.RUnlock()

	return allActiveTokenSymbolLower
}

//=========================================================
// TokenSymbolNameMap token symbol-name map
var tokenSymbolNameMapLock sync.RWMutex
var tokenSymbolNameMap = map[string]string{}

func GetTokenSymbolNameMap(name string) string {
	// tokenSymbolNameMapLock.RLock()
	// defer tokenSymbolNameMapLock.RUnlock()

	return tokenSymbolNameMap[name]
}

func ExistTokenSymbolNameMap(name string) bool {
	// tokenSymbolNameMapLock.RLock()
	// defer tokenSymbolNameMapLock.RUnlock()

	_, ret := tokenSymbolNameMap[name]
	return ret
}

func lockAll() {
	symbolIdToNameLowerMapLock.Lock()
	symbolIdToNameUpperMapLock.Lock()
	symbolNameToIdMapLock.Lock()
	symbolNameToNameLowerMapLock.Lock()
	symbolNameToNameUpperMapLock.Lock()
	symbolInitZeroMapLock.Lock()
	allActiveTokenSymbolLowerLock.Lock()
	tokenSymbolNameMapLock.Lock()
}

func unlockAll() {
	symbolIdToNameLowerMapLock.Unlock()
	symbolIdToNameUpperMapLock.Unlock()
	symbolNameToIdMapLock.Unlock()
	symbolNameToNameLowerMapLock.Unlock()
	symbolNameToNameUpperMapLock.Unlock()
	symbolInitZeroMapLock.Unlock()
	allActiveTokenSymbolLowerLock.Unlock()
	tokenSymbolNameMapLock.Unlock()
}

func RLockAll() {
	symbolIdToNameLowerMapLock.RLock()
	symbolIdToNameUpperMapLock.RLock()
	symbolNameToIdMapLock.RLock()
	symbolNameToNameLowerMapLock.RLock()
	symbolNameToNameUpperMapLock.RLock()
	symbolInitZeroMapLock.RLock()
	allActiveTokenSymbolLowerLock.RLock()
	tokenSymbolNameMapLock.RLock()
}

func RUnlockAll() {
	symbolIdToNameLowerMapLock.RUnlock()
	symbolIdToNameUpperMapLock.RUnlock()
	symbolNameToIdMapLock.RUnlock()
	symbolNameToNameLowerMapLock.RUnlock()
	symbolNameToNameUpperMapLock.RUnlock()
	symbolInitZeroMapLock.RUnlock()
	allActiveTokenSymbolLowerLock.RUnlock()
	tokenSymbolNameMapLock.RUnlock()
}
