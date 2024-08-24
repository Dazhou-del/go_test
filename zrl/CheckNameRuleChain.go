package zrl

import (
	"fmt"
	"golang.org/x/net/context"
)

type CheckNameRuleChain struct {
	BaseRuleChain
}

func NewCheckNameRuleChain(next RuleChain) RuleChain {
	return &CheckNameRuleChain{
		BaseRuleChain{next: next},
	}
}

func (c *CheckNameRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	if params["daZhou"] != "daZhou" {
		return fmt.Errorf("invalidage name :%s", params["token"])
	}

	if err := c.NextApply(ctx, params); err != nil {
		fmt.Println("剩下错误处理...")
		return err
	}

	fmt.Println("年龄检查成功")
	return nil
}
