// An extension to the MIB module defined in
// RFC 1850 for managing OSPF implimentation. 
// Most of the MIB definitions are based on
// the IETF draft 
// < draft-ietf-ospf-mib-update-05.txt > . 
// Support for OSPF Sham link is also added
package cisco_ospf_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ospf_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-OSPF-MIB CISCO-OSPF-MIB}", reflect.TypeOf(CISCOOSPFMIB{}))
    ydk.RegisterEntity("CISCO-OSPF-MIB:CISCO-OSPF-MIB", reflect.TypeOf(CISCOOSPFMIB{}))
}

// CISCOOSPFMIB
type CISCOOSPFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cospfgeneralgroup CISCOOSPFMIB_Cospfgeneralgroup

    // The OSPF Process's Link State Database. This  table is meant for Opaque
    // LSA's.
    Cospflsdbtable CISCOOSPFMIB_Cospflsdbtable

    // Information about this router's sham links.
    Cospfshamlinktable CISCOOSPFMIB_Cospfshamlinktable

    // The OSPF Process's Link-Local Link State Database for non-virtual links.
    Cospflocallsdbtable CISCOOSPFMIB_Cospflocallsdbtable

    // The OSPF Process's Link-Local Link State Database for virtual links.
    Cospfvirtlocallsdbtable CISCOOSPFMIB_Cospfvirtlocallsdbtable

    // A table of sham link neighbor information.
    Cospfshamlinknbrtable CISCOOSPFMIB_Cospfshamlinknbrtable

    // Information about this router's sham links.
    Cospfshamlinkstable CISCOOSPFMIB_Cospfshamlinkstable
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetEntityData() *types.CommonEntityData {
    cISCOOSPFMIB.EntityData.YFilter = cISCOOSPFMIB.YFilter
    cISCOOSPFMIB.EntityData.YangName = "CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOOSPFMIB.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.SegmentPath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOOSPFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOOSPFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOOSPFMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOOSPFMIB.EntityData.Children["cospfGeneralGroup"] = types.YChild{"Cospfgeneralgroup", &cISCOOSPFMIB.Cospfgeneralgroup}
    cISCOOSPFMIB.EntityData.Children["cospfLsdbTable"] = types.YChild{"Cospflsdbtable", &cISCOOSPFMIB.Cospflsdbtable}
    cISCOOSPFMIB.EntityData.Children["cospfShamLinkTable"] = types.YChild{"Cospfshamlinktable", &cISCOOSPFMIB.Cospfshamlinktable}
    cISCOOSPFMIB.EntityData.Children["cospfLocalLsdbTable"] = types.YChild{"Cospflocallsdbtable", &cISCOOSPFMIB.Cospflocallsdbtable}
    cISCOOSPFMIB.EntityData.Children["cospfVirtLocalLsdbTable"] = types.YChild{"Cospfvirtlocallsdbtable", &cISCOOSPFMIB.Cospfvirtlocallsdbtable}
    cISCOOSPFMIB.EntityData.Children["cospfShamLinkNbrTable"] = types.YChild{"Cospfshamlinknbrtable", &cISCOOSPFMIB.Cospfshamlinknbrtable}
    cISCOOSPFMIB.EntityData.Children["cospfShamLinksTable"] = types.YChild{"Cospfshamlinkstable", &cISCOOSPFMIB.Cospfshamlinkstable}
    cISCOOSPFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOOSPFMIB.EntityData)
}

// CISCOOSPFMIB_Cospfgeneralgroup
type CISCOOSPFMIB_Cospfgeneralgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates metrics used to choose among multiple AS- external-LSAs. When
    // cospfRFC1583Compatibility is set to enabled, only cost will be used when
    // choosing among multiple AS-external-LSAs advertising the same destination.
    // When cospfRFC1583Compatibility is set to disabled, preference will be
    // driven first by type of path using cost only to break ties. The type is
    // bool.
    Cospfrfc1583Compatibility interface{}

    // The router's support for Opaque LSA types. The type is bool.
    Cospfopaquelsasupport interface{}

    // The router's support for OSPF traffic engineering. The type is bool.
    Cospftrafficengineeringsupport interface{}

    // The total number of Opaque AS link-state advertisements in the link state
    // database. The type is interface{} with range: 0..4294967295.
    Cospfopaqueaslsacount interface{}

    // The 32-bit unsigned sum of the Opaque AS  link-state advertisements' LS
    // checksums contained link state database. The type is interface{} with
    // range: 0..4294967295.
    Cospfopaqueaslsacksumsum interface{}
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetEntityData() *types.CommonEntityData {
    cospfgeneralgroup.EntityData.YFilter = cospfgeneralgroup.YFilter
    cospfgeneralgroup.EntityData.YangName = "cospfGeneralGroup"
    cospfgeneralgroup.EntityData.BundleName = "cisco_ios_xe"
    cospfgeneralgroup.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfgeneralgroup.EntityData.SegmentPath = "cospfGeneralGroup"
    cospfgeneralgroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfgeneralgroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfgeneralgroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfgeneralgroup.EntityData.Children = make(map[string]types.YChild)
    cospfgeneralgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    cospfgeneralgroup.EntityData.Leafs["cospfRFC1583Compatibility"] = types.YLeaf{"Cospfrfc1583Compatibility", cospfgeneralgroup.Cospfrfc1583Compatibility}
    cospfgeneralgroup.EntityData.Leafs["cospfOpaqueLsaSupport"] = types.YLeaf{"Cospfopaquelsasupport", cospfgeneralgroup.Cospfopaquelsasupport}
    cospfgeneralgroup.EntityData.Leafs["cospfTrafficEngineeringSupport"] = types.YLeaf{"Cospftrafficengineeringsupport", cospfgeneralgroup.Cospftrafficengineeringsupport}
    cospfgeneralgroup.EntityData.Leafs["cospfOpaqueASLsaCount"] = types.YLeaf{"Cospfopaqueaslsacount", cospfgeneralgroup.Cospfopaqueaslsacount}
    cospfgeneralgroup.EntityData.Leafs["cospfOpaqueASLsaCksumSum"] = types.YLeaf{"Cospfopaqueaslsacksumsum", cospfgeneralgroup.Cospfopaqueaslsacksumsum}
    return &(cospfgeneralgroup.EntityData)
}

// CISCOOSPFMIB_Cospflsdbtable
// The OSPF Process's Link State Database. This 
// table is meant for Opaque LSA's
type CISCOOSPFMIB_Cospflsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry.
    Cospflsdbentry []CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetEntityData() *types.CommonEntityData {
    cospflsdbtable.EntityData.YFilter = cospflsdbtable.YFilter
    cospflsdbtable.EntityData.YangName = "cospfLsdbTable"
    cospflsdbtable.EntityData.BundleName = "cisco_ios_xe"
    cospflsdbtable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospflsdbtable.EntityData.SegmentPath = "cospfLsdbTable"
    cospflsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospflsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospflsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospflsdbtable.EntityData.Children = make(map[string]types.YChild)
    cospflsdbtable.EntityData.Children["cospfLsdbEntry"] = types.YChild{"Cospflsdbentry", nil}
    for i := range cospflsdbtable.Cospflsdbentry {
        cospflsdbtable.EntityData.Children[types.GetSegmentPath(&cospflsdbtable.Cospflsdbentry[i])] = types.YChild{"Cospflsdbentry", &cospflsdbtable.Cospflsdbentry[i]}
    }
    cospflsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospflsdbtable.EntityData)
}

// CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbareaid
    Ospflsdbareaid interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Cospflsdbtype.
    Cospflsdbtype interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdblsid
    Ospflsdblsid interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbrouterid
    Ospflsdbrouterid interface{}

    // The sequence number field is a  signed  32-bit integer.   It  is used to
    // detect old and duplicate link state advertisements.  The  space  of
    // sequence  numbers  is  linearly  ordered.   The larger the sequence number
    // the more recent  the advertisement. The type is interface{} with range:
    // 1..147483647.
    Cospflsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: 0..2147483647.
    Cospflsdbage interface{}

    // This field is the  checksum  of  the  complete contents  of  the 
    // advertisement, excepting the age field.  The age field is excepted  so 
    // that an   advertisement's  age  can  be  incremented without updating the 
    // checksum.   The  checksum used  is  the same that is used for ISO
    // connectionless datagrams; it is commonly referred  to as the Fletcher
    // checksum. The type is interface{} with range: 0..2147483647.
    Cospflsdbchecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    Cospflsdbadvertisement interface{}
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetEntityData() *types.CommonEntityData {
    cospflsdbentry.EntityData.YFilter = cospflsdbentry.YFilter
    cospflsdbentry.EntityData.YangName = "cospfLsdbEntry"
    cospflsdbentry.EntityData.BundleName = "cisco_ios_xe"
    cospflsdbentry.EntityData.ParentYangName = "cospfLsdbTable"
    cospflsdbentry.EntityData.SegmentPath = "cospfLsdbEntry" + "[ospfLsdbAreaId='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdbareaid) + "']" + "[cospfLsdbType='" + fmt.Sprintf("%v", cospflsdbentry.Cospflsdbtype) + "']" + "[ospfLsdbLsid='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdblsid) + "']" + "[ospfLsdbRouterId='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdbrouterid) + "']"
    cospflsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospflsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospflsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospflsdbentry.EntityData.Children = make(map[string]types.YChild)
    cospflsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospflsdbentry.EntityData.Leafs["ospfLsdbAreaId"] = types.YLeaf{"Ospflsdbareaid", cospflsdbentry.Ospflsdbareaid}
    cospflsdbentry.EntityData.Leafs["cospfLsdbType"] = types.YLeaf{"Cospflsdbtype", cospflsdbentry.Cospflsdbtype}
    cospflsdbentry.EntityData.Leafs["ospfLsdbLsid"] = types.YLeaf{"Ospflsdblsid", cospflsdbentry.Ospflsdblsid}
    cospflsdbentry.EntityData.Leafs["ospfLsdbRouterId"] = types.YLeaf{"Ospflsdbrouterid", cospflsdbentry.Ospflsdbrouterid}
    cospflsdbentry.EntityData.Leafs["cospfLsdbSequence"] = types.YLeaf{"Cospflsdbsequence", cospflsdbentry.Cospflsdbsequence}
    cospflsdbentry.EntityData.Leafs["cospfLsdbAge"] = types.YLeaf{"Cospflsdbage", cospflsdbentry.Cospflsdbage}
    cospflsdbentry.EntityData.Leafs["cospfLsdbChecksum"] = types.YLeaf{"Cospflsdbchecksum", cospflsdbentry.Cospflsdbchecksum}
    cospflsdbentry.EntityData.Leafs["cospfLsdbAdvertisement"] = types.YLeaf{"Cospflsdbadvertisement", cospflsdbentry.Cospflsdbadvertisement}
    return &(cospflsdbentry.EntityData)
}

// CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype string

const (
    CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype_areaOpaqueLink CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype = "areaOpaqueLink"

    CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype_asOpaqueLink CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype = "asOpaqueLink"
)

// CISCOOSPFMIB_Cospfshamlinktable
// Information about this router's sham links
type CISCOOSPFMIB_Cospfshamlinktable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry.
    Cospfshamlinkentry []CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetEntityData() *types.CommonEntityData {
    cospfshamlinktable.EntityData.YFilter = cospfshamlinktable.YFilter
    cospfshamlinktable.EntityData.YangName = "cospfShamLinkTable"
    cospfshamlinktable.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinktable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfshamlinktable.EntityData.SegmentPath = "cospfShamLinkTable"
    cospfshamlinktable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinktable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinktable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinktable.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinktable.EntityData.Children["cospfShamLinkEntry"] = types.YChild{"Cospfshamlinkentry", nil}
    for i := range cospfshamlinktable.Cospfshamlinkentry {
        cospfshamlinktable.EntityData.Children[types.GetSegmentPath(&cospfshamlinktable.Cospfshamlinkentry[i])] = types.YChild{"Cospfshamlinkentry", &cospfshamlinktable.Cospfshamlinkentry[i]}
    }
    cospfshamlinktable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospfshamlinktable.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry
// Information about a single sham link
type CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The  Transit  Area  that  the   Virtual   Link
    // traverses.  By definition, this is not 0.0.0.0. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinkareaid interface{}

    // This attribute is a key. The Local IP address of the sham link. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinklocalipaddress interface{}

    // This attribute is a key. The Router ID of the other end router of the sham
    // link. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinkneighborid interface{}

    // The number of seconds between  link-state  advertisement retransmissions, 
    // for  adjacencies belonging to this  link.   This  value  is also used when
    // retransmitting database description   and  link-state  request  packets.
    // This value   should  be well over the expected round trip time. The type is
    // interface{} with range: 0..3600.
    Cospfshamlinkretransinterval interface{}

    // The length of time, in  seconds,  between  the Hello  packets that the
    // router sends on the sham link. The type is interface{} with range:
    // 1..65535.
    Cospfshamlinkhellointerval interface{}

    // The number of seconds that  a  router's  Hello packets  have  not been seen
    // before it's neighbors declare the router down.  This  should  be some 
    // multiple  of  the  Hello  interval. The type is interface{} with range:
    // 0..2147483647.
    Cospfshamlinkrtrdeadinterval interface{}

    // OSPF sham link states. The type is Cospfshamlinkstate.
    Cospfshamlinkstate interface{}

    // The number of state changes or error events on this sham link. The type is
    // interface{} with range: 0..4294967295.
    Cospfshamlinkevents interface{}

    // The Metric to be advertised. The type is interface{} with range: 0..65535.
    Cospfshamlinkmetric interface{}
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetEntityData() *types.CommonEntityData {
    cospfshamlinkentry.EntityData.YFilter = cospfshamlinkentry.YFilter
    cospfshamlinkentry.EntityData.YangName = "cospfShamLinkEntry"
    cospfshamlinkentry.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinkentry.EntityData.ParentYangName = "cospfShamLinkTable"
    cospfshamlinkentry.EntityData.SegmentPath = "cospfShamLinkEntry" + "[cospfShamLinkAreaId='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinkareaid) + "']" + "[cospfShamLinkLocalIpAddress='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinklocalipaddress) + "']" + "[cospfShamLinkNeighborId='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinkneighborid) + "']"
    cospfshamlinkentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinkentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinkentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinkentry.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinkentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkAreaId"] = types.YLeaf{"Cospfshamlinkareaid", cospfshamlinkentry.Cospfshamlinkareaid}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkLocalIpAddress"] = types.YLeaf{"Cospfshamlinklocalipaddress", cospfshamlinkentry.Cospfshamlinklocalipaddress}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkNeighborId"] = types.YLeaf{"Cospfshamlinkneighborid", cospfshamlinkentry.Cospfshamlinkneighborid}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkRetransInterval"] = types.YLeaf{"Cospfshamlinkretransinterval", cospfshamlinkentry.Cospfshamlinkretransinterval}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkHelloInterval"] = types.YLeaf{"Cospfshamlinkhellointerval", cospfshamlinkentry.Cospfshamlinkhellointerval}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkRtrDeadInterval"] = types.YLeaf{"Cospfshamlinkrtrdeadinterval", cospfshamlinkentry.Cospfshamlinkrtrdeadinterval}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkState"] = types.YLeaf{"Cospfshamlinkstate", cospfshamlinkentry.Cospfshamlinkstate}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkEvents"] = types.YLeaf{"Cospfshamlinkevents", cospfshamlinkentry.Cospfshamlinkevents}
    cospfshamlinkentry.EntityData.Leafs["cospfShamLinkMetric"] = types.YLeaf{"Cospfshamlinkmetric", cospfshamlinkentry.Cospfshamlinkmetric}
    return &(cospfshamlinkentry.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate represents OSPF sham link states.
type CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate string

const (
    CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate_down CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate = "down"

    CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate_pointToPoint CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry_Cospfshamlinkstate = "pointToPoint"
)

// CISCOOSPFMIB_Cospflocallsdbtable
// The OSPF Process's Link-Local Link State Database
// for non-virtual links.
type CISCOOSPFMIB_Cospflocallsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry.
    Cospflocallsdbentry []CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetEntityData() *types.CommonEntityData {
    cospflocallsdbtable.EntityData.YFilter = cospflocallsdbtable.YFilter
    cospflocallsdbtable.EntityData.YangName = "cospfLocalLsdbTable"
    cospflocallsdbtable.EntityData.BundleName = "cisco_ios_xe"
    cospflocallsdbtable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospflocallsdbtable.EntityData.SegmentPath = "cospfLocalLsdbTable"
    cospflocallsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospflocallsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospflocallsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospflocallsdbtable.EntityData.Children = make(map[string]types.YChild)
    cospflocallsdbtable.EntityData.Children["cospfLocalLsdbEntry"] = types.YChild{"Cospflocallsdbentry", nil}
    for i := range cospflocallsdbtable.Cospflocallsdbentry {
        cospflocallsdbtable.EntityData.Children[types.GetSegmentPath(&cospflocallsdbtable.Cospflocallsdbentry[i])] = types.YChild{"Cospflocallsdbentry", &cospflocallsdbtable.Cospflocallsdbentry[i]}
    }
    cospflocallsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospflocallsdbtable.EntityData)
}

// CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP Address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospflocallsdbipaddress interface{}

    // This attribute is a key. The Interface Index of the interface from which
    // the LSA was received if the interface is unnumbered. The type is
    // interface{} with range: 0..2147483647.
    Cospflocallsdbaddresslessif interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Cospflocallsdbtype.
    Cospflocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32 bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospflocallsdblsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospflocallsdbrouterid interface{}

    // The sequence number field is a signed 32-bit integer. It is used to detect
    // old and duplicate link state advertisements. The space of sequence numbers
    // is linearly ordered. The larger the sequence number the more recent the
    // advertisement. The type is interface{} with range: -2147483647..2147483647.
    Cospflocallsdbsequence interface{}

    // This field is the age of the link state advertisement  in seconds. The type
    // is interface{} with range: 0..3600.
    Cospflocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field. The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum. The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred  to as the Fletcher checksum. The type is interface{}
    // with range: 0..4294967295.
    Cospflocallsdbchecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    Cospflocallsdbadvertisement interface{}
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetEntityData() *types.CommonEntityData {
    cospflocallsdbentry.EntityData.YFilter = cospflocallsdbentry.YFilter
    cospflocallsdbentry.EntityData.YangName = "cospfLocalLsdbEntry"
    cospflocallsdbentry.EntityData.BundleName = "cisco_ios_xe"
    cospflocallsdbentry.EntityData.ParentYangName = "cospfLocalLsdbTable"
    cospflocallsdbentry.EntityData.SegmentPath = "cospfLocalLsdbEntry" + "[cospfLocalLsdbIpAddress='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbipaddress) + "']" + "[cospfLocalLsdbAddressLessIf='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbaddresslessif) + "']" + "[cospfLocalLsdbType='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbtype) + "']" + "[cospfLocalLsdbLsid='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdblsid) + "']" + "[cospfLocalLsdbRouterId='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbrouterid) + "']"
    cospflocallsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospflocallsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospflocallsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospflocallsdbentry.EntityData.Children = make(map[string]types.YChild)
    cospflocallsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbIpAddress"] = types.YLeaf{"Cospflocallsdbipaddress", cospflocallsdbentry.Cospflocallsdbipaddress}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbAddressLessIf"] = types.YLeaf{"Cospflocallsdbaddresslessif", cospflocallsdbentry.Cospflocallsdbaddresslessif}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbType"] = types.YLeaf{"Cospflocallsdbtype", cospflocallsdbentry.Cospflocallsdbtype}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbLsid"] = types.YLeaf{"Cospflocallsdblsid", cospflocallsdbentry.Cospflocallsdblsid}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbRouterId"] = types.YLeaf{"Cospflocallsdbrouterid", cospflocallsdbentry.Cospflocallsdbrouterid}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbSequence"] = types.YLeaf{"Cospflocallsdbsequence", cospflocallsdbentry.Cospflocallsdbsequence}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbAge"] = types.YLeaf{"Cospflocallsdbage", cospflocallsdbentry.Cospflocallsdbage}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbChecksum"] = types.YLeaf{"Cospflocallsdbchecksum", cospflocallsdbentry.Cospflocallsdbchecksum}
    cospflocallsdbentry.EntityData.Leafs["cospfLocalLsdbAdvertisement"] = types.YLeaf{"Cospflocallsdbadvertisement", cospflocallsdbentry.Cospflocallsdbadvertisement}
    return &(cospflocallsdbentry.EntityData)
}

// CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype string

const (
    CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype_localOpaqueLink CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype = "localOpaqueLink"
)

// CISCOOSPFMIB_Cospfvirtlocallsdbtable
// The OSPF Process's Link-Local Link State Database
// for virtual links.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry.
    Cospfvirtlocallsdbentry []CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetEntityData() *types.CommonEntityData {
    cospfvirtlocallsdbtable.EntityData.YFilter = cospfvirtlocallsdbtable.YFilter
    cospfvirtlocallsdbtable.EntityData.YangName = "cospfVirtLocalLsdbTable"
    cospfvirtlocallsdbtable.EntityData.BundleName = "cisco_ios_xe"
    cospfvirtlocallsdbtable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfvirtlocallsdbtable.EntityData.SegmentPath = "cospfVirtLocalLsdbTable"
    cospfvirtlocallsdbtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfvirtlocallsdbtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfvirtlocallsdbtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfvirtlocallsdbtable.EntityData.Children = make(map[string]types.YChild)
    cospfvirtlocallsdbtable.EntityData.Children["cospfVirtLocalLsdbEntry"] = types.YChild{"Cospfvirtlocallsdbentry", nil}
    for i := range cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry {
        cospfvirtlocallsdbtable.EntityData.Children[types.GetSegmentPath(&cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry[i])] = types.YChild{"Cospfvirtlocallsdbentry", &cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry[i]}
    }
    cospfvirtlocallsdbtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospfvirtlocallsdbtable.EntityData)
}

// CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Transit Area that the Virtual Link traverses.
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfvirtlocallsdbtransitarea interface{}

    // This attribute is a key. The Router ID of the Virtual Neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfvirtlocallsdbneighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each 
    // link state type has a separate advertisement format. The type is
    // Cospfvirtlocallsdbtype.
    Cospfvirtlocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32 bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfvirtlocallsdblsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfvirtlocallsdbrouterid interface{}

    // The sequence number field is a  signed  32-bit integer. It is used to
    // detect old and duplicate link state advertisements. The space of sequence
    // numbers is linearly ordered. The larger the sequence number the more recent
    // the advertisement. The type is interface{} with range:
    // -2147483647..2147483647.
    Cospfvirtlocallsdbsequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: 0..3600.
    Cospfvirtlocallsdbage interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field. The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum. The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred  to as the Fletcher checksum. The type is interface{}
    // with range: 0..4294967295.
    Cospfvirtlocallsdbchecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    Cospfvirtlocallsdbadvertisement interface{}
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetEntityData() *types.CommonEntityData {
    cospfvirtlocallsdbentry.EntityData.YFilter = cospfvirtlocallsdbentry.YFilter
    cospfvirtlocallsdbentry.EntityData.YangName = "cospfVirtLocalLsdbEntry"
    cospfvirtlocallsdbentry.EntityData.BundleName = "cisco_ios_xe"
    cospfvirtlocallsdbentry.EntityData.ParentYangName = "cospfVirtLocalLsdbTable"
    cospfvirtlocallsdbentry.EntityData.SegmentPath = "cospfVirtLocalLsdbEntry" + "[cospfVirtLocalLsdbTransitArea='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtransitarea) + "']" + "[cospfVirtLocalLsdbNeighbor='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbneighbor) + "']" + "[cospfVirtLocalLsdbType='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtype) + "']" + "[cospfVirtLocalLsdbLsid='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdblsid) + "']" + "[cospfVirtLocalLsdbRouterId='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbrouterid) + "']"
    cospfvirtlocallsdbentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfvirtlocallsdbentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfvirtlocallsdbentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfvirtlocallsdbentry.EntityData.Children = make(map[string]types.YChild)
    cospfvirtlocallsdbentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbTransitArea"] = types.YLeaf{"Cospfvirtlocallsdbtransitarea", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtransitarea}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbNeighbor"] = types.YLeaf{"Cospfvirtlocallsdbneighbor", cospfvirtlocallsdbentry.Cospfvirtlocallsdbneighbor}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbType"] = types.YLeaf{"Cospfvirtlocallsdbtype", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtype}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbLsid"] = types.YLeaf{"Cospfvirtlocallsdblsid", cospfvirtlocallsdbentry.Cospfvirtlocallsdblsid}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbRouterId"] = types.YLeaf{"Cospfvirtlocallsdbrouterid", cospfvirtlocallsdbentry.Cospfvirtlocallsdbrouterid}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbSequence"] = types.YLeaf{"Cospfvirtlocallsdbsequence", cospfvirtlocallsdbentry.Cospfvirtlocallsdbsequence}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbAge"] = types.YLeaf{"Cospfvirtlocallsdbage", cospfvirtlocallsdbentry.Cospfvirtlocallsdbage}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbChecksum"] = types.YLeaf{"Cospfvirtlocallsdbchecksum", cospfvirtlocallsdbentry.Cospfvirtlocallsdbchecksum}
    cospfvirtlocallsdbentry.EntityData.Leafs["cospfVirtLocalLsdbAdvertisement"] = types.YLeaf{"Cospfvirtlocallsdbadvertisement", cospfvirtlocallsdbentry.Cospfvirtlocallsdbadvertisement}
    return &(cospfvirtlocallsdbentry.EntityData)
}

// CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype represents Each  link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype string

const (
    CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype_localOpaqueLink CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype = "localOpaqueLink"
)

// CISCOOSPFMIB_Cospfshamlinknbrtable
// A table of sham link neighbor information.
type CISCOOSPFMIB_Cospfshamlinknbrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sham link neighbor information. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry.
    Cospfshamlinknbrentry []CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetEntityData() *types.CommonEntityData {
    cospfshamlinknbrtable.EntityData.YFilter = cospfshamlinknbrtable.YFilter
    cospfshamlinknbrtable.EntityData.YangName = "cospfShamLinkNbrTable"
    cospfshamlinknbrtable.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinknbrtable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfshamlinknbrtable.EntityData.SegmentPath = "cospfShamLinkNbrTable"
    cospfshamlinknbrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinknbrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinknbrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinknbrtable.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinknbrtable.EntityData.Children["cospfShamLinkNbrEntry"] = types.YChild{"Cospfshamlinknbrentry", nil}
    for i := range cospfshamlinknbrtable.Cospfshamlinknbrentry {
        cospfshamlinknbrtable.EntityData.Children[types.GetSegmentPath(&cospfshamlinknbrtable.Cospfshamlinknbrentry[i])] = types.YChild{"Cospfshamlinknbrentry", &cospfshamlinknbrtable.Cospfshamlinknbrentry[i]}
    }
    cospfshamlinknbrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospfshamlinknbrtable.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry
// Sham link neighbor information.
type CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinkslocalipaddrtype
    Cospfshamlinkslocalipaddrtype interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinkslocalipaddr
    Cospfshamlinkslocalipaddr interface{}

    // This attribute is a key. The area to which the sham link is part of. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinknbrarea interface{}

    // This attribute is a key. The type of internet address of this sham link
    // neighbor's IP address. The type is InetAddressType.
    Cospfshamlinknbripaddrtype interface{}

    // This attribute is a key. The IP address this sham link neighbor is using.
    // The type is string with length: 0..255.
    Cospfshamlinknbripaddr interface{}

    // A 32-bit integer uniquely identifying the neighboring router. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinknbrrtrid interface{}

    // A Bit Mask corresponding to the neighbor's options field.  Bit 1, if set,
    // indicates that the  system  will operate  on  Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will  ignore  all metrics except the TOS
    // 0 metric.  Bit 2, if set, indicates  that  the  system  is Network 
    // Multicast  capable; ie, that it implements  OSPF Multicast Routing. The
    // type is interface{} with range: 0..255.
    Cospfshamlinknbroptions interface{}

    // The state of this sham link neighbor relation- ship. The type is
    // Cospfshamlinknbrstate.
    Cospfshamlinknbrstate interface{}

    // The number of  times  this sham link has changed state or an error has
    // occurred. The type is interface{} with range: 0..4294967295.
    Cospfshamlinknbrevents interface{}

    // The  current  length  of  the   retransmission queue. The retransmission
    // queue is maintained for LSAs that have been flooded but not acknowledged on
    // this adjacency. The type is interface{} with range: 0..4294967295.
    Cospfshamlinknbrlsretransqlen interface{}

    // Indicates whether Hellos are being  suppressed to the neighbor. The type is
    // bool.
    Cospfshamlinknbrhellosuppressed interface{}
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetEntityData() *types.CommonEntityData {
    cospfshamlinknbrentry.EntityData.YFilter = cospfshamlinknbrentry.YFilter
    cospfshamlinknbrentry.EntityData.YangName = "cospfShamLinkNbrEntry"
    cospfshamlinknbrentry.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinknbrentry.EntityData.ParentYangName = "cospfShamLinkNbrTable"
    cospfshamlinknbrentry.EntityData.SegmentPath = "cospfShamLinkNbrEntry" + "[cospfShamLinksLocalIpAddrType='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinkslocalipaddrtype) + "']" + "[cospfShamLinksLocalIpAddr='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinkslocalipaddr) + "']" + "[cospfShamLinkNbrArea='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbrarea) + "']" + "[cospfShamLinkNbrIpAddrType='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbripaddrtype) + "']" + "[cospfShamLinkNbrIpAddr='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbripaddr) + "']"
    cospfshamlinknbrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinknbrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinknbrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinknbrentry.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinknbrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinksLocalIpAddrType"] = types.YLeaf{"Cospfshamlinkslocalipaddrtype", cospfshamlinknbrentry.Cospfshamlinkslocalipaddrtype}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinksLocalIpAddr"] = types.YLeaf{"Cospfshamlinkslocalipaddr", cospfshamlinknbrentry.Cospfshamlinkslocalipaddr}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrArea"] = types.YLeaf{"Cospfshamlinknbrarea", cospfshamlinknbrentry.Cospfshamlinknbrarea}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrIpAddrType"] = types.YLeaf{"Cospfshamlinknbripaddrtype", cospfshamlinknbrentry.Cospfshamlinknbripaddrtype}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrIpAddr"] = types.YLeaf{"Cospfshamlinknbripaddr", cospfshamlinknbrentry.Cospfshamlinknbripaddr}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrRtrId"] = types.YLeaf{"Cospfshamlinknbrrtrid", cospfshamlinknbrentry.Cospfshamlinknbrrtrid}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrOptions"] = types.YLeaf{"Cospfshamlinknbroptions", cospfshamlinknbrentry.Cospfshamlinknbroptions}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrState"] = types.YLeaf{"Cospfshamlinknbrstate", cospfshamlinknbrentry.Cospfshamlinknbrstate}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrEvents"] = types.YLeaf{"Cospfshamlinknbrevents", cospfshamlinknbrentry.Cospfshamlinknbrevents}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrLsRetransQLen"] = types.YLeaf{"Cospfshamlinknbrlsretransqlen", cospfshamlinknbrentry.Cospfshamlinknbrlsretransqlen}
    cospfshamlinknbrentry.EntityData.Leafs["cospfShamLinkNbrHelloSuppressed"] = types.YLeaf{"Cospfshamlinknbrhellosuppressed", cospfshamlinknbrentry.Cospfshamlinknbrhellosuppressed}
    return &(cospfshamlinknbrentry.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate represents ship.
type CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate string

const (
    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_down CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "down"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_attempt CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "attempt"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_init CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "init"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_twoWay CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "twoWay"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_exchangeStart CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "exchangeStart"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_exchange CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "exchange"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_loading CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "loading"

    CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate_full CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry_Cospfshamlinknbrstate = "full"
)

// CISCOOSPFMIB_Cospfshamlinkstable
// Information about this router's sham links.
type CISCOOSPFMIB_Cospfshamlinkstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry.
    Cospfshamlinksentry []CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetEntityData() *types.CommonEntityData {
    cospfshamlinkstable.EntityData.YFilter = cospfshamlinkstable.YFilter
    cospfshamlinkstable.EntityData.YangName = "cospfShamLinksTable"
    cospfshamlinkstable.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinkstable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfshamlinkstable.EntityData.SegmentPath = "cospfShamLinksTable"
    cospfshamlinkstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinkstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinkstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinkstable.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinkstable.EntityData.Children["cospfShamLinksEntry"] = types.YChild{"Cospfshamlinksentry", nil}
    for i := range cospfshamlinkstable.Cospfshamlinksentry {
        cospfshamlinkstable.EntityData.Children[types.GetSegmentPath(&cospfshamlinkstable.Cospfshamlinksentry[i])] = types.YChild{"Cospfshamlinksentry", &cospfshamlinkstable.Cospfshamlinksentry[i]}
    }
    cospfshamlinkstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cospfshamlinkstable.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry
// Information about a single sham link.
type CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The area that this sham link is part of. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Cospfshamlinksareaid interface{}

    // This attribute is a key. The type of internet address of this sham link's
    // local IP address. The type is InetAddressType.
    Cospfshamlinkslocalipaddrtype interface{}

    // This attribute is a key. The Local IP address of the sham link. The type is
    // string with length: 0..255.
    Cospfshamlinkslocalipaddr interface{}

    // This attribute is a key. The type of internet address of this sham link's
    // remote IP address. The type is InetAddressType.
    Cospfshamlinksremoteipaddrtype interface{}

    // This attribute is a key. The IP address of the other end router of the sham
    // link. The type is string with length: 0..255.
    Cospfshamlinksremoteipaddr interface{}

    // The number of seconds between  link-state  advertisement retransmissions,
    // for adjacencies belonging to this link. This value is also used when
    // retransmitting database  description and link-state request packets. This
    // value should be well over the expected round trip time. The type is
    // interface{} with range: 0..3600.
    Cospfshamlinksretransinterval interface{}

    // The length of time, in  seconds,  between  the Hello  packets that the
    // router sends on the sham link. The type is interface{} with range:
    // 1..65535.
    Cospfshamlinkshellointerval interface{}

    // The number of seconds that  a  router's  Hello packets  have  not been seen
    // before it's neighbors declare the router down.  This  should  be some 
    // multiple  of  the  Hello  interval. The type is interface{} with range:
    // 0..2147483647.
    Cospfshamlinksrtrdeadinterval interface{}

    // OSPF sham link states. The type is Cospfshamlinksstate.
    Cospfshamlinksstate interface{}

    // The number of state changes or error events on this sham link. The type is
    // interface{} with range: 0..4294967295.
    Cospfshamlinksevents interface{}

    // The Metric to be advertised. The type is interface{} with range: 0..65535.
    Cospfshamlinksmetric interface{}
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetEntityData() *types.CommonEntityData {
    cospfshamlinksentry.EntityData.YFilter = cospfshamlinksentry.YFilter
    cospfshamlinksentry.EntityData.YangName = "cospfShamLinksEntry"
    cospfshamlinksentry.EntityData.BundleName = "cisco_ios_xe"
    cospfshamlinksentry.EntityData.ParentYangName = "cospfShamLinksTable"
    cospfshamlinksentry.EntityData.SegmentPath = "cospfShamLinksEntry" + "[cospfShamLinksAreaId='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksareaid) + "']" + "[cospfShamLinksLocalIpAddrType='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinkslocalipaddrtype) + "']" + "[cospfShamLinksLocalIpAddr='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinkslocalipaddr) + "']" + "[cospfShamLinksRemoteIpAddrType='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksremoteipaddrtype) + "']" + "[cospfShamLinksRemoteIpAddr='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksremoteipaddr) + "']"
    cospfshamlinksentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfshamlinksentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfshamlinksentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfshamlinksentry.EntityData.Children = make(map[string]types.YChild)
    cospfshamlinksentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksAreaId"] = types.YLeaf{"Cospfshamlinksareaid", cospfshamlinksentry.Cospfshamlinksareaid}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksLocalIpAddrType"] = types.YLeaf{"Cospfshamlinkslocalipaddrtype", cospfshamlinksentry.Cospfshamlinkslocalipaddrtype}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksLocalIpAddr"] = types.YLeaf{"Cospfshamlinkslocalipaddr", cospfshamlinksentry.Cospfshamlinkslocalipaddr}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksRemoteIpAddrType"] = types.YLeaf{"Cospfshamlinksremoteipaddrtype", cospfshamlinksentry.Cospfshamlinksremoteipaddrtype}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksRemoteIpAddr"] = types.YLeaf{"Cospfshamlinksremoteipaddr", cospfshamlinksentry.Cospfshamlinksremoteipaddr}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksRetransInterval"] = types.YLeaf{"Cospfshamlinksretransinterval", cospfshamlinksentry.Cospfshamlinksretransinterval}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksHelloInterval"] = types.YLeaf{"Cospfshamlinkshellointerval", cospfshamlinksentry.Cospfshamlinkshellointerval}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksRtrDeadInterval"] = types.YLeaf{"Cospfshamlinksrtrdeadinterval", cospfshamlinksentry.Cospfshamlinksrtrdeadinterval}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksState"] = types.YLeaf{"Cospfshamlinksstate", cospfshamlinksentry.Cospfshamlinksstate}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksEvents"] = types.YLeaf{"Cospfshamlinksevents", cospfshamlinksentry.Cospfshamlinksevents}
    cospfshamlinksentry.EntityData.Leafs["cospfShamLinksMetric"] = types.YLeaf{"Cospfshamlinksmetric", cospfshamlinksentry.Cospfshamlinksmetric}
    return &(cospfshamlinksentry.EntityData)
}

// CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate represents OSPF sham link states.
type CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate string

const (
    CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate_down CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate = "down"

    CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate_pointToPoint CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate = "pointToPoint"
)

