package media

import "github.com/prodbox/weasn/kernel/message"

type MediaUpload struct {
	message.JsonErrorMessage
	MediaType string   `json:"type"`
	MediaId   string   `json:"media_id"`
	CreatedAt int64    `json:"created_at"`
	Item      []string `json:"item"`
}
