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

    
    CospfGeneralGroup CISCOOSPFMIB_CospfGeneralGroup

    // The OSPF Process's Link State Database. This  table is meant for Opaque
    // LSA's.
    CospfLsdbTable CISCOOSPFMIB_CospfLsdbTable

    // Information about this router's sham links.
    CospfShamLinkTable CISCOOSPFMIB_CospfShamLinkTable

    // The OSPF Process's Link-Local Link State Database for non-virtual links.
    CospfLocalLsdbTable CISCOOSPFMIB_CospfLocalLsdbTable

    // The OSPF Process's Link-Local Link State Database for virtual links.
    CospfVirtLocalLsdbTable CISCOOSPFMIB_CospfVirtLocalLsdbTable

    // A table of sham link neighbor information.
    CospfShamLinkNbrTable CISCOOSPFMIB_CospfShamLinkNbrTable

    // Information about this router's sham links.
    CospfShamLinksTable CISCOOSPFMIB_CospfShamLinksTable
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetEntityData() *types.CommonEntityData {
    cISCOOSPFMIB.EntityData.YFilter = cISCOOSPFMIB.YFilter
    cISCOOSPFMIB.EntityData.YangName = "CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOOSPFMIB.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.SegmentPath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB"
    cISCOOSPFMIB.EntityData.AbsolutePath = cISCOOSPFMIB.EntityData.SegmentPath
    cISCOOSPFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOOSPFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOOSPFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOOSPFMIB.EntityData.Children = types.NewOrderedMap()
    cISCOOSPFMIB.EntityData.Children.Append("cospfGeneralGroup", types.YChild{"CospfGeneralGroup", &cISCOOSPFMIB.CospfGeneralGroup})
    cISCOOSPFMIB.EntityData.Children.Append("cospfLsdbTable", types.YChild{"CospfLsdbTable", &cISCOOSPFMIB.CospfLsdbTable})
    cISCOOSPFMIB.EntityData.Children.Append("cospfShamLinkTable", types.YChild{"CospfShamLinkTable", &cISCOOSPFMIB.CospfShamLinkTable})
    cISCOOSPFMIB.EntityData.Children.Append("cospfLocalLsdbTable", types.YChild{"CospfLocalLsdbTable", &cISCOOSPFMIB.CospfLocalLsdbTable})
    cISCOOSPFMIB.EntityData.Children.Append("cospfVirtLocalLsdbTable", types.YChild{"CospfVirtLocalLsdbTable", &cISCOOSPFMIB.CospfVirtLocalLsdbTable})
    cISCOOSPFMIB.EntityData.Children.Append("cospfShamLinkNbrTable", types.YChild{"CospfShamLinkNbrTable", &cISCOOSPFMIB.CospfShamLinkNbrTable})
    cISCOOSPFMIB.EntityData.Children.Append("cospfShamLinksTable", types.YChild{"CospfShamLinksTable", &cISCOOSPFMIB.CospfShamLinksTable})
    cISCOOSPFMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOOSPFMIB.EntityData.YListKeys = []string {}

    return &(cISCOOSPFMIB.EntityData)
}

// CISCOOSPFMIB_CospfGeneralGroup
type CISCOOSPFMIB_CospfGeneralGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates metrics used to choose among multiple AS- external-LSAs. When
    // cospfRFC1583Compatibility is set to enabled, only cost will be used when
    // choosing among multiple AS-external-LSAs advertising the same destination.
    // When cospfRFC1583Compatibility is set to disabled, preference will be
    // driven first by type of path using cost only to break ties. The type is
    // bool.
    CospfRFC1583Compatibility interface{}

    // The router's support for Opaque LSA types. The type is bool.
    CospfOpaqueLsaSupport interface{}

    // The router's support for OSPF traffic engineering. The type is bool.
    CospfTrafficEngineeringSupport interface{}

    // The total number of Opaque AS link-state advertisements in the link state
    // database. The type is interface{} with range: 0..4294967295.
    CospfOpaqueASLsaCount interface{}

    // The 32-bit unsigned sum of the Opaque AS  link-state advertisements' LS
    // checksums contained link state database. The type is interface{} with
    // range: 0..4294967295.
    CospfOpaqueASLsaCksumSum interface{}
}

func (cospfGeneralGroup *CISCOOSPFMIB_CospfGeneralGroup) GetEntityData() *types.CommonEntityData {
    cospfGeneralGroup.EntityData.YFilter = cospfGeneralGroup.YFilter
    cospfGeneralGroup.EntityData.YangName = "cospfGeneralGroup"
    cospfGeneralGroup.EntityData.BundleName = "cisco_ios_xe"
    cospfGeneralGroup.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfGeneralGroup.EntityData.SegmentPath = "cospfGeneralGroup"
    cospfGeneralGroup.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfGeneralGroup.EntityData.SegmentPath
    cospfGeneralGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfGeneralGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfGeneralGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfGeneralGroup.EntityData.Children = types.NewOrderedMap()
    cospfGeneralGroup.EntityData.Leafs = types.NewOrderedMap()
    cospfGeneralGroup.EntityData.Leafs.Append("cospfRFC1583Compatibility", types.YLeaf{"CospfRFC1583Compatibility", cospfGeneralGroup.CospfRFC1583Compatibility})
    cospfGeneralGroup.EntityData.Leafs.Append("cospfOpaqueLsaSupport", types.YLeaf{"CospfOpaqueLsaSupport", cospfGeneralGroup.CospfOpaqueLsaSupport})
    cospfGeneralGroup.EntityData.Leafs.Append("cospfTrafficEngineeringSupport", types.YLeaf{"CospfTrafficEngineeringSupport", cospfGeneralGroup.CospfTrafficEngineeringSupport})
    cospfGeneralGroup.EntityData.Leafs.Append("cospfOpaqueASLsaCount", types.YLeaf{"CospfOpaqueASLsaCount", cospfGeneralGroup.CospfOpaqueASLsaCount})
    cospfGeneralGroup.EntityData.Leafs.Append("cospfOpaqueASLsaCksumSum", types.YLeaf{"CospfOpaqueASLsaCksumSum", cospfGeneralGroup.CospfOpaqueASLsaCksumSum})

    cospfGeneralGroup.EntityData.YListKeys = []string {}

    return &(cospfGeneralGroup.EntityData)
}

// CISCOOSPFMIB_CospfLsdbTable
// The OSPF Process's Link State Database. This 
// table is meant for Opaque LSA's
type CISCOOSPFMIB_CospfLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry.
    CospfLsdbEntry []*CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry
}

func (cospfLsdbTable *CISCOOSPFMIB_CospfLsdbTable) GetEntityData() *types.CommonEntityData {
    cospfLsdbTable.EntityData.YFilter = cospfLsdbTable.YFilter
    cospfLsdbTable.EntityData.YangName = "cospfLsdbTable"
    cospfLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    cospfLsdbTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfLsdbTable.EntityData.SegmentPath = "cospfLsdbTable"
    cospfLsdbTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfLsdbTable.EntityData.SegmentPath
    cospfLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfLsdbTable.EntityData.Children = types.NewOrderedMap()
    cospfLsdbTable.EntityData.Children.Append("cospfLsdbEntry", types.YChild{"CospfLsdbEntry", nil})
    for i := range cospfLsdbTable.CospfLsdbEntry {
        cospfLsdbTable.EntityData.Children.Append(types.GetSegmentPath(cospfLsdbTable.CospfLsdbEntry[i]), types.YChild{"CospfLsdbEntry", cospfLsdbTable.CospfLsdbEntry[i]})
    }
    cospfLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    cospfLsdbTable.EntityData.YListKeys = []string {}

    return &(cospfLsdbTable.EntityData)
}

// CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry
// A single Link State Advertisement.
type CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbAreaId
    OspfLsdbAreaId interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // CospfLsdbType.
    CospfLsdbType interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbLsid
    OspfLsdbLsid interface{}

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // Refers to ospf_mib.OSPFMIB_OspfLsdbTable_OspfLsdbEntry_OspfLsdbRouterId
    OspfLsdbRouterId interface{}

    // The sequence number field is a  signed  32-bit integer.   It  is used to
    // detect old and duplicate link state advertisements.  The  space  of
    // sequence  numbers  is  linearly  ordered.   The larger the sequence number
    // the more recent  the advertisement. The type is interface{} with range:
    // 1..147483647.
    CospfLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: 0..2147483647.
    CospfLsdbAge interface{}

    // This field is the  checksum  of  the  complete contents  of  the 
    // advertisement, excepting the age field.  The age field is excepted  so 
    // that an   advertisement's  age  can  be  incremented without updating the 
    // checksum.   The  checksum used  is  the same that is used for ISO
    // connectionless datagrams; it is commonly referred  to as the Fletcher
    // checksum. The type is interface{} with range: 0..2147483647.
    CospfLsdbChecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    CospfLsdbAdvertisement interface{}
}

func (cospfLsdbEntry *CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry) GetEntityData() *types.CommonEntityData {
    cospfLsdbEntry.EntityData.YFilter = cospfLsdbEntry.YFilter
    cospfLsdbEntry.EntityData.YangName = "cospfLsdbEntry"
    cospfLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfLsdbEntry.EntityData.ParentYangName = "cospfLsdbTable"
    cospfLsdbEntry.EntityData.SegmentPath = "cospfLsdbEntry" + types.AddKeyToken(cospfLsdbEntry.OspfLsdbAreaId, "ospfLsdbAreaId") + types.AddKeyToken(cospfLsdbEntry.CospfLsdbType, "cospfLsdbType") + types.AddKeyToken(cospfLsdbEntry.OspfLsdbLsid, "ospfLsdbLsid") + types.AddKeyToken(cospfLsdbEntry.OspfLsdbRouterId, "ospfLsdbRouterId")
    cospfLsdbEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfLsdbTable/" + cospfLsdbEntry.EntityData.SegmentPath
    cospfLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfLsdbEntry.EntityData.Children = types.NewOrderedMap()
    cospfLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbAreaId", types.YLeaf{"OspfLsdbAreaId", cospfLsdbEntry.OspfLsdbAreaId})
    cospfLsdbEntry.EntityData.Leafs.Append("cospfLsdbType", types.YLeaf{"CospfLsdbType", cospfLsdbEntry.CospfLsdbType})
    cospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbLsid", types.YLeaf{"OspfLsdbLsid", cospfLsdbEntry.OspfLsdbLsid})
    cospfLsdbEntry.EntityData.Leafs.Append("ospfLsdbRouterId", types.YLeaf{"OspfLsdbRouterId", cospfLsdbEntry.OspfLsdbRouterId})
    cospfLsdbEntry.EntityData.Leafs.Append("cospfLsdbSequence", types.YLeaf{"CospfLsdbSequence", cospfLsdbEntry.CospfLsdbSequence})
    cospfLsdbEntry.EntityData.Leafs.Append("cospfLsdbAge", types.YLeaf{"CospfLsdbAge", cospfLsdbEntry.CospfLsdbAge})
    cospfLsdbEntry.EntityData.Leafs.Append("cospfLsdbChecksum", types.YLeaf{"CospfLsdbChecksum", cospfLsdbEntry.CospfLsdbChecksum})
    cospfLsdbEntry.EntityData.Leafs.Append("cospfLsdbAdvertisement", types.YLeaf{"CospfLsdbAdvertisement", cospfLsdbEntry.CospfLsdbAdvertisement})

    cospfLsdbEntry.EntityData.YListKeys = []string {"OspfLsdbAreaId", "CospfLsdbType", "OspfLsdbLsid", "OspfLsdbRouterId"}

    return &(cospfLsdbEntry.EntityData)
}

// CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType string

const (
    CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType_areaOpaqueLink CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType = "areaOpaqueLink"

    CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType_asOpaqueLink CISCOOSPFMIB_CospfLsdbTable_CospfLsdbEntry_CospfLsdbType = "asOpaqueLink"
)

// CISCOOSPFMIB_CospfShamLinkTable
// Information about this router's sham links
type CISCOOSPFMIB_CospfShamLinkTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry.
    CospfShamLinkEntry []*CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry
}

func (cospfShamLinkTable *CISCOOSPFMIB_CospfShamLinkTable) GetEntityData() *types.CommonEntityData {
    cospfShamLinkTable.EntityData.YFilter = cospfShamLinkTable.YFilter
    cospfShamLinkTable.EntityData.YangName = "cospfShamLinkTable"
    cospfShamLinkTable.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinkTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfShamLinkTable.EntityData.SegmentPath = "cospfShamLinkTable"
    cospfShamLinkTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfShamLinkTable.EntityData.SegmentPath
    cospfShamLinkTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinkTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinkTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinkTable.EntityData.Children = types.NewOrderedMap()
    cospfShamLinkTable.EntityData.Children.Append("cospfShamLinkEntry", types.YChild{"CospfShamLinkEntry", nil})
    for i := range cospfShamLinkTable.CospfShamLinkEntry {
        cospfShamLinkTable.EntityData.Children.Append(types.GetSegmentPath(cospfShamLinkTable.CospfShamLinkEntry[i]), types.YChild{"CospfShamLinkEntry", cospfShamLinkTable.CospfShamLinkEntry[i]})
    }
    cospfShamLinkTable.EntityData.Leafs = types.NewOrderedMap()

    cospfShamLinkTable.EntityData.YListKeys = []string {}

    return &(cospfShamLinkTable.EntityData)
}

// CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry
// Information about a single sham link
type CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The  Transit  Area  that  the   Virtual   Link
    // traverses.  By definition, this is not 0.0.0.0. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinkAreaId interface{}

    // This attribute is a key. The Local IP address of the sham link. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinkLocalIpAddress interface{}

    // This attribute is a key. The Router ID of the other end router of the sham
    // link. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinkNeighborId interface{}

    // The number of seconds between  link-state  advertisement retransmissions, 
    // for  adjacencies belonging to this  link.   This  value  is also used when
    // retransmitting database description   and  link-state  request  packets.
    // This value   should  be well over the expected round trip time. The type is
    // interface{} with range: 0..3600.
    CospfShamLinkRetransInterval interface{}

    // The length of time, in  seconds,  between  the Hello  packets that the
    // router sends on the sham link. The type is interface{} with range:
    // 1..65535.
    CospfShamLinkHelloInterval interface{}

    // The number of seconds that  a  router's  Hello packets  have  not been seen
    // before it's neighbors declare the router down.  This  should  be some 
    // multiple  of  the  Hello  interval. The type is interface{} with range:
    // 0..2147483647.
    CospfShamLinkRtrDeadInterval interface{}

    // OSPF sham link states. The type is CospfShamLinkState.
    CospfShamLinkState interface{}

    // The number of state changes or error events on this sham link. The type is
    // interface{} with range: 0..4294967295.
    CospfShamLinkEvents interface{}

    // The Metric to be advertised. The type is interface{} with range: 0..65535.
    CospfShamLinkMetric interface{}
}

func (cospfShamLinkEntry *CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry) GetEntityData() *types.CommonEntityData {
    cospfShamLinkEntry.EntityData.YFilter = cospfShamLinkEntry.YFilter
    cospfShamLinkEntry.EntityData.YangName = "cospfShamLinkEntry"
    cospfShamLinkEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinkEntry.EntityData.ParentYangName = "cospfShamLinkTable"
    cospfShamLinkEntry.EntityData.SegmentPath = "cospfShamLinkEntry" + types.AddKeyToken(cospfShamLinkEntry.CospfShamLinkAreaId, "cospfShamLinkAreaId") + types.AddKeyToken(cospfShamLinkEntry.CospfShamLinkLocalIpAddress, "cospfShamLinkLocalIpAddress") + types.AddKeyToken(cospfShamLinkEntry.CospfShamLinkNeighborId, "cospfShamLinkNeighborId")
    cospfShamLinkEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfShamLinkTable/" + cospfShamLinkEntry.EntityData.SegmentPath
    cospfShamLinkEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinkEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinkEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinkEntry.EntityData.Children = types.NewOrderedMap()
    cospfShamLinkEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkAreaId", types.YLeaf{"CospfShamLinkAreaId", cospfShamLinkEntry.CospfShamLinkAreaId})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkLocalIpAddress", types.YLeaf{"CospfShamLinkLocalIpAddress", cospfShamLinkEntry.CospfShamLinkLocalIpAddress})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkNeighborId", types.YLeaf{"CospfShamLinkNeighborId", cospfShamLinkEntry.CospfShamLinkNeighborId})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkRetransInterval", types.YLeaf{"CospfShamLinkRetransInterval", cospfShamLinkEntry.CospfShamLinkRetransInterval})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkHelloInterval", types.YLeaf{"CospfShamLinkHelloInterval", cospfShamLinkEntry.CospfShamLinkHelloInterval})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkRtrDeadInterval", types.YLeaf{"CospfShamLinkRtrDeadInterval", cospfShamLinkEntry.CospfShamLinkRtrDeadInterval})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkState", types.YLeaf{"CospfShamLinkState", cospfShamLinkEntry.CospfShamLinkState})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkEvents", types.YLeaf{"CospfShamLinkEvents", cospfShamLinkEntry.CospfShamLinkEvents})
    cospfShamLinkEntry.EntityData.Leafs.Append("cospfShamLinkMetric", types.YLeaf{"CospfShamLinkMetric", cospfShamLinkEntry.CospfShamLinkMetric})

    cospfShamLinkEntry.EntityData.YListKeys = []string {"CospfShamLinkAreaId", "CospfShamLinkLocalIpAddress", "CospfShamLinkNeighborId"}

    return &(cospfShamLinkEntry.EntityData)
}

// CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState represents OSPF sham link states.
type CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState string

const (
    CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState_down CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState = "down"

    CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState_pointToPoint CISCOOSPFMIB_CospfShamLinkTable_CospfShamLinkEntry_CospfShamLinkState = "pointToPoint"
)

// CISCOOSPFMIB_CospfLocalLsdbTable
// The OSPF Process's Link-Local Link State Database
// for non-virtual links.
type CISCOOSPFMIB_CospfLocalLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry.
    CospfLocalLsdbEntry []*CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry
}

func (cospfLocalLsdbTable *CISCOOSPFMIB_CospfLocalLsdbTable) GetEntityData() *types.CommonEntityData {
    cospfLocalLsdbTable.EntityData.YFilter = cospfLocalLsdbTable.YFilter
    cospfLocalLsdbTable.EntityData.YangName = "cospfLocalLsdbTable"
    cospfLocalLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    cospfLocalLsdbTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfLocalLsdbTable.EntityData.SegmentPath = "cospfLocalLsdbTable"
    cospfLocalLsdbTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfLocalLsdbTable.EntityData.SegmentPath
    cospfLocalLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfLocalLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfLocalLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfLocalLsdbTable.EntityData.Children = types.NewOrderedMap()
    cospfLocalLsdbTable.EntityData.Children.Append("cospfLocalLsdbEntry", types.YChild{"CospfLocalLsdbEntry", nil})
    for i := range cospfLocalLsdbTable.CospfLocalLsdbEntry {
        cospfLocalLsdbTable.EntityData.Children.Append(types.GetSegmentPath(cospfLocalLsdbTable.CospfLocalLsdbEntry[i]), types.YChild{"CospfLocalLsdbEntry", cospfLocalLsdbTable.CospfLocalLsdbEntry[i]})
    }
    cospfLocalLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    cospfLocalLsdbTable.EntityData.YListKeys = []string {}

    return &(cospfLocalLsdbTable.EntityData)
}

// CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry
// A single Link State Advertisement.
type CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The IP Address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfLocalLsdbIpAddress interface{}

    // This attribute is a key. The Interface Index of the interface from which
    // the LSA was received if the interface is unnumbered. The type is
    // interface{} with range: 0..2147483647.
    CospfLocalLsdbAddressLessIf interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // CospfLocalLsdbType.
    CospfLocalLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32 bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfLocalLsdbLsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfLocalLsdbRouterId interface{}

    // The sequence number field is a signed 32-bit integer. It is used to detect
    // old and duplicate link state advertisements. The space of sequence numbers
    // is linearly ordered. The larger the sequence number the more recent the
    // advertisement. The type is interface{} with range: -2147483647..2147483647.
    CospfLocalLsdbSequence interface{}

    // This field is the age of the link state advertisement  in seconds. The type
    // is interface{} with range: 0..3600.
    CospfLocalLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field. The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum. The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred  to as the Fletcher checksum. The type is interface{}
    // with range: 0..4294967295.
    CospfLocalLsdbChecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    CospfLocalLsdbAdvertisement interface{}
}

func (cospfLocalLsdbEntry *CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry) GetEntityData() *types.CommonEntityData {
    cospfLocalLsdbEntry.EntityData.YFilter = cospfLocalLsdbEntry.YFilter
    cospfLocalLsdbEntry.EntityData.YangName = "cospfLocalLsdbEntry"
    cospfLocalLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfLocalLsdbEntry.EntityData.ParentYangName = "cospfLocalLsdbTable"
    cospfLocalLsdbEntry.EntityData.SegmentPath = "cospfLocalLsdbEntry" + types.AddKeyToken(cospfLocalLsdbEntry.CospfLocalLsdbIpAddress, "cospfLocalLsdbIpAddress") + types.AddKeyToken(cospfLocalLsdbEntry.CospfLocalLsdbAddressLessIf, "cospfLocalLsdbAddressLessIf") + types.AddKeyToken(cospfLocalLsdbEntry.CospfLocalLsdbType, "cospfLocalLsdbType") + types.AddKeyToken(cospfLocalLsdbEntry.CospfLocalLsdbLsid, "cospfLocalLsdbLsid") + types.AddKeyToken(cospfLocalLsdbEntry.CospfLocalLsdbRouterId, "cospfLocalLsdbRouterId")
    cospfLocalLsdbEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfLocalLsdbTable/" + cospfLocalLsdbEntry.EntityData.SegmentPath
    cospfLocalLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfLocalLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfLocalLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfLocalLsdbEntry.EntityData.Children = types.NewOrderedMap()
    cospfLocalLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbIpAddress", types.YLeaf{"CospfLocalLsdbIpAddress", cospfLocalLsdbEntry.CospfLocalLsdbIpAddress})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbAddressLessIf", types.YLeaf{"CospfLocalLsdbAddressLessIf", cospfLocalLsdbEntry.CospfLocalLsdbAddressLessIf})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbType", types.YLeaf{"CospfLocalLsdbType", cospfLocalLsdbEntry.CospfLocalLsdbType})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbLsid", types.YLeaf{"CospfLocalLsdbLsid", cospfLocalLsdbEntry.CospfLocalLsdbLsid})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbRouterId", types.YLeaf{"CospfLocalLsdbRouterId", cospfLocalLsdbEntry.CospfLocalLsdbRouterId})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbSequence", types.YLeaf{"CospfLocalLsdbSequence", cospfLocalLsdbEntry.CospfLocalLsdbSequence})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbAge", types.YLeaf{"CospfLocalLsdbAge", cospfLocalLsdbEntry.CospfLocalLsdbAge})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbChecksum", types.YLeaf{"CospfLocalLsdbChecksum", cospfLocalLsdbEntry.CospfLocalLsdbChecksum})
    cospfLocalLsdbEntry.EntityData.Leafs.Append("cospfLocalLsdbAdvertisement", types.YLeaf{"CospfLocalLsdbAdvertisement", cospfLocalLsdbEntry.CospfLocalLsdbAdvertisement})

    cospfLocalLsdbEntry.EntityData.YListKeys = []string {"CospfLocalLsdbIpAddress", "CospfLocalLsdbAddressLessIf", "CospfLocalLsdbType", "CospfLocalLsdbLsid", "CospfLocalLsdbRouterId"}

    return &(cospfLocalLsdbEntry.EntityData)
}

// CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry_CospfLocalLsdbType represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry_CospfLocalLsdbType string

const (
    CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry_CospfLocalLsdbType_localOpaqueLink CISCOOSPFMIB_CospfLocalLsdbTable_CospfLocalLsdbEntry_CospfLocalLsdbType = "localOpaqueLink"
)

// CISCOOSPFMIB_CospfVirtLocalLsdbTable
// The OSPF Process's Link-Local Link State Database
// for virtual links.
type CISCOOSPFMIB_CospfVirtLocalLsdbTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry.
    CospfVirtLocalLsdbEntry []*CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry
}

func (cospfVirtLocalLsdbTable *CISCOOSPFMIB_CospfVirtLocalLsdbTable) GetEntityData() *types.CommonEntityData {
    cospfVirtLocalLsdbTable.EntityData.YFilter = cospfVirtLocalLsdbTable.YFilter
    cospfVirtLocalLsdbTable.EntityData.YangName = "cospfVirtLocalLsdbTable"
    cospfVirtLocalLsdbTable.EntityData.BundleName = "cisco_ios_xe"
    cospfVirtLocalLsdbTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfVirtLocalLsdbTable.EntityData.SegmentPath = "cospfVirtLocalLsdbTable"
    cospfVirtLocalLsdbTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfVirtLocalLsdbTable.EntityData.SegmentPath
    cospfVirtLocalLsdbTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfVirtLocalLsdbTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfVirtLocalLsdbTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfVirtLocalLsdbTable.EntityData.Children = types.NewOrderedMap()
    cospfVirtLocalLsdbTable.EntityData.Children.Append("cospfVirtLocalLsdbEntry", types.YChild{"CospfVirtLocalLsdbEntry", nil})
    for i := range cospfVirtLocalLsdbTable.CospfVirtLocalLsdbEntry {
        cospfVirtLocalLsdbTable.EntityData.Children.Append(types.GetSegmentPath(cospfVirtLocalLsdbTable.CospfVirtLocalLsdbEntry[i]), types.YChild{"CospfVirtLocalLsdbEntry", cospfVirtLocalLsdbTable.CospfVirtLocalLsdbEntry[i]})
    }
    cospfVirtLocalLsdbTable.EntityData.Leafs = types.NewOrderedMap()

    cospfVirtLocalLsdbTable.EntityData.YListKeys = []string {}

    return &(cospfVirtLocalLsdbTable.EntityData)
}

// CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry
// A single Link State Advertisement.
type CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Transit Area that the Virtual Link traverses.
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfVirtLocalLsdbTransitArea interface{}

    // This attribute is a key. The Router ID of the Virtual Neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfVirtLocalLsdbNeighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each 
    // link state type has a separate advertisement format. The type is
    // CospfVirtLocalLsdbType.
    CospfVirtLocalLsdbType interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32 bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfVirtLocalLsdbLsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfVirtLocalLsdbRouterId interface{}

    // The sequence number field is a  signed  32-bit integer. It is used to
    // detect old and duplicate link state advertisements. The space of sequence
    // numbers is linearly ordered. The larger the sequence number the more recent
    // the advertisement. The type is interface{} with range:
    // -2147483647..2147483647.
    CospfVirtLocalLsdbSequence interface{}

    // This field is the age of the link state advertisement in seconds. The type
    // is interface{} with range: 0..3600.
    CospfVirtLocalLsdbAge interface{}

    // This field is the checksum of the complete contents of the advertisement,
    // excepting the age field. The age field is excepted so that an
    // advertisement's age can be incremented without updating the checksum. The
    // checksum used is the same that is used for ISO connectionless datagrams; it
    // is commonly referred  to as the Fletcher checksum. The type is interface{}
    // with range: 0..4294967295.
    CospfVirtLocalLsdbChecksum interface{}

    // The entire Link State Advertisement, including its header. The type is
    // string with length: 1..65535.
    CospfVirtLocalLsdbAdvertisement interface{}
}

func (cospfVirtLocalLsdbEntry *CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry) GetEntityData() *types.CommonEntityData {
    cospfVirtLocalLsdbEntry.EntityData.YFilter = cospfVirtLocalLsdbEntry.YFilter
    cospfVirtLocalLsdbEntry.EntityData.YangName = "cospfVirtLocalLsdbEntry"
    cospfVirtLocalLsdbEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfVirtLocalLsdbEntry.EntityData.ParentYangName = "cospfVirtLocalLsdbTable"
    cospfVirtLocalLsdbEntry.EntityData.SegmentPath = "cospfVirtLocalLsdbEntry" + types.AddKeyToken(cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbTransitArea, "cospfVirtLocalLsdbTransitArea") + types.AddKeyToken(cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbNeighbor, "cospfVirtLocalLsdbNeighbor") + types.AddKeyToken(cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbType, "cospfVirtLocalLsdbType") + types.AddKeyToken(cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbLsid, "cospfVirtLocalLsdbLsid") + types.AddKeyToken(cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbRouterId, "cospfVirtLocalLsdbRouterId")
    cospfVirtLocalLsdbEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfVirtLocalLsdbTable/" + cospfVirtLocalLsdbEntry.EntityData.SegmentPath
    cospfVirtLocalLsdbEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfVirtLocalLsdbEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfVirtLocalLsdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfVirtLocalLsdbEntry.EntityData.Children = types.NewOrderedMap()
    cospfVirtLocalLsdbEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbTransitArea", types.YLeaf{"CospfVirtLocalLsdbTransitArea", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbTransitArea})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbNeighbor", types.YLeaf{"CospfVirtLocalLsdbNeighbor", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbNeighbor})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbType", types.YLeaf{"CospfVirtLocalLsdbType", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbType})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbLsid", types.YLeaf{"CospfVirtLocalLsdbLsid", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbLsid})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbRouterId", types.YLeaf{"CospfVirtLocalLsdbRouterId", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbRouterId})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbSequence", types.YLeaf{"CospfVirtLocalLsdbSequence", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbSequence})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbAge", types.YLeaf{"CospfVirtLocalLsdbAge", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbAge})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbChecksum", types.YLeaf{"CospfVirtLocalLsdbChecksum", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbChecksum})
    cospfVirtLocalLsdbEntry.EntityData.Leafs.Append("cospfVirtLocalLsdbAdvertisement", types.YLeaf{"CospfVirtLocalLsdbAdvertisement", cospfVirtLocalLsdbEntry.CospfVirtLocalLsdbAdvertisement})

    cospfVirtLocalLsdbEntry.EntityData.YListKeys = []string {"CospfVirtLocalLsdbTransitArea", "CospfVirtLocalLsdbNeighbor", "CospfVirtLocalLsdbType", "CospfVirtLocalLsdbLsid", "CospfVirtLocalLsdbRouterId"}

    return &(cospfVirtLocalLsdbEntry.EntityData)
}

// CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry_CospfVirtLocalLsdbType represents Each  link state type has a separate advertisement format.
type CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry_CospfVirtLocalLsdbType string

const (
    CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry_CospfVirtLocalLsdbType_localOpaqueLink CISCOOSPFMIB_CospfVirtLocalLsdbTable_CospfVirtLocalLsdbEntry_CospfVirtLocalLsdbType = "localOpaqueLink"
)

// CISCOOSPFMIB_CospfShamLinkNbrTable
// A table of sham link neighbor information.
type CISCOOSPFMIB_CospfShamLinkNbrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sham link neighbor information. The type is slice of
    // CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry.
    CospfShamLinkNbrEntry []*CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry
}

func (cospfShamLinkNbrTable *CISCOOSPFMIB_CospfShamLinkNbrTable) GetEntityData() *types.CommonEntityData {
    cospfShamLinkNbrTable.EntityData.YFilter = cospfShamLinkNbrTable.YFilter
    cospfShamLinkNbrTable.EntityData.YangName = "cospfShamLinkNbrTable"
    cospfShamLinkNbrTable.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinkNbrTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfShamLinkNbrTable.EntityData.SegmentPath = "cospfShamLinkNbrTable"
    cospfShamLinkNbrTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfShamLinkNbrTable.EntityData.SegmentPath
    cospfShamLinkNbrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinkNbrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinkNbrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinkNbrTable.EntityData.Children = types.NewOrderedMap()
    cospfShamLinkNbrTable.EntityData.Children.Append("cospfShamLinkNbrEntry", types.YChild{"CospfShamLinkNbrEntry", nil})
    for i := range cospfShamLinkNbrTable.CospfShamLinkNbrEntry {
        cospfShamLinkNbrTable.EntityData.Children.Append(types.GetSegmentPath(cospfShamLinkNbrTable.CospfShamLinkNbrEntry[i]), types.YChild{"CospfShamLinkNbrEntry", cospfShamLinkNbrTable.CospfShamLinkNbrEntry[i]})
    }
    cospfShamLinkNbrTable.EntityData.Leafs = types.NewOrderedMap()

    cospfShamLinkNbrTable.EntityData.YListKeys = []string {}

    return &(cospfShamLinkNbrTable.EntityData)
}

// CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry
// Sham link neighbor information.
type CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksLocalIpAddrType
    CospfShamLinksLocalIpAddrType interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksLocalIpAddr
    CospfShamLinksLocalIpAddr interface{}

    // This attribute is a key. The area to which the sham link is part of. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinkNbrArea interface{}

    // This attribute is a key. The type of internet address of this sham link
    // neighbor's IP address. The type is InetAddressType.
    CospfShamLinkNbrIpAddrType interface{}

    // This attribute is a key. The IP address this sham link neighbor is using.
    // The type is string with length: 0..255.
    CospfShamLinkNbrIpAddr interface{}

    // A 32-bit integer uniquely identifying the neighboring router. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinkNbrRtrId interface{}

    // A Bit Mask corresponding to the neighbor's options field.  Bit 1, if set,
    // indicates that the  system  will operate  on  Type of Service metrics other
    // than TOS 0.  If zero, the neighbor will  ignore  all metrics except the TOS
    // 0 metric.  Bit 2, if set, indicates  that  the  system  is Network 
    // Multicast  capable; ie, that it implements  OSPF Multicast Routing. The
    // type is interface{} with range: 0..255.
    CospfShamLinkNbrOptions interface{}

    // The state of this sham link neighbor relation- ship. The type is
    // CospfShamLinkNbrState.
    CospfShamLinkNbrState interface{}

    // The number of  times  this sham link has changed state or an error has
    // occurred. The type is interface{} with range: 0..4294967295.
    CospfShamLinkNbrEvents interface{}

    // The  current  length  of  the   retransmission queue. The retransmission
    // queue is maintained for LSAs that have been flooded but not acknowledged on
    // this adjacency. The type is interface{} with range: 0..4294967295.
    CospfShamLinkNbrLsRetransQLen interface{}

    // Indicates whether Hellos are being  suppressed to the neighbor. The type is
    // bool.
    CospfShamLinkNbrHelloSuppressed interface{}
}

func (cospfShamLinkNbrEntry *CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry) GetEntityData() *types.CommonEntityData {
    cospfShamLinkNbrEntry.EntityData.YFilter = cospfShamLinkNbrEntry.YFilter
    cospfShamLinkNbrEntry.EntityData.YangName = "cospfShamLinkNbrEntry"
    cospfShamLinkNbrEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinkNbrEntry.EntityData.ParentYangName = "cospfShamLinkNbrTable"
    cospfShamLinkNbrEntry.EntityData.SegmentPath = "cospfShamLinkNbrEntry" + types.AddKeyToken(cospfShamLinkNbrEntry.CospfShamLinksLocalIpAddrType, "cospfShamLinksLocalIpAddrType") + types.AddKeyToken(cospfShamLinkNbrEntry.CospfShamLinksLocalIpAddr, "cospfShamLinksLocalIpAddr") + types.AddKeyToken(cospfShamLinkNbrEntry.CospfShamLinkNbrArea, "cospfShamLinkNbrArea") + types.AddKeyToken(cospfShamLinkNbrEntry.CospfShamLinkNbrIpAddrType, "cospfShamLinkNbrIpAddrType") + types.AddKeyToken(cospfShamLinkNbrEntry.CospfShamLinkNbrIpAddr, "cospfShamLinkNbrIpAddr")
    cospfShamLinkNbrEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfShamLinkNbrTable/" + cospfShamLinkNbrEntry.EntityData.SegmentPath
    cospfShamLinkNbrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinkNbrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinkNbrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinkNbrEntry.EntityData.Children = types.NewOrderedMap()
    cospfShamLinkNbrEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinksLocalIpAddrType", types.YLeaf{"CospfShamLinksLocalIpAddrType", cospfShamLinkNbrEntry.CospfShamLinksLocalIpAddrType})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinksLocalIpAddr", types.YLeaf{"CospfShamLinksLocalIpAddr", cospfShamLinkNbrEntry.CospfShamLinksLocalIpAddr})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrArea", types.YLeaf{"CospfShamLinkNbrArea", cospfShamLinkNbrEntry.CospfShamLinkNbrArea})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrIpAddrType", types.YLeaf{"CospfShamLinkNbrIpAddrType", cospfShamLinkNbrEntry.CospfShamLinkNbrIpAddrType})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrIpAddr", types.YLeaf{"CospfShamLinkNbrIpAddr", cospfShamLinkNbrEntry.CospfShamLinkNbrIpAddr})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrRtrId", types.YLeaf{"CospfShamLinkNbrRtrId", cospfShamLinkNbrEntry.CospfShamLinkNbrRtrId})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrOptions", types.YLeaf{"CospfShamLinkNbrOptions", cospfShamLinkNbrEntry.CospfShamLinkNbrOptions})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrState", types.YLeaf{"CospfShamLinkNbrState", cospfShamLinkNbrEntry.CospfShamLinkNbrState})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrEvents", types.YLeaf{"CospfShamLinkNbrEvents", cospfShamLinkNbrEntry.CospfShamLinkNbrEvents})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrLsRetransQLen", types.YLeaf{"CospfShamLinkNbrLsRetransQLen", cospfShamLinkNbrEntry.CospfShamLinkNbrLsRetransQLen})
    cospfShamLinkNbrEntry.EntityData.Leafs.Append("cospfShamLinkNbrHelloSuppressed", types.YLeaf{"CospfShamLinkNbrHelloSuppressed", cospfShamLinkNbrEntry.CospfShamLinkNbrHelloSuppressed})

    cospfShamLinkNbrEntry.EntityData.YListKeys = []string {"CospfShamLinksLocalIpAddrType", "CospfShamLinksLocalIpAddr", "CospfShamLinkNbrArea", "CospfShamLinkNbrIpAddrType", "CospfShamLinkNbrIpAddr"}

    return &(cospfShamLinkNbrEntry.EntityData)
}

// CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState represents ship.
type CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState string

const (
    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_down CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "down"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_attempt CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "attempt"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_init CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "init"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_twoWay CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "twoWay"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_exchangeStart CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "exchangeStart"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_exchange CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "exchange"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_loading CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "loading"

    CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState_full CISCOOSPFMIB_CospfShamLinkNbrTable_CospfShamLinkNbrEntry_CospfShamLinkNbrState = "full"
)

// CISCOOSPFMIB_CospfShamLinksTable
// Information about this router's sham links.
type CISCOOSPFMIB_CospfShamLinksTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry.
    CospfShamLinksEntry []*CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry
}

func (cospfShamLinksTable *CISCOOSPFMIB_CospfShamLinksTable) GetEntityData() *types.CommonEntityData {
    cospfShamLinksTable.EntityData.YFilter = cospfShamLinksTable.YFilter
    cospfShamLinksTable.EntityData.YangName = "cospfShamLinksTable"
    cospfShamLinksTable.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinksTable.EntityData.ParentYangName = "CISCO-OSPF-MIB"
    cospfShamLinksTable.EntityData.SegmentPath = "cospfShamLinksTable"
    cospfShamLinksTable.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/" + cospfShamLinksTable.EntityData.SegmentPath
    cospfShamLinksTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinksTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinksTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinksTable.EntityData.Children = types.NewOrderedMap()
    cospfShamLinksTable.EntityData.Children.Append("cospfShamLinksEntry", types.YChild{"CospfShamLinksEntry", nil})
    for i := range cospfShamLinksTable.CospfShamLinksEntry {
        cospfShamLinksTable.EntityData.Children.Append(types.GetSegmentPath(cospfShamLinksTable.CospfShamLinksEntry[i]), types.YChild{"CospfShamLinksEntry", cospfShamLinksTable.CospfShamLinksEntry[i]})
    }
    cospfShamLinksTable.EntityData.Leafs = types.NewOrderedMap()

    cospfShamLinksTable.EntityData.YListKeys = []string {}

    return &(cospfShamLinksTable.EntityData)
}

// CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry
// Information about a single sham link.
type CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The area that this sham link is part of. The type
    // is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CospfShamLinksAreaId interface{}

    // This attribute is a key. The type of internet address of this sham link's
    // local IP address. The type is InetAddressType.
    CospfShamLinksLocalIpAddrType interface{}

    // This attribute is a key. The Local IP address of the sham link. The type is
    // string with length: 0..255.
    CospfShamLinksLocalIpAddr interface{}

    // This attribute is a key. The type of internet address of this sham link's
    // remote IP address. The type is InetAddressType.
    CospfShamLinksRemoteIpAddrType interface{}

    // This attribute is a key. The IP address of the other end router of the sham
    // link. The type is string with length: 0..255.
    CospfShamLinksRemoteIpAddr interface{}

    // The number of seconds between  link-state  advertisement retransmissions,
    // for adjacencies belonging to this link. This value is also used when
    // retransmitting database  description and link-state request packets. This
    // value should be well over the expected round trip time. The type is
    // interface{} with range: 0..3600.
    CospfShamLinksRetransInterval interface{}

    // The length of time, in  seconds,  between  the Hello  packets that the
    // router sends on the sham link. The type is interface{} with range:
    // 1..65535.
    CospfShamLinksHelloInterval interface{}

    // The number of seconds that  a  router's  Hello packets  have  not been seen
    // before it's neighbors declare the router down.  This  should  be some 
    // multiple  of  the  Hello  interval. The type is interface{} with range:
    // 0..2147483647.
    CospfShamLinksRtrDeadInterval interface{}

    // OSPF sham link states. The type is CospfShamLinksState.
    CospfShamLinksState interface{}

    // The number of state changes or error events on this sham link. The type is
    // interface{} with range: 0..4294967295.
    CospfShamLinksEvents interface{}

    // The Metric to be advertised. The type is interface{} with range: 0..65535.
    CospfShamLinksMetric interface{}
}

func (cospfShamLinksEntry *CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry) GetEntityData() *types.CommonEntityData {
    cospfShamLinksEntry.EntityData.YFilter = cospfShamLinksEntry.YFilter
    cospfShamLinksEntry.EntityData.YangName = "cospfShamLinksEntry"
    cospfShamLinksEntry.EntityData.BundleName = "cisco_ios_xe"
    cospfShamLinksEntry.EntityData.ParentYangName = "cospfShamLinksTable"
    cospfShamLinksEntry.EntityData.SegmentPath = "cospfShamLinksEntry" + types.AddKeyToken(cospfShamLinksEntry.CospfShamLinksAreaId, "cospfShamLinksAreaId") + types.AddKeyToken(cospfShamLinksEntry.CospfShamLinksLocalIpAddrType, "cospfShamLinksLocalIpAddrType") + types.AddKeyToken(cospfShamLinksEntry.CospfShamLinksLocalIpAddr, "cospfShamLinksLocalIpAddr") + types.AddKeyToken(cospfShamLinksEntry.CospfShamLinksRemoteIpAddrType, "cospfShamLinksRemoteIpAddrType") + types.AddKeyToken(cospfShamLinksEntry.CospfShamLinksRemoteIpAddr, "cospfShamLinksRemoteIpAddr")
    cospfShamLinksEntry.EntityData.AbsolutePath = "CISCO-OSPF-MIB:CISCO-OSPF-MIB/cospfShamLinksTable/" + cospfShamLinksEntry.EntityData.SegmentPath
    cospfShamLinksEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cospfShamLinksEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cospfShamLinksEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cospfShamLinksEntry.EntityData.Children = types.NewOrderedMap()
    cospfShamLinksEntry.EntityData.Leafs = types.NewOrderedMap()
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksAreaId", types.YLeaf{"CospfShamLinksAreaId", cospfShamLinksEntry.CospfShamLinksAreaId})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksLocalIpAddrType", types.YLeaf{"CospfShamLinksLocalIpAddrType", cospfShamLinksEntry.CospfShamLinksLocalIpAddrType})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksLocalIpAddr", types.YLeaf{"CospfShamLinksLocalIpAddr", cospfShamLinksEntry.CospfShamLinksLocalIpAddr})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksRemoteIpAddrType", types.YLeaf{"CospfShamLinksRemoteIpAddrType", cospfShamLinksEntry.CospfShamLinksRemoteIpAddrType})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksRemoteIpAddr", types.YLeaf{"CospfShamLinksRemoteIpAddr", cospfShamLinksEntry.CospfShamLinksRemoteIpAddr})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksRetransInterval", types.YLeaf{"CospfShamLinksRetransInterval", cospfShamLinksEntry.CospfShamLinksRetransInterval})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksHelloInterval", types.YLeaf{"CospfShamLinksHelloInterval", cospfShamLinksEntry.CospfShamLinksHelloInterval})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksRtrDeadInterval", types.YLeaf{"CospfShamLinksRtrDeadInterval", cospfShamLinksEntry.CospfShamLinksRtrDeadInterval})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksState", types.YLeaf{"CospfShamLinksState", cospfShamLinksEntry.CospfShamLinksState})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksEvents", types.YLeaf{"CospfShamLinksEvents", cospfShamLinksEntry.CospfShamLinksEvents})
    cospfShamLinksEntry.EntityData.Leafs.Append("cospfShamLinksMetric", types.YLeaf{"CospfShamLinksMetric", cospfShamLinksEntry.CospfShamLinksMetric})

    cospfShamLinksEntry.EntityData.YListKeys = []string {"CospfShamLinksAreaId", "CospfShamLinksLocalIpAddrType", "CospfShamLinksLocalIpAddr", "CospfShamLinksRemoteIpAddrType", "CospfShamLinksRemoteIpAddr"}

    return &(cospfShamLinksEntry.EntityData)
}

// CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState represents OSPF sham link states.
type CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState string

const (
    CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState_down CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState = "down"

    CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState_pointToPoint CISCOOSPFMIB_CospfShamLinksTable_CospfShamLinksEntry_CospfShamLinksState = "pointToPoint"
)

