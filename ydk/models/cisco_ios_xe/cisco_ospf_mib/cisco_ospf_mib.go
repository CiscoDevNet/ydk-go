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
    parent types.Entity
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

func (cISCOOSPFMIB *CISCOOSPFMIB) GetFilter() yfilter.YFilter { return cISCOOSPFMIB.YFilter }

func (cISCOOSPFMIB *CISCOOSPFMIB) SetFilter(yf yfilter.YFilter) { cISCOOSPFMIB.YFilter = yf }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetGoName(yname string) string {
    if yname == "cospfGeneralGroup" { return "Cospfgeneralgroup" }
    if yname == "cospfLsdbTable" { return "Cospflsdbtable" }
    if yname == "cospfShamLinkTable" { return "Cospfshamlinktable" }
    if yname == "cospfLocalLsdbTable" { return "Cospflocallsdbtable" }
    if yname == "cospfVirtLocalLsdbTable" { return "Cospfvirtlocallsdbtable" }
    if yname == "cospfShamLinkNbrTable" { return "Cospfshamlinknbrtable" }
    if yname == "cospfShamLinksTable" { return "Cospfshamlinkstable" }
    return ""
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetSegmentPath() string {
    return "CISCO-OSPF-MIB:CISCO-OSPF-MIB"
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfGeneralGroup" {
        return &cISCOOSPFMIB.Cospfgeneralgroup
    }
    if childYangName == "cospfLsdbTable" {
        return &cISCOOSPFMIB.Cospflsdbtable
    }
    if childYangName == "cospfShamLinkTable" {
        return &cISCOOSPFMIB.Cospfshamlinktable
    }
    if childYangName == "cospfLocalLsdbTable" {
        return &cISCOOSPFMIB.Cospflocallsdbtable
    }
    if childYangName == "cospfVirtLocalLsdbTable" {
        return &cISCOOSPFMIB.Cospfvirtlocallsdbtable
    }
    if childYangName == "cospfShamLinkNbrTable" {
        return &cISCOOSPFMIB.Cospfshamlinknbrtable
    }
    if childYangName == "cospfShamLinksTable" {
        return &cISCOOSPFMIB.Cospfshamlinkstable
    }
    return nil
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cospfGeneralGroup"] = &cISCOOSPFMIB.Cospfgeneralgroup
    children["cospfLsdbTable"] = &cISCOOSPFMIB.Cospflsdbtable
    children["cospfShamLinkTable"] = &cISCOOSPFMIB.Cospfshamlinktable
    children["cospfLocalLsdbTable"] = &cISCOOSPFMIB.Cospflocallsdbtable
    children["cospfVirtLocalLsdbTable"] = &cISCOOSPFMIB.Cospfvirtlocallsdbtable
    children["cospfShamLinkNbrTable"] = &cISCOOSPFMIB.Cospfshamlinknbrtable
    children["cospfShamLinksTable"] = &cISCOOSPFMIB.Cospfshamlinkstable
    return children
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOOSPFMIB *CISCOOSPFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetYangName() string { return "CISCO-OSPF-MIB" }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOOSPFMIB *CISCOOSPFMIB) SetParent(parent types.Entity) { cISCOOSPFMIB.parent = parent }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetParent() types.Entity { return cISCOOSPFMIB.parent }

func (cISCOOSPFMIB *CISCOOSPFMIB) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospfgeneralgroup
type CISCOOSPFMIB_Cospfgeneralgroup struct {
    parent types.Entity
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

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetFilter() yfilter.YFilter { return cospfgeneralgroup.YFilter }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) SetFilter(yf yfilter.YFilter) { cospfgeneralgroup.YFilter = yf }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetGoName(yname string) string {
    if yname == "cospfRFC1583Compatibility" { return "Cospfrfc1583Compatibility" }
    if yname == "cospfOpaqueLsaSupport" { return "Cospfopaquelsasupport" }
    if yname == "cospfTrafficEngineeringSupport" { return "Cospftrafficengineeringsupport" }
    if yname == "cospfOpaqueASLsaCount" { return "Cospfopaqueaslsacount" }
    if yname == "cospfOpaqueASLsaCksumSum" { return "Cospfopaqueaslsacksumsum" }
    return ""
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetSegmentPath() string {
    return "cospfGeneralGroup"
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfRFC1583Compatibility"] = cospfgeneralgroup.Cospfrfc1583Compatibility
    leafs["cospfOpaqueLsaSupport"] = cospfgeneralgroup.Cospfopaquelsasupport
    leafs["cospfTrafficEngineeringSupport"] = cospfgeneralgroup.Cospftrafficengineeringsupport
    leafs["cospfOpaqueASLsaCount"] = cospfgeneralgroup.Cospfopaqueaslsacount
    leafs["cospfOpaqueASLsaCksumSum"] = cospfgeneralgroup.Cospfopaqueaslsacksumsum
    return leafs
}

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetBundleName() string { return "cisco_ios_xe" }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetYangName() string { return "cospfGeneralGroup" }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) SetParent(parent types.Entity) { cospfgeneralgroup.parent = parent }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetParent() types.Entity { return cospfgeneralgroup.parent }

func (cospfgeneralgroup *CISCOOSPFMIB_Cospfgeneralgroup) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospflsdbtable
// The OSPF Process's Link State Database. This 
// table is meant for Opaque LSA's
type CISCOOSPFMIB_Cospflsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry.
    Cospflsdbentry []CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetFilter() yfilter.YFilter { return cospflsdbtable.YFilter }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) SetFilter(yf yfilter.YFilter) { cospflsdbtable.YFilter = yf }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetGoName(yname string) string {
    if yname == "cospfLsdbEntry" { return "Cospflsdbentry" }
    return ""
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetSegmentPath() string {
    return "cospfLsdbTable"
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfLsdbEntry" {
        for _, c := range cospflsdbtable.Cospflsdbentry {
            if cospflsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry{}
        cospflsdbtable.Cospflsdbentry = append(cospflsdbtable.Cospflsdbentry, child)
        return &cospflsdbtable.Cospflsdbentry[len(cospflsdbtable.Cospflsdbentry)-1]
    }
    return nil
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospflsdbtable.Cospflsdbentry {
        children[cospflsdbtable.Cospflsdbentry[i].GetSegmentPath()] = &cospflsdbtable.Cospflsdbentry[i]
    }
    return children
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetYangName() string { return "cospfLsdbTable" }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) SetParent(parent types.Entity) { cospflsdbtable.parent = parent }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetParent() types.Entity { return cospflsdbtable.parent }

func (cospflsdbtable *CISCOOSPFMIB_Cospflsdbtable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to ospf_mib.OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdbareaid
    Ospflsdbareaid interface{}

    // This attribute is a key. The type of the link state advertisement. Each
    // link state type has a separate advertisement format. The type is
    // Cospflsdbtype.
    Cospflsdbtype interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to ospf_mib.OSPFMIB_Ospflsdbtable_Ospflsdbentry_Ospflsdblsid
    Ospflsdblsid interface{}

    // This attribute is a key. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetFilter() yfilter.YFilter { return cospflsdbentry.YFilter }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) SetFilter(yf yfilter.YFilter) { cospflsdbentry.YFilter = yf }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetGoName(yname string) string {
    if yname == "ospfLsdbAreaId" { return "Ospflsdbareaid" }
    if yname == "cospfLsdbType" { return "Cospflsdbtype" }
    if yname == "ospfLsdbLsid" { return "Ospflsdblsid" }
    if yname == "ospfLsdbRouterId" { return "Ospflsdbrouterid" }
    if yname == "cospfLsdbSequence" { return "Cospflsdbsequence" }
    if yname == "cospfLsdbAge" { return "Cospflsdbage" }
    if yname == "cospfLsdbChecksum" { return "Cospflsdbchecksum" }
    if yname == "cospfLsdbAdvertisement" { return "Cospflsdbadvertisement" }
    return ""
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetSegmentPath() string {
    return "cospfLsdbEntry" + "[ospfLsdbAreaId='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdbareaid) + "']" + "[cospfLsdbType='" + fmt.Sprintf("%v", cospflsdbentry.Cospflsdbtype) + "']" + "[ospfLsdbLsid='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdblsid) + "']" + "[ospfLsdbRouterId='" + fmt.Sprintf("%v", cospflsdbentry.Ospflsdbrouterid) + "']"
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ospfLsdbAreaId"] = cospflsdbentry.Ospflsdbareaid
    leafs["cospfLsdbType"] = cospflsdbentry.Cospflsdbtype
    leafs["ospfLsdbLsid"] = cospflsdbentry.Ospflsdblsid
    leafs["ospfLsdbRouterId"] = cospflsdbentry.Ospflsdbrouterid
    leafs["cospfLsdbSequence"] = cospflsdbentry.Cospflsdbsequence
    leafs["cospfLsdbAge"] = cospflsdbentry.Cospflsdbage
    leafs["cospfLsdbChecksum"] = cospflsdbentry.Cospflsdbchecksum
    leafs["cospfLsdbAdvertisement"] = cospflsdbentry.Cospflsdbadvertisement
    return leafs
}

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetYangName() string { return "cospfLsdbEntry" }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) SetParent(parent types.Entity) { cospflsdbentry.parent = parent }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetParent() types.Entity { return cospflsdbentry.parent }

func (cospflsdbentry *CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry) GetParentYangName() string { return "cospfLsdbTable" }

// CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype string

const (
    CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype_areaOpaqueLink CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype = "areaOpaqueLink"

    CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype_asOpaqueLink CISCOOSPFMIB_Cospflsdbtable_Cospflsdbentry_Cospflsdbtype = "asOpaqueLink"
)

// CISCOOSPFMIB_Cospfshamlinktable
// Information about this router's sham links
type CISCOOSPFMIB_Cospfshamlinktable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry.
    Cospfshamlinkentry []CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetFilter() yfilter.YFilter { return cospfshamlinktable.YFilter }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) SetFilter(yf yfilter.YFilter) { cospfshamlinktable.YFilter = yf }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetGoName(yname string) string {
    if yname == "cospfShamLinkEntry" { return "Cospfshamlinkentry" }
    return ""
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetSegmentPath() string {
    return "cospfShamLinkTable"
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfShamLinkEntry" {
        for _, c := range cospfshamlinktable.Cospfshamlinkentry {
            if cospfshamlinktable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry{}
        cospfshamlinktable.Cospfshamlinkentry = append(cospfshamlinktable.Cospfshamlinkentry, child)
        return &cospfshamlinktable.Cospfshamlinkentry[len(cospfshamlinktable.Cospfshamlinkentry)-1]
    }
    return nil
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospfshamlinktable.Cospfshamlinkentry {
        children[cospfshamlinktable.Cospfshamlinkentry[i].GetSegmentPath()] = &cospfshamlinktable.Cospfshamlinkentry[i]
    }
    return children
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetYangName() string { return "cospfShamLinkTable" }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) SetParent(parent types.Entity) { cospfshamlinktable.parent = parent }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetParent() types.Entity { return cospfshamlinktable.parent }

func (cospfshamlinktable *CISCOOSPFMIB_Cospfshamlinktable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry
// Information about a single sham link
type CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The  Transit  Area  that  the   Virtual   Link
    // traverses.  By definition, this is not 0.0.0.0. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfshamlinkareaid interface{}

    // This attribute is a key. The Local IP address of the sham link. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfshamlinklocalipaddress interface{}

    // This attribute is a key. The Router ID of the other end router of the sham
    // link. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetFilter() yfilter.YFilter { return cospfshamlinkentry.YFilter }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) SetFilter(yf yfilter.YFilter) { cospfshamlinkentry.YFilter = yf }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetGoName(yname string) string {
    if yname == "cospfShamLinkAreaId" { return "Cospfshamlinkareaid" }
    if yname == "cospfShamLinkLocalIpAddress" { return "Cospfshamlinklocalipaddress" }
    if yname == "cospfShamLinkNeighborId" { return "Cospfshamlinkneighborid" }
    if yname == "cospfShamLinkRetransInterval" { return "Cospfshamlinkretransinterval" }
    if yname == "cospfShamLinkHelloInterval" { return "Cospfshamlinkhellointerval" }
    if yname == "cospfShamLinkRtrDeadInterval" { return "Cospfshamlinkrtrdeadinterval" }
    if yname == "cospfShamLinkState" { return "Cospfshamlinkstate" }
    if yname == "cospfShamLinkEvents" { return "Cospfshamlinkevents" }
    if yname == "cospfShamLinkMetric" { return "Cospfshamlinkmetric" }
    return ""
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetSegmentPath() string {
    return "cospfShamLinkEntry" + "[cospfShamLinkAreaId='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinkareaid) + "']" + "[cospfShamLinkLocalIpAddress='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinklocalipaddress) + "']" + "[cospfShamLinkNeighborId='" + fmt.Sprintf("%v", cospfshamlinkentry.Cospfshamlinkneighborid) + "']"
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfShamLinkAreaId"] = cospfshamlinkentry.Cospfshamlinkareaid
    leafs["cospfShamLinkLocalIpAddress"] = cospfshamlinkentry.Cospfshamlinklocalipaddress
    leafs["cospfShamLinkNeighborId"] = cospfshamlinkentry.Cospfshamlinkneighborid
    leafs["cospfShamLinkRetransInterval"] = cospfshamlinkentry.Cospfshamlinkretransinterval
    leafs["cospfShamLinkHelloInterval"] = cospfshamlinkentry.Cospfshamlinkhellointerval
    leafs["cospfShamLinkRtrDeadInterval"] = cospfshamlinkentry.Cospfshamlinkrtrdeadinterval
    leafs["cospfShamLinkState"] = cospfshamlinkentry.Cospfshamlinkstate
    leafs["cospfShamLinkEvents"] = cospfshamlinkentry.Cospfshamlinkevents
    leafs["cospfShamLinkMetric"] = cospfshamlinkentry.Cospfshamlinkmetric
    return leafs
}

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetYangName() string { return "cospfShamLinkEntry" }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) SetParent(parent types.Entity) { cospfshamlinkentry.parent = parent }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetParent() types.Entity { return cospfshamlinkentry.parent }

func (cospfshamlinkentry *CISCOOSPFMIB_Cospfshamlinktable_Cospfshamlinkentry) GetParentYangName() string { return "cospfShamLinkTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry.
    Cospflocallsdbentry []CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetFilter() yfilter.YFilter { return cospflocallsdbtable.YFilter }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) SetFilter(yf yfilter.YFilter) { cospflocallsdbtable.YFilter = yf }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetGoName(yname string) string {
    if yname == "cospfLocalLsdbEntry" { return "Cospflocallsdbentry" }
    return ""
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetSegmentPath() string {
    return "cospfLocalLsdbTable"
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfLocalLsdbEntry" {
        for _, c := range cospflocallsdbtable.Cospflocallsdbentry {
            if cospflocallsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry{}
        cospflocallsdbtable.Cospflocallsdbentry = append(cospflocallsdbtable.Cospflocallsdbentry, child)
        return &cospflocallsdbtable.Cospflocallsdbentry[len(cospflocallsdbtable.Cospflocallsdbentry)-1]
    }
    return nil
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospflocallsdbtable.Cospflocallsdbentry {
        children[cospflocallsdbtable.Cospflocallsdbentry[i].GetSegmentPath()] = &cospflocallsdbtable.Cospflocallsdbentry[i]
    }
    return children
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetYangName() string { return "cospfLocalLsdbTable" }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) SetParent(parent types.Entity) { cospflocallsdbtable.parent = parent }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetParent() types.Entity { return cospflocallsdbtable.parent }

func (cospflocallsdbtable *CISCOOSPFMIB_Cospflocallsdbtable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP Address of the interface from which the LSA
    // was received if the interface is numbered. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospflocallsdblsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetFilter() yfilter.YFilter { return cospflocallsdbentry.YFilter }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) SetFilter(yf yfilter.YFilter) { cospflocallsdbentry.YFilter = yf }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetGoName(yname string) string {
    if yname == "cospfLocalLsdbIpAddress" { return "Cospflocallsdbipaddress" }
    if yname == "cospfLocalLsdbAddressLessIf" { return "Cospflocallsdbaddresslessif" }
    if yname == "cospfLocalLsdbType" { return "Cospflocallsdbtype" }
    if yname == "cospfLocalLsdbLsid" { return "Cospflocallsdblsid" }
    if yname == "cospfLocalLsdbRouterId" { return "Cospflocallsdbrouterid" }
    if yname == "cospfLocalLsdbSequence" { return "Cospflocallsdbsequence" }
    if yname == "cospfLocalLsdbAge" { return "Cospflocallsdbage" }
    if yname == "cospfLocalLsdbChecksum" { return "Cospflocallsdbchecksum" }
    if yname == "cospfLocalLsdbAdvertisement" { return "Cospflocallsdbadvertisement" }
    return ""
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetSegmentPath() string {
    return "cospfLocalLsdbEntry" + "[cospfLocalLsdbIpAddress='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbipaddress) + "']" + "[cospfLocalLsdbAddressLessIf='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbaddresslessif) + "']" + "[cospfLocalLsdbType='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbtype) + "']" + "[cospfLocalLsdbLsid='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdblsid) + "']" + "[cospfLocalLsdbRouterId='" + fmt.Sprintf("%v", cospflocallsdbentry.Cospflocallsdbrouterid) + "']"
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfLocalLsdbIpAddress"] = cospflocallsdbentry.Cospflocallsdbipaddress
    leafs["cospfLocalLsdbAddressLessIf"] = cospflocallsdbentry.Cospflocallsdbaddresslessif
    leafs["cospfLocalLsdbType"] = cospflocallsdbentry.Cospflocallsdbtype
    leafs["cospfLocalLsdbLsid"] = cospflocallsdbentry.Cospflocallsdblsid
    leafs["cospfLocalLsdbRouterId"] = cospflocallsdbentry.Cospflocallsdbrouterid
    leafs["cospfLocalLsdbSequence"] = cospflocallsdbentry.Cospflocallsdbsequence
    leafs["cospfLocalLsdbAge"] = cospflocallsdbentry.Cospflocallsdbage
    leafs["cospfLocalLsdbChecksum"] = cospflocallsdbentry.Cospflocallsdbchecksum
    leafs["cospfLocalLsdbAdvertisement"] = cospflocallsdbentry.Cospflocallsdbadvertisement
    return leafs
}

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetYangName() string { return "cospfLocalLsdbEntry" }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) SetParent(parent types.Entity) { cospflocallsdbentry.parent = parent }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetParent() types.Entity { return cospflocallsdbentry.parent }

func (cospflocallsdbentry *CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry) GetParentYangName() string { return "cospfLocalLsdbTable" }

// CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype represents Each link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype string

const (
    CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype_localOpaqueLink CISCOOSPFMIB_Cospflocallsdbtable_Cospflocallsdbentry_Cospflocallsdbtype = "localOpaqueLink"
)

// CISCOOSPFMIB_Cospfvirtlocallsdbtable
// The OSPF Process's Link-Local Link State Database
// for virtual links.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single Link State Advertisement. The type is slice of
    // CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry.
    Cospfvirtlocallsdbentry []CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetFilter() yfilter.YFilter { return cospfvirtlocallsdbtable.YFilter }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) SetFilter(yf yfilter.YFilter) { cospfvirtlocallsdbtable.YFilter = yf }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetGoName(yname string) string {
    if yname == "cospfVirtLocalLsdbEntry" { return "Cospfvirtlocallsdbentry" }
    return ""
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetSegmentPath() string {
    return "cospfVirtLocalLsdbTable"
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfVirtLocalLsdbEntry" {
        for _, c := range cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry {
            if cospfvirtlocallsdbtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry{}
        cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry = append(cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry, child)
        return &cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry[len(cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry)-1]
    }
    return nil
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry {
        children[cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry[i].GetSegmentPath()] = &cospfvirtlocallsdbtable.Cospfvirtlocallsdbentry[i]
    }
    return children
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetBundleName() string { return "cisco_ios_xe" }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetYangName() string { return "cospfVirtLocalLsdbTable" }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) SetParent(parent types.Entity) { cospfvirtlocallsdbtable.parent = parent }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetParent() types.Entity { return cospfvirtlocallsdbtable.parent }

func (cospfvirtlocallsdbtable *CISCOOSPFMIB_Cospfvirtlocallsdbtable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry
// A single Link State Advertisement.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Transit Area that the Virtual Link traverses.
    // By definition, this is not 0.0.0.0. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfvirtlocallsdbtransitarea interface{}

    // This attribute is a key. The Router ID of the Virtual Neighbor. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfvirtlocallsdbneighbor interface{}

    // This attribute is a key. The type of the link state advertisement. Each 
    // link state type has a separate advertisement format. The type is
    // Cospfvirtlocallsdbtype.
    Cospfvirtlocallsdbtype interface{}

    // This attribute is a key. The Link State ID is an LS Type Specific field
    // containing a 32 bit identifier in IP address format; it identifies the
    // piece of the routing domain that is being described by the advertisement.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfvirtlocallsdblsid interface{}

    // This attribute is a key. The 32 bit number that uniquely identifies the
    // originating router in the Autonomous System. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetFilter() yfilter.YFilter { return cospfvirtlocallsdbentry.YFilter }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) SetFilter(yf yfilter.YFilter) { cospfvirtlocallsdbentry.YFilter = yf }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetGoName(yname string) string {
    if yname == "cospfVirtLocalLsdbTransitArea" { return "Cospfvirtlocallsdbtransitarea" }
    if yname == "cospfVirtLocalLsdbNeighbor" { return "Cospfvirtlocallsdbneighbor" }
    if yname == "cospfVirtLocalLsdbType" { return "Cospfvirtlocallsdbtype" }
    if yname == "cospfVirtLocalLsdbLsid" { return "Cospfvirtlocallsdblsid" }
    if yname == "cospfVirtLocalLsdbRouterId" { return "Cospfvirtlocallsdbrouterid" }
    if yname == "cospfVirtLocalLsdbSequence" { return "Cospfvirtlocallsdbsequence" }
    if yname == "cospfVirtLocalLsdbAge" { return "Cospfvirtlocallsdbage" }
    if yname == "cospfVirtLocalLsdbChecksum" { return "Cospfvirtlocallsdbchecksum" }
    if yname == "cospfVirtLocalLsdbAdvertisement" { return "Cospfvirtlocallsdbadvertisement" }
    return ""
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetSegmentPath() string {
    return "cospfVirtLocalLsdbEntry" + "[cospfVirtLocalLsdbTransitArea='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtransitarea) + "']" + "[cospfVirtLocalLsdbNeighbor='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbneighbor) + "']" + "[cospfVirtLocalLsdbType='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbtype) + "']" + "[cospfVirtLocalLsdbLsid='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdblsid) + "']" + "[cospfVirtLocalLsdbRouterId='" + fmt.Sprintf("%v", cospfvirtlocallsdbentry.Cospfvirtlocallsdbrouterid) + "']"
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfVirtLocalLsdbTransitArea"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbtransitarea
    leafs["cospfVirtLocalLsdbNeighbor"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbneighbor
    leafs["cospfVirtLocalLsdbType"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbtype
    leafs["cospfVirtLocalLsdbLsid"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdblsid
    leafs["cospfVirtLocalLsdbRouterId"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbrouterid
    leafs["cospfVirtLocalLsdbSequence"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbsequence
    leafs["cospfVirtLocalLsdbAge"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbage
    leafs["cospfVirtLocalLsdbChecksum"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbchecksum
    leafs["cospfVirtLocalLsdbAdvertisement"] = cospfvirtlocallsdbentry.Cospfvirtlocallsdbadvertisement
    return leafs
}

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetYangName() string { return "cospfVirtLocalLsdbEntry" }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) SetParent(parent types.Entity) { cospfvirtlocallsdbentry.parent = parent }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetParent() types.Entity { return cospfvirtlocallsdbentry.parent }

func (cospfvirtlocallsdbentry *CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry) GetParentYangName() string { return "cospfVirtLocalLsdbTable" }

// CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype represents Each  link state type has a separate advertisement format.
type CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype string

const (
    CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype_localOpaqueLink CISCOOSPFMIB_Cospfvirtlocallsdbtable_Cospfvirtlocallsdbentry_Cospfvirtlocallsdbtype = "localOpaqueLink"
)

// CISCOOSPFMIB_Cospfshamlinknbrtable
// A table of sham link neighbor information.
type CISCOOSPFMIB_Cospfshamlinknbrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sham link neighbor information. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry.
    Cospfshamlinknbrentry []CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetFilter() yfilter.YFilter { return cospfshamlinknbrtable.YFilter }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) SetFilter(yf yfilter.YFilter) { cospfshamlinknbrtable.YFilter = yf }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetGoName(yname string) string {
    if yname == "cospfShamLinkNbrEntry" { return "Cospfshamlinknbrentry" }
    return ""
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetSegmentPath() string {
    return "cospfShamLinkNbrTable"
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfShamLinkNbrEntry" {
        for _, c := range cospfshamlinknbrtable.Cospfshamlinknbrentry {
            if cospfshamlinknbrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry{}
        cospfshamlinknbrtable.Cospfshamlinknbrentry = append(cospfshamlinknbrtable.Cospfshamlinknbrentry, child)
        return &cospfshamlinknbrtable.Cospfshamlinknbrentry[len(cospfshamlinknbrtable.Cospfshamlinknbrentry)-1]
    }
    return nil
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospfshamlinknbrtable.Cospfshamlinknbrentry {
        children[cospfshamlinknbrtable.Cospfshamlinknbrentry[i].GetSegmentPath()] = &cospfshamlinknbrtable.Cospfshamlinknbrentry[i]
    }
    return children
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetYangName() string { return "cospfShamLinkNbrTable" }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) SetParent(parent types.Entity) { cospfshamlinknbrtable.parent = parent }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetParent() types.Entity { return cospfshamlinknbrtable.parent }

func (cospfshamlinknbrtable *CISCOOSPFMIB_Cospfshamlinknbrtable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry
// Sham link neighbor information.
type CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is InetAddressType. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinkslocalipaddrtype
    Cospfshamlinkslocalipaddrtype interface{}

    // This attribute is a key. The type is string with length: 0..255. Refers to
    // cisco_ospf_mib.CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinkslocalipaddr
    Cospfshamlinkslocalipaddr interface{}

    // This attribute is a key. The area to which the sham link is part of. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Cospfshamlinknbrarea interface{}

    // This attribute is a key. The type of internet address of this sham link
    // neighbor's IP address. The type is InetAddressType.
    Cospfshamlinknbripaddrtype interface{}

    // This attribute is a key. The IP address this sham link neighbor is using.
    // The type is string with length: 0..255.
    Cospfshamlinknbripaddr interface{}

    // A 32-bit integer uniquely identifying the neighboring router. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetFilter() yfilter.YFilter { return cospfshamlinknbrentry.YFilter }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) SetFilter(yf yfilter.YFilter) { cospfshamlinknbrentry.YFilter = yf }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetGoName(yname string) string {
    if yname == "cospfShamLinksLocalIpAddrType" { return "Cospfshamlinkslocalipaddrtype" }
    if yname == "cospfShamLinksLocalIpAddr" { return "Cospfshamlinkslocalipaddr" }
    if yname == "cospfShamLinkNbrArea" { return "Cospfshamlinknbrarea" }
    if yname == "cospfShamLinkNbrIpAddrType" { return "Cospfshamlinknbripaddrtype" }
    if yname == "cospfShamLinkNbrIpAddr" { return "Cospfshamlinknbripaddr" }
    if yname == "cospfShamLinkNbrRtrId" { return "Cospfshamlinknbrrtrid" }
    if yname == "cospfShamLinkNbrOptions" { return "Cospfshamlinknbroptions" }
    if yname == "cospfShamLinkNbrState" { return "Cospfshamlinknbrstate" }
    if yname == "cospfShamLinkNbrEvents" { return "Cospfshamlinknbrevents" }
    if yname == "cospfShamLinkNbrLsRetransQLen" { return "Cospfshamlinknbrlsretransqlen" }
    if yname == "cospfShamLinkNbrHelloSuppressed" { return "Cospfshamlinknbrhellosuppressed" }
    return ""
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetSegmentPath() string {
    return "cospfShamLinkNbrEntry" + "[cospfShamLinksLocalIpAddrType='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinkslocalipaddrtype) + "']" + "[cospfShamLinksLocalIpAddr='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinkslocalipaddr) + "']" + "[cospfShamLinkNbrArea='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbrarea) + "']" + "[cospfShamLinkNbrIpAddrType='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbripaddrtype) + "']" + "[cospfShamLinkNbrIpAddr='" + fmt.Sprintf("%v", cospfshamlinknbrentry.Cospfshamlinknbripaddr) + "']"
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfShamLinksLocalIpAddrType"] = cospfshamlinknbrentry.Cospfshamlinkslocalipaddrtype
    leafs["cospfShamLinksLocalIpAddr"] = cospfshamlinknbrentry.Cospfshamlinkslocalipaddr
    leafs["cospfShamLinkNbrArea"] = cospfshamlinknbrentry.Cospfshamlinknbrarea
    leafs["cospfShamLinkNbrIpAddrType"] = cospfshamlinknbrentry.Cospfshamlinknbripaddrtype
    leafs["cospfShamLinkNbrIpAddr"] = cospfshamlinknbrentry.Cospfshamlinknbripaddr
    leafs["cospfShamLinkNbrRtrId"] = cospfshamlinknbrentry.Cospfshamlinknbrrtrid
    leafs["cospfShamLinkNbrOptions"] = cospfshamlinknbrentry.Cospfshamlinknbroptions
    leafs["cospfShamLinkNbrState"] = cospfshamlinknbrentry.Cospfshamlinknbrstate
    leafs["cospfShamLinkNbrEvents"] = cospfshamlinknbrentry.Cospfshamlinknbrevents
    leafs["cospfShamLinkNbrLsRetransQLen"] = cospfshamlinknbrentry.Cospfshamlinknbrlsretransqlen
    leafs["cospfShamLinkNbrHelloSuppressed"] = cospfshamlinknbrentry.Cospfshamlinknbrhellosuppressed
    return leafs
}

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetYangName() string { return "cospfShamLinkNbrEntry" }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) SetParent(parent types.Entity) { cospfshamlinknbrentry.parent = parent }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetParent() types.Entity { return cospfshamlinknbrentry.parent }

func (cospfshamlinknbrentry *CISCOOSPFMIB_Cospfshamlinknbrtable_Cospfshamlinknbrentry) GetParentYangName() string { return "cospfShamLinkNbrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single sham link. The type is slice of
    // CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry.
    Cospfshamlinksentry []CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetFilter() yfilter.YFilter { return cospfshamlinkstable.YFilter }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) SetFilter(yf yfilter.YFilter) { cospfshamlinkstable.YFilter = yf }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetGoName(yname string) string {
    if yname == "cospfShamLinksEntry" { return "Cospfshamlinksentry" }
    return ""
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetSegmentPath() string {
    return "cospfShamLinksTable"
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cospfShamLinksEntry" {
        for _, c := range cospfshamlinkstable.Cospfshamlinksentry {
            if cospfshamlinkstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry{}
        cospfshamlinkstable.Cospfshamlinksentry = append(cospfshamlinkstable.Cospfshamlinksentry, child)
        return &cospfshamlinkstable.Cospfshamlinksentry[len(cospfshamlinkstable.Cospfshamlinksentry)-1]
    }
    return nil
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cospfshamlinkstable.Cospfshamlinksentry {
        children[cospfshamlinkstable.Cospfshamlinksentry[i].GetSegmentPath()] = &cospfshamlinkstable.Cospfshamlinksentry[i]
    }
    return children
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetYangName() string { return "cospfShamLinksTable" }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) SetParent(parent types.Entity) { cospfshamlinkstable.parent = parent }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetParent() types.Entity { return cospfshamlinkstable.parent }

func (cospfshamlinkstable *CISCOOSPFMIB_Cospfshamlinkstable) GetParentYangName() string { return "CISCO-OSPF-MIB" }

// CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry
// Information about a single sham link.
type CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The area that this sham link is part of. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetFilter() yfilter.YFilter { return cospfshamlinksentry.YFilter }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) SetFilter(yf yfilter.YFilter) { cospfshamlinksentry.YFilter = yf }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetGoName(yname string) string {
    if yname == "cospfShamLinksAreaId" { return "Cospfshamlinksareaid" }
    if yname == "cospfShamLinksLocalIpAddrType" { return "Cospfshamlinkslocalipaddrtype" }
    if yname == "cospfShamLinksLocalIpAddr" { return "Cospfshamlinkslocalipaddr" }
    if yname == "cospfShamLinksRemoteIpAddrType" { return "Cospfshamlinksremoteipaddrtype" }
    if yname == "cospfShamLinksRemoteIpAddr" { return "Cospfshamlinksremoteipaddr" }
    if yname == "cospfShamLinksRetransInterval" { return "Cospfshamlinksretransinterval" }
    if yname == "cospfShamLinksHelloInterval" { return "Cospfshamlinkshellointerval" }
    if yname == "cospfShamLinksRtrDeadInterval" { return "Cospfshamlinksrtrdeadinterval" }
    if yname == "cospfShamLinksState" { return "Cospfshamlinksstate" }
    if yname == "cospfShamLinksEvents" { return "Cospfshamlinksevents" }
    if yname == "cospfShamLinksMetric" { return "Cospfshamlinksmetric" }
    return ""
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetSegmentPath() string {
    return "cospfShamLinksEntry" + "[cospfShamLinksAreaId='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksareaid) + "']" + "[cospfShamLinksLocalIpAddrType='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinkslocalipaddrtype) + "']" + "[cospfShamLinksLocalIpAddr='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinkslocalipaddr) + "']" + "[cospfShamLinksRemoteIpAddrType='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksremoteipaddrtype) + "']" + "[cospfShamLinksRemoteIpAddr='" + fmt.Sprintf("%v", cospfshamlinksentry.Cospfshamlinksremoteipaddr) + "']"
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cospfShamLinksAreaId"] = cospfshamlinksentry.Cospfshamlinksareaid
    leafs["cospfShamLinksLocalIpAddrType"] = cospfshamlinksentry.Cospfshamlinkslocalipaddrtype
    leafs["cospfShamLinksLocalIpAddr"] = cospfshamlinksentry.Cospfshamlinkslocalipaddr
    leafs["cospfShamLinksRemoteIpAddrType"] = cospfshamlinksentry.Cospfshamlinksremoteipaddrtype
    leafs["cospfShamLinksRemoteIpAddr"] = cospfshamlinksentry.Cospfshamlinksremoteipaddr
    leafs["cospfShamLinksRetransInterval"] = cospfshamlinksentry.Cospfshamlinksretransinterval
    leafs["cospfShamLinksHelloInterval"] = cospfshamlinksentry.Cospfshamlinkshellointerval
    leafs["cospfShamLinksRtrDeadInterval"] = cospfshamlinksentry.Cospfshamlinksrtrdeadinterval
    leafs["cospfShamLinksState"] = cospfshamlinksentry.Cospfshamlinksstate
    leafs["cospfShamLinksEvents"] = cospfshamlinksentry.Cospfshamlinksevents
    leafs["cospfShamLinksMetric"] = cospfshamlinksentry.Cospfshamlinksmetric
    return leafs
}

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetBundleName() string { return "cisco_ios_xe" }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetYangName() string { return "cospfShamLinksEntry" }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) SetParent(parent types.Entity) { cospfshamlinksentry.parent = parent }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetParent() types.Entity { return cospfshamlinksentry.parent }

func (cospfshamlinksentry *CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry) GetParentYangName() string { return "cospfShamLinksTable" }

// CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate represents OSPF sham link states.
type CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate string

const (
    CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate_down CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate = "down"

    CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate_pointToPoint CISCOOSPFMIB_Cospfshamlinkstable_Cospfshamlinksentry_Cospfshamlinksstate = "pointToPoint"
)

