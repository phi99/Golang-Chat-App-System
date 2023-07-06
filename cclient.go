package main
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var name string
	fmt.Println("Starting ChatClient...")
	fmt.Println("Enter name?")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s, connecting to chat system... \n", name)
	conn, err := net.Dial("tcp",":17000")
	if err!=nil{
		log.Fatal("Could not connect to chat system", err)
	}
	fmt.Println("Connected to chat system")

	name +=":"
	defer conn.Close()

	go func() {
		scanner:=bufio.NewScanner(conn)
		for scanner.Scan(){
			fmt.Println(scanner.Text())
		}
	}()

	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg:=scanner.Text()
		_,err=fmt.Fprintf(conn, name+msg+"\n")		
	}
}
