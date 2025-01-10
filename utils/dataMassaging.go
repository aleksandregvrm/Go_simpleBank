package util

import "fmt"

type MyDataMessaging struct {
	name            string
	surname         string
	dateOfBirthYear int16
}

var myname interface{} = "MyStringName"

func (dataMessaging *MyDataMessaging) NewMyDataMessaging() *MyDataMessaging {
	return &MyDataMessaging{
		name:            dataMessaging.name,
		surname:         dataMessaging.surname,
		dateOfBirthYear: dataMessaging.dateOfBirthYear,
	}
}

func (messaging *MyDataMessaging) Messaging(name string) {
	value, ok := myname.(string)
	if !ok {
		fmt.Println("this is not a string")
		return
	}
	fmt.Println(value)
	fmt.Println(messaging.name + name)
}
