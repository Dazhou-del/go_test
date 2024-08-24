package zrl

import (
	"context"
	"fmt"
)

type CheckAgeRuleChain struct {
	BaseRuleChain
}

func NewCheckAgeRuleChain(next RuleChain) RuleChain {
	return &CheckAgeRuleChain{
		BaseRuleChain{
			next: next},
	}
}

func (c *CheckAgeRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	age, _ := params["age"].(int)
	if age <= 18 {
		return fmt.Errorf("invalidage age :%d", params["token"])
	}

	if err := c.NextApply(ctx, params); err != nil {
		fmt.Println("名字检查错误处理...")
		return err
	}

	fmt.Println("名字检查成功")
	return nil
}
