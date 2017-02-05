package main

import (
    "fmt"
    // "log"
    "errors"
    "net/http"
    // "net/rpc"
    "time"
    "sync"
    // "net"
    // "encoding/json"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
    "github.com/gorilla/mux"
  )

type Dict int
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

var dict map[string]*Word
var mutex = &sync.Mutex{}

// Return AlreadyExists or OtherServerSideError

func (t *Dict) InsertWord(word *LoneWord, reply *int) error { 
    fmt.Println("Server: InsertWord")
    _, exists := dict[word.Key]
    if exists {
      return errors.New("AlreadyExists")
    }

    var res Word  
    res.Key = word.Key
    res.Meaning = word.Meaning
    res.Pos = word.Pos
    synonyms := word.Synonyms
    res.Synonyms = make([]*Word, len(synonyms), len(synonyms))
    for i, w := range synonyms {
      syn, exist := dict[w]
      if !exist {
        return errors.New("SynonymUndefined")
      }
      res.Synonyms[i] = syn
    }
    mutex.Lock()
    dict[res.Key] = &res
    mutex.Unlock()
    *reply = 1
    return nil
}

// The error contains either UnknownWord or OtherServerSideError.  
// Make this procedure take 5 seconds using sleep. 
// RemoveWord should also remove this as the synonym from all other words that has this word as its synonym.

func (t* Dict) RemoveWord(word *Word, reply *int) error {
  fmt.Println("Remove: ", word.Key)
  _, exists  := dict[word.Key]
  if !exists {
    return errors.New("UnknownWord")
  }

  mutex.Lock()
  for v := range dict {
    for i, s := range dict[v].Synonyms {
      if s.Key == word.Key {
        dict[v].Synonyms[i] = dict[v].Synonyms[len(dict[v].Synonyms) - 1]
        dict[v].Synonyms = dict[v].Synonyms[:len(dict[v].Synonyms) - 1]
        fmt.Println("Remove synonym from ", v)
      }
    }
  } 
  // TODO: return OtherServerSideError
  delete(dict, word.Key)
  mutex.Unlock()
  *reply = 0

  time.Sleep(5000 * time.Millisecond) 
  return nil

} 

// The error contains either UnknownWord or OtherServerSideError.
func (t* Dict) LookUpWord(key *string, res *Word) error {
  fmt.Println("Lookup: ", *key)
  mutex.Lock()
  w, exists  := dict[*key]
  mutex.Unlock()
  if !exists {
    return errors.New("UnknownWord")
  }
  *res = *w 
  // TODO: return OtherServerSideError
  return nil
} 

func main() {

  // dictionary
  dict = make(map[string]*Word)

  // register Dict
  dictionary := new(Dict)

    rpcServer := rpc.NewServer()

    rpcServer.RegisterCodec(json.NewCodec(), "application/json")
    rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")


    rpcServer.RegisterService(dictionary, "dict")

    router := mux.NewRouter()
    router.Handle("/rpc", rpcServer)
    fmt.Println("Launching http server @ 6000...")

    http.ListenAndServe(":6000", router)


  // s := rpc.NewServer()
  // s.RegisterCodec(json.NewCodec(), "application/json")
  // s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

  // s.RegisterService(dictionary, "Dict")
  // http.Handle("/rpc", s)
  // fmt.Println("Launching http server @ 6000...")

  // http.ListenAndServe(":6000", s)
  // rpc.Register(dictionary)
  // rpc.HandleHTTP()

  // go http.ListenAndServe(":6000", nil)
  // fmt.Println("Launching http server @ 6000...")

  // // if e != nil {
  // //   log.Fatal("listen error:", e)
  // // } 

  // // tcp
  // server := rpc.NewServer()
  // server.Register(dictionary)
  // server.RegisterCodec(json.NewCodec(), "application/json")

  // fmt.Println("Launching tcp server @ 5000...")

  // l, e := net.Listen("tcp", ":5000")
  // if e != nil {
  //   log.Fatal("listen error:", e)
  // }

  // server.Accept(l)

}