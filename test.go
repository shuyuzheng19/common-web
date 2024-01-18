package main

import (
	"common-web-framework/models"
	"encoding/json"
	"fmt"
)

func main() {
	var user *models.User

	json.Unmarshal([]byte(`{"id":1}`), &user)

	fmt.Println(user)

}
