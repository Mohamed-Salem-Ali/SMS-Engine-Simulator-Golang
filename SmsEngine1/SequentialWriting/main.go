package main

import (
	"log"
	"os"
	"time"
)

func AddToFile(f string, s []string) {
	file, err := os.OpenFile(f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("we got error", err)
	}
	log.SetOutput(file)
	for i := 0; i < len(s); i++ {
		//time.Sleep(time.Second / 2)
		log.Println(s[i])
	}
}
func main() {
	names := []string{"salem", "sasa", "lala", "bala", "messi", "ronaldo"}
	start := time.Now()
	AddToFile("file.txt", names)
	end := time.Since(start)
	log.Println("Execution Time::", end)
}
