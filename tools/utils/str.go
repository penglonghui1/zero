package utils

import "strconv"

func IfStr(condition bool, str1, str2 string) string {
	if condition {
		return str1
	}
	return str2
}

func StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func IsInArray(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func StrOr(str string, arr ...string) string {
	if str != "" {
		return str
	}

	for _, v := range arr {
		if "" != v {
			return v
		}
	}
	return ""
}

func StrMapKeys(strMap map[string]string) []string {
	var ids []string
	for id := range strMap {
		ids = append(ids, id)
	}
	return ids
}
