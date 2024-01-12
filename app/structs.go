package main

type CommonFields struct {
	ID string `json:"id"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Ping struct {
	Key     string `json:"key"`
	SvgName string `json:"svg-name"`

	*CommonFields
}

type User struct {
	Username string `json:"username"`

	*CommonFields
}
