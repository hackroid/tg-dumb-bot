package utils

//
type BaseTimerTask interface {
	Init()
	GetMsg() string
}

func GetTimerTasksList() []BaseTimerTask {
	var list []BaseTimerTask
	// New timer task will be added here. (crawler...etc)
	list = append(list, GetWeiboCrawler())
	return list
}
