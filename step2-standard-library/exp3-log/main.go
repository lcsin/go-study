package main

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	// 设置log日志的前缀
	log.SetPrefix("[Golang]: ")

	// 自定义错误输出到错误日志文件
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm|os.ModeAppend)
	if err != nil {
		log.Fatalln("failed open err file,err: ", err)
	}

	// 将日志输出到文件
	//log.SetOutput(file)

	// 设置log的标识
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// 自定义log
	Trace = log.New(ioutil.Discard, "DIY LOG TRACE: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	Info = log.New(os.Stdout, "DIY LOG INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	Warning = log.New(os.Stdout, "DIY LOG WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	Error = log.New(file, "DIY LOG ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}

func main() {
	/*
		log包的三种输出方式：
			1. log.Print()系列：一般信息
			2. log.Fatal()系列：重大错误
			3. log.Panic()系列：恐慌
	*/
	log.Print("Print: hello,world!")
	log.Printf("Printf: hello,world!")
	log.Println("Println: hello,world!")

	log.Fatal("Fatal: hello,world!")
	log.Fatalf("Fatalf: hello,world!")
	log.Fatalln("Fatalln: hello,world!")

	log.Panic("Panic: hello,world!")
	log.Panicf("Panicf: hello,world!")
	log.Panicln("Panicln: hello,world!")

	// 使用自定义log输出
	Trace.Println("Trace...")
	Info.Println("Info...")
	Warning.Println("Warning...")
	Error.Println("Error...")
}
