package conf

import "time"

var Domain = "127.0.0.1:8081"
var SecretKey = []byte("SecretKey")
var Issuer = []byte("Issuer")
var ExpiresAt = time.Hour * 24
var DatabasePath = "./clanbattle.db"
