// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2"

	"github.com/fuyibing/es/finance"
)

func TestFinanceBillCreate(t *testing.T) {
	ctx := log.NewContext()
	res, err := finance.Manager.
		Bill().
		Create(ctx, "20210101", map[string]interface{}{
			"bill_id":        95437,
			"bill_no":        "210823000098",
			"invoice_amount": 78187.20,
			"invoice_title":  "紫金财产保险股份有限公司河北分公司",
		})

	if err != nil {
		t.Errorf("bill create error: %v.", err)
		return
	}

	t.Logf("bill create result: %s.", res.DocumentId)

	// billId	billNo	statementNo	claimNo	insureCompanyId	insureCompany	policyNo	policyType	claimMethod	isOnline	invoiceType	orderCount	claimedCount	goodsCount	invoiceAmount	taxRate	taxAmount	amountWithoutTax	temporaryAmount	billStatus	lastClaimTime	claimStatus	lastOrderId	lastOrderType	invoiceTitle	taxpayerIdNumber	addressAndPhone	bankAndAccount	remark	completeTime	downloadNum	isDel	gmtCreated	gmtUpdated	directClaim	billType	purchaserId	sellerId	uploadStatus	receiptedAmount	receiptedStatus	receiveDate	receiveTime	claimFormStatus
	// 95437	210823000098	DS20210823124868		165102	紫金财产保险股份有限公司河北分公司	20656213000021000012	0	2	1	2	7	0	7	78187.20	0.00	0.000000	0.000000	0	1	1970-01-01 08:00:00	0	0	2	紫金财产保险股份有限公司河北分公司	91130100564864728C	河北省石家庄市裕华区裕华东路133号方北大厦B座12层1201-1208室、11层1107-1108室 0311-89297705	中国建设银行股份有限公司石家庄广安街支行 13001615236050507345	210823000098		0	0	2021-08-23 14:53:26	2021-08-23 14:53:35	1	0	165102	160060	0	0.00	0			0
}

func TestFinanceBillDelete(t *testing.T) {}

func TestFinanceBillUpdate(t *testing.T) {}

func TestFinanceBillSearch(t *testing.T) {}
