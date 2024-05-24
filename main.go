package InvestecAPIGo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestInvestec() {
	fmt.Println("Hello Investec API v0.0.1")
}

func GetToken() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
    log.Fatal(err)
  }

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(body))
}
