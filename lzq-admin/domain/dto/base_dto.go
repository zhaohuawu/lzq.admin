package dto

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/05
 * @Version 1.0.0
 */

type BaseDto struct {
	ID string `json:"id" label:"主键"`
}

type KeyAndValueDto struct {
	Key   string `json:"key" label:"编码"`
	Value string `json:"value" label:"值"`
}
