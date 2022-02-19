package utility

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/1
 * @Version 1.0.0
 */

import uuid "github.com/satori/go.uuid"

func UuidCreate() string {
	return uuid.NewV4().String()
}

func StringToUuid(str string) (uuid.UUID, error) {
	uuid, err := uuid.FromString(str)
	return uuid, err
}
