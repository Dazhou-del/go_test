package main

import (
	"context"
	"fmt"
)

type RuleHandler func(ctx context.Context, params map[string]interface{}) error

func main() {

	if err := TestRuleChainV1(); err != nil {
		fmt.Println("失败")
	}

	// 领取奖励
	fmt.Println("ok")
}

func TestRuleChainV1() error {
	var checkTokenRule RuleHandler = func(ctx context.Context, params map[string]interface{}) error {
		// 校验token是否合法
		if params["token"] != "abc" {
			return fmt.Errorf("token不合法")
		}

		return nil
	}

	var checkAgeRule RuleHandler = func(ctx context.Context, params map[string]interface{}) error {
		// 校验年龄是否合法
		if params["age"] != "ccc" {
			return fmt.Errorf("年龄不合法")
		}
		return nil
	}

	var checkAuthorizedRule RuleHandler = func(ctx context.Context, params map[string]interface{}) error {
		// 校验授权是否合法
		if params["authorized"] != "ddd" {
			return fmt.Errorf("授权不合法")
		}
		return nil
	}

	// 组装调用链
	ruleChain := []RuleHandler{
		checkTokenRule,
		checkAgeRule,
		checkAuthorizedRule,
	}

	ctx := context.Background()

	params := map[string]interface{}{
		"token":      "abc",
		"age":        "ccc",
		"authorized": "ddd",
	}

	for _, rule := range ruleChain {
		if err := rule(ctx, params); err != nil {
			// 校验未通过，终止发奖流程

			return fmt.Errorf("chun ", err)
		}
	}

	return nil
}
