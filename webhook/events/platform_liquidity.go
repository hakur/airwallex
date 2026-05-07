// Package events provides typed webhook event structures for the platform liquidity domain.
// Platform Liquidity 领域类型化 webhook 事件结构体。
//
// 官方文档: https://www.airwallex.com/docs/developer-tools/webhooks/listen-for-webhook-events/platform-liquidity-program.md
//
// 事件映射表:
//
//	platform_liquidity_program.low_balance → PlatformLiquidityProgramLowBalanceEvent (Data: PlatformLiquidityProgramLowBalanceData)
//
// 更新工作流：当官方文档更新时，阅读上述 URL 并重新生成此文件。
package events

// --- Platform Liquidity Program: Low Balance ---

// PlatformLiquidityProgramLowBalanceEvent represents the platform_liquidity_program.low_balance webhook event.
// 当 Platform Liquidity Program 可用余额低于阈值时触发。
type PlatformLiquidityProgramLowBalanceEvent struct {
	Event
	Data PlatformLiquidityProgramLowBalanceData `json:"data"`
}

// PlatformLiquidityProgramLowBalanceData is the payload data for platform_liquidity_program.low_balance event.
type PlatformLiquidityProgramLowBalanceData struct {
	ID                  string                                      `json:"id"`
	Name                string                                      `json:"name"`
	Status              string                                      `json:"status"`
	LimitBalances       []PlatformLiquidityProgramLimitBalance      `json:"limit_balances"`
	SupportedCurrencies []PlatformLiquidityProgramSupportedCurrency `json:"supported_currencies"`
}

// PlatformLiquidityProgramLimitBalance represents a limit balance in the platform liquidity program payload.
type PlatformLiquidityProgramLimitBalance struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
	Total     float64 `json:"total"`
	Currency  string  `json:"currency"`
}

// PlatformLiquidityProgramSupportedCurrency represents a supported currency configuration in the platform liquidity program payload.
type PlatformLiquidityProgramSupportedCurrency struct {
	Currency            string  `json:"currency"`
	LowBalanceThreshold float64 `json:"low_balance_threshold"`
}
