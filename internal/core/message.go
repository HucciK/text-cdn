package core

type Message struct {
	Type     string `json:"type" bson:"type"`
	Language string `json:"language" bson:"language"`
	Text     string `json:"text" bson:"text,omitempty"`
}
