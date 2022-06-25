package memory

import (
	"fmt"
	"github.com/zbitech/common/pkg/utils"
	"log"
	"sync"
)

type MemoryStore struct {
	memory map[string]interface{}
	lock   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		memory: map[string]interface{}{},
	}
}

func (m *MemoryStore) StoreItem(id string, item interface{}) {

	m.lock.RLock()
	defer m.lock.RUnlock()

	m.memory[id] = item
}

func (m *MemoryStore) AddItem(id string, item interface{}) error {

	m.lock.RLock()
	defer m.lock.RUnlock()

	array, ok := m.memory[id].([]interface{})
	if !ok {
		log.Printf("Store is not an array - %s - %v", array, ok)
		return fmt.Errorf("Store is not an array")
	}

	log.Printf("Adding item %s to %s", item, id)
	array = append(array, item)
	m.memory[id] = array

	return nil
}

func (m *MemoryStore) RemoveItem(id string) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	m.memory[id] = nil
}

func (m *MemoryStore) GetItem(id string) (interface{}, error) {

	log.Print("START MemoryStore.GetItem")
	defer log.Print("END MemoryStore.GetItem")

	m.lock.RLock()
	defer m.lock.RUnlock()

	if item, ok := m.memory[id]; ok {
		log.Printf("Retrieved object - %s - %v", utils.MarshalObject(item), item)
		return item, nil
	} else {
		log.Printf("Item %s not found", id)
		return nil, fmt.Errorf("Item %s does not exist", id)
	}
}

func (m *MemoryStore) GetItems() map[string]interface{} {

	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.memory
}
