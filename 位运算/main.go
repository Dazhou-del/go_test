package main

import "fmt"

type SpuInfoSign int32

const (
	ProductInfoSignSku SpuInfoSign = 1 << 0 // SKU信息

	ProductInfoSignTarget SpuInfoSign = 1 << 1 // 商品条件

	ProductInfoSignService SpuInfoSign = 1 << 2 // 商品服务

	ProductInfoSignPicture SpuInfoSign = 1 << 3 // 商品图片

	ProductInfoSignLogistics SpuInfoSign = 1 << 4 // 商品物流信息

	ProductInfoSignShop SpuInfoSign = 1 << 5 // 商家信息

	ProductInfoSignCategory SpuInfoSign = 1 << 6 // 商品分类信息

	ProductInfoSignDetail SpuInfoSign = 1 << 7 // 商品详情信息

	ProductInfoSignMsInfo SpuInfoSign = 1 << 8 // 码商信息

	ProductInfoSignRemind SpuInfoSign = 1 << 9 // 短信提醒信息

	ProductInfoSignControl SpuInfoSign = 1 << 10 // 商品管控记录

	ProductInfoSignLogisticsDetail SpuInfoSign = 1 << 11 // 商品物流信息-详细

	ProductInfoSignLogisticsFast   SpuInfoSign = 1 << 12 // 商品极速发货信息
	ProductInfoSignCategoryRiseSku SpuInfoSign = 1 << 13 // 分类套餐信息
	ProduceInfoSignLogisticsDetail SpuInfoSign = 1 << 14 // 商品v2物流信息
	ProductThirdLogistics          SpuInfoSign = 1 << 15 // 第三方物流
	ProductVersion                 SpuInfoSign = 1 << 16 // 商品版本信息
	ProductInfoSignBrand           SpuInfoSign = 1 << 17 // 品牌
	ProductInfoSignModel           SpuInfoSign = 1 << 18 // 型号

	ProductInfoSignServiceDesc SpuInfoSign = 1 << 19 // 租赁服务设备说明
)

func main() {
	css := int32(ProductInfoSignSku |
		ProductInfoSignTarget |
		ProductInfoSignService |
		ProductInfoSignPicture |
		ProductInfoSignMsInfo |
		ProductInfoSignLogistics |
		ProductInfoSignDetail,
	)
	fmt.Println(css)
	spuSign := 9
	c := 1
	spuSign |= c
	fmt.Println(spuSign)

	spuSign &^= c

	fmt.Println(spuSign)
}
