package serialize

import (
	erro "github.com/LeorlikK/socketio/internal/errors"
)

const (
	ErrUnsupportedUseRead erro.String = "Serialize() method unsupported, use the Read() method instead"
	ErrUnsupported        erro.State  = "method: unsupported"
)
