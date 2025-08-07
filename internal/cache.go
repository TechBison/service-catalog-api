package internal

import (
	"sync"

	"gorm.io/gorm"
)

var (
	serviceCache = make(map[int]Service)
	cacheMutex   = sync.RWMutex{}
)

func LoadCache(db *gorm.DB) {
	var services []Service
	db.Preload("Versions").Find(&services)

	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	for _, s := range services {
		serviceCache[s.ID] = s
	}
}

func GetCachedServices() []Service {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	services := make([]Service, 0, len(serviceCache))
	for _, s := range serviceCache {
		services = append(services, s)
	}
	return services
}

func GetCachedServiceByID(id int) (Service, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	s, ok := serviceCache[id]
	return s, ok
}
