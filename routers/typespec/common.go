package typespec

type ErrResult struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
