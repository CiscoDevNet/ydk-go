// This module contains a collection of YANG groupings 
// in filter configurations for policy model.
package policy_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package policy_types"))
}

type Control struct {
}

func (id Control) String() string {
	return "policy-types:control"
}

type InputInterface struct {
}

func (id InputInterface) String() string {
	return "policy-types:input-interface"
}

type SrcMac struct {
}

func (id SrcMac) String() string {
	return "policy-types:src-mac"
}

type Qos struct {
}

func (id Qos) String() string {
	return "policy-types:qos"
}

type PerfMon struct {
}

func (id PerfMon) String() string {
	return "policy-types:perf-mon"
}

type Application struct {
}

func (id Application) String() string {
	return "policy-types:application"
}

type SecurityGroupName struct {
}

func (id SecurityGroupName) String() string {
	return "policy-types:security-group-name"
}

type PacketService struct {
}

func (id PacketService) String() string {
	return "policy-types:packet-service"
}

type QosClass struct {
}

func (id QosClass) String() string {
	return "policy-types:qos-class"
}

type Ipv4AclName struct {
}

func (id Ipv4AclName) String() string {
	return "policy-types:ipv4-acl-name"
}

type FlowDlci struct {
}

func (id FlowDlci) String() string {
	return "policy-types:flow-dlci"
}

type ControlClass struct {
}

func (id ControlClass) String() string {
	return "policy-types:control-class"
}

type InspectClass struct {
}

func (id InspectClass) String() string {
	return "policy-types:inspect-class"
}

type AppnavClass struct {
}

func (id AppnavClass) String() string {
	return "policy-types:appnav-class"
}

type Service struct {
}

func (id Service) String() string {
	return "policy-types:service"
}

type Dei struct {
}

func (id Dei) String() string {
	return "policy-types:dei"
}

type Prec struct {
}

func (id Prec) String() string {
	return "policy-types:prec"
}

type AccessControlClass struct {
}

func (id AccessControlClass) String() string {
	return "policy-types:access-control-class"
}

type PacketLength struct {
}

func (id PacketLength) String() string {
	return "policy-types:packet-length"
}

type Ipv4Acl struct {
}

func (id Ipv4Acl) String() string {
	return "policy-types:ipv4-acl"
}

type FlowDe struct {
}

func (id FlowDe) String() string {
	return "policy-types:flow-de"
}

type FlowIp struct {
}

func (id FlowIp) String() string {
	return "policy-types:flow-ip"
}

type FlowRecord struct {
}

func (id FlowRecord) String() string {
	return "policy-types:flow-record"
}

type VlanInner struct {
}

func (id VlanInner) String() string {
	return "policy-types:vlan-inner"
}

type AccessControl struct {
}

func (id AccessControl) String() string {
	return "policy-types:access-control"
}

type Metadata struct {
}

func (id Metadata) String() string {
	return "policy-types:metadata"
}

type Vlan struct {
}

func (id Vlan) String() string {
	return "policy-types:vlan"
}

type AtmVci struct {
}

func (id AtmVci) String() string {
	return "policy-types:atm-vci"
}

type Appnav struct {
}

func (id Appnav) String() string {
	return "policy-types:appnav"
}

type Inspect struct {
}

func (id Inspect) String() string {
	return "policy-types:inspect"
}

type ClassMap struct {
}

func (id ClassMap) String() string {
	return "policy-types:class-map"
}

type QosGroup struct {
}

func (id QosGroup) String() string {
	return "policy-types:qos-group"
}

type WlanUserPriority struct {
}

func (id WlanUserPriority) String() string {
	return "policy-types:wlan-user-priority"
}

type IpRtp struct {
}

func (id IpRtp) String() string {
	return "policy-types:ip-rtp"
}

type Ipv6Acl struct {
}

func (id Ipv6Acl) String() string {
	return "policy-types:ipv6-acl"
}

type AtmClp struct {
}

func (id AtmClp) String() string {
	return "policy-types:atm-clp"
}

type DstMac struct {
}

func (id DstMac) String() string {
	return "policy-types:dst-mac"
}

type Cos struct {
}

func (id Cos) String() string {
	return "policy-types:cos"
}

type Pbr struct {
}

func (id Pbr) String() string {
	return "policy-types:pbr"
}

type DeiInner struct {
}

func (id DeiInner) String() string {
	return "policy-types:dei-inner"
}

type MplsExpTop struct {
}

func (id MplsExpTop) String() string {
	return "policy-types:mpls-exp-top"
}

type CosInner struct {
}

func (id CosInner) String() string {
	return "policy-types:cos-inner"
}

type Ipv6AclName struct {
}

func (id Ipv6AclName) String() string {
	return "policy-types:ipv6-acl-name"
}

type MplsExpImp struct {
}

func (id MplsExpImp) String() string {
	return "policy-types:mpls-exp-imp"
}

type SecurityGroupTag struct {
}

func (id SecurityGroupTag) String() string {
	return "policy-types:security-group-tag"
}

type ClassType struct {
}

func (id ClassType) String() string {
	return "policy-types:class-type"
}

type DiscardClass struct {
}

func (id DiscardClass) String() string {
	return "policy-types:discard-class"
}

type Vpls struct {
}

func (id Vpls) String() string {
	return "policy-types:vpls"
}

type PolicyType struct {
}

func (id PolicyType) String() string {
	return "policy-types:policy-type"
}

// Metric represents metric
type Metric string

const (
    Metric_none Metric = "none"

    Metric_peta Metric = "peta"

    Metric_tera Metric = "tera"

    Metric_giga Metric = "giga"

    Metric_mega Metric = "mega"

    Metric_kilo Metric = "kilo"

    Metric_milli Metric = "milli"

    Metric_nano Metric = "nano"
)

// Direction represents 
type Direction string

const (
    Direction_inbound Direction = "inbound"

    Direction_outbound Direction = "outbound"
)

// RateUnit represents ratio:   ratio
type RateUnit string

const (
    RateUnit_pps RateUnit = "pps"

    RateUnit_cps RateUnit = "cps"

    RateUnit_bps RateUnit = "bps"

    RateUnit_perc RateUnit = "perc"

    RateUnit_ratio RateUnit = "ratio"
)

