package api
import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func SendMessage(message, botID string, chat_id int64){
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
	req,_ := http.NewRequest("POST",urls,strings.NewReader(string(js)))
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
