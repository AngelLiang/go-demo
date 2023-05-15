package main

import (
	"fmt"
	"io/ioutil"
    "encoding/json"
    "strings"
	"net/http"
	// "net/url"
)

const GithubUserUrl = "https://api.github.com/user"

// type GithubUserUrlResp struct {
//     Code int `json:"current_user_url"`
//     Msg string `json:"msg"`
//     // Data CreateTXData `json:"data"`
// }

type GithubUserUrlResp struct {
    Message string `json:"message"`
    DocumentationUrl string `json:"documentation_url"`
    // Data CreateTXData `json:"data"`
}


// json 转结构体
func BuildGithubUserUrlResp(body []byte) (*GithubUserUrlResp, error){
    var resp GithubUserUrlResp
    // err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&resp)
	err := json.Unmarshal(body, &resp)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}


func BuildHttpQuery(params map[string]string) (paramStr string) {
	paramsArr := make([]string, 0, len(params))
	for k, v := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", k, v))
	}
	paramStr = strings.Join(paramsArr, "&")
	return paramStr
}

// doGET 发起一个HTTP请求，使用GET方法
// 参数:
//   url: HTTP请求的URL
//   params: querystring的参数，以键值对形式传递
// 返回值:
//   body: HTTP响应的内容，类型为[]byte
//   err: 发起请求过程中的错误，类型为error
func doGET(url string, params map[string]string) ([]byte, error) {
	// 构建完整的URL
	requestURL := url + "?"
	requestURL += BuildHttpQuery(params)

	// 发起GET请求
	response, err := http.Get(requestURL)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}


func main() {
	// 示例用法
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	body, err := doGET(GithubUserUrl, params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println("body:", body)

	resp, err := BuildGithubUserUrlResp(body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("resp:", resp)
}
