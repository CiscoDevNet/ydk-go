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
    parent types.Entity
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

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetFilter() yfilter.YFilter { return cISCOIETFPWMIB.YFilter }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFPWMIB.YFilter = yf }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetGoName(yname string) string {
    if yname == "cpwVcObjects" { return "Cpwvcobjects" }
    if yname == "cpwVcTable" { return "Cpwvctable" }
    if yname == "cpwVcPerfCurrentTable" { return "Cpwvcperfcurrenttable" }
    if yname == "cpwVcPerfIntervalTable" { return "Cpwvcperfintervaltable" }
    if yname == "cpwVcPerfTotalTable" { return "Cpwvcperftotaltable" }
    if yname == "cpwVcIdMappingTable" { return "Cpwvcidmappingtable" }
    if yname == "cpwVcPeerMappingTable" { return "Cpwvcpeermappingtable" }
    return ""
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetSegmentPath() string {
    return "CISCO-IETF-PW-MIB:CISCO-IETF-PW-MIB"
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcObjects" {
        return &cISCOIETFPWMIB.Cpwvcobjects
    }
    if childYangName == "cpwVcTable" {
        return &cISCOIETFPWMIB.Cpwvctable
    }
    if childYangName == "cpwVcPerfCurrentTable" {
        return &cISCOIETFPWMIB.Cpwvcperfcurrenttable
    }
    if childYangName == "cpwVcPerfIntervalTable" {
        return &cISCOIETFPWMIB.Cpwvcperfintervaltable
    }
    if childYangName == "cpwVcPerfTotalTable" {
        return &cISCOIETFPWMIB.Cpwvcperftotaltable
    }
    if childYangName == "cpwVcIdMappingTable" {
        return &cISCOIETFPWMIB.Cpwvcidmappingtable
    }
    if childYangName == "cpwVcPeerMappingTable" {
        return &cISCOIETFPWMIB.Cpwvcpeermappingtable
    }
    return nil
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpwVcObjects"] = &cISCOIETFPWMIB.Cpwvcobjects
    children["cpwVcTable"] = &cISCOIETFPWMIB.Cpwvctable
    children["cpwVcPerfCurrentTable"] = &cISCOIETFPWMIB.Cpwvcperfcurrenttable
    children["cpwVcPerfIntervalTable"] = &cISCOIETFPWMIB.Cpwvcperfintervaltable
    children["cpwVcPerfTotalTable"] = &cISCOIETFPWMIB.Cpwvcperftotaltable
    children["cpwVcIdMappingTable"] = &cISCOIETFPWMIB.Cpwvcidmappingtable
    children["cpwVcPeerMappingTable"] = &cISCOIETFPWMIB.Cpwvcpeermappingtable
    return children
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetYangName() string { return "CISCO-IETF-PW-MIB" }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) SetParent(parent types.Entity) { cISCOIETFPWMIB.parent = parent }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetParent() types.Entity { return cISCOIETFPWMIB.parent }

func (cISCOIETFPWMIB *CISCOIETFPWMIB) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcobjects
type CISCOIETFPWMIB_Cpwvcobjects struct {
    parent types.Entity
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

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetFilter() yfilter.YFilter { return cpwvcobjects.YFilter }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) SetFilter(yf yfilter.YFilter) { cpwvcobjects.YFilter = yf }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetGoName(yname string) string {
    if yname == "cpwVcIndexNext" { return "Cpwvcindexnext" }
    if yname == "cpwVcPerfTotalErrorPackets" { return "Cpwvcperftotalerrorpackets" }
    if yname == "cpwVcUpDownNotifEnable" { return "Cpwvcupdownnotifenable" }
    if yname == "cpwVcNotifRate" { return "Cpwvcnotifrate" }
    return ""
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetSegmentPath() string {
    return "cpwVcObjects"
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndexNext"] = cpwvcobjects.Cpwvcindexnext
    leafs["cpwVcPerfTotalErrorPackets"] = cpwvcobjects.Cpwvcperftotalerrorpackets
    leafs["cpwVcUpDownNotifEnable"] = cpwvcobjects.Cpwvcupdownnotifenable
    leafs["cpwVcNotifRate"] = cpwvcobjects.Cpwvcnotifrate
    return leafs
}

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetYangName() string { return "cpwVcObjects" }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) SetParent(parent types.Entity) { cpwvcobjects.parent = parent }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetParent() types.Entity { return cpwvcobjects.parent }

func (cpwvcobjects *CISCOIETFPWMIB_Cpwvcobjects) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvctable
// This table specifies information for connecting various 
// emulated services to various tunnel type.
type CISCOIETFPWMIB_Cpwvctable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A row in this table represents an emulated virtual  connection (VC) across
    // a packet network. It is indexed by  cpwVcIndex, which uniquely identifying
    // a singular   connection. . The type is slice of
    // CISCOIETFPWMIB_Cpwvctable_Cpwvcentry.
    Cpwvcentry []CISCOIETFPWMIB_Cpwvctable_Cpwvcentry
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetFilter() yfilter.YFilter { return cpwvctable.YFilter }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) SetFilter(yf yfilter.YFilter) { cpwvctable.YFilter = yf }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetGoName(yname string) string {
    if yname == "cpwVcEntry" { return "Cpwvcentry" }
    return ""
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetSegmentPath() string {
    return "cpwVcTable"
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcEntry" {
        for _, c := range cpwvctable.Cpwvcentry {
            if cpwvctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvctable_Cpwvcentry{}
        cpwvctable.Cpwvcentry = append(cpwvctable.Cpwvcentry, child)
        return &cpwvctable.Cpwvcentry[len(cpwvctable.Cpwvcentry)-1]
    }
    return nil
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvctable.Cpwvcentry {
        children[cpwvctable.Cpwvcentry[i].GetSegmentPath()] = &cpwvctable.Cpwvcentry[i]
    }
    return children
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetYangName() string { return "cpwVcTable" }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) SetParent(parent types.Entity) { cpwvctable.parent = parent }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetParent() types.Entity { return cpwvctable.parent }

func (cpwvctable *CISCOIETFPWMIB_Cpwvctable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvctable_Cpwvcentry
// A row in this table represents an emulated virtual 
// connection (VC) across a packet network. It is indexed by 
// cpwVcIndex, which uniquely identifying a singular  
// connection. 
type CISCOIETFPWMIB_Cpwvctable_Cpwvcentry struct {
    parent types.Entity
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

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetFilter() yfilter.YFilter { return cpwvcentry.YFilter }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) SetFilter(yf yfilter.YFilter) { cpwvcentry.YFilter = yf }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcType" { return "Cpwvctype" }
    if yname == "cpwVcOwner" { return "Cpwvcowner" }
    if yname == "cpwVcPsnType" { return "Cpwvcpsntype" }
    if yname == "cpwVcSetUpPriority" { return "Cpwvcsetuppriority" }
    if yname == "cpwVcHoldingPriority" { return "Cpwvcholdingpriority" }
    if yname == "cpwVcInboundMode" { return "Cpwvcinboundmode" }
    if yname == "cpwVcPeerAddrType" { return "Cpwvcpeeraddrtype" }
    if yname == "cpwVcPeerAddr" { return "Cpwvcpeeraddr" }
    if yname == "cpwVcID" { return "Cpwvcid" }
    if yname == "cpwVcLocalGroupID" { return "Cpwvclocalgroupid" }
    if yname == "cpwVcControlWord" { return "Cpwvccontrolword" }
    if yname == "cpwVcLocalIfMtu" { return "Cpwvclocalifmtu" }
    if yname == "cpwVcLocalIfString" { return "Cpwvclocalifstring" }
    if yname == "cpwVcRemoteGroupID" { return "Cpwvcremotegroupid" }
    if yname == "cpwVcRemoteControlWord" { return "Cpwvcremotecontrolword" }
    if yname == "cpwVcRemoteIfMtu" { return "Cpwvcremoteifmtu" }
    if yname == "cpwVcRemoteIfString" { return "Cpwvcremoteifstring" }
    if yname == "cpwVcOutboundVcLabel" { return "Cpwvcoutboundvclabel" }
    if yname == "cpwVcInboundVcLabel" { return "Cpwvcinboundvclabel" }
    if yname == "cpwVcName" { return "Cpwvcname" }
    if yname == "cpwVcDescr" { return "Cpwvcdescr" }
    if yname == "cpwVcCreateTime" { return "Cpwvccreatetime" }
    if yname == "cpwVcUpTime" { return "Cpwvcuptime" }
    if yname == "cpwVcAdminStatus" { return "Cpwvcadminstatus" }
    if yname == "cpwVcOperStatus" { return "Cpwvcoperstatus" }
    if yname == "cpwVcInboundOperStatus" { return "Cpwvcinboundoperstatus" }
    if yname == "cpwVcOutboundOperStatus" { return "Cpwvcoutboundoperstatus" }
    if yname == "cpwVcTimeElapsed" { return "Cpwvctimeelapsed" }
    if yname == "cpwVcValidIntervals" { return "Cpwvcvalidintervals" }
    if yname == "cpwVcRowStatus" { return "Cpwvcrowstatus" }
    if yname == "cpwVcStorageType" { return "Cpwvcstoragetype" }
    return ""
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetSegmentPath() string {
    return "cpwVcEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcentry.Cpwvcindex) + "']"
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcentry.Cpwvcindex
    leafs["cpwVcType"] = cpwvcentry.Cpwvctype
    leafs["cpwVcOwner"] = cpwvcentry.Cpwvcowner
    leafs["cpwVcPsnType"] = cpwvcentry.Cpwvcpsntype
    leafs["cpwVcSetUpPriority"] = cpwvcentry.Cpwvcsetuppriority
    leafs["cpwVcHoldingPriority"] = cpwvcentry.Cpwvcholdingpriority
    leafs["cpwVcInboundMode"] = cpwvcentry.Cpwvcinboundmode
    leafs["cpwVcPeerAddrType"] = cpwvcentry.Cpwvcpeeraddrtype
    leafs["cpwVcPeerAddr"] = cpwvcentry.Cpwvcpeeraddr
    leafs["cpwVcID"] = cpwvcentry.Cpwvcid
    leafs["cpwVcLocalGroupID"] = cpwvcentry.Cpwvclocalgroupid
    leafs["cpwVcControlWord"] = cpwvcentry.Cpwvccontrolword
    leafs["cpwVcLocalIfMtu"] = cpwvcentry.Cpwvclocalifmtu
    leafs["cpwVcLocalIfString"] = cpwvcentry.Cpwvclocalifstring
    leafs["cpwVcRemoteGroupID"] = cpwvcentry.Cpwvcremotegroupid
    leafs["cpwVcRemoteControlWord"] = cpwvcentry.Cpwvcremotecontrolword
    leafs["cpwVcRemoteIfMtu"] = cpwvcentry.Cpwvcremoteifmtu
    leafs["cpwVcRemoteIfString"] = cpwvcentry.Cpwvcremoteifstring
    leafs["cpwVcOutboundVcLabel"] = cpwvcentry.Cpwvcoutboundvclabel
    leafs["cpwVcInboundVcLabel"] = cpwvcentry.Cpwvcinboundvclabel
    leafs["cpwVcName"] = cpwvcentry.Cpwvcname
    leafs["cpwVcDescr"] = cpwvcentry.Cpwvcdescr
    leafs["cpwVcCreateTime"] = cpwvcentry.Cpwvccreatetime
    leafs["cpwVcUpTime"] = cpwvcentry.Cpwvcuptime
    leafs["cpwVcAdminStatus"] = cpwvcentry.Cpwvcadminstatus
    leafs["cpwVcOperStatus"] = cpwvcentry.Cpwvcoperstatus
    leafs["cpwVcInboundOperStatus"] = cpwvcentry.Cpwvcinboundoperstatus
    leafs["cpwVcOutboundOperStatus"] = cpwvcentry.Cpwvcoutboundoperstatus
    leafs["cpwVcTimeElapsed"] = cpwvcentry.Cpwvctimeelapsed
    leafs["cpwVcValidIntervals"] = cpwvcentry.Cpwvcvalidintervals
    leafs["cpwVcRowStatus"] = cpwvcentry.Cpwvcrowstatus
    leafs["cpwVcStorageType"] = cpwvcentry.Cpwvcstoragetype
    return leafs
}

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetYangName() string { return "cpwVcEntry" }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) SetParent(parent types.Entity) { cpwvcentry.parent = parent }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetParent() types.Entity { return cpwvcentry.parent }

func (cpwvcentry *CISCOIETFPWMIB_Cpwvctable_Cpwvcentry) GetParentYangName() string { return "cpwVcTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for  every VC. The type is
    // slice of CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry.
    Cpwvcperfcurrententry []CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetFilter() yfilter.YFilter { return cpwvcperfcurrenttable.YFilter }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) SetFilter(yf yfilter.YFilter) { cpwvcperfcurrenttable.YFilter = yf }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetGoName(yname string) string {
    if yname == "cpwVcPerfCurrentEntry" { return "Cpwvcperfcurrententry" }
    return ""
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetSegmentPath() string {
    return "cpwVcPerfCurrentTable"
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcPerfCurrentEntry" {
        for _, c := range cpwvcperfcurrenttable.Cpwvcperfcurrententry {
            if cpwvcperfcurrenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry{}
        cpwvcperfcurrenttable.Cpwvcperfcurrententry = append(cpwvcperfcurrenttable.Cpwvcperfcurrententry, child)
        return &cpwvcperfcurrenttable.Cpwvcperfcurrententry[len(cpwvcperfcurrenttable.Cpwvcperfcurrententry)-1]
    }
    return nil
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcperfcurrenttable.Cpwvcperfcurrententry {
        children[cpwvcperfcurrenttable.Cpwvcperfcurrententry[i].GetSegmentPath()] = &cpwvcperfcurrenttable.Cpwvcperfcurrententry[i]
    }
    return children
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetYangName() string { return "cpwVcPerfCurrentTable" }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) SetParent(parent types.Entity) { cpwvcperfcurrenttable.parent = parent }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetParent() types.Entity { return cpwvcperfcurrenttable.parent }

func (cpwvcperfcurrenttable *CISCOIETFPWMIB_Cpwvcperfcurrenttable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry
// An entry in this table is created by the agent for 
// every VC.
type CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry struct {
    parent types.Entity
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

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetFilter() yfilter.YFilter { return cpwvcperfcurrententry.YFilter }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) SetFilter(yf yfilter.YFilter) { cpwvcperfcurrententry.YFilter = yf }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcPerfCurrentInHCPackets" { return "Cpwvcperfcurrentinhcpackets" }
    if yname == "cpwVcPerfCurrentInHCBytes" { return "Cpwvcperfcurrentinhcbytes" }
    if yname == "cpwVcPerfCurrentOutHCPackets" { return "Cpwvcperfcurrentouthcpackets" }
    if yname == "cpwVcPerfCurrentOutHCBytes" { return "Cpwvcperfcurrentouthcbytes" }
    return ""
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetSegmentPath() string {
    return "cpwVcPerfCurrentEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperfcurrententry.Cpwvcindex) + "']"
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcperfcurrententry.Cpwvcindex
    leafs["cpwVcPerfCurrentInHCPackets"] = cpwvcperfcurrententry.Cpwvcperfcurrentinhcpackets
    leafs["cpwVcPerfCurrentInHCBytes"] = cpwvcperfcurrententry.Cpwvcperfcurrentinhcbytes
    leafs["cpwVcPerfCurrentOutHCPackets"] = cpwvcperfcurrententry.Cpwvcperfcurrentouthcpackets
    leafs["cpwVcPerfCurrentOutHCBytes"] = cpwvcperfcurrententry.Cpwvcperfcurrentouthcbytes
    return leafs
}

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetYangName() string { return "cpwVcPerfCurrentEntry" }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) SetParent(parent types.Entity) { cpwvcperfcurrententry.parent = parent }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetParent() types.Entity { return cpwvcperfcurrententry.parent }

func (cpwvcperfcurrententry *CISCOIETFPWMIB_Cpwvcperfcurrenttable_Cpwvcperfcurrententry) GetParentYangName() string { return "cpwVcPerfCurrentTable" }

// CISCOIETFPWMIB_Cpwvcperfintervaltable
// This table provides per-VC performance information for  
// each interval.
type CISCOIETFPWMIB_Cpwvcperfintervaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry.
    Cpwvcperfintervalentry []CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetFilter() yfilter.YFilter { return cpwvcperfintervaltable.YFilter }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) SetFilter(yf yfilter.YFilter) { cpwvcperfintervaltable.YFilter = yf }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetGoName(yname string) string {
    if yname == "cpwVcPerfIntervalEntry" { return "Cpwvcperfintervalentry" }
    return ""
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetSegmentPath() string {
    return "cpwVcPerfIntervalTable"
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcPerfIntervalEntry" {
        for _, c := range cpwvcperfintervaltable.Cpwvcperfintervalentry {
            if cpwvcperfintervaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry{}
        cpwvcperfintervaltable.Cpwvcperfintervalentry = append(cpwvcperfintervaltable.Cpwvcperfintervalentry, child)
        return &cpwvcperfintervaltable.Cpwvcperfintervalentry[len(cpwvcperfintervaltable.Cpwvcperfintervalentry)-1]
    }
    return nil
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcperfintervaltable.Cpwvcperfintervalentry {
        children[cpwvcperfintervaltable.Cpwvcperfintervalentry[i].GetSegmentPath()] = &cpwvcperfintervaltable.Cpwvcperfintervalentry[i]
    }
    return children
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetYangName() string { return "cpwVcPerfIntervalTable" }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) SetParent(parent types.Entity) { cpwvcperfintervaltable.parent = parent }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetParent() types.Entity { return cpwvcperfintervaltable.parent }

func (cpwvcperfintervaltable *CISCOIETFPWMIB_Cpwvcperfintervaltable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry struct {
    parent types.Entity
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

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetFilter() yfilter.YFilter { return cpwvcperfintervalentry.YFilter }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) SetFilter(yf yfilter.YFilter) { cpwvcperfintervalentry.YFilter = yf }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcPerfIntervalNumber" { return "Cpwvcperfintervalnumber" }
    if yname == "cpwVcPerfIntervalValidData" { return "Cpwvcperfintervalvaliddata" }
    if yname == "cpwVcPerfIntervalTimeElapsed" { return "Cpwvcperfintervaltimeelapsed" }
    if yname == "cpwVcPerfIntervalInHCPackets" { return "Cpwvcperfintervalinhcpackets" }
    if yname == "cpwVcPerfIntervalInHCBytes" { return "Cpwvcperfintervalinhcbytes" }
    if yname == "cpwVcPerfIntervalOutHCPackets" { return "Cpwvcperfintervalouthcpackets" }
    if yname == "cpwVcPerfIntervalOutHCBytes" { return "Cpwvcperfintervalouthcbytes" }
    return ""
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetSegmentPath() string {
    return "cpwVcPerfIntervalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperfintervalentry.Cpwvcindex) + "']" + "[cpwVcPerfIntervalNumber='" + fmt.Sprintf("%v", cpwvcperfintervalentry.Cpwvcperfintervalnumber) + "']"
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcperfintervalentry.Cpwvcindex
    leafs["cpwVcPerfIntervalNumber"] = cpwvcperfintervalentry.Cpwvcperfintervalnumber
    leafs["cpwVcPerfIntervalValidData"] = cpwvcperfintervalentry.Cpwvcperfintervalvaliddata
    leafs["cpwVcPerfIntervalTimeElapsed"] = cpwvcperfintervalentry.Cpwvcperfintervaltimeelapsed
    leafs["cpwVcPerfIntervalInHCPackets"] = cpwvcperfintervalentry.Cpwvcperfintervalinhcpackets
    leafs["cpwVcPerfIntervalInHCBytes"] = cpwvcperfintervalentry.Cpwvcperfintervalinhcbytes
    leafs["cpwVcPerfIntervalOutHCPackets"] = cpwvcperfintervalentry.Cpwvcperfintervalouthcpackets
    leafs["cpwVcPerfIntervalOutHCBytes"] = cpwvcperfintervalentry.Cpwvcperfintervalouthcbytes
    return leafs
}

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetYangName() string { return "cpwVcPerfIntervalEntry" }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) SetParent(parent types.Entity) { cpwvcperfintervalentry.parent = parent }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetParent() types.Entity { return cpwvcperfintervalentry.parent }

func (cpwvcperfintervalentry *CISCOIETFPWMIB_Cpwvcperfintervaltable_Cpwvcperfintervalentry) GetParentYangName() string { return "cpwVcPerfIntervalTable" }

// CISCOIETFPWMIB_Cpwvcperftotaltable
// This table provides per-VC Performance information from VC  
// start time.
type CISCOIETFPWMIB_Cpwvcperftotaltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created agent for every VC. The type is slice of
    // CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry.
    Cpwvcperftotalentry []CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetFilter() yfilter.YFilter { return cpwvcperftotaltable.YFilter }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) SetFilter(yf yfilter.YFilter) { cpwvcperftotaltable.YFilter = yf }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetGoName(yname string) string {
    if yname == "cpwVcPerfTotalEntry" { return "Cpwvcperftotalentry" }
    return ""
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetSegmentPath() string {
    return "cpwVcPerfTotalTable"
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcPerfTotalEntry" {
        for _, c := range cpwvcperftotaltable.Cpwvcperftotalentry {
            if cpwvcperftotaltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry{}
        cpwvcperftotaltable.Cpwvcperftotalentry = append(cpwvcperftotaltable.Cpwvcperftotalentry, child)
        return &cpwvcperftotaltable.Cpwvcperftotalentry[len(cpwvcperftotaltable.Cpwvcperftotalentry)-1]
    }
    return nil
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcperftotaltable.Cpwvcperftotalentry {
        children[cpwvcperftotaltable.Cpwvcperftotalentry[i].GetSegmentPath()] = &cpwvcperftotaltable.Cpwvcperftotalentry[i]
    }
    return children
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetYangName() string { return "cpwVcPerfTotalTable" }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) SetParent(parent types.Entity) { cpwvcperftotaltable.parent = parent }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetParent() types.Entity { return cpwvcperftotaltable.parent }

func (cpwvcperftotaltable *CISCOIETFPWMIB_Cpwvcperftotaltable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry
// An entry in this table is created agent for every VC.
type CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry struct {
    parent types.Entity
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

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetFilter() yfilter.YFilter { return cpwvcperftotalentry.YFilter }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) SetFilter(yf yfilter.YFilter) { cpwvcperftotalentry.YFilter = yf }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetGoName(yname string) string {
    if yname == "cpwVcIndex" { return "Cpwvcindex" }
    if yname == "cpwVcPerfTotalInHCPackets" { return "Cpwvcperftotalinhcpackets" }
    if yname == "cpwVcPerfTotalInHCBytes" { return "Cpwvcperftotalinhcbytes" }
    if yname == "cpwVcPerfTotalOutHCPackets" { return "Cpwvcperftotalouthcpackets" }
    if yname == "cpwVcPerfTotalOutHCBytes" { return "Cpwvcperftotalouthcbytes" }
    if yname == "cpwVcPerfTotalDiscontinuityTime" { return "Cpwvcperftotaldiscontinuitytime" }
    return ""
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetSegmentPath() string {
    return "cpwVcPerfTotalEntry" + "[cpwVcIndex='" + fmt.Sprintf("%v", cpwvcperftotalentry.Cpwvcindex) + "']"
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIndex"] = cpwvcperftotalentry.Cpwvcindex
    leafs["cpwVcPerfTotalInHCPackets"] = cpwvcperftotalentry.Cpwvcperftotalinhcpackets
    leafs["cpwVcPerfTotalInHCBytes"] = cpwvcperftotalentry.Cpwvcperftotalinhcbytes
    leafs["cpwVcPerfTotalOutHCPackets"] = cpwvcperftotalentry.Cpwvcperftotalouthcpackets
    leafs["cpwVcPerfTotalOutHCBytes"] = cpwvcperftotalentry.Cpwvcperftotalouthcbytes
    leafs["cpwVcPerfTotalDiscontinuityTime"] = cpwvcperftotalentry.Cpwvcperftotaldiscontinuitytime
    return leafs
}

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetYangName() string { return "cpwVcPerfTotalEntry" }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) SetParent(parent types.Entity) { cpwvcperftotalentry.parent = parent }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetParent() types.Entity { return cpwvcperftotalentry.parent }

func (cpwvcperftotalentry *CISCOIETFPWMIB_Cpwvcperftotaltable_Cpwvcperftotalentry) GetParentYangName() string { return "cpwVcPerfTotalTable" }

// CISCOIETFPWMIB_Cpwvcidmappingtable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_Cpwvcidmappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured by
    // the cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry.
    Cpwvcidmappingentry []CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetFilter() yfilter.YFilter { return cpwvcidmappingtable.YFilter }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) SetFilter(yf yfilter.YFilter) { cpwvcidmappingtable.YFilter = yf }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetGoName(yname string) string {
    if yname == "cpwVcIdMappingEntry" { return "Cpwvcidmappingentry" }
    return ""
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetSegmentPath() string {
    return "cpwVcIdMappingTable"
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcIdMappingEntry" {
        for _, c := range cpwvcidmappingtable.Cpwvcidmappingentry {
            if cpwvcidmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry{}
        cpwvcidmappingtable.Cpwvcidmappingentry = append(cpwvcidmappingtable.Cpwvcidmappingentry, child)
        return &cpwvcidmappingtable.Cpwvcidmappingentry[len(cpwvcidmappingtable.Cpwvcidmappingentry)-1]
    }
    return nil
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcidmappingtable.Cpwvcidmappingentry {
        children[cpwvcidmappingtable.Cpwvcidmappingentry[i].GetSegmentPath()] = &cpwvcidmappingtable.Cpwvcidmappingentry[i]
    }
    return children
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetYangName() string { return "cpwVcIdMappingTable" }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) SetParent(parent types.Entity) { cpwvcidmappingtable.parent = parent }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetParent() types.Entity { return cpwvcidmappingtable.parent }

func (cpwvcidmappingtable *CISCOIETFPWMIB_Cpwvcidmappingtable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry
// An entry in this table is created by the agent for every  
// VC configured by the cpwVcTable.
type CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry struct {
    parent types.Entity
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

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetFilter() yfilter.YFilter { return cpwvcidmappingentry.YFilter }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) SetFilter(yf yfilter.YFilter) { cpwvcidmappingentry.YFilter = yf }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetGoName(yname string) string {
    if yname == "cpwVcIdMappingVcType" { return "Cpwvcidmappingvctype" }
    if yname == "cpwVcIdMappingVcID" { return "Cpwvcidmappingvcid" }
    if yname == "cpwVcIdMappingPeerAddrType" { return "Cpwvcidmappingpeeraddrtype" }
    if yname == "cpwVcIdMappingPeerAddr" { return "Cpwvcidmappingpeeraddr" }
    if yname == "cpwVcIdMappingVcIndex" { return "Cpwvcidmappingvcindex" }
    return ""
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetSegmentPath() string {
    return "cpwVcIdMappingEntry" + "[cpwVcIdMappingVcType='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvctype) + "']" + "[cpwVcIdMappingVcID='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvcid) + "']" + "[cpwVcIdMappingPeerAddrType='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingpeeraddrtype) + "']" + "[cpwVcIdMappingPeerAddr='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingpeeraddr) + "']" + "[cpwVcIdMappingVcIndex='" + fmt.Sprintf("%v", cpwvcidmappingentry.Cpwvcidmappingvcindex) + "']"
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcIdMappingVcType"] = cpwvcidmappingentry.Cpwvcidmappingvctype
    leafs["cpwVcIdMappingVcID"] = cpwvcidmappingentry.Cpwvcidmappingvcid
    leafs["cpwVcIdMappingPeerAddrType"] = cpwvcidmappingentry.Cpwvcidmappingpeeraddrtype
    leafs["cpwVcIdMappingPeerAddr"] = cpwvcidmappingentry.Cpwvcidmappingpeeraddr
    leafs["cpwVcIdMappingVcIndex"] = cpwvcidmappingentry.Cpwvcidmappingvcindex
    return leafs
}

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetYangName() string { return "cpwVcIdMappingEntry" }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) SetParent(parent types.Entity) { cpwvcidmappingentry.parent = parent }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetParent() types.Entity { return cpwvcidmappingentry.parent }

func (cpwvcidmappingentry *CISCOIETFPWMIB_Cpwvcidmappingtable_Cpwvcidmappingentry) GetParentYangName() string { return "cpwVcIdMappingTable" }

// CISCOIETFPWMIB_Cpwvcpeermappingtable
// This table provides reverse mapping of the existing VCs  
// based on vc type and VC ID ordering. This table is  
// typically useful for EMS ordered query of existing VCs.
type CISCOIETFPWMIB_Cpwvcpeermappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the agent for every   VC configured in
    // cpwVcTable. The type is slice of
    // CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry.
    Cpwvcpeermappingentry []CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetFilter() yfilter.YFilter { return cpwvcpeermappingtable.YFilter }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) SetFilter(yf yfilter.YFilter) { cpwvcpeermappingtable.YFilter = yf }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetGoName(yname string) string {
    if yname == "cpwVcPeerMappingEntry" { return "Cpwvcpeermappingentry" }
    return ""
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetSegmentPath() string {
    return "cpwVcPeerMappingTable"
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpwVcPeerMappingEntry" {
        for _, c := range cpwvcpeermappingtable.Cpwvcpeermappingentry {
            if cpwvcpeermappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry{}
        cpwvcpeermappingtable.Cpwvcpeermappingentry = append(cpwvcpeermappingtable.Cpwvcpeermappingentry, child)
        return &cpwvcpeermappingtable.Cpwvcpeermappingentry[len(cpwvcpeermappingtable.Cpwvcpeermappingentry)-1]
    }
    return nil
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpwvcpeermappingtable.Cpwvcpeermappingentry {
        children[cpwvcpeermappingtable.Cpwvcpeermappingentry[i].GetSegmentPath()] = &cpwvcpeermappingtable.Cpwvcpeermappingentry[i]
    }
    return children
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetYangName() string { return "cpwVcPeerMappingTable" }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) SetParent(parent types.Entity) { cpwvcpeermappingtable.parent = parent }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetParent() types.Entity { return cpwvcpeermappingtable.parent }

func (cpwvcpeermappingtable *CISCOIETFPWMIB_Cpwvcpeermappingtable) GetParentYangName() string { return "CISCO-IETF-PW-MIB" }

// CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry
// An entry in this table is created by the agent for every  
// VC configured in cpwVcTable.
type CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry struct {
    parent types.Entity
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

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetFilter() yfilter.YFilter { return cpwvcpeermappingentry.YFilter }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) SetFilter(yf yfilter.YFilter) { cpwvcpeermappingentry.YFilter = yf }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetGoName(yname string) string {
    if yname == "cpwVcPeerMappingPeerAddrType" { return "Cpwvcpeermappingpeeraddrtype" }
    if yname == "cpwVcPeerMappingPeerAddr" { return "Cpwvcpeermappingpeeraddr" }
    if yname == "cpwVcPeerMappingVcType" { return "Cpwvcpeermappingvctype" }
    if yname == "cpwVcPeerMappingVcID" { return "Cpwvcpeermappingvcid" }
    if yname == "cpwVcPeerMappingVcIndex" { return "Cpwvcpeermappingvcindex" }
    return ""
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetSegmentPath() string {
    return "cpwVcPeerMappingEntry" + "[cpwVcPeerMappingPeerAddrType='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddrtype) + "']" + "[cpwVcPeerMappingPeerAddr='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingpeeraddr) + "']" + "[cpwVcPeerMappingVcType='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvctype) + "']" + "[cpwVcPeerMappingVcID='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvcid) + "']" + "[cpwVcPeerMappingVcIndex='" + fmt.Sprintf("%v", cpwvcpeermappingentry.Cpwvcpeermappingvcindex) + "']"
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cpwVcPeerMappingPeerAddrType"] = cpwvcpeermappingentry.Cpwvcpeermappingpeeraddrtype
    leafs["cpwVcPeerMappingPeerAddr"] = cpwvcpeermappingentry.Cpwvcpeermappingpeeraddr
    leafs["cpwVcPeerMappingVcType"] = cpwvcpeermappingentry.Cpwvcpeermappingvctype
    leafs["cpwVcPeerMappingVcID"] = cpwvcpeermappingentry.Cpwvcpeermappingvcid
    leafs["cpwVcPeerMappingVcIndex"] = cpwvcpeermappingentry.Cpwvcpeermappingvcindex
    return leafs
}

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetYangName() string { return "cpwVcPeerMappingEntry" }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) SetParent(parent types.Entity) { cpwvcpeermappingentry.parent = parent }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetParent() types.Entity { return cpwvcpeermappingentry.parent }

func (cpwvcpeermappingentry *CISCOIETFPWMIB_Cpwvcpeermappingtable_Cpwvcpeermappingentry) GetParentYangName() string { return "cpwVcPeerMappingTable" }

