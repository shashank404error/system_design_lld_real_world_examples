package main

import "fmt"

type server struct {
	id     string
	weight int
}

type lb struct {
	servers     []*server
	serverQueue []*server
}

func (d *lb) initServerQueue() {
	var serverQueue []*server
	for _, s := range d.servers {
		var serverQueueIdentical []*server
		for i := 0; i < s.weight; i++ {
			serverQueueIdentical = append(serverQueueIdentical, &server{id: s.id})
		}
		serverQueue = append(serverQueue, serverQueueIdentical...)
	}
	d.serverQueue = serverQueue
}

func (d *lb) nextServer() *server {
	nextServer := d.serverQueue[0]
	d.serverQueue = append(d.serverQueue[1:len(d.serverQueue)], nextServer)
	return nextServer
}

func main() {
	loadBalancer := &lb{
		servers: []*server{
			{
				id:     "server_1",
				weight: 5,
			},
			{
				id:     "server_2",
				weight: 3,
			},
			{
				id:     "server_3",
				weight: 2,
			},
		},
	}
	loadBalancer.initServerQueue()

	for i := 0; i < 20; i++ {
		nextServer := loadBalancer.nextServer()
		fmt.Printf("Request id: %d will be served by %s\n", i, nextServer.id)
	}

}
