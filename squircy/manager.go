package squircy

import (
	"github.com/go-martini/martini"
)

type Manager struct {
	*martini.ClassicMartini
}

func NewManager() (manager *Manager) {
	manager = &Manager{martini.Classic()}
	manager.Map(NewConfiguration("config.json"))
	manager.invokeAndMap(newIrcConnection)
	manager.invokeAndMap(newHandlerCollection)
	manager.invokeAndMap(newRedisClient)
	
	return
}

func (manager *Manager) invokeAndMap(fn interface{}) {
	res, err := manager.Invoke(fn)
	if err != nil {
		panic(err)
	}
	manager.Map(res[0].Interface())
}