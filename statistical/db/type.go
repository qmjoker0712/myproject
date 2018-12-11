package db

import "time"

type ProfitGraph struct {
	Address               string `gorm:"column:address;" json:"address"`
	JYType                string `gorm:"column:type;" json:"type"`
	TxHash                string `gorm:"column:tx_hash;" json:"tx_hash"`
	TotalD4D              string `gorm:"column:total_d4d;" json:"total_d4d"`
	TokenID               int    `gorm:"column:token_id;" json:"token_id"`
	TaDividends           string `gorm:"column:ta_dividends;" json:"ta_dividends"`
	RefDividends          string `gorm:"column:ref_dividends;" json:"ref_dividends"`
	RefAddress            string `gorm:"column:ref_address;" json:"ref_address"`
	PcD4DCount            string `gorm:"column:pc_d4d_count;" json:"pc_d4d_count"`
	PcTokenCount          string `gorm:"column:pc_token_count;" json:"pc_token_count"`
	RtD4DCount            string `gorm:"column:rt_d4d_count;" json:"rt_d4d_count"`
	RtTokenCount          string `gorm:"column:rt_token_count;" json:"rt_token_count"`
	SellD4DCount          string `gorm:"column:sell_d4d_count;" json:"sell_d4d_count"`
	SellGettokenCount     string `gorm:"column:sell_gettoken_count;" json:"sell_gettoken_count"`
	WithdrawalsTokenCount string `gorm:"column:withdrawals_token_count;" json:"withdrawals_token_count"`
	SwapAddCount          string `gorm:"column:swap_add_count;" json:"swap_add_count"`
	SwapReduced4dCount    string `gorm:"column:swap_reduced4d_count;" json:"swap_reduced4d_count"`
	CreateTime            uint64 `gorm:"column:create_time;" json:"create_time"`
	TotalBalance          string `gorm:"column:total_balance;" json:"total_balance"`
	InsertTime            int64  `gorm:"column:insert_time;" json:"insert_time"`
}
type ProfitGraphTD struct {
	Time                 int64   ` json:"time"`
	TransactionDividends float64 ` json:"transaction_dividends"`
}
type ProfitGraphRD struct {
	Time               int64               ` json:"time"`
	ReferralDividends  float64             ` json:"referral_dividends"`
	MasternodeReferral float64             ` json:"masternode_referral"`
	MR                 []map[string]string ` json:"mr_list"`
}
type ProfitGraphTB struct {
	Time               int64                ` json:"time"`
	TokenBalance       float64              ` json:"token_balance"`
	Withdrawals        float64              ` json:"withdrawal"`
	SwapAddTokenCount  float64              ` json:"swap_add_token_count"`
	SwapReduced4dCount float64              ` json:"swap_reduced4d_count"`
	WithdrawalsList    []map[string]float64 ` json:"withdrawals_list"`
	SwapList           []map[string]float64 ` json:"swap_list"`
}
type ProfitGraphDEX4D struct {
	Time               int64                ` json:"time"`
	DEX4D              float64              ` json:"myproject"`
	PurchasesD4D       float64              ` json:"purchasesd4d"`
	PurchasesToken     float64              ` json:"purchases_token"`
	ReinvestmentsD4D   float64              ` json:"reinvestmentsd4d"`
	ReinvestmentsToken float64              ` json:"reinvestments_token"`
	SellD4D            float64              ` json:"selld4d"`
	SellToken          float64              ` json:"sell_token"`
	SwapAddTokenCount  float64              ` json:"swap_add_token_count"`
	SwapReduced4dCount float64              ` json:"swap_reduced4d_count"`
	BuyList            []map[string]float64 ` json:"buy_list"`
	SellList           []map[string]float64 ` json:"sell_list"`
	SwapList           []map[string]float64 ` json:"swap_list"`
	Reinvestment       []map[string]float64 ` json:"reinvestment_list"`
}
type AggregatedStatistic struct {
	Address               string `gorm:"column:address;" json:"address"`
	TokenID               int    `gorm:"column:token_id;" json:"token_id"`
	TokenName             string `gorm:"column:token_name;" json:"token_name"`
	EnteringTime          uint64 `gorm:"column:entering_time;" json:"entering_time"`
	EnteringRank          int    `gorm:"column:entering_rank;" json:"entering_rank"`
	TotalD4DCount         string `gorm:"column:total_d4d_count;" json:"d4d"`
	D4DShare              string `gorm:"column:d4d_share;" json:"share"`
	D4DRank               int    `gorm:"column:d4d_rank;" json:"d4d_rank"`
	InvestedTokenCount    string `gorm:"column:invested_token_count;" json:"invested_token_count"`
	InvestedTimes         int    `gorm:"column:invested_times;" json:"invested_times"`
	ReinvestedTokenCount  string `gorm:"column:reinvested_token_count;" json:"reinvested_token_count"`
	ReinvestedTimes       int    `gorm:"column:reinvested_times;" json:"reinvested_times"`
	WithdrawnTokenCount   string `gorm:"column:withdrawn_token_count;" json:"withdrawn_token_count"`
	WithdrawnTokenTimes   int    `gorm:"column:withdrawn_token_times;" json:"withdrawn_token_times"`
	RfDividendsTokenCount string `gorm:"column:rf_dividends_token_count;" json:"rf_dividends_token_count"`
	RfDividendsTokenTimes int    `gorm:"column:rf_dividends_token_times;" json:"rf_dividends_token_times"`
	TaDividendsTokenCount string `gorm:"column:ta_dividends_token_count;" json:"ta_dividends_token_count"`
	SellD4DCount          string `gorm:"column:sell_d4d_count;" json:"sell_d4d_count"`
	SellD4DTimes          int    `gorm:"column:sell_d4d_times;" json:"sell_d4d_times"`
	TokenBalance          string `gorm:"column:token_balance;" json:"total："`
}
type MarketSequence struct {
	TokenID               int     `gorm:"column:token_id;" json:"token_id"`
	TokenName             string  `gorm:"column:token_name;" json:"token_name"`
	TokenBalance          string  `gorm:"column:token_balance;" json:"token_balance"`
	TokenBalancePrice     string  `gorm:"column:token_balance_price;" json:"token_balance_price"`
	TokenBalanceShare     string  `gorm:"column:token_balance_share;" json:"token_balance_share"`
	TokenRealize          string  `gorm:"column:token_realize;" json:"token_realize"`
	TokenRealizePrice     string  `gorm:"column:token_realize_price;" json:"token_realize_price"`
	TokenRealizeShare     string  `gorm:"column:token_realize_share;" json:"token_realize_share"`
	TokenTransaction      string  `gorm:"column:token_transaction;" json:"token_transaction"`
	TokenTransactionPrice string  `gorm:"column:token_transaction_price;" json:"token_transaction_price"`
	TokenTransactionShare string  `gorm:"column:token_transaction_share;" json:"token_transaction_share"`
	TokenReferralPrice    string  `gorm:"column:token_referral_price;" json:"token_referral_price"`
	TokenReferralShare    string  `gorm:"column:token_referral_share;" json:"token_referral_share"`
	TokenReferral         string  `gorm:"column:token_referral;" json:"token_referral"`
	InsertTime            int64   `gorm:"column:insert_time;" json:"insert_time"`
	USD                   float64 `gorm:"column:usd;" json:"usd"`
}
type MarketBalance struct {
	TokenID       int    `gorm:"column:token_id;" json:"token_id"`
	TokenBalance  string `gorm:"column:token_balance;" json:"token_balance"`
	Devidends     string `gorm:"column:devidends;" json:"devidends"`
	Referral      string `gorm:"column:referral;" json:"referral"`
	SelloutAmount string `gorm:"column:sellout_amount;" json:"sellout_amount"`
}
type MarketData struct {
	Time  int64  `json:"time"`
	Value string `json:"value"`
}
type TokenTimeList struct {
	TokenName        string `json:"token_name"`
	TokenBalance     string `json:"token_balance"`
	TokenRealize     string `json:"token_realize"`
	TokenTransaction string `json:"token_transaction"`
	TokenReferral    string `json:"token_referral"`
	Time             int64  `json:"time"`
}
type TokenMarketList struct {
	TokenName             string `json:"token_name"`
	TokenBalance          string `json:"token_balance"`
	TokenRealize          string `json:"token_realize"`
	TokenTransaction      string `json:"token_transaction"`
	TokenReferral         string `json:"token_referral"`
	TokenBalanceShare     string `json:"token_balance_share"`
	TokenRealizeShare     string `json:"token_realize_share"`
	TokenTransactionShare string `json:"token_transaction_share"`
	TokenReferralShare    string `json:"token_referral_share"`
	Time                  int64  `json:"time"`
}
type TokenMarketData struct {
	BalanceShare string `gorm:"column:balance_share;" json:"balance_share"`
	InsertTime   int64  `gorm:"column:insert_time;" json:"insert_time"`
	TokenName    string `gorm:"column:token_name;" json:"token_name"`
	TokenID      int    `gorm:"column:token_id;" json:"token_id"`
}
type TokenMarketChart struct {
	TokenName    string  `json:"token_name"`
	TokenID      int     `json:"token_id"`
	InsertTime   int64   `json:"insert_time"`
	BalanceShare float64 `json:"balance_share"`
}
type TokenMarketChartList struct {
	InsertTime   int64   `json:"insert_time"`
	BalanceShare float64 `json:"balance_share"`
}
type BiggestOrder struct {
	TokenID        int     `gorm:"column:token_id;" json:"token_id"`
	TokenName      string  `gorm:"column:token_name;" json:"token_name"`
	EarnTokenName  string  `gorm:"column:earn_token_name;" json:"earn_token_name"`
	EarnTokenID    int     `gorm:"column:earn_token_id;" json:"earn_token_id"`
	Address        string  `gorm:"column:address;" json:"address"`
	TokenType      string  `gorm:"column:type;" json:"type"`
	InsertTime     int64   `gorm:"column:insert_time;" json:"insert_time"`
	OriginTime     uint64  `gorm:"column:origin_time;" json:"origin_time"`
	TokenCNYAmount string  `gorm:"column:token_cny_amount;" json:"token_cny_amount"`
	TokenUSDAmount string  `gorm:"column:token_usd_amount;" json:"token_usd_amount"`
	TokenAmount    string  `gorm:"column:token_amount;" json:"token_amount"`
	D4DAmount      string  `gorm:"column:d4d_amount;" json:"d4d_amount"`
	USD            float64 `gorm:"column:usd;" json:"usd"`
	CNY            float64 `gorm:"column:cny;" json:"cny"`
	TXHash         string  `gorm:"column:tx_hash;" json:"tx_hash"`
}
type Rank struct {
	Address    string `gorm:"column:address;" json:"address"`
	TokenID    int    `gorm:"column:token_id;" json:"token_id"`
	Time       uint64 `gorm:"column:time;" json:"time"`
	Rank       int    `gorm:"column:ranks;" json:"rank"`
	Number     int    `gorm:"column:number;" json:"number"`
	UserCount  int    `gorm:"column:user_count;" json:"user_count"`
	OriginTime int64  `gorm:"column:origin_time;" json:"origin_time"`
}
type TransactionDividendsData struct {
	Day                       uint64 ` json:"day"`
	DailyTransactionDividends string ` json:"daily_transaction_dividends"`
	RollingAverage            string ` json:"rolling_average"`
}
type TransactionDividend struct {
	Day                       uint64 `gorm:"column:day;" json:"day"`
	TokenID                   int    `gorm:"column:token_id;" json:"token_id"`
	Address                   string `gorm:"column:user_addr;" json:"user_addr"`
	DailyTransactionDividends string `gorm:"column:daily_transaction_dividends;" json:"daily_transaction_dividends"`
}
type AddressCount struct {
	Count int `gorm:"column:address_count;" json:"address_count"`
}
type TotalD4D struct {
	Count string `gorm:"column:d4d_count;" json:"d4d_count"`
}
type PersonTotalD4D struct {
	Address       string `gorm:"column:address;" json:"address"`
	TotalD4DCount string `gorm:"column:total_d4d_count;" json:"total_d4d_count"`
}
type TotalBalance struct {
	SelloutAmount string `gorm:"column:sellout_amount;" json:"sellout_amount"`
	Devidends     string `gorm:"column:devidends;" json:"devidends"`
	Referral      string `gorm:"column:referral;" json:"referral"`
}
type HeadInfo struct {
	Address       string  `gorm:"column:address;" json:"address"`
	TotalD4DCount float64 `gorm:"column:total_d4d_count;" json:"total_d4d_count"`
	TokenBalance  float64 `gorm:"column:token_balance;" json:"token_balance"`
	D4DShare      float64 `gorm:"column:d4d_share;" json:"d4d_share"`
}
type TotalCount struct {
	TotalCount int `gorm:"column:total;" json:"total"`
}
type RankHistroy struct {
	Time uint64 `gorm:"column:time;" json:"time"`
	Rank int    `gorm:"column:ranks;" json:"rank"`
}
type BiggestOrderList struct {
	TokenName   string `gorm:"column:token_name;" json:"token_name"`
	Address     string `gorm:"column:address;" json:"address"`
	TokenType   string `gorm:"column:type;" json:"type"`
	OriginTime  uint64 `gorm:"column:origin_time;" json:"origin_time"`
	TokenAmount string `gorm:"column:token_amount;" json:"token_amount"`
	D4DAmount   string `gorm:"column:d4d_amount;" json:"d4d_amount"`
}
type BiggestOrderChart struct {
	TokenName string `json:"token_name"`
	// Value     []BiggestOrderPrice `json:"value"`
	// Key      string              `json:"key"`
	Buy      []BiggestOrderPrice `json:"buy"`
	Sell     []BiggestOrderPrice `json:"sell"`
	Swap     []BiggestOrderPrice `json:"swap"`
	Reinvest []BiggestOrderPrice `json:"reinvest"`
}
type BiggestOrderPrice struct {
	Price string `gorm:"column:price;" json:"amounts"`
}
type TotalTokenBalance struct {
	Total string `gorm:"column:total;" json:"total"`
}
type BiggestOrderCount struct {
	OrderCount int `gorm:"column:order_count;" json:"order_count"`
}
type HolderStatistics struct {
	TokenRank             int    ` json:"d4d_rank"`
	TokenName             string ` json:"token_name"`
	Address               string ` json:"address"`
	EnteringTime          uint64 ` json:"entering_time"`
	TokenBalance          string ` json:"token_balance"`
	TokenShare            string ` json:"d4d_share"`
	InvestedTokenCount    string ` json:"invested_token_count"`
	InvestedTimes         int    ` json:"invested_times"`
	ReinvestedTokenCount  string ` json:"reinvested_token_count"`
	ReinvestedTimes       int    ` json:"reinvested_times"`
	WithdrawnTokenCount   string ` json:"withdrawn_token_count"`
	WithdrawnTokenTimes   int    ` json:"withdrawn_token_times"`
	RfDividendsTokenCount string ` json:"rf_dividends_token_count"`
	RfDividendsTokenTimes int    ` json:"rf_dividends_token_times"`
	SellD4DCount          string ` json:"sell_d4d_count"`
	SellD4DTimes          int    ` json:"sell_d4d_times"`
}

func TimeSelect() int {
	now := time.Now()
	todayZeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	switch {
	case todayZeroTime <= now.Unix() && now.Unix() < todayZeroTime+1*int64(7200):
		return 1
	case todayZeroTime+1*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+2*int64(7200):
		return 2
	case todayZeroTime+2*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+3*int64(7200):
		return 3
	case todayZeroTime+3*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+4*int64(7200):
		return 4
	case todayZeroTime+4*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+5*int64(7200):
		return 5
	case todayZeroTime+5*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+6*int64(7200):
		return 6
	case todayZeroTime+6*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+7*int64(7200):
		return 7
	case todayZeroTime+7*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+8*int64(7200):
		return 8
	case todayZeroTime+8*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+9*int64(7200):
		return 9
	case todayZeroTime+9*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+10*int64(7200):
		return 10
	case todayZeroTime+10*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+11*int64(7200):
		return 11
	case todayZeroTime+11*int64(7200) <= now.Unix() && now.Unix() < todayZeroTime+12*int64(7200):
		return 12
	}
	return 1

}
