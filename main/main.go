package main

func main() {
	bs := new(BotServer)
	bs.initBot()
	bs.addMessageHandler(MSG_TYPE_TEXT, normalTextMessage)
	bs.addMessageHandler(MSG_TYPE_CMD, normalCommandMessage)
	bs.pollingChannelUpdates()
}
