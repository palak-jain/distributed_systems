
/**
* ArithClient
 */

package main

import (
  "net/rpc"
  "fmt"
  // "log"
  // "os"
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
  insertCall := client.Go("Dict.InsertWord", a, &reply, nil)
  // fmt.Println("Insert word")
  for {
    select {
    case <- insertCall.Done:
      fmt.Println("Success Inserting word")
      return
    default:
      fmt.Println("130050031")
    }
  }
}
                
func LookUpWord(s string) {
  var w Word
  lookCall := client.Go("Dict.LookUpWord", &s, &w, nil)
  // fmt.Println("Insert word")
  for {
    select {
    case <- lookCall.Done:
      fmt.Println("Success Lookup")
      PrintWord(&w)
      return
    default:
      fmt.Println("130050031")
    }
  }
 
}

func RemoveWord(w *Word) {
  var reply int
  removeCall := client.Go("Dict.RemoveWord", w, &reply, nil)
  // fmt.Println("Remove word")
  for {
    select {
    case <- removeCall.Done:
      fmt.Println("Success removing word")
      return
    default:
      fmt.Println("Palak Jain    130050031")
    }
  }
}

var client *rpc.Client

func main() {
  client, _ = rpc.DialHTTP("tcp", "localhost"+":6000")


  // Asynchronous call
  a := &LoneWord{Key: "Barney2", Meaning: "Awesome", Pos: Noun}
  b := &LoneWord{Key: "Ted2", Meaning: "Sweet", Synonyms: []string{"Barney"}, Pos: Noun}
  c := &LoneWord{Key: "Robin2", Meaning: "Daring", Synonyms: []string{"Barney", "Ted"}, Pos: Adj}
  d := &Word{Key: "Barney2"}

  InsertWord(a)
  //  InsertWord(a) // throws alreadyexists error
  InsertWord(b)
  LookUpWord(a.Key)
  InsertWord(c)
  LookUpWord(c.Key)
  RemoveWord(d)
  LookUpWord(c.Key)
  LookUpWord(b.Key)
  // LookUpWord("yo") throws unknown word error
  d.Key = b.Key
  RemoveWord(d)
  d.Key = c.Key
  RemoveWord(d)
  // RemoveWord(d) throws unknown word error

  
}