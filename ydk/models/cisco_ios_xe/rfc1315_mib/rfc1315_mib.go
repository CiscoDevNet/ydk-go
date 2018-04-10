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
    Frdlcmitable RFC1315MIB_Frdlcmitable

    // A table containing information about specific Data Link Connection
    // Identifiers and corresponding virtual circuits.
    Frcircuittable RFC1315MIB_Frcircuittable

    // A table containing information about Errors on the Frame Relay interface.
    Frerrtable RFC1315MIB_Frerrtable
}

func (rFC1315MIB *RFC1315MIB) GetEntityData() *types.CommonEntityData {
    rFC1315MIB.EntityData.YFilter = rFC1315MIB.YFilter
    rFC1315MIB.EntityData.YangName = "RFC1315-MIB"
    rFC1315MIB.EntityData.BundleName = "cisco_ios_xe"
    rFC1315MIB.EntityData.ParentYangName = "RFC1315-MIB"
    rFC1315MIB.EntityData.SegmentPath = "RFC1315-MIB:RFC1315-MIB"
    rFC1315MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rFC1315MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rFC1315MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rFC1315MIB.EntityData.Children = make(map[string]types.YChild)
    rFC1315MIB.EntityData.Children["frame-relay-globals"] = types.YChild{"FrameRelayGlobals", &rFC1315MIB.FrameRelayGlobals}
    rFC1315MIB.EntityData.Children["frDlcmiTable"] = types.YChild{"Frdlcmitable", &rFC1315MIB.Frdlcmitable}
    rFC1315MIB.EntityData.Children["frCircuitTable"] = types.YChild{"Frcircuittable", &rFC1315MIB.Frcircuittable}
    rFC1315MIB.EntityData.Children["frErrTable"] = types.YChild{"Frerrtable", &rFC1315MIB.Frerrtable}
    rFC1315MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rFC1315MIB.EntityData)
}

// RFC1315MIB_FrameRelayGlobals
type RFC1315MIB_FrameRelayGlobals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable  indicates  whether  the  system produces the
    // frDLCIStatusChange trap. The type is Frtrapstate.
    Frtrapstate interface{}
}

func (frameRelayGlobals *RFC1315MIB_FrameRelayGlobals) GetEntityData() *types.CommonEntityData {
    frameRelayGlobals.EntityData.YFilter = frameRelayGlobals.YFilter
    frameRelayGlobals.EntityData.YangName = "frame-relay-globals"
    frameRelayGlobals.EntityData.BundleName = "cisco_ios_xe"
    frameRelayGlobals.EntityData.ParentYangName = "RFC1315-MIB"
    frameRelayGlobals.EntityData.SegmentPath = "frame-relay-globals"
    frameRelayGlobals.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    frameRelayGlobals.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    frameRelayGlobals.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    frameRelayGlobals.EntityData.Children = make(map[string]types.YChild)
    frameRelayGlobals.EntityData.Leafs = make(map[string]types.YLeaf)
    frameRelayGlobals.EntityData.Leafs["frTrapState"] = types.YLeaf{"Frtrapstate", frameRelayGlobals.Frtrapstate}
    return &(frameRelayGlobals.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Parameters for a particular Data Link Con- nection Management
    // Interface. The type is slice of RFC1315MIB_Frdlcmitable_Frdlcmientry.
    Frdlcmientry []RFC1315MIB_Frdlcmitable_Frdlcmientry
}

func (frdlcmitable *RFC1315MIB_Frdlcmitable) GetEntityData() *types.CommonEntityData {
    frdlcmitable.EntityData.YFilter = frdlcmitable.YFilter
    frdlcmitable.EntityData.YangName = "frDlcmiTable"
    frdlcmitable.EntityData.BundleName = "cisco_ios_xe"
    frdlcmitable.EntityData.ParentYangName = "RFC1315-MIB"
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

// RFC1315MIB_Frdlcmitable_Frdlcmientry
// The Parameters for a particular Data Link Con-
// nection Management Interface.
type RFC1315MIB_Frdlcmitable_Frdlcmientry struct {
    EntityData types.CommonEntityData
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

func (frdlcmientry *RFC1315MIB_Frdlcmitable_Frdlcmientry) GetEntityData() *types.CommonEntityData {
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
    return &(frdlcmientry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The information regarding a single  Data  Link Connection Identifier. The
    // type is slice of RFC1315MIB_Frcircuittable_Frcircuitentry.
    Frcircuitentry []RFC1315MIB_Frcircuittable_Frcircuitentry
}

func (frcircuittable *RFC1315MIB_Frcircuittable) GetEntityData() *types.CommonEntityData {
    frcircuittable.EntityData.YFilter = frcircuittable.YFilter
    frcircuittable.EntityData.YangName = "frCircuitTable"
    frcircuittable.EntityData.BundleName = "cisco_ios_xe"
    frcircuittable.EntityData.ParentYangName = "RFC1315-MIB"
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

// RFC1315MIB_Frcircuittable_Frcircuitentry
// The information regarding a single  Data  Link
// Connection Identifier.
type RFC1315MIB_Frcircuittable_Frcircuitentry struct {
    EntityData types.CommonEntityData
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

func (frcircuitentry *RFC1315MIB_Frcircuittable_Frcircuitentry) GetEntityData() *types.CommonEntityData {
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
    return &(frcircuitentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The error information for a single frame relay interface. The type is slice
    // of RFC1315MIB_Frerrtable_Frerrentry.
    Frerrentry []RFC1315MIB_Frerrtable_Frerrentry
}

func (frerrtable *RFC1315MIB_Frerrtable) GetEntityData() *types.CommonEntityData {
    frerrtable.EntityData.YFilter = frerrtable.YFilter
    frerrtable.EntityData.YangName = "frErrTable"
    frerrtable.EntityData.BundleName = "cisco_ios_xe"
    frerrtable.EntityData.ParentYangName = "RFC1315-MIB"
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

// RFC1315MIB_Frerrtable_Frerrentry
// The error information for a single frame relay
// interface.
type RFC1315MIB_Frerrtable_Frerrentry struct {
    EntityData types.CommonEntityData
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

func (frerrentry *RFC1315MIB_Frerrtable_Frerrentry) GetEntityData() *types.CommonEntityData {
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
    return &(frerrentry.EntityData)
}

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

