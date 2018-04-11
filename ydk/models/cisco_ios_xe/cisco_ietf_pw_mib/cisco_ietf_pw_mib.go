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

    
    Cpwvcobjects CISCOIETFPWMIB_Cpwvcobjects

    // This table specifies information for connecting various  emulated services
    // to various tunnel type.
    Cpwvctable CISCOIETFPWMIB_Cpwvctable

    // This table provides per-VC performance information for the   current
    // interval.
    Cpwvcperfcurrenttable CISCOIETFPWMIB_Cpwvcperfcurrenttable

    // This table provides per-VC performance information for   each interval.
    Cpwvcperfintervaltable CISCOIETFPWMIB_Cpwvcperfintervaltable

    // This table provides per-VC Performance information from VC   start time.
    Cpwvcperftotaltable CISCOIETFPWMIB_Cpwvcperftotaltable

    // This table provides reverse mapping of the existing VCs   based on vc type
    // and VC ID ordering. This table is   typically useful for EMS ordered query
    // of existing VCs.
    Cpwvcidmappingtable CISCOIETFPWMIB_Cpwvcidmappingtable

    // This table provides reverse mapping of the existing VCs   based on vc type
    // and VC ID ordering. This table is   typically useful for EMS ordered query
    // of existing VCs.
    Cpwvcpeermappingtable CISCOIETFPWMIB_Cpwvcpeermappingtable
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFPWMIB.EntityData.YFilter = cISCOIETFPWMIB.YFilter
    cISCOIETFPWMIB.EntityData.YangName = "CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFPWMIB.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.SegmentPath = "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB"
    cISCOIETFPWMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFPWMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFPWMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFPWMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFPWMIB.EntityData.Children["cpwVcObjects"] = types.YChild{"Cpwvcobjects", &cISCOIETFPWMIB.Cpwvcobjects}
    cISCOIETFPWMIB.EntityData.Children["cpwVcTable"] = types.YChild{"Cpwvctable", &cISCOIETFPWMIB.Cpwvctable}
    cISCOIETFPWMIB.EntityData.Children["cpwVcPerfCurrentTable"] = types.YChild{"Cpwvcperfcurrenttable", &cISCOIETFPWMIB.Cpwvcperfcurrenttable}
    cISCOIETFPWMIB.EntityData.Children["cpwVcPerfIntervalTable"] = types.YChild{"Cpwvcperfintervaltable", &cISCOIETFPWMIB.Cpwvcperfintervaltable}
    cISCOIETFPWMIB.EntityData.Children["cpwVcPerfTotalTable"] = types.YChild{"Cpwvcperftotaltable", &cISCOIETFPWMIB.Cpwvcperftotaltable}
    cISCOIETFPWMIB.EntityData.Children["cpwVcIdMappingTable"] = types.YChild{"Cpwvcidmappingtable", &cISCOIETFPWMIB.Cpwvcidmappingtable}
    cISCOIETFPWMIB.EntityData.Children["cpwVcPeerMappingTable"] = types.YChild{"Cpwvcpeermappingtable", &cISCOIETFPWMIB.Cpwvcpeermappingtable}
    cISCOIETFPWMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFPWMIB.EntityData)
}

// CISCOIETFPWMIB_Cpwvcobjects
type CISCOIETFPWMIB_Cpwvcobjects struct {
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
    Cpwvcindexnext interface{}

    // Counter for number of error at VC level processing, for   example packets
    // received with unknown VC label. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwvcperftotalerrorpackets interface{}

    // If this object is set to true(1), then it enables the emission of cpwVcUp
    // and cpwVcDown notifications; otherwise these notifications are not emitted.
    // The type is bool.
    Cpwvcupdownnotifenable interface{}

    // This object defines the maximum number of PW VC notifications that can be
    // emitted from the device per second. The type is interface{} with range:
    // 0..4294967295.
    Cpwvcnotifrate interface{}
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetEntityData() *types.CommonEntityData {
    cpwvcobjects.EntityData.YFilter = cpwvcobjects.YFilter
    cpwvcobjects.EntityData.YangName = "cpwVcObjects"
    cpwvcobjects.EntityData.BundleName = "cisco_ios_xe"
    cpwvcobjects.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcobjects.EntityData.SegmentPath = "cpwVcObjects"
    cpwvcobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcobjects.EntityData.Children = make(map[string]types.YChild)
    cpwvcobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcobjects.EntityData.Leafs["cpwVcIndexNext"] = types.YLeaf{"Cpwvcindexnext", cpwvcobjects.Cpwvcindexnext}
    cpwvcobjects.EntityData.Leafs["cpwVcPerfTotalErrorPackets"] = types.YLeaf{"Cpwvcperftotalerrorpackets", cpwvcobjects.Cpwvcperftotalerrorpackets}
    cpwvcobjects.EntityData.Leafs["cpwVcUpDownNotifEnable"] = types.YLeaf{"Cpwvcupdownnotifenable", cpwvcobjects.Cpwvcupdownnotifenable}
    cpwvcobjects.EntityData.Leafs["cpwVcNotifRate"] = types.YLeaf{"Cpwvcnotifrate", cpwvcobjects.Cpwvcnotifrate}
    return &(cpwvcobjects.EntityData)
}

// CISCOIETFPWMIB_Cpwvctable
// This table specifies information for connecting various 
// emulated services to various tunnel type.
type CISCOIETFPWMIB_Cpwvctable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in this table represents an emulated virtual  connection (VC) across
    // a packet network. It is indexed by  cpwVcIndex, which uniquely identifying
    // a singular   connection. . The type is slice of
    // CISCOIETFPWMIB_Cpwvctable_Cpwvcentry.
    Cpwvcentry []CISCOIETFPWMIB_Cpwvctable_Cpwvcentry
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetEntityData() *types.CommonEntityData {
    cpwvctable.EntityData.YFilter = cpwvctable.YFilter
    cpwvctable.EntityData.YangName = "cpwVcTable"
    cpwvctable.EntityData.BundleName = "cisco_ios_xe"
    cpwvctable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvctable.EntityData.SegmentPath = "cpwVcTable"
    cpwvctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvctable.EntityData.Children = make(map[string]types.YChild)
    cpwvctable.EntityData.Children["cpwVcEntry"] = types.YChild{"Cpwvcentry", nil}
    for i := range cpwvctable.Cpwvcentry {
        cpwvctable.EntityData.Children[types.GetSegmentPath(&cpwvctable.Cpwvcentry[i])] = types.YChild{"Cpwvcentry", &cpwvctable.Cpwvcentry[i]}
    }
    cpwvctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvctable.EntityData)
}

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry
// A row in this table represents an emulated virtual 
// connection (VC) across a packet network. It is indexed by 
// cpwVcIndex, which uniquely identifying a singular  
// connection. 
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index for the conceptual row identifying a VC
    // within   this PW Emulation VC table. The type is interface{} with range:
    // 0..4294967295.
    Cpwvcindex interface{}

    // This value indicate the service to be carried over  this VC.   Note: the
    // exact set of VC types is yet to be worked   out by the WG. . The type is
    // CpwVcType.
    Cpwvctype interface{}

    // Set by the operator to indicate the protocol responsible   for establishing
    // this VC. Value 'manual' is used in all  cases where no maintenance protocol
    // (PW signaling) is used   to set-up the VC, i.e. require configuration of
    // entries in  the VC tables including VC labels, etc. The value  
    // 'maintenanceProtocol' is used in case of standard  signaling of the VC for
    // the specific PSN, for example LDP  for MPLS PSN as specified in <draft-
    // draft-martini-  l2circuit-trans-mpls> or L2TP control protocol.   Value
    // 'other' is used for other types of signaling. The type is Cpwvcowner.
    Cpwvcowner interface{}

    // Set by the operator to indicate the PSN type on which this   VC will be
    // carried. Based on this object, the relevant PSN   table entries are created
    // in the in the PSN specific MIB   modules. For example, if mpls(1) is
    // defined, the agent   create an entry in cpwVcMplsTable, which further
    // define the   MPLS PSN configuration.  Note: the exact set of PSN types is
    // yet to be worked   out by the WG. . The type is Cpwvcpsntype.
    Cpwvcpsntype interface{}

    // This object define the relative set-up priority of the VC    in a
    // lowest-to-highest fashion, where 0 is the highest   priority. VCs with the
    // same priority are treated with  equal priority. Dropped VC will be set
    // 'dormant' (as  indicated in cpwVcOperStatus).  This value is significant if
    // there are competing resources  between VCs and the implementation support
    // this feature.  If not supported or not relevant, the value of zero MUST  be
    // used. The type is interface{} with range: 0..7.
    Cpwvcsetuppriority interface{}

    // This object define the relative holding priority of the VC    in a
    // lowest-to-highest fashion, where 0 is the highest   priority. VCs with the
    // same priority are treated with  equal priority. Dropped VC will be set
    // 'dormant' (as  indicated in cpwVcOperStatus).  This value is significant if
    // there are competing resources  between VCs and the implementation support
    // this feature.  If not supported or not relevant, the value of zero MUST  be
    // used. The type is interface{} with range: 0..7.
    Cpwvcholdingpriority interface{}

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
    // Cpwvcinboundmode.
    Cpwvcinboundmode interface{}

    // Denotes the address type of the peer node maintenance  protocol (signaling)
    // address if PW maintenance protocol is  used for the VC creation. It should
    // be set to   'unknown' if PE/PW maintenance protocol is not used,   i.e.
    // cpwVcOwner is set to 'manual'. . The type is InetAddressType.
    Cpwvcpeeraddrtype interface{}

    // This object contains the value of of the peer node address  of the PW/PE
    // maintenance protocol entity. This object   should contain a value of 0 if
    // not relevant (manual   configuration of the VC). The type is string with
    // length: 0..255.
    Cpwvcpeeraddr interface{}

    // Used in the outgoing VC ID field within the 'Virtual  Circuit FEC Element'
    // when LDP signaling is used or PW ID   AVP for L2TP. The type is interface{}
    // with range: 0..4294967295.
    Cpwvcid interface{}

    // Used in the Group ID field sent to the peer PWES   within the maintenance
    // protocol used for VC setup,   zero if not used. The type is interface{}
    // with range: 0..4294967295.
    Cpwvclocalgroupid interface{}

    // Define if the control word will be sent with each packet by   the local
    // node. The type is bool.
    Cpwvccontrolword interface{}

    // If not equal zero, the optional IfMtu object in the   maintenance protocol
    // will be sent with this value,   representing the locally supported MTU size
    // over the   interface (or the virtual interface) associated with the   VC.
    // The type is interface{} with range: 0..65535.
    Cpwvclocalifmtu interface{}

    // Each VC is associated to an interface (or a virtual   interface) in the
    // ifTable of the node as part of the  service configuration. This object
    // defines if the   maintenance protocol will send the interface's name as 
    // appears on the ifTable in the name object as part of the  maintenance
    // protocol. If set to false, the optional element  will not be sent. The type
    // is bool.
    Cpwvclocalifstring interface{}

    // Obtained from the Group ID field as received via the   maintenance protocol
    // used for VC setup, zero if not used.   Value of 0xFFFF shall be used if the
    // object is yet to be   defined by the VC maintenance protocol. The type is
    // interface{} with range: 0..4294967295.
    Cpwvcremotegroupid interface{}

    // If maintenance protocol is used for VC establishment, this   parameter
    // indicates the received status of the control word   usage, i.e. if packets
    // will be received with control word  or not. The value of 'notYetKnown' is
    // used while the  maintenance protocol has not yet received the indication  
    // from the remote node.  In manual configuration of the VC this parameters
    // indicate   to the local node what is the expected encapsulation for  the
    // received packets. . The type is Cpwvcremotecontrolword.
    Cpwvcremotecontrolword interface{}

    // The remote interface MTU as (optionally) received from the  remote node via
    // the maintenance protocol. Should be zero if  this parameter is not
    // available or not used. The type is interface{} with range: 0..4294967295.
    Cpwvcremoteifmtu interface{}

    // Indicate the interface description string as received by  the maintenance
    // protocol, MUST be NULL string if not   applicable or not known yet. The
    // type is string with length: 0..80.
    Cpwvcremoteifstring interface{}

    // The VC label used in the outbound direction (i.e. toward   the PSN). It may
    // be set up manually if owner is 'manual' or   automatically otherwise.
    // Examples: For MPLS PSN, it   represents the 20 bits of VC tag, for L2TP it
    // represent the  32 bits Session ID.  If the label is not yet known
    // (signaling in process), the   object should return a value of 0xFFFF. The
    // type is interface{} with range: 0..4294967295.
    Cpwvcoutboundvclabel interface{}

    // The VC label used in the inbound direction (i.e. packets   received from
    // the PSN. It may be set up manually if owner  is 'manual' or automatically
    // otherwise.   Examples: For MPLS PSN, it represents the 20 bits of VC  tag,
    // for L2TP it represent the 32 bits Session ID.  If the label is not yet
    // known (signaling in process), the   object should return a value of 0xFFFF.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcinboundvclabel interface{}

    // The canonical name assigned to the VC. The type is string.
    Cpwvcname interface{}

    // A textual string containing information about the VC.   If there is no
    // description this object contains a zero  length string. The type is string.
    Cpwvcdescr interface{}

    // System time when this VC was created. The type is interface{} with range:
    // 0..4294967295.
    Cpwvccreatetime interface{}

    // Number of consecutive ticks this VC has been 'up' in  both directions
    // together (i.e. 'up' is observed in   cpwVcOperStatus.). The type is
    // interface{} with range: 0..4294967295.
    Cpwvcuptime interface{}

    // The desired operational status of this VC. The type is Cpwvcadminstatus.
    Cpwvcadminstatus interface{}

    // Indicates the actual combined operational status of this   VC. It is 'up'
    // if both cpwVcInboundOperStatus and   cpwVcOutboundOperStatus are in 'up'
    // state. For all other   values, if the VCs in both directions are of the
    // same  value it reflects that value, otherwise it is set to the  most severe
    // status out of the two statuses. The order of   severance from most severe
    // to less severe is: unknown,   notPresent, down, lowerLayerDown, dormant,
    // testing, up.  The operator may consult the per direction OperStatus for 
    // fault isolation per direction. The type is CpwOperStatus.
    Cpwvcoperstatus interface{}

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
    Cpwvcinboundoperstatus interface{}

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
    Cpwvcoutboundoperstatus interface{}

    // The number of seconds, including partial seconds,  that have elapsed since
    // the beginning of the current  measurement period. If, for some reason, such
    // as an  adjustment in the system's time-of-day clock, the  current interval
    // exceeds the maximum value, the  agent will return the maximum value. The
    // type is interface{} with range: 1..900.
    Cpwvctimeelapsed interface{}

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
    Cpwvcvalidintervals interface{}

    // For creating, modifying, and deleting this row. The type is RowStatus.
    Cpwvcrowstatus interface{}

    // This variable indicates the storage type for this  object. The type is
    // StorageType.
    Cpwvcstoragetype interface{}
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetEntityData() *types.CommonEntityData {
    cpwvcentry.EntityData.YFilter = cpwvcentry.YFilter
    cpwvcentry.EntityData.YangName = "cpwVcEntry"
    cpwvcentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcentry.EntityData.ParentYangName = "cpwVcTable"
    cpwvcentry.EntityData.SegmentPath = "cpwVcEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcentry.Cpwvcindex) + "']"
    cpwvcentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcentry.Cpwvcindex}
    cpwvcentry.EntityData.Leafs["cpwVcType"] = types.YLeaf{"Cpwvctype", cpwvcentry.Cpwvctype}
    cpwvcentry.EntityData.Leafs["cpwVcOwner"] = types.YLeaf{"Cpwvcowner", cpwvcentry.Cpwvcowner}
    cpwvcentry.EntityData.Leafs["cpwVcPsnType"] = types.YLeaf{"Cpwvcpsntype", cpwvcentry.Cpwvcpsntype}
    cpwvcentry.EntityData.Leafs["cpwVcSetUpPriority"] = types.YLeaf{"Cpwvcsetuppriority", cpwvcentry.Cpwvcsetuppriority}
    cpwvcentry.EntityData.Leafs["cpwVcHoldingPriority"] = types.YLeaf{"Cpwvcholdingpriority", cpwvcentry.Cpwvcholdingpriority}
    cpwvcentry.EntityData.Leafs["cpwVcInboundMode"] = types.YLeaf{"Cpwvcinboundmode", cpwvcentry.Cpwvcinboundmode}
    cpwvcentry.EntityData.Leafs["cpwVcPeerAddrType"] = types.YLeaf{"Cpwvcpeeraddrtype", cpwvcentry.Cpwvcpeeraddrtype}
    cpwvcentry.EntityData.Leafs["cpwVcPeerAddr"] = types.YLeaf{"Cpwvcpeeraddr", cpwvcentry.Cpwvcpeeraddr}
    cpwvcentry.EntityData.Leafs["cpwVcID"] = types.YLeaf{"Cpwvcid", cpwvcentry.Cpwvcid}
    cpwvcentry.EntityData.Leafs["cpwVcLocalGroupID"] = types.YLeaf{"Cpwvclocalgroupid", cpwvcentry.Cpwvclocalgroupid}
    cpwvcentry.EntityData.Leafs["cpwVcControlWord"] = types.YLeaf{"Cpwvccontrolword", cpwvcentry.Cpwvccontrolword}
    cpwvcentry.EntityData.Leafs["cpwVcLocalIfMtu"] = types.YLeaf{"Cpwvclocalifmtu", cpwvcentry.Cpwvclocalifmtu}
    cpwvcentry.EntityData.Leafs["cpwVcLocalIfString"] = types.YLeaf{"Cpwvclocalifstring", cpwvcentry.Cpwvclocalifstring}
    cpwvcentry.EntityData.Leafs["cpwVcRemoteGroupID"] = types.YLeaf{"Cpwvcremotegroupid", cpwvcentry.Cpwvcremotegroupid}
    cpwvcentry.EntityData.Leafs["cpwVcRemoteControlWord"] = types.YLeaf{"Cpwvcremotecontrolword", cpwvcentry.Cpwvcremotecontrolword}
    cpwvcentry.EntityData.Leafs["cpwVcRemoteIfMtu"] = types.YLeaf{"Cpwvcremoteifmtu", cpwvcentry.Cpwvcremoteifmtu}
    cpwvcentry.EntityData.Leafs["cpwVcRemoteIfString"] = types.YLeaf{"Cpwvcremoteifstring", cpwvcentry.Cpwvcremoteifstring}
    cpwvcentry.EntityData.Leafs["cpwVcOutboundVcLabel"] = types.YLeaf{"Cpwvcoutboundvclabel", cpwvcentry.Cpwvcoutboundvclabel}
    cpwvcentry.EntityData.Leafs["cpwVcInboundVcLabel"] = types.YLeaf{"Cpwvcinboundvclabel", cpwvcentry.Cpwvcinboundvclabel}
    cpwvcentry.EntityData.Leafs["cpwVcName"] = types.YLeaf{"Cpwvcname", cpwvcentry.Cpwvcname}
    cpwvcentry.EntityData.Leafs["cpwVcDescr"] = types.YLeaf{"Cpwvcdescr", cpwvcentry.Cpwvcdescr}
    cpwvcentry.EntityData.Leafs["cpwVcCreateTime"] = types.YLeaf{"Cpwvccreatetime", cpwvcentry.Cpwvccreatetime}
    cpwvcentry.EntityData.Leafs["cpwVcUpTime"] = types.YLeaf{"Cpwvcuptime", cpwvcentry.Cpwvcuptime}
    cpwvcentry.EntityData.Leafs["cpwVcAdminStatus"] = types.YLeaf{"Cpwvcadminstatus", cpwvcentry.Cpwvcadminstatus}
    cpwvcentry.EntityData.Leafs["cpwVcOperStatus"] = types.YLeaf{"Cpwvcoperstatus", cpwvcentry.Cpwvcoperstatus}
    cpwvcentry.EntityData.Leafs["cpwVcInboundOperStatus"] = types.YLeaf{"Cpwvcinboundoperstatus", cpwvcentry.Cpwvcinboundoperstatus}
    cpwvcentry.EntityData.Leafs["cpwVcOutboundOperStatus"] = types.YLeaf{"Cpwvcoutboundoperstatus", cpwvcentry.Cpwvcoutboundoperstatus}
    cpwvcentry.EntityData.Leafs["cpwVcTimeElapsed"] = types.YLeaf{"Cpwvctimeelapsed", cpwvcentry.Cpwvctimeelapsed}
    cpwvcentry.EntityData.Leafs["cpwVcValidIntervals"] = types.YLeaf{"Cpwvcvalidintervals", cpwvcentry.Cpwvcvalidintervals}
    cpwvcentry.EntityData.Leafs["cpwVcRowStatus"] = types.YLeaf{"Cpwvcrowstatus", cpwvcentry.Cpwvcrowstatus}
    cpwvcentry.EntityData.Leafs["cpwVcStorageType"] = types.YLeaf{"Cpwvcstoragetype", cpwvcentry.Cpwvcstoragetype}
    return &(cpwvcentry.EntityData)
}

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus represents The desired operational status of this VC.
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus string

const (
    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus_up CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus = "up"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus_down CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus = "down"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus_testing CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcadminstatus = "testing"
)

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode represents regardless of the outer tunnel used to carry the VC.
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode string

const (
    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode_loose CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode = "loose"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode_strict CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcinboundmode = "strict"
)

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner represents Value 'other' is used for other types of signaling.
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner string

const (
    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner_manual CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner = "manual"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner_maintenanceProtocol CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner = "maintenanceProtocol"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner_other CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcowner = "other"
)

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype represents out by the WG. 
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype string

const (
    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_mpls CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "mpls"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_l2tp CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "l2tp"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_ip CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "ip"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_mplsOverIp CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "mplsOverIp"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_gre CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "gre"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype_other CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcpsntype = "other"
)

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword represents the received packets. 
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword string

const (
    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword_noControlWord CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword = "noControlWord"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword_withControlWord CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword = "withControlWord"

    CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword_notYetKnown CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcremotecontrolword = "notYetKnown"
)

// CISCOIETFPWMIB_Cpwvcperfcurrenttable
// This table provides per-VC performance information for the  
// current interval.
type CISCOIETFPWMIB_Cpwvcperfcurrenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for  every VC. The type is
    // slice of CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry.
    Cpwvcperfcurrententry []CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetEntityData() *types.CommonEntityData {
    cpwvcperfcurrenttable.EntityData.YFilter = cpwvcperfcurrenttable.YFilter
    cpwvcperfcurrenttable.EntityData.YangName = "cpwVcPerfCurrentTable"
    cpwvcperfcurrenttable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperfcurrenttable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcperfcurrenttable.EntityData.SegmentPath = "cpwVcPerfCurrentTable"
    cpwvcperfcurrenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperfcurrenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperfcurrenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperfcurrenttable.EntityData.Children = make(map[string]types.YChild)
    cpwvcperfcurrenttable.EntityData.Children["cpwVcPerfCurrentEntry"] = types.YChild{"Cpwvcperfcurrententry", nil}
    for i := range cpwvcperfcurrenttable.Cpwvcperfcurrententry {
        cpwvcperfcurrenttable.EntityData.Children[types.GetSegmentPath(&cpwvcperfcurrenttable.Cpwvcperfcurrententry[i])] = types.YChild{"Cpwvcperfcurrententry", &cpwvcperfcurrenttable.Cpwvcperfcurrententry[i]}
    }
    cpwvcperfcurrenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcperfcurrenttable.EntityData)
}

// CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry
// An entry in this table is created by the agent for 
// every VC.
type CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // High capacity counter for number of packets received  by the VC (from the
    // PSN) in the current 15 minute  interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfcurrentinhcpackets interface{}

    // High capacity counter for number of bytes received  by the VC (from the
    // PSN) in the current 15 minute  interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfcurrentinhcbytes interface{}

    // High capacity counter for number of packets forwarded  by the VC (to the
    // PSN) in the current 15 minute interval. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwvcperfcurrentouthcpackets interface{}

    // High capacity counter for number of bytes forwarded  by the VC (to the PSN)
    // in the current 15 minute interval. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpwvcperfcurrentouthcbytes interface{}
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetEntityData() *types.CommonEntityData {
    cpwvcperfcurrententry.EntityData.YFilter = cpwvcperfcurrententry.YFilter
    cpwvcperfcurrententry.EntityData.YangName = "cpwVcPerfCurrentEntry"
    cpwvcperfcurrententry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperfcurrententry.EntityData.ParentYangName = "cpwVcPerfCurrentTable"
    cpwvcperfcurrententry.EntityData.SegmentPath = "cpwVcPerfCurrentEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperfcurrententry.Cpwvcindex) + "']"
    cpwvcperfcurrententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperfcurrententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperfcurrententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperfcurrententry.EntityData.Children = make(map[string]types.YChild)
    cpwvcperfcurrententry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcperfcurrententry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcperfcurrententry.Cpwvcindex}
    cpwvcperfcurrententry.EntityData.Leafs["cpwVcPerfCurrentInHCPackets"] = types.YLeaf{"Cpwvcperfcurrentinhcpackets", cpwvcperfcurrententry.Cpwvcperfcurrentinhcpackets}
    cpwvcperfcurrententry.EntityData.Leafs["cpwVcPerfCurrentInHCBytes"] = types.YLeaf{"Cpwvcperfcurrentinhcbytes", cpwvcperfcurrententry.Cpwvcperfcurrentinhcbytes}
    cpwvcperfcurrententry.EntityData.Leafs["cpwVcPerfCurrentOutHCPackets"] = types.YLeaf{"Cpwvcperfcurrentouthcpackets", cpwvcperfcurrententry.Cpwvcperfcurrentouthcpackets}
    cpwvcperfcurrententry.EntityData.Leafs["cpwVcPerfCurrentOutHCBytes"] = types.YLeaf{"Cpwvcperfcurrentouthcbytes", cpwvcperfcurrententry.Cpwvcperfcurrentouthcbytes}
    return &(cpwvcperfcurrententry.EntityData)
}

// CISCOIETFPWMIB_Cpwvcperfintervaltable
// This table provides per-VC performance information for  
// each interval.
type CISCOIETFPWMIB_Cpwvcperfintervaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry.
    Cpwvcperfintervalentry []CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetEntityData() *types.CommonEntityData {
    cpwvcperfintervaltable.EntityData.YFilter = cpwvcperfintervaltable.YFilter
    cpwvcperfintervaltable.EntityData.YangName = "cpwVcPerfIntervalTable"
    cpwvcperfintervaltable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperfintervaltable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcperfintervaltable.EntityData.SegmentPath = "cpwVcPerfIntervalTable"
    cpwvcperfintervaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperfintervaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperfintervaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperfintervaltable.EntityData.Children = make(map[string]types.YChild)
    cpwvcperfintervaltable.EntityData.Children["cpwVcPerfIntervalEntry"] = types.YChild{"Cpwvcperfintervalentry", nil}
    for i := range cpwvcperfintervaltable.Cpwvcperfintervalentry {
        cpwvcperfintervaltable.EntityData.Children[types.GetSegmentPath(&cpwvcperfintervaltable.Cpwvcperfintervalentry[i])] = types.YChild{"Cpwvcperfintervalentry", &cpwvcperfintervaltable.Cpwvcperfintervalentry[i]}
    }
    cpwvcperfintervaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcperfintervaltable.EntityData)
}

// CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // This attribute is a key. A number N, between 1 and 96, which identifies the
    // interval for which the set of statistics is available.  The interval
    // identified by 1 is the most recently  completed 15 minute interval, and the
    // interval identified  by N is the interval immediately preceding the one 
    // identified by N-1.  The minimum range of N is 1 through 4. The default
    // range  is 1 to 32. The maximum range of N is 1 through 96. . The type is
    // interface{} with range: 1..96.
    Cpwvcperfintervalnumber interface{}

    // This variable indicates if the data for this interval  is valid. The type
    // is bool.
    Cpwvcperfintervalvaliddata interface{}

    // The duration of a particular interval in seconds.  Adjustments in the
    // system's time-of-day clock, may  cause the interval to be greater or less
    // than the  normal value. Therefore this actual interval value  is provided.
    // The type is interface{} with range: -2147483648..2147483647.
    Cpwvcperfintervaltimeelapsed interface{}

    // High capacity counter for number of packets received by  the VC (from the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfintervalinhcpackets interface{}

    // High capacity counter for number of bytes received by the   VC (from the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfintervalinhcbytes interface{}

    // High capacity counter for number of packets forwarded by   the VC (to the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfintervalouthcpackets interface{}

    // High capacity counter for number of bytes forwarded by the   VC (to the
    // PSN) in a particular 15-minute interval. The type is interface{} with
    // range: 0..18446744073709551615.
    Cpwvcperfintervalouthcbytes interface{}
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetEntityData() *types.CommonEntityData {
    cpwvcperfintervalentry.EntityData.YFilter = cpwvcperfintervalentry.YFilter
    cpwvcperfintervalentry.EntityData.YangName = "cpwVcPerfIntervalEntry"
    cpwvcperfintervalentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperfintervalentry.EntityData.ParentYangName = "cpwVcPerfIntervalTable"
    cpwvcperfintervalentry.EntityData.SegmentPath = "cpwVcPerfIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperfintervalentry.Cpwvcindex) + "']" + "[cpwVcPerfIntervalNumber='" + fmt.Sprintf("%v", cpwvcperfintervalentry.Cpwvcperfintervalnumber) + "']"
    cpwvcperfintervalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperfintervalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperfintervalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperfintervalentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcperfintervalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcperfintervalentry.Cpwvcindex}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalNumber"] = types.YLeaf{"Cpwvcperfintervalnumber", cpwvcperfintervalentry.Cpwvcperfintervalnumber}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalValidData"] = types.YLeaf{"Cpwvcperfintervalvaliddata", cpwvcperfintervalentry.Cpwvcperfintervalvaliddata}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalTimeElapsed"] = types.YLeaf{"Cpwvcperfintervaltimeelapsed", cpwvcperfintervalentry.Cpwvcperfintervaltimeelapsed}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalInHCPackets"] = types.YLeaf{"Cpwvcperfintervalinhcpackets", cpwvcperfintervalentry.Cpwvcperfintervalinhcpackets}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalInHCBytes"] = types.YLeaf{"Cpwvcperfintervalinhcbytes", cpwvcperfintervalentry.Cpwvcperfintervalinhcbytes}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalOutHCPackets"] = types.YLeaf{"Cpwvcperfintervalouthcpackets", cpwvcperfintervalentry.Cpwvcperfintervalouthcpackets}
    cpwvcperfintervalentry.EntityData.Leafs["cpwVcPerfIntervalOutHCBytes"] = types.YLeaf{"Cpwvcperfintervalouthcbytes", cpwvcperfintervalentry.Cpwvcperfintervalouthcbytes}
    return &(cpwvcperfintervalentry.EntityData)
}

// CISCOIETFPWMIB_Cpwvcperftotaltable
// This table provides per-VC Performance information from VC  
// start time.
type CISCOIETFPWMIB_Cpwvcperftotaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry.
    Cpwvcperftotalentry []CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetEntityData() *types.CommonEntityData {
    cpwvcperftotaltable.EntityData.YFilter = cpwvcperftotaltable.YFilter
    cpwvcperftotaltable.EntityData.YangName = "cpwVcPerfTotalTable"
    cpwvcperftotaltable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperftotaltable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcperftotaltable.EntityData.SegmentPath = "cpwVcPerfTotalTable"
    cpwvcperftotaltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperftotaltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperftotaltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperftotaltable.EntityData.Children = make(map[string]types.YChild)
    cpwvcperftotaltable.EntityData.Children["cpwVcPerfTotalEntry"] = types.YChild{"Cpwvcperftotalentry", nil}
    for i := range cpwvcperftotaltable.Cpwvcperftotalentry {
        cpwvcperftotaltable.EntityData.Children[types.GetSegmentPath(&cpwvcperftotaltable.Cpwvcperftotalentry[i])] = types.YChild{"Cpwvcperftotalentry", &cpwvcperftotaltable.Cpwvcperftotalentry[i]}
    }
    cpwvcperftotaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcperftotaltable.EntityData)
}

// CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to cisco_ietf_pw_mib.CISCOIETFPWMIB_Cpwvctable_Cpwvcentry_Cpwvcindex
    Cpwvcindex interface{}

    // High capacity counter for number of packets received by the   VC (from the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    Cpwvcperftotalinhcpackets interface{}

    // High capacity counter for number of bytes received by the   VC (from the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    Cpwvcperftotalinhcbytes interface{}

    // High capacity counter for number of packets forwarded by   the VC (to the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    Cpwvcperftotalouthcpackets interface{}

    // High capacity counter for number of bytes forwarded by the   VC (to the
    // PSN). The type is interface{} with range: 0..18446744073709551615.
    Cpwvcperftotalouthcbytes interface{}

    // The value of sysUpTime on the most recent occasion at  which any one or
    // more of this row Counter32 or  Counter64 suffered a discontinuity. If no
    // such  discontinuities have occurred since the last re-  initialization of
    // the local management subsystem, then  this object contains a zero value.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcperftotaldiscontinuitytime interface{}
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetEntityData() *types.CommonEntityData {
    cpwvcperftotalentry.EntityData.YFilter = cpwvcperftotalentry.YFilter
    cpwvcperftotalentry.EntityData.YangName = "cpwVcPerfTotalEntry"
    cpwvcperftotalentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcperftotalentry.EntityData.ParentYangName = "cpwVcPerfTotalTable"
    cpwvcperftotalentry.EntityData.SegmentPath = "cpwVcPerfTotalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperftotalentry.Cpwvcindex) + "']"
    cpwvcperftotalentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcperftotalentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcperftotalentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcperftotalentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcperftotalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcperftotalentry.EntityData.Leafs["cpwVcIndex"] = types.YLeaf{"Cpwvcindex", cpwvcperftotalentry.Cpwvcindex}
    cpwvcperftotalentry.EntityData.Leafs["cpwVcPerfTotalInHCPackets"] = types.YLeaf{"Cpwvcperftotalinhcpackets", cpwvcperftotalentry.Cpwvcperftotalinhcpackets}
    cpwvcperftotalentry.EntityData.Leafs["cpwVcPerfTotalInHCBytes"] = types.YLeaf{"Cpwvcperftotalinhcbytes", cpwvcperftotalentry.Cpwvcperftotalinhcbytes}
    cpwvcperftotalentry.EntityData.Leafs["cpwVcPerfTotalOutHCPackets"] = types.YLeaf{"Cpwvcperftotalouthcpackets", cpwvcperftotalentry.Cpwvcperftotalouthcpackets}
    cpwvcperftotalentry.EntityData.Leafs["cpwVcPerfTotalOutHCBytes"] = types.YLeaf{"Cpwvcperftotalouthcbytes", cpwvcperftotalentry.Cpwvcperftotalouthcbytes}
    cpwvcperftotalentry.EntityData.Leafs["cpwVcPerfTotalDiscontinuityTime"] = types.YLeaf{"Cpwvcperftotaldiscontinuitytime", cpwvcperftotalentry.Cpwvcperftotaldiscontinuitytime}
    return &(cpwvcperftotalentry.EntityData)
}

// CISCOIETFPWMIB_Cpwvcidmappingtable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_Cpwvcidmappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured by
    // the cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry.
    Cpwvcidmappingentry []CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetEntityData() *types.CommonEntityData {
    cpwvcidmappingtable.EntityData.YFilter = cpwvcidmappingtable.YFilter
    cpwvcidmappingtable.EntityData.YangName = "cpwVcIdMappingTable"
    cpwvcidmappingtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcidmappingtable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcidmappingtable.EntityData.SegmentPath = "cpwVcIdMappingTable"
    cpwvcidmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcidmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcidmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcidmappingtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcidmappingtable.EntityData.Children["cpwVcIdMappingEntry"] = types.YChild{"Cpwvcidmappingentry", nil}
    for i := range cpwvcidmappingtable.Cpwvcidmappingentry {
        cpwvcidmappingtable.EntityData.Children[types.GetSegmentPath(&cpwvcidmappingtable.Cpwvcidmappingentry[i])] = types.YChild{"Cpwvcidmappingentry", &cpwvcidmappingtable.Cpwvcidmappingentry[i]}
    }
    cpwvcidmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcidmappingtable.EntityData)
}

// CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry
// An entry in this table is created by the agent for every  
// VC configured by the cpwVcTable.
type CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The VC type (indicate the service) of this VC. The
    // type is CpwVcType.
    Cpwvcidmappingvctype interface{}

    // This attribute is a key. The VC ID of this VC. Zero if the VC is configured
    // manually. The type is interface{} with range: 0..4294967295.
    Cpwvcidmappingvcid interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // InetAddressType.
    Cpwvcidmappingpeeraddrtype interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // string with length: 0..255.
    Cpwvcidmappingpeeraddr interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcidmappingvcindex interface{}
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetEntityData() *types.CommonEntityData {
    cpwvcidmappingentry.EntityData.YFilter = cpwvcidmappingentry.YFilter
    cpwvcidmappingentry.EntityData.YangName = "cpwVcIdMappingEntry"
    cpwvcidmappingentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcidmappingentry.EntityData.ParentYangName = "cpwVcIdMappingTable"
    cpwvcidmappingentry.EntityData.SegmentPath = "cpwVcIdMappingEntry" + "[cpwVcIdMappingVcType='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvctype) + "']" + "[cpwVcIdMappingVcID='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvcid) + "']" + "[cpwVcIdMappingPeerAddrType='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingpeeraddrtype) + "']" + "[cpwVcIdMappingPeerAddr='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingpeeraddr) + "']" + "[cpwVcIdMappingVcIndex='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvcindex) + "']"
    cpwvcidmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcidmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcidmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcidmappingentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcidmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcidmappingentry.EntityData.Leafs["cpwVcIdMappingVcType"] = types.YLeaf{"Cpwvcidmappingvctype", cpwvcidmappingentry.Cpwvcidmappingvctype}
    cpwvcidmappingentry.EntityData.Leafs["cpwVcIdMappingVcID"] = types.YLeaf{"Cpwvcidmappingvcid", cpwvcidmappingentry.Cpwvcidmappingvcid}
    cpwvcidmappingentry.EntityData.Leafs["cpwVcIdMappingPeerAddrType"] = types.YLeaf{"Cpwvcidmappingpeeraddrtype", cpwvcidmappingentry.Cpwvcidmappingpeeraddrtype}
    cpwvcidmappingentry.EntityData.Leafs["cpwVcIdMappingPeerAddr"] = types.YLeaf{"Cpwvcidmappingpeeraddr", cpwvcidmappingentry.Cpwvcidmappingpeeraddr}
    cpwvcidmappingentry.EntityData.Leafs["cpwVcIdMappingVcIndex"] = types.YLeaf{"Cpwvcidmappingvcindex", cpwvcidmappingentry.Cpwvcidmappingvcindex}
    return &(cpwvcidmappingentry.EntityData)
}

// CISCOIETFPWMIB_Cpwvcpeermappingtable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_Cpwvcpeermappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured in
    // cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry.
    Cpwvcpeermappingentry []CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetEntityData() *types.CommonEntityData {
    cpwvcpeermappingtable.EntityData.YFilter = cpwvcpeermappingtable.YFilter
    cpwvcpeermappingtable.EntityData.YangName = "cpwVcPeerMappingTable"
    cpwvcpeermappingtable.EntityData.BundleName = "cisco_ios_xe"
    cpwvcpeermappingtable.EntityData.ParentYangName = "CISCO-IETF-PW-MIB"
    cpwvcpeermappingtable.EntityData.SegmentPath = "cpwVcPeerMappingTable"
    cpwvcpeermappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcpeermappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcpeermappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcpeermappingtable.EntityData.Children = make(map[string]types.YChild)
    cpwvcpeermappingtable.EntityData.Children["cpwVcPeerMappingEntry"] = types.YChild{"Cpwvcpeermappingentry", nil}
    for i := range cpwvcpeermappingtable.Cpwvcpeermappingentry {
        cpwvcpeermappingtable.EntityData.Children[types.GetSegmentPath(&cpwvcpeermappingtable.Cpwvcpeermappingentry[i])] = types.YChild{"Cpwvcpeermappingentry", &cpwvcpeermappingtable.Cpwvcpeermappingentry[i]}
    }
    cpwvcpeermappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpwvcpeermappingtable.EntityData)
}

// CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry
// An entry in this table is created by the agent for every  
// VC configured in cpwVcTable.
type CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address type of the peer node. The type is
    // InetAddressType.
    Cpwvcpeermappingpeeraddrtype interface{}

    // This attribute is a key. IP address type of the peer node. The type is
    // string with length: 0..255.
    Cpwvcpeermappingpeeraddr interface{}

    // This attribute is a key. The VC type (indicate the service) of this VC. The
    // type is CpwVcType.
    Cpwvcpeermappingvctype interface{}

    // This attribute is a key. The VC ID of this VC. Zero if the VC is configured
    // manually. The type is interface{} with range: 0..4294967295.
    Cpwvcpeermappingvcid interface{}

    // This attribute is a key. The value that represent the VC in the cpwVcTable.
    // The type is interface{} with range: 0..4294967295.
    Cpwvcpeermappingvcindex interface{}
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetEntityData() *types.CommonEntityData {
    cpwvcpeermappingentry.EntityData.YFilter = cpwvcpeermappingentry.YFilter
    cpwvcpeermappingentry.EntityData.YangName = "cpwVcPeerMappingEntry"
    cpwvcpeermappingentry.EntityData.BundleName = "cisco_ios_xe"
    cpwvcpeermappingentry.EntityData.ParentYangName = "cpwVcPeerMappingTable"
    cpwvcpeermappingentry.EntityData.SegmentPath = "cpwVcPeerMappingEntry" + "[cpwVcPeerMappingPeerAddrType='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddrtype) + "']" + "[cpwVcPeerMappingPeerAddr='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddr) + "']" + "[cpwVcPeerMappingVcType='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvctype) + "']" + "[cpwVcPeerMappingVcID='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvcid) + "']" + "[cpwVcPeerMappingVcIndex='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvcindex) + "']"
    cpwvcpeermappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpwvcpeermappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpwvcpeermappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpwvcpeermappingentry.EntityData.Children = make(map[string]types.YChild)
    cpwvcpeermappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpwvcpeermappingentry.EntityData.Leafs["cpwVcPeerMappingPeerAddrType"] = types.YLeaf{"Cpwvcpeermappingpeeraddrtype", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddrtype}
    cpwvcpeermappingentry.EntityData.Leafs["cpwVcPeerMappingPeerAddr"] = types.YLeaf{"Cpwvcpeermappingpeeraddr", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddr}
    cpwvcpeermappingentry.EntityData.Leafs["cpwVcPeerMappingVcType"] = types.YLeaf{"Cpwvcpeermappingvctype", cpwvcpeermappingentry.Cpwvcpeermappingvctype}
    cpwvcpeermappingentry.EntityData.Leafs["cpwVcPeerMappingVcID"] = types.YLeaf{"Cpwvcpeermappingvcid", cpwvcpeermappingentry.Cpwvcpeermappingvcid}
    cpwvcpeermappingentry.EntityData.Leafs["cpwVcPeerMappingVcIndex"] = types.YLeaf{"Cpwvcpeermappingvcindex", cpwvcpeermappingentry.Cpwvcpeermappingvcindex}
    return &(cpwvcpeermappingentry.EntityData)
}

