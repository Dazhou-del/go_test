package zrl

import (
	"context"
	"fmt"
)

// 抽象的责任链接口
type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}

type BaseRuleChain struct {
	next RuleChain
}

func (br *BaseRuleChain) Next() RuleChain {
	return br.next
}

func (br *BaseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("implement me")
}

func (br *BaseRuleChain) NextApply(ctx context.Context, params map[string]interface{}) error {
	if br.Next() != nil {
		return br.Next().Apply(ctx, params)
	}

	return nil
}

func main() {
	ctx := context.Background()

	params := map[string]interface{}{
		"token": "abc",
		"age":   17,
		"name":  "daZhou",
	}

	if err := invokeRuleChain(ctx, params); err != nil {
		fmt.Println("链式调用出现错误2", err)
	}

	// 成功
}

func invokeRuleChain(ctx context.Context, params map[string]interface{}) error {
	// 组装成链 token->age->name 校验
	checkNameRuleChain := NewCheckNameRuleChain(nil)
	ageRuleChain := NewCheckAgeRuleChain(checkNameRuleChain)
	chain := NewCheckTokenRuleChain(ageRuleChain)

	if err := chain.Apply(ctx, params); err != nil {
		// 可return
		fmt.Println("链式调用出现错误", err)
	}

	// 成功
	return nil
}
