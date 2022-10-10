package cq

var api_url = "http://127.0.0.1:5700" // + url

func SetAiUrl(url string) {
	api_url = url
}

func SetAiPort(port string) {
	api_url = "http://127.0.0.1:" + port
}
