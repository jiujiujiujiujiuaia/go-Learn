package main

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	numberMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for index := 0; index < len(s); index++ {
		if next := index + 1; next < len(s) && numberMap[string(s[next])] >
			numberMap[string(s[index])] {
			res += numberMap[string(s[next])] - numberMap[string(s[index])]
			index = next
		} else {
			res += numberMap[string(s[index])]
		}
	}
	return res
}

// @lc code=end
