// The MIB module for managing Power Source Equipment
// (PSE) working according to the IEEE 802.af Powered
// Ethernet (DTE Power via MDI) standard.
// 
//  The following terms are used throughout this
//  MIB module.  For complete formal definitions,
//  the IEEE 802.3 standards should be consulted
//  wherever possible:
// 
//  Group - A recommended, but optional, entity
//  defined by the IEEE 802.3 management standard,
//  in order to support a modular numbering scheme.
//  The classical example allows an implementor to
//  represent field-replaceable units as groups of
//  ports, with the port numbering matching the
//  modular hardware implementation.
// 
// Port - This entity identifies the port within the group
// for which this entry contains information.  The numbering
// scheme for ports is implementation specific.
// 
// Copyright (c) The Internet Society (2003).  This version
// of this MIB module is part of RFC 3621; See the RFC
// itself for full legal notices.
package power_ethernet_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package power_ethernet_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:POWER-ETHERNET-MIB POWER-ETHERNET-MIB}", reflect.TypeOf(POWERETHERNETMIB{}))
    ydk.RegisterEntity("POWER-ETHERNET-MIB:POWER-ETHERNET-MIB", reflect.TypeOf(POWERETHERNETMIB{}))
}

// POWERETHERNETMIB
type POWERETHERNETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table of objects that display and control the power characteristics of
    // power Ethernet ports on a Power Source Entity (PSE) device.  This group
    // will be implemented in managed power Ethernet switches and mid-span
    // devices. Values of all read-write objects in this table are persistent at
    // restart/reboot.
    PethPsePortTable POWERETHERNETMIB_PethPsePortTable

    // A table of objects that display and control attributes of the main power
    // source in a PSE  device.  Ethernet switches are one example of boxes that
    // would support these objects. Values of all read-write objects in this table
    // are persistent at restart/reboot.
    PethMainPseTable POWERETHERNETMIB_PethMainPseTable

    // A table of objects that display and control the Notification on a PSE 
    // device. Values of all read-write objects in this table are persistent at
    // restart/reboot.
    PethNotificationControlTable POWERETHERNETMIB_PethNotificationControlTable
}

func (pOWERETHERNETMIB *POWERETHERNETMIB) GetEntityData() *types.CommonEntityData {
    pOWERETHERNETMIB.EntityData.YFilter = pOWERETHERNETMIB.YFilter
    pOWERETHERNETMIB.EntityData.YangName = "POWER-ETHERNET-MIB"
    pOWERETHERNETMIB.EntityData.BundleName = "cisco_ios_xe"
    pOWERETHERNETMIB.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pOWERETHERNETMIB.EntityData.SegmentPath = "POWER-ETHERNET-MIB:POWER-ETHERNET-MIB"
    pOWERETHERNETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pOWERETHERNETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pOWERETHERNETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pOWERETHERNETMIB.EntityData.Children = types.NewOrderedMap()
    pOWERETHERNETMIB.EntityData.Children.Append("pethPsePortTable", types.YChild{"PethPsePortTable", &pOWERETHERNETMIB.PethPsePortTable})
    pOWERETHERNETMIB.EntityData.Children.Append("pethMainPseTable", types.YChild{"PethMainPseTable", &pOWERETHERNETMIB.PethMainPseTable})
    pOWERETHERNETMIB.EntityData.Children.Append("pethNotificationControlTable", types.YChild{"PethNotificationControlTable", &pOWERETHERNETMIB.PethNotificationControlTable})
    pOWERETHERNETMIB.EntityData.Leafs = types.NewOrderedMap()

    pOWERETHERNETMIB.EntityData.YListKeys = []string {}

    return &(pOWERETHERNETMIB.EntityData)
}

// POWERETHERNETMIB_PethPsePortTable
// A table of objects that display and control the power
// characteristics of power Ethernet ports on a Power Source
// Entity (PSE) device.  This group will be implemented in
// managed power Ethernet switches and mid-span devices.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_PethPsePortTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that display and control the power characteristics of a
    // power Ethernet PSE port. The type is slice of
    // POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry.
    PethPsePortEntry []*POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry
}

func (pethPsePortTable *POWERETHERNETMIB_PethPsePortTable) GetEntityData() *types.CommonEntityData {
    pethPsePortTable.EntityData.YFilter = pethPsePortTable.YFilter
    pethPsePortTable.EntityData.YangName = "pethPsePortTable"
    pethPsePortTable.EntityData.BundleName = "cisco_ios_xe"
    pethPsePortTable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethPsePortTable.EntityData.SegmentPath = "pethPsePortTable"
    pethPsePortTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethPsePortTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethPsePortTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethPsePortTable.EntityData.Children = types.NewOrderedMap()
    pethPsePortTable.EntityData.Children.Append("pethPsePortEntry", types.YChild{"PethPsePortEntry", nil})
    for i := range pethPsePortTable.PethPsePortEntry {
        pethPsePortTable.EntityData.Children.Append(types.GetSegmentPath(pethPsePortTable.PethPsePortEntry[i]), types.YChild{"PethPsePortEntry", pethPsePortTable.PethPsePortEntry[i]})
    }
    pethPsePortTable.EntityData.Leafs = types.NewOrderedMap()

    pethPsePortTable.EntityData.YListKeys = []string {}

    return &(pethPsePortTable.EntityData)
}

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry
// A set of objects that display and control the power
// characteristics of a power Ethernet PSE port.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group
    // containing the port to which a power Ethernet PSE is connected.  Group
    // means box in the stack, module in a rack and the value 1 MUST be used for
    // non-modular devices. Furthermore, the same value MUST be used in this
    // variable, pethMainPseGroupIndex, and pethNotificationControlGroupIndex to
    // refer to a given box in a stack or module in the rack. The type is
    // interface{} with range: 1..2147483647.
    PethPsePortGroupIndex interface{}

    // This attribute is a key. This variable uniquely identifies the power
    // Ethernet PSE port within group pethPsePortGroupIndex to which the power
    // Ethernet PSE entry is connected. The type is interface{} with range:
    // 1..2147483647.
    PethPsePortIndex interface{}

    // true (1) An interface which can provide the PSE functions. false(2) The
    // interface will act as it would if it had no PSE function. The type is bool.
    PethPsePortAdminEnable interface{}

    // Describes the capability of controlling the power pairs functionality to
    // switch pins for sourcing power. The value true indicate that the device has
    // the capability to control the power pairs.  When false the PSE Pinout
    // Alternative used cannot be controlled through the PethPsePortAdminEnable
    // attribute. The type is bool.
    PethPsePortPowerPairsControlAbility interface{}

    // Describes or controls the pairs in use.  If the value of
    // pethPsePortPowerPairsControl is true, this object is writable. A value of
    // signal(1) means that the signal pairs only are in use. A value of spare(2)
    // means that the spare pairs only are in use. The type is
    // PethPsePortPowerPairs.
    PethPsePortPowerPairs interface{}

    // Describes the operational status of the port PD detection. A value of
    // disabled(1)- indicates that the PSE State diagram is in the state DISABLED.
    // A value of deliveringPower(3) - indicates that the PSE State diagram is in
    // the state POWER_ON for a duration greater than tlim max (see IEEE Std
    // 802.3af Table 33-5 tlim). A value of fault(4) - indicates that the PSE
    // State diagram is in the state TEST_ERROR. A value of test(5) - indicates
    // that the PSE State diagram is in the state TEST_MODE. A value of
    // otherFault(6) - indicates that the PSE State diagram is in the state IDLE
    // due to the variable error_conditions. A value of searching(2)- indicates
    // the PSE State diagram is in a state other than those listed above. The type
    // is PethPsePortDetectionStatus.
    PethPsePortDetectionStatus interface{}

    // This object controls the priority of the port from the point of view of a
    // power management algorithm.  The priority that is set by this variable
    // could be used by a control mechanism that prevents over current situations
    // by disconnecting first ports with lower power priority.  Ports that connect
    // devices critical to the operation of the network - like the E911 telephones
    // ports - should be set to higher priority. The type is
    // PethPsePortPowerPriority.
    PethPsePortPowerPriority interface{}

    // This counter is incremented when the PSE state diagram transitions directly
    // from the state POWER_ON to the state IDLE due to tmpdo_timer_done being
    // asserted. The type is interface{} with range: 0..4294967295.
    PethPsePortMPSAbsentCounter interface{}

    // A manager will set the value of this variable to indicate the type of
    // powered device that is connected to the port. The default value supplied by
    // the agent if no value has ever been set should be a zero-length octet
    // string. The type is string.
    PethPsePortType interface{}

    // Classification is a way to tag different terminals on the Power over LAN
    // network according to their power consumption. Devices such as IP
    // telephones, WLAN access points and others, will be classified according to
    // their power requirements.  The meaning of the classification labels is
    // defined in the IEEE specification.  This variable is valid only while a PD
    // is being powered, that is, while the attribute pethPsePortDetectionStatus
    // is reporting the enumeration deliveringPower. The type is
    // PethPsePortPowerClassifications.
    PethPsePortPowerClassifications interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // SIGNATURE_INVALID. The type is interface{} with range: 0..4294967295.
    PethPsePortInvalidSignatureCounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // POWER_DENIED. The type is interface{} with range: 0..4294967295.
    PethPsePortPowerDeniedCounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // ERROR_DELAY_OVER. The type is interface{} with range: 0..4294967295.
    PethPsePortOverLoadCounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // ERROR_DELAY_SHORT. The type is interface{} with range: 0..4294967295.
    PethPsePortShortCounter interface{}

    // This object is an extension of the pethPsePortAdminEnable object from
    // RFC3621. It allows the user to be more specific when enabling the PSE
    // functions. The states, 'auto', 'static' and 'limit' correspond to a value
    // of 'true' for the pethPsePortAdminEnable object. The state 'disable'
    // corresponds to a value of 'false' for the pethPsePortAdminEnable object. 
    // Setting this value to 'auto' enables Powered Device discovery on the
    // interface and the amount of power allocated depends on the Powered Device
    // discovered. If pethPsePortAdminEnable was 'false' prior to this set
    // operation, then it will become 'true'.  Setting this value to 'static' will
    // also enable Powered Device discovery. However, this is different from
    // 'auto' in that the amount of power is pre-allocated based on the
    // configuration on the Power Sourcing Equipment. If pethPsePortAdminEnable
    // was 'false' prior to this set operation, then it will become 'true'. 
    // Setting this value to 'limit' enables Powered Device discovery on the
    // interface. The amount of power allocated depends on the Powered Device
    // discovered and the value of cpeExtPsePortPwrMax. The lower value will be
    // used. If pethPsePortAdminEnable was 'false' prior to this set operation,
    // then it will become 'true'.  Setting this value to 'disable' disables the
    // PSE functions. The pethPsePortAdminEnable object will adopt the value of
    // 'false' if it was 'true' prior to this set operation. When setting the
    // pethPsePortAdminEnable object to 'false' this object cpeExtPsePortEnable
    // will adopt the value of 'disable'.  If cpeExtPsePortPolicingCapable of the
    // PSE port, or cpeExtMainPsePwrMonitorCapable of the PSE port's main group,
    // has the value of 'false', this object can only be set to 'auto', 'static'
    // or 'disable'. Otherwise, this object can be set to 'auto', 'static',
    // 'limit' or 'disable'. The type is CpeExtPsePortEnable.
    CpeExtPsePortEnable interface{}

    // This object indicates the discover mode used by the system to discover the
    // PD.  A value of 'unknown' indicates that the discover mode on the interface
    // is unknown.  A value of 'off' indicates that discovery is disabled on the
    // interface.  A value of 'ieee' indicates that the discover mode on the
    // interface is IEEE based.  A value of 'cisco' indicates that the discover
    // mode on the interface is Cisco based.  A value of 'ieeeAndCisco' indicates
    // that the discover mode on the interface is both IEEE and Cisco based. The
    // type is CpeExtPsePortDiscoverMode.
    CpeExtPsePortDiscoverMode interface{}

    // This object indicates if a Powered Device (PD) has been detected on the
    // interface.  A value of 'true' indicates that a PD has been detected on the
    // interface.  A value of 'false' indicates that no PD has been detected on
    // the interface. The type is bool.
    CpeExtPsePortDeviceDetected interface{}

    // This object indicates whether the Powered Device attached to the interface
    // is an IEEE compliant Powered Device or not.  A value of 'true' indicates
    // the attached Powered Device is an IEEE compliant Powered Device.  A value
    // of 'false' indicates the attached Powered Device is not an IEEE compliant
    // Powered Device. This also means that the value of the corresponding object
    // from the pethPsePortTable, pethPsePortPowerClassifications is irrelevant.
    // The type is bool.
    CpeExtPsePortIeeePd interface{}

    // This object is an extension of the pethPsePortDetectionStatus object from
    // RFC3621 and provides additional status information.  deny: When set, the PD
    // attached to the interface is being       denied power due to insufficient
    // power resources on       the Power Sourcing Equipment.  overdraw: When set,
    // the PD attached to the interface is            being denied power because
    // the PD is trying            to consume more power than it has been         
    // configured to consume.  overdrawLog: When set, the PD attached to the
    // interface               is trying to consume more power than it has        
    // been configured to consume, but is not being               denied power.
    // The type is map[string]bool.
    CpeExtPsePortAdditionalStatus interface{}

    // This indicates the maximum amount of power that the PSE will make available
    // to the PD connected to this interface. The type is interface{} with range:
    // 0..4294967295. Units are milliwatts.
    CpeExtPsePortPwrMax interface{}

    // This object indicates the amount of power allocated from the PSE for the
    // PD. The type is interface{} with range: 0..4294967295. Units are
    // milliwatts.
    CpeExtPsePortPwrAllocated interface{}

    // This object indicates the amount of power available for the PD to use. This
    // value may differ from the value cpeExtPsePortPwrAllocated due to the
    // efficiency issues of delivering the power from the PSE to the PD.  When
    // sufficient power is available to power a PD, this value should be equal to
    // the lower of the two objects, cpeExtDefaultAllocation and
    // cpeExtPsePortPwrMax. The type is interface{} with range: 0..4294967295.
    // Units are milliwatts.
    CpeExtPsePortPwrAvailable interface{}

    // This indicates the actual power consumption of the PD connected to this
    // interface. It may not necessarily be equal to the value of
    // cpeExtPsePortPwrAvailable. The type is interface{} with range:
    // 0..4294967295. Units are milliwatts.
    CpeExtPsePortPwrConsumption interface{}

    // This indicates the maximum amount of power drawn by the PD connected to
    // this interface, since it was powered on. The type is interface{} with
    // range: 0..4294967295. Units are milliwatts.
    CpeExtPsePortMaxPwrDrawn interface{}

    // The entPhysicalIndex value that uniquely identifies the PSE port. If the
    // PSE port does not have a corresponding  physical entry in entPhysicalTable
    // or if the  entPhysicalTable is not supported by the management system, this
    // object has the value of zero. The type is interface{} with range:
    // 0..2147483647.
    CpeExtPsePortEntPhyIndex interface{}

    // This object indicates whether the PSE port hardware is capable of policing
    // the port for proper power consumption  based on the allocated value. The
    // type is bool.
    CpeExtPsePortPolicingCapable interface{}

    // This object allows the user to turn on or turn off the power policing of
    // the PSE port. If the instance value of  cpeExtPsePortPolicingCapable is
    // 'TRUE', the user is allowed to set this object to 'on' or 'off'. Otherwise,
    // this object is read-only and always has the value of 'off'. The type is
    // CpeExtPsePortPolicingEnable.
    CpeExtPsePortPolicingEnable interface{}

    // This object specifies the power policing action that the device will take
    // on this PSE port when the real-time  power consumption exceeds its max
    // power allocation if  the value of cpeExtPsePortPolicingEnable is 'on'.     
    // 'deny'          - the device will deny the power to                        
    // the PSE port       'logOnly'       - the device will not deny the power    
    // to the PSE port. The type is CpeExtPsePortPolicingAction.
    CpeExtPsePortPolicingAction interface{}

    // This object specifies the manual power allocation that the PSE will
    // allocate to the PD connected to this  interface regardless of the amount
    // requested via CDP  or IEEE.   Setting this object value to zero disables
    // the manual  power allocation.   Warning: Misconfiguring this manual power
    // allocation may  cause damage to the system and void the warranty. Take 
    // precautions not to oversubscribe the power supply. The type is interface{}
    // with range: 0..4294967295. Units are milliwatts.
    CpeExtPsePortPwrManAlloc interface{}

    // This object indicates the PSE functionality that this port supports.  If
    // the 'policing' BIT is set, then this PSE port is capable of policing the
    // port for proper power consumption based on the allocated value.  If the
    // 'poePlus' BIT is set, then this PSE port supports PoE Plus functions. The
    // type is map[string]bool.
    CpeExtPsePortCapabilities interface{}
}

func (pethPsePortEntry *POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry) GetEntityData() *types.CommonEntityData {
    pethPsePortEntry.EntityData.YFilter = pethPsePortEntry.YFilter
    pethPsePortEntry.EntityData.YangName = "pethPsePortEntry"
    pethPsePortEntry.EntityData.BundleName = "cisco_ios_xe"
    pethPsePortEntry.EntityData.ParentYangName = "pethPsePortTable"
    pethPsePortEntry.EntityData.SegmentPath = "pethPsePortEntry" + types.AddKeyToken(pethPsePortEntry.PethPsePortGroupIndex, "pethPsePortGroupIndex") + types.AddKeyToken(pethPsePortEntry.PethPsePortIndex, "pethPsePortIndex")
    pethPsePortEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethPsePortEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethPsePortEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethPsePortEntry.EntityData.Children = types.NewOrderedMap()
    pethPsePortEntry.EntityData.Leafs = types.NewOrderedMap()
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortGroupIndex", types.YLeaf{"PethPsePortGroupIndex", pethPsePortEntry.PethPsePortGroupIndex})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortIndex", types.YLeaf{"PethPsePortIndex", pethPsePortEntry.PethPsePortIndex})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortAdminEnable", types.YLeaf{"PethPsePortAdminEnable", pethPsePortEntry.PethPsePortAdminEnable})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortPowerPairsControlAbility", types.YLeaf{"PethPsePortPowerPairsControlAbility", pethPsePortEntry.PethPsePortPowerPairsControlAbility})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortPowerPairs", types.YLeaf{"PethPsePortPowerPairs", pethPsePortEntry.PethPsePortPowerPairs})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortDetectionStatus", types.YLeaf{"PethPsePortDetectionStatus", pethPsePortEntry.PethPsePortDetectionStatus})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortPowerPriority", types.YLeaf{"PethPsePortPowerPriority", pethPsePortEntry.PethPsePortPowerPriority})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortMPSAbsentCounter", types.YLeaf{"PethPsePortMPSAbsentCounter", pethPsePortEntry.PethPsePortMPSAbsentCounter})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortType", types.YLeaf{"PethPsePortType", pethPsePortEntry.PethPsePortType})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortPowerClassifications", types.YLeaf{"PethPsePortPowerClassifications", pethPsePortEntry.PethPsePortPowerClassifications})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortInvalidSignatureCounter", types.YLeaf{"PethPsePortInvalidSignatureCounter", pethPsePortEntry.PethPsePortInvalidSignatureCounter})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortPowerDeniedCounter", types.YLeaf{"PethPsePortPowerDeniedCounter", pethPsePortEntry.PethPsePortPowerDeniedCounter})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortOverLoadCounter", types.YLeaf{"PethPsePortOverLoadCounter", pethPsePortEntry.PethPsePortOverLoadCounter})
    pethPsePortEntry.EntityData.Leafs.Append("pethPsePortShortCounter", types.YLeaf{"PethPsePortShortCounter", pethPsePortEntry.PethPsePortShortCounter})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortEnable", types.YLeaf{"CpeExtPsePortEnable", pethPsePortEntry.CpeExtPsePortEnable})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortDiscoverMode", types.YLeaf{"CpeExtPsePortDiscoverMode", pethPsePortEntry.CpeExtPsePortDiscoverMode})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortDeviceDetected", types.YLeaf{"CpeExtPsePortDeviceDetected", pethPsePortEntry.CpeExtPsePortDeviceDetected})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortIeeePd", types.YLeaf{"CpeExtPsePortIeeePd", pethPsePortEntry.CpeExtPsePortIeeePd})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortAdditionalStatus", types.YLeaf{"CpeExtPsePortAdditionalStatus", pethPsePortEntry.CpeExtPsePortAdditionalStatus})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPwrMax", types.YLeaf{"CpeExtPsePortPwrMax", pethPsePortEntry.CpeExtPsePortPwrMax})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPwrAllocated", types.YLeaf{"CpeExtPsePortPwrAllocated", pethPsePortEntry.CpeExtPsePortPwrAllocated})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPwrAvailable", types.YLeaf{"CpeExtPsePortPwrAvailable", pethPsePortEntry.CpeExtPsePortPwrAvailable})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPwrConsumption", types.YLeaf{"CpeExtPsePortPwrConsumption", pethPsePortEntry.CpeExtPsePortPwrConsumption})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortMaxPwrDrawn", types.YLeaf{"CpeExtPsePortMaxPwrDrawn", pethPsePortEntry.CpeExtPsePortMaxPwrDrawn})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortEntPhyIndex", types.YLeaf{"CpeExtPsePortEntPhyIndex", pethPsePortEntry.CpeExtPsePortEntPhyIndex})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPolicingCapable", types.YLeaf{"CpeExtPsePortPolicingCapable", pethPsePortEntry.CpeExtPsePortPolicingCapable})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPolicingEnable", types.YLeaf{"CpeExtPsePortPolicingEnable", pethPsePortEntry.CpeExtPsePortPolicingEnable})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPolicingAction", types.YLeaf{"CpeExtPsePortPolicingAction", pethPsePortEntry.CpeExtPsePortPolicingAction})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortPwrManAlloc", types.YLeaf{"CpeExtPsePortPwrManAlloc", pethPsePortEntry.CpeExtPsePortPwrManAlloc})
    pethPsePortEntry.EntityData.Leafs.Append("cpeExtPsePortCapabilities", types.YLeaf{"CpeExtPsePortCapabilities", pethPsePortEntry.CpeExtPsePortCapabilities})

    pethPsePortEntry.EntityData.YListKeys = []string {"PethPsePortGroupIndex", "PethPsePortIndex"}

    return &(pethPsePortEntry.EntityData)
}

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode represents the interface is both IEEE and Cisco based.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode_unknown POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode = "unknown"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode_off POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode = "off"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode_ieee POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode = "ieee"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode_cisco POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode = "cisco"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode_ieeeAndCisco POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortDiscoverMode = "ieeeAndCisco"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable represents 'limit' or 'disable'.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable_auto POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable = "auto"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable_static POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable = "static"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable_limit POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable = "limit"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable_disable POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortEnable = "disable"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction represents                        to the PSE port
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction_deny POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction = "deny"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction_logOnly POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingAction = "logOnly"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable represents object is read-only and always has the value of 'off'.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable_on POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable = "on"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable_off POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_CpeExtPsePortPolicingEnable = "off"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus represents in a state other than those listed above.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_disabled POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "disabled"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_searching POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "searching"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_deliveringPower POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "deliveringPower"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_fault POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "fault"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_test POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "test"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus_otherFault POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortDetectionStatus = "otherFault"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications represents is reporting the enumeration deliveringPower.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications_class0 POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications = "class0"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications_class1 POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications = "class1"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications_class2 POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications = "class2"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications_class3 POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications = "class3"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications_class4 POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerClassifications = "class4"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs represents only are in use.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs_signal POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs = "signal"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs_spare POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPairs = "spare"
)

// POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority represents telephones ports - should be set to higher priority.
type POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority string

const (
    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority_critical POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority = "critical"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority_high POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority = "high"

    POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority_low POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortPowerPriority = "low"
)

// POWERETHERNETMIB_PethMainPseTable
// A table of objects that display and control attributes
// of the main power source in a PSE  device.  Ethernet
// switches are one example of boxes that would support
// these objects.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_PethMainPseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that display and control the Main power of a PSE. . The
    // type is slice of POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry.
    PethMainPseEntry []*POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry
}

func (pethMainPseTable *POWERETHERNETMIB_PethMainPseTable) GetEntityData() *types.CommonEntityData {
    pethMainPseTable.EntityData.YFilter = pethMainPseTable.YFilter
    pethMainPseTable.EntityData.YangName = "pethMainPseTable"
    pethMainPseTable.EntityData.BundleName = "cisco_ios_xe"
    pethMainPseTable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethMainPseTable.EntityData.SegmentPath = "pethMainPseTable"
    pethMainPseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethMainPseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethMainPseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethMainPseTable.EntityData.Children = types.NewOrderedMap()
    pethMainPseTable.EntityData.Children.Append("pethMainPseEntry", types.YChild{"PethMainPseEntry", nil})
    for i := range pethMainPseTable.PethMainPseEntry {
        pethMainPseTable.EntityData.Children.Append(types.GetSegmentPath(pethMainPseTable.PethMainPseEntry[i]), types.YChild{"PethMainPseEntry", pethMainPseTable.PethMainPseEntry[i]})
    }
    pethMainPseTable.EntityData.Leafs = types.NewOrderedMap()

    pethMainPseTable.EntityData.YListKeys = []string {}

    return &(pethMainPseTable.EntityData)
}

// POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry
// A set of objects that display and control the Main
// power of a PSE. 
type POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group to
    // which power Ethernet PSE is connected.  Group means (box in the stack,
    // module in a rack) and the value 1 MUST be used for non-modular devices. 
    // Furthermore, the same value MUST be used in this variable,
    // pethPsePortGroupIndex, and pethNotificationControlGroupIndex to refer to a
    // given box in a stack or module in a rack. The type is interface{} with
    // range: 1..2147483647.
    PethMainPseGroupIndex interface{}

    // The nominal power of the PSE expressed in Watts. The type is interface{}
    // with range: 1..65535. Units are Watts.
    PethMainPsePower interface{}

    // The operational status of the main PSE. The type is PethMainPseOperStatus.
    PethMainPseOperStatus interface{}

    // Measured usage power expressed in Watts. The type is interface{} with
    // range: 0..4294967295. Units are Watts.
    PethMainPseConsumptionPower interface{}

    // The usage threshold expressed in percents for comparing the measured power
    // and initiating an alarm if the threshold is exceeded. The type is
    // interface{} with range: 1..99. Units are %.
    PethMainPseUsageThreshold interface{}
}

func (pethMainPseEntry *POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry) GetEntityData() *types.CommonEntityData {
    pethMainPseEntry.EntityData.YFilter = pethMainPseEntry.YFilter
    pethMainPseEntry.EntityData.YangName = "pethMainPseEntry"
    pethMainPseEntry.EntityData.BundleName = "cisco_ios_xe"
    pethMainPseEntry.EntityData.ParentYangName = "pethMainPseTable"
    pethMainPseEntry.EntityData.SegmentPath = "pethMainPseEntry" + types.AddKeyToken(pethMainPseEntry.PethMainPseGroupIndex, "pethMainPseGroupIndex")
    pethMainPseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethMainPseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethMainPseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethMainPseEntry.EntityData.Children = types.NewOrderedMap()
    pethMainPseEntry.EntityData.Leafs = types.NewOrderedMap()
    pethMainPseEntry.EntityData.Leafs.Append("pethMainPseGroupIndex", types.YLeaf{"PethMainPseGroupIndex", pethMainPseEntry.PethMainPseGroupIndex})
    pethMainPseEntry.EntityData.Leafs.Append("pethMainPsePower", types.YLeaf{"PethMainPsePower", pethMainPseEntry.PethMainPsePower})
    pethMainPseEntry.EntityData.Leafs.Append("pethMainPseOperStatus", types.YLeaf{"PethMainPseOperStatus", pethMainPseEntry.PethMainPseOperStatus})
    pethMainPseEntry.EntityData.Leafs.Append("pethMainPseConsumptionPower", types.YLeaf{"PethMainPseConsumptionPower", pethMainPseEntry.PethMainPseConsumptionPower})
    pethMainPseEntry.EntityData.Leafs.Append("pethMainPseUsageThreshold", types.YLeaf{"PethMainPseUsageThreshold", pethMainPseEntry.PethMainPseUsageThreshold})

    pethMainPseEntry.EntityData.YListKeys = []string {"PethMainPseGroupIndex"}

    return &(pethMainPseEntry.EntityData)
}

// POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus represents The operational status of the main PSE.
type POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus string

const (
    POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus_on POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus = "on"

    POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus_off POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus = "off"

    POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus_faulty POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseOperStatus = "faulty"
)

// POWERETHERNETMIB_PethNotificationControlTable
// A table of objects that display and control the
// Notification on a PSE  device.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_PethNotificationControlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that control the Notification events. The type is slice of
    // POWERETHERNETMIB_PethNotificationControlTable_PethNotificationControlEntry.
    PethNotificationControlEntry []*POWERETHERNETMIB_PethNotificationControlTable_PethNotificationControlEntry
}

func (pethNotificationControlTable *POWERETHERNETMIB_PethNotificationControlTable) GetEntityData() *types.CommonEntityData {
    pethNotificationControlTable.EntityData.YFilter = pethNotificationControlTable.YFilter
    pethNotificationControlTable.EntityData.YangName = "pethNotificationControlTable"
    pethNotificationControlTable.EntityData.BundleName = "cisco_ios_xe"
    pethNotificationControlTable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethNotificationControlTable.EntityData.SegmentPath = "pethNotificationControlTable"
    pethNotificationControlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethNotificationControlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethNotificationControlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethNotificationControlTable.EntityData.Children = types.NewOrderedMap()
    pethNotificationControlTable.EntityData.Children.Append("pethNotificationControlEntry", types.YChild{"PethNotificationControlEntry", nil})
    for i := range pethNotificationControlTable.PethNotificationControlEntry {
        pethNotificationControlTable.EntityData.Children.Append(types.GetSegmentPath(pethNotificationControlTable.PethNotificationControlEntry[i]), types.YChild{"PethNotificationControlEntry", pethNotificationControlTable.PethNotificationControlEntry[i]})
    }
    pethNotificationControlTable.EntityData.Leafs = types.NewOrderedMap()

    pethNotificationControlTable.EntityData.YListKeys = []string {}

    return &(pethNotificationControlTable.EntityData)
}

// POWERETHERNETMIB_PethNotificationControlTable_PethNotificationControlEntry
// A set of objects that control the Notification events.
type POWERETHERNETMIB_PethNotificationControlTable_PethNotificationControlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group. 
    // Group means box in the stack, module in a rack and the value 1 MUST be used
    // for non-modular devices.  Furthermore, the same value MUST be used in this
    // variable, pethPsePortGroupIndex, and pethMainPseGroupIndex to refer to a
    // given box in a stack or module in a rack. . The type is interface{} with
    // range: 1..2147483647.
    PethNotificationControlGroupIndex interface{}

    // This object controls, on a per-group basis, whether or not notifications
    // from the agent are enabled.  The value true(1) means that notifications are
    // enabled; the value false(2) means that they are not. The type is bool.
    PethNotificationControlEnable interface{}
}

func (pethNotificationControlEntry *POWERETHERNETMIB_PethNotificationControlTable_PethNotificationControlEntry) GetEntityData() *types.CommonEntityData {
    pethNotificationControlEntry.EntityData.YFilter = pethNotificationControlEntry.YFilter
    pethNotificationControlEntry.EntityData.YangName = "pethNotificationControlEntry"
    pethNotificationControlEntry.EntityData.BundleName = "cisco_ios_xe"
    pethNotificationControlEntry.EntityData.ParentYangName = "pethNotificationControlTable"
    pethNotificationControlEntry.EntityData.SegmentPath = "pethNotificationControlEntry" + types.AddKeyToken(pethNotificationControlEntry.PethNotificationControlGroupIndex, "pethNotificationControlGroupIndex")
    pethNotificationControlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethNotificationControlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethNotificationControlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethNotificationControlEntry.EntityData.Children = types.NewOrderedMap()
    pethNotificationControlEntry.EntityData.Leafs = types.NewOrderedMap()
    pethNotificationControlEntry.EntityData.Leafs.Append("pethNotificationControlGroupIndex", types.YLeaf{"PethNotificationControlGroupIndex", pethNotificationControlEntry.PethNotificationControlGroupIndex})
    pethNotificationControlEntry.EntityData.Leafs.Append("pethNotificationControlEnable", types.YLeaf{"PethNotificationControlEnable", pethNotificationControlEntry.PethNotificationControlEnable})

    pethNotificationControlEntry.EntityData.YListKeys = []string {"PethNotificationControlGroupIndex"}

    return &(pethNotificationControlEntry.EntityData)
}

