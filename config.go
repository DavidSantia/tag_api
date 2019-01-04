package tag_api

// Folders and credentials
const (
	// Retries to wait for docker DB instance
	DbConnectRetries = 5

	// NATS server
	NHost = "nats://localhost:4222"
	NSub  = "update"
)

// 16-byte JSON Web Token encryption key
var JwtKey = []byte{194, 164, 235, 6, 138, 248, 171, 239, 24, 216, 11, 22, 137, 199, 215, 133}

// Session key
var SessionKey = []byte("something-very-secret")
