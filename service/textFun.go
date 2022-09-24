package service

import (
	"fmt"
	"go-filedrop/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Get(c *gin.Context) {
	c.JSON(200, "ok")
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	global.WSConn = append(global.WSConn, ws)
	defer func() {
		global.WSConn = DeleteSlice(ws, global.WSConn)
		ws.Close()
	}()

	ws.WriteMessage(websocket.TextMessage, global.TextTemp)
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 保存到全局变量
		global.TextTemp = message
		fmt.Println("conn num", len(global.WSConn))
		for _, wss := range global.WSConn {
			//写入ws数据
			err = wss.WriteMessage(mt, global.TextTemp)
			if err != nil {
				break
			}
		}

	}
}

func DeleteSlice(conn *websocket.Conn, a []*websocket.Conn) []*websocket.Conn {
	for i := 0; i < len(a); i++ {
		if a[i] == conn {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}
