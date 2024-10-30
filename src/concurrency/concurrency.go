package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BodyResponse struct {
	UserId int8
	Id int16
	Title string
	Completed bool
}

type Ouput struct {
	Status uint16
	Json BodyResponse
	error string
}

// func hi(num int) {
// 	fmt.Println("Hola ", num)
// 	time.Sleep(1000 * time.Millisecond)
// }

func get(num int, c chan <- Ouput) {

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", num)

	resp, err := http.Get(url);

	if err != nil {
		c <- Ouput{ Status: uint16(resp.StatusCode), error: err.Error() }
		return
	}
	defer resp.Body.Close();

	var body BodyResponse;

	errDeco := json.NewDecoder(resp.Body).Decode(&body)

	if errDeco != nil {
		c <- Ouput{ Status: 500, error: errDeco.Error() }
		return
	}

	c <- Ouput{ Status: uint16(resp.StatusCode), Json: body }
}

func Concurrency() {

	c := make(chan Ouput, 100);

	for i := 0; i < 100; i++ {
		go get(i, c)
	}

	for i := 0; i < 100; i++ {
		val := <-c

		fmt.Printf("idx: %d, and value: %v \n", i, val)
	}

	close(c)

}