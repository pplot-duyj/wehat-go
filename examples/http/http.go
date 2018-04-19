package main

import (
	"fmt"
	"net/http"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"strings"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	//配置微信参数
	config := &wechat.Config{
		AppID:     "wxf2e00ca264aa374a",
		AppSecret: "ecdcfd4bbd98894b829b6c326e34078f",
		Token:     "8_esvvoAMa-IGRPJgbnrpROlSaqmpLIGq_nA4GbgiQp6X7AxpTY6faEqfxk_cdao9MXlAKSQKC_7rFFwZNg6LMfKqKT4jnrW-EiPIKRP4TZTR10mJsVM3IS-6C31X5rdG3WNarctbHmxKs7Ep_LOGgAAACFX",

		EncodingAESKey: "",
	}

	fmt.Print(req.PostForm)
	for v := range req.PostForm {
		fmt.Print("==" + v + ":==" + req.Form[v][0])
	}

	if strings.EqualFold("GET", req.Method) {
		req.ParseForm()
		if len(req.Form["echostr"]) > 0 {
			echostr := req.Form["echostr"][0]
			fmt.Println("echostr:" + echostr)
			rw.Write([]byte(echostr))
			return
		}
		return
	}

	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		fmt.Print(msg)
		fmt.Print("msg")
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}
