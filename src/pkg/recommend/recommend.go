package recommend

import (
	"math"
)

type Re struct {
	UserRatings map[int64]map[int64]float64
}

// PearsonCorrelation 计算皮尔逊相关系数
func (r Re) PearsonCorrelation(user1, user2 int64) float64 {
	// 找出两个用户共同评分的项目
	commonItems := make([]int64, 0)
	for item := range r.UserRatings[user1] {
		if _, ok := r.UserRatings[user2][item]; ok {
			commonItems = append(commonItems, item)
		}
	}

	if len(commonItems) == 0 {
		return 0 // 没有共同评分的项目
	}

	// 计算评分的平均值
	var sum1, sum2 float64
	for _, item := range commonItems {
		sum1 += r.UserRatings[user1][item]
		sum2 += r.UserRatings[user2][item]
	}
	mean1 := sum1 / float64(len(commonItems))
	mean2 := sum2 / float64(len(commonItems))

	// 计算皮尔逊相关系数的分子和分母
	var numerator, denominator1, denominator2 float64
	for _, item := range commonItems {
		numerator += (r.UserRatings[user1][item] - mean1) * (r.UserRatings[user2][item] - mean2)
		denominator1 += math.Pow(r.UserRatings[user1][item]-mean1, 2)
		denominator2 += math.Pow(r.UserRatings[user2][item]-mean2, 2)
	}

	denominator := math.Sqrt(denominator1) * math.Sqrt(denominator2)

	if denominator == 0 {
		return 0 // 避免除零错误
	}

	// 计算皮尔逊相关系数
	return numerator / denominator
}

// FindMostSimilarUser 获取与目标用户最相似的用户
func (r Re) FindMostSimilarUser(targetUser int64) int64 {
	var mostSimilarUser int64
	maxCorrelation := -1.0

	for user := range r.UserRatings {
		if user != targetUser {
			correlation := r.PearsonCorrelation(targetUser, user)
			if correlation > maxCorrelation {
				maxCorrelation = correlation
				mostSimilarUser = user
			}
		}
	}

	return mostSimilarUser
}
