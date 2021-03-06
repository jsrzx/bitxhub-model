package pb

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/meshplus/bitxhub-kit/types"
)

func (m *IBTP) ID() string {
	return fmt.Sprintf("%s-%s-%d", m.From, m.To, m.Index)
}

func (m *IBTP) Hash() types.Hash {
	body, err := json.Marshal([]interface{}{
		m.From,
		m.To,
		m.Index,
		m.Type,
		m.Timestamp,
		m.Payload,
		m.Extra,
	})
	if err != nil {
		panic(err)
	}

	data := sha256.Sum256(body)

	return types.Bytes2Hash(data[:])
}