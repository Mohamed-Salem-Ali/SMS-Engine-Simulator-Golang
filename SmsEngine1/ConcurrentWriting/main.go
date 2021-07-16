package main

import (
	"log"
	"os"
	"time"
)

func process(file *os.File, s []string, out chan string) {
	defer close(out)
	for i := 0; i < len(s); i++ {
		out <- (s[i])
	}
}
func main() {
	names := []string{"salem", "sasa", "lala", "bala", "messi", "ronaldo"}
	start := time.Now()
	out1 := make(chan string)
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("we got error", err)
	}
	go process(file, names, out1)
	log.SetOutput(file)
	for msg := range out1 {
		//time.Sleep(time.Second / 2)
		log.Println(msg)
	}
	end := time.Since(start)
	log.Println("Execution Time::", end)
}
