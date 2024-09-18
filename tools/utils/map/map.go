package _map

import "github.com/pengcainiao/zero/tools/utils"

type Map map[string]interface {
}

// 判断是否存在指定 key
func (m Map) Exits(key string) bool {
	_, exists := m[key]
	return exists
}

// 获取指定 key 下的 value，支持默认值
func (m Map) Get(key string, defaultValue ...interface{}) interface{} {
	value, exists := m[key]
	if exists {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

// 设置 key
func (m Map) Set(key string, value interface{}) Map {
	m[key] = value

	return m
}

// 移除 key
func (m Map) Remove(key string) Map {
	delete(m, key)

	return m
}

// 获取字符串
func (m Map) GetString(key string, defaultValue ...string) string {
	return m.Get(key, utils.FirstStringOr(defaultValue, "")).(string)
}

// 获取int
func (m Map) GetInt(key string, defaultValue ...int) int {
	return m.Get(key, utils.FirstIntOr(defaultValue, 0)).(int)
}

// 获取int64
func (m Map) GetInt64(key string, defaultValue ...int64) int64 {
	return m.Get(key, utils.FirstInt64Or(defaultValue, 0)).(int64)
}

// 获取32位浮点型数字
func (m Map) GetFloat(key string, defaultValue ...float32) float32 {
	return m.Get(key, utils.FirstFloat32Or(defaultValue, 0)).(float32)
}

// 获取64位浮点型数字
func (m Map) GetFloat64(key string, defaultValue ...float64) float64 {
	return m.Get(key, utils.FirstFloat64Or(defaultValue, 0)).(float64)
}

// 获取bool
func (m Map) GetBool(key string, defaultValue ...bool) float32 {
	return m.Get(key, utils.FirstBoolOr(defaultValue, false)).(float32)
}
