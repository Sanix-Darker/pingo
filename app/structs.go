package main

type commonFields struct {
	//pk id depending on the struct
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Ping struct {
	*commonFields

	// Key: will be the {key}.svg file
	// and will be the key generated (human readable)
	Key    string `json:"key"`
	UserId string `json:"user_id"`
	// Will be the whole counts of strikes linked to a single ping
	Strikes int `json:"strikes"`
}

type Strike struct {
	PingId string `json:"ping_id"`
	// From: the website link.
	From string `json:"from"`
	// StrikedAt: The date of the strike.
	StrikedAt string `json:"strike_at"`
}

type User struct {
	*commonFields

	// Username: extracted from the email or just the email.
	Username string `json:"username"`
	// from the github / google login
	LoginProvider string `json:"login_provider"`
}

// TODO: methods related to each struct (DTO) for building and structuring
// each model to be ready for any database interface.
