package utility

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"fmt"
	"os"
)

// GetRootPath 获取项目根目录
func GetRootPath() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		fmt.Printf("path err %v", err)
	}

	//if strings.Contains(path, "hsf.basic") {
	//	path = strings.Split(path, "hsf.basic")[0] + "hsf.basic"
	//}

	return
}
