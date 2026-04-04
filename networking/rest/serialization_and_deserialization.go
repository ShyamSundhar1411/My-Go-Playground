package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func SerializationAndDeserialization() {
	user := User{
		Name: "Alice",
		Age:  30,
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error serializing user:", err)	
	}
	fmt.Println("Serialized JSON:", string(jsonData))
	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal("Error deserializing JSON:", err)
	}
	fmt.Println("Deserialized User:", user1)
	jsonData1 := `{"name":"Bob","age":25}`
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)
	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	fmt.Println("Decoded User from Reader:", user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user)
	if err != nil {
		log.Fatal("Error encoding user:", err)
	}
	fmt.Println("Encoded User to Buffer:", buf.String())
}