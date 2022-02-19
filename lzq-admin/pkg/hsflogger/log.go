package hsflogger

/**
 * @Author  糊涂的老知青
 * @Date    2021/12/17
 * @Version 1.0.0
 */

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"lzq-admin/config"
	"lzq-admin/pkg/utility"
	"log"
	"os"
	"strings"
	"time"
)

var hsfLog = logrus.New()

func Init() {
	// 为当前logrus实例设置消息输出格式为json格式
	hsfLog.Formatter = &logrus.JSONFormatter{}
	// 设置日志级别为warn以上
	//hsfLog.SetLevel(logrus.DebugLevel)
	// 设置在输出日志中添加文件名和方法信息
	hsfLog.SetReportCaller(true)
	hook := initHook()
	hsfLog.AddHook(hook)
}

//esHook 自定义的ES hook
type hsfLogHook struct {
	cmd    string // 记录启动的命令
	client *elastic.Client
	osFile *os.File
}

//initHook 初始化
func initHook() *hsfLogHook {
	hook := &hsfLogHook{cmd: strings.Join(os.Args, " ")}
	if config.Config.LogConfig.UseElasticSearch {
		hook.client = initEs()
	} else {
		hook.osFile, _ = initFile()
		hsfLog.Out = hook.osFile
	}
	return hook
}

// initFile  初始化本地文件
func initFile() (*os.File, error) {
	rootPath, err := utility.GetRootPath()
	if err != nil {
		fmt.Println("get RootPath failed, err:", err)
		return nil, err
	}
	fmt.Println("rootPath", rootPath)
	today := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile(fmt.Sprintf("%v/logs/%v.log", rootPath, today), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return logFile, err
	}
	return logFile, nil
}

//InitEs 初始化ES
func initEs() *elastic.Client {
	es, err := elastic.NewClient(
		elastic.SetURL(config.Config.LogConfig.ElasticSearchUrl),
		elastic.SetBasicAuth(config.Config.LogConfig.ElasticSearchUser, config.Config.LogConfig.ElasticSearchPwd),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(15*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ES:", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "ES:", log.LstdFlags)),
	)

	if err != nil {
		log.Fatal("failed to create Elastic V6 Client: ", err)
	}
	return es
}

//Fire logrus hook interface 方法
func (hook *hsfLogHook) Fire(entry *logrus.Entry) error {
	doc := newLogEntity(entry)
	doc["cmd"] = hook.cmd
	hook.sendLog(doc)
	return nil
}

//Levels logrus hook interface 方法
func (hook *hsfLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//sendLog 异步发送日志到es
func (hook *hsfLogHook) sendLog(doc appLogDocModel) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("send entry to es failed: ", r)
		}
	}()
	if config.Config.LogConfig.UseElasticSearch {
		_, err := hook.client.Index().Index(doc.indexName()).Type("logEvent").BodyJson(doc).Do(context.Background())
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(doc)
	}
}

//appLogDocModel es model
type appLogDocModel map[string]interface{}

func newLogEntity(e *logrus.Entry) appLogDocModel {
	ins := map[string]interface{}{}
	for kk, vv := range e.Data {
		ins[kk] = vv
	}
	ins["@times"] = time.Now().Local()
	ins["level"] = e.Level
	ins["message"] = e.Message
	ins["caller"] = fmt.Sprintf("%s:%d  %#v", e.Caller.File, e.Caller.Line, e.Caller.Func)
	return ins
}

// indexName es index name 时间分割
func (m *appLogDocModel) indexName() string {
	return config.Config.LogConfig.ApplicationName + "-" + time.Now().Local().Format("2006-01")
}
