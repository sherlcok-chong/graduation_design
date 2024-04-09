package v1

import (
	"GraduationDesign/src/dao"
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/global"
	"GraduationDesign/src/logic"
	mid "GraduationDesign/src/middleware"
	"GraduationDesign/src/model"
	"GraduationDesign/src/model/chat"
	"GraduationDesign/src/myerr"
	"context"
	"encoding/json"
	"fmt"
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math"
	"net/http"
	"sync"
)

type ws struct {
}

// gin 不支持websocket，需升级http请求为webSocket协议
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mu sync.Mutex
var idCounter int64

//map记录所有ws客户端
var (
	clients      = make(map[int64]Client)
	userCli      = make(map[int64]int64)
	broadcastMsg = make(chan []byte, 100)
)

type Client struct {
	ID   int64
	Conn *websocket.Conn
}

func generateID() int64 {
	mu.Lock()
	defer mu.Unlock()
	if idCounter == math.MaxInt64 {
		idCounter = 0
	}
	idCounter++
	return idCounter
}

func (c *Client) sendClientID() error {
	msg := map[string]int64{"clientID": c.ID}
	msg["msg_type"] = 2
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	global.Logger.Info("link start")
	return c.Conn.WriteMessage(websocket.TextMessage, data)
}

func (c *Client) sendConnectMsg(m string) error {
	msg := map[string]string{"msg": m}
	//global.Logger.Error("err in check token")
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.Conn.WriteMessage(websocket.TextMessage, data)
}

//WebSocket handle
func (ws) WebSocket(c *gin.Context) {
	webs, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("websocket connect err")
		return
	}
	clientID := generateID()
	defer func() {
		//连接断开，删除无效client
		delete(clients, clientID)
		webs.Close()
	}()

	//记录每一个新连接
	mu.Lock()
	clients[clientID] = Client{
		ID:   clientID,
		Conn: webs,
	}
	mu.Unlock()
	fmt.Println("连接成功")
	cl := clients[clientID]
	cl.sendClientID()
	for {
		//读取ws中的数据
		_, msg, err := webs.ReadMessage()
		if err != nil {
			break
		}

		//加入广播消息
		broadcastMsg <- msg
	}
}

// Broadcast 广播(群发)
func (ws) Broadcast() {
	for {
		v, ok := <-broadcastMsg
		if !ok {
			break
		}
		msg := &chat.MsgSend{}
		if err := json.Unmarshal(v, msg); err != nil {
			global.Logger.Error(err.Error())
			continue
		}

		go func() {
			// 读/发送成功消息逻辑
			ctx := context.Background()
			if msg.IsRead == true && msg.FUserID == msg.TUserID {
				err := dao.Group.Mysql.ReadMessage(ctx, msg.ID)
				if err != nil {
					global.Logger.Error("err in read message")
				}
				return
			}

			clf := clients[msg.FClientID]
			var userID int64
			if msg.MsgType == 2 {
				global.Logger.Info("check sign")
				// 验签
				token, err := mid.MustAccount(msg.Text)
				if err != nil {
					clf.sendConnectMsg("check token fault")
					delete(clients, msg.FClientID)
					clf.Conn.Close()
					return
				}
				userID = token.Content.ID
				userCli[userID] = clf.ID
				clf.sendConnectMsg("check token success")
			} else if msg.MsgType == 3 {
				clf.sendConnectMsg("ok")
			} else {
				if exists, _ := dao.Group.Mysql.ExistsUserByID(ctx, msg.TUserID); !exists {
					global.Logger.Error("not find toId in mysql")
					clf.sendConnectMsg("not find toId in mysql")
					return
				}
				if err := clf.Conn.WriteMessage(websocket.TextMessage, v); err != nil {
					//发送失败，客户端异常（断线...）
					delete(clients, clf.ID)
				}
				clt, ok := clients[userCli[msg.TUserID]]
				if ok {
					if err := clt.Conn.WriteMessage(websocket.TextMessage, v); err != nil {
						//发送失败，客户端异常（断线...）
						delete(clients, clt.ID)
					} else {
						log := fmt.Sprintf("%v send msg to %v,data : %v", msg.FUserID, msg.TUserID, msg.CreateAt)
						global.Logger.Info(log)
					}
				}
				err := dao.Group.Mysql.CreateNewMessage(ctx, db.CreateNewMessageParams{
					Fid:      msg.FUserID,
					Tid:      msg.TUserID,
					IsFile:   msg.IsFile,
					IsRead:   false,
					Texts:    msg.Text,
					Createat: msg.CreateAt,
				})
				if err != nil {
					global.Logger.Error("add new message error")
					return
				}
			}

		}()
	}
}

func (ws) GetNotReadMsg(c *gin.Context) {
	rly := app.NewResponse(c)
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	data, err := logic.Group.WS.GetNotReadMsg(c, content.ID)
	rly.Reply(err, data)
}

func (ws) ReadAllMessage(c *gin.Context) {
	rly := app.NewResponse(c)
	req := &chat.FromUserID{}
	if err := c.ShouldBindQuery(req); err != nil {
		rly.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	content, ok := mid.GetTokenContent(c)
	if !ok || content.Type != model.UserToken {
		rly.Reply(myerr.AuthNotExist)
		return
	}
	err := logic.Group.WS.ReadUserMsg(c, req.UserID, content.ID)
	rly.Reply(err, nil)
}
