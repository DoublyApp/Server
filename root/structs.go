package main

type Message struct {
  MessageID int
  SenderID int
  ReceiverID int
  MessageText int
  TimeCreated string
}

type User struct {
  UserID int
  UserName string
  Email string
  Password string
  DOB string
  Gender string
}

type UserInterest struct {
  UserID int
  InterestID int
}

type UsersFriends struct {
  UserID int
  FriendID int
  FriendStatus string
}

type GPS struct {
  UserID int
  Latitude float32
  Longitude float32
}

type Interest struct {
  InterestID int
  InterestName string
}
