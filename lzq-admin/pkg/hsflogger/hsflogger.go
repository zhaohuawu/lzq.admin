package hsflogger

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"github.com/sirupsen/logrus"
)

var Fields logrus.Fields

func logWithField(objs ...interface{}) *logrus.Entry {
	if len(objs) > 0 {
		//for _,v:=range objs{
		//	fields[v]=v
		//}

	}
	return hsfLog.WithFields(Fields)
}

func LogInformation(msg string, obj ...interface{}) {
	logWithField(obj).Info(msg)
}

func LogError(msg string, err error, obj ...interface{}) {
	logWithField(obj).WithField("err", err).Error(msg)
}

func LogDebug(msg string, obj ...interface{}) {
	logWithField(obj).Debug(msg)
}
