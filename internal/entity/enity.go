package entity

type TextKey struct {
	Text string `json:"text" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

type KeyValue struct {
	Key   string `json:"key" binding:"required"`
	Value int64  `json:"value" binding:"required"`
}

type User struct {
	Name string `json:"name" binding:"required"`
	Age  uint8  `json:"age" binding:"required"`
}
