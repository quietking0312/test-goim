package comet

import (
	"github.com/Terry-Mao/goim/internal/comet/errors"
	"sync"

	"github.com/Terry-Mao/goim/api/protocol"
)

// Room is a room and store channel room info.
type Room struct {
	ID          string
	channelPool sync.Map
	drop        bool
	Online      int32 // dirty read is ok
	AllOnline   int32
}

// NewRoom new a room struct, store channel room info.
func NewRoom(id string) (r *Room) {
	r = new(Room)
	r.ID = id
	r.drop = false
	r.Online = 0
	return
}

// Put put channel into the room.
// chKey 是 channel 的唯一标识。 确定它的数据类型是字符串
func (r *Room) Put(chKey string, ch *Channel) (err error) {
	if !r.drop {
		r.channelPool.Store(chKey, ch)
		r.Online++
	} else {
		err = errors.ErrRoomDroped
	}
	return
}

// Del delete channel from the room.
func (r *Room) Del(chKey string) bool {
	r.channelPool.Delete(chKey)
	r.Online--
	r.drop = r.Online == 0
	return r.drop
}

// Push push msg to the room, if chan full discard it.
func (r *Room) Push(p *protocol.Proto) {
	r.channelPool.Range(func(key, value interface{}) bool {
		ch, ok := value.(*Channel)
		if !ok {
			r.Del(key.(string))
			return true
		}
		_ = ch.Push(p)
		return true
	})
}

// Close close the room.
func (r *Room) Close() {
	r.channelPool.Range(func(key, value interface{}) bool {
		ch, ok := value.(*Channel)
		if !ok {
			r.Del(key.(string))
			return true
		}
		ch.Close()
		r.Del(key.(string))
		return true
	})
}

// OnlineNum the room all online.
func (r *Room) OnlineNum() int32 {
	if r.AllOnline > 0 {
		return r.AllOnline
	}
	return r.Online
}
