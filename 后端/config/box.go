package config

import "github.com/wike2019/wike_go/lib/utils"

// 定义一个泛型结构体
type Box[T any] struct {
	Content *utils.MapSync[T]
}
