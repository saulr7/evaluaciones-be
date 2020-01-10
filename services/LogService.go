package services

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func LogToFile(Titulo string, x interface{}) {

	currentTime := time.Now()

	f, err := os.OpenFile(currentTime.Format("2006-01-02")+".log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, Titulo+" ", log.LstdFlags)

	out, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	logger.Println(string(out))
}
