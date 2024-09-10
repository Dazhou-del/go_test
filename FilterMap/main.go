package main

import (
	"fmt"
	"github.com/samber/lo"
)

type SkuAudit struct {
	age  int
	name string
}

func main() {
	Sku := make([]SkuAudit, 0, 20)

	skua := SkuAudit{
		age:  0,
		name: "daz",
	}

	Sku = append(Sku, skua)

	b := SkuAudit{
		age:  2,
		name: "daz",
	}
	Sku = append(Sku, b)
	c := SkuAudit{
		age:  2,
		name: "daz",
	}
	Sku = append(Sku, c)
	d := SkuAudit{
		age:  4,
		name: "daz",
	}
	Sku = append(Sku, d)

	storageRentalRuleIdList := lo.FilterMap(Sku, func(item SkuAudit, index int) (int, bool) {
		if item.age > 0 {
			return item.age, true
		}
		return 0, false
	})

	fmt.Println(storageRentalRuleIdList)
}
