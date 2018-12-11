package kmq

import (
	"github.com/Shopify/sarama"
)

// Offset - interface for control your own offset
type Offset interface {
	// Load be called when creating new consumer
	// Init your offset in this method
	Load() (int64, error)

	// Save be called when get a new message
	// Persistence your offset value in this method
	Save(offset int64) error
}

//=======================================================
// NewestOffset return newest offset
type NewestOffset struct {
}

// Load - return sarama.OffsetNewest
func (n *NewestOffset) Load() (int64, error) {
	return sarama.OffsetNewest, nil
}

// Save - nil
func (n *NewestOffset) Save(offset int64) error {
	return nil
}

//=======================================================
// OldestOffset return oldest offset
type OldestOffset struct {
}

// Load - return sarama.OffsetOldest
func (n *OldestOffset) Load() (int64, error) {
	return sarama.OffsetOldest, nil
}

// Save - nil
func (n *OldestOffset) Save(offset int64) error {
	return nil
}
