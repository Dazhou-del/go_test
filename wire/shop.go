package wire

import "fmt"

type ShopCat struct {
	productList *ProductList
}

type ProductList struct{}

func NewShop(pro *ProductList) (*ShopCat, error) {
	return &ShopCat{productList: pro}, nil
}

func NewProductList() (*ProductList, error) {
	return &ProductList{}, nil
}

func initShopCat() (cat *ShopCat, err error) {
	// 构建ProductList对象
	list, err := NewProductList()
	if err != nil {
		return nil, err
	}

	shopCat, err := NewShop(list)
	if err != nil {
		return nil, err
	}

	return shopCat, nil
}

func callShopCat(cat *ShopCat) {
	fmt.Println("call shopcat")
}
