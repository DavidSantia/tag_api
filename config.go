package tag_api

// Folders and credentials
const (
	// MySQL DB info
	DbHost = "localhost"
	DbPort = "6603"
	DbUser = "demo"
	DbPass = "welcome1"
	DbName = "tagdemo"
)

// 16-byte JSON Web Token encryption key
var JwtKey = []byte{194, 164, 235, 6, 138, 248, 171, 239, 24, 216, 11, 22, 137, 199, 215, 133}

// Session key
var SessionKey = []byte("something-very-secret")
