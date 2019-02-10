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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    FrameRelayGlobals RFC1315MIB_FrameRelayGlobals

    // The Parameters for the Data Link Connection Management Interface for the
    // frame relay service on this interface.
    FrDlcmiTable RFC1315MIB_FrDlcmiTable

    // A table containing information about specific Data Link Connection
    // Identifiers and corresponding virtual circuits.
    FrCircuitTable RFC1315MIB_FrCircuitTable

    // A table containing information about Errors on the Frame Relay interface.
    FrErrTable RFC1315MIB_FrErrTable
}

func (rFC1315MIB *RFC1315MIB) GetEntityData() *types.CommonEntityData {
    rFC1315MIB.EntityData.YFilter = rFC1315MIB.YFilter
    rFC1315MIB.EntityData.YangName = "RFC1315-MIB"
    rFC1315MIB.EntityData.BundleName = "cisco_ios_xe"
    rFC1315MIB.EntityData.ParentYangName = "RFC1315-MIB"
    rFC1315MIB.EntityData.SegmentPath = "RFC1315-MIB:RFC1315-MIB"
    rFC1315MIB.EntityData.AbsolutePath = rFC1315MIB.EntityData.SegmentPath
    rFC1315MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rFC1315MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rFC1315MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rFC1315MIB.EntityData.Children = types.NewOrderedMap()
    rFC1315MIB.EntityData.Children.Append("frame-relay-globals", types.YChild{"FrameRelayGlobals", &rFC1315MIB.FrameRelayGlobals})
    rFC1315MIB.EntityData.Children.Append("frDlcmiTable", types.YChild{"FrDlcmiTable", &rFC1315MIB.FrDlcmiTable})
    rFC1315MIB.EntityData.Children.Append("frCircuitTable", types.YChild{"FrCircuitTable", &rFC1315MIB.FrCircuitTable})
    rFC1315MIB.EntityData.Children.Append("frErrTable", types.YChild{"FrErrTable", &rFC1315MIB.FrErrTable})
    rFC1315MIB.EntityData.Leafs = types.NewOrderedMap()

    rFC1315MIB.EntityData.YListKeys = []string {}

    return &(rFC1315MIB.EntityData)
}

// RFC1315MIB_FrameRelayGlobals
type RFC1315MIB_FrameRelayGlobals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable  indicates  whether  the  system produces the
    // frDLCIStatusChange trap. The type is FrTrapState.
    FrTrapState interface{}
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetEntityData() *types.CommonEntityData {
    frameRelayGlobals.EntityData.YFilter = frameRelayGlobals.YFilter
    frameRelayGlobals.EntityData.YangName = "frame-relay-globals"
    frameRelayGlobals.EntityData.BundleName = "cisco_ios_xe"
    frameRelayGlobals.EntityData.ParentYangName = "RFC1315-MIB"
    frameRelayGlobals.EntityData.SegmentPath = "frame-relay-globals"
    frameRelayGlobals.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/" + frameRelayGlobals.EntityData.SegmentPath
    frameRelayGlobals.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frameRelayGlobals.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frameRelayGlobals.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frameRelayGlobals.EntityData.Children = types.NewOrderedMap()
    frameRelayGlobals.EntityData.Leafs = types.NewOrderedMap()
    frameRelayGlobals.EntityData.Leafs.Append("frTrapState", types.YLeaf{"FrTrapState", frameRelayGlobals.FrTrapState})

    frameRelayGlobals.EntityData.YListKeys = []string {}

    return &(frameRelayGlobals.EntityData)
}

// RFC1315MIB_FrameRelayGlobals_FrTrapState represents produces the frDLCIStatusChange trap.
type RFC1315MIB_FrameRelayGlobals_FrTrapState string

const (
    RFC1315MIB_FrameRelayGlobals_FrTrapState_enabled RFC1315MIB_FrameRelayGlobals_FrTrapState = "enabled"

    RFC1315MIB_FrameRelayGlobals_FrTrapState_disabled RFC1315MIB_FrameRelayGlobals_FrTrapState = "disabled"
)

// RFC1315MIB_FrDlcmiTable
// The Parameters for the Data Link Connection Management
// Interface for the frame relay service on this
// interface.
type RFC1315MIB_FrDlcmiTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Con- nection Management
    // Interface. The type is slice of RFC1315MIB_FrDlcmiTable_FrDlcmiEntry.
    FrDlcmiEntry []*RFC1315MIB_FrDlcmiTable_FrDlcmiEntry
}

func (frDlcmiTable *RFC1315MIB_FrDlcmiTable) GetEntityData() *types.CommonEntityData {
    frDlcmiTable.EntityData.YFilter = frDlcmiTable.YFilter
    frDlcmiTable.EntityData.YangName = "frDlcmiTable"
    frDlcmiTable.EntityData.BundleName = "cisco_ios_xe"
    frDlcmiTable.EntityData.ParentYangName = "RFC1315-MIB"
    frDlcmiTable.EntityData.SegmentPath = "frDlcmiTable"
    frDlcmiTable.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/" + frDlcmiTable.EntityData.SegmentPath
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

// RFC1315MIB_FrDlcmiTable_FrDlcmiEntry
// The Parameters for a particular Data Link Con-
// nection Management Interface.
type RFC1315MIB_FrDlcmiTable_FrDlcmiEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ifIndex value of the  corresponding  ifEn-
    // try. The type is interface{} with range: -2147483648..2147483647.
    FrDlcmiIfIndex interface{}

    // This variable states which Data  Link  Connec- tion Management scheme is
    // active (and by impli- cation, what DLCI it uses) on the  Frame  Relay
    // interface. The type is FrDlcmiState.
    FrDlcmiState interface{}

    // This variable states which address  format  is in use on the Frame Relay
    // interface. The type is FrDlcmiAddress.
    FrDlcmiAddress interface{}

    // This variable states which address  length  in octets.  In the case of Q922
    // format, the length indicates the entire length of the address  in- cluding
    // the control portion. The type is FrDlcmiAddressLen.
    FrDlcmiAddressLen interface{}

    // This is the number of seconds between  succes- sive status enquiry
    // messages. The type is interface{} with range: 5..30.
    FrDlcmiPollingInterval interface{}

    // Number of status enquiry intervals  that  pass before  issuance  of a full
    // status enquiry mes- sage. The type is interface{} with range: 1..255.
    FrDlcmiFullEnquiryInterval interface{}

    // This  is  the  maximum  number  of  unanswered Status Enquiries the
    // equipment shall accept be- fore declaring the interface down. The type is
    // interface{} with range: 1..10.
    FrDlcmiErrorThreshold interface{}

    // This is the number of status polling intervals over which the error
    // threshold is counted.  For example, if within 'MonitoredEvents' number  of
    // events  the  station  receives 'ErrorThreshold' number of errors, the
    // interface  is  marked  as down. The type is interface{} with range: 1..10.
    FrDlcmiMonitoredEvents interface{}

    // The maximum number of Virtual Circuits allowed for  this  interface.  
    // Usually dictated by the Frame Relay network.  In response to a SET, if a
    // value less than zero or  higher  than the agent's maximal capability is
    // configured, the agent  should  respond  bad- Value. The type is interface{}
    // with range: -2147483648..2147483647.
    FrDlcmiMaxSupportedVCs interface{}

    // This indicates whether the Frame Relay  inter- face is using a multicast
    // service. The type is FrDlcmiMulticast.
    FrDlcmiMulticast interface{}
}

func (frDlcmiEntry *RFC1315MIB_FrDlcmiTable_FrDlcmiEntry) GetEntityData() *types.CommonEntityData {
    frDlcmiEntry.EntityData.YFilter = frDlcmiEntry.YFilter
    frDlcmiEntry.EntityData.YangName = "frDlcmiEntry"
    frDlcmiEntry.EntityData.BundleName = "cisco_ios_xe"
    frDlcmiEntry.EntityData.ParentYangName = "frDlcmiTable"
    frDlcmiEntry.EntityData.SegmentPath = "frDlcmiEntry" + types.AddKeyToken(frDlcmiEntry.FrDlcmiIfIndex, "frDlcmiIfIndex")
    frDlcmiEntry.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/frDlcmiTable/" + frDlcmiEntry.EntityData.SegmentPath
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

    frDlcmiEntry.EntityData.YListKeys = []string {"FrDlcmiIfIndex"}

    return &(frDlcmiEntry.EntityData)
}

// RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress represents in use on the Frame Relay interface.
type RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress string

const (
    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q921 RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q921"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922March90 RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922March90"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922November90 RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922November90"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress_q922 RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddress = "q922"
)

// RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen represents cluding the control portion.
type RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen string

const (
    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_two_octets RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "two-octets"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_three_octets RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "three-octets"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen_four_octets RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiAddressLen = "four-octets"
)

// RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast represents face is using a multicast service.
type RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast string

const (
    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast_nonBroadcast RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast = "nonBroadcast"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast_broadcast RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiMulticast = "broadcast"
)

// RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState represents interface.
type RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState string

const (
    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_noLmiConfigured RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "noLmiConfigured"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_lmiRev1 RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "lmiRev1"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_ansiT1_617_D RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "ansiT1-617-D"

    RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState_ansiT1_617_B RFC1315MIB_FrDlcmiTable_FrDlcmiEntry_FrDlcmiState = "ansiT1-617-B"
)

// RFC1315MIB_FrCircuitTable
// A table containing information about specific Data
// Link Connection Identifiers and corresponding virtual
// circuits.
type RFC1315MIB_FrCircuitTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single  Data  Link Connection Identifier. The
    // type is slice of RFC1315MIB_FrCircuitTable_FrCircuitEntry.
    FrCircuitEntry []*RFC1315MIB_FrCircuitTable_FrCircuitEntry
}

func (frCircuitTable *RFC1315MIB_FrCircuitTable) GetEntityData() *types.CommonEntityData {
    frCircuitTable.EntityData.YFilter = frCircuitTable.YFilter
    frCircuitTable.EntityData.YangName = "frCircuitTable"
    frCircuitTable.EntityData.BundleName = "cisco_ios_xe"
    frCircuitTable.EntityData.ParentYangName = "RFC1315-MIB"
    frCircuitTable.EntityData.SegmentPath = "frCircuitTable"
    frCircuitTable.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/" + frCircuitTable.EntityData.SegmentPath
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

// RFC1315MIB_FrCircuitTable_FrCircuitEntry
// The information regarding a single  Data  Link
// Connection Identifier.
type RFC1315MIB_FrCircuitTable_FrCircuitEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ifIndex Value of the ifEntry this  virtual
    // circuit is layered onto. The type is interface{} with range:
    // -2147483648..2147483647.
    FrCircuitIfIndex interface{}

    // This attribute is a key. The Data Link Connection Identifier  for  this
    // virtual circuit. The type is interface{} with range:
    // -2147483648..2147483647.
    FrCircuitDlci interface{}

    // Indicates whether the particular virtual  cir- cuit  is operational.  In
    // the absence of a Data Link Connection Management  Interface,  virtual
    // circuit  entries  (rows) may be created by set- ting virtual  circuit 
    // state  to  'active',  or deleted by changing Circuit state to 'invalid'.
    // Whether or not the row actually  disappears  is left  to the
    // implementation, so this object may actually read as 'invalid' for  some 
    // arbitrary length  of  time.   It is also legal to set the state of a
    // virtual  circuit  to  'inactive'  to temporarily disable a given circuit.
    // The type is FrCircuitState.
    FrCircuitState interface{}

    // Number of frames received from the network in- dicating  forward 
    // congestion since the virtual circuit was created. The type is interface{}
    // with range: 0..4294967295.
    FrCircuitReceivedFECNs interface{}

    // Number of frames received from the network in- dicating  backward
    // congestion since the virtual circuit was created. The type is interface{}
    // with range: 0..4294967295.
    FrCircuitReceivedBECNs interface{}

    // The number of frames sent  from  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    FrCircuitSentFrames interface{}

    // The number of octets sent  from  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    FrCircuitSentOctets interface{}

    // Number of frames received  over  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    FrCircuitReceivedFrames interface{}

    // Number of octets received  over  this  virtual circuit since it was
    // created. The type is interface{} with range: 0..4294967295.
    FrCircuitReceivedOctets interface{}

    // The value of sysUpTime when the  virtual  cir- cuit was created, whether by
    // the Data Link Con- nection Management Interface  or  by  a  SetRe- quest.
    // The type is interface{} with range: 0..4294967295.
    FrCircuitCreationTime interface{}

    // The value of sysUpTime when last there  was  a change in the virtual
    // circuit state. The type is interface{} with range: 0..4294967295.
    FrCircuitLastTimeChange interface{}

    // This variable indicates the maximum amount  of data,  in  bits,  that  the 
    // network  agrees to transfer under normal  conditions,  during  the
    // measurement interval. The type is interface{} with range:
    // -2147483648..2147483647.
    FrCircuitCommittedBurst interface{}

    // This variable indicates the maximum amount  of uncommitted data bits that
    // the network will at- tempt to deliver over the measurement interval.  By
    // default, if not configured when creating the entry, the Excess Information
    // Burst Size is set to the value of ifSpeed. The type is interface{} with
    // range: -2147483648..2147483647.
    FrCircuitExcessBurst interface{}

    // Throughput is the average number of 'Frame Re- lay  Information  Field' 
    // bits  transferred per second across a user network interface  in  one
    // direction, measured over the measurement inter- val.  If the  configured 
    // committed  burst  rate  and throughput  are  both non-zero, the measurement
    // interval T=frCircuitCommittedBurst/frCircuitThroughput.  If the  configured
    // committed  burst  rate  and throughput  are  both zero, the measurement in-
    // terval        T=frCircuitExcessBurst/ifSpeed. The type is interface{} with
    // range: -2147483648..2147483647.
    FrCircuitThroughput interface{}
}

func (frCircuitEntry *RFC1315MIB_FrCircuitTable_FrCircuitEntry) GetEntityData() *types.CommonEntityData {
    frCircuitEntry.EntityData.YFilter = frCircuitEntry.YFilter
    frCircuitEntry.EntityData.YangName = "frCircuitEntry"
    frCircuitEntry.EntityData.BundleName = "cisco_ios_xe"
    frCircuitEntry.EntityData.ParentYangName = "frCircuitTable"
    frCircuitEntry.EntityData.SegmentPath = "frCircuitEntry" + types.AddKeyToken(frCircuitEntry.FrCircuitIfIndex, "frCircuitIfIndex") + types.AddKeyToken(frCircuitEntry.FrCircuitDlci, "frCircuitDlci")
    frCircuitEntry.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/frCircuitTable/" + frCircuitEntry.EntityData.SegmentPath
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

    frCircuitEntry.EntityData.YListKeys = []string {"FrCircuitIfIndex", "FrCircuitDlci"}

    return &(frCircuitEntry.EntityData)
}

// RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState represents temporarily disable a given circuit.
type RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState string

const (
    RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_invalid RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "invalid"

    RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_active RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "active"

    RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState_inactive RFC1315MIB_FrCircuitTable_FrCircuitEntry_FrCircuitState = "inactive"
)

// RFC1315MIB_FrErrTable
// A table containing information about Errors on the
// Frame Relay interface.
type RFC1315MIB_FrErrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of RFC1315MIB_FrErrTable_FrErrEntry.
    FrErrEntry []*RFC1315MIB_FrErrTable_FrErrEntry
}

func (frErrTable *RFC1315MIB_FrErrTable) GetEntityData() *types.CommonEntityData {
    frErrTable.EntityData.YFilter = frErrTable.YFilter
    frErrTable.EntityData.YangName = "frErrTable"
    frErrTable.EntityData.BundleName = "cisco_ios_xe"
    frErrTable.EntityData.ParentYangName = "RFC1315-MIB"
    frErrTable.EntityData.SegmentPath = "frErrTable"
    frErrTable.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/" + frErrTable.EntityData.SegmentPath
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

// RFC1315MIB_FrErrTable_FrErrEntry
// The error information for a single frame relay
// interface.
type RFC1315MIB_FrErrTable_FrErrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ifIndex Value of the  corresponding  ifEn-
    // try. The type is interface{} with range: -2147483648..2147483647.
    FrErrIfIndex interface{}

    // The type of error that was last seen  on  this interface. The type is
    // FrErrType.
    FrErrType interface{}

    // An octet string containing as much of the  er- ror  packet as possible.  As
    // a minimum, it must contain the Q.922 Address or  as  much  as  was
    // delivered.   It is desirable to include all in- formation up to the PDU.
    // The type is string.
    FrErrData interface{}

    // The value of sysUpTime at which the error  was detected. The type is
    // interface{} with range: 0..4294967295.
    FrErrTime interface{}
}

func (frErrEntry *RFC1315MIB_FrErrTable_FrErrEntry) GetEntityData() *types.CommonEntityData {
    frErrEntry.EntityData.YFilter = frErrEntry.YFilter
    frErrEntry.EntityData.YangName = "frErrEntry"
    frErrEntry.EntityData.BundleName = "cisco_ios_xe"
    frErrEntry.EntityData.ParentYangName = "frErrTable"
    frErrEntry.EntityData.SegmentPath = "frErrEntry" + types.AddKeyToken(frErrEntry.FrErrIfIndex, "frErrIfIndex")
    frErrEntry.EntityData.AbsolutePath = "RFC1315-MIB:RFC1315-MIB/frErrTable/" + frErrEntry.EntityData.SegmentPath
    frErrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frErrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frErrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frErrEntry.EntityData.Children = types.NewOrderedMap()
    frErrEntry.EntityData.Leafs = types.NewOrderedMap()
    frErrEntry.EntityData.Leafs.Append("frErrIfIndex", types.YLeaf{"FrErrIfIndex", frErrEntry.FrErrIfIndex})
    frErrEntry.EntityData.Leafs.Append("frErrType", types.YLeaf{"FrErrType", frErrEntry.FrErrType})
    frErrEntry.EntityData.Leafs.Append("frErrData", types.YLeaf{"FrErrData", frErrEntry.FrErrData})
    frErrEntry.EntityData.Leafs.Append("frErrTime", types.YLeaf{"FrErrTime", frErrEntry.FrErrTime})

    frErrEntry.EntityData.YListKeys = []string {"FrErrIfIndex"}

    return &(frErrEntry.EntityData)
}

// RFC1315MIB_FrErrTable_FrErrEntry_FrErrType represents interface.
type RFC1315MIB_FrErrTable_FrErrEntry_FrErrType string

const (
    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_unknownError RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "unknownError"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_receiveShort RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "receiveShort"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_receiveLong RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "receiveLong"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_illegalDLCI RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "illegalDLCI"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_unknownDLCI RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "unknownDLCI"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_dlcmiProtoErr RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiProtoErr"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_dlcmiUnknownIE RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiUnknownIE"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_dlcmiSequenceErr RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiSequenceErr"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_dlcmiUnknownRpt RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "dlcmiUnknownRpt"

    RFC1315MIB_FrErrTable_FrErrEntry_FrErrType_noErrorSinceReset RFC1315MIB_FrErrTable_FrErrEntry_FrErrType = "noErrorSinceReset"
)

