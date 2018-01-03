package charge

var Wx_pub_scan = map[string]interface{}{
	// 必填，客户端软件中展示的条码值，扫码设备扫描获取。
	"scan_code": "286801346868493272",
	// 可选，终端号，要求不同终端此号码不一样，会显示在对账单中，如A01、SH008等。
	"terminal_id": "SH008",
	// 可选，指定支付方式，指定不能使用信用卡支付可设置为 no_credit 。
	"limit_pay": "no_credit",
	// 可选，商品标记，代金券或立减优惠功能的参数。
	//"goods_tag": "your goods_tag",
}
