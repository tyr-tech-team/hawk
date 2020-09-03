package etcd

import (
	"hawk/config/encoder"
	"strings"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

func makeMap(e encoder.Encoder, kv []*mvccpb.KeyValue, stripPrefix string) map[string]interface{} {
	data := make(map[string]interface{})

	for _, v := range kv {
		data = update(e, data, v, "put", stripPrefix)
	}

	return data
}

func update(e encoder.Encoder, data map[string]interface{}, v *mvccpb.KeyValue, action, stripPrefix string) map[string]interface{} {
	// remove prefix if non empty, and ensure leading / is removed as well
	vkey := strings.TrimPrefix(strings.TrimPrefix(string(v.Key), stripPrefix), "/")
	// split on prefix
	haveSplit := strings.Contains(vkey, "/")
	keys := strings.Split(vkey, "/")

	var vals interface{}
	e.Decode(v.Value, &vals)

	if !haveSplit && len(keys) == 1 {
		switch action {
		case "delete":
			data = make(map[string]interface{})
		default:
			v, ok := vals.(map[string]interface{})
			if ok {
				data = v
			}
		}
		return data
	}

	// set data for first iteration
	kvals := data
	// iterate the keys and make maps
	for i, k := range keys {
		kval, ok := kvals[k].(map[string]interface{})
		if !ok {
			// create next map
			kval = make(map[string]interface{})
			// set it
			kvals[k] = kval
		}

		// last key: write vals
		if l := len(keys) - 1; i == l {
			switch action {
			case "delete":
				delete(kvals, k)
			default:
				kvals[k] = vals
			}
			break
		}

		// set kvals for next iterator
		kvals = kval
	}

	return data
}
