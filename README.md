# permutation in string

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

## Examples

```
Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
 

Constraints:

1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.
```

## 解析

題目要確認 s2 的子字串包含 s1字串已各種變換字元次序的形式

簡單說就是判斷 s2 具有 s1字串長度相同的子字串 並且組成的字元出現次數相同

所以如果主要就是要找出 s1 字串內所有字元的出現次數

然後透過 slide window 的方式，每次比對與 s1 字元長度相同 slide windows 檢查字元出現次數是否相同

如果全部 slide window 長度接不同則代表沒有

因為所有字元都是小寫 a - z 所以只要用長度 26 的整數陣列 即可儲存所有頻率

先 loop s1 算出所有字元出現次數 在 freq

初始化左界 left 與右界 right 為 0, count = len(s1)

每次檢查 freq [ s[right]- 'a'] 如果大於 1, 更新 count = count - 1

更新 freq [ s[right]- 'a'] -= 1 // 把遇到的字元次數遞減

更新 right += 1

如果 count == 0 代表 已經找到所有出現過的字元 回傳 true

當 right - left >= len(s1)

檢查 freq [ s[left] -'a'] >= 0 如果是 則 更新 count = count + 1

更新 freq [ s[left] -'a'] += 1, left += 1


## 實作

```golang
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
```