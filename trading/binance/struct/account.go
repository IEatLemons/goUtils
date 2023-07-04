package binance_struct

type ApiRestrictions struct {
	IpRestrict                     bool `json:"ipRestrict"`                     // 是否限制ip访问
	CreateTime                     int  `json:"createTime"`                     // 创建时间
	EnableWithdrawals              bool `json:"enableWithdrawals"`              // 此选项允许通过此api提现。开启提现选项必须添加IP访问限制过滤器
	EnableInternalTransfer         bool `json:"enableInternalTransfer"`         // 此选项授权此密钥在您的母账户和子账户之间划转资金
	PermitsUniversalTransfer       bool `json:"permitsUniversalTransfer"`       // 授权该密钥可用于专用的万向划转接口，用以操作其支持的多种类型资金划转。各业务自身的划转接口使用权限，不受本授权影响
	EnableVanillaOptions           bool `json:"enableVanillaOptions"`           // 欧式期权交易权限
	EnableReading                  bool `json:"enableReading"`                  // 整体的读取权限
	EnableFutures                  bool `json:"enableFutures"`                  // 合约交易权限，需注意开通合约账户之前创建的API Key不支持合约API功能
	EnableMargin                   bool `json:"enableMargin"`                   // 此选项在全仓账户完成划转后可编辑
	EnableSpotAndMarginTrading     bool `json:"enableSpotAndMarginTrading"`     // 现货和杠杆交易权限
	TradingAuthorityExpirationTime int  `json:"tradingAuthorityExpirationTime"` // 现货和杠杆交易权限到期时间，如果没有则不返回该字段
}

// AccountInfo
type AccountInfo struct {
	MakerCommission            int             `json:"makerCommission"`            // 制造商委员会
	TakerCommission            int             `json:"takerCommission"`            // 接受者委员会
	BuyerCommission            int             `json:"buyerCommission"`            // 买家委员会
	SellerCommission           int             `json:"sellerCommission"`           // 卖家的佣金
	CommissionRates            CommissionRates `json:"commissionRates"`            // 佣金率
	CanTrade                   bool            `json:"canTrade"`                   // 可以交易
	CanWithdraw                bool            `json:"canWithdraw"`                // 可以提币
	CanDeposit                 bool            `json:"canDeposit"`                 // 可以充值
	Brokered                   bool            `json:"brokered"`                   // 是否代理
	RequireSelfTradePrevention bool            `json:"requireSelfTradePrevention"` // 要求自我贸易防范
	UpdateTime                 int             `json:"updateTime"`                 // 更新时间
	AccountType                string          `json:"accountType"`                // 账户类型
	Balances                   []Balance       `json:"balances"`                   // 账户余额
	Permissions                []string        `json:"permissions"`                // 权限
}
type CommissionRates struct {
	Maker  string `json:"maker"`  // 制造商
	Taker  string `json:"taker"`  // 接受者
	Buyer  string `json:"buyer"`  // 买家
	Seller string `json:"seller"` // 卖方
}
type Balance struct {
	Asset  string `json:"asset"`  // 资产名称（币名）
	Free   string `json:"free"`   // 可以金额
	Locked string `json:"locked"` // 冻结金额
}
