package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Response struct {
 Pair     string `json:"Pair"`
 Total   string `json:"Total"`
 Volume string    `json:"Volume"`
 Fees string `json:"Fees"`
}
func main() {
 fmt.Println("Calling API...")
client := &http.Client{}
 req, err := http.NewRequest("GET", "http://127.0.0.1:5000/v2/0x63607de7ae773638d012561a01383ab8ac321371", nil)
 if err != nil {
  fmt.Print(err.Error())
 }
 //req.Header.Add("Accept", "application/json")
 //req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)
 if err != nil {
  fmt.Print(err.Error())
 }
defer resp.Body.Close()
 bodyBytes, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Print(err.Error())
 }
var responseObject Response
 json.Unmarshal(bodyBytes, &responseObject)
 fmt.Printf("API Response as struct %+v\n", responseObject)
}