package rfc1315_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rfc1315_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:RFC1315-MIB RFC1315-MIB}", reflect.TypeOf(RFC1315MIB{}))
    ydk.RegisterEntity("RFC1315-MIB:RFC1315-MIB", reflect.TypeOf(RFC1315MIB{}))
}

// RFC1315MIB
type RFC1315MIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    FrameRelayGlobals RFC1315MIB_FrameRelayGlobals

    // The Parameters for the Data Link Connection Management Interface for the
    // frame relay service on this interface.
    Frdlcmitable RFC1315MIB_Frdlcmitable

    // A table containing information about specific Data Link Connection
    // Identifiers and corresponding virtual circuits.
    Frcircuittable RFC1315MIB_Frcircuittable

    // A table containing information about Errors on the Frame Relay interface.
    Frerrtable RFC1315MIB_Frerrtable
}

func (rFC1315MIB *RFC1315MIB) GetFilter() yfilter.YFilter { return rFC1315MIB.YFilter }

func (rFC1315MIB *RFC1315MIB) SetFilter(yf yfilter.YFilter) { rFC1315MIB.YFilter = yf }

func (rFC1315MIB *RFC1315MIB) GetGoName(yname string) string {
    if yname == "frame-relay-globals" { return "FrameRelayGlobals" }
    if yname == "frDlcmiTable" { return "Frdlcmitable" }
    if yname == "frCircuitTable" { return "Frcircuittable" }
    if yname == "frErrTable" { return "Frerrtable" }
    return ""
}

func (rFC1315MIB *RFC1315MIB) GetSegmentPath() string {
    return "RFC1315-MIB:RFC1315-MIB"
}

func (rFC1315MIB *RFC1315MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frame-relay-globals" {
        return &rFC1315MIB.FrameRelayGlobals
    }
    if childYangName == "frDlcmiTable" {
        return &rFC1315MIB.Frdlcmitable
    }
    if childYangName == "frCircuitTable" {
        return &rFC1315MIB.Frcircuittable
    }
    if childYangName == "frErrTable" {
        return &rFC1315MIB.Frerrtable
    }
    return nil
}

func (rFC1315MIB *RFC1315MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frame-relay-globals"] = &rFC1315MIB.FrameRelayGlobals
    children["frDlcmiTable"] = &rFC1315MIB.Frdlcmitable
    children["frCircuitTable"] = &rFC1315MIB.Frcircuittable
    children["frErrTable"] = &rFC1315MIB.Frerrtable
    return children
}

func (rFC1315MIB *RFC1315MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rFC1315MIB *RFC1315MIB) GetBundleName() string { return "cisco_ios_xe" }

func (rFC1315MIB *RFC1315MIB) GetYangName() string { return "RFC1315-MIB" }

func (rFC1315MIB *RFC1315MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rFC1315MIB *RFC1315MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rFC1315MIB *RFC1315MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rFC1315MIB *RFC1315MIB) SetParent(parent types.Entity) { rFC1315MIB.parent = parent }

func (rFC1315MIB *RFC1315MIB) GetParent() types.Entity { return rFC1315MIB.parent }

func (rFC1315MIB *RFC1315MIB) GetParentYangName() string { return "RFC1315-MIB" }

// RFC1315MIB_FrameRelayGlobals
type RFC1315MIB_FrameRelayGlobals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable  indicates  whether  the  system produces the
    // frDLCIStatusChange trap. The type is Frtrapstate.
    Frtrapstate interface{}
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetFilter() yfilter.YFilter { return frameRelayGlobals.YFilter }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) SetFilter(yf yfilter.YFilter) { frameRelayGlobals.YFilter = yf }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetGoName(yname string) string {
    if yname == "frTrapState" { return "Frtrapstate" }
    return ""
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetSegmentPath() string {
    return "frame-relay-globals"
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frTrapState"] = frameRelayGlobals.Frtrapstate
    return leafs
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetBundleName() string { return "cisco_ios_xe" }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetYangName() string { return "frame-relay-globals" }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) SetParent(parent types.Entity) { frameRelayGlobals.parent = parent }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetParent() types.Entity { return frameRelayGlobals.parent }

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetParentYangName() string { return "RFC1315-MIB" }

// RFC1315MIB_FrameRelayGlobals_Frtrapstate represents produces the frDLCIStatusChange trap.
type RFC1315MIB_FrameRelayGlobals_Frtrapstate string

const (
    RFC1315MIB_FrameRelayGlobals_Frtrapstate_enabled RFC1315MIB_FrameRelayGlobals_Frtrapstate = "enabled"

    RFC1315MIB_FrameRelayGlobals_Frtrapstate_disabled RFC1315MIB_FrameRelayGlobals_Frtrapstate = "disabled"
)

// RFC1315MIB_Frdlcmitable
// The Parameters for the Data Link Connection Management
// Interface for the frame relay service on this
// interface.
type RFC1315MIB_Frdlcmitable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Con- nection Management
    // Interface. The type is slice of RFC1315MIB_Frdlcmitable_Frdlcmientry.
    Frdlcmientry []RFC1315MIB_Frdlcmitable_Frdlcmientry
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetFilter() yfilter.YFilter { return frdlcmitable.YFilter }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) SetFilter(yf yfilter.YFilter) { frdlcmitable.YFilter = yf }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetGoName(yname string) string {
    if yname == "frDlcmiEntry" { return "Frdlcmientry" }
    return ""
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetSegmentPath() string {
    return "frDlcmiTable"
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frDlcmiEntry" {
        for _, c := range frdlcmitable.Frdlcmientry {
            if frdlcmitable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1315MIB_Frdlcmitable_Frdlcmientry{}
        frdlcmitable.Frdlcmientry = append(frdlcmitable.Frdlcmientry, child)
        return &frdlcmitable.Frdlcmientry[len(frdlcmitable.Frdlcmientry)-1]
    }
    return nil
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frdlcmitable.Frdlcmientry {
        children[frdlcmitable.Frdlcmientry[i].GetSegmentPath()] = &frdlcmitable.Frdlcmientry[i]
    }
    return children
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetBundleName() string { return "cisco_ios_xe" }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetYangName() string { return "frDlcmiTable" }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) SetParent(parent types.Entity) { frdlcmitable.parent = parent }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetParent() types.Entity { return frdlcmitable.parent }

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetParentYangName() string { return "RFC1315-MIB" }

// RFC1315MIB_Frdlcmitable_Frdlcmientry
// The Parameters for a particular Data Link Con-
// nection Management Interface.
type RFC1315MIB_Frdlcmitable_Frdlcmientry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the  corresponding  ifEn-
    // try. The type is interface{} with range: -2147483648..2147483647.
    Frdlcmiifindex interface{}

    // This variable states which Data  Link  Connec- tion Management scheme is
    // active (and by impli- cation, what DLCI it uses) on the  Frame  Relay
    // interface. The type is Frdlcmistate.
    Frdlcmistate interface{}

    // This variable states which address  format  is in use on the Frame Relay
    // interface. The type is Frdlcmiaddress.
    Frdlcmiaddress interface{}

    // This variable states which address  length  in octets.  In the case of Q922
    // format, the length indicates the entire length of the address  in- cluding
    // the control portion. The type is Frdlcmiaddresslen.
    Frdlcmiaddresslen interface{}

    // This is the number of seconds between  succes- sive status enquiry
    // messages. The type is interface{} with range: 5..30.
    Frdlcmipollinginterval interface{}

    // Number of status enquiry intervals  that  pass before  issuance  of a full
    // status enquiry mes- sage. The type is interface{} with range: 1..255.
    Frdlcmifullenquiryinterval interface{}

    // This  is  the  maximum  number  of  unanswered Status Enquiries the
    // equipment shall accept be- fore declaring the interface down. The type is
    // interface{} with range: 1..10.
    Frdlcmierrorthreshold interface{}

    // This is the number of status polling intervals over which the error
    // threshold is counted.  For example, if within 'MonitoredEvents' number  of
    // events  the  station  receives 'ErrorThreshold' number of errors, the
    // interface  is  marked  as down. The type is interface{} with range: 1..10.
    Frdlcmimonitoredevents interface{}

    // The maximum number of Virtual Circuits allowed for  this  interface.  
    // Usually dictated by the Frame Relay network.  In response to a SET, if a
    // value less than zero or  higher  than the agent's maximal capability is
    // configured, the agent  should  respond  bad- Value. The type is interface{}
    // with range: -2147483648..2147483647.
    Frdlcmimaxsupportedvcs interface{}

    // This indicates whether the Frame Relay  inter- face is using a multicast
    // service. The type is Frdlcmimulticast.
    Frdlcmimulticast interface{}
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetFilter() yfilter.YFilter { return frdlcmientry.YFilter }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) SetFilter(yf yfilter.YFilter) { frdlcmientry.YFilter = yf }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetGoName(yname string) string {
    if yname == "frDlcmiIfIndex" { return "Frdlcmiifindex" }
    if yname == "frDlcmiState" { return "Frdlcmistate" }
    if yname == "frDlcmiAddress" { return "Frdlcmiaddress" }
    if yname == "frDlcmiAddressLen" { return "Frdlcmiaddresslen" }
    if yname == "frDlcmiPollingInterval" { return "Frdlcmipollinginterval" }
    if yname == "frDlcmiFullEnquiryInterval" { return "Frdlcmifullenquiryinterval" }
    if yname == "frDlcmiErrorThreshold" { return "Frdlcmierrorthreshold" }
    if yname == "frDlcmiMonitoredEvents" { return "Frdlcmimonitoredevents" }
    if yname == "frDlcmiMaxSupportedVCs" { return "Frdlcmimaxsupportedvcs" }
    if yname == "frDlcmiMulticast" { return "Frdlcmimulticast" }
    return ""
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetSegmentPath() string {
    return "frDlcmiEntry" + "[frDlcmiIfIndex='" + fmt.Sprintf("%v", frdlcmientry.Frdlcmiifindex) + "']"
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frDlcmiIfIndex"] = frdlcmientry.Frdlcmiifindex
    leafs["frDlcmiState"] = frdlcmientry.Frdlcmistate
    leafs["frDlcmiAddress"] = frdlcmientry.Frdlcmiaddress
    leafs["frDlcmiAddressLen"] = frdlcmientry.Frdlcmiaddresslen
    leafs["frDlcmiPollingInterval"] = frdlcmientry.Frdlcmipollinginterval
    leafs["frDlcmiFullEnquiryInterval"] = frdlcmientry.Frdlcmifullenquiryinterval
    leafs["frDlcmiErrorThreshold"] = frdlcmientry.Frdlcmierrorthreshold
    leafs["frDlcmiMonitoredEvents"] = frdlcmientry.Frdlcmimonitoredevents
    leafs["frDlcmiMaxSupportedVCs"] = frdlcmientry.Frdlcmimaxsupportedvcs
    leafs["frDlcmiMulticast"] = frdlcmientry.Frdlcmimulticast
    return leafs
}

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetBundleName() string { return "cisco_ios_xe" }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetYangName() string { return "frDlcmiEntry" }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) SetParent(parent types.Entity) { frdlcmientry.parent = parent }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetParent() types.Entity { return frdlcmientry.parent }

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetParentYangName() string { return "frDlcmiTable" }

// RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress represents in use on the Frame Relay interface.
type RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress string

const (
    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q921 RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q921"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922March90 RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922March90"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922November90 RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922November90"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922 RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922"
)

// RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen represents cluding the control portion.
type RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen string

const (
    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_two_octets RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "two-octets"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_three_octets RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "three-octets"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_four_octets RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "four-octets"
)

// RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast represents face is using a multicast service.
type RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast string

const (
    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast_nonBroadcast RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast = "nonBroadcast"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast_broadcast RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast = "broadcast"
)

// RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate represents interface.
type RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate string

const (
    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_noLmiConfigured RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "noLmiConfigured"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_lmiRev1 RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "lmiRev1"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_ansiT1_617_D RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "ansiT1-617-D"

    RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_ansiT1_617_B RFC1315MIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "ansiT1-617-B"
)

// RFC1315MIB_Frcircuittable
// A table containing information about specific Data
// Link Connection Identifiers and corresponding virtual
// circuits.
type RFC1315MIB_Frcircuittable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single  Data  Link Connection Identifier. The
    // type is slice of RFC1315MIB_Frcircuittable_Frcircuitentry.
    Frcircuitentry []RFC1315MIB_Frcircuittable_Frcircuitentry
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetFilter() yfilter.YFilter { return frcircuittable.YFilter }

func (frcircuittable *RFC1315MIB_Frcircuittable) SetFilter(yf yfilter.YFilter) { frcircuittable.YFilter = yf }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetGoName(yname string) string {
    if yname == "frCircuitEntry" { return "Frcircuitentry" }
    return ""
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetSegmentPath() string {
    return "frCircuitTable"
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frCircuitEntry" {
        for _, c := range frcircuittable.Frcircuitentry {
            if frcircuittable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1315MIB_Frcircuittable_Frcircuitentry{}
        frcircuittable.Frcircuitentry = append(frcircuittable.Frcircuitentry, child)
        return &frcircuittable.Frcircuitentry[len(frcircuittable.Frcircuitentry)-1]
    }
    return nil
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frcircuittable.Frcircuitentry {
        children[frcircuittable.Frcircuitentry[i].GetSegmentPath()] = &frcircuittable.Frcircuitentry[i]
    }
    return children
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetBundleName() string { return "cisco_ios_xe" }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetYangName() string { return "frCircuitTable" }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frcircuittable *RFC1315MIB_Frcircuittable) SetParent(parent types.Entity) { frcircuittable.parent = parent }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetParent() types.Entity { return frcircuittable.parent }

func (frcircuittable *RFC1315MIB_Frcircuittable) GetParentYangName() string { return "RFC1315-MIB" }

// RFC1315MIB_Frcircuittable_Frcircuitentry
// The information regarding a single  Data  Link
// Connection Identifier.
type RFC1315MIB_Frcircuittable_Frcircuitentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the ifEntry this  virtual
    // circuit is layered onto. The type is interface{} with range:
    // -2147483648..2147483647.
    Frcircuitifindex interface{}

    // This attribute is a key. The Data Link Connection Identifier  for  this
    // virtual circuit. The type is interface{} with range:
    // -2147483648..2147483647.
    Frcircuitdlci interface{}

    // Indicates whether the particular virtual  cir- cuit  is operational.  In
    // the absence of a Data Link Connection Management  Interface,  virtual
    // circuit  entries  (rows) may be created by set- ting virtual  circuit 
    // state  to  'active',  or deleted by changing Circuit state to 'invalid'.
    // Whether or not the row actually  disappears  is left  to the
    // implementation, so this object may actually read as 'invalid' for  some 
    // arbitrary length  of  time.   It is also legal to set the state of a
    // virtual  circuit  to  'inactive'  to temporarily disable a given circuit.
    // The type is Frcircuitstate.
    Frcircuitstate interface{}

    // Number of frames received from the network in- dicating  forward 
    // congestion since the virtual circuit was created. The type is interface{}
    // with range: 0..4294967295.
    Frcircuitreceivedfecns interface{}

    // Number of frames received from the network in- dicating  backward
    // congestion since the virtual circuit was created. The type is interface{}
    // with range: 0..4294967295.
    Frcircuitreceivedbecns interface{}

    // The number of frames sent  from  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    Frcircuitsentframes interface{}

    // The number of octets sent  from  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    Frcircuitsentoctets interface{}

    // Number of frames received  over  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    Frcircuitreceivedframes interface{}

    // Number of octets received  over  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    Frcircuitreceivedoctets interface{}

    // The value of sysUpTime when the  virtual  cir- cuit was created, whether by
    // the Data Link Con- nection Management Interface  or  by  a  SetRe- quest.
    // The type is interface{} with range: 0..4294967295.
    Frcircuitcreationtime interface{}

    // The value of sysUpTime when last there  was  a change in the virtual
    // circuit state. The type is interface{} with range: 0..4294967295.
    Frcircuitlasttimechange interface{}

    // This variable indicates the maximum amount  of data,  in  bits,  that  the 
    // network  agrees to transfer under normal  conditions,  during  the
    // measurement interval. The type is interface{} with range:
    // -2147483648..2147483647.
    Frcircuitcommittedburst interface{}

    // This variable indicates the maximum amount  of uncommitted data bits that
    // the network will at- tempt to deliver over the measurement interval.  By
    // default, if not configured when creating the entry, the Excess Information
    // Burst Size is set to the value of ifSpeed. The type is interface{} with
    // range: -2147483648..2147483647.
    Frcircuitexcessburst interface{}

    // Throughput is the average number of 'Frame Re- lay  Information  Field' 
    // bits  transferred per second across a user network interface  in  one
    // direction, measured over the measurement inter- val.  If the  configured 
    // committed  burst  rate  and throughput  are  both non-zero, the measurement
    // interval T=frCircuitCommittedBurst/frCircuitThroughput.  If the  configured
    // committed  burst  rate  and throughput  are  both zero, the measurement in-
    // terval        T=frCircuitExcessBurst/ifSpeed. The type is interface{} with
    // range: -2147483648..2147483647.
    Frcircuitthroughput interface{}
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetFilter() yfilter.YFilter { return frcircuitentry.YFilter }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) SetFilter(yf yfilter.YFilter) { frcircuitentry.YFilter = yf }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetGoName(yname string) string {
    if yname == "frCircuitIfIndex" { return "Frcircuitifindex" }
    if yname == "frCircuitDlci" { return "Frcircuitdlci" }
    if yname == "frCircuitState" { return "Frcircuitstate" }
    if yname == "frCircuitReceivedFECNs" { return "Frcircuitreceivedfecns" }
    if yname == "frCircuitReceivedBECNs" { return "Frcircuitreceivedbecns" }
    if yname == "frCircuitSentFrames" { return "Frcircuitsentframes" }
    if yname == "frCircuitSentOctets" { return "Frcircuitsentoctets" }
    if yname == "frCircuitReceivedFrames" { return "Frcircuitreceivedframes" }
    if yname == "frCircuitReceivedOctets" { return "Frcircuitreceivedoctets" }
    if yname == "frCircuitCreationTime" { return "Frcircuitcreationtime" }
    if yname == "frCircuitLastTimeChange" { return "Frcircuitlasttimechange" }
    if yname == "frCircuitCommittedBurst" { return "Frcircuitcommittedburst" }
    if yname == "frCircuitExcessBurst" { return "Frcircuitexcessburst" }
    if yname == "frCircuitThroughput" { return "Frcircuitthroughput" }
    return ""
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetSegmentPath() string {
    return "frCircuitEntry" + "[frCircuitIfIndex='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitifindex) + "']" + "[frCircuitDlci='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitdlci) + "']"
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frCircuitIfIndex"] = frcircuitentry.Frcircuitifindex
    leafs["frCircuitDlci"] = frcircuitentry.Frcircuitdlci
    leafs["frCircuitState"] = frcircuitentry.Frcircuitstate
    leafs["frCircuitReceivedFECNs"] = frcircuitentry.Frcircuitreceivedfecns
    leafs["frCircuitReceivedBECNs"] = frcircuitentry.Frcircuitreceivedbecns
    leafs["frCircuitSentFrames"] = frcircuitentry.Frcircuitsentframes
    leafs["frCircuitSentOctets"] = frcircuitentry.Frcircuitsentoctets
    leafs["frCircuitReceivedFrames"] = frcircuitentry.Frcircuitreceivedframes
    leafs["frCircuitReceivedOctets"] = frcircuitentry.Frcircuitreceivedoctets
    leafs["frCircuitCreationTime"] = frcircuitentry.Frcircuitcreationtime
    leafs["frCircuitLastTimeChange"] = frcircuitentry.Frcircuitlasttimechange
    leafs["frCircuitCommittedBurst"] = frcircuitentry.Frcircuitcommittedburst
    leafs["frCircuitExcessBurst"] = frcircuitentry.Frcircuitexcessburst
    leafs["frCircuitThroughput"] = frcircuitentry.Frcircuitthroughput
    return leafs
}

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetBundleName() string { return "cisco_ios_xe" }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetYangName() string { return "frCircuitEntry" }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) SetParent(parent types.Entity) { frcircuitentry.parent = parent }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetParent() types.Entity { return frcircuitentry.parent }

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetParentYangName() string { return "frCircuitTable" }

// RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate represents temporarily disable a given circuit.
type RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate string

const (
    RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate_invalid RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "invalid"

    RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate_active RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "active"

    RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate_inactive RFC1315MIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "inactive"
)

// RFC1315MIB_Frerrtable
// A table containing information about Errors on the
// Frame Relay interface.
type RFC1315MIB_Frerrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of RFC1315MIB_Frerrtable_Frerrentry.
    Frerrentry []RFC1315MIB_Frerrtable_Frerrentry
}

func (frerrtable *RFC1315MIB_Frerrtable) GetFilter() yfilter.YFilter { return frerrtable.YFilter }

func (frerrtable *RFC1315MIB_Frerrtable) SetFilter(yf yfilter.YFilter) { frerrtable.YFilter = yf }

func (frerrtable *RFC1315MIB_Frerrtable) GetGoName(yname string) string {
    if yname == "frErrEntry" { return "Frerrentry" }
    return ""
}

func (frerrtable *RFC1315MIB_Frerrtable) GetSegmentPath() string {
    return "frErrTable"
}

func (frerrtable *RFC1315MIB_Frerrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frErrEntry" {
        for _, c := range frerrtable.Frerrentry {
            if frerrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := RFC1315MIB_Frerrtable_Frerrentry{}
        frerrtable.Frerrentry = append(frerrtable.Frerrentry, child)
        return &frerrtable.Frerrentry[len(frerrtable.Frerrentry)-1]
    }
    return nil
}

func (frerrtable *RFC1315MIB_Frerrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frerrtable.Frerrentry {
        children[frerrtable.Frerrentry[i].GetSegmentPath()] = &frerrtable.Frerrentry[i]
    }
    return children
}

func (frerrtable *RFC1315MIB_Frerrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frerrtable *RFC1315MIB_Frerrtable) GetBundleName() string { return "cisco_ios_xe" }

func (frerrtable *RFC1315MIB_Frerrtable) GetYangName() string { return "frErrTable" }

func (frerrtable *RFC1315MIB_Frerrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frerrtable *RFC1315MIB_Frerrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frerrtable *RFC1315MIB_Frerrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frerrtable *RFC1315MIB_Frerrtable) SetParent(parent types.Entity) { frerrtable.parent = parent }

func (frerrtable *RFC1315MIB_Frerrtable) GetParent() types.Entity { return frerrtable.parent }

func (frerrtable *RFC1315MIB_Frerrtable) GetParentYangName() string { return "RFC1315-MIB" }

// RFC1315MIB_Frerrtable_Frerrentry
// The error information for a single frame relay
// interface.
type RFC1315MIB_Frerrtable_Frerrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the  corresponding  ifEn-
    // try. The type is interface{} with range: -2147483648..2147483647.
    Frerrifindex interface{}

    // The type of error that was last seen  on  this interface. The type is
    // Frerrtype.
    Frerrtype interface{}

    // An octet string containing as much of the  er- ror  packet as possible.  As
    // a minimum, it must contain the Q.922 Address or  as  much  as  was
    // delivered.   It is desirable to include all in- formation up to the PDU.
    // The type is string.
    Frerrdata interface{}

    // The value of sysUpTime at which the error  was detected. The type is
    // interface{} with range: 0..4294967295.
    Frerrtime interface{}
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetFilter() yfilter.YFilter { return frerrentry.YFilter }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) SetFilter(yf yfilter.YFilter) { frerrentry.YFilter = yf }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetGoName(yname string) string {
    if yname == "frErrIfIndex" { return "Frerrifindex" }
    if yname == "frErrType" { return "Frerrtype" }
    if yname == "frErrData" { return "Frerrdata" }
    if yname == "frErrTime" { return "Frerrtime" }
    return ""
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetSegmentPath() string {
    return "frErrEntry" + "[frErrIfIndex='" + fmt.Sprintf("%v", frerrentry.Frerrifindex) + "']"
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frErrIfIndex"] = frerrentry.Frerrifindex
    leafs["frErrType"] = frerrentry.Frerrtype
    leafs["frErrData"] = frerrentry.Frerrdata
    leafs["frErrTime"] = frerrentry.Frerrtime
    return leafs
}

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetBundleName() string { return "cisco_ios_xe" }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetYangName() string { return "frErrEntry" }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) SetParent(parent types.Entity) { frerrentry.parent = parent }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetParent() types.Entity { return frerrentry.parent }

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetParentYangName() string { return "frErrTable" }

// RFC1315MIB_Frerrtable_Frerrentry_Frerrtype represents interface.
type RFC1315MIB_Frerrtable_Frerrentry_Frerrtype string

const (
    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_unknownError RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "unknownError"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_receiveShort RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "receiveShort"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_receiveLong RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "receiveLong"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_illegalDLCI RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "illegalDLCI"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_unknownDLCI RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "unknownDLCI"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_dlcmiProtoErr RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiProtoErr"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_dlcmiUnknownIE RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiUnknownIE"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_dlcmiSequenceErr RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiSequenceErr"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_dlcmiUnknownRpt RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiUnknownRpt"

    RFC1315MIB_Frerrtable_Frerrentry_Frerrtype_noErrorSinceReset RFC1315MIB_Frerrtable_Frerrentry_Frerrtype = "noErrorSinceReset"
)

