package main

import (
  "fmt"
  "log"
  "bytes"
  "regexp"
  "net/http"
  "os"
  "io/ioutil"
  "encoding/json"
  "github.com/aperturescience/go-speech/utils"
)

func main() {

  var err error

  // 1. read audio file in memory
  audio, err := ioutil.ReadFile("test.wav")

  if err != nil {
    log.Fatalf("Error reading file:\n%v\n", err)
  }

  reader := bytes.NewReader(audio)

  // 2. send binary data to API endpoint
  endpoint := "https://www.google.com/speech-api/v2/recognize?"

  params := map[string]string {
    "output"  : "json",
    "lang"    : "en-us",
    "key"     : os.Getenv("GOOGLE_SPEECH_API_KEY"),
  }

  resp, err := http.Post(endpoint + utils.Queryify(params), "audio/l16; rate=16000;", reader)

  if err != nil {
    log.Fatalf("Error POST-ing data to Google Speech Endpoint:\n %v\n", err)
  }

  // 3. read response from API and unmarshal to struct
  body, _ := ioutil.ReadAll(resp.Body)

  if resp.StatusCode != 200 {
    log.Fatalf("Google Speech Error:\n%v\n", string(body))
  }

  // 4. remove false response
  rgx := regexp.MustCompile("{\"result\":\\[\\]}")
  sanitized := rgx.ReplaceAllLiteralString(string(body), "")

  // 5. create response struct via unmarshal
  var response utils.Response

  err = json.Unmarshal([]byte(sanitized), &response)

  if err != nil {
    log.Fatalf("Error Unmarshaling JSON:\n%v\n", err)
  }

  fmt.Printf("%+v\n", response)
}