// This MIB describes objects used for managing Cisco
// Unified Border Element (CUBE).
// 
// The Cisco Unified Border Element (CUBE) is a Cisco 
// IOS Session Border Controller (SBC) that interconnects
// independent voice over IP (VoIP) and video over IP 
// networks for data, voice, and video transport
package cisco_ube_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ube_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-UBE-MIB CISCO-UBE-MIB}", reflect.TypeOf(CISCOUBEMIB{}))
    ydk.RegisterEntity("CISCO-UBE-MIB:CISCO-UBE-MIB", reflect.TypeOf(CISCOUBEMIB{}))
}

// CISCOUBEMIB
type CISCOUBEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoUbeMIBObjects CISCOUBEMIB_CiscoUbeMIBObjects
}

func (cISCOUBEMIB *CISCOUBEMIB) GetEntityData() *types.CommonEntityData {
    cISCOUBEMIB.EntityData.YFilter = cISCOUBEMIB.YFilter
    cISCOUBEMIB.EntityData.YangName = "CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOUBEMIB.EntityData.ParentYangName = "CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.SegmentPath = "CISCO-UBE-MIB:CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.AbsolutePath = cISCOUBEMIB.EntityData.SegmentPath
    cISCOUBEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOUBEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOUBEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOUBEMIB.EntityData.Children = types.NewOrderedMap()
    cISCOUBEMIB.EntityData.Children.Append("ciscoUbeMIBObjects", types.YChild{"CiscoUbeMIBObjects", &cISCOUBEMIB.CiscoUbeMIBObjects})
    cISCOUBEMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOUBEMIB.EntityData.YListKeys = []string {}

    return &(cISCOUBEMIB.EntityData)
}

// CISCOUBEMIB_CiscoUbeMIBObjects
type CISCOUBEMIB_CiscoUbeMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object represents, whether the Cisco Unified Border Element (CUBE) is
    // enabled  on the device or not.  The value 'true' means that the CUBE
    // feature  is enabled on the device.  The value 'false' means that the CUBE
    // feature  is disabled. The type is bool.
    CubeEnabled interface{}

    // This object represents the version of Cisco Unified Border Element on the
    // device. The type is string.
    CubeVersion interface{}

    // This object provides the total number of CUBE session allowed on the
    // device. The value zero  means no sessions are allowed with CUBE. The type
    // is interface{} with range: 0..999999. Units are session.
    CubeTotalSessionAllowed interface{}
}

func (ciscoUbeMIBObjects *CISCOUBEMIB_CiscoUbeMIBObjects) GetEntityData() *types.CommonEntityData {
    ciscoUbeMIBObjects.EntityData.YFilter = ciscoUbeMIBObjects.YFilter
    ciscoUbeMIBObjects.EntityData.YangName = "ciscoUbeMIBObjects"
    ciscoUbeMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoUbeMIBObjects.EntityData.ParentYangName = "CISCO-UBE-MIB"
    ciscoUbeMIBObjects.EntityData.SegmentPath = "ciscoUbeMIBObjects"
    ciscoUbeMIBObjects.EntityData.AbsolutePath = "CISCO-UBE-MIB:CISCO-UBE-MIB/" + ciscoUbeMIBObjects.EntityData.SegmentPath
    ciscoUbeMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoUbeMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoUbeMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoUbeMIBObjects.EntityData.Children = types.NewOrderedMap()
    ciscoUbeMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoUbeMIBObjects.EntityData.Leafs.Append("cubeEnabled", types.YLeaf{"CubeEnabled", ciscoUbeMIBObjects.CubeEnabled})
    ciscoUbeMIBObjects.EntityData.Leafs.Append("cubeVersion", types.YLeaf{"CubeVersion", ciscoUbeMIBObjects.CubeVersion})
    ciscoUbeMIBObjects.EntityData.Leafs.Append("cubeTotalSessionAllowed", types.YLeaf{"CubeTotalSessionAllowed", ciscoUbeMIBObjects.CubeTotalSessionAllowed})

    ciscoUbeMIBObjects.EntityData.YListKeys = []string {}

    return &(ciscoUbeMIBObjects.EntityData)
}

