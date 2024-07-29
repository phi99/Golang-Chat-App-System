..***In Progress**

	 	 	   ----------------------------------------------------------------
	     	    	     ** Chat System - Golang **
	  	 	   ----------------------------------------------------------------

```text
Chat App Architecture/Traffic Flow Overview
-------------------------------------------------
The chat systems utilizes client-server architecture, TCP/IP communication protocol, Golang concurrency features, and implemented using pipeline pattern design which is basically a series of stages connected via channels, where each stage is groups of goroutines running similar function.

client1 (connects to server to join chat room, and is allocated a channel - wchan1) -> TCP connection1 -> [r.MsgChan]-> wchan1 -> TCP connection1 -> client1
                                                                                
client2 (connects to server to join chat room, and is allocated a channel - wchan2) -> TCP connection2 -> [r.MsgChan]-> wchan1 -> TCP connection2 -> client2                                                                                           
```

```text
Functionality/Mechanism
------------------------
When client connects to server to join chat room via TCP/IP, it would be allocated a channel at the server (wchan) which serve as a pipeline that connects the message that the server would like to broadcast to the client in the chat room and the client connection. The server also create a channel MsgCh to receive the contents of the message from the client's TCP connection
```

```text
Usage
------------------------
Run the server and then run several clients, it would then automatically join the chat room and prompt the client to enter their name, after typing client name, when each client send message to the chat room, all the client in the chat room will receive the message 
```
