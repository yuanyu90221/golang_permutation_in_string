package perm_in_string

import "log"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 || len(s1) > len(s2) {
		return false
	}
	// charMap
	var freq [26]int
	for idx := 0; idx < len(s1); idx++ {
		freq[s1[idx]-'a']++
	}
	left := 0
	count := len(s1)
	right := 0
	for right < len(s2) {
		// 右界根據遇到 char 做更新
		log.Printf("freq: %v", freq)
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']-- // 這邊計次需要把每次遇到有在 freq 內的都做紀錄 就算是 -1也要紀錄
		right++
		if count == 0 { // 右界遇到所有的 s1
			return true
		}
		if right-left >= len(s1) { // 當 slide window 大等跟等於 len(s1) 左界右移
			if freq[s2[left]-'a'] >= 0 { // 回補 曾經出現過在 freq的字元
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
