/*
Copyright 2019 The IoTSH Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package message

import (
	"context"
	"fmt"
)

// Message describes a message entry.
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ListMessages returns all messages.
func ListMessages(ctx context.Context, count int) ([]Message, error) {
	messages := make([]Message, count)
	for i := 0; i < count; i++ {
		messages[i].ID = i
		messages[i].Title = fmt.Sprintf("Example %d", i)
		messages[i].Content = fmt.Sprintf("Content of example %d", i)
	}
	return messages, nil
}

// GetMessage returns a message by id.
func GetMessage(ctx context.Context, id int) (*Message, error) {
	return &Message{
		ID:      id,
		Title:   "This is an example",
		Content: "Example content",
	}, nil
}
