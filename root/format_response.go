package main

import (
  "encoding/json"
  "bytes"
)

func FormatMessages(messages ...Message) string {
  var buffer bytes.Buffer
  for _, message := range messages {
    tmp, _ := json.Marshal(message)
    buffer.WriteString(string(tmp))
  }
  return buffer.String()
}

func FormatUsers(users []User) string {
  var buffer bytes.Buffer
  for _, user := range users {
    tmp, _ := json.Marshal(user)
    buffer.WriteString(string(tmp))
  }

  return buffer.String()
}
