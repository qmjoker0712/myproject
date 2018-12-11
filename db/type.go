package db

import "time"

type TokenTable struct {
	TokenID     int64  `gorm:"column:token_id;" json:"token_id"`
	Name        string `gorm:"column:name;" json:"name"`
	TokenAddr   string `gorm:"column:token_addr;" json:"token_addr"`
	State       int64  `gorm:"column:state;" json:"state"`
	MarketToken string `gorm:"column:market_token;" json:"market_token"`
}
type PayMent struct {
	PayMent float64 `gorm:"column:payment;" json:"payment"`
}

type GameTable struct {
	TokenID     int64  `gorm:"column:token_id" json:"token_id"`
	TokenName   string `gorm:"column:token_name" json:"token_name"`
	GameNameEN  string `gorm:"column:game_name_en" json:"game_name_en"`
	GameNameCH  string `gorm:"column:game_name_ch" json:"game_name_ch"`
	DescEN      string `gorm:"column:desc_en" json:"desc_en"`
	DescCH      string `gorm:"column:desc_ch" json:"desc_ch"`
	Address     string `gorm:"column:address" json:"address"`
	Logo        string `gorm:"column:logo" json:"logo"`
	Background  string `gorm:"column:background" json:"background"`
	State       int8   `gorm:"column:state" json:"state"`
	CreateTime  uint64 `gorm:"column:create_time;type:bigint(13)" json:"create_time"`
	PublishTime uint64 `gorm:"column:pulish_time;type:bigint(13)" json:"pulish_time"`
	OfflineTime uint64 `gorm:"column:offline_time;type:bigint(13)" json:"offline_time"`
}

type MsgQueueOffset struct {
	ConsumerID string `gorm:"column:consumer_id;type:varchar(64);primary_key" json:"consumer_id"`
	Offset     int64  `gorm:"column:offset" json:"offset"`
	SaveTime   uint64 `gorm:"column:save_time;type:bigint(13)" json:"save_time"`
}

type Ambassador struct {
	TokenName      string `gorm:"column:token_name;type:varchar(20);primary_key" json:"token_name"`
	UserAddress    string `gorm:"column:user_addr;type:varchar(42);primary_key" json:"user_addr"`
	IncomingTokens string `gorm:"column:incoming_tokens;type:varchar(40)" json:"incoming_tokens"`
	D4DMinted      string `gorm:"column:d4d_minted;type:varchar(40)" json:"d4d_minted"`
}

type LoginInfo struct {
	UserAddress   string `gorm:"column:user_addr;type:varchar(42)" json:"user_addr"`
	LoginType     string `gorm:"column:login_type" json:"login_type"`
	DeviceType    string `gorm:"column:device_type" json:"device_type"`
	DeviceID      string `gorm:"column:device_id" json:"device_id"`
	DeviceVersion string `gorm:"column:device_version" json:"device_version"`
	OSVersion     string `gorm:"column:os_version" json:"os_version"`
	BrowserInfo   string `gorm:"column:browser_info" json:"browser_info"`
	CreatedAt     time.Time
}
