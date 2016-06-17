package main

import (
   "encoding/json"
    "io/ioutil"
    "testing"
)

func TestJSON(t *testing.T) {
    buf, _ := ioutil.ReadFile("plugins.json")
    s:= string(buf)
    if isJSON(s){
       //nothing here since pass by default
    } else{
        t.Fail()
    }
    
}

func isJSON( str string ) bool {
     var temp interface{}
     return json.Unmarshal([]byte(str), &temp) == nil
 }

