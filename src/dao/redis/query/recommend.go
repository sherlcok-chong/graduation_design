package query

import (
	"GraduationDesign/src/global"
	"GraduationDesign/src/pkg/recommend"
	"context"
	"fmt"
	"log"
	"math"
)

func (q Queries) ReadRE(ctx context.Context) {
	keys, err := q.rdb.Keys(ctx, "ratings:user:*:item:*").Result()
	if err != nil {
		log.Fatal(err)
	}
	global.Re = recommend.Re{UserRatings: make(map[int64]map[int64]float64)}
	for _, key := range keys {
		value, err := q.rdb.Get(ctx, key).Float64()
		if err != nil {
			log.Fatal(err)
		}

		var userID, itemID int64
		n, err := fmt.Sscanf(key, "ratings:user:%d:item:%d", &userID, &itemID)
		if n != 2 || err != nil {
			log.Fatalf("failed to parse key: %s", key)
		}

		if _, ok := global.Re.UserRatings[userID]; !ok {
			global.Re.UserRatings[userID] = make(map[int64]float64)
		}

		global.Re.UserRatings[userID][itemID] = value
	}
}
func (q Queries) UpdateRating(userID, itemID int64, newRating float64) error {
	userRatings, ok := global.Re.UserRatings[userID]
	if !ok {
		return fmt.Errorf("user ID %d not found", userID)
	}
	userRatings[itemID] = newRating

	// 将修改后的数据存储到 Redis 中
	key := fmt.Sprintf("ratings:user:%d:item:%d", userID, itemID)
	ctx := context.Background()
	err := q.rdb.Set(ctx, key, newRating, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetUserPreferences 获取用户喜好的物品列表
func (q Queries) GetUserPreferences(userID int64) ([]int, error) {
	ctx := context.Background()
	// 获取以 ratings:user:UserID:item:ItemID 格式存储的键的列表
	keys, err := q.rdb.Keys(ctx, fmt.Sprintf("ratings:user:%d:item:*", userID)).Result()
	if err != nil {
		return nil, err
	}

	// 解析键中的物品 ID
	var itemIDs []int
	for _, key := range keys {
		var itemID int
		// 解析出 ItemID
		n, err := fmt.Sscanf(key, "ratings:user:%d:item:%d", &userID, &itemID)
		if n != 2 || err != nil {
			return nil, fmt.Errorf("failed to parse key: %s", key)
		}
		itemIDs = append(itemIDs, itemID)
	}

	return itemIDs, nil
}

// CalculateItemSimilarity 计算物品之间的相似度
func (q Queries) CalculateItemSimilarity(item1ID, item2ID int) (float64, error) {
	ctx := context.Background()

	// 获取所有评分过物品1的用户的列表
	keys1, err := q.rdb.Keys(ctx, fmt.Sprintf("ratings:user:*:item:%d", item1ID)).Result()
	if err != nil {
		return 0, err
	}
	// 获取所有评分过物品2的用户的列表
	keys2, err := q.rdb.Keys(ctx, fmt.Sprintf("ratings:user:*:item:%d", item2ID)).Result()
	if err != nil {
		return 0, err
	}

	// 计算共同评分的用户数
	commonUsers := 0
	// 遍历所有评分过物品1的用户的列表
	for _, key1 := range keys1 {
		// 遍历所有评分过物品2的用户的列表
		for _, key2 := range keys2 {
			// 如果某个用户同时评分过物品1和物品2，则认为存在共同评分用户
			if key1 == key2 {
				commonUsers++
				break
			}
		}
	}

	// 计算余弦相似度
	similarity := float64(commonUsers) / (math.Sqrt(float64(len(keys1))) * math.Sqrt(float64(len(keys2))))

	return similarity, nil
}
