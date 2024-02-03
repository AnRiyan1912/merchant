package models

type RegisterRequest struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
  Person   struct {
    Fullname string `json:"fullname" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Address  string `json:"address" binding:"required"`
  } `json:"person" binding:"required"`
}
