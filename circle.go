package engine

import (
	"sync"

	"github.com/Monibuca/engine/avformat"
)

const CIRCLE_SIZE = 512

type CircleItem struct {
	*avformat.AVPacket
	next  *CircleItem
	pre   *CircleItem
	index int
	*sync.RWMutex
}

func CreateCircle() (p *CircleItem) {
	p = &CircleItem{AVPacket: new(avformat.AVPacket), RWMutex: new(sync.RWMutex)}
	first := p
	for i := 0; i < CIRCLE_SIZE; i++ {
		p.next = &CircleItem{pre: p, index: i, AVPacket: new(avformat.AVPacket), RWMutex: new(sync.RWMutex)}
		p = p.next
	}
	first.pre = p
	p.next = first
	return
}
