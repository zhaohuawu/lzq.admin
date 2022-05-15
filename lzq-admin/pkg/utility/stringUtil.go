package utility

import "strings"

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/15
 * @Version 1.0.0
 */

func UrlJoint(path ...string) string {
	var url strings.Builder
	for _, v := range path {
		if strings.HasSuffix(v, "/") {
			url.WriteString(v)
		} else {
			url.WriteString(v)
			url.WriteString("/")
		}
	}
	return strings.TrimSuffix(url.String(), "/")
}
