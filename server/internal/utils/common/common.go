package utils_common

import "math"

// 计算较昨日的增长/减少百分比
// 参数：今日收益，昨日收益
// 返回值：百分比数字（正数表示增长，负数表示减少）
func CalculateChangePercent(today, yesterday float64) float64 {
	// 如果昨日收益为0，特殊处理
	if yesterday == 0 {
		// 今日收益也为0，变化率为0
		if today == 0 {
			return 0.0
		}
		// 今日收益>0，返回正无穷
		if today > 0 {
			return math.Inf(1)
		}
		// 今日收益<0，返回负无穷
		return math.Inf(-1)
	}

	// 正常计算百分比变化
	// 使用昨日的绝对值作为分母，这样百分比变化更直观
	return ((today - yesterday) / math.Abs(yesterday)) * 100
}
