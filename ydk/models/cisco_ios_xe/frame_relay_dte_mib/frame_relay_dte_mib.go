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

    
    FrameRelayTrapControl FRAMERELAYDTEMIB_FrameRelayTrapControl

    // The Parameters for the Data Link Connection Management Interface for the
    // frame relay service on this interface.
    FrDlcmiTable FRAMERELAYDTEMIB_FrDlcmiTable

    // A table containing information about specific Data Link Connections (DLC)
    // or virtual circuits.
    FrCircuitTable FRAMERELAYDTEMIB_FrCircuitTable

    // A table containing information about Errors on the Frame Relay interface. 
    // Discontinuities in the counters contained in this table are the same as
    // apply to the ifEntry associated with the Interface.
    FrErrTable FRAMERELAYDTEMIB_FrErrTable
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

    fRAMERELAYDTEMIB.EntityData.Children = types.NewOrderedMap()
    fRAMERELAYDTEMIB.EntityData.Children.Append("frameRelayTrapControl", types.YChild{"FrameRelayTrapControl", &fRAMERELAYDTEMIB.FrameRelayTrapControl})
    fRAMERELAYDTEMIB.EntityData.Children.Append("frDlcmiTable", types.YChild{"FrDlcmiTable", &fRAMERELAYDTEMIB.FrDlcmiTable})
    fRAMERELAYDTEMIB.EntityData.Children.Append("frCircuitTable", types.YChild{"FrCircuitTable", &fRAMERELAYDTEMIB.FrCircuitTable})
    fRAMERELAYDTEMIB.EntityData.Children.Append("frErrTable", types.YChild{"FrErrTable", &fRAMERELAYDTEMIB.FrErrTable})
    fRAMERELAYDTEMIB.EntityData.Leafs = types.NewOrderedMap()

    fRAMERELAYDTEMIB.EntityData.YListKeys = []string {}

    return &(fRAMERELAYDTEMIB.EntityData)
}

// FRAMERELAYDTEMIB_FrameRelayTrapControl
type FRAMERELAYDTEMIB_FrameRelayTrapControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the frDLCIStatusChange
    // trap. The type is FrTrapState.
    FrTrapState interface{}

    // This variable indicates the number of milliseconds that must elapse between
    // trap emissions.  If events occur more rapidly, the impementation may simply
    // fail to trap, or may queue traps until an appropriate time. The type is
    // interface{} with range: 0..3600000.
    FrTrapMaxRate interface{}
}

func (frameRelayTrapControl *FRAMERELAYDTEMIB_FrameRelayTrapControl) GetEntityData() *types.CommonEntityData {
    frameRelayTrapControl.EntityData.YFilter = frameRelayTrapControl.YFilter
    frameRelayTrapControl.EntityData.YangName = "frameRelayTrapControl"
    frameRelayTrapControl.EntityData.BundleName = "cisco_ios_xe"
    frameRelayTrapControl.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frameRelayTrapControl.EntityData.SegmentPath = "frameRelayTrapControl"
    frameRelayTrapControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frameRelayTrapControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frameRelayTrapControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frameRelayTrapControl.EntityData.Children = types.NewOrderedMap()
    frameRelayTrapControl.EntityData.Leafs = types.NewOrderedMap()
    frameRelayTrapControl.EntityData.Leafs.Append("frTrapState", types.YLeaf{"FrTrapState", frameRelayTrapControl.FrTrapState})
    frameRelayTrapControl.EntityData.Leafs.Append("frTrapMaxRate", types.YLeaf{"FrTrapMaxRate", frameRelayTrapControl.FrTrapMaxRate})

    frameRelayTrapControl.EntityData.YListKeys = []string {}

    return &(frameRelayTrapControl.EntityData)
}

// FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState represents the frDLCIStatusChange trap.
type FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState string

const (
    FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState_enabled FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState = "enabled"

    FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState_disabled FRAMERELAYDTEMIB_FrameRelayTrapControl_FrTrapState = "disabled"
)

// FRAMERELAYDTEMIB_FrDlcmiTable
// The Parameters for the Data Link Connection Management
// Interface for the frame relay service on this
// interface.
type FRAMERELAYDTEMIB_FrDlcmiTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Connection Management Interface.
    // The type is slice of FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry.
    FrDlcmiEntry []*FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry
}

func (frDlcmiTable *FRAMERELAYDTEMIB_FrDlcmiTable) GetEntityData() *types.CommonEntityData {
    frDlcmiTable.EntityData.YFilter = frDlcmiTable.YFilter
    frDlcmiTable.EntityData.YangName = "frDlcmiTable"
    frDlcmiTable.EntityData.BundleName = "cisco_ios_xe"
    frDlcmiTable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frDlcmiTable.EntityData.SegmentPath = "frDlcmiTable"
    frDlcmiTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frDlcmiTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frDlcmiTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frDlcmiTable.EntityData.Children = types.NewOrderedMap()
    frDlcmiTable.EntityData.Children.Append("frDlcmiEntry", types.YChild{"FrDlcmiEntry", nil})
    for i := range frDlcmiTable.FrDlcmiEntry {
        frDlcmiTable.EntityData.Children.Append(types.GetSegmentPath(frDlcmiTable.FrDlcmiEntry[i]), types.YChild{"FrDlcmiEntry", frDlcmiTable.FrDlcmiEntry[i]})
    }
    frDlcmiTable.EntityData.Leafs = types.NewOrderedMap()

    frDlcmiTable.EntityData.YListKeys = []string {}

    return &(frDlcmiTable.EntityData)
}

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry
// The Parameters for a particular Data Link Connection
// Management Interface.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the corresponding ifEntry.
    // The type is interface{} with range: 1..2147483647.
    FrDlcmiIfIndex interface{}

    // This variable states which Data Link Connection Management scheme is active
    // (and by implication, what DLCI it uses) on the Frame Relay interface. The
    // type is FrDlcmiState.
    FrDlcmiState interface{}

    // This variable states which address format is in use on the Frame Relay
    // interface. The type is FrDlcmiAddress.
    FrDlcmiAddress interface{}

    // This variable states the address length in octets.  In the case of Q922
    // format, the length indicates the entire length of the address including the
    // control portion. The type is FrDlcmiAddressLen.
    FrDlcmiAddressLen interface{}

    // This is the number of seconds between successive status enquiry messages.
    // The type is interface{} with range: 5..30. Units are seconds.
    FrDlcmiPollingInterval interface{}

    // Number of status enquiry intervals that pass before issuance of a full
    // status enquiry message. The type is interface{} with range: 1..255.
    FrDlcmiFullEnquiryInterval interface{}

    // This is the maximum number of unanswered Status Enquiries the equipment
    // shall accept before declaring the interface down. The type is interface{}
    // with range: 1..10.
    FrDlcmiErrorThreshold interface{}

    // This is the number of status polling intervals over which the error
    // threshold is counted.  For example, if within 'MonitoredEvents' number of
    // events the station receives 'ErrorThreshold' number of errors, the
    // interface is marked as down. The type is interface{} with range: 1..10.
    FrDlcmiMonitoredEvents interface{}

    // The maximum number of Virtual Circuits allowed for this interface.  Usually
    // dictated by the Frame Relay network.  In response to a SET, if a value less
    // than zero or higher than the agent's maximal capability is configured, the
    // agent should respond badValue. The type is interface{} with range:
    // 0..8388607.
    FrDlcmiMaxSupportedVCs interface{}

    // This indicates whether the Frame Relay interface is using a multicast
    // service. The type is FrDlcmiMulticast.
    FrDlcmiMulticast interface{}

    // This indicates the status of the Frame Relay interface as determined by the
    // performance of the dlcmi.  If no dlcmi is running, the Frame Relay
    // interface will stay in the running state indefinitely. The type is
    // FrDlcmiStatus.
    FrDlcmiStatus interface{}

    // SNMP Version 2 Row Status Variable.  Writable objects in the table may be
    // written in any RowStatus state. The type is RowStatus.
    FrDlcmiRowStatus interface{}
}

func (frDlcmiEntry *FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry) GetEntityData() *types.CommonEntityData {
    frDlcmiEntry.EntityData.YFilter = frDlcmiEntry.YFilter
    frDlcmiEntry.EntityData.YangName = "frDlcmiEntry"
    frDlcmiEntry.EntityData.BundleName = "cisco_ios_xe"
    frDlcmiEntry.EntityData.ParentYangName = "frDlcmiTable"
    frDlcmiEntry.EntityData.SegmentPath = "frDlcmiEntry" + types.AddKeyToken(frDlcmiEntry.FrDlcmiIfIndex, "frDlcmiIfIndex")
    frDlcmiEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frDlcmiEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frDlcmiEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frDlcmiEntry.EntityData.Children = types.NewOrderedMap()
    frDlcmiEntry.EntityData.Leafs = types.NewOrderedMap()
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiIfIndex", types.YLeaf{"FrDlcmiIfIndex", frDlcmiEntry.FrDlcmiIfIndex})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiState", types.YLeaf{"FrDlcmiState", frDlcmiEntry.FrDlcmiState})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiAddress", types.YLeaf{"FrDlcmiAddress", frDlcmiEntry.FrDlcmiAddress})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiAddressLen", types.YLeaf{"FrDlcmiAddressLen", frDlcmiEntry.FrDlcmiAddressLen})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiPollingInterval", types.YLeaf{"FrDlcmiPollingInterval", frDlcmiEntry.FrDlcmiPollingInterval})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiFullEnquiryInterval", types.YLeaf{"FrDlcmiFullEnquiryInterval", frDlcmiEntry.FrDlcmiFullEnquiryInterval})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiErrorThreshold", types.YLeaf{"FrDlcmiErrorThreshold", frDlcmiEntry.FrDlcmiErrorThreshold})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiMonitoredEvents", types.YLeaf{"FrDlcmiMonitoredEvents", frDlcmiEntry.FrDlcmiMonitoredEvents})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiMaxSupportedVCs", types.YLeaf{"FrDlcmiMaxSupportedVCs", frDlcmiEntry.FrDlcmiMaxSupportedVCs})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiMulticast", types.YLeaf{"FrDlcmiMulticast", frDlcmiEntry.FrDlcmiMulticast})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiStatus", types.YLeaf{"FrDlcmiStatus", frDlcmiEntry.FrDlcmiStatus})
    frDlcmiEntry.EntityData.Leafs.Append("frDlcmiRowStatus", types.YLeaf{"FrDlcmiRowStatus", frDlcmiEntry.FrDlcmiRowStatus})

    frDlcmiEntry.EntityData.YListKeys = []string {"FrDlcmiIfIndex"}

    return &(frDlcmiEntry.EntityData)
}

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress represents the Frame Relay interface.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress string

const (
    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q921 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q921"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922March90 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922March90"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922November90 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922November90"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922"
)

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen represents portion.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen string

const (
    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_twoOctets FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "twoOctets"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_threeOctets FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "threeOctets"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_fourOctets FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "fourOctets"
)

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast represents using a multicast service.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast string

const (
    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast_nonBroadcast FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast = "nonBroadcast"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast_broadcast FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast = "broadcast"
)

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState represents DLCI it uses) on the Frame Relay interface.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState string

const (
    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_noLmiConfigured FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "noLmiConfigured"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_lmiRev1 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "lmiRev1"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_ansiT1617D FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "ansiT1617D"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_ansiT1617B FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "ansiT1617B"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_itut933A FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "itut933A"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_ansiT1617D1994 FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "ansiT1617D1994"
)

// FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus represents in the running state indefinitely.
type FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus string

const (
    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus_running FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus = "running"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus_fault FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus = "fault"

    FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus_initializing FRAMERELAYDTEMIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiStatus = "initializing"
)

// FRAMERELAYDTEMIB_FrCircuitTable
// A table containing information about specific Data
// Link Connections (DLC) or virtual circuits.
type FRAMERELAYDTEMIB_FrCircuitTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single Data Link Connection.  Discontinuities
    // in the counters contained in this table are indicated by the value in
    // frCircuitCreationTime. The type is slice of
    // FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry.
    FrCircuitEntry []*FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry
}

func (frCircuitTable *FRAMERELAYDTEMIB_FrCircuitTable) GetEntityData() *types.CommonEntityData {
    frCircuitTable.EntityData.YFilter = frCircuitTable.YFilter
    frCircuitTable.EntityData.YangName = "frCircuitTable"
    frCircuitTable.EntityData.BundleName = "cisco_ios_xe"
    frCircuitTable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frCircuitTable.EntityData.SegmentPath = "frCircuitTable"
    frCircuitTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frCircuitTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frCircuitTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frCircuitTable.EntityData.Children = types.NewOrderedMap()
    frCircuitTable.EntityData.Children.Append("frCircuitEntry", types.YChild{"FrCircuitEntry", nil})
    for i := range frCircuitTable.FrCircuitEntry {
        frCircuitTable.EntityData.Children.Append(types.GetSegmentPath(frCircuitTable.FrCircuitEntry[i]), types.YChild{"FrCircuitEntry", frCircuitTable.FrCircuitEntry[i]})
    }
    frCircuitTable.EntityData.Leafs = types.NewOrderedMap()

    frCircuitTable.EntityData.YListKeys = []string {}

    return &(frCircuitTable.EntityData)
}

// FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry
// The information regarding a single Data Link
// Connection.  Discontinuities in the counters contained
// in this table are indicated by the value in
// frCircuitCreationTime.
type FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the ifEntry this virtual
    // circuit is layered onto. The type is interface{} with range: 1..2147483647.
    FrCircuitIfIndex interface{}

    // This attribute is a key. The Data Link Connection Identifier for this
    // virtual circuit. The type is interface{} with range: 0..8388607.
    FrCircuitDlci interface{}

    // Indicates whether the particular virtual circuit is operational.  In the
    // absence of a Data Link Connection Management Interface, virtual circuit
    // entries (rows) may be created by setting virtual circuit state to 'active',
    // or deleted by changing Circuit state to 'invalid'.  Whether or not the row
    // actually disappears is left to the implementation, so this object may
    // actually read as 'invalid' for some arbitrary length of time.  It is also
    // legal to set the state of a virtual circuit to 'inactive' to temporarily
    // disable a given circuit.  The use of 'invalid' is deprecated in this SNMP
    // Version 2 MIB, in favor of frCircuitRowStatus. The type is FrCircuitState.
    FrCircuitState interface{}

    // Number of frames received from the network indicating forward congestion
    // since the virtual circuit was created.  This occurs when the remote DTE
    // sets the FECN flag, or when a switch in the network enqueues the frame to a
    // trunk whose transmission queue is congested. The type is interface{} with
    // range: 0..4294967295.
    FrCircuitReceivedFECNs interface{}

    // Number of frames received from the network indicating backward congestion
    // since the virtual circuit was created.  This occurs when the remote DTE
    // sets the BECN flag, or when a switch in the network receives the frame from
    // a trunk whose transmission queue is congested. The type is interface{} with
    // range: 0..4294967295.
    FrCircuitReceivedBECNs interface{}

    // The number of frames sent from this virtual circuit since it was created.
    // The type is interface{} with range: 0..4294967295.
    FrCircuitSentFrames interface{}

    // The number of octets sent from this virtual circuit since it was created. 
    // Octets counted are the full frame relay header and the payload, but do not
    // include the flag characters or CRC. The type is interface{} with range:
    // 0..4294967295.
    FrCircuitSentOctets interface{}

    // Number of frames received over this virtual circuit since it was created.
    // The type is interface{} with range: 0..4294967295.
    FrCircuitReceivedFrames interface{}

    // Number of octets received over this virtual circuit since it was created. 
    // Octets counted include the full frame relay header, but do not include the
    // flag characters or the CRC. The type is interface{} with range:
    // 0..4294967295.
    FrCircuitReceivedOctets interface{}

    // The value of sysUpTime when the virtual circuit was created, whether by the
    // Data Link Connection Management Interface or by a SetRequest. The type is
    // interface{} with range: 0..4294967295.
    FrCircuitCreationTime interface{}

    // The value of sysUpTime when last there was a change in the virtual circuit
    // state. The type is interface{} with range: 0..4294967295.
    FrCircuitLastTimeChange interface{}

    // This variable indicates the maximum amount of data, in bits, that the
    // network agrees to transfer under normal conditions, during the measurement
    // interval. The type is interface{} with range: 0..2147483647.
    FrCircuitCommittedBurst interface{}

    // This variable indicates the maximum amount of uncommitted data bits that
    // the network will attempt to deliver over the measurement interval.  By
    // default, if not configured when creating the entry, the Excess Information
    // Burst Size is set to the value of ifSpeed. The type is interface{} with
    // range: 0..2147483647.
    FrCircuitExcessBurst interface{}

    // Throughput is the average number of 'Frame Relay Information Field' bits
    // transferred per second across a user network interface in one direction,
    // measured over the measurement interval.  If the configured committed burst
    // rate and throughput are both non-zero, the measurement interval, T, is    
    // T=frCircuitCommittedBurst/frCircuitThroughput.  If the configured committed
    // burst rate and throughput are both zero, the measurement interval, T, is   
    // T=frCircuitExcessBurst/ifSpeed. The type is interface{} with range:
    // 0..2147483647.
    FrCircuitThroughput interface{}

    // This indicates whether this VC is used as a unicast VC (i.e. not multicast)
    // or the type of multicast service subscribed to. The type is
    // FrCircuitMulticast.
    FrCircuitMulticast interface{}

    // Indication of whether the VC was manually created (static), or dynamically
    // created (dynamic) via the data link control management interface. The type
    // is FrCircuitType.
    FrCircuitType interface{}

    // The number of inbound frames dropped because of format errors, or because
    // the VC is inactive. The type is interface{} with range: 0..4294967295.
    FrCircuitDiscards interface{}

    // Number of frames received from the network indicating that they were
    // eligible for discard since the virtual circuit was created.  This occurs
    // when the remote DTE sets the DE flag, or when in remote DTE's switch
    // detects that the frame was received as Excess Burst data. The type is
    // interface{} with range: 0..4294967295.
    FrCircuitReceivedDEs interface{}

    // Number of frames sent to the network indicating that they were eligible for
    // discard since the virtual circuit was created.   This occurs when the local
    // DTE sets the DE flag, indicating that during Network congestion situations
    // those frames should be discarded in preference of other frames sent without
    // the DE bit set. The type is interface{} with range: 0..4294967295.
    FrCircuitSentDEs interface{}

    // Normally the same value as frDlcmiIfIndex, but different when an
    // implementation associates a virtual ifEntry with a DLC or set of DLCs in
    // order to associate higher layer objects such as the ipAddrEntry with a
    // subset of the virtual circuits on a Frame Relay interface. The type of such
    // ifEntries is defined by the higher layer object; for example, if PPP/Frame
    // Relay is implemented, the ifType of this ifEntry would be PPP. If it is not
    // so defined, as would be the case with an ipAddrEntry, it should be of type
    // Other. The type is interface{} with range: 1..2147483647.
    FrCircuitLogicalIfIndex interface{}

    // This object is used to create a new row or modify or destroy an existing
    // row in the manner described in the definition of the RowStatus textual
    // convention. Writable objects in the table may be written in any RowStatus
    // state. The type is RowStatus.
    FrCircuitRowStatus interface{}
}

func (frCircuitEntry *FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry) GetEntityData() *types.CommonEntityData {
    frCircuitEntry.EntityData.YFilter = frCircuitEntry.YFilter
    frCircuitEntry.EntityData.YangName = "frCircuitEntry"
    frCircuitEntry.EntityData.BundleName = "cisco_ios_xe"
    frCircuitEntry.EntityData.ParentYangName = "frCircuitTable"
    frCircuitEntry.EntityData.SegmentPath = "frCircuitEntry" + types.AddKeyToken(frCircuitEntry.FrCircuitIfIndex, "frCircuitIfIndex") + types.AddKeyToken(frCircuitEntry.FrCircuitDlci, "frCircuitDlci")
    frCircuitEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frCircuitEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frCircuitEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frCircuitEntry.EntityData.Children = types.NewOrderedMap()
    frCircuitEntry.EntityData.Leafs = types.NewOrderedMap()
    frCircuitEntry.EntityData.Leafs.Append("frCircuitIfIndex", types.YLeaf{"FrCircuitIfIndex", frCircuitEntry.FrCircuitIfIndex})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitDlci", types.YLeaf{"FrCircuitDlci", frCircuitEntry.FrCircuitDlci})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitState", types.YLeaf{"FrCircuitState", frCircuitEntry.FrCircuitState})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitReceivedFECNs", types.YLeaf{"FrCircuitReceivedFECNs", frCircuitEntry.FrCircuitReceivedFECNs})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitReceivedBECNs", types.YLeaf{"FrCircuitReceivedBECNs", frCircuitEntry.FrCircuitReceivedBECNs})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitSentFrames", types.YLeaf{"FrCircuitSentFrames", frCircuitEntry.FrCircuitSentFrames})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitSentOctets", types.YLeaf{"FrCircuitSentOctets", frCircuitEntry.FrCircuitSentOctets})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitReceivedFrames", types.YLeaf{"FrCircuitReceivedFrames", frCircuitEntry.FrCircuitReceivedFrames})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitReceivedOctets", types.YLeaf{"FrCircuitReceivedOctets", frCircuitEntry.FrCircuitReceivedOctets})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitCreationTime", types.YLeaf{"FrCircuitCreationTime", frCircuitEntry.FrCircuitCreationTime})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitLastTimeChange", types.YLeaf{"FrCircuitLastTimeChange", frCircuitEntry.FrCircuitLastTimeChange})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitCommittedBurst", types.YLeaf{"FrCircuitCommittedBurst", frCircuitEntry.FrCircuitCommittedBurst})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitExcessBurst", types.YLeaf{"FrCircuitExcessBurst", frCircuitEntry.FrCircuitExcessBurst})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitThroughput", types.YLeaf{"FrCircuitThroughput", frCircuitEntry.FrCircuitThroughput})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitMulticast", types.YLeaf{"FrCircuitMulticast", frCircuitEntry.FrCircuitMulticast})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitType", types.YLeaf{"FrCircuitType", frCircuitEntry.FrCircuitType})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitDiscards", types.YLeaf{"FrCircuitDiscards", frCircuitEntry.FrCircuitDiscards})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitReceivedDEs", types.YLeaf{"FrCircuitReceivedDEs", frCircuitEntry.FrCircuitReceivedDEs})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitSentDEs", types.YLeaf{"FrCircuitSentDEs", frCircuitEntry.FrCircuitSentDEs})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitLogicalIfIndex", types.YLeaf{"FrCircuitLogicalIfIndex", frCircuitEntry.FrCircuitLogicalIfIndex})
    frCircuitEntry.EntityData.Leafs.Append("frCircuitRowStatus", types.YLeaf{"FrCircuitRowStatus", frCircuitEntry.FrCircuitRowStatus})

    frCircuitEntry.EntityData.YListKeys = []string {"FrCircuitIfIndex", "FrCircuitDlci"}

    return &(frCircuitEntry.EntityData)
}

// FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast represents subscribed to
type FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast string

const (
    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast_unicast FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast = "unicast"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast_oneWay FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast = "oneWay"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast_twoWay FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast = "twoWay"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast_nWay FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitMulticast = "nWay"
)

// FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState represents 2 MIB, in favor of frCircuitRowStatus.
type FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState string

const (
    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_invalid FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "invalid"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_active FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "active"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_inactive FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "inactive"
)

// FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType represents link control management interface.
type FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType string

const (
    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType_static FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType = "static"

    FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType_dynamic FRAMERELAYDTEMIB_FrCircuitTable_FrCircuitEntry_FrCircuitType = "dynamic"
)

// FRAMERELAYDTEMIB_FrErrTable
// A table containing information about Errors on the
// Frame Relay interface.  Discontinuities in the counters
// contained in this table are the same as apply to the
// ifEntry associated with the Interface.
type FRAMERELAYDTEMIB_FrErrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of FRAMERELAYDTEMIB_FrErrTable_FrErrEntry.
    FrErrEntry []*FRAMERELAYDTEMIB_FrErrTable_FrErrEntry
}

func (frErrTable *FRAMERELAYDTEMIB_FrErrTable) GetEntityData() *types.CommonEntityData {
    frErrTable.EntityData.YFilter = frErrTable.YFilter
    frErrTable.EntityData.YangName = "frErrTable"
    frErrTable.EntityData.BundleName = "cisco_ios_xe"
    frErrTable.EntityData.ParentYangName = "FRAME-RELAY-DTE-MIB"
    frErrTable.EntityData.SegmentPath = "frErrTable"
    frErrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frErrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frErrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frErrTable.EntityData.Children = types.NewOrderedMap()
    frErrTable.EntityData.Children.Append("frErrEntry", types.YChild{"FrErrEntry", nil})
    for i := range frErrTable.FrErrEntry {
        frErrTable.EntityData.Children.Append(types.GetSegmentPath(frErrTable.FrErrEntry[i]), types.YChild{"FrErrEntry", frErrTable.FrErrEntry[i]})
    }
    frErrTable.EntityData.Leafs = types.NewOrderedMap()

    frErrTable.EntityData.YListKeys = []string {}

    return &(frErrTable.EntityData)
}

// FRAMERELAYDTEMIB_FrErrTable_FrErrEntry
// The error information for a single frame relay
// interface.
type FRAMERELAYDTEMIB_FrErrTable_FrErrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex Value of the corresponding ifEntry.
    // The type is interface{} with range: 1..2147483647.
    FrErrIfIndex interface{}

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
    // The type is FrErrType.
    FrErrType interface{}

    // An octet string containing as much of the error packet as possible.  As a
    // minimum, it must contain the Q.922 Address or as much as was delivered.  It
    // is desirable to include all header and demultiplexing information. The type
    // is string with length: 1..1600.
    FrErrData interface{}

    // The value of sysUpTime at which the error was detected. The type is
    // interface{} with range: 0..4294967295.
    FrErrTime interface{}

    // The number of times the interface has gone down since it was initialized.
    // The type is interface{} with range: 0..4294967295.
    FrErrFaults interface{}

    // The value of sysUpTime at the time when the interface was taken down due to
    // excessive errors.  Excessive errors is defined as the time when a DLCMI
    // exceeds the frDlcmiErrorThreshold number of errors within
    // frDlcmiMonitoredEvents. See FrDlcmiEntry for further details. The type is
    // interface{} with range: 0..4294967295.
    FrErrFaultTime interface{}
}

func (frErrEntry *FRAMERELAYDTEMIB_FrErrTable_FrErrEntry) GetEntityData() *types.CommonEntityData {
    frErrEntry.EntityData.YFilter = frErrEntry.YFilter
    frErrEntry.EntityData.YangName = "frErrEntry"
    frErrEntry.EntityData.BundleName = "cisco_ios_xe"
    frErrEntry.EntityData.ParentYangName = "frErrTable"
    frErrEntry.EntityData.SegmentPath = "frErrEntry" + types.AddKeyToken(frErrEntry.FrErrIfIndex, "frErrIfIndex")
    frErrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frErrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frErrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frErrEntry.EntityData.Children = types.NewOrderedMap()
    frErrEntry.EntityData.Leafs = types.NewOrderedMap()
    frErrEntry.EntityData.Leafs.Append("frErrIfIndex", types.YLeaf{"FrErrIfIndex", frErrEntry.FrErrIfIndex})
    frErrEntry.EntityData.Leafs.Append("frErrType", types.YLeaf{"FrErrType", frErrEntry.FrErrType})
    frErrEntry.EntityData.Leafs.Append("frErrData", types.YLeaf{"FrErrData", frErrEntry.FrErrData})
    frErrEntry.EntityData.Leafs.Append("frErrTime", types.YLeaf{"FrErrTime", frErrEntry.FrErrTime})
    frErrEntry.EntityData.Leafs.Append("frErrFaults", types.YLeaf{"FrErrFaults", frErrEntry.FrErrFaults})
    frErrEntry.EntityData.Leafs.Append("frErrFaultTime", types.YLeaf{"FrErrFaultTime", frErrEntry.FrErrFaultTime})

    frErrEntry.EntityData.YListKeys = []string {"FrErrIfIndex"}

    return &(frErrEntry.EntityData)
}

// FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType represents                    cold start or warm start.
type FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType string

const (
    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_unknownError FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "unknownError"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_receiveShort FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "receiveShort"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_receiveLong FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "receiveLong"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_illegalAddress FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "illegalAddress"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_unknownAddress FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "unknownAddress"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_dlcmiProtoErr FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiProtoErr"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_dlcmiUnknownIE FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiUnknownIE"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_dlcmiSequenceErr FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiSequenceErr"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_dlcmiUnknownRpt FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiUnknownRpt"

    FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType_noErrorSinceReset FRAMERELAYDTEMIB_FrErrTable_FrErrEntry_FrErrType = "noErrorSinceReset"
)

