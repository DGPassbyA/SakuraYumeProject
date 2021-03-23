package conf

import "time"

var Domain = "127.0.0.1:8081"
var SecretKey = []byte("TestKey")
var Issuer = []byte("SakuraYume")
var ExpiresAt = time.Hour * 24
var DatabasePath = "./clanbattle.db"
