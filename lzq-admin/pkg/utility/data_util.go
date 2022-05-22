package utility

import "time"

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/22
 * @Version 1.0.0
 */

func TimeToDateStr(time time.Time) string {
	return time.Format("2006-01-02")
}

func TimeToDateTimeStr(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
