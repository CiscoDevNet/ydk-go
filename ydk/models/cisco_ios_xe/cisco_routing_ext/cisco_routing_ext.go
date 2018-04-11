// This YANG module is an extention to 'ietf-routing'
// module and describes addtional operational
// data for route list
package cisco_routing_ext

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_routing_ext"))
}

type Rip struct {
}

func (id Rip) String() string {
	return "cisco-routing-ext:rip"
}

type IsIs struct {
}

func (id IsIs) String() string {
	return "cisco-routing-ext:is-is"
}

type Bgp struct {
}

func (id Bgp) String() string {
	return "cisco-routing-ext:bgp"
}

type Eigrp struct {
}

func (id Eigrp) String() string {
	return "cisco-routing-ext:eigrp"
}

type Mobile struct {
}

func (id Mobile) String() string {
	return "cisco-routing-ext:mobile"
}

