package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zbitech/common/pkg/model/ztypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ZcashNodeConf struct {
	list       map[string][]string
	properties map[string]string
}

func NewZcashConf(network ztypes.NetworkType, txIndex bool, miner bool) *ZcashNodeConf {

	conf := ZcashNodeConf{
		list:       map[string][]string{},
		properties: map[string]string{},
	}

	switch network {
	case ztypes.TESTNET_TYPE:
		conf.setProperty("testnet", "1")
		conf.setList("addnodes", ztypes.TESTNET)
	case ztypes.MAINNET_TYPE:
		conf.setProperty("testnet", "0")
		conf.setList("addnodes", ztypes.MAINNET)
	}

	conf.setProperty("listen", "1")
	conf.setProperty("printtoconsole", "1")
	conf.setProperty("server", "1")

	conf.setProperty("showmetrics", "1")
	conf.setProperty("logips", "1")

	conf.setProperty("rpcclienttimeout", ztypes.ZCASH_RPCCLIENT_TIMEOUT)
	conf.setProperty("maxconnections", ztypes.ZCASH_MAX_CONNECTIONS)

	if txIndex {
		conf.setProperty("txindex", "1")
	}

	if miner {
		conf.setProperty("gen", "1")
		conf.setProperty("genproclimit", "1")
		conf.setProperty("equihashsolver", ztypes.ZCASH_SOLVER)
	}

	conf.setList("rpcallowip", []string{"0.0.0.0/0"})
	conf.setList("connect", []string{})

	return &conf
}

func (z ZcashNodeConf) SetPort(port int32) {
	z.setProperty("rpcbind", "0.0.0.0")
	z.setProperty("rpcport", strconv.FormatInt(int64(port), 10))
}

func (z ZcashNodeConf) setProperty(name, val string) {
	z.properties[name] = val
}

func (z ZcashNodeConf) getProperty(name, def string) string {
	value, ok := z.properties[name]

	if ok {
		return value
	}

	return def
}

func (z ZcashNodeConf) setList(name string, values []string) {
	z.list[name] = values
}

func (z ZcashNodeConf) getList(name string) []string {
	if value, ok := z.list[name]; ok {
		return value
	} else {
		return []string{}
	}
}

func (z ZcashNodeConf) AddPeer(peer string) {
	peers := z.getList("connect")
	peers = append(peers, peer)
	z.setList("connect", peers)
}

func (z ZcashNodeConf) Value() string {

	b := new(bytes.Buffer)

	// write out properties
	for key, value := range z.properties {
		if value != "" && value != "0" {
			fmt.Fprintf(b, "    %s=%s\n", key, value)
		}
	}

	// write out list as individual items
	for key, list := range z.list {
		for _, item := range list {
			fmt.Fprintf(b, "    %s=%s\n", key, item)
		}
	}

	return b.String()
}

func (z *ZcashNodeConf) toMap() map[string]interface{} {
	values := map[string]interface{}{}
	for key, value := range z.properties {
		if value != "" && value != "0" {
			values[key] = value
		}
	}

	for key, list := range z.list {
		values[key] = list
	}

	return values
}

func (z *ZcashNodeConf) fromMap(values map[string]interface{}) {
	for key, value := range values {
		if key == "connect" || key == "addnodes" || key == "rpcallowip" {
			a := value.(primitive.A)
			arr := []interface{}(a)
			//arr := value.([]interface{})
			//			log.Printf("Value of %s is %s - %s", key, value, arr)
			//			z.list[key] = arr
			z.list[key] = make([]string, len(arr))
			for idx, item := range arr {
				z.list[key][idx] = item.(string)
			}
			//			z.list[key] = arr
		} else {
			z.properties[key] = value.(string)
		}
	}
}

func (z *ZcashNodeConf) MarshalJSON() ([]byte, error) {
	return json.Marshal(z.toMap())
}

func (z *ZcashNodeConf) UnmarshalJSON(data []byte) error {
	values := map[string]interface{}{}
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}

	z.fromMap(values)
	return nil
}

// func (z *ZcashNodeConf) MarshalBSON() ([]byte, error) {
// 	return bson.Marshal(z.toMap())
// }

// func (z *ZcashNodeConf) UnmarshalBSON(data []byte) error {

// 	conf := NewZcashConf(ztypes.MAINNET_TYPE, false, false)

// 	values := map[string]interface{}{}
// 	err := bson.Unmarshal(data, &values)
// 	if err != nil {
// 		return err
// 	}

// 	conf.fromMap(values)
// 	log.Printf("Zcash.conf = %s, %s", conf.properties, conf.list)
// 	z.list = conf.list
// 	z.properties = conf.properties

// 	log.Printf("Zcash.conf = %s, %s", z.properties, z.list)

// 	return nil
// }
