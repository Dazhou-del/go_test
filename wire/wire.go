//go:build wireinject

package wire

import "github.com/google/wire"

func initShopCart() (cat *ShopCat, err error) {
	wire.Build(
		NewShop,
		NewProductList,
	)

	return nil, nil
}
