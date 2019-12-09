package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func creatLog() {
	fp, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	myPath := "/go/src/github.com/martinrocket/nature/"
	logPath := "logs/"

	folderPath := fp + myPath + logPath
	dt := time.Now()
	logName := dt.Format("20060102_") + "application.log"

	startLog(folderPath, logName)
	log.Println("Log Startup... ", time.Now().String())

}

func startLog(p, l string) {
	fmt.Println(p + " " + l)

	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	mkLog := filepath.Join(p, l)
	file, err := os.OpenFile(mkLog, os.O_RDWR|os.O_APPEND, 0660)
	if os.Stat(mkLog); os.IsNotExist(err) {
		file, err = os.Create(mkLog)

	} else {
		file, err = os.OpenFile(mkLog, os.O_RDWR|os.O_APPEND, 0660)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}
	}
	log.SetOutput(file)
	log.Println("Log created at:", mkLog)

}
