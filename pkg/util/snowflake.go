package util

import "github.com/bwmarrin/snowflake"

// 使用雪花算法产生唯一id
func GenerateUniqueId(roomid uint) (string, error) {
	node, err := snowflake.NewNode(int64(roomid))
	if err != nil {
		return "", err
	}
	id := node.Generate()
	return id.String(), nil
}
