package main

import "fmt"

/**
给定一个正整数n ，你可以做如下操作：

如果n是偶数，则用n / 2替换n 。
如果n是奇数，则可以用n + 1或n - 1替换n 。
n变为 1 所需的最小替换次数是多少？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

// 记忆化搜索
var memo map[int]int

func integerReplacement1(n int) int {
	if n == 1 {
		return 0
	}
	res, ok := memo[n]
	if ok {
		return res
	}
	defer func() {
		memo[n] = res
	}()
	if n%2 == 0 {
		res = 1 + integerReplacement1(n/2)
	} else {
		res = 2 + min(integerReplacement1(n/2), integerReplacement1(n/2+1))
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(integerReplacement(11))
	fmt.Println(integerReplacement(22))
	fmt.Println(integerReplacement(33))
	fmt.Println(integerReplacement(44))
	fmt.Println(integerReplacement(55))
}
