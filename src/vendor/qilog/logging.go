package qilog
import(
	"github.com/op/go-logging"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func init(){
	SetMode()
}

var Logger *logging.Logger
var LogDetails LoggerType
var Mode string
var	Monitoring	bool
var Monitorlog *log.Logger

type LoggerType struct{
	Module string
	Level string
	Prefix	string
	Console bool
	FileName string
	Backend []logging.Backend
	FormatString string
	Format  logging.Formatter
	LeveledBackend logging.LeveledBackend
}

func Init(){
	ReadJSON("log.json")
	SetModule()
	SetOutput()
	SetFormatter()
	LevelBackend()
	SetLog()
	SetLevel()
	SetMonitor()
}

func SetMode(){
	Mode = "Production"
	Monitoring = false
	Init()
}

func ReadJSON(fileName string){
	fetchJSON := make(map[string]LoggerType) 
	data, err := ioutil.ReadFile(fileName)
	if err != nil{
		log.Fatal(err)
	}	

	err = json.Unmarshal([]byte(data), &fetchJSON)
	if err != nil{
		log.Fatal(err)
	}
	var OK bool
	LogDetails, OK = fetchJSON[Mode]
	if !OK {
		os.Stderr.Write([]byte("Invalid Flag"))
		os.Exit(2)
	}
}

func SetModule(){
	Logger = logging.MustGetLogger(LogDetails.Module)
}

//sets output for logging
func SetOutput(){
	if LogDetails.Console {
		LogDetails.Backend = append(LogDetails.Backend, logging.NewLogBackend(os.Stderr, "", 0))		
	}

	file, _ := os.OpenFile(LogDetails.FileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	LogDetails.Backend = append(LogDetails.Backend, logging.NewLogBackend(file, "", 0))
}

func SetFormatter(){
	LogDetails.Format = logging.MustStringFormatter(LogDetails.FormatString)

	LogDetails.Backend[0] = logging.NewBackendFormatter(LogDetails.Backend[0], LogDetails.Format)
	LogDetails.Backend[1] = logging.NewBackendFormatter(LogDetails.Backend[1], LogDetails.Format)
}

func LevelBackend(){
	LogDetails.LeveledBackend = logging.MultiLogger(LogDetails.Backend[0], LogDetails.Backend[1])
}

func SetLog(){
	Logger.SetBackend(LogDetails.LeveledBackend)
}

func SetLevel(){
	level, _ := logging.LogLevel(LogDetails.Level)
	logging.SetLevel(level, Logger.Module)
}

func SetMonitor(){
	file, _ := os.OpenFile(LogDetails.FileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)	
	Monitorlog = log.New(file, "Monitor: ", log.LstdFlags)
}


