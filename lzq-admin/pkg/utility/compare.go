package utility

import "reflect"

/**
 * @Author  糊涂的老知青
 * @Date    2022/2/13
 * @Version 1.0.0
 */

func IsEmptyOfStruct(v interface{}, vEmpty interface{}) bool {
	return reflect.DeepEqual(v, vEmpty)
}
