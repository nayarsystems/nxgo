package nxcore

import (
	"encoding/json"

	"github.com/jaracil/ei"
)

type LockInfo struct {
	Id    string `json:"id"`
	Owner string `json:"owner"`
}

// Lock tries to get a lock.
// Returns lock success/failure or error.
func (nc *NexusConn) Lock(lock string) (interface{}, error) {
	par := ei.M{
		"lock": lock,
	}
	res, err := nc.Exec("sync.lock", par)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Unlock tries to free a lock.
// Returns unlock success/failure or error.
func (nc *NexusConn) Unlock(lock string) (interface{}, error) {
	par := ei.M{
		"lock": lock,
	}
	res, err := nc.Exec("sync.unlock", par)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LockList lists locks from Nexus.
// Returns a list of LockInfo or error.
func (nc *NexusConn) LockList(prefix string, limit int, skip int, opts ...*ListOpts) ([]LockInfo, error) {
	par := map[string]interface{}{
		"prefix": prefix,
		"limit":  limit,
		"skip":   skip,
	}
	if len(opts) > 0 {
		if opts[0].LimitByDepth {
			par["depth"] = opts[0].Depth
		}
		if opts[0].Filter != "" {
			par["filter"] = opts[0].Filter
		}
	}
	res, err := nc.Exec("sync.list", par)
	if err != nil {
		return nil, err
	}
	locks := make([]LockInfo, 0)
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &locks)
	if err != nil {
		return nil, err
	}

	return locks, nil
}

// LockCount counts locks from Nexus.
// Returns the response object from Nexus or error.
func (nc *NexusConn) LockCount(prefix string, opts ...*CountOpts) (interface{}, error) {
	par := map[string]interface{}{
		"prefix": prefix,
	}
	if len(opts) > 0 {
		if opts[0].Subprefixes {
			par["subprefixes"] = opts[0].Subprefixes
		}
		if opts[0].Filter != "" {
			par["filter"] = opts[0].Filter
		}
	}
	return nc.Exec("sync.count", par)
}
