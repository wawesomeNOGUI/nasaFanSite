package main

import (
  "fmt"
  "net/http"
  "io"
  //"io/ioutil"
  "strings"
  //"flag"
)

func getFirstImage(siteCode string) string {
  picSrcStart := strings.Index(siteCode, "IMG SRC=")
  if picSrcStart < 0 {
    fmt.Println("oh no couldn't find picture :()")
  }

  picSrcStart += 9  //length of "IMG SRC=""

  //get rid of rest of sitecode before IMG SRC
  siteCode = siteCode[picSrcStart:]

  picSrcEnd := strings.Index(siteCode, "\"")

  siteCode = siteCode[:picSrcEnd]

  //return img src value
  return siteCode
}

func main(){
  url := "https://apod.nasa.gov/apod/astropix.html?"

  var b []byte  //Stores byte text data
  var err error

  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  b, err = io.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  imgSrc := getFirstImage(string(b))
  imgSrc = "https://apod.nasa.gov/apod/" + imgSrc

  fmt.Println(imgSrc)




  fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	err = http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
		panic(err)
	}
}
