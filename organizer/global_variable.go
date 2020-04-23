package organizer

import (
	"os/user"
)

var StockPath string
var FileServerPath string
var User *user.User
var Port string
var AllowedOrigin string