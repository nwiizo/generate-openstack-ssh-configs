package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

/** JSONデコード用に構造体定義 */
type Server struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Networks string `json:"networks"`
	Image    string `json:"image"`
	Flavor   string `json:"flavor"`
}

func main() {
	if f, err := os.Stat("./sample.json"); os.IsNotExist(err) || f.IsDir() {
		cmdstr := "openstack server list -f json > sample.json"
		err := exec.Command("sh", "-c", cmdstr).Run()
		if err != nil {
			fmt.Println("Command Exec Error.")
		}
	}
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var servers []Server
	if err := json.Unmarshal(bytes, &servers); err != nil {
		log.Fatal(err)
	}
	// デコードしたデータを表示
	for _, s := range servers {
		fmt.Printf("# Connection ID:%s\n", s.Id)
		fmt.Printf("Host %s\n", s.Name)
		if strings.Contains(s.Networks, ";") == true {
			fmt.Printf("# Change it Pick HostName\n")
			networks := strings.Split(s.Networks, ";")
			for network := range networks {
				ipadd := strings.Split(networks[network], "=")
				fmt.Printf("#  HostName %s\n", ipadd[1])
			}
		} else {
			network := strings.Split(s.Networks, "=")
			fmt.Printf("  HostName %s\n", network[1])
		}
		fmt.Printf("  user root\n")
		fmt.Printf("  port 22\n")
		fmt.Printf("  IdentityFile ~/.ssh/keys/key\n\n")
	}
}
