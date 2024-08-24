package zrl

import (
	"context"
	"fmt"
)

type CheckTokenRuleChain struct {
	BaseRuleChain
}

func NewCheckTokenRuleChain(next RuleChain) RuleChain {
	return &CheckTokenRuleChain{
		BaseRuleChain: BaseRuleChain{
			next: next},
	}
}

func (c *CheckTokenRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	if params["token"] != "token" {
		return fmt.Errorf("invalidage token :%s", params["token"])
	}

	if err := c.NextApply(ctx, params); err != nil {
		fmt.Println("检查年龄错误后处理...")
		return err
	}

	fmt.Println("校验年龄通过....")
	return nil
}
