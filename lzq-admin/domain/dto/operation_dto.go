package dto

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/3
 * @Version 1.0.0
 */

type OperationDto struct {
	Name   string `json:"name"`
	Text   string `json:"text"`
	Policy string `json:"policy"`
}

func GetOperationButton(name, text, policy string) OperationDto {
	return OperationDto{
		Name:   name,
		Text:   text,
		Policy: policy,
	}
}
