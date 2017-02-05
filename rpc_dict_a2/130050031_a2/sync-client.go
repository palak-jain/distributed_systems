
/**
* ArithClient
 */

package main

import (
  "net/rpc"
  "fmt"
  "log"
  // "time"
)

type PartOfSpeech int

const (
    Noun PartOfSpeech = iota
    Verb 
    Adj     
)

type Word struct {
  Key, Meaning string
  Synonyms []*Word
  Pos PartOfSpeech
}


type LoneWord struct {
  Key, Meaning string
  Synonyms []string
  Pos PartOfSpeech
}
 

func PrintWord(w *Word) {
  fmt.Println("Key: ", w.Key, "Meaning: ", w.Meaning, "Pos: ", w.Pos)
  fmt.Printf("Synonyms: ")
  for _, v := range w.Synonyms {
      fmt.Printf(v.Key)
      fmt.Printf(" ")

  }
  fmt.Printf("\n")
}

func InsertWord(a *LoneWord) {
  var reply int
  err := client.Call("Dict.InsertWord", a, &reply)
  if err != nil {
    log.Fatal("[Dict]InsertWord Error: ", err)
  }
  fmt.Println("Success InsertWord")
}

func LookUpWord(s string) {
  var w Word
  err := client.Call("Dict.LookUpWord", &s, &w)
  if err != nil {
    log.Fatal("[Dict]LookUpWord Error: ", err)
  }
  fmt.Println("Success LookUpWord")
  PrintWord(&w)
  return 
}

func RemoveWord(w *Word) {
  var reply int
  err := client.Call("Dict.RemoveWord", w, &reply)
  if err != nil {
    log.Fatal("[Dict] RemoveWord Error: ", err)
  }
  fmt.Println("Success Removeword")
}

var client *rpc.Client

func main() {
  client, _ = rpc.Dial("tcp", "localhost"+":5000")
  // if err != nil {
  //   log.Fatal("dialing:", err)
  // }

  // Synchronous call
  a := &LoneWord{Key: "Barney", Meaning: "Awesome", Pos: Noun}
  b := &LoneWord{Key: "Ted", Meaning: "Sweet", Synonyms: []string{"Barney"}, Pos: Noun}
  c := &LoneWord{Key: "Robin", Meaning: "Daring", Synonyms: []string{"Barney", "Ted"}, Pos: Adj}
  d := &Word{Key: "Barney"}

  InsertWord(a)
  InsertWord(b)
  LookUpWord(a.Key)
  InsertWord(c)
  LookUpWord(c.Key)
  RemoveWord(d)
  LookUpWord(c.Key)
  LookUpWord(b.Key)
  // LookUpWord("yo")
  d.Key = b.Key
  RemoveWord(d)
  d.Key = c.Key
  // time.Sleep(100* time.Millisecond) // change to 5000
  RemoveWord(d)

  
}