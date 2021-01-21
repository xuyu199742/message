package utils

import (
	"strconv"

	"github.com/cstockton/go-conv"
	"github.com/fatih/structs"
	"time"
)

// StructToMap 转换struct为map
func StructToMap(s interface{}) map[string]interface{} {
	return structs.Map(s)
}

// StringToInt 字符串转数值
func StringToInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func UnixMilliToTime(mse int64) time.Time {
	str, _ := conv.String(mse)
	var (
		s  int64
		ms int64
	)
	s, _ = conv.Int64(str[:10])
	if len(str) >= 13 {
		ms, _ = conv.Int64(str[10:])
	}
	return time.Unix(s, ms)
}
