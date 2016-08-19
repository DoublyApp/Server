package main

import (
  "encoding/json"
  "bytes"
)

func formatMessages(messages ...Message) string {
  var buffer bytes.Buffer
  for _, message := range messages {
    tmp, err := json.Marshal(message)
    buffer.WriteString(tmp)
  }
  return buffer.String()
}

func formatUsers(users ...User) string {
  var buffer bytes.Buffer
  for _, user := range users {
    tmp, err := json.Marshal(user)
    buffer.WriteString(tmp)
  }
  return buffer.String()
}
