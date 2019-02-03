package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UNIT int64

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	LOG int = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

const (
	OS_LINUX = iota
	OS_X
	OS_WIN
	OS_OTHERS
)

type _FILE struct {
	dir      string
	filename string
	mu       *sync.RWMutex
	logfile  *os.File
	json     bool
	lg       *log.Logger
	_suffix  int
	_date    *time.Time
}

var logObj *_FILE
var logLevel int = 1
var maxFileSize int64
var maxFileCount int32
var RollingFile bool = false
var dailyRolling bool = true
var consoleAppender bool = false

const DATEFORMAT = "2006-01-02"

var logFormat string = "%s %s"
var consoleFormat string = "%s %s:%d %s %s"

func SetConsole(isConsole bool) {
	consoleAppender = isConsole
}

func SetLevel(_level int) {
	logLevel = _level
}

func SetRollingFile(fileDir, fileName string, maxNumber int32, maxSize int64, _unit UNIT) {
	maxFileCount = maxNumber
	maxFileSize = maxSize * int64(_unit)
	RollingFile = true
	dailyRolling = false
	logObj = &_FILE{dir: fileDir, filename: fileName, mu: new(sync.RWMutex)}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()
	for i := 1; i <= int(maxNumber); i++ {
		if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
			logObj._suffix = i
		} else {
			break
		}
	}
	if !logObj.isMustRename() {
		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		logObj.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logObj.rename()
	}
	go fileMonitor()
}

func SetRollingDaily(fileDir, fileName string) {
	RollingFile = false
	dailyRolling = true
	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
	logObj = &_FILE{dir: fileDir, filename: fileName, _date: &t, mu: new(sync.RWMutex)}
	logObj.mu.Lock()
	defer logObj.mu.Unlock()

	if !logObj.isMustRename() {
		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		logObj.lg = log.New(logObj.logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logObj.rename()
	}
}

func GetTraceLevelName(level int) string {
	switch level {
	case LOG:
		return "LOG"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func GetOsFlag() int {
	switch os := runtime.GOOS; os {
	case "darwin":
		return OS_X
	case "linux":
		return OS_LINUX
	case "windows":
		return OS_WIN
	default:
		return OS_OTHERS
	}
}

func GetOsEol() string {
	if GetOsFlag() == OS_WIN {
		return "\r\n"
	}
	return "\n"
}

func Concat(delimiter string, input ...interface{}) string {
	buffer := bytes.Buffer{}
	l := len(input)
	for i := 0; i < l; i++ {
		buffer.WriteString(fmt.Sprint(input[i]))
		if i < l-1 {
			buffer.WriteString(delimiter)
		}
	}
	return buffer.String()
}

func console(msg string) {
	if consoleAppender {
		log.Print(msg)
	}
}

var curStackFlag bool = false
var curStackPath string
var curStackLine int

func markCurStack() {
	if !curStackFlag {
		curStackFlag = true
		_, curStackPath, curStackLine, _ = runtime.Caller(0)
	}
}

func getCurStackPath() string {
	if !curStackFlag {
		markCurStack()
	}
	return curStackPath
}

func getStack(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}

func detectStack() (string, int) {
	curPath := getCurStackPath()
	for skip := 0; ; skip++ {
		_, path, line, ok := runtime.Caller(skip)
		if path != curPath {
			return path, line
		}
		if !ok {
			break
		}
	}
	return "", 0
}

func splitDirFile(path string) (string, string) {
	return filepath.Dir(path), filepath.Base(path)
}

func getTraceDirInfo(dir string) string {
	if GetOsFlag() == OS_WIN {
		split := strings.Split(dir, "\\")
		if len(split) > 2 {
			return split[0] + "\\" + split[1] + "\\...\\" + split[len(split)-1] + "\\"
		} else {
			return dir + "\\"
		}
	} else {
		split := strings.Split(dir, "/")
		if len(split) > 2 {
			return split[0] + "/.../" + split[len(split)-1]
		} else {
			return dir + "/"
		}
	}
}

func getTraceFileLine() (string, int) {
	path, line := detectStack()
	dir, file := splitDirFile(path)
	dir = getTraceDirInfo(dir)
	return file, line
}

func buildConsoleMessage(level int, msg string) string {
	file, line := getTraceFileLine()
	return fmt.Sprintf(consoleFormat+GetOsEol(), time.Now().Format("2006/01/02 15:04:05"), file, line, GetTraceLevelName(level), msg)
}

func buildLogMessage(level int, msg string) string {
	return fmt.Sprintf(logFormat+GetOsEol(), GetTraceLevelName(level), msg)
}

func catchError() {
	if err := recover(); err != nil {
		log.Println("err", err)
	}
}

func Trace(level int, v ...interface{}) bool {
	if dailyRolling {
		fileCheck()
	}
	defer catchError()
	logObj.mu.RLock()
	defer logObj.mu.RUnlock()
	// if level == ERROR || level == INFO {
	// 	if len(v) >= 4 && v[0] != nil && v[0].(string) == "remote" {
	// 		go httpLog(level, v[1].(string), v[2].(string), v[3].(string))
	// 	}
	// }
	msg := Concat(" ", v...)

	logStr := buildConsoleMessage(level, msg)
	console(logStr)
	if level >= logLevel {
		logObj.lg.Output(3, msg)
	}
	return true
}

func Log(v ...interface{}) bool {
	return Trace(LOG, v...)
}

func Debug(v ...interface{}) bool {
	return Trace(DEBUG, v...)
}

func Info(v ...interface{}) bool {
	return Trace(INFO, v...)
}

func Warn(v ...interface{}) bool {
	return Trace(WARN, v...)
}

func Error(v ...interface{}) bool {
	return Trace(ERROR, v...)
}

func Fatal(v ...interface{}) bool {
	return Trace(FATAL, v...)
}

func (f *_FILE) isMustRename() bool {
	if dailyRolling {
		t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
		if t.After(*f._date) {
			return true
		}
	} else {
		if maxFileCount > 1 {
			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
				return true
			}
		}
	}
	return false
}

func (f *_FILE) rename() {
	if dailyRolling {
		fn := f.dir + "/" + f.filename + "." + f._date.Format(DATEFORMAT)
		if !isExist(fn) && f.isMustRename() {
			if f.logfile != nil {
				f.logfile.Close()
			}
			err := os.Rename(f.dir+"/"+f.filename, fn)
			if err != nil {
				f.lg.Println("rename err", err.Error())
			}
			t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
			f._date = &t
			f.logfile, _ = os.Create(f.dir + "/" + f.filename)
			f.lg = log.New(logObj.logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
		}
	} else {
		f.coverNextOne()
	}
}

func (f *_FILE) nextSuffix() int {
	return int(f._suffix%int(maxFileCount) + 1)
}

func (f *_FILE) coverNextOne() {
	f._suffix = f.nextSuffix()
	if f.logfile != nil {
		f.logfile.Close()
	}
	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix))) {
		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix)))
	}
	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(int(f._suffix)))
	f.logfile, _ = os.Create(f.dir + "/" + f.filename)
	f.lg = log.New(logObj.logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
}

func fileSize(file string) int64 {
	f, e := os.Stat(file)
	if e != nil {
		fmt.Println(e.Error())
		return 0
	}
	return f.Size()
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func fileMonitor() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			fileCheck()
		}
	}
}

func fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if logObj != nil && logObj.isMustRename() {
		logObj.mu.Lock()
		defer logObj.mu.Unlock()
		logObj.rename()
	}
}

/*
func flog(i int) {
	Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
	Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//指定是否控制台打印，默认为true
	SetConsole(true)
	//指定日志文件备份方式为文件大小的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//第三个参数为备份文件最大数量
	//第四个参数为备份文件大小
	//第五个参数为文件大小的单位
	SetRollingFile("logtest", "test.log", 10, 5, KB)

	//指定日志文件备份方式为日期的方式
	//第一个参数为日志文件存放目录
	//第二个参数为日志文件命名
	//logger.SetRollingDaily("logtest", "test.log")

	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	//一般习惯是测试阶段为debug，生成环境为info以上
	SetLevel(ERROR)
	flog(1)

	for i := 10000; i > 0; i-- {
		go flog(i)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)

}
*/
