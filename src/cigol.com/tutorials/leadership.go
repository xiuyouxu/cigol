package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var server string = "http://192.168.13.109:2379/v2/keys"
	go watchAll(server)

	var leaderRoot string = "/cigol-leader"
	data, err := get(server + leaderRoot)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var stat map[string]interface{}
	if err = json.Unmarshal(data, &stat); err != nil {
		panic(err)
	}
	fmt.Println(stat)
	if _, exist := stat["errorCode"]; exist {
		fmt.Println("leader root does not exist:", leaderRoot)
		// put the leader root
		_, err = postOrPut("PUT", server+leaderRoot, "dir=true")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//		fmt.Println(string(data))
	}
	data, err = get(server + leaderRoot)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	var stat2 map[string]interface{}
	if err = json.Unmarshal(data, &stat2); err != nil {
		panic(err)
	}
	if _, exist := stat2["errorCode"]; exist {
		fmt.Println("can not find the leader root, exit")
		os.Exit(1)
	}

	//create the queue node
	data, err = postOrPut("POST", server+leaderRoot, "value=cigol", "ttl=5")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	if err = json.Unmarshal(data, &stat); err != nil {
		panic(err)
	}
	stat = stat["node"].(map[string]interface{})
	go refresh(server, stat["key"].(string))
	go watch(server, stat["key"].(string), leaderRoot)

	time.Sleep(time.Hour * 1600)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading response body error", err)
		return nil, err
	}
	return body, nil
}

func postOrPut(method, url string, params ...string) ([]byte, error) {
	client := &http.Client{}
	param := ""
	for _, p := range params {
		param = param + p + "&"
	}
	if len(param) > 0 {
		param = param[0 : len(param)-1]
	}
	fmt.Println("put params:", param)
	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	// need to set the content-type, or the params not work
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading response body error", err)
		return nil, err
	}
	return body, nil
}

func watch(server, realName, leaderRoot string) {
	fmt.Println("in watch(), realName is", realName)

	body, err := get(server + leaderRoot + "?recursive=true&sorted=true")
	if err != nil {
		fmt.Println(realName, "watch error", err)
		watch(server, realName, leaderRoot)
	} else {
		q := realName[len(leaderRoot)+1:]
		me, err := strconv.Atoi(q)
		if err != nil {
			fmt.Println("failed to get self queue number", err)
			return
		}
		fmt.Println("me=", me)
		var list map[string]interface{}
		if err = json.Unmarshal(body, &list); err != nil {
			fmt.Println("unmarshalling watch list error", err)
			return
		}
		nodes := extract(list, "node", "nodes")
		if nodes != nil {
			n := -1
			watchPath := ""

			queues := nodes.([]interface{})
			for _, queue := range queues {
				q = queue.(map[string]interface{})["key"].(string)[len(leaderRoot)+1:]
				k, err := strconv.Atoi(q)
				if err != nil {
					fmt.Println("failed to get self queue number", err)
					return
				}
				fmt.Println("k=", k, "n=", n, "watchPath=", watchPath)
				if k > n && k < me {
					n = k
					watchPath = queue.(map[string]interface{})["key"].(string)
				}
			}
			if watchPath == "" {
				fmt.Println(realName, "is the leader now...")
			} else {
				fmt.Println(realName, "is watching ", watchPath)
				// watch the prev node
				get(server + watchPath + "?wait=true")
				// rewatch
				watch(server, realName, leaderRoot)
			}
		}
	}
}

func refresh(server, realName string) {
	ticker := time.NewTicker(time.Millisecond * 5000)
	for t := range ticker.C {
		fmt.Println("refresh", realName, "at", t)
		doRefresh(server, realName)
	}
}

func doRefresh(server, realName string) {
	postOrPut("PUT", server+realName, "ttl=5", "refresh=true", "prevExist=true")
}

func extract(m map[string]interface{}, keys ...string) interface{} {
	t := m
	var val interface{} = nil
	for _, key := range keys {
		if v, e := t[key]; e {
			t, _ = v.(map[string]interface{})
			val = v
		}
	}
	return val
}

func watchAll(server string) {
	body, err := get(server + "/?recursive=true&wait=true")
	if err != nil {
		fmt.Println("reading response body error", err)
	} else {
		fmt.Println(string(body))
	}
	watchAll(server)
}
