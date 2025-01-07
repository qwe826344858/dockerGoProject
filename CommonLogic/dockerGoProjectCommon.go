package CommonLogic

import (
	"fmt"
	"log"
	"os"
)


func LoggerInit(sName string){
	env := os.Getenv("SYSTEM_ENV")
	strLogDir := ""
	if env == "docker"{
		strLogDir = os.Getenv("GO_LOG_DIR")
	}else{
		strLogDir = "/home/lighthouse/LogInfo/DockerGoLog"
	}

	sFilePath := fmt.Sprintf("%s/Service_%s.log",strLogDir,sName)
	log.Printf("env:%v strLogDir:%v sFilePath:%v",env,strLogDir,sFilePath)
	logFile, err := os.OpenFile(sFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
		return
	}
	log.SetOutput(logFile)
}