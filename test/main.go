package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	err := status.Error(codes.InvalidArgument, "qwe")
	fmt.Println(err)
	m, e := json.Marshal(err)
	fmt.Println(m, e)
	var responseErr error
	e = json.Unmarshal(m, &responseErr)
	fmt.Println(responseErr, e)
}
