package main

import (
  "fmt"
  "net/http"
  "io"
  "os"
  //"io/ioutil"
  "strings"
  //"flag"
)

func getStuffHTTP(url string) []byte {
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

  return b
}

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
  b := getStuffHTTP(url)
  imgSrc := getFirstImage(string(b))

  imgSrc = "https://apod.nasa.gov/apod/" + imgSrc
  imgType := imgSrc[strings.LastIndex(imgSrc, "."):]
  fmt.Println(imgType)
  imgData := getStuffHTTP(imgSrc)

  //fmt.Println(imgSrc)

  //save image of the day in public folder
  err := os.WriteFile("./public/nasaImg" + imgType, imgData, 0666)
  //check(err)




  fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	err = http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
		panic(err)
	}
}
