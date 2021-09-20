package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func  Authorizing(botID string){
	client := &http.Client{}
	telegram_id :=botID
	var values map[string]string
	values = make(map[string]string)
	values["Accept"]="application/json"
	values["Content-Type"]="application/json"
	urls :="https://api.telegram.org/" +"bot" + telegram_id +"/getMe"
	Url,err :=url.Parse(urls)
	if err != nil {
		return
	}
	urlPath := Url.String()
	req,_ := http.NewRequest("GET",urlPath,nil)
	for key,value := range values {
		req.Header.Add(key,value)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}
func getUpdates(botID string){
	client := &http.Client{}
	telegram_id :=botID
	var values map[string]string
	values = make(map[string]string)
	values["Accept"]="application/json"
	values["Content-Type"]="application/json"
	urls :="https://api.telegram.org/" +"bot" + telegram_id +"/getUpdates"
	Url,err :=url.Parse(urls)
	if err != nil {
		return
	}
	urlPath := Url.String()
	req,_ := http.NewRequest("GET",urlPath,nil)
	for key,value := range values {
		req.Header.Add(key,value)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func SendMessage(message, botID string, chat_id int64)(s string,err error){
	client := &http.Client{}
	telegram_id :=botID
	var values map[string]string
	values = make(map[string]string)
	var values1 map[string]interface{}
	values1 = make(map[string]interface{})

	values["Accept"]="application/json"
	values["Content-Type"]="application/json"
	urls :="https://api.telegram.org/" +"bot" + telegram_id +"/sendMessage"

	values1["chat_id"] = -chat_id
	values1["text"] = message
	js,_ := json.Marshal(values1)
	req,err := http.NewRequest("POST",urls,strings.NewReader(string(js)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	for key,value := range values {
		req.Header.Add(key,value)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}

	//fmt.Println(string(content))
	return string(content),nil
}

func SendDocument(document, botID string, chat_id int64)(s string,err error){
	/**
	document :文件或者图片的路径
	botID :机器人的ID
	chat_id ：群ID
	*/

	bodyBuf := &bytes.Buffer{} //创建缓存

	bodyWriter := multipart.NewWriter(bodyBuf) // 创建part的writer

	//关键的一步操作，fwimage自行看上图抓包里的，而且这里最好用filepath.Base取文件名不要带路径
	fileWriter, err := bodyWriter.CreateFormFile("document", filepath.Base(document))
	if err != nil {
		fmt.Println("error writing to buffer",err)
	}
	fh, err := os.Open(document)
	if err != nil {
		fmt.Println("error opening file",err)
	}
	defer fh.Close()
	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println(err)
	}
	//添加chat_id
	string_chat_id := strconv.FormatInt(chat_id,10)
	use_chat_id :="-"+string_chat_id
	bodyWriter.WriteField("chat_id", use_chat_id)
	bodyWriter.Close()

	client := &http.Client{}
	telegram_id :=botID
	var values map[string]string
	values = make(map[string]string)
	values["Accept"]="*/*"
	values["Content-Type"]=bodyWriter.FormDataContentType()
	urls :="https://api.telegram.org/" +"bot" + telegram_id +"/sendDocument"
	req,err := http.NewRequest("POST",urls,bodyBuf)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	for key,value := range values {
		req.Header.Add(key,value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	//fmt.Println(string(content))
	return string(content),nil
}

func SendPhoto(photo, botID string, chat_id int64)(s string,err error){
	/**
	photo :文件的路径
	botID :机器人的ID
	chat_id ：群ID
	*/
	bodyBuf := &bytes.Buffer{} //创建缓存
	bodyWriter := multipart.NewWriter(bodyBuf) // 创建part的writer
	//关键的一步操作，fwimage自行看上图抓包里的，而且这里最好用filepath.Base取文件名不要带路径
	fileWriter, err := bodyWriter.CreateFormFile("photo", filepath.Base(photo))
	if err != nil {
		fmt.Println("error writing to buffer",err)
	}
	fh, err := os.Open(photo)
	if err != nil {
		fmt.Println("error opening file",err)
	}
	defer fh.Close()
	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println(err)
	}
	//添加 chat_id
	string_chat_id := strconv.FormatInt(chat_id,10)
	use_chat_id :="-"+string_chat_id
	bodyWriter.WriteField("chat_id", use_chat_id)
	bodyWriter.Close()

	client := &http.Client{}
	telegram_id :=botID
	var values map[string]string
	values = make(map[string]string)
	values["Accept"]="*/*"
	values["Content-Type"]=bodyWriter.FormDataContentType()
	urls :="https://api.telegram.org/" +"bot" + telegram_id +"/sendPhoto"
	req,err := http.NewRequest("POST",urls,bodyBuf)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	for key,value := range values {
		req.Header.Add(key,value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return "",err
	}
	fmt.Println(string(content))
	return string(content),nil
}
