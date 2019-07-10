package qilog

import "github.com/op/go-logging"


func Errorf(format string, args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Errorf(CallPath(2) + format, args...)
			return true
		}
	}
	return false
}

func Error(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Error(CallPath(2), args)
			return true
		}
	}
	return false
}

func Fatalf(format string, args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Fatalf(CallPath(2) + format, args...)
			return true
		}
	}
	return false
}

func Fatal(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Fatal(CallPath(2), args)
			return true
		}
	}
	return false
}

func Panicf(format string, args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Panicf(CallPath(2) + format, args...)
			return true
		}
	}
	return false
}

func Panic(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Panic(CallPath(2), args)
			return true
		}
	}
	return false
}

func Warningf(format string, args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Warningf(CallPath(2) + format, args...)
			return true
		}
	}
	return false
}

func Warning(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Warning(CallPath(2), args)
			return true
		}
	}
	return false
}

func Criticalf(format string, args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Criticalf(CallPath(2) + format, args...)
			return true
		}
	}
	return false
}

func Critical(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			Logger.Critical(CallPath(2), args)
			return true
		}
	}
	return false
}

func Check(args ...interface{}) bool{
	for _, arg := range args {
		if arg != nil{
			return true
		}
	}
	return false
}


func Infof(format string, args ...interface{}){
	Logger.Infof(CallPath(2) + format, args...)
}

func Info(args ...interface{}){
	Logger.Info(CallPath(2), args)
}

func Debugf(format string, args ...interface{}){
	Logger.Debugf(CallPath(2) + format, args...)
}

func Debug(args ...interface{}){
	Logger.Debug(CallPath(2), args)
}

func Monitorf(format string, args ...interface{}){
	if Monitoring == true{
		Monitorlog.Printf(CallPath(2) + format, args...)
	}
}

func Monitor(args ...interface{}){
	if Monitoring == true{
		Monitorlog.Println(CallPath(2), args)
	}
}

func SetLogLevel(level string){
	l, _ := logging.LogLevel(level)
	logging.SetLevel(l, Logger.Module)
}