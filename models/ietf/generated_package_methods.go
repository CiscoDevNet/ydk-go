package ietf


import (
    "runtime"
    "path"
)

func GetCapabilities() map[string]string {
    caps := make(map[string]string)
    caps["iana-crypt-hash"] = "2014-08-06"
    caps["iana-if-type"] = "2014-05-08"
    caps["ietf-diffserv-action"] = "2015-04-07"
    caps["ietf-diffserv-classifier"] = "2015-04-07"
    caps["ietf-diffserv-policy"] = "2015-04-07"
    caps["ietf-diffserv-target"] = "2015-04-07"
    caps["ietf-event-notifications"] = "2016-10-27"
    caps["ietf-inet-types"] = "2013-07-15"
    caps["ietf-interfaces"] = "2014-05-08"
    caps["ietf-interfaces-ext"] = "None"
    caps["ietf-ip"] = "2014-06-16"
    caps["ietf-ipv4-unicast-routing"] = "2015-05-25"
    caps["ietf-ipv6-unicast-routing"] = "2015-05-25"
    caps["ietf-key-chain"] = "2015-02-24"
    caps["ietf-netconf"] = "2011-06-01"
    caps["ietf-netconf-acm"] = "2012-02-22"
    caps["ietf-netconf-monitoring"] = "2010-10-04"
    caps["ietf-netconf-notifications"] = "2012-02-06"
    caps["ietf-netconf-with-defaults"] = "2011-06-01"
    caps["ietf-ospf"] = "2015-03-09"
    caps["ietf-restconf-monitoring"] = "2016-08-15"
    caps["ietf-routing"] = "2015-05-25"
    caps["ietf-syslog-types"] = "2015-11-09"
    caps["ietf-yang-library"] = "2016-06-21"
    caps["ietf-yang-push"] = "2016-10-28"
    caps["ietf-yang-smiv2"] = "2012-06-22"
    caps["ietf-yang-types"] = "2013-07-15"
    caps["policy-types"] = "2013-10-07"
    return caps
}

func GetNamespaces() map[string]string {
    namespaces := make(map[string]string)
    namespaces["iana-crypt-hash"] = "urn:ietf:params:xml:ns:yang:iana-crypt-hash"
    namespaces["iana-if-type"] = "urn:ietf:params:xml:ns:yang:iana-if-type"
    namespaces["ietf-diffserv-action"] = "urn:ietf:params:xml:ns:yang:ietf-diffserv-action"
    namespaces["ietf-diffserv-classifier"] = "urn:ietf:params:xml:ns:yang:ietf-diffserv-classifier"
    namespaces["ietf-diffserv-policy"] = "urn:ietf:params:xml:ns:yang:ietf-diffserv-policy"
    namespaces["ietf-diffserv-target"] = "urn:ietf:params:xml:ns:yang:ietf-diffserv-target"
    namespaces["ietf-event-notifications"] = "urn:ietf:params:xml:ns:yang:ietf-event-notifications"
    namespaces["ietf-inet-types"] = "urn:ietf:params:xml:ns:yang:ietf-inet-types"
    namespaces["ietf-interfaces"] = "urn:ietf:params:xml:ns:yang:ietf-interfaces"
    namespaces["ietf-interfaces-ext"] = "urn:ietf:params:xml:ns:yang:ietf-interfaces-ext"
    namespaces["ietf-ip"] = "urn:ietf:params:xml:ns:yang:ietf-ip"
    namespaces["ietf-ipv4-unicast-routing"] = "urn:ietf:params:xml:ns:yang:ietf-ipv4-unicast-routing"
    namespaces["ietf-ipv6-unicast-routing"] = "urn:ietf:params:xml:ns:yang:ietf-ipv6-unicast-routing"
    namespaces["ietf-key-chain"] = "urn:ietf:params:xml:ns:yang:ietf-key-chain"
    namespaces["ietf-netconf"] = "urn:ietf:params:xml:ns:netconf:base:1.0"
    namespaces["ietf-netconf-acm"] = "urn:ietf:params:xml:ns:yang:ietf-netconf-acm"
    namespaces["ietf-netconf-monitoring"] = "urn:ietf:params:xml:ns:yang:ietf-netconf-monitoring"
    namespaces["ietf-netconf-notifications"] = "urn:ietf:params:xml:ns:yang:ietf-netconf-notifications"
    namespaces["ietf-netconf-with-defaults"] = "urn:ietf:params:xml:ns:yang:ietf-netconf-with-defaults"
    namespaces["ietf-ospf"] = "urn:ietf:params:xml:ns:yang:ietf-ospf"
    namespaces["ietf-restconf-monitoring"] = "urn:ietf:params:xml:ns:yang:ietf-restconf-monitoring"
    namespaces["ietf-routing"] = "urn:ietf:params:xml:ns:yang:ietf-routing"
    namespaces["ietf-syslog-types"] = "urn:ietf:params:xml:ns:yang:ietf-syslog-types"
    namespaces["ietf-yang-library"] = "urn:ietf:params:xml:ns:yang:ietf-yang-library"
    namespaces["ietf-yang-push"] = "urn:ietf:params:xml:ns:yang:ietf-yang-push"
    namespaces["ietf-yang-smiv2"] = "urn:ietf:params:xml:ns:yang:ietf-yang-smiv2"
    namespaces["ietf-yang-types"] = "urn:ietf:params:xml:ns:yang:ietf-yang-types"
    namespaces["policy-types"] = "urn:ietf:params:xml:ns:yang:c3pl-types"
    return namespaces
}

func GetModelsPath() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }
    return path.Join(path.Dir(filename), "_yang")
}

