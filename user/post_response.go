package user

type PostResponse struct {
	ID      int    `json:"id"`
	IdUser  int    `json:"user_id"`
	Content string `json:"content"`
}

type PostListResponse struct {
	Data []Post `json:"data"`
}
