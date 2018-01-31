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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscoubemibobjects CISCOUBEMIB_Ciscoubemibobjects
}

func (cISCOUBEMIB *CISCOUBEMIB) GetFilter() yfilter.YFilter { return cISCOUBEMIB.YFilter }

func (cISCOUBEMIB *CISCOUBEMIB) SetFilter(yf yfilter.YFilter) { cISCOUBEMIB.YFilter = yf }

func (cISCOUBEMIB *CISCOUBEMIB) GetGoName(yname string) string {
    if yname == "ciscoUbeMIBObjects" { return "Ciscoubemibobjects" }
    return ""
}

func (cISCOUBEMIB *CISCOUBEMIB) GetSegmentPath() string {
    return "CISCO-UBE-MIB:CISCO-UBE-MIB"
}

func (cISCOUBEMIB *CISCOUBEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoUbeMIBObjects" {
        return &cISCOUBEMIB.Ciscoubemibobjects
    }
    return nil
}

func (cISCOUBEMIB *CISCOUBEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoUbeMIBObjects"] = &cISCOUBEMIB.Ciscoubemibobjects
    return children
}

func (cISCOUBEMIB *CISCOUBEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOUBEMIB *CISCOUBEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOUBEMIB *CISCOUBEMIB) GetYangName() string { return "CISCO-UBE-MIB" }

func (cISCOUBEMIB *CISCOUBEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOUBEMIB *CISCOUBEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOUBEMIB *CISCOUBEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOUBEMIB *CISCOUBEMIB) SetParent(parent types.Entity) { cISCOUBEMIB.parent = parent }

func (cISCOUBEMIB *CISCOUBEMIB) GetParent() types.Entity { return cISCOUBEMIB.parent }

func (cISCOUBEMIB *CISCOUBEMIB) GetParentYangName() string { return "CISCO-UBE-MIB" }

// CISCOUBEMIB_Ciscoubemibobjects
type CISCOUBEMIB_Ciscoubemibobjects struct {
    parent types.Entity
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

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetFilter() yfilter.YFilter { return ciscoubemibobjects.YFilter }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) SetFilter(yf yfilter.YFilter) { ciscoubemibobjects.YFilter = yf }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetGoName(yname string) string {
    if yname == "cubeEnabled" { return "Cubeenabled" }
    if yname == "cubeVersion" { return "Cubeversion" }
    if yname == "cubeTotalSessionAllowed" { return "Cubetotalsessionallowed" }
    return ""
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetSegmentPath() string {
    return "ciscoUbeMIBObjects"
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cubeEnabled"] = ciscoubemibobjects.Cubeenabled
    leafs["cubeVersion"] = ciscoubemibobjects.Cubeversion
    leafs["cubeTotalSessionAllowed"] = ciscoubemibobjects.Cubetotalsessionallowed
    return leafs
}

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetYangName() string { return "ciscoUbeMIBObjects" }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) SetParent(parent types.Entity) { ciscoubemibobjects.parent = parent }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetParent() types.Entity { return ciscoubemibobjects.parent }

func (ciscoubemibobjects *CISCOUBEMIB_Ciscoubemibobjects) GetParentYangName() string { return "CISCO-UBE-MIB" }

