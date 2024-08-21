package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(s1, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	s1RuneArr := []rune(s1)
	s2RuneArr := []rune(s2)
	sort.Slice(s1RuneArr, func(i, j int) bool { return s1RuneArr[i] > s1RuneArr[j] })
	sort.Slice(s2RuneArr, func(i, j int) bool { return s2RuneArr[i] > s2RuneArr[j] })
	return string(s1RuneArr) == string(s2RuneArr)
}
