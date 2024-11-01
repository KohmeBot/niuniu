package prob

import (
	"math/rand/v2"
)

// HitProb 命中概率
func HitProb(pro float64) bool {
	return rand.Float64() < pro
}

type Value[T any] struct {
	Prob float64
	V    T
}

type ProbabilityGroup[T any] []Value[T]

// Hit 命中(返回组内命中的值)
func (g ProbabilityGroup[T]) Hit() Value[T] {
	// 计算总权重
	totalProb := 0.0
	for _, value := range g {
		totalProb += value.Prob
	}

	randomValue := rand.Float64() * totalProb

	// 根据随机值选择命中的值
	for _, value := range g {
		if randomValue < value.Prob {
			return value
		}
		randomValue -= value.Prob
	}

	// 如果没有找到命中值，返回零值（根据你的需求可以调整）
	return Value[T]{}

}
