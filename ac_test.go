package goAcAutoMachine

import (
	"fmt"
	"testing"
)

func TestAc(t *testing.T) {
	//content := "JT11111111，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
	//content := "JT1取消拦截,JT44取消改地址"
	//content := "JT11111111JT11111111撤消改地址，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
	content := "JT11111111撤消改地址,JT11111111撤消改地址"
	ac := NewAcAutoMachine()
	ac.AddPattern("拦截")
	ac.AddPattern("取消拦截")
	ac.AddPattern("取消拦")
	ac.AddPattern("改地址")
	ac.AddPattern("撤消改地址")
	ac.AddPattern("取消改地址")
	ac.AddPattern("取消改地")
	ac.AddPattern("消改地")
	ac.AddPattern("重量")
	ac.AddPattern("查重量")
	ac.AddPattern("催促")
	ac.AddPattern("撤")
	ac.Build()
	results := ac.Query(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		fmt.Println(result)
	}

	results = ac.QueryLast(content)
	fmt.Println("=============")
	for _, result := range results {
		fmt.Println(result)
	}
}
