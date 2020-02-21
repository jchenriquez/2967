package threesumclosest

import (
  "testing"
  "encoding/json"
  "bufio"
  "io"
  "os"
)

type Test struct {
  Input []int `json:"input"`
  Target int `json:"target"`
  Output int `json:"output"`
}

func TestThreeSumClosests(t *testing.T) {
  f, err := os.Open("tests.json")

  if err != nil {
    t.Error(err)
    return
  }

  defer f.Close()

  reader := bufio.NewReader(f)
  decoder := json.NewDecoder(reader)
  var tests map[string]Test

  for {
    err = decoder.Decode(&tests)

    if err == nil {
      for name, test := range tests {
        t.Run(name, func (st *testing.T) {
          res := ThreeSumClosest(test.Input, test.Target)

          if res != test.Output {
            st.Errorf("returned answer was %d", res)
          }
        })
      }
    } else if err == io.EOF {
      break
    } else {
      t.Error(err)
      break
    }
  }
}