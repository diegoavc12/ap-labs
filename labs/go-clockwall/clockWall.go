package main

import (
	"log"
	"net"
	"os"
	"strings"
	"bufio"
	"fmt"
)


func main() { 
	done := make(chan int)
	for i:=1; i<len(os.Args); i++{
		
		arr:=strings.Split(os.Args[i], "=")
		conn, err := net.Dial("tcp", arr[1])
    	if err != nil {
        	log.Fatal(err)
    	}

		buff:=bufio.NewReader(conn)
    	
    	go func() {
        	//io.Copy(os.Stdout, conn) // NOTE: ignoring errors
        	//log.Println("done")
			for{
				str, err:=buff.ReadString('\n')
				
				if err != nil {
					log.Fatal(err)
				}
				str=strings.ReplaceAll(str,"\n","")
				for j:=0; j<i; j++{
					fmt.Println();
				}		
				fmt.Printf("%s",str)
		
			
				
				
			}
        	done <- 2 // signal the main goroutine
    	}()


    	
	}
	x := 1
    x = <-done // wait for background goroutine to finish
    log.Println("Channel Closed with value: ", x)
    close(done)
	close(ch)
}