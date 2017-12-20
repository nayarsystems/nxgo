package nxcore

import (
	"encoding/json"
	"time"
)

// UserCreate creates new user in Nexus user's table.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserCreate(user, pass string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"pass": pass,
	}
	return nc.Exec("user.create", par)
}

// UserDelete removes user from Nexus user's table.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserDelete(user string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
	}
	return nc.Exec("user.delete", par)
}

// UserRename changes the name of a user from Nexus user's table.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserRename(user string, new string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"new":  new,
	}
	return nc.Exec("user.rename", par)
}

type UserInfo struct {
	User        string                            `json:"user"`
	CreatedAt   time.Time                         `json:"createdAt"`
	Tags        map[string]map[string]interface{} `json:"tags"`
	Templates   []string                          `json:"templates"`
	Whitelist   []string                          `json:"whitelist"`
	Blacklist   []string                          `json:"blacklist"`
	MaxSessions int                               `json:"maxsessions"`
	Disabled    bool                              `json:"disabled"`
}

// UserList lists users from Nexus user's table.
// Returns a list of UserInfo or error.
func (nc *NexusConn) UserList(prefix string, limit int, skip int, opts ...*ListOpts) ([]UserInfo, error) {
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
	res, err := nc.Exec("user.list", par)
	if err != nil {
		return nil, err
	}
	users := make([]UserInfo, 0)
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UserCount counts users from Nexus user's table.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserCount(prefix string, opts ...*CountOpts) (interface{}, error) {
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
	return nc.Exec("user.count", par)
}

// UserGetTags gets the user tags
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserGetTags(user string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
	}
	return nc.Exec("user.getTags", par)
}

// UserGetEffectiveTags gets the user effective tags on a prefix (higher tags on hierachy will overwrite lower ones)
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserGetEffectiveTags(user string, prefix string) (interface{}, error) {
	par := map[string]interface{}{
		"user":   user,
		"prefix": prefix,
	}
	return nc.Exec("user.getEffectiveTags", par)
}

// UserSetTags set tags on user's prefix.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserSetTags(user string, prefix string, tags map[string]interface{}) (interface{}, error) {
	par := map[string]interface{}{
		"user":   user,
		"prefix": prefix,
		"tags":   tags,
	}
	return nc.Exec("user.setTags", par)
}

// UserDelTags remove tags from user's prefix.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserDelTags(user string, prefix string, tags []string) (interface{}, error) {
	par := map[string]interface{}{
		"user":   user,
		"prefix": prefix,
		"tags":   tags,
	}
	return nc.Exec("user.delTags", par)
}

// UserSetPass sets new user password.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserSetPass(user string, pass string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"pass": pass,
	}
	return nc.Exec("user.setPass", par)
}

// UserAddTemplate adds a new template to the user.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserAddTemplate(user, template string) (interface{}, error) {
	par := map[string]interface{}{
		"user":     user,
		"template": template,
	}
	return nc.Exec("user.addTemplate", par)
}

// UserDelTemplate removes a template from the user.
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserDelTemplate(user, template string) (interface{}, error) {
	par := map[string]interface{}{
		"user":     user,
		"template": template,
	}
	return nc.Exec("user.delTemplate", par)
}

// UserAddWhitelist adds an IP to the user's whitelist.
// IP is a regex that will be matched against the client source address
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserAddWhitelist(user, ip string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"ip":   ip,
	}
	return nc.Exec("user.addWhitelist", par)
}

// UserDelWhitelist removes an IP from the user's whitelist.
// IP is a regex that will be matched against the client source address
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserDelWhitelist(user, ip string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"ip":   ip,
	}
	return nc.Exec("user.delWhitelist", par)
}

// UserAddBlacklist adds an IP to the user's blacklist.
// IP is a regex that will be matched against the client source address
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserAddBlacklist(user, ip string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"ip":   ip,
	}
	return nc.Exec("user.addBlacklist", par)
}

// UserDelBlacklist removes an IP from the user's whitelist.
// IP is a regex that will be matched against the client source address
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserDelBlacklist(user, ip string) (interface{}, error) {
	par := map[string]interface{}{
		"user": user,
		"ip":   ip,
	}
	return nc.Exec("user.delBlacklist", par)
}

// UserSetMaxSessions set the maximum number of sessions a client can open
// Setting the value lower than the current number of sessions won't kill any session
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserSetMaxSessions(user string, sessions int) (interface{}, error) {
	par := map[string]interface{}{
		"user":        user,
		"maxsessions": sessions,
	}
	return nc.Exec("user.setMaxSessions", par)
}

// UserSetDisable set disable flag on an user
// Setting the value to true will prevent any login from this user
// Returns the response object from Nexus or error.
func (nc *NexusConn) UserSetDisabled(user string, disabled bool) (interface{}, error) {
	par := map[string]interface{}{
		"user":     user,
		"disabled": disabled,
	}
	return nc.Exec("user.setDisabled", par)
}
