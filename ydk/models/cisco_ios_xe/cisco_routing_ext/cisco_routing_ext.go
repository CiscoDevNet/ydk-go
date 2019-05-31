// This YANG module is an extention to 'ietf-routing'
// module and describes addtional operational
// data for route list.
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
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

type Igrp struct {
}

func (id Igrp) String() string {
	return "cisco-routing-ext:igrp"
}

type Nhrp struct {
}

func (id Nhrp) String() string {
	return "cisco-routing-ext:nhrp"
}

type Hsrp struct {
}

func (id Hsrp) String() string {
	return "cisco-routing-ext:hsrp"
}

type Lisp struct {
}

func (id Lisp) String() string {
	return "cisco-routing-ext:lisp"
}

