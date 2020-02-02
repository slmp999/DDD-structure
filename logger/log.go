package logging

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	logs "github.com/sirupsen/logrus"

	logss "log"
	//"google.golang.org/api/option"
	"os"
)

var log = logs.New()

var now string

func init() {
	// strArray := [5]string{"Line", "time", "z-msg", "func"}
	log.Out = os.Stderr
	log.Level = logs.InfoLevel
	log.Formatter = &logs.TextFormatter{
		DisableColors:             false,
		DisableTimestamp:          true,
		ForceColors:               false,
		DisableLevelTruncation:    false,
		EnvironmentOverrideColors: false,
		DisableSorting:            false,
		QuoteEmptyFields:          true,
	}

	// log. = &logs.Logger{
	// 	Out:   os.Stderr,
	// 	Level: logs.InfoLevel,
	// 	Formatter: &logs.TextFormatter{
	// 		FieldMap: logs.FieldMap{
	// 			"time": now,
	// 		},
	// 		DisableColors:             true,
	// 		DisableTimestamp:          true,
	// 		ForceColors:               true,
	// 		DisableLevelTruncation:    false,
	// 		EnvironmentOverrideColors: true,
	// 	},
	// }
	// log.SetFormatter(logs.WithFields(logs.Fields{"time": now}))

}

// Print logs a message at level Info on the standard log.
func Print(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "info", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project}).Print(fmt.Sprintln(args...))
}

// Panic logs a message at level Panic on the standard log.
func Panic(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Panic("")
}

// Fatal logs a message at level Fatal on the standard log.
func Fatal(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "fatal", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Fatal("")
}

func Warn(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "panic", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"warn": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Warn("")
}

// Warning logs a message at level Warn on the standard log.
func Warning(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "warn", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Warning("")
}

func sendLine(message string, token string) (string, error) {
	client := &http.Client{}
	url := "https://notify-api.line.me/api/notify"
	payload := strings.NewReader("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"message\"\r\n\r\n" + message + "\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--")
	cReq, err := http.NewRequest("POST", url, payload)
	cReq.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	// cReq.Header.Add("authorization", "Bearer djp2CpuGYUb8slBPE5YFxMQ4s1u5u4DU0Fpoe25gbZS")
	cReq.Header.Add("authorization", "Bearer "+token)
	cReq.Header.Add("cache-control", "no-cache")
	response, err := client.Do(cReq)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	var returnBody struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(body, &returnBody)
	if err != nil {
		log.Error("send line notify error : ", err)
	}
	if returnBody.Status != 200 {
		log.Error("send line notify error : ", returnBody)
	}
	return returnBody.Message, nil
}

func Error(args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project}).Error(fmt.Sprint(args...))

	message := fmt.Sprintf(
		"ERROR : \nวันที่ %v เวลา %v \n func: %v \n line: %v \n message: %v",
		time.Now().In(loc).Format("2006-01-02"),
		time.Now().In(loc).Format("15:04:05"),
		project,
		fn+":"+strconv.Itoa(line),
		fmt.Sprint(args...),
	)
	go sendLine(message, "V7QQFIjZ1VEAKcRcTC8TlfDttv4S85rZUwyhetn0zHH")
}

func Infof(format string, args ...interface{}) {

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Infof("")
}

func Infoln(args ...interface{}) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "info", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Infoln("")
}

// Printf logs a message at level Info on the standard log.
func Printf(format string, args ...interface{}) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Printf("")

}

// Panicf logs a message at level Panic on the standard log.
func Panicf(format string, args ...interface{}) {
	// _, fn, line, _ := runtime.Caller(1)
	// payload := ":[" + now + "] " + fn + ":" + strconv.Itoa(line) + ": " + getHostName() + " : " + fmt.Sprintf(format, args...)
	// log.Panicf(format, payload)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-msg": fmt.Sprint(args...)}).Panicf("")
}

// Fatalf logs a message at level Fatal on the standard log.
func Fatalf(format string, args ...interface{}) {
	// _, fn, line, _ := runtime.Caller(1)

	// payload := fn + ":[" + now + "] " + fn + ":" + strconv.Itoa(line) + ": " + getHostName() + " : " + fmt.Sprintf(format, args...)
	// log.Fatalf(format, payload)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	// payload := fn + ":" + strconv.Itoa(line) + ": " + getHostName() + " : " + fmt.Sprint(args...)
	// log.Println(payload)

	log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project}).Fatalf(format, fmt.Sprint(args...))
}

func GetFillName(fillname string) string {
	// dir, _ := filepath.Abs(filepath.Dir(fillname))
	// prefixPath := fillname[:len(fillname)-len("/main.go")]
	// var ss []string
	// if runtime.GOOS == "windows" {
	// 	ss = strings.Split(prefixPath, "/")
	// } else {
	// 	ss = strings.Split(prefixPath, "/")
	parts := strings.Split(fillname, ".")

	//log.Println(parts)
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	// }
	return packageName
}

// Println logs a message at level Info on the standard log.
func Println(args ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	// payload := fn + ":" + strconv.Itoa(line) + ": " + getHostName() + " : " + fmt.Sprint(args...)
	// log.Println(payload)

	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project}).Println(fmt.Sprintln(args...))
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "info", fn+":"+strconv.Itoa(line), project)

}
func Log(l *logss.Logger, msg string, level string, line string, funcs string) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	l.SetPrefix(`level=` + level + ` time="` + time.Now().In(loc).Format("2006-01-02 15:04:05") + `" func="` + funcs + `" line="` + line + `" `)
	switch level {
	case "info":
		l.Print(msg)
	case "warn":
		l.Print(msg)
	case "debug":
		l.Print(msg)
	case "warning":
		l.Print(msg)
	case "fatal":
		l.Print(msg)
	default:
		l.Print(msg)
		// freebsd, openbsd,
		// plan9, windows...
	}
}

// Panicln logs a message at level Panic on the standard log.
func Panicln(args ...interface{}) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	// payload := fn + ":" + strconv.Itoa(line) + ": " + getHostName() + " : " + fmt.Sprint(args...)
	// log.Println(payload)
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "panic", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-meg": fmt.Sprint(args...)}).Panicln("")
}

// Fatalln logs a message at level Fatal on the standard log.
func Fatalln(args ...interface{}) {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	pc, fn, line, _ := runtime.Caller(1)
	fuc := runtime.FuncForPC(pc)
	project := GetFillName(fuc.Name())
	l := logss.New(os.Stdout, "", 0)
	Log(l, fmt.Sprint(args...), "fatal", fn+":"+strconv.Itoa(line), project)
	// log.WithFields(logs.Fields{"time": now, "Line": fn + ":" + strconv.Itoa(line), "func": project, "z-meg": fmt.Sprint(args...)}).Fatalln("")
}

func getHostName() string {
	host, _ := os.Hostname()
	return host
}
