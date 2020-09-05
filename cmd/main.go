package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type AddReq struct {
	OperandA int
	OperandB int
}

type AddReply struct {
	Result int
}

func main() {
	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		var add_req AddReq
		json.NewDecoder(r.Body).Decode(&add_req)
		log.Println("req: ", add_req.OperandA, add_req.OperandB)
		var add_reply AddReply
		add_reply.Result = add_req.OperandA + add_req.OperandB
		json.NewEncoder(w).Encode(add_reply)
	})

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("../web/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
	// curl -s -XPOST -d'{"OperandA":10,"OperandB":20}' http://localhost:8080/api/add
}
