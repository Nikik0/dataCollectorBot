package cache

import (
	"container/list"
	"github.com/Nikik0/dataCollectorBot/internal/model"
	"sync"
)

type Item struct {
	Key   string
	Value model.User
}

type LRU struct {
	mutex    *sync.RWMutex
	capacity int
	queue    *list.List
	items    map[string]*list.Element
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		mutex:    new(sync.RWMutex),
		capacity: capacity,
		queue:    list.New(),
		items:    make(map[string]*list.Element),
	}
}

func (c *LRU) Add(key string, value model.User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return
	}

	if c.queue.Len() == c.capacity {
		c.clear()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element
}

func (c *LRU) Get(key string) (model.User, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	element, exists := c.items[key]
	if !exists {
		return //todo
	}

	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value, nil
}

func (c *LRU) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, found := c.items[key]; found {
		c.deleteItem(val)
	}
}

func (c *LRU) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.items)
}

func (c *LRU) clear() {
	if element := c.queue.Back(); element != nil {
		c.deleteItem(element)
	}
}

func (c *LRU) deleteItem(element *list.Element) {
	item := c.queue.Remove(element).(*Item)
	delete(c.items, item.Key)
}
