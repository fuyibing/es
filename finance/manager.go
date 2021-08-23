// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package finance

var Manager = &Management{}

type Management struct{}

// 开票单管理.
func (o *Management) Bill() *BillManager {
	return NewBillManager()
}

// 订单管理.
func (o *Management) Order(m int) *OrderManager {
	return NewOrderManager(m)
}
