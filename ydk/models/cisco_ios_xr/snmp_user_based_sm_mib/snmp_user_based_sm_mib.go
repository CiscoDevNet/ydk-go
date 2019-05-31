package snmp_user_based_sm_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_user_based_sm_mib"))
    ydk.RegisterEntity("{http://tail-f.com/ns/mibs/SNMP-USER-BASED-SM-MIB/200210160000Z SNMP-USER-BASED-SM-MIB}", reflect.TypeOf(SNMPUSERBASEDSMMIB{}))
    ydk.RegisterEntity("SNMP-USER-BASED-SM-MIB:SNMP-USER-BASED-SM-MIB", reflect.TypeOf(SNMPUSERBASEDSMMIB{}))
}

// SNMPUSERBASEDSMMIB
type SNMPUSERBASEDSMMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    UsmStats SNMPUSERBASEDSMMIB_UsmStats

    
    UsmUserTable SNMPUSERBASEDSMMIB_UsmUserTable
}

func (sNMPUSERBASEDSMMIB *SNMPUSERBASEDSMMIB) GetEntityData() *types.CommonEntityData {
    sNMPUSERBASEDSMMIB.EntityData.YFilter = sNMPUSERBASEDSMMIB.YFilter
    sNMPUSERBASEDSMMIB.EntityData.YangName = "SNMP-USER-BASED-SM-MIB"
    sNMPUSERBASEDSMMIB.EntityData.BundleName = "cisco_ios_xr"
    sNMPUSERBASEDSMMIB.EntityData.ParentYangName = "SNMP-USER-BASED-SM-MIB"
    sNMPUSERBASEDSMMIB.EntityData.SegmentPath = "SNMP-USER-BASED-SM-MIB:SNMP-USER-BASED-SM-MIB"
    sNMPUSERBASEDSMMIB.EntityData.AbsolutePath = sNMPUSERBASEDSMMIB.EntityData.SegmentPath
    sNMPUSERBASEDSMMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sNMPUSERBASEDSMMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sNMPUSERBASEDSMMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sNMPUSERBASEDSMMIB.EntityData.Children = types.NewOrderedMap()
    sNMPUSERBASEDSMMIB.EntityData.Children.Append("usmStats", types.YChild{"UsmStats", &sNMPUSERBASEDSMMIB.UsmStats})
    sNMPUSERBASEDSMMIB.EntityData.Children.Append("usmUserTable", types.YChild{"UsmUserTable", &sNMPUSERBASEDSMMIB.UsmUserTable})
    sNMPUSERBASEDSMMIB.EntityData.Leafs = types.NewOrderedMap()

    sNMPUSERBASEDSMMIB.EntityData.YListKeys = []string {}

    return &(sNMPUSERBASEDSMMIB.EntityData)
}

// SNMPUSERBASEDSMMIB_UsmStats
type SNMPUSERBASEDSMMIB_UsmStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    UsmStatsUnsupportedSecLevels interface{}

    // The type is interface{} with range: 0..4294967295.
    UsmStatsNotInTimeWindows interface{}

    // The type is interface{} with range: 0..4294967295.
    UsmStatsUnknownUserNames interface{}

    // The type is interface{} with range: 0..4294967295.
    UsmStatsUnknownEngineIDs interface{}

    // The type is interface{} with range: 0..4294967295.
    UsmStatsWrongDigests interface{}

    // The type is interface{} with range: 0..4294967295.
    UsmStatsDecryptionErrors interface{}
}

func (usmStats *SNMPUSERBASEDSMMIB_UsmStats) GetEntityData() *types.CommonEntityData {
    usmStats.EntityData.YFilter = usmStats.YFilter
    usmStats.EntityData.YangName = "usmStats"
    usmStats.EntityData.BundleName = "cisco_ios_xr"
    usmStats.EntityData.ParentYangName = "SNMP-USER-BASED-SM-MIB"
    usmStats.EntityData.SegmentPath = "usmStats"
    usmStats.EntityData.AbsolutePath = "SNMP-USER-BASED-SM-MIB:SNMP-USER-BASED-SM-MIB/" + usmStats.EntityData.SegmentPath
    usmStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usmStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usmStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usmStats.EntityData.Children = types.NewOrderedMap()
    usmStats.EntityData.Leafs = types.NewOrderedMap()
    usmStats.EntityData.Leafs.Append("usmStatsUnsupportedSecLevels", types.YLeaf{"UsmStatsUnsupportedSecLevels", usmStats.UsmStatsUnsupportedSecLevels})
    usmStats.EntityData.Leafs.Append("usmStatsNotInTimeWindows", types.YLeaf{"UsmStatsNotInTimeWindows", usmStats.UsmStatsNotInTimeWindows})
    usmStats.EntityData.Leafs.Append("usmStatsUnknownUserNames", types.YLeaf{"UsmStatsUnknownUserNames", usmStats.UsmStatsUnknownUserNames})
    usmStats.EntityData.Leafs.Append("usmStatsUnknownEngineIDs", types.YLeaf{"UsmStatsUnknownEngineIDs", usmStats.UsmStatsUnknownEngineIDs})
    usmStats.EntityData.Leafs.Append("usmStatsWrongDigests", types.YLeaf{"UsmStatsWrongDigests", usmStats.UsmStatsWrongDigests})
    usmStats.EntityData.Leafs.Append("usmStatsDecryptionErrors", types.YLeaf{"UsmStatsDecryptionErrors", usmStats.UsmStatsDecryptionErrors})

    usmStats.EntityData.YListKeys = []string {}

    return &(usmStats.EntityData)
}

// SNMPUSERBASEDSMMIB_UsmUserTable
type SNMPUSERBASEDSMMIB_UsmUserTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SNMPUSERBASEDSMMIB_UsmUserTable_UsmUserEntry.
    UsmUserEntry []*SNMPUSERBASEDSMMIB_UsmUserTable_UsmUserEntry
}

func (usmUserTable *SNMPUSERBASEDSMMIB_UsmUserTable) GetEntityData() *types.CommonEntityData {
    usmUserTable.EntityData.YFilter = usmUserTable.YFilter
    usmUserTable.EntityData.YangName = "usmUserTable"
    usmUserTable.EntityData.BundleName = "cisco_ios_xr"
    usmUserTable.EntityData.ParentYangName = "SNMP-USER-BASED-SM-MIB"
    usmUserTable.EntityData.SegmentPath = "usmUserTable"
    usmUserTable.EntityData.AbsolutePath = "SNMP-USER-BASED-SM-MIB:SNMP-USER-BASED-SM-MIB/" + usmUserTable.EntityData.SegmentPath
    usmUserTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usmUserTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usmUserTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usmUserTable.EntityData.Children = types.NewOrderedMap()
    usmUserTable.EntityData.Children.Append("usmUserEntry", types.YChild{"UsmUserEntry", nil})
    for i := range usmUserTable.UsmUserEntry {
        usmUserTable.EntityData.Children.Append(types.GetSegmentPath(usmUserTable.UsmUserEntry[i]), types.YChild{"UsmUserEntry", usmUserTable.UsmUserEntry[i]})
    }
    usmUserTable.EntityData.Leafs = types.NewOrderedMap()

    usmUserTable.EntityData.YListKeys = []string {}

    return &(usmUserTable.EntityData)
}

// SNMPUSERBASEDSMMIB_UsmUserTable_UsmUserEntry
type SNMPUSERBASEDSMMIB_UsmUserTable_UsmUserEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserEngineID interface{}

    // This attribute is a key. The type is string with length: 1..32.
    UsmUserName interface{}

    // The type is string with length: 0..255. This attribute is mandatory.
    UsmUserSecurityName interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    UsmUserCloneFrom interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    // The default value is 1.3.6.1.6.3.10.1.1.1.
    UsmUserAuthProtocol interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserAuthKeyChange interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserOwnAuthKeyChange interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    // The default value is 1.3.6.1.6.3.10.1.2.1.
    UsmUserPrivProtocol interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserPrivKeyChange interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserOwnPrivKeyChange interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserPublic interface{}

    // The type is StorageType. The default value is nonVolatile.
    UsmUserStorageType interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserAuthKey interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    UsmUserPrivKey interface{}
}

func (usmUserEntry *SNMPUSERBASEDSMMIB_UsmUserTable_UsmUserEntry) GetEntityData() *types.CommonEntityData {
    usmUserEntry.EntityData.YFilter = usmUserEntry.YFilter
    usmUserEntry.EntityData.YangName = "usmUserEntry"
    usmUserEntry.EntityData.BundleName = "cisco_ios_xr"
    usmUserEntry.EntityData.ParentYangName = "usmUserTable"
    usmUserEntry.EntityData.SegmentPath = "usmUserEntry" + types.AddKeyToken(usmUserEntry.UsmUserEngineID, "usmUserEngineID") + types.AddKeyToken(usmUserEntry.UsmUserName, "usmUserName")
    usmUserEntry.EntityData.AbsolutePath = "SNMP-USER-BASED-SM-MIB:SNMP-USER-BASED-SM-MIB/usmUserTable/" + usmUserEntry.EntityData.SegmentPath
    usmUserEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usmUserEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usmUserEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usmUserEntry.EntityData.Children = types.NewOrderedMap()
    usmUserEntry.EntityData.Leafs = types.NewOrderedMap()
    usmUserEntry.EntityData.Leafs.Append("usmUserEngineID", types.YLeaf{"UsmUserEngineID", usmUserEntry.UsmUserEngineID})
    usmUserEntry.EntityData.Leafs.Append("usmUserName", types.YLeaf{"UsmUserName", usmUserEntry.UsmUserName})
    usmUserEntry.EntityData.Leafs.Append("usmUserSecurityName", types.YLeaf{"UsmUserSecurityName", usmUserEntry.UsmUserSecurityName})
    usmUserEntry.EntityData.Leafs.Append("usmUserCloneFrom", types.YLeaf{"UsmUserCloneFrom", usmUserEntry.UsmUserCloneFrom})
    usmUserEntry.EntityData.Leafs.Append("usmUserAuthProtocol", types.YLeaf{"UsmUserAuthProtocol", usmUserEntry.UsmUserAuthProtocol})
    usmUserEntry.EntityData.Leafs.Append("usmUserAuthKeyChange", types.YLeaf{"UsmUserAuthKeyChange", usmUserEntry.UsmUserAuthKeyChange})
    usmUserEntry.EntityData.Leafs.Append("usmUserOwnAuthKeyChange", types.YLeaf{"UsmUserOwnAuthKeyChange", usmUserEntry.UsmUserOwnAuthKeyChange})
    usmUserEntry.EntityData.Leafs.Append("usmUserPrivProtocol", types.YLeaf{"UsmUserPrivProtocol", usmUserEntry.UsmUserPrivProtocol})
    usmUserEntry.EntityData.Leafs.Append("usmUserPrivKeyChange", types.YLeaf{"UsmUserPrivKeyChange", usmUserEntry.UsmUserPrivKeyChange})
    usmUserEntry.EntityData.Leafs.Append("usmUserOwnPrivKeyChange", types.YLeaf{"UsmUserOwnPrivKeyChange", usmUserEntry.UsmUserOwnPrivKeyChange})
    usmUserEntry.EntityData.Leafs.Append("usmUserPublic", types.YLeaf{"UsmUserPublic", usmUserEntry.UsmUserPublic})
    usmUserEntry.EntityData.Leafs.Append("usmUserStorageType", types.YLeaf{"UsmUserStorageType", usmUserEntry.UsmUserStorageType})
    usmUserEntry.EntityData.Leafs.Append("usmUserAuthKey", types.YLeaf{"UsmUserAuthKey", usmUserEntry.UsmUserAuthKey})
    usmUserEntry.EntityData.Leafs.Append("usmUserPrivKey", types.YLeaf{"UsmUserPrivKey", usmUserEntry.UsmUserPrivKey})

    usmUserEntry.EntityData.YListKeys = []string {"UsmUserEngineID", "UsmUserName"}

    return &(usmUserEntry.EntityData)
}

