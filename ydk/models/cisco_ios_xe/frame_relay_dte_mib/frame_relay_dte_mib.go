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
    EntityData types.CommonEntityData
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

func (fRAMERELAYDTEMIB *FRAMERELAYDTEMIB) GetEntityData() *types.CommonEntityData {
    fRAMERELAYDTEMIB.EntityData.YFilter = fRAMERELAYDTEMIB.YFilter
    fRAMERELAYDTEMIB.EntityData.YangName = "FRAME-RELAY-DTE-MIB"
    fRAMERELAYDTEMIB.EntityData.BundleName = "cisco_ios_xe"
    fRAMERELAYDTEMIB.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    fRAMERELAYDTEMIB.EntityData.SegmentPath = "FRAME-RELAY-DTE-MIB:FRAME-RELAY-DTE-MIB"
    fRAMERELAYDTEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fRAMERELAYDTEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fRAMERELAYDTEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fRAMERELAYDTEMIB.EntityData.Children = make(map[string]types.YChild)
    fRAMERELAYDTEMIB.EntityData.Children["frameRelayTrapControl"] = types.YChild{"Framerelaytrapcontrol", &fRAMERELAYDTEMIB.Framerelaytrapcontrol}
    fRAMERELAYDTEMIB.EntityData.Children["frDlcmiTable"] = types.YChild{"Frdlcmitable", &fRAMERELAYDTEMIB.Frdlcmitable}
    fRAMERELAYDTEMIB.EntityData.Children["frCircuitTable"] = types.YChild{"Frcircuittable", &fRAMERELAYDTEMIB.Frcircuittable}
    fRAMERELAYDTEMIB.EntityData.Children["frErrTable"] = types.YChild{"Frerrtable", &fRAMERELAYDTEMIB.Frerrtable}
    fRAMERELAYDTEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fRAMERELAYDTEMIB.EntityData)
}

// FRAMERELAYDTEMIB_Framerelaytrapcontrol
type FRAMERELAYDTEMIB_Framerelaytrapcontrol struct {
    EntityData types.CommonEntityData
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

func (framerelaytrapcontrol *FRAMERELAYDTEMIB_Framerelaytrapcontrol) GetEntityData() *types.CommonEntityData {
    framerelaytrapcontrol.EntityData.YFilter = framerelaytrapcontrol.YFilter
    framerelaytrapcontrol.EntityData.YangName = "frameRelayTrapControl"
    framerelaytrapcontrol.EntityData.BundleName = "cisco_ios_xe"
    framerelaytrapcontrol.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    framerelaytrapcontrol.EntityData.SegmentPath = "frameRelayTrapControl"
    framerelaytrapcontrol.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    framerelaytrapcontrol.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    framerelaytrapcontrol.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    framerelaytrapcontrol.EntityData.Children = make(map[string]types.YChild)
    framerelaytrapcontrol.EntityData.Leafs = make(map[string]types.YLeaf)
    framerelaytrapcontrol.EntityData.Leafs["frTrapState"] = types.YLeaf{"Frtrapstate", framerelaytrapcontrol.Frtrapstate}
    framerelaytrapcontrol.EntityData.Leafs["frTrapMaxRate"] = types.YLeaf{"Frtrapmaxrate", framerelaytrapcontrol.Frtrapmaxrate}
    return &(framerelaytrapcontrol.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Connection Management Interface.
    // The type is slice of FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry.
    Frdlcmientry []FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry
}

func (frdlcmitable *FRAMERELAYDTEMIB_Frdlcmitable) GetEntityData() *types.CommonEntityData {
    frdlcmitable.EntityData.YFilter = frdlcmitable.YFilter
    frdlcmitable.EntityData.YangName = "frDlcmiTable"
    frdlcmitable.EntityData.BundleName = "cisco_ios_xe"
    frdlcmitable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frdlcmitable.EntityData.SegmentPath = "frDlcmiTable"
    frdlcmitable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frdlcmitable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frdlcmitable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frdlcmitable.EntityData.Children = make(map[string]types.YChild)
    frdlcmitable.EntityData.Children["frDlcmiEntry"] = types.YChild{"Frdlcmientry", nil}
    for i := range frdlcmitable.Frdlcmientry {
        frdlcmitable.EntityData.Children[types.GetSegmentPath(&frdlcmitable.Frdlcmientry[i])] = types.YChild{"Frdlcmientry", &frdlcmitable.Frdlcmientry[i]}
    }
    frdlcmitable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frdlcmitable.EntityData)
}

// FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry
// The Parameters for a particular Data Link Connection
// Management Interface.
type FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry struct {
    EntityData types.CommonEntityData
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

func (frdlcmientry *FRAMERELAYDTEMIB_Frdlcmitable_Frdlcmientry) GetEntityData() *types.CommonEntityData {
    frdlcmientry.EntityData.YFilter = frdlcmientry.YFilter
    frdlcmientry.EntityData.YangName = "frDlcmiEntry"
    frdlcmientry.EntityData.BundleName = "cisco_ios_xe"
    frdlcmientry.EntityData.ParentYangName = "frDlcmiTable"
    frdlcmientry.EntityData.SegmentPath = "frDlcmiEntry" + "[frDlcmiIfIndex='" + fmt.Sprintf("%v", frdlcmientry.Frdlcmiifindex) + "']"
    frdlcmientry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frdlcmientry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frdlcmientry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frdlcmientry.EntityData.Children = make(map[string]types.YChild)
    frdlcmientry.EntityData.Leafs = make(map[string]types.YLeaf)
    frdlcmientry.EntityData.Leafs["frDlcmiIfIndex"] = types.YLeaf{"Frdlcmiifindex", frdlcmientry.Frdlcmiifindex}
    frdlcmientry.EntityData.Leafs["frDlcmiState"] = types.YLeaf{"Frdlcmistate", frdlcmientry.Frdlcmistate}
    frdlcmientry.EntityData.Leafs["frDlcmiAddress"] = types.YLeaf{"Frdlcmiaddress", frdlcmientry.Frdlcmiaddress}
    frdlcmientry.EntityData.Leafs["frDlcmiAddressLen"] = types.YLeaf{"Frdlcmiaddresslen", frdlcmientry.Frdlcmiaddresslen}
    frdlcmientry.EntityData.Leafs["frDlcmiPollingInterval"] = types.YLeaf{"Frdlcmipollinginterval", frdlcmientry.Frdlcmipollinginterval}
    frdlcmientry.EntityData.Leafs["frDlcmiFullEnquiryInterval"] = types.YLeaf{"Frdlcmifullenquiryinterval", frdlcmientry.Frdlcmifullenquiryinterval}
    frdlcmientry.EntityData.Leafs["frDlcmiErrorThreshold"] = types.YLeaf{"Frdlcmierrorthreshold", frdlcmientry.Frdlcmierrorthreshold}
    frdlcmientry.EntityData.Leafs["frDlcmiMonitoredEvents"] = types.YLeaf{"Frdlcmimonitoredevents", frdlcmientry.Frdlcmimonitoredevents}
    frdlcmientry.EntityData.Leafs["frDlcmiMaxSupportedVCs"] = types.YLeaf{"Frdlcmimaxsupportedvcs", frdlcmientry.Frdlcmimaxsupportedvcs}
    frdlcmientry.EntityData.Leafs["frDlcmiMulticast"] = types.YLeaf{"Frdlcmimulticast", frdlcmientry.Frdlcmimulticast}
    frdlcmientry.EntityData.Leafs["frDlcmiStatus"] = types.YLeaf{"Frdlcmistatus", frdlcmientry.Frdlcmistatus}
    frdlcmientry.EntityData.Leafs["frDlcmiRowStatus"] = types.YLeaf{"Frdlcmirowstatus", frdlcmientry.Frdlcmirowstatus}
    return &(frdlcmientry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Data Link Connection.  Discontinuities
    // in the counters contained in this table are indicated by the value in
    // frCircuitCreationTime. The type is slice of
    // FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry.
    Frcircuitentry []FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry
}

func (frcircuittable *FRAMERELAYDTEMIB_Frcircuittable) GetEntityData() *types.CommonEntityData {
    frcircuittable.EntityData.YFilter = frcircuittable.YFilter
    frcircuittable.EntityData.YangName = "frCircuitTable"
    frcircuittable.EntityData.BundleName = "cisco_ios_xe"
    frcircuittable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frcircuittable.EntityData.SegmentPath = "frCircuitTable"
    frcircuittable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frcircuittable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frcircuittable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frcircuittable.EntityData.Children = make(map[string]types.YChild)
    frcircuittable.EntityData.Children["frCircuitEntry"] = types.YChild{"Frcircuitentry", nil}
    for i := range frcircuittable.Frcircuitentry {
        frcircuittable.EntityData.Children[types.GetSegmentPath(&frcircuittable.Frcircuitentry[i])] = types.YChild{"Frcircuitentry", &frcircuittable.Frcircuitentry[i]}
    }
    frcircuittable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frcircuittable.EntityData)
}

// FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry
// The information regarding a single Data Link
// Connection.  Discontinuities in the counters contained
// in this table are indicated by the value in
// frCircuitCreationTime.
type FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry struct {
    EntityData types.CommonEntityData
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

func (frcircuitentry *FRAMERELAYDTEMIB_Frcircuittable_Frcircuitentry) GetEntityData() *types.CommonEntityData {
    frcircuitentry.EntityData.YFilter = frcircuitentry.YFilter
    frcircuitentry.EntityData.YangName = "frCircuitEntry"
    frcircuitentry.EntityData.BundleName = "cisco_ios_xe"
    frcircuitentry.EntityData.ParentYangName = "frCircuitTable"
    frcircuitentry.EntityData.SegmentPath = "frCircuitEntry" + "[frCircuitIfIndex='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitifindex) + "']" + "[frCircuitDlci='" + fmt.Sprintf("%v", frcircuitentry.Frcircuitdlci) + "']"
    frcircuitentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frcircuitentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frcircuitentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frcircuitentry.EntityData.Children = make(map[string]types.YChild)
    frcircuitentry.EntityData.Leafs = make(map[string]types.YLeaf)
    frcircuitentry.EntityData.Leafs["frCircuitIfIndex"] = types.YLeaf{"Frcircuitifindex", frcircuitentry.Frcircuitifindex}
    frcircuitentry.EntityData.Leafs["frCircuitDlci"] = types.YLeaf{"Frcircuitdlci", frcircuitentry.Frcircuitdlci}
    frcircuitentry.EntityData.Leafs["frCircuitState"] = types.YLeaf{"Frcircuitstate", frcircuitentry.Frcircuitstate}
    frcircuitentry.EntityData.Leafs["frCircuitReceivedFECNs"] = types.YLeaf{"Frcircuitreceivedfecns", frcircuitentry.Frcircuitreceivedfecns}
    frcircuitentry.EntityData.Leafs["frCircuitReceivedBECNs"] = types.YLeaf{"Frcircuitreceivedbecns", frcircuitentry.Frcircuitreceivedbecns}
    frcircuitentry.EntityData.Leafs["frCircuitSentFrames"] = types.YLeaf{"Frcircuitsentframes", frcircuitentry.Frcircuitsentframes}
    frcircuitentry.EntityData.Leafs["frCircuitSentOctets"] = types.YLeaf{"Frcircuitsentoctets", frcircuitentry.Frcircuitsentoctets}
    frcircuitentry.EntityData.Leafs["frCircuitReceivedFrames"] = types.YLeaf{"Frcircuitreceivedframes", frcircuitentry.Frcircuitreceivedframes}
    frcircuitentry.EntityData.Leafs["frCircuitReceivedOctets"] = types.YLeaf{"Frcircuitreceivedoctets", frcircuitentry.Frcircuitreceivedoctets}
    frcircuitentry.EntityData.Leafs["frCircuitCreationTime"] = types.YLeaf{"Frcircuitcreationtime", frcircuitentry.Frcircuitcreationtime}
    frcircuitentry.EntityData.Leafs["frCircuitLastTimeChange"] = types.YLeaf{"Frcircuitlasttimechange", frcircuitentry.Frcircuitlasttimechange}
    frcircuitentry.EntityData.Leafs["frCircuitCommittedBurst"] = types.YLeaf{"Frcircuitcommittedburst", frcircuitentry.Frcircuitcommittedburst}
    frcircuitentry.EntityData.Leafs["frCircuitExcessBurst"] = types.YLeaf{"Frcircuitexcessburst", frcircuitentry.Frcircuitexcessburst}
    frcircuitentry.EntityData.Leafs["frCircuitThroughput"] = types.YLeaf{"Frcircuitthroughput", frcircuitentry.Frcircuitthroughput}
    frcircuitentry.EntityData.Leafs["frCircuitMulticast"] = types.YLeaf{"Frcircuitmulticast", frcircuitentry.Frcircuitmulticast}
    frcircuitentry.EntityData.Leafs["frCircuitType"] = types.YLeaf{"Frcircuittype", frcircuitentry.Frcircuittype}
    frcircuitentry.EntityData.Leafs["frCircuitDiscards"] = types.YLeaf{"Frcircuitdiscards", frcircuitentry.Frcircuitdiscards}
    frcircuitentry.EntityData.Leafs["frCircuitReceivedDEs"] = types.YLeaf{"Frcircuitreceiveddes", frcircuitentry.Frcircuitreceiveddes}
    frcircuitentry.EntityData.Leafs["frCircuitSentDEs"] = types.YLeaf{"Frcircuitsentdes", frcircuitentry.Frcircuitsentdes}
    frcircuitentry.EntityData.Leafs["frCircuitLogicalIfIndex"] = types.YLeaf{"Frcircuitlogicalifindex", frcircuitentry.Frcircuitlogicalifindex}
    frcircuitentry.EntityData.Leafs["frCircuitRowStatus"] = types.YLeaf{"Frcircuitrowstatus", frcircuitentry.Frcircuitrowstatus}
    return &(frcircuitentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of FRAMERELAYDTEMIB_Frerrtable_Frerrentry.
    Frerrentry []FRAMERELAYDTEMIB_Frerrtable_Frerrentry
}

func (frerrtable *FRAMERELAYDTEMIB_Frerrtable) GetEntityData() *types.CommonEntityData {
    frerrtable.EntityData.YFilter = frerrtable.YFilter
    frerrtable.EntityData.YangName = "frErrTable"
    frerrtable.EntityData.BundleName = "cisco_ios_xe"
    frerrtable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frerrtable.EntityData.SegmentPath = "frErrTable"
    frerrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frerrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frerrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frerrtable.EntityData.Children = make(map[string]types.YChild)
    frerrtable.EntityData.Children["frErrEntry"] = types.YChild{"Frerrentry", nil}
    for i := range frerrtable.Frerrentry {
        frerrtable.EntityData.Children[types.GetSegmentPath(&frerrtable.Frerrentry[i])] = types.YChild{"Frerrentry", &frerrtable.Frerrentry[i]}
    }
    frerrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frerrtable.EntityData)
}

// FRAMERELAYDTEMIB_Frerrtable_Frerrentry
// The error information for a single frame relay
// interface.
type FRAMERELAYDTEMIB_Frerrtable_Frerrentry struct {
    EntityData types.CommonEntityData
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

func (frerrentry *FRAMERELAYDTEMIB_Frerrtable_Frerrentry) GetEntityData() *types.CommonEntityData {
    frerrentry.EntityData.YFilter = frerrentry.YFilter
    frerrentry.EntityData.YangName = "frErrEntry"
    frerrentry.EntityData.BundleName = "cisco_ios_xe"
    frerrentry.EntityData.ParentYangName = "frErrTable"
    frerrentry.EntityData.SegmentPath = "frErrEntry" + "[frErrIfIndex='" + fmt.Sprintf("%v", frerrentry.Frerrifindex) + "']"
    frerrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frerrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frerrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frerrentry.EntityData.Children = make(map[string]types.YChild)
    frerrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    frerrentry.EntityData.Leafs["frErrIfIndex"] = types.YLeaf{"Frerrifindex", frerrentry.Frerrifindex}
    frerrentry.EntityData.Leafs["frErrType"] = types.YLeaf{"Frerrtype", frerrentry.Frerrtype}
    frerrentry.EntityData.Leafs["frErrData"] = types.YLeaf{"Frerrdata", frerrentry.Frerrdata}
    frerrentry.EntityData.Leafs["frErrTime"] = types.YLeaf{"Frerrtime", frerrentry.Frerrtime}
    frerrentry.EntityData.Leafs["frErrFaults"] = types.YLeaf{"Frerrfaults", frerrentry.Frerrfaults}
    frerrentry.EntityData.Leafs["frErrFaultTime"] = types.YLeaf{"Frerrfaulttime", frerrentry.Frerrfaulttime}
    return &(frerrentry.EntityData)
}

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

