package tools

import "log"

func MessageError(err error)  {
	if err!=nil {
		log.Println(err.Error())
		return
	}
}