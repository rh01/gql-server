package graphql

// Config ..
type Config struct {
	Store  Store  `json:"store"`
	Logger Logger `json:"logger"`
}

// Store ..
type Store struct {
	Server    string
	Name      string
	Username  string
	Password  string
	Timeout   string
	Monotonic bool
	Retry     string // duration

	// old info
	IdleConnection int
	MaxConnection  int
	Log            bool
}

// Logger ..x``
type Logger struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}
