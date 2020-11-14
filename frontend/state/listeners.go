package state

import (
	"github.com/aclisp/godashboard/frontend/state/util"
)

// Listeners is the listeners that will be invoked when the store changes.
var listeners = util.NewListenerRegistry()

// AddListener adds a listener function only if the key does not exist.
func AddListener(key interface{}, listener func()) {
	if listeners.Has(key) {
		return
	}
	listeners.Add(key, listener)
}
