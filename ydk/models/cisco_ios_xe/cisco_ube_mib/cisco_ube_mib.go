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

    
    Ciscoubemibobjects CISCOUBEMIB_Ciscoubemibobjects
}

func (cISCOUBEMIB *CISCOUBEMIB) GetEntityData() *types.CommonEntityData {
    cISCOUBEMIB.EntityData.YFilter = cISCOUBEMIB.YFilter
    cISCOUBEMIB.EntityData.YangName = "CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOUBEMIB.EntityData.ParentYangName = "CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.SegmentPath = "CISCO-UBE-MIB:CISCO-UBE-MIB"
    cISCOUBEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOUBEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOUBEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOUBEMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOUBEMIB.EntityData.Children["ciscoUbeMIBObjects"] = types.YChild{"Ciscoubemibobjects", &cISCOUBEMIB.Ciscoubemibobjects}
    cISCOUBEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOUBEMIB.EntityData)
}

// CISCOUBEMIB_Ciscoubemibobjects
type CISCOUBEMIB_Ciscoubemibobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object represents, whether the Cisco Unified Border Element (CUBE) is
    // enabled  on the device or not.  The value 'true' means that the CUBE
    // feature  is enabled on the device.  The value 'false' means that the CUBE
    // feature  is disabled. The type is bool.
    Cubeenabled interface{}

    // This object represents the version of Cisco Unified Border Element on the
    // device. The type is string.
    Cubeversion interface{}

    // This object provides the total number of CUBE session allowed on the
    // device. The value zero  means no sessions are allowed with CUBE. The type
    // is interface{} with range: 0..999999. Units are session.
    Cubetotalsessionallowed interface{}
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetEntityData() *types.CommonEntityData {
    ciscoubemibobjects.EntityData.YFilter = ciscoubemibobjects.YFilter
    ciscoubemibobjects.EntityData.YangName = "ciscoUbeMIBObjects"
    ciscoubemibobjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoubemibobjects.EntityData.ParentYangName = "CISCO-UBE-MIB"
    ciscoubemibobjects.EntityData.SegmentPath = "ciscoUbeMIBObjects"
    ciscoubemibobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoubemibobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoubemibobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoubemibobjects.EntityData.Children = make(map[string]types.YChild)
    ciscoubemibobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoubemibobjects.EntityData.Leafs["cubeEnabled"] = types.YLeaf{"Cubeenabled", ciscoubemibobjects.Cubeenabled}
    ciscoubemibobjects.EntityData.Leafs["cubeVersion"] = types.YLeaf{"Cubeversion", ciscoubemibobjects.Cubeversion}
    ciscoubemibobjects.EntityData.Leafs["cubeTotalSessionAllowed"] = types.YLeaf{"Cubetotalsessionallowed", ciscoubemibobjects.Cubetotalsessionallowed}
    return &(ciscoubemibobjects.EntityData)
}

