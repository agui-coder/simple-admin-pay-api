package payno

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

const (
	PayNo         = "PAY_NO:"
	OrderNoPrefix = "P"
)

// Generate 生成支付编号
func Generate(db *redis.Redis, prefix string) (string, error) {
	// 构建没有前缀的编号
	noPrefix := prefix + time.Now().Format("20060102150405")
	key := PayNo + noPrefix
	no, err := db.Incr(key)
	if err != nil {
		return "", err
	}
	err = db.Expire(key, 60)
	if err != nil {
		return "", err
	}
	// 将int64转换为字符串
	noStr := strconv.FormatInt(no, 10)
	return noPrefix + noStr, nil
}
