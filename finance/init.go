// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package finance

// 数据类别.
//
// 索引`finance`/`fin`下, 以字段`_catalog`区分类别, 类别定义如下.
// 一经确定不可修改.
const (
	CatalogOrder         = 10 // 订单数据.
	CatalogGoodsOrigin   = 11 // 订单原始商品
	CatalogGoodsReplaced = 12 // 订单替换后商品
	CatalogEquity        = 20 // 权益数据.
	CatalogBill          = 30 // 开票单数据.
)

// ES索引定义.
const (
	DocumentIndex = "finance" // 索引名.
	DocumentType  = "fin"     // 索引类型名.
)

// 结算模型枚举.
//
// 索引`finance`/`fin`下, 以字段`_model`区分模型, 类别定义如下.
// 一经确定不可修改.
const (
	ModelDirect     = 101 // [1]直付结算/应付.
	ModelRenewal    = 102 // [2]换新结算/应付.
	ModelIntegral   = 103 // [3]积分结算/应付.
	ModelCommission = 104 // [4]佣金结算/应付.
	ModelHealthy    = 105 // [5]健康服务结算/应付.
	ModelPurchase   = 201 // [6]应收采购结算/应收.
)
