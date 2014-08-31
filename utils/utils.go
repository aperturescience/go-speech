package utils

import (
  "net/url"
)

type Response struct {
  Results       Results  `json:"result"`
  ResultIndex   int      `json:"result_index"`
}
type Responses []Response

type Result struct {
  Alternatives  Alternatives  `json:"alternative"`
  Final         bool          `json:"final"`
}
type Results []Result

type Alternative struct {
  Transcript    string  `json:"transcript"`
  Confidence    float32 `json:"confidence"`
}
type Alternatives []Alternative

/**
 * Convert map[string]string to a URL querystring
 */
func Queryify(params map[string]string) string {

  values := url.Values{}

  for k, v := range params {
    values.Add(k, v)
  }

  paramsString := values.Encode()

  return paramsString
}