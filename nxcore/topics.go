package nxcore

import (
	"encoding/json"

	"github.com/jaracil/ei"
)

type TopicInfo struct {
	Topic       string `json:"topic"`
	Subscribers int    `json:"subscribers"`
}

// TopicSubscribe subscribes a pipe to a topic.
// Returns the response object from Nexus or error.
func (nc *NexusConn) TopicSubscribe(pipe *Pipe, topic string) (interface{}, error) {
	par := ei.M{
		"pipeid": pipe.Id(),
		"topic":  topic,
	}
	return nc.Exec("topic.sub", par)
}

// TopicUnsubscribe unsubscribes a pipe from a topic.
// Returns the response object from Nexus or error.
func (nc *NexusConn) TopicUnsubscribe(pipe *Pipe, topic string) (interface{}, error) {
	par := ei.M{
		"pipeid": pipe.Id(),
		"topic":  topic,
	}
	return nc.Exec("topic.unsub", par)
}

// TopicPublish publishes message to a topic.
// Returns the response object from Nexus or error.
func (nc *NexusConn) TopicPublish(topic string, msg interface{}) (interface{}, error) {
	par := ei.M{
		"topic": topic,
		"msg":   msg,
	}
	return nc.Exec("topic.pub", par)
}

// TopicList lists subscriptions to topics from Nexus.
// Returns a list of TopicInfo or error.
func (nc *NexusConn) TopicList(prefix string, limit int, skip int, opts ...*ListOpts) ([]TopicInfo, error) {
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
	res, err := nc.Exec("topic.list", par)
	if err != nil {
		return nil, err
	}
	topics := make([]TopicInfo, 0)
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &topics)
	if err != nil {
		return nil, err
	}

	return topics, nil
}

// TopicCount counts subscriptions to topics from Nexus.
// Returns the response object from Nexus or error.
func (nc *NexusConn) TopicCount(prefix string, opts ...*CountOpts) (interface{}, error) {
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
	return nc.Exec("topic.count", par)
}
