// This MIB contains managed object definitions for Pseudo 
// Wire operation as in: Pate, P., et al, <draft-ietf-pwe3- 
// framework>, Xiao, X., et al, <draft-ietf-pwe3- 
// requirements>, Martini, L., et al, <draft-martini- 
// l2circuit-trans-mpls>, and Martini, L., et al, 
// <draft-martini-l2circuit-encap-mpls>. 
// 
// The indexes for this MIB are also used to index the PSN- 
// specific tables and the VC-specific tables. The VC Type 
// dictates which VC-specific MIB to use. For example, a  
// 'cep' VC Type requires the use the configuration and status  
// tables within the CEP-MIB. 
// 
// This MIB enable the use of any underlying packet switched  
// network (PSN). Specific tables for the MPLS PSN is 
// currently defined in a separate CISCO-IETF-PW-MPLS-MIB. 
// Tables to support other PSNs (IP, L2TP for example) will 
// be added to this MIB in future revisions. 
// 
// At the time of publication of this version, there are no  
// PWE3 WG documents for all features and objects in this MIB, 
// and the MIB is therefore subject to change based on the WG  
// progress.
package cisco_ietf_pw_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-PW-MIB CISCO-IETF-PW-MIB}", reflect.TypeOf(CISCOIETFPWMIB{}))
    ydk.RegisterEntity("CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB", reflect.TypeOf(CISCOIETFPWMIB{}))
}

// CISCOIETFPWMIB
type CISCOIETFPWMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CpwVcObjects CISCOIETFPWMIB_CpwVcObjects

    // This table specifies information for connecting various  emulated services
    // to various tunnel type.
    CpwVcTable CISCOIETFPWMIB_CpwVcTable

    // This table provides per-VC performance information for the   current
    // interval.
    CpwVcPerfCurrentTable CISCOIETFPWMIB_CpwVcPerfCurrentTable

    // This table provides per-VC performance information for   each interval.
    CpwVcPerfIntervalTable CISCOIETFPWMIB_CpwVcPerfIntervalTable

    // This table provides per-VC Performance information from VC   start time.
    CpwVcPerfTotalTable CISCOIETFPWMIB_CpwVcPerfTotalTable

    // This table provides reverse mapping of the existing VCs   based on vc type
    // and VC ID ordering. This table is   typically useful for EMS ordered query
    // of existing VCs.
    CpwVcIdMappingTable CISCOIETFPWMIB_CpwVcIdMappingTable

    // This table provides reverse mapping of the existing VCs   based on vc type
    // and VC ID ordering. This table is   typically useful for EMS ordered query
    // of existing VCs.
    CpwVcPeerMappingTable CISCOIETFPWMIB_CpwVcPeerMappingTable
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWMIB.EntityData.YFilter = cISCOIETFPWMIB.YFilter
    cISCOIETFPWMIB.EntityData.YangName = "CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWMIB.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.SegmentPath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.AbsolutePath = cISCOIETFPWMIB.EntityData.SegmentPath
    cISCOIETFPWMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcObjects", types.YChild{"CpwVcObjects", &cISCOIETFPWMIB.CpwVcObjects})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcTable", types.YChild{"CpwVcTable", &cISCOIETFPWMIB.CpwVcTable})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcPerfCurrentTable", types.YChild{"CpwVcPerfCurrentTable", &cISCOIETFPWMIB.CpwVcPerfCurrentTable})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcPerfIntervalTable", types.YChild{"CpwVcPerfIntervalTable", &cISCOIETFPWMIB.CpwVcPerfIntervalTable})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcPerfTotalTable", types.YChild{"CpwVcPerfTotalTable", &cISCOIETFPWMIB.CpwVcPerfTotalTable})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcIdMappingTable", types.YChild{"CpwVcIdMappingTable", &cISCOIETFPWMIB.CpwVcIdMappingTable})
    cISCOIETFPWMIB.EntityData.Children.Append("cpwVcPeerMappingTable", types.YChild{"CpwVcPeerMappingTable", &cISCOIETFPWMIB.CpwVcPeerMappingTable})
    cISCOIETFPWMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFPWMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFPWMIB.EntityData)
}

// CISCOIETFPWMIB_CpwVcObjects
type CISCOIETFPWMIB_CpwVcObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an appropriate value to be used  for cpwVcIndex when
    // creating entries in the  cpwVcTable. The value 0 indicates that no 
    // unassigned entries are available.  To obtain the  value of cpwVcIndex for a
    // new entry in the  cpwVcTable, the manager issues a management  protocol
    // retrieval operation to obtain the current  value of cpwVcIndex.  After each
    // retrieval  operation, the agent should modify the value to  reflect the
    // next unassigned index.  After a manager  retrieves a value the agent will
    // determine through  its local policy when this index value will be made 
    // available for reuse. The type is interface{} with range: 0..4294967295.
    CpwVcIndexNext interface{}

    // Counter for number of error at VC level processing, for   example packets
    // received with unknown VC label. The type is interface{} with range:
    // 0..18446744073709551615.
    CpwVcPerfTotalErrorPackets interface{}

    // If this object is set to true(1), then it enables the emission of cpwVcUp
    // and cpwVcDown notifications; otherwise these notifications are not emitted.
    // The type is bool.
    CpwVcUpDownNotifEnable interface{}

    // This object defines the maximum number of PW VC notifications that can be
    // emitted from the device per second. The type is interface{} with range:
    // 0..4294967295.
    CpwVcNotifRate interface{}
}

func (cpwVcObjects *CISCOIETFPWMIB_CpwVcObjects) GetEntityData() *types.CommonEntityData {
    cpwVcObjects.EntityData.YFilter = cpwVcObjects.YFilter
    cpwVcObjects.EntityData.YangName = "cpwVcObjects"
    cpwVcObjects.EntityData.BundleName = "cisco_ios_xe"
    cpwVcObjects.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcObjects.EntityData.SegmentPath = "cpwVcObjects"
    cpwVcObjects.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcObjects.EntityData.SegmentPath
    cpwVcObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcObjects.EntityData.Children = types.NewOrderedMap()
    cpwVcObjects.EntityData.Leafs = types.NewOrderedMap()
    cpwVcObjects.EntityData.Leafs.Append("cpwVcIndexNext", types.YLeaf{"CpwVcIndexNext", cpwVcObjects.CpwVcIndexNext})
    cpwVcObjects.EntityData.Leafs.Append("cpwVcPerfTotalErrorPackets", types.YLeaf{"CpwVcPerfTotalErrorPackets", cpwVcObjects.CpwVcPerfTotalErrorPackets})
    cpwVcObjects.EntityData.Leafs.Append("cpwVcUpDownNotifEnable", types.YLeaf{"CpwVcUpDownNotifEnable", cpwVcObjects.CpwVcUpDownNotifEnable})
    cpwVcObjects.EntityData.Leafs.Append("cpwVcNotifRate", types.YLeaf{"CpwVcNotifRate", cpwVcObjects.CpwVcNotifRate})

    cpwVcObjects.EntityData.YListKeys = []string {}

    return &(cpwVcObjects.EntityData)
}

// CISCOIETFPWMIB_CpwVcTable
// This table specifies information for connecting various 
// emulated services to various tunnel type.
type CISCOIETFPWMIB_CpwVcTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents an emulated virtual  connection (VC) across
    // a packet network. It is indexed by  cpwVcIndex, which uniquely identifying
    // a singular   connection. . The type is slice of
    // CISCOIETFPWMIB_CpwVcTable_CpwVcEntry.
    CpwVcEntry []*CISCOIETFPWMIB_CpwVcTable_CpwVcEntry
}

func (cpwVcTable *CISCOIETFPWMIB_CpwVcTable) GetEntityData() *types.CommonEntityData {
    cpwVcTable.EntityData.YFilter = cpwVcTable.YFilter
    cpwVcTable.EntityData.YangName = "cpwVcTable"
    cpwVcTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcTable.EntityData.SegmentPath = "cpwVcTable"
    cpwVcTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcTable.EntityData.SegmentPath
    cpwVcTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcTable.EntityData.Children = types.NewOrderedMap()
    cpwVcTable.EntityData.Children.Append("cpwVcEntry", types.YChild{"CpwVcEntry", nil})
    for i := range cpwVcTable.CpwVcEntry {
        cpwVcTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcTable.CpwVcEntry[i]), types.YChild{"CpwVcEntry", cpwVcTable.CpwVcEntry[i]})
    }
    cpwVcTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcTable.EntityData.YListKeys = []string {}

    return &(cpwVcTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry
// A row in this table represents an emulated virtual 
// connection (VC) across a packet network. It is indexed by 
// cpwVcIndex, which uniquely identifying a singular  
// connection. 
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index for the conceptual row identifying a VC
    // within   this PW Emulation VC table. The type is interface{} with range:
    // 0..4294967295.
    CpwVcIndex interface{}

    // This value indicate the service to be carried over  this VC.   Note: the
    // exact set of VC types is yet to be worked   out by the WG. . The type is
    // CpwVcType.
    CpwVcType interface{}

    // Set by the operator to indicate the protocol responsible   for establishing
    // this VC. Value 'manual' is used in all  cases where no maintenance protocol
    // (PW signaling) is used   to set-up the VC, i.e. require configuration of
    // entries in  the VC tables including VC labels, etc. The value  
    // 'maintenanceProtocol' is used in case of standard  signaling of the VC for
    // the specific PSN, for example LDP  for MPLS PSN as specified in <draft-
    // draft-martini-  l2circuit-trans-mpls> or L2TP control protocol.   Value
    // 'other' is used for other types of signaling. The type is CpwVcOwner.
    CpwVcOwner interface{}

    // Set by the operator to indicate the PSN type on which this   VC will be
    // carried. Based on this object, the relevant PSN   table entries are created
    // in the in the PSN specific MIB   modules. For example, if mpls(1) is
    // defined, the agent   create an entry in cpwVcMplsTable, which further
    // define the   MPLS PSN configuration.  Note: the exact set of PSN types is
    // yet to be worked   out by the WG. . The type is CpwVcPsnType.
    CpwVcPsnType interface{}

    // This object define the relative set-up priority of the VC    in a
    // lowest-to-highest fashion, where 0 is the highest   priority. VCs with the
    // same priority are treated with  equal priority. Dropped VC will be set
    // 'dormant' (as  indicated in cpwVcOperStatus).  This value is significant if
    // there are competing resources  between VCs and the implementation support
    // this feature.  If not supported or not relevant, the value of zero MUST  be
    // used. The type is interface{} with range: 0..7.
    CpwVcSetUpPriority interface{}

    // This object define the relative holding priority of the VC    in a
    // lowest-to-highest fashion, where 0 is the highest   priority. VCs with the
    // same priority are treated with  equal priority. Dropped VC will be set
    // 'dormant' (as  indicated in cpwVcOperStatus).  This value is significant if
    // there are competing resources  between VCs and the implementation support
    // this feature.  If not supported or not relevant, the value of zero MUST  be
    // used. The type is interface{} with range: 0..7.
    CpwVcHoldingPriority interface{}

    // This object is used to enable greater security for   implementation that
    // use per platform VC label space. In   strict mode, packets coming from the
    // PSN are accepted only   from tunnels that are associated to the same VC via
    // the   inbound tunnel table in the case of MPLS, or as identified   by the
    // source IP address in case of L2TP or IP PSN. The   entries in the inbound
    // tunnel table are either explicitly   configured or implicitly known by the
    // maintenance protocol   used for VC set-up.   If such association is not
    // known, not configured or not   desired, loose mode should be configured,
    // and the node   should accept the packet based on the VC label only  
    // regardless of the outer tunnel used to carry the VC. The type is
    // CpwVcInboundMode.
    CpwVcInboundMode interface{}

    // Denotes the address type of the peer node maintenance  protocol (signaling)
    // address if PW maintenance protocol is  used for the VC creation. It should
    // be set to   'unknown' if PE/PW maintenance protocol is not used,   i.e.
    // cpwVcOwner is set to 'manual'. . The type is InetAddressType.
    CpwVcPeerAddrType interface{}

    // This object contains the value of of the peer node address  of the PW/PE
    // maintenance protocol entity. This object   should contain a value of 0 if
    // not relevant (manual   configuration of the VC). The type is string with
    // length: 0..255.
    CpwVcPeerAddr interface{}

    // Used in the outgoing VC ID field within the 'Virtual  Circuit FEC Element'
    // when LDP signaling is used or PW ID   AVP for L2TP. The type is interface{}
    // with range: 0..4294967295.
    CpwVcID interface{}

    // Used in the Group ID field sent to the peer PWES   within the maintenance
    // protocol used for VC setup,   zero if not used. The type is interface{}
    // with range: 0..4294967295.
    CpwVcLocalGroupID interface{}

    // Define if the control word will be sent with each packet by   the local
    // node. The type is bool.
    CpwVcControlWord interface{}

    // If not equal zero, the optional IfMtu object in the   maintenance protocol
    // will be sent with this value,   representing the locally supported MTU size
    // over the   interface (or the virtual interface) associated with the   VC.
    // The type is interface{} with range: 0..65535.
    CpwVcLocalIfMtu interface{}

    // Each VC is associated to an interface (or a virtual   interface) in the
    // ifTable of the node as part of the  service configuration. This object
    // defines if the   maintenance protocol will send the interface's name as 
    // appears on the ifTable in the name object as part of the  maintenance
    // protocol. If set to false, the optional element  will not be sent. The type
    // is bool.
    CpwVcLocalIfString interface{}

    // Obtained from the Group ID field as received via the   maintenance protocol
    // used for VC setup, zero if not used.   Value of 0xFFFF shall be used if the
    // object is yet to be   defined by the VC maintenance protocol. The type is
    // interface{} with range: 0..4294967295.
    CpwVcRemoteGroupID interface{}

    // If maintenance protocol is used for VC establishment, this   parameter
    // indicates the received status of the control word   usage, i.e. if packets
    // will be received with control word  or not. The value of 'notYetKnown' is
    // used while the  maintenance protocol has not yet received the indication  
    // from the remote node.  In manual configuration of the VC this parameters
    // indicate   to the local node what is the expected encapsulation for  the
    // received packets. . The type is CpwVcRemoteControlWord.
    CpwVcRemoteControlWord interface{}

    // The remote interface MTU as (optionally) received from the  remote node via
    // the maintenance protocol. Should be zero if  this parameter is not
    // available or not used. The type is interface{} with range: 0..4294967295.
    CpwVcRemoteIfMtu interface{}

    // Indicate the interface description string as received by  the maintenance
    // protocol, MUST be NULL string if not   applicable or not known yet. The
    // type is string with length: 0..80.
    CpwVcRemoteIfString interface{}

    // The VC label used in the outbound direction (i.e. toward   the PSN). It may
    // be set up manually if owner is 'manual' or   automatically otherwise.
    // Examples: For MPLS PSN, it   represents the 20 bits of VC tag, for L2TP it
    // represent the  32 bits Session ID.  If the label is not yet known
    // (signaling in process), the   object should return a value of 0xFFFF. The
    // type is interface{} with range: 0..4294967295.
    CpwVcOutboundVcLabel interface{}

    // The VC label used in the inbound direction (i.e. packets   received from
    // the PSN. It may be set up manually if owner  is 'manual' or automatically
    // otherwise.   Examples: For MPLS PSN, it represents the 20 bits of VC  tag,
    // for L2TP it represent the 32 bits Session ID.  If the label is not yet
    // known (signaling in process), the   object should return a value of 0xFFFF.
    // The type is interface{} with range: 0..4294967295.
    CpwVcInboundVcLabel interface{}

    // The canonical name assigned to the VC. The type is string.
    CpwVcName interface{}

    // A textual string containing information about the VC.   If there is no
    // description this object contains a zero  length string. The type is string.
    CpwVcDescr interface{}

    // System time when this VC was created. The type is interface{} with range:
    // 0..4294967295.
    CpwVcCreateTime interface{}

    // Number of consecutive ticks this VC has been 'up' in  both directions
    // together (i.e. 'up' is observed in   cpwVcOperStatus.). The type is
    // interface{} with range: 0..4294967295.
    CpwVcUpTime interface{}

    // The desired operational status of this VC. The type is CpwVcAdminStatus.
    CpwVcAdminStatus interface{}

    // Indicates the actual combined operational status of this   VC. It is 'up'
    // if both cpwVcInboundOperStatus and   cpwVcOutboundOperStatus are in 'up'
    // state. For all other   values, if the VCs in both directions are of the
    // same  value it reflects that value, otherwise it is set to the  most severe
    // status out of the two statuses. The order of   severance from most severe
    // to less severe is: unknown,   notPresent, down, lowerLayerDown, dormant,
    // testing, up.  The operator may consult the per direction OperStatus for 
    // fault isolation per direction. The type is CpwOperStatus.
    CpwVcOperStatus interface{}

    // Indicates the actual operational status of this VC in the   inbound
    // direction.   - down:           if PW signaling has not yet finished, or    
    // indications available at the service                     level indicate
    // that the VC is not                     passing packets.  - testing:       
    // if AdminStatus at the VC level is set to                     test.  -
    // dormant:        The VC is not available because of the                   
    // required resources are occupied VC with                     higher priority
    // VCs .  - notPresent:     Some component is missing to accomplish           
    // the set up of the VC.  - lowerLayerDown: The underlying PSN is not in
    // OperStatus                     'up'.  . The type is CpwOperStatus.
    CpwVcInboundOperStatus interface{}

    // Indicates the actual operational status of this VC in the   outbound
    // direction  - down:           if PW signaling has not yet finished, or      
    // indications available at the service                     level indicate
    // that the VC is not                    passing packets.  - testing:       
    // if AdminStatus at the VC level is set to                     test.  -
    // dormant:        The VC is not available because of the                   
    // required resources are occupied VC with                     higher priority
    // VCs .  - notPresent:     Some component is missing to accomplish           
    // the set up of the VC.  - lowerLayerDown: The underlying PSN is not in
    // OperStatus                     'up'.  . The type is CpwOperStatus.
    CpwVcOutboundOperStatus interface{}

    // The number of seconds, including partial seconds,  that have elapsed since
    // the beginning of the current  measurement period. If, for some reason, such
    // as an  adjustment in the system's time-of-day clock, the  current interval
    // exceeds the maximum value, the  agent will return the maximum value. The
    // type is interface{} with range: 1..900.
    CpwVcTimeElapsed interface{}

    // The number of previous 15-minute intervals  for which data was collected.  
    // An agent with PW capability must be capable of supporting at   least n
    // intervals. The minimum value of n is 4, The default   of n is 32 and the
    // maximum value of n is 96.  The value will be <n> unless the measurement was
    // (re-)   started within the last (<n>*15) minutes, in which case the   value
    // will be the number of complete 15 minute intervals for   which the agent
    // has at least some data. In certain cases   (e.g., in the case where the
    // agent is a proxy) it is   possible that some intervals are unavailable.  In
    // this case,   this interval is the maximum interval number for which data  
    // is available. . The type is interface{} with range: 0..96.
    CpwVcValidIntervals interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    CpwVcRowStatus interface{}

    // This variable indicates the storage type for this  object. The type is
    // StorageType.
    CpwVcStorageType interface{}
}

func (cpwVcEntry *CISCOIETFPWMIB_CpwVcTable_CpwVcEntry) GetEntityData() *types.CommonEntityData {
    cpwVcEntry.EntityData.YFilter = cpwVcEntry.YFilter
    cpwVcEntry.EntityData.YangName = "cpwVcEntry"
    cpwVcEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcEntry.EntityData.ParentYangName = "cpwVcTable"
    cpwVcEntry.EntityData.SegmentPath = "cpwVcEntry" + types.AddKeyToken(cpwVcEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcTable/" + cpwVcEntry.EntityData.SegmentPath
    cpwVcEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcEntry.CpwVcIndex})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcType", types.YLeaf{"CpwVcType", cpwVcEntry.CpwVcType})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcOwner", types.YLeaf{"CpwVcOwner", cpwVcEntry.CpwVcOwner})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcPsnType", types.YLeaf{"CpwVcPsnType", cpwVcEntry.CpwVcPsnType})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcSetUpPriority", types.YLeaf{"CpwVcSetUpPriority", cpwVcEntry.CpwVcSetUpPriority})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcHoldingPriority", types.YLeaf{"CpwVcHoldingPriority", cpwVcEntry.CpwVcHoldingPriority})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcInboundMode", types.YLeaf{"CpwVcInboundMode", cpwVcEntry.CpwVcInboundMode})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcPeerAddrType", types.YLeaf{"CpwVcPeerAddrType", cpwVcEntry.CpwVcPeerAddrType})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcPeerAddr", types.YLeaf{"CpwVcPeerAddr", cpwVcEntry.CpwVcPeerAddr})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcID", types.YLeaf{"CpwVcID", cpwVcEntry.CpwVcID})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcLocalGroupID", types.YLeaf{"CpwVcLocalGroupID", cpwVcEntry.CpwVcLocalGroupID})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcControlWord", types.YLeaf{"CpwVcControlWord", cpwVcEntry.CpwVcControlWord})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcLocalIfMtu", types.YLeaf{"CpwVcLocalIfMtu", cpwVcEntry.CpwVcLocalIfMtu})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcLocalIfString", types.YLeaf{"CpwVcLocalIfString", cpwVcEntry.CpwVcLocalIfString})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcRemoteGroupID", types.YLeaf{"CpwVcRemoteGroupID", cpwVcEntry.CpwVcRemoteGroupID})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcRemoteControlWord", types.YLeaf{"CpwVcRemoteControlWord", cpwVcEntry.CpwVcRemoteControlWord})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcRemoteIfMtu", types.YLeaf{"CpwVcRemoteIfMtu", cpwVcEntry.CpwVcRemoteIfMtu})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcRemoteIfString", types.YLeaf{"CpwVcRemoteIfString", cpwVcEntry.CpwVcRemoteIfString})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcOutboundVcLabel", types.YLeaf{"CpwVcOutboundVcLabel", cpwVcEntry.CpwVcOutboundVcLabel})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcInboundVcLabel", types.YLeaf{"CpwVcInboundVcLabel", cpwVcEntry.CpwVcInboundVcLabel})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcName", types.YLeaf{"CpwVcName", cpwVcEntry.CpwVcName})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcDescr", types.YLeaf{"CpwVcDescr", cpwVcEntry.CpwVcDescr})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcCreateTime", types.YLeaf{"CpwVcCreateTime", cpwVcEntry.CpwVcCreateTime})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcUpTime", types.YLeaf{"CpwVcUpTime", cpwVcEntry.CpwVcUpTime})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcAdminStatus", types.YLeaf{"CpwVcAdminStatus", cpwVcEntry.CpwVcAdminStatus})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcOperStatus", types.YLeaf{"CpwVcOperStatus", cpwVcEntry.CpwVcOperStatus})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcInboundOperStatus", types.YLeaf{"CpwVcInboundOperStatus", cpwVcEntry.CpwVcInboundOperStatus})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcOutboundOperStatus", types.YLeaf{"CpwVcOutboundOperStatus", cpwVcEntry.CpwVcOutboundOperStatus})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcTimeElapsed", types.YLeaf{"CpwVcTimeElapsed", cpwVcEntry.CpwVcTimeElapsed})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcValidIntervals", types.YLeaf{"CpwVcValidIntervals", cpwVcEntry.CpwVcValidIntervals})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcRowStatus", types.YLeaf{"CpwVcRowStatus", cpwVcEntry.CpwVcRowStatus})
    cpwVcEntry.EntityData.Leafs.Append("cpwVcStorageType", types.YLeaf{"CpwVcStorageType", cpwVcEntry.CpwVcStorageType})

    cpwVcEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcEntry.EntityData)
}

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus represents The desired operational status of this VC.
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus string

const (
    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus_up CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus = "up"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus_down CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus = "down"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus_testing CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcAdminStatus = "testing"
)

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode represents regardless of the outer tunnel used to carry the VC.
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode string

const (
    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode_loose CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode = "loose"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode_strict CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcInboundMode = "strict"
)

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner represents Value 'other' is used for other types of signaling.
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner string

const (
    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner_manual CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner = "manual"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner_maintenanceProtocol CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner = "maintenanceProtocol"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner_other CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcOwner = "other"
)

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType represents out by the WG. 
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType string

const (
    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_mpls CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "mpls"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_l2tp CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "l2tp"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_ip CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "ip"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_mplsOverIp CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "mplsOverIp"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_gre CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "gre"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType_other CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcPsnType = "other"
)

// CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord represents the received packets. 
type CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord string

const (
    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord_noControlWord CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord = "noControlWord"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord_withControlWord CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord = "withControlWord"

    CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord_notYetKnown CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcRemoteControlWord = "notYetKnown"
)

// CISCOIETFPWMIB_CpwVcPerfCurrentTable
// This table provides per-VC performance information for the  
// current interval.
type CISCOIETFPWMIB_CpwVcPerfCurrentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for  every VC. The type is
    // slice of CISCOIETFPWMIB_CpwVcPerfCurrentTable_CpwVcPerfCurrentEntry.
    CpwVcPerfCurrentEntry []*CISCOIETFPWMIB_CpwVcPerfCurrentTable_CpwVcPerfCurrentEntry
}

func (cpwVcPerfCurrentTable *CISCOIETFPWMIB_CpwVcPerfCurrentTable) GetEntityData() *types.CommonEntityData {
    cpwVcPerfCurrentTable.EntityData.YFilter = cpwVcPerfCurrentTable.YFilter
    cpwVcPerfCurrentTable.EntityData.YangName = "cpwVcPerfCurrentTable"
    cpwVcPerfCurrentTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfCurrentTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcPerfCurrentTable.EntityData.SegmentPath = "cpwVcPerfCurrentTable"
    cpwVcPerfCurrentTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcPerfCurrentTable.EntityData.SegmentPath
    cpwVcPerfCurrentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfCurrentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfCurrentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfCurrentTable.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfCurrentTable.EntityData.Children.Append("cpwVcPerfCurrentEntry", types.YChild{"CpwVcPerfCurrentEntry", nil})
    for i := range cpwVcPerfCurrentTable.CpwVcPerfCurrentEntry {
        cpwVcPerfCurrentTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcPerfCurrentTable.CpwVcPerfCurrentEntry[i]), types.YChild{"CpwVcPerfCurrentEntry", cpwVcPerfCurrentTable.CpwVcPerfCurrentEntry[i]})
    }
    cpwVcPerfCurrentTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcPerfCurrentTable.EntityData.YListKeys = []string {}

    return &(cpwVcPerfCurrentTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcPerfCurrentTable_CpwVcPerfCurrentEntry
// An entry in this table is created by the agent for 
// every VC.
type CISCOIETFPWMIB_CpwVcPerfCurrentTable_CpwVcPerfCurrentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // High capacity counter for number of packets received  by the VC (from the
    // PSN) in the current 15 minute  interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfCurrentInHCPackets interface{}

    // High capacity counter for number of bytes received  by the VC (from the
    // PSN) in the current 15 minute  interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfCurrentInHCBytes interface{}

    // High capacity counter for number of packets forwarded  by the VC (to the
    // PSN) in the current 15 minute interval. The type is interface{} with range:
    // 0..18446744073709551615.
    CpwVcPerfCurrentOutHCPackets interface{}

    // High capacity counter for number of bytes forwarded  by the VC (to the PSN)
    // in the current 15 minute interval. The type is interface{} with range:
    // 0..18446744073709551615.
    CpwVcPerfCurrentOutHCBytes interface{}
}

func (cpwVcPerfCurrentEntry *CISCOIETFPWMIB_CpwVcPerfCurrentTable_CpwVcPerfCurrentEntry) GetEntityData() *types.CommonEntityData {
    cpwVcPerfCurrentEntry.EntityData.YFilter = cpwVcPerfCurrentEntry.YFilter
    cpwVcPerfCurrentEntry.EntityData.YangName = "cpwVcPerfCurrentEntry"
    cpwVcPerfCurrentEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfCurrentEntry.EntityData.ParentYangName = "cpwVcPerfCurrentTable"
    cpwVcPerfCurrentEntry.EntityData.SegmentPath = "cpwVcPerfCurrentEntry" + types.AddKeyToken(cpwVcPerfCurrentEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcPerfCurrentEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcPerfCurrentTable/" + cpwVcPerfCurrentEntry.EntityData.SegmentPath
    cpwVcPerfCurrentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfCurrentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfCurrentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfCurrentEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfCurrentEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcPerfCurrentEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcPerfCurrentEntry.CpwVcIndex})
    cpwVcPerfCurrentEntry.EntityData.Leafs.Append("cpwVcPerfCurrentInHCPackets", types.YLeaf{"CpwVcPerfCurrentInHCPackets", cpwVcPerfCurrentEntry.CpwVcPerfCurrentInHCPackets})
    cpwVcPerfCurrentEntry.EntityData.Leafs.Append("cpwVcPerfCurrentInHCBytes", types.YLeaf{"CpwVcPerfCurrentInHCBytes", cpwVcPerfCurrentEntry.CpwVcPerfCurrentInHCBytes})
    cpwVcPerfCurrentEntry.EntityData.Leafs.Append("cpwVcPerfCurrentOutHCPackets", types.YLeaf{"CpwVcPerfCurrentOutHCPackets", cpwVcPerfCurrentEntry.CpwVcPerfCurrentOutHCPackets})
    cpwVcPerfCurrentEntry.EntityData.Leafs.Append("cpwVcPerfCurrentOutHCBytes", types.YLeaf{"CpwVcPerfCurrentOutHCBytes", cpwVcPerfCurrentEntry.CpwVcPerfCurrentOutHCBytes})

    cpwVcPerfCurrentEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcPerfCurrentEntry.EntityData)
}

// CISCOIETFPWMIB_CpwVcPerfIntervalTable
// This table provides per-VC performance information for  
// each interval.
type CISCOIETFPWMIB_CpwVcPerfIntervalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_CpwVcPerfIntervalTable_CpwVcPerfIntervalEntry.
    CpwVcPerfIntervalEntry []*CISCOIETFPWMIB_CpwVcPerfIntervalTable_CpwVcPerfIntervalEntry
}

func (cpwVcPerfIntervalTable *CISCOIETFPWMIB_CpwVcPerfIntervalTable) GetEntityData() *types.CommonEntityData {
    cpwVcPerfIntervalTable.EntityData.YFilter = cpwVcPerfIntervalTable.YFilter
    cpwVcPerfIntervalTable.EntityData.YangName = "cpwVcPerfIntervalTable"
    cpwVcPerfIntervalTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfIntervalTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcPerfIntervalTable.EntityData.SegmentPath = "cpwVcPerfIntervalTable"
    cpwVcPerfIntervalTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcPerfIntervalTable.EntityData.SegmentPath
    cpwVcPerfIntervalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfIntervalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfIntervalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfIntervalTable.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfIntervalTable.EntityData.Children.Append("cpwVcPerfIntervalEntry", types.YChild{"CpwVcPerfIntervalEntry", nil})
    for i := range cpwVcPerfIntervalTable.CpwVcPerfIntervalEntry {
        cpwVcPerfIntervalTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcPerfIntervalTable.CpwVcPerfIntervalEntry[i]), types.YChild{"CpwVcPerfIntervalEntry", cpwVcPerfIntervalTable.CpwVcPerfIntervalEntry[i]})
    }
    cpwVcPerfIntervalTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcPerfIntervalTable.EntityData.YListKeys = []string {}

    return &(cpwVcPerfIntervalTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcPerfIntervalTable_CpwVcPerfIntervalEntry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_CpwVcPerfIntervalTable_CpwVcPerfIntervalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // This attribute is a key. A number N, between 1 and 96, which identifies the
    // interval for which the set of statistics is available.  The interval
    // identified by 1 is the most recently  completed 15 minute interval, and the
    // interval identified  by N is the interval immediately preceding the one 
    // identified by N-1.  The minimum range of N is 1 through 4. The default
    // range  is 1 to 32. The maximum range of N is 1 through 96. . The type is
    // interface{} with range: 1..96.
    CpwVcPerfIntervalNumber interface{}

    // This variable indicates if the data for this interval  is valid. The type
    // is bool.
    CpwVcPerfIntervalValidData interface{}

    // The duration of a particular interval in seconds.  Adjustments in the
    // system's time-of-day clock, may  cause the interval to be greater or less
    // than the  normal value. Therefore this actual interval value  is provided.
    // The type is interface{} with range: -2147483648..2147483647.
    CpwVcPerfIntervalTimeElapsed interface{}

    // High capacity counter for number of packets received by  the VC (from the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfIntervalInHCPackets interface{}

    // High capacity counter for number of bytes received by the   VC (from the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfIntervalInHCBytes interface{}

    // High capacity counter for number of packets forwarded by   the VC (to the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfIntervalOutHCPackets interface{}

    // High capacity counter for number of bytes forwarded by the   VC (to the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    CpwVcPerfIntervalOutHCBytes interface{}
}

func (cpwVcPerfIntervalEntry *CISCOIETFPWMIB_CpwVcPerfIntervalTable_CpwVcPerfIntervalEntry) GetEntityData() *types.CommonEntityData {
    cpwVcPerfIntervalEntry.EntityData.YFilter = cpwVcPerfIntervalEntry.YFilter
    cpwVcPerfIntervalEntry.EntityData.YangName = "cpwVcPerfIntervalEntry"
    cpwVcPerfIntervalEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfIntervalEntry.EntityData.ParentYangName = "cpwVcPerfIntervalTable"
    cpwVcPerfIntervalEntry.EntityData.SegmentPath = "cpwVcPerfIntervalEntry" + types.AddKeyToken(cpwVcPerfIntervalEntry.CpwVcIndex, "cpwVcIndex") + types.AddKeyToken(cpwVcPerfIntervalEntry.CpwVcPerfIntervalNumber, "cpwVcPerfIntervalNumber")
    cpwVcPerfIntervalEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcPerfIntervalTable/" + cpwVcPerfIntervalEntry.EntityData.SegmentPath
    cpwVcPerfIntervalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfIntervalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfIntervalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfIntervalEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfIntervalEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcPerfIntervalEntry.CpwVcIndex})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalNumber", types.YLeaf{"CpwVcPerfIntervalNumber", cpwVcPerfIntervalEntry.CpwVcPerfIntervalNumber})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalValidData", types.YLeaf{"CpwVcPerfIntervalValidData", cpwVcPerfIntervalEntry.CpwVcPerfIntervalValidData})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalTimeElapsed", types.YLeaf{"CpwVcPerfIntervalTimeElapsed", cpwVcPerfIntervalEntry.CpwVcPerfIntervalTimeElapsed})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalInHCPackets", types.YLeaf{"CpwVcPerfIntervalInHCPackets", cpwVcPerfIntervalEntry.CpwVcPerfIntervalInHCPackets})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalInHCBytes", types.YLeaf{"CpwVcPerfIntervalInHCBytes", cpwVcPerfIntervalEntry.CpwVcPerfIntervalInHCBytes})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalOutHCPackets", types.YLeaf{"CpwVcPerfIntervalOutHCPackets", cpwVcPerfIntervalEntry.CpwVcPerfIntervalOutHCPackets})
    cpwVcPerfIntervalEntry.EntityData.Leafs.Append("cpwVcPerfIntervalOutHCBytes", types.YLeaf{"CpwVcPerfIntervalOutHCBytes", cpwVcPerfIntervalEntry.CpwVcPerfIntervalOutHCBytes})

    cpwVcPerfIntervalEntry.EntityData.YListKeys = []string {"CpwVcIndex", "CpwVcPerfIntervalNumber"}

    return &(cpwVcPerfIntervalEntry.EntityData)
}

// CISCOIETFPWMIB_CpwVcPerfTotalTable
// This table provides per-VC Performance information from VC  
// start time.
type CISCOIETFPWMIB_CpwVcPerfTotalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_CpwVcPerfTotalTable_CpwVcPerfTotalEntry.
    CpwVcPerfTotalEntry []*CISCOIETFPWMIB_CpwVcPerfTotalTable_CpwVcPerfTotalEntry
}

func (cpwVcPerfTotalTable *CISCOIETFPWMIB_CpwVcPerfTotalTable) GetEntityData() *types.CommonEntityData {
    cpwVcPerfTotalTable.EntityData.YFilter = cpwVcPerfTotalTable.YFilter
    cpwVcPerfTotalTable.EntityData.YangName = "cpwVcPerfTotalTable"
    cpwVcPerfTotalTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfTotalTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcPerfTotalTable.EntityData.SegmentPath = "cpwVcPerfTotalTable"
    cpwVcPerfTotalTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcPerfTotalTable.EntityData.SegmentPath
    cpwVcPerfTotalTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfTotalTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfTotalTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfTotalTable.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfTotalTable.EntityData.Children.Append("cpwVcPerfTotalEntry", types.YChild{"CpwVcPerfTotalEntry", nil})
    for i := range cpwVcPerfTotalTable.CpwVcPerfTotalEntry {
        cpwVcPerfTotalTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcPerfTotalTable.CpwVcPerfTotalEntry[i]), types.YChild{"CpwVcPerfTotalEntry", cpwVcPerfTotalTable.CpwVcPerfTotalEntry[i]})
    }
    cpwVcPerfTotalTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcPerfTotalTable.EntityData.YListKeys = []string {}

    return &(cpwVcPerfTotalTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcPerfTotalTable_CpwVcPerfTotalEntry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_CpwVcPerfTotalTable_CpwVcPerfTotalEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_CpwVcTable_CpwVcEntry_CpwVcIndex
    CpwVcIndex interface{}

    // High capacity counter for number of packets received by the   VC (from the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    CpwVcPerfTotalInHCPackets interface{}

    // High capacity counter for number of bytes received by the   VC (from the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    CpwVcPerfTotalInHCBytes interface{}

    // High capacity counter for number of packets forwarded by   the VC (to the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    CpwVcPerfTotalOutHCPackets interface{}

    // High capacity counter for number of bytes forwarded by the   VC (to the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    CpwVcPerfTotalOutHCBytes interface{}

    // The value of sysUpTime on the most recent occasion at  which any one or
    // more of this row Counter32 or  Counter64 suffered a discontinuity. If no
    // such  discontinuities have occurred since the last re-  initialization of
    // the local management subsystem, then  this object contains a zero value.
    // The type is interface{} with range: 0..4294967295.
    CpwVcPerfTotalDiscontinuityTime interface{}
}

func (cpwVcPerfTotalEntry *CISCOIETFPWMIB_CpwVcPerfTotalTable_CpwVcPerfTotalEntry) GetEntityData() *types.CommonEntityData {
    cpwVcPerfTotalEntry.EntityData.YFilter = cpwVcPerfTotalEntry.YFilter
    cpwVcPerfTotalEntry.EntityData.YangName = "cpwVcPerfTotalEntry"
    cpwVcPerfTotalEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPerfTotalEntry.EntityData.ParentYangName = "cpwVcPerfTotalTable"
    cpwVcPerfTotalEntry.EntityData.SegmentPath = "cpwVcPerfTotalEntry" + types.AddKeyToken(cpwVcPerfTotalEntry.CpwVcIndex, "cpwVcIndex")
    cpwVcPerfTotalEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcPerfTotalTable/" + cpwVcPerfTotalEntry.EntityData.SegmentPath
    cpwVcPerfTotalEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPerfTotalEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPerfTotalEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPerfTotalEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcPerfTotalEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcIndex", types.YLeaf{"CpwVcIndex", cpwVcPerfTotalEntry.CpwVcIndex})
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcPerfTotalInHCPackets", types.YLeaf{"CpwVcPerfTotalInHCPackets", cpwVcPerfTotalEntry.CpwVcPerfTotalInHCPackets})
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcPerfTotalInHCBytes", types.YLeaf{"CpwVcPerfTotalInHCBytes", cpwVcPerfTotalEntry.CpwVcPerfTotalInHCBytes})
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcPerfTotalOutHCPackets", types.YLeaf{"CpwVcPerfTotalOutHCPackets", cpwVcPerfTotalEntry.CpwVcPerfTotalOutHCPackets})
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcPerfTotalOutHCBytes", types.YLeaf{"CpwVcPerfTotalOutHCBytes", cpwVcPerfTotalEntry.CpwVcPerfTotalOutHCBytes})
    cpwVcPerfTotalEntry.EntityData.Leafs.Append("cpwVcPerfTotalDiscontinuityTime", types.YLeaf{"CpwVcPerfTotalDiscontinuityTime", cpwVcPerfTotalEntry.CpwVcPerfTotalDiscontinuityTime})

    cpwVcPerfTotalEntry.EntityData.YListKeys = []string {"CpwVcIndex"}

    return &(cpwVcPerfTotalEntry.EntityData)
}

// CISCOIETFPWMIB_CpwVcIdMappingTable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_CpwVcIdMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured by
    // the cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_CpwVcIdMappingTable_CpwVcIdMappingEntry.
    CpwVcIdMappingEntry []*CISCOIETFPWMIB_CpwVcIdMappingTable_CpwVcIdMappingEntry
}

func (cpwVcIdMappingTable *CISCOIETFPWMIB_CpwVcIdMappingTable) GetEntityData() *types.CommonEntityData {
    cpwVcIdMappingTable.EntityData.YFilter = cpwVcIdMappingTable.YFilter
    cpwVcIdMappingTable.EntityData.YangName = "cpwVcIdMappingTable"
    cpwVcIdMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcIdMappingTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcIdMappingTable.EntityData.SegmentPath = "cpwVcIdMappingTable"
    cpwVcIdMappingTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcIdMappingTable.EntityData.SegmentPath
    cpwVcIdMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcIdMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcIdMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcIdMappingTable.EntityData.Children = types.NewOrderedMap()
    cpwVcIdMappingTable.EntityData.Children.Append("cpwVcIdMappingEntry", types.YChild{"CpwVcIdMappingEntry", nil})
    for i := range cpwVcIdMappingTable.CpwVcIdMappingEntry {
        cpwVcIdMappingTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcIdMappingTable.CpwVcIdMappingEntry[i]), types.YChild{"CpwVcIdMappingEntry", cpwVcIdMappingTable.CpwVcIdMappingEntry[i]})
    }
    cpwVcIdMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcIdMappingTable.EntityData.YListKeys = []string {}

    return &(cpwVcIdMappingTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcIdMappingTable_CpwVcIdMappingEntry
// An entry in this table is created by the agent for every  
// VC configured by the cpwVcTable.
type CISCOIETFPWMIB_CpwVcIdMappingTable_CpwVcIdMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The VC type (indicate the service) of this VC. The
    // type is CpwVcType.
    CpwVcIdMappingVcType interface{}

    // This attribute is a key. The VC ID of this VC. Zero if the VC is configured
    // manually. The type is interface{} with range: 0..4294967295.
    CpwVcIdMappingVcID interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // InetAddressType.
    CpwVcIdMappingPeerAddrType interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // string with length: 0..255.
    CpwVcIdMappingPeerAddr interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    CpwVcIdMappingVcIndex interface{}
}

func (cpwVcIdMappingEntry *CISCOIETFPWMIB_CpwVcIdMappingTable_CpwVcIdMappingEntry) GetEntityData() *types.CommonEntityData {
    cpwVcIdMappingEntry.EntityData.YFilter = cpwVcIdMappingEntry.YFilter
    cpwVcIdMappingEntry.EntityData.YangName = "cpwVcIdMappingEntry"
    cpwVcIdMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcIdMappingEntry.EntityData.ParentYangName = "cpwVcIdMappingTable"
    cpwVcIdMappingEntry.EntityData.SegmentPath = "cpwVcIdMappingEntry" + types.AddKeyToken(cpwVcIdMappingEntry.CpwVcIdMappingVcType, "cpwVcIdMappingVcType") + types.AddKeyToken(cpwVcIdMappingEntry.CpwVcIdMappingVcID, "cpwVcIdMappingVcID") + types.AddKeyToken(cpwVcIdMappingEntry.CpwVcIdMappingPeerAddrType, "cpwVcIdMappingPeerAddrType") + types.AddKeyToken(cpwVcIdMappingEntry.CpwVcIdMappingPeerAddr, "cpwVcIdMappingPeerAddr") + types.AddKeyToken(cpwVcIdMappingEntry.CpwVcIdMappingVcIndex, "cpwVcIdMappingVcIndex")
    cpwVcIdMappingEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcIdMappingTable/" + cpwVcIdMappingEntry.EntityData.SegmentPath
    cpwVcIdMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcIdMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcIdMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcIdMappingEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcIdMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcIdMappingEntry.EntityData.Leafs.Append("cpwVcIdMappingVcType", types.YLeaf{"CpwVcIdMappingVcType", cpwVcIdMappingEntry.CpwVcIdMappingVcType})
    cpwVcIdMappingEntry.EntityData.Leafs.Append("cpwVcIdMappingVcID", types.YLeaf{"CpwVcIdMappingVcID", cpwVcIdMappingEntry.CpwVcIdMappingVcID})
    cpwVcIdMappingEntry.EntityData.Leafs.Append("cpwVcIdMappingPeerAddrType", types.YLeaf{"CpwVcIdMappingPeerAddrType", cpwVcIdMappingEntry.CpwVcIdMappingPeerAddrType})
    cpwVcIdMappingEntry.EntityData.Leafs.Append("cpwVcIdMappingPeerAddr", types.YLeaf{"CpwVcIdMappingPeerAddr", cpwVcIdMappingEntry.CpwVcIdMappingPeerAddr})
    cpwVcIdMappingEntry.EntityData.Leafs.Append("cpwVcIdMappingVcIndex", types.YLeaf{"CpwVcIdMappingVcIndex", cpwVcIdMappingEntry.CpwVcIdMappingVcIndex})

    cpwVcIdMappingEntry.EntityData.YListKeys = []string {"CpwVcIdMappingVcType", "CpwVcIdMappingVcID", "CpwVcIdMappingPeerAddrType", "CpwVcIdMappingPeerAddr", "CpwVcIdMappingVcIndex"}

    return &(cpwVcIdMappingEntry.EntityData)
}

// CISCOIETFPWMIB_CpwVcPeerMappingTable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_CpwVcPeerMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured in
    // cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_CpwVcPeerMappingTable_CpwVcPeerMappingEntry.
    CpwVcPeerMappingEntry []*CISCOIETFPWMIB_CpwVcPeerMappingTable_CpwVcPeerMappingEntry
}

func (cpwVcPeerMappingTable *CISCOIETFPWMIB_CpwVcPeerMappingTable) GetEntityData() *types.CommonEntityData {
    cpwVcPeerMappingTable.EntityData.YFilter = cpwVcPeerMappingTable.YFilter
    cpwVcPeerMappingTable.EntityData.YangName = "cpwVcPeerMappingTable"
    cpwVcPeerMappingTable.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPeerMappingTable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwVcPeerMappingTable.EntityData.SegmentPath = "cpwVcPeerMappingTable"
    cpwVcPeerMappingTable.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/" + cpwVcPeerMappingTable.EntityData.SegmentPath
    cpwVcPeerMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPeerMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPeerMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPeerMappingTable.EntityData.Children = types.NewOrderedMap()
    cpwVcPeerMappingTable.EntityData.Children.Append("cpwVcPeerMappingEntry", types.YChild{"CpwVcPeerMappingEntry", nil})
    for i := range cpwVcPeerMappingTable.CpwVcPeerMappingEntry {
        cpwVcPeerMappingTable.EntityData.Children.Append(types.GetSegmentPath(cpwVcPeerMappingTable.CpwVcPeerMappingEntry[i]), types.YChild{"CpwVcPeerMappingEntry", cpwVcPeerMappingTable.CpwVcPeerMappingEntry[i]})
    }
    cpwVcPeerMappingTable.EntityData.Leafs = types.NewOrderedMap()

    cpwVcPeerMappingTable.EntityData.YListKeys = []string {}

    return &(cpwVcPeerMappingTable.EntityData)
}

// CISCOIETFPWMIB_CpwVcPeerMappingTable_CpwVcPeerMappingEntry
// An entry in this table is created by the agent for every  
// VC configured in cpwVcTable.
type CISCOIETFPWMIB_CpwVcPeerMappingTable_CpwVcPeerMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address type of the peer node. The type is
    // InetAddressType.
    CpwVcPeerMappingPeerAddrType interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // string with length: 0..255.
    CpwVcPeerMappingPeerAddr interface{}

    // This attribute is a key. The VC type (indicate the service) of this VC. The
    // type is CpwVcType.
    CpwVcPeerMappingVcType interface{}

    // This attribute is a key. The VC ID of this VC. Zero if the VC is configured
    // manually. The type is interface{} with range: 0..4294967295.
    CpwVcPeerMappingVcID interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    CpwVcPeerMappingVcIndex interface{}
}

func (cpwVcPeerMappingEntry *CISCOIETFPWMIB_CpwVcPeerMappingTable_CpwVcPeerMappingEntry) GetEntityData() *types.CommonEntityData {
    cpwVcPeerMappingEntry.EntityData.YFilter = cpwVcPeerMappingEntry.YFilter
    cpwVcPeerMappingEntry.EntityData.YangName = "cpwVcPeerMappingEntry"
    cpwVcPeerMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    cpwVcPeerMappingEntry.EntityData.ParentYangName = "cpwVcPeerMappingTable"
    cpwVcPeerMappingEntry.EntityData.SegmentPath = "cpwVcPeerMappingEntry" + types.AddKeyToken(cpwVcPeerMappingEntry.CpwVcPeerMappingPeerAddrType, "cpwVcPeerMappingPeerAddrType") + types.AddKeyToken(cpwVcPeerMappingEntry.CpwVcPeerMappingPeerAddr, "cpwVcPeerMappingPeerAddr") + types.AddKeyToken(cpwVcPeerMappingEntry.CpwVcPeerMappingVcType, "cpwVcPeerMappingVcType") + types.AddKeyToken(cpwVcPeerMappingEntry.CpwVcPeerMappingVcID, "cpwVcPeerMappingVcID") + types.AddKeyToken(cpwVcPeerMappingEntry.CpwVcPeerMappingVcIndex, "cpwVcPeerMappingVcIndex")
    cpwVcPeerMappingEntry.EntityData.AbsolutePath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB/cpwVcPeerMappingTable/" + cpwVcPeerMappingEntry.EntityData.SegmentPath
    cpwVcPeerMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwVcPeerMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwVcPeerMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwVcPeerMappingEntry.EntityData.Children = types.NewOrderedMap()
    cpwVcPeerMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    cpwVcPeerMappingEntry.EntityData.Leafs.Append("cpwVcPeerMappingPeerAddrType", types.YLeaf{"CpwVcPeerMappingPeerAddrType", cpwVcPeerMappingEntry.CpwVcPeerMappingPeerAddrType})
    cpwVcPeerMappingEntry.EntityData.Leafs.Append("cpwVcPeerMappingPeerAddr", types.YLeaf{"CpwVcPeerMappingPeerAddr", cpwVcPeerMappingEntry.CpwVcPeerMappingPeerAddr})
    cpwVcPeerMappingEntry.EntityData.Leafs.Append("cpwVcPeerMappingVcType", types.YLeaf{"CpwVcPeerMappingVcType", cpwVcPeerMappingEntry.CpwVcPeerMappingVcType})
    cpwVcPeerMappingEntry.EntityData.Leafs.Append("cpwVcPeerMappingVcID", types.YLeaf{"CpwVcPeerMappingVcID", cpwVcPeerMappingEntry.CpwVcPeerMappingVcID})
    cpwVcPeerMappingEntry.EntityData.Leafs.Append("cpwVcPeerMappingVcIndex", types.YLeaf{"CpwVcPeerMappingVcIndex", cpwVcPeerMappingEntry.CpwVcPeerMappingVcIndex})

    cpwVcPeerMappingEntry.EntityData.YListKeys = []string {"CpwVcPeerMappingPeerAddrType", "CpwVcPeerMappingPeerAddr", "CpwVcPeerMappingVcType", "CpwVcPeerMappingVcID", "CpwVcPeerMappingVcIndex"}

    return &(cpwVcPeerMappingEntry.EntityData)
}

