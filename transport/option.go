package transport

import with "github.com/LeorlikK/socketio/internal/option"

type option = with.Option
type optionWith = with.OptionWith

func WithSocketRoomFilter(fn func(Namespace, Room, SocketID) (bool, error)) option {
	return func(o optionWith) {
		if ary, ok := o.(*SocketArray); ok {
			ary.filterOnRoom = fn
		}
	}
}
