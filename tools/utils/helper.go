package utils

// IfI 三元运算
func IfI(condition bool, i, d interface{}) interface{} {
	if condition {
		return i
	}
	return d
}

// 获取第一个字符串
func FirstStringOr(arr []string, defaultValue string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个int
func FirstIntOr(arr []int, defaultValue int) int {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个int8
func FirstInt8Or(arr []int8, defaultValue int8) int8 {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个int16
func FirstInt16Or(arr []int16, defaultValue int16) int16 {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个int64
func FirstInt64Or(arr []int64, defaultValue int64) int64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个float32
func FirstFloat32Or(arr []float32, defaultValue float32) float32 {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个字符串
func FirstFloat64Or(arr []float64, defaultValue float64) float64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}

// 获取第一个bool
func FirstBoolOr(arr []bool, defaultValue bool) bool {
	if len(arr) > 0 {
		return arr[0]
	}
	return defaultValue
}
