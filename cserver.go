package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type room struct {
	name string 
	Msgchan chan string
	clients map[chan<- string]struct{}
	*sync.Mutex
}

func (r *room) broadcastMsg(msg string) {
	r.Lock()
	defer r.Unlock()
	fmt.Println("Received message: ",msg)
	for wchan,_ := range r.clients {  // 1
		go func(wchan chan<- string) {
			wchan <- msg
		}(wchan)
	}
}

type client struct {
	*bufio.Reader
	*bufio.Writer
	wchan chan string
}

func (c *client) writeMonitor() {
	go func() {
		for s:= range c.wchan { //once a msg is ready in the client write ch, we'll write it to the underlying client network buffer
			c.WriteString(s+"\n") //these are mebedded functions from bufio/io packages so we can call them directly, the backslash n is to indicate End of Line. To write the content string from buffer to the writer object 
			c.Flush() //Flushes the data to the target resource from the writer
		}
	}()
 }

func StartClient(msgChan chan<- string, cn net.Conn) (chan<- string) {
	c:=new(client)
	c.Reader=bufio.NewReader(cn)
	c.Writer=bufio.NewWriter(cn)
	c.wchan=make(chan string)

	go func() {
		scanner:=bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			msgChan <- scanner.Text()
		}
	}()
	c.writeMonitor()

	return c.wchan
}

func (r *room) AddClient(c net.Conn) {
	r.Lock()
	wchan := StartClient(r.Msgchan,c)
	//define keys for the clients map
	r.clients[wchan]=struct{}{} // this will trigger the check for wc value in the for range in 1
	r.Unlock()
}

func CreateRoom(name string) *room {
	r:=&room{
		name: name,
		Msgchan: make(chan string), //initialize the msg ch
		Mutex: new(sync.Mutex),
		clients: make(map[chan<- string]struct{}),// initialize clients map or set
	}

	r.Run()
	return r
}

func (r *room) Run() {
	fmt.Println("Starting chat room",r.name)
	go func(){
		for msg:=range r.Msgchan{
			r.broadcastMsg(msg)
		}
	}()
}

func main(){
	r:=CreateRoom("Simple Chat Room")
	l, err:=net.Listen("tcp",":17000")
	if err!=nil{
		fmt.Println("Error connecting to chat client", err)
	}

	for {
		//the accept function will make the run function a blocking function as it waits for incoming connection and wait for the for loop to break
		conn, err:=l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection from chat client",err)
			break
		}

		//use go routine to handle connection so it doesn't affect or disturb the for loop
		go handleConnection(r,conn)
	}
}

func handleConnection(r * room, c net.Conn) {
	r.AddClient(c)
}

