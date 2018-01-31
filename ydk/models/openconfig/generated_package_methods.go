package openconfig


import (
    "runtime"
    "path"
)

func GetCapabilities() map[string]string {
    caps := make(map[string]string)
    caps["cisco-xr-openconfig-interfaces-types"] = "2016-10-12"
    caps["openconfig-bgp"] = "2016-06-21"
    caps["openconfig-bgp-common"] = "2016-06-21"
    caps["openconfig-bgp-common-multiprotocol"] = "2016-06-21"
    caps["openconfig-bgp-common-structure"] = "2016-06-21"
    caps["openconfig-bgp-global"] = "2016-06-21"
    caps["openconfig-bgp-neighbor"] = "2016-06-21"
    caps["openconfig-bgp-peer-group"] = "2016-06-21"
    caps["openconfig-bgp-policy"] = "2016-06-21"
    caps["openconfig-bgp-types"] = "2016-06-21"
    caps["openconfig-channel-monitor"] = "2017-07-08"
    caps["openconfig-extensions"] = "2015-10-09"
    caps["openconfig-if-aggregate"] = "2016-05-26"
    caps["openconfig-if-ethernet"] = "2016-05-26"
    caps["openconfig-if-ip"] = "2016-05-26"
    caps["openconfig-interfaces"] = "2016-05-26"
    caps["openconfig-lacp"] = "2016-05-26"
    caps["openconfig-local-routing"] = "2016-05-11"
    caps["openconfig-mpls"] = "2015-11-05"
    caps["openconfig-mpls-igp"] = "2015-11-05"
    caps["openconfig-mpls-ldp"] = "2015-11-05"
    caps["openconfig-mpls-rsvp"] = "2015-11-05"
    caps["openconfig-mpls-sr"] = "2015-11-05"
    caps["openconfig-mpls-static"] = "2015-11-05"
    caps["openconfig-mpls-te"] = "2015-11-05"
    caps["openconfig-mpls-types"] = "2015-11-05"
    caps["openconfig-optical-amplifier"] = "2017-07-08"
    caps["openconfig-platform"] = "2016-06-06"
    caps["openconfig-platform-transceiver"] = "2016-05-24"
    caps["openconfig-platform-types"] = "2016-06-06"
    caps["openconfig-policy-types"] = "2016-05-12"
    caps["openconfig-rib-bgp"] = "2016-04-11"
    caps["openconfig-rib-bgp-types"] = "2016-04-11"
    caps["openconfig-routing-policy"] = "2016-05-12"
    caps["openconfig-telemetry"] = "2016-02-04"
    caps["openconfig-terminal-device"] = "2016-06-17"
    caps["openconfig-transport-line-common"] = "2017-07-08"
    caps["openconfig-transport-line-protection"] = "2017-03-28"
    caps["openconfig-transport-types"] = "2016-06-17"
    caps["openconfig-types"] = "2017-01-13"
    caps["openconfig-vlan"] = "2016-05-26"
    caps["openconfig-vlan-types"] = "2016-05-26"
    return caps
}

func GetNamespaces() map[string]string {
    namespaces := make(map[string]string)
    namespaces["cisco-xr-openconfig-interfaces-types"] = "http://cisco.com/ns/yang/cisco-xr-openconfig-interfaces-types"
    namespaces["openconfig-bgp"] = "http://openconfig.net/yang/bgp"
    namespaces["openconfig-bgp-policy"] = "http://openconfig.net/yang/bgp-policy"
    namespaces["openconfig-bgp-types"] = "http://openconfig.net/yang/bgp-types"
    namespaces["openconfig-channel-monitor"] = "http://openconfig.net/yang/channel-monitor"
    namespaces["openconfig-extensions"] = "http://openconfig.net/yang/openconfig-ext"
    namespaces["openconfig-if-aggregate"] = "http://openconfig.net/yang/interfaces/aggregate"
    namespaces["openconfig-if-ethernet"] = "http://openconfig.net/yang/interfaces/ethernet"
    namespaces["openconfig-if-ip"] = "http://openconfig.net/yang/interfaces/ip"
    namespaces["openconfig-interfaces"] = "http://openconfig.net/yang/interfaces"
    namespaces["openconfig-lacp"] = "http://openconfig.net/yang/lacp"
    namespaces["openconfig-local-routing"] = "http://openconfig.net/yang/local-routing"
    namespaces["openconfig-mpls"] = "http://openconfig.net/yang/mpls"
    namespaces["openconfig-mpls-ldp"] = "http://openconfig.net/yang/ldp"
    namespaces["openconfig-mpls-rsvp"] = "http://openconfig.net/yang/rsvp"
    namespaces["openconfig-mpls-sr"] = "http://openconfig.net/yang/sr"
    namespaces["openconfig-mpls-types"] = "http://openconfig.net/yang/mpls-types"
    namespaces["openconfig-optical-amplifier"] = "http://openconfig.net/yang/optical-amplfier"
    namespaces["openconfig-platform"] = "http://openconfig.net/yang/platform"
    namespaces["openconfig-platform-transceiver"] = "http://openconfig.net/yang/platform/transceiver"
    namespaces["openconfig-platform-types"] = "http://openconfig.net/yang/platform-types"
    namespaces["openconfig-policy-types"] = "http://openconfig.net/yang/policy-types"
    namespaces["openconfig-rib-bgp"] = "http://openconfig.net/yang/rib/bgp"
    namespaces["openconfig-rib-bgp-types"] = "http://openconfig.net/yang/rib/bgp-types"
    namespaces["openconfig-routing-policy"] = "http://openconfig.net/yang/routing-policy"
    namespaces["openconfig-telemetry"] = "http://openconfig.net/yang/telemetry"
    namespaces["openconfig-terminal-device"] = "http://openconfig.net/yang/terminal-device"
    namespaces["openconfig-transport-line-common"] = "http://openconfig.net/yang/transport-line-common"
    namespaces["openconfig-transport-line-protection"] = "http://openconfig.net/yang/optical-transport-line-protection"
    namespaces["openconfig-transport-types"] = "http://openconfig.net/yang/transport-types"
    namespaces["openconfig-types"] = "http://openconfig.net/yang/openconfig-types"
    namespaces["openconfig-vlan"] = "http://openconfig.net/yang/vlan"
    namespaces["openconfig-vlan-types"] = "http://openconfig.net/yang/vlan-types"
    return namespaces
}

func GetModelsPath() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    return path.Join(path.Dir(filename), "_yang")
}

