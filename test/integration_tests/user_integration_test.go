package endtoend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hunter32292/go-server-example/cmd"
	"github.com/hunter32292/go-server-example/pkg/models"
)

var (
	Server *http.Server
)

func setupServer() {
	// Execute in Root directory of project
	os.Chdir("../..")
	go func(server *http.Server) {
		cmd.StartServer(server)
	}(Server)
	// Wait for Server Startup
	time.Sleep(time.Second * 5)
}

func cleanupServer() error {
	var err error
	// Thread this off since we're running this as a process in the backend
	err = cmd.StopServer(Server)

	if err != nil {
		return err
	}
	return nil
}

// Integration Test Example
func TestRequestingUserData(t *testing.T) {
	Server = cmd.GetServer()
	setupServer()

	t.Cleanup(func() {
		fmt.Println("Cleaning up integration test...")
		err := cleanupServer()
		if err != nil {
			fmt.Println(err)
		}
	})

	url := "http://localhost:8080/users"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	var dataList []*models.User

	err = json.Unmarshal(body, &dataList)
	if err != nil {
		t.Error(err)
	}

	if len(dataList) <= 0 {
		t.Error(err)
	}
}
