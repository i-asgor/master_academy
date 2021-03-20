package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	
	//\x00smbiz.temp@gmail.com\x00Tem@2020
	resp := []byte("\x00" + "asgor.ice@gmail.com" + "\x00" + "asgor@ice#01735586261")
	//fmt.Println(string(resp),resp)
	
	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}