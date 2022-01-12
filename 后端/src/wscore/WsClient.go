package wscore

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WsClient struct {
	conn *websocket.Conn
	readChan chan *WsMessage  //读队列 (chan)
	closeChan chan byte  // 失败队列
	mu *sync.Mutex

}
func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn,readChan:make(chan *WsMessage),closeChan:make(chan byte),mu: &sync.Mutex{}}
}
func(this *WsClient) Ping(waittime time.Duration){
	defer this.mu.Unlock()
	for {
		this.mu.Lock()
		time.Sleep(waittime)
		err:=this.conn.WriteMessage(websocket.TextMessage,[]byte("ping"))
		if err!=nil{
			ClientMap.Remove(this.conn)
			return
		}
		this.mu.Unlock()
	}
}
func(this *WsClient) ReadLoop(){
	for{
		t,data,err:=this.conn.ReadMessage()
		if err!=nil{
			this.conn.Close()
			ClientMap.Remove(this.conn)
			this.closeChan<-1
			break
		}
		this.readChan<-NewWsMessage(t,data)
	}
}



