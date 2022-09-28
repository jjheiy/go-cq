package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Json map[string]interface{}
type List []Json

func Post(url string, params []byte) (Json, error) {
	// 1. 创建http客户端实例
	client := &http.Client{}
	// 2. 创建请求实例
	req, err := http.NewRequest("POST", url, strings.NewReader(string(params)))
	if err != nil {
		panic(err)
	}
	// 3. 设置请求头，可以设置多个
	req.Header.Set("Host", " ")
	req.Header.Set("Content-Type", "application/json")
	// 4. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 5. 一次性读取请求到的数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var v interface{}
	json.Unmarshal(body, &v)
	if value, ok := v.([]interface{}); ok {
		data := make(map[string]any)
		temp := make(List, 0)
		for _, v := range value {
			temp = append(temp, v.(map[string]interface{}))
		}
		data["list"] = temp
		return data, nil
	}
	data := v.(map[string]interface{})
	return data, nil
}

func Get(url string) (Json, error) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	json.Unmarshal(b, &v)
	if value, ok := v.([]interface{}); ok {
		data := make(Json)
		temp := make(List, 0)
		for _, v := range value {
			temp = append(temp, v.(map[string]interface{}))
		}
		data["data"] = temp
		return data, nil
	}
	data := v.(map[string]interface{})
	return data, nil
}
