package saver

import (
	"github.com/ozoncp/ocp-tenant-api/internal/flusher"
	"github.com/ozoncp/ocp-tenant-api/internal/tenant"
	"time"
)

type Saver interface {
	Save(entity tenant.Tenant)
	Close()
}

type RemoveRule uint8

const (
	All RemoveRule = iota
	Last
)

type saver struct {
	capacity   uint
	flusher    flusher.Flusher
	removeRule RemoveRule
	buffer     []tenant.Tenant
	toSave     chan tenant.Tenant
	toClose    chan struct{}
	close      bool
	ticker     *time.Ticker
}

func (s saver) work() {
	for {
		select {
		case entity := <-s.toSave:
			s.appendToSave(entity)
		case <-s.ticker.C:
			s.flush()
		case <-s.toClose:
			s.flush()
			break
		}
	}
}

func (s saver) appendToSave(entity tenant.Tenant) {
	if len(s.buffer) == cap(s.buffer) {
		switch s.removeRule {
		case All:
			s.buffer = s.buffer[:0]
		case Last:
			s.buffer = s.buffer[:len(s.buffer)-1]
		}
	}
	s.buffer = append(s.buffer, entity)
}

func (s saver) flush() {
	if len(s.buffer) > 0 {
		flushResult := s.flusher.Flush(s.buffer)
		if flushResult == nil {
			s.buffer = s.buffer[:0]
		} else {
			s.buffer = s.buffer[:copy(s.buffer, flushResult)]
		}
	}
}

func (s saver) Save(entity tenant.Tenant) {
	if s.close {
		panic("Saver closed.")
	}
	s.toSave <- entity
}

func (s saver) Close() {
	s.ticker.Stop()
	s.toClose <- struct{}{}
	s.close = true
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	removeRule RemoveRule,
	period time.Duration) Saver {
	result := saver{
		capacity:   capacity,
		flusher:    flusher,
		removeRule: removeRule,
	}
	result.ticker = time.NewTicker(period)
	go result.work()
	return result
}
