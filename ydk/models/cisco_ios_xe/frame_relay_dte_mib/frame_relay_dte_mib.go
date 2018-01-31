// The MIB module to describe the use of a Frame Relay
// interface by a DTE.
package frame_relay_dte_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package frame_relay_dte_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:FRAME-RELAY-DTE-MIB FRAME-RELAY-DTE-MIB}", reflect.TypeOf(FRAMERELAYDTEMIB{}))
    ydk.RegisterEntity("FRAME-RELAY-DTE-MIB:FRAME-RELAY-DTE-MIB", reflect.TypeOf(FRAMERELAYDTEMIB{}))
}

// FRAMERELAYDTEMIB
type FRAMERELAYDTEMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Framerelaytrapcontrol FRAMERELAYDTEMIB_Framerelaytrapcontrol

    // The Parameters for the Data Link Connection Management Interface for the
    // frame relay service on this interface.
    Frdlcmitable FRAMERELAYDTEMIB_Frdlcmitable

    // A table containing information about specific Data Link Connections (DLC)
    // or virtual circuits.
    Frcircuittable FRAMERELAYDTEMIB_Frcircuittable

    // A table containing information about Errors on the Frame Relay interface. 
    // Discontinuities in the counters contained in this table are the same as
    // apply to the ifEntry associated with the Interface.
    Frerrtable FRAMERELAYDTEMIB_Frerrtable
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetFilter() yfilter.YFilter { return fRAMERELAYDTEMIB.YFilter }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) SetFilter(yf yfilter.YFilter) { fRAMERELAYDTEMIB.YFilter = yf }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetGoName(yname string) string {
    if yname == "frameRelayTrapControl" { return "Framerelaytrapcontrol" }
    if yname == "frDlcmiTable" { return "Frdlcmitable" }
    if yname == "frCircuitTable" { return "Frcircuittable" }
    if yname == "frErrTable" { return "Frerrtable" }
    return ""
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetSegmentPath() string {
    return "FRAME-RELAY-DTE-MIB:FRAME-RELAY-DTE-MIB"
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frameRelayTrapControl" {
        return &fRAMERELAYDTEMIB.Framerelaytrapcontrol
    }
    if childYangName == "frDlcmiTable" {
        return &fRAMERELAYDTEMIB.Frdlcmitable
    }
    if childYangName == "frCircuitTable" {
        return &fRAMERELAYDTEMIB.Frcircuittable
    }
    if childYangName == "frErrTable" {
        return &fRAMERELAYDTEMIB.Frerrtable
    }
    return nil
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frameRelayTrapControl"] = &fRAMERELAYDTEMIB.Framerelaytrapcontrol
    children["frDlcmiTable"] = &fRAMERELAYDTEMIB.Frdlcmitable
    children["frCircuitTable"] = &fRAMERELAYDTEMIB.Frcircuittable
    children["frErrTable"] = &fRAMERELAYDTEMIB.Frerrtable
    return children
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetYangName() string { return "FRAME-RELAY-DTE-MIB" }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) SetParent(parent types.Entity) { fRAMERELAYDTEMIB.parent = parent }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetParent() types.Entity { return fRAMERELAYDTEMIB.parent }

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetParentYangName() string { return "FRAME-RELAY-DTE-MIB" }

// FRAMERELAYDTEMIB_Framerelaytrapcontrol
type FRAMERELAYDTEMIB_Framerelaytrapcontrol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the frDLCIStatusChange
    // trap. The type is Frtrapstate.
    Frtrapstate interface{}

    // This variable indicates the number of milliseconds that must elapse between
    // trap emissions.  If events occur more rapidly, the impementation may simply
    // fail to trap, or may queue traps until an appropriate time. The type is
    // interface{} with range: 0..3600000.
    Frtrapmaxrate interface{}
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetFilter() yfilter.YFilter { return framerelaytrapcontrol.YFilter }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) SetFilter(yf yfilter.YFilter) { framerelaytrapcontrol.YFilter = yf }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetGoName(yname string) string {
    if yname == "frTrapState" { return "Frtrapstate" }
    if yname == "frTrapMaxRate" { return "Frtrapmaxrate" }
    return ""
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetSegmentPath() string {
    return "frameRelayTrapControl"
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frTrapState"] = framerelaytrapcontrol.Frtrapstate
    leafs["frTrapMaxRate"] = framerelaytrapcontrol.Frtrapmaxrate
    return leafs
}

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetBundleName() string { return "cisco_ios_xe" }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetYangName() string { return "frameRelayTrapControl" }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) SetParent(parent types.Entity) { framerelaytrapcontrol.parent = parent }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetParent() types.Entity { return framerelaytrapcontrol.parent }

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetParentYangName() string { return "FRAME-RELAY-DTE-MIB" }

// FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate represents the frDLCIStatusChange trap.
type FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate string

const (
    FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate_enabled FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate = "enabled"

    FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate_disabled FRAMERELAYDTEMIB_Framerelaytrapcontrol_Frtrapstate = "disabled"
)

// FRAMERELAYDTEMIB_Frdlcmitable
// The Parameters for the Data Link Connection Management
// Interface for the frame relay service on this
// interface.
type FRAMERELAYDTEMIB_Frdlcmitable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Connection Management Interface.
    // The type is slice of FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry.
    Frdlcmientry []FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetFilter() yfilter.YFilter { return frdlcmitable.YFilter }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) SetFilter(yf yfilter.YFilter) { frdlcmitable.YFilter = yf }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetGoName(yname string) string {
    if yname == "frDlcmiEntry" { return "Frdlcmientry" }
    return ""
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetSegmentPath() string {
    return "frDlcmiTable"
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frDlcmiEntry" {
        for _, c := range frdlcmitable.Frdlcmientry {
            if frdlcmitable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry{}
        frdlcmitable.Frdlcmientry = append(frdlcmitable.Frdlcmientry, child)
        return &frdlcmitable.Frdlcmientry[len(frdlcmitable.Frdlcmientry)-1]
    }
    return nil
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frdlcmitable.Frdlcmientry {
        children[frdlcmitable.Frdlcmientry[i].GetSegmentPath()] = &frdlcmitable.Frdlcmientry[i]
    }
    return children
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetBundleName() string { return "cisco_ios_xe" }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetYangName() string { return "frDlcmiTable" }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) SetParent(parent types.Entity) { frdlcmitable.parent = parent }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetParent() types.Entity { return frdlcmitable.parent }

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetParentYangName() string { return "FRAME-RELAY-DTE-MIB" }

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry
// The Parameters for a particular Data Link Connection
// Management Interface.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the corresponding ifEntry.
    // The type is interface{} with range: 1..2147483647.
    Frdlcmiifindex interface{}

    // This variable states which Data Link Connection Management scheme is active
    // (and by implication, what DLCI it uses) on the Frame Relay interface. The
    // type is Frdlcmistate.
    Frdlcmistate interface{}

    // This variable states which address format is in use on the Frame Relay
    // interface. The type is Frdlcmiaddress.
    Frdlcmiaddress interface{}

    // This variable states the address length in octets.  In the case of Q922
    // format, the length indicates the entire length of the address including the
    // control portion. The type is Frdlcmiaddresslen.
    Frdlcmiaddresslen interface{}

    // This is the number of seconds between successive status enquiry messages.
    // The type is interface{} with range: 5..30. Units are seconds.
    Frdlcmipollinginterval interface{}

    // Number of status enquiry intervals that pass before issuance of a full
    // status enquiry message. The type is interface{} with range: 1..255.
    Frdlcmifullenquiryinterval interface{}

    // This is the maximum number of unanswered Status Enquiries the equipment
    // shall accept before declaring the interface down. The type is interface{}
    // with range: 1..10.
    Frdlcmierrorthreshold interface{}

    // This is the number of status polling intervals over which the error
    // threshold is counted.  For example, if within 'MonitoredEvents' number of
    // events the station receives 'ErrorThreshold' number of errors, the
    // interface is marked as down. The type is interface{} with range: 1..10.
    Frdlcmimonitoredevents interface{}

    // The maximum number of Virtual Circuits allowed for this interface.  Usually
    // dictated by the Frame Relay network.  In response to a SET, if a value less
    // than zero or higher than the agent's maximal capability is configured, the
    // agent should respond badValue. The type is interface{} with range:
    // 0..8388607.
    Frdlcmimaxsupportedvcs interface{}

    // This indicates whether the Frame Relay interface is using a multicast
    // service. The type is Frdlcmimulticast.
    Frdlcmimulticast interface{}

    // This indicates the status of the Frame Relay interface as determined by the
    // performance of the dlcmi.  If no dlcmi is running, the Frame Relay
    // interface will stay in the running state indefinitely. The type is
    // Frdlcmistatus.
    Frdlcmistatus interface{}

    // SNMP Version 2 Row Status Variable.  Writable objects in the table may be
    // written in any RowStatus state. The type is RowStatus.
    Frdlcmirowstatus interface{}
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetFilter() yfilter.YFilter { return frdlcmientry.YFilter }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) SetFilter(yf yfilter.YFilter) { frdlcmientry.YFilter = yf }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetGoName(yname string) string {
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
    if yname == "frDlcmiStatus" { return "Frdlcmistatus" }
    if yname == "frDlcmiRowStatus" { return "Frdlcmirowstatus" }
    return ""
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetSegmentPath() string {
    return "frDlcmiEntry" + "[frDlcmiIfIndex='" + fmt.Sprintf("%v", frdlcmientry.Frdlcmiifindex) + "']"
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetLeafs() map[string]interface{} {
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
    leafs["frDlcmiStatus"] = frdlcmientry.Frdlcmistatus
    leafs["frDlcmiRowStatus"] = frdlcmientry.Frdlcmirowstatus
    return leafs
}

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetBundleName() string { return "cisco_ios_xe" }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetYangName() string { return "frDlcmiEntry" }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) SetParent(parent types.Entity) { frdlcmientry.parent = parent }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetParent() types.Entity { return frdlcmientry.parent }

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetParentYangName() string { return "frDlcmiTable" }

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress represents the Frame Relay interface.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress string

const (
    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q921 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q921"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922March90 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922March90"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922November90 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922November90"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress_q922 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddress = "q922"
)

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen represents portion.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen string

const (
    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_twoOctets FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "twoOctets"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_threeOctets FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "threeOctets"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen_fourOctets FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmiaddresslen = "fourOctets"
)

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast represents using a multicast service.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast string

const (
    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast_nonBroadcast FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast = "nonBroadcast"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast_broadcast FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmimulticast = "broadcast"
)

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate represents DLCI it uses) on the Frame Relay interface.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate string

const (
    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_noLmiConfigured FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "noLmiConfigured"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_lmiRev1 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "lmiRev1"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_ansiT1617D FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "ansiT1617D"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_ansiT1617B FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "ansiT1617B"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_itut933A FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "itut933A"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate_ansiT1617D1994 FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistate = "ansiT1617D1994"
)

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus represents in the running state indefinitely.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus string

const (
    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus_running FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus = "running"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus_fault FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus = "fault"

    FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus_initializing FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry_Frdlcmistatus = "initializing"
)

// FRAMERELAYDTEMIB_Frcircuittable
// A table containing information about specific Data
// Link Connections (DLC) or virtual circuits.
type FRAMERELAYDTEMIB_Frcircuittable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The information regarding a single Data Link Connection.  Discontinuities
    // in the counters contained in this table are indicated by the value in
    // frCircuitCreationTime. The type is slice of
    // FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry.
    Frcircuitentry []FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetFilter() yfilter.YFilter { return frcircuittable.YFilter }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) SetFilter(yf yfilter.YFilter) { frcircuittable.YFilter = yf }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetGoName(yname string) string {
    if yname == "frCircuitEntry" { return "Frcircuitentry" }
    return ""
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetSegmentPath() string {
    return "frCircuitTable"
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frCircuitEntry" {
        for _, c := range frcircuittable.Frcircuitentry {
            if frcircuittable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry{}
        frcircuittable.Frcircuitentry = append(frcircuittable.Frcircuitentry, child)
        return &frcircuittable.Frcircuitentry[len(frcircuittable.Frcircuitentry)-1]
    }
    return nil
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frcircuittable.Frcircuitentry {
        children[frcircuittable.Frcircuitentry[i].GetSegmentPath()] = &frcircuittable.Frcircuitentry[i]
    }
    return children
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetBundleName() string { return "cisco_ios_xe" }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetYangName() string { return "frCircuitTable" }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) SetParent(parent types.Entity) { frcircuittable.parent = parent }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetParent() types.Entity { return frcircuittable.parent }

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetParentYangName() string { return "FRAME-RELAY-DTE-MIB" }

// FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry
// The information regarding a single Data Link
// Connection.  Discontinuities in the counters contained
// in this table are indicated by the value in
// frCircuitCreationTime.
type FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the ifEntry this virtual
    // circuit is layered onto. The type is interface{} with range: 1..2147483647.
    Frcircuitifindex interface{}

    // This attribute is a key. The Data Link Connection Identifier for this
    // virtual circuit. The type is interface{} with range: 0..8388607.
    Frcircuitdlci interface{}

    // Indicates whether the particular virtual circuit is operational.  In the
    // absence of a Data Link Connection Management Interface, virtual circuit
    // entries (rows) may be created by setting virtual circuit state to 'active',
    // or deleted by changing Circuit state to 'invalid'.  Whether or not the row
    // actually disappears is left to the implementation, so this object may
    // actually read as 'invalid' for some arbitrary length of time.  It is also
    // legal to set the state of a virtual circuit to 'inactive' to temporarily
    // disable a given circuit.  The use of 'invalid' is deprecated in this SNMP
    // Version 2 MIB, in favor of frCircuitRowStatus. The type is Frcircuitstate.
    Frcircuitstate interface{}

    // Number of frames received from the network indicating forward congestion
    // since the virtual circuit was created.  This occurs when the remote DTE
    // sets the FECN flag, or when a switch in the network enqueues the frame to a
    // trunk whose transmission queue is congested. The type is interface{} with
    // range: 0..4294967295.
    Frcircuitreceivedfecns interface{}

    // Number of frames received from the network indicating backward congestion
    // since the virtual circuit was created.  This occurs when the remote DTE
    // sets the BECN flag, or when a switch in the network receives the frame from
    // a trunk whose transmission queue is congested. The type is interface{} with
    // range: 0..4294967295.
    Frcircuitreceivedbecns interface{}

    // The number of frames sent from this virtual circuit since it was created.
    // The type is interface{} with range: 0..4294967295.
    Frcircuitsentframes interface{}

    // The number of octets sent from this virtual circuit since it was created. 
    // Octets counted are the full frame relay header and the payload, but do not
    // include the flag characters or CRC. The type is interface{} with range:
    // 0..4294967295.
    Frcircuitsentoctets interface{}

    // Number of frames received over this virtual circuit since it was created.
    // The type is interface{} with range: 0..4294967295.
    Frcircuitreceivedframes interface{}

    // Number of octets received over this virtual circuit since it was created. 
    // Octets counted include the full frame relay header, but do not include the
    // flag characters or the CRC. The type is interface{} with range:
    // 0..4294967295.
    Frcircuitreceivedoctets interface{}

    // The value of sysUpTime when the virtual circuit was created, whether by the
    // Data Link Connection Management Interface or by a SetRequest. The type is
    // interface{} with range: 0..4294967295.
    Frcircuitcreationtime interface{}

    // The value of sysUpTime when last there was a change in the virtual circuit
    // state. The type is interface{} with range: 0..4294967295.
    Frcircuitlasttimechange interface{}

    // This variable indicates the maximum amount of data, in bits, that the
    // network agrees to transfer under normal conditions, during the measurement
    // interval. The type is interface{} with range: 0..2147483647.
    Frcircuitcommittedburst interface{}

    // This variable indicates the maximum amount of uncommitted data bits that
    // the network will attempt to deliver over the measurement interval.  By
    // default, if not configured when creating the entry, the Excess Information
    // Burst Size is set to the value of ifSpeed. The type is interface{} with
    // range: 0..2147483647.
    Frcircuitexcessburst interface{}

    // Throughput is the average number of 'Frame Relay Information Field' bits
    // transferred per second across a user network interface in one direction,
    // measured over the measurement interval.  If the configured committed burst
    // rate and throughput are both non-zero, the measurement interval, T, is    
    // T=frCircuitCommittedBurst/frCircuitThroughput.  If the configured committed
    // burst rate and throughput are both zero, the measurement interval, T, is   
    // T=frCircuitExcessBurst/ifSpeed. The type is interface{} with range:
    // 0..2147483647.
    Frcircuitthroughput interface{}

    // This indicates whether this VC is used as a unicast VC (i.e. not multicast)
    // or the type of multicast service subscribed to. The type is
    // Frcircuitmulticast.
    Frcircuitmulticast interface{}

    // Indication of whether the VC was manually created (static), or dynamically
    // created (dynamic) via the data link control management interface. The type
    // is Frcircuittype.
    Frcircuittype interface{}

    // The number of inbound frames dropped because of format errors, or because
    // the VC is inactive. The type is interface{} with range: 0..4294967295.
    Frcircuitdiscards interface{}

    // Number of frames received from the network indicating that they were
    // eligible for discard since the virtual circuit was created.  This occurs
    // when the remote DTE sets the DE flag, or when in remote DTE's switch
    // detects that the frame was received as Excess Burst data. The type is
    // interface{} with range: 0..4294967295.
    Frcircuitreceiveddes interface{}

    // Number of frames sent to the network indicating that they were eligible for
    // discard since the virtual circuit was created.   This occurs when the local
    // DTE sets the DE flag, indicating that during Network congestion situations
    // those frames should be discarded in preference of other frames sent without
    // the DE bit set. The type is interface{} with range: 0..4294967295.
    Frcircuitsentdes interface{}

    // Normally the same value as frDlcmiIfIndex, but different when an
    // implementation associates a virtual ifEntry with a DLC or set of DLCs in
    // order to associate higher layer objects such as the ipAddrEntry with a
    // subset of the virtual circuits on a Frame Relay interface. The type of such
    // ifEntries is defined by the higher layer object; for example, if PPP/Frame
    // Relay is implemented, the ifType of this ifEntry would be PPP. If it is not
    // so defined, as would be the case with an ipAddrEntry, it should be of type
    // Other. The type is interface{} with range: 1..2147483647.
    Frcircuitlogicalifindex interface{}

    // This object is used to create a new row or modify or destroy an existing
    // row in the manner described in the definition of the RowStatus textual
    // convention. Writable objects in the table may be written in any RowStatus
    // state. The type is RowStatus.
    Frcircuitrowstatus interface{}
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetFilter() yfilter.YFilter { return frcircuitentry.YFilter }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) SetFilter(yf yfilter.YFilter) { frcircuitentry.YFilter = yf }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetGoName(yname string) string {
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
    if yname == "frCircuitMulticast" { return "Frcircuitmulticast" }
    if yname == "frCircuitType" { return "Frcircuittype" }
    if yname == "frCircuitDiscards" { return "Frcircuitdiscards" }
    if yname == "frCircuitReceivedDEs" { return "Frcircuitreceiveddes" }
    if yname == "frCircuitSentDEs" { return "Frcircuitsentdes" }
    if yname == "frCircuitLogicalIfIndex" { return "Frcircuitlogicalifindex" }
    if yname == "frCircuitRowStatus" { return "Frcircuitrowstatus" }
    return ""
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetSegmentPath() string {
    return "frCircuitEntry" + "[frCircuitIfIndex='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitifindex) + "']" + "[frCircuitDlci='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitdlci) + "']"
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetLeafs() map[string]interface{} {
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
    leafs["frCircuitMulticast"] = frcircuitentry.Frcircuitmulticast
    leafs["frCircuitType"] = frcircuitentry.Frcircuittype
    leafs["frCircuitDiscards"] = frcircuitentry.Frcircuitdiscards
    leafs["frCircuitReceivedDEs"] = frcircuitentry.Frcircuitreceiveddes
    leafs["frCircuitSentDEs"] = frcircuitentry.Frcircuitsentdes
    leafs["frCircuitLogicalIfIndex"] = frcircuitentry.Frcircuitlogicalifindex
    leafs["frCircuitRowStatus"] = frcircuitentry.Frcircuitrowstatus
    return leafs
}

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetBundleName() string { return "cisco_ios_xe" }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetYangName() string { return "frCircuitEntry" }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) SetParent(parent types.Entity) { frcircuitentry.parent = parent }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetParent() types.Entity { return frcircuitentry.parent }

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetParentYangName() string { return "frCircuitTable" }

// FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast represents subscribed to
type FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast string

const (
    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast_unicast FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast = "unicast"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast_oneWay FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast = "oneWay"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast_twoWay FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast = "twoWay"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast_nWay FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitmulticast = "nWay"
)

// FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate represents 2 MIB, in favor of frCircuitRowStatus.
type FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate string

const (
    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate_invalid FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "invalid"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate_active FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "active"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate_inactive FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuitstate = "inactive"
)

// FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype represents link control management interface.
type FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype string

const (
    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype_static FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype = "static"

    FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype_dynamic FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry_Frcircuittype = "dynamic"
)

// FRAMERELAYDTEMIB_Frerrtable
// A table containing information about Errors on the
// Frame Relay interface.  Discontinuities in the counters
// contained in this table are the same as apply to the
// ifEntry associated with the Interface.
type FRAMERELAYDTEMIB_Frerrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of FRAMERELAYDTEMIB_Frerrtable_Frerrentry.
    Frerrentry []FRAMERELAYDTEMIB_Frerrtable_Frerrentry
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetFilter() yfilter.YFilter { return frerrtable.YFilter }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) SetFilter(yf yfilter.YFilter) { frerrtable.YFilter = yf }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetGoName(yname string) string {
    if yname == "frErrEntry" { return "Frerrentry" }
    return ""
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetSegmentPath() string {
    return "frErrTable"
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frErrEntry" {
        for _, c := range frerrtable.Frerrentry {
            if frerrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FRAMERELAYDTEMIB_Frerrtable_Frerrentry{}
        frerrtable.Frerrentry = append(frerrtable.Frerrentry, child)
        return &frerrtable.Frerrentry[len(frerrtable.Frerrentry)-1]
    }
    return nil
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frerrtable.Frerrentry {
        children[frerrtable.Frerrentry[i].GetSegmentPath()] = &frerrtable.Frerrentry[i]
    }
    return children
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetBundleName() string { return "cisco_ios_xe" }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetYangName() string { return "frErrTable" }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) SetParent(parent types.Entity) { frerrtable.parent = parent }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetParent() types.Entity { return frerrtable.parent }

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetParentYangName() string { return "FRAME-RELAY-DTE-MIB" }

// FRAMERELAYDTEMIB_Frerrtable_Frerrentry
// The error information for a single frame relay
// interface.
type FRAMERELAYDTEMIB_Frerrtable_Frerrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the corresponding ifEntry.
    // The type is interface{} with range: 1..2147483647.
    Frerrifindex interface{}

    // The type of error that was last seen  on  this interface:  receiveShort:
    // frame was not long enough to allow         demultiplexing - the address
    // field was incomplete,         or for virtual circuits using Multiprotocol
    // over         Frame Relay, the protocol identifier was missing         or
    // incomplete.  receiveLong: frame exceeded maximum length configured for this
    // interface.  illegalAddress: address field did not match configured format. 
    // unknownAddress: frame received on a virtual circuit which was not          
    // active or administratively disabled.  dlcmiProtoErr: unspecified error
    // occurred when attempting to                interpret link maintenance
    // frame.  dlcmiUnknownIE: link maintenance frame contained an Information    
    // Element type which is not valid for the                 configured link
    // maintenance protocol.  dlcmiSequenceErr: link maintenance frame contained a
    // sequence                   number other than the expected value. 
    // dlcmiUnknownRpt: link maintenance frame contained a Report Type            
    // Information Element whose value was not valid                  for the
    // configured link maintenance protocol.  noErrorSinceReset: no errors have
    // been detected since the last                    cold start or warm start.
    // The type is Frerrtype.
    Frerrtype interface{}

    // An octet string containing as much of the error packet as possible.  As a
    // minimum, it must contain the Q.922 Address or as much as was delivered.  It
    // is desirable to include all header and demultiplexing information. The type
    // is string with length: 1..1600.
    Frerrdata interface{}

    // The value of sysUpTime at which the error was detected. The type is
    // interface{} with range: 0..4294967295.
    Frerrtime interface{}

    // The number of times the interface has gone down since it was initialized.
    // The type is interface{} with range: 0..4294967295.
    Frerrfaults interface{}

    // The value of sysUpTime at the time when the interface was taken down due to
    // excessive errors.  Excessive errors is defined as the time when a DLCMI
    // exceeds the frDlcmiErrorThreshold number of errors within
    // frDlcmiMonitoredEvents. See FrDlcmiEntry for further details. The type is
    // interface{} with range: 0..4294967295.
    Frerrfaulttime interface{}
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetFilter() yfilter.YFilter { return frerrentry.YFilter }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) SetFilter(yf yfilter.YFilter) { frerrentry.YFilter = yf }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetGoName(yname string) string {
    if yname == "frErrIfIndex" { return "Frerrifindex" }
    if yname == "frErrType" { return "Frerrtype" }
    if yname == "frErrData" { return "Frerrdata" }
    if yname == "frErrTime" { return "Frerrtime" }
    if yname == "frErrFaults" { return "Frerrfaults" }
    if yname == "frErrFaultTime" { return "Frerrfaulttime" }
    return ""
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetSegmentPath() string {
    return "frErrEntry" + "[frErrIfIndex='" + fmt.Sprintf("%v", frerrentry.Frerrifindex) + "']"
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frErrIfIndex"] = frerrentry.Frerrifindex
    leafs["frErrType"] = frerrentry.Frerrtype
    leafs["frErrData"] = frerrentry.Frerrdata
    leafs["frErrTime"] = frerrentry.Frerrtime
    leafs["frErrFaults"] = frerrentry.Frerrfaults
    leafs["frErrFaultTime"] = frerrentry.Frerrfaulttime
    return leafs
}

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetBundleName() string { return "cisco_ios_xe" }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetYangName() string { return "frErrEntry" }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) SetParent(parent types.Entity) { frerrentry.parent = parent }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetParent() types.Entity { return frerrentry.parent }

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetParentYangName() string { return "frErrTable" }

// FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype represents                    cold start or warm start.
type FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype string

const (
    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_unknownError FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "unknownError"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_receiveShort FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "receiveShort"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_receiveLong FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "receiveLong"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_illegalAddress FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "illegalAddress"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_unknownAddress FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "unknownAddress"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_dlcmiProtoErr FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiProtoErr"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_dlcmiUnknownIE FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiUnknownIE"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_dlcmiSequenceErr FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiSequenceErr"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_dlcmiUnknownRpt FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "dlcmiUnknownRpt"

    FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype_noErrorSinceReset FRAMERELAYDTEMIB_Frerrtable_Frerrentry_Frerrtype = "noErrorSinceReset"
)

