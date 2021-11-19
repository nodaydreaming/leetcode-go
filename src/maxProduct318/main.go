package main

/**
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(words []string) int {
	maskMap := make(map[int]int)
	for _, word := range words {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		maskMap[mask] = max(len(word), maskMap[mask])
	}

	ans := 0
	for x, lenX := range maskMap {
		for y, lenY := range maskMap {
			if x&y == 0 {
				ans = max(ans, lenX*lenY)
			}
		}
	}
	return ans
}
