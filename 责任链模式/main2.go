package main

import (
	"context"
	"fmt"
)

type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}

type BaseRuleChain struct {
	next RuleChain
}

func (b *BaseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("notimplement")
}

func (b *BaseRuleChain) Next() RuleChain {
	return b.next
}

func (b *BaseRuleChain) applyNext(ctx context.Context, params map[string]interface{}) error {
	if b.Next() != nil {
		return b.Next().Apply(ctx, params)
	}
	return nil
}

type CheckTokenRule struct {
	BaseRuleChain
}

func NewCheckTokenRule(next RuleChain) *CheckTokenRule {
	return &CheckTokenRule{
		BaseRuleChain: BaseRuleChain{next: next},
	}
}
func (c *CheckTokenRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验token是否合法
	if params["token"] != "abc" {
		return fmt.Errorf("token不合法")
	}

	if err := c.applyNext(ctx, params); err != nil {
		fmt.Println("check token rule err post process....")

		return err
	}

	fmt.Println("check token rule common post process....")

	return nil
}

type CheckAgeRule struct {
	BaseRuleChain
}

func NewCheckAgeRule(next RuleChain) RuleChain {
	return &CheckAgeRule{
		BaseRuleChain: BaseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAgeRule) Apply(ctx context.Context, params map[string]interface{}) error {
	//校验age是否合法
	age, _ := params["age"].(int)
	if age < 18 {
		return fmt.Errorf("invalidage:%d", age)
	}

	if err := c.applyNext(ctx, params); err != nil {
		//err post process
		fmt.Println("check age rule err post process...")
		return err
	}

	fmt.Println("check age rule common post process....")
	return nil
}

type CheckAuthorizedStatusRule struct {
	BaseRuleChain
}

func NewCheckAuthorizedStatusRule(next RuleChain) RuleChain {
	return &CheckAuthorizedStatusRule{
		BaseRuleChain: BaseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAuthorizedStatusRule) Apply(ctx context.Context, params map[string]interface{}) error {
	//校验age是否合法
	age, _ := params["age"].(int)
	if age < 18 {
		return fmt.Errorf("invalidage:%d", age)
	}

	if err := c.applyNext(ctx, params); err != nil {
		//err post process
		fmt.Println("check age rule err post process...")
		return err
	}

	fmt.Println("check age rule common post process....")
	return nil
}

func main() {
	checkAuthorizedStatusRule := NewCheckAuthorizedStatusRule(nil)
	checkAgeRule := NewCheckAgeRule(checkAuthorizedStatusRule)
	tokenRule := NewCheckTokenRule(checkAgeRule)

	if err := tokenRule.Apply(context.Background(), map[string]interface{}{
		"token": "abc",
		"age":   12,
	}); err != nil {
		// 可以return
		fmt.Println("失败")
	}

	// 成功
}
