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
    Pethpseporttable POWERETHERNETMIB_Pethpseporttable

    // A table of objects that display and control attributes of the main power
    // source in a PSE  device.  Ethernet switches are one example of boxes that
    // would support these objects. Values of all read-write objects in this table
    // are persistent at restart/reboot.
    Pethmainpsetable POWERETHERNETMIB_Pethmainpsetable

    // A table of objects that display and control the Notification on a PSE 
    // device. Values of all read-write objects in this table are persistent at
    // restart/reboot.
    Pethnotificationcontroltable POWERETHERNETMIB_Pethnotificationcontroltable
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

    pOWERETHERNETMIB.EntityData.Children = make(map[string]types.YChild)
    pOWERETHERNETMIB.EntityData.Children["pethPsePortTable"] = types.YChild{"Pethpseporttable", &pOWERETHERNETMIB.Pethpseporttable}
    pOWERETHERNETMIB.EntityData.Children["pethMainPseTable"] = types.YChild{"Pethmainpsetable", &pOWERETHERNETMIB.Pethmainpsetable}
    pOWERETHERNETMIB.EntityData.Children["pethNotificationControlTable"] = types.YChild{"Pethnotificationcontroltable", &pOWERETHERNETMIB.Pethnotificationcontroltable}
    pOWERETHERNETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pOWERETHERNETMIB.EntityData)
}

// POWERETHERNETMIB_Pethpseporttable
// A table of objects that display and control the power
// characteristics of power Ethernet ports on a Power Source
// Entity (PSE) device.  This group will be implemented in
// managed power Ethernet switches and mid-span devices.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_Pethpseporttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that display and control the power characteristics of a
    // power Ethernet PSE port. The type is slice of
    // POWERETHERNETMIB_Pethpseporttable_Pethpseportentry.
    Pethpseportentry []POWERETHERNETMIB_Pethpseporttable_Pethpseportentry
}

func (pethpseporttable *POWERETHERNETMIB_Pethpseporttable) GetEntityData() *types.CommonEntityData {
    pethpseporttable.EntityData.YFilter = pethpseporttable.YFilter
    pethpseporttable.EntityData.YangName = "pethPsePortTable"
    pethpseporttable.EntityData.BundleName = "cisco_ios_xe"
    pethpseporttable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethpseporttable.EntityData.SegmentPath = "pethPsePortTable"
    pethpseporttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethpseporttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethpseporttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethpseporttable.EntityData.Children = make(map[string]types.YChild)
    pethpseporttable.EntityData.Children["pethPsePortEntry"] = types.YChild{"Pethpseportentry", nil}
    for i := range pethpseporttable.Pethpseportentry {
        pethpseporttable.EntityData.Children[types.GetSegmentPath(&pethpseporttable.Pethpseportentry[i])] = types.YChild{"Pethpseportentry", &pethpseporttable.Pethpseportentry[i]}
    }
    pethpseporttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pethpseporttable.EntityData)
}

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry
// A set of objects that display and control the power
// characteristics of a power Ethernet PSE port.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group
    // containing the port to which a power Ethernet PSE is connected.  Group
    // means box in the stack, module in a rack and the value 1 MUST be used for
    // non-modular devices. Furthermore, the same value MUST be used in this
    // variable, pethMainPseGroupIndex, and pethNotificationControlGroupIndex to
    // refer to a given box in a stack or module in the rack. The type is
    // interface{} with range: 1..2147483647.
    Pethpseportgroupindex interface{}

    // This attribute is a key. This variable uniquely identifies the power
    // Ethernet PSE port within group pethPsePortGroupIndex to which the power
    // Ethernet PSE entry is connected. The type is interface{} with range:
    // 1..2147483647.
    Pethpseportindex interface{}

    // true (1) An interface which can provide the PSE functions. false(2) The
    // interface will act as it would if it had no PSE function. The type is bool.
    Pethpseportadminenable interface{}

    // Describes the capability of controlling the power pairs functionality to
    // switch pins for sourcing power. The value true indicate that the device has
    // the capability to control the power pairs.  When false the PSE Pinout
    // Alternative used cannot be controlled through the PethPsePortAdminEnable
    // attribute. The type is bool.
    Pethpseportpowerpairscontrolability interface{}

    // Describes or controls the pairs in use.  If the value of
    // pethPsePortPowerPairsControl is true, this object is writable. A value of
    // signal(1) means that the signal pairs only are in use. A value of spare(2)
    // means that the spare pairs only are in use. The type is
    // Pethpseportpowerpairs.
    Pethpseportpowerpairs interface{}

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
    // is Pethpseportdetectionstatus.
    Pethpseportdetectionstatus interface{}

    // This object controls the priority of the port from the point of view of a
    // power management algorithm.  The priority that is set by this variable
    // could be used by a control mechanism that prevents over current situations
    // by disconnecting first ports with lower power priority.  Ports that connect
    // devices critical to the operation of the network - like the E911 telephones
    // ports - should be set to higher priority. The type is
    // Pethpseportpowerpriority.
    Pethpseportpowerpriority interface{}

    // This counter is incremented when the PSE state diagram transitions directly
    // from the state POWER_ON to the state IDLE due to tmpdo_timer_done being
    // asserted. The type is interface{} with range: 0..4294967295.
    Pethpseportmpsabsentcounter interface{}

    // A manager will set the value of this variable to indicate the type of
    // powered device that is connected to the port. The default value supplied by
    // the agent if no value has ever been set should be a zero-length octet
    // string. The type is string.
    Pethpseporttype interface{}

    // Classification is a way to tag different terminals on the Power over LAN
    // network according to their power consumption. Devices such as IP
    // telephones, WLAN access points and others, will be classified according to
    // their power requirements.  The meaning of the classification labels is
    // defined in the IEEE specification.  This variable is valid only while a PD
    // is being powered, that is, while the attribute pethPsePortDetectionStatus
    // is reporting the enumeration deliveringPower. The type is
    // Pethpseportpowerclassifications.
    Pethpseportpowerclassifications interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // SIGNATURE_INVALID. The type is interface{} with range: 0..4294967295.
    Pethpseportinvalidsignaturecounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // POWER_DENIED. The type is interface{} with range: 0..4294967295.
    Pethpseportpowerdeniedcounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // ERROR_DELAY_OVER. The type is interface{} with range: 0..4294967295.
    Pethpseportoverloadcounter interface{}

    // This counter is incremented when the PSE state diagram enters the state
    // ERROR_DELAY_SHORT. The type is interface{} with range: 0..4294967295.
    Pethpseportshortcounter interface{}

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
    // 'limit' or 'disable'. The type is Cpeextpseportenable.
    Cpeextpseportenable interface{}

    // This object indicates the discover mode used by the system to discover the
    // PD.  A value of 'unknown' indicates that the discover mode on the interface
    // is unknown.  A value of 'off' indicates that discovery is disabled on the
    // interface.  A value of 'ieee' indicates that the discover mode on the
    // interface is IEEE based.  A value of 'cisco' indicates that the discover
    // mode on the interface is Cisco based.  A value of 'ieeeAndCisco' indicates
    // that the discover mode on the interface is both IEEE and Cisco based. The
    // type is Cpeextpseportdiscovermode.
    Cpeextpseportdiscovermode interface{}

    // This object indicates if a Powered Device (PD) has been detected on the
    // interface.  A value of 'true' indicates that a PD has been detected on the
    // interface.  A value of 'false' indicates that no PD has been detected on
    // the interface. The type is bool.
    Cpeextpseportdevicedetected interface{}

    // This object indicates whether the Powered Device attached to the interface
    // is an IEEE compliant Powered Device or not.  A value of 'true' indicates
    // the attached Powered Device is an IEEE compliant Powered Device.  A value
    // of 'false' indicates the attached Powered Device is not an IEEE compliant
    // Powered Device. This also means that the value of the corresponding object
    // from the pethPsePortTable, pethPsePortPowerClassifications is irrelevant.
    // The type is bool.
    Cpeextpseportieeepd interface{}

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
    Cpeextpseportadditionalstatus interface{}

    // This indicates the maximum amount of power that the PSE will make available
    // to the PD connected to this interface. The type is interface{} with range:
    // 0..4294967295. Units are milliwatts.
    Cpeextpseportpwrmax interface{}

    // This object indicates the amount of power allocated from the PSE for the
    // PD. The type is interface{} with range: 0..4294967295. Units are
    // milliwatts.
    Cpeextpseportpwrallocated interface{}

    // This object indicates the amount of power available for the PD to use. This
    // value may differ from the value cpeExtPsePortPwrAllocated due to the
    // efficiency issues of delivering the power from the PSE to the PD.  When
    // sufficient power is available to power a PD, this value should be equal to
    // the lower of the two objects, cpeExtDefaultAllocation and
    // cpeExtPsePortPwrMax. The type is interface{} with range: 0..4294967295.
    // Units are milliwatts.
    Cpeextpseportpwravailable interface{}

    // This indicates the actual power consumption of the PD connected to this
    // interface. It may not necessarily be equal to the value of
    // cpeExtPsePortPwrAvailable. The type is interface{} with range:
    // 0..4294967295. Units are milliwatts.
    Cpeextpseportpwrconsumption interface{}

    // This indicates the maximum amount of power drawn by the PD connected to
    // this interface, since it was powered on. The type is interface{} with
    // range: 0..4294967295. Units are milliwatts.
    Cpeextpseportmaxpwrdrawn interface{}

    // The entPhysicalIndex value that uniquely identifies the PSE port. If the
    // PSE port does not have a corresponding  physical entry in entPhysicalTable
    // or if the  entPhysicalTable is not supported by the management system, this
    // object has the value of zero. The type is interface{} with range:
    // 0..2147483647.
    Cpeextpseportentphyindex interface{}

    // This object indicates whether the PSE port hardware is capable of policing
    // the port for proper power consumption  based on the allocated value. The
    // type is bool.
    Cpeextpseportpolicingcapable interface{}

    // This object allows the user to turn on or turn off the power policing of
    // the PSE port. If the instance value of  cpeExtPsePortPolicingCapable is
    // 'TRUE', the user is allowed to set this object to 'on' or 'off'. Otherwise,
    // this object is read-only and always has the value of 'off'. The type is
    // Cpeextpseportpolicingenable.
    Cpeextpseportpolicingenable interface{}

    // This object specifies the power policing action that the device will take
    // on this PSE port when the real-time  power consumption exceeds its max
    // power allocation if  the value of cpeExtPsePortPolicingEnable is 'on'.     
    // 'deny'          - the device will deny the power to                        
    // the PSE port       'logOnly'       - the device will not deny the power    
    // to the PSE port. The type is Cpeextpseportpolicingaction.
    Cpeextpseportpolicingaction interface{}

    // This object specifies the manual power allocation that the PSE will
    // allocate to the PD connected to this  interface regardless of the amount
    // requested via CDP  or IEEE.   Setting this object value to zero disables
    // the manual  power allocation.   Warning: Misconfiguring this manual power
    // allocation may  cause damage to the system and void the warranty. Take 
    // precautions not to oversubscribe the power supply. The type is interface{}
    // with range: 0..4294967295. Units are milliwatts.
    Cpeextpseportpwrmanalloc interface{}

    // This object indicates the PSE functionality that this port supports.  If
    // the 'policing' BIT is set, then this PSE port is capable of policing the
    // port for proper power consumption based on the allocated value.  If the
    // 'poePlus' BIT is set, then this PSE port supports PoE Plus functions. The
    // type is map[string]bool.
    Cpeextpseportcapabilities interface{}
}

func (pethpseportentry *POWERETHERNETMIB_Pethpseporttable_Pethpseportentry) GetEntityData() *types.CommonEntityData {
    pethpseportentry.EntityData.YFilter = pethpseportentry.YFilter
    pethpseportentry.EntityData.YangName = "pethPsePortEntry"
    pethpseportentry.EntityData.BundleName = "cisco_ios_xe"
    pethpseportentry.EntityData.ParentYangName = "pethPsePortTable"
    pethpseportentry.EntityData.SegmentPath = "pethPsePortEntry" + "[pethPsePortGroupIndex='" + fmt.Sprintf("%v", pethpseportentry.Pethpseportgroupindex) + "']" + "[pethPsePortIndex='" + fmt.Sprintf("%v", pethpseportentry.Pethpseportindex) + "']"
    pethpseportentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethpseportentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethpseportentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethpseportentry.EntityData.Children = make(map[string]types.YChild)
    pethpseportentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pethpseportentry.EntityData.Leafs["pethPsePortGroupIndex"] = types.YLeaf{"Pethpseportgroupindex", pethpseportentry.Pethpseportgroupindex}
    pethpseportentry.EntityData.Leafs["pethPsePortIndex"] = types.YLeaf{"Pethpseportindex", pethpseportentry.Pethpseportindex}
    pethpseportentry.EntityData.Leafs["pethPsePortAdminEnable"] = types.YLeaf{"Pethpseportadminenable", pethpseportentry.Pethpseportadminenable}
    pethpseportentry.EntityData.Leafs["pethPsePortPowerPairsControlAbility"] = types.YLeaf{"Pethpseportpowerpairscontrolability", pethpseportentry.Pethpseportpowerpairscontrolability}
    pethpseportentry.EntityData.Leafs["pethPsePortPowerPairs"] = types.YLeaf{"Pethpseportpowerpairs", pethpseportentry.Pethpseportpowerpairs}
    pethpseportentry.EntityData.Leafs["pethPsePortDetectionStatus"] = types.YLeaf{"Pethpseportdetectionstatus", pethpseportentry.Pethpseportdetectionstatus}
    pethpseportentry.EntityData.Leafs["pethPsePortPowerPriority"] = types.YLeaf{"Pethpseportpowerpriority", pethpseportentry.Pethpseportpowerpriority}
    pethpseportentry.EntityData.Leafs["pethPsePortMPSAbsentCounter"] = types.YLeaf{"Pethpseportmpsabsentcounter", pethpseportentry.Pethpseportmpsabsentcounter}
    pethpseportentry.EntityData.Leafs["pethPsePortType"] = types.YLeaf{"Pethpseporttype", pethpseportentry.Pethpseporttype}
    pethpseportentry.EntityData.Leafs["pethPsePortPowerClassifications"] = types.YLeaf{"Pethpseportpowerclassifications", pethpseportentry.Pethpseportpowerclassifications}
    pethpseportentry.EntityData.Leafs["pethPsePortInvalidSignatureCounter"] = types.YLeaf{"Pethpseportinvalidsignaturecounter", pethpseportentry.Pethpseportinvalidsignaturecounter}
    pethpseportentry.EntityData.Leafs["pethPsePortPowerDeniedCounter"] = types.YLeaf{"Pethpseportpowerdeniedcounter", pethpseportentry.Pethpseportpowerdeniedcounter}
    pethpseportentry.EntityData.Leafs["pethPsePortOverLoadCounter"] = types.YLeaf{"Pethpseportoverloadcounter", pethpseportentry.Pethpseportoverloadcounter}
    pethpseportentry.EntityData.Leafs["pethPsePortShortCounter"] = types.YLeaf{"Pethpseportshortcounter", pethpseportentry.Pethpseportshortcounter}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortEnable"] = types.YLeaf{"Cpeextpseportenable", pethpseportentry.Cpeextpseportenable}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortDiscoverMode"] = types.YLeaf{"Cpeextpseportdiscovermode", pethpseportentry.Cpeextpseportdiscovermode}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortDeviceDetected"] = types.YLeaf{"Cpeextpseportdevicedetected", pethpseportentry.Cpeextpseportdevicedetected}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortIeeePd"] = types.YLeaf{"Cpeextpseportieeepd", pethpseportentry.Cpeextpseportieeepd}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortAdditionalStatus"] = types.YLeaf{"Cpeextpseportadditionalstatus", pethpseportentry.Cpeextpseportadditionalstatus}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPwrMax"] = types.YLeaf{"Cpeextpseportpwrmax", pethpseportentry.Cpeextpseportpwrmax}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPwrAllocated"] = types.YLeaf{"Cpeextpseportpwrallocated", pethpseportentry.Cpeextpseportpwrallocated}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPwrAvailable"] = types.YLeaf{"Cpeextpseportpwravailable", pethpseportentry.Cpeextpseportpwravailable}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPwrConsumption"] = types.YLeaf{"Cpeextpseportpwrconsumption", pethpseportentry.Cpeextpseportpwrconsumption}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortMaxPwrDrawn"] = types.YLeaf{"Cpeextpseportmaxpwrdrawn", pethpseportentry.Cpeextpseportmaxpwrdrawn}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortEntPhyIndex"] = types.YLeaf{"Cpeextpseportentphyindex", pethpseportentry.Cpeextpseportentphyindex}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPolicingCapable"] = types.YLeaf{"Cpeextpseportpolicingcapable", pethpseportentry.Cpeextpseportpolicingcapable}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPolicingEnable"] = types.YLeaf{"Cpeextpseportpolicingenable", pethpseportentry.Cpeextpseportpolicingenable}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPolicingAction"] = types.YLeaf{"Cpeextpseportpolicingaction", pethpseportentry.Cpeextpseportpolicingaction}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortPwrManAlloc"] = types.YLeaf{"Cpeextpseportpwrmanalloc", pethpseportentry.Cpeextpseportpwrmanalloc}
    pethpseportentry.EntityData.Leafs["cpeExtPsePortCapabilities"] = types.YLeaf{"Cpeextpseportcapabilities", pethpseportentry.Cpeextpseportcapabilities}
    return &(pethpseportentry.EntityData)
}

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode represents the interface is both IEEE and Cisco based.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode_unknown POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode = "unknown"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode_off POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode = "off"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode_ieee POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode = "ieee"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode_cisco POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode = "cisco"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode_ieeeAndCisco POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportdiscovermode = "ieeeAndCisco"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable represents 'limit' or 'disable'.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable_auto POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable = "auto"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable_static POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable = "static"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable_limit POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable = "limit"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable_disable POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportenable = "disable"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction represents                        to the PSE port
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction_deny POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction = "deny"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction_logOnly POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingaction = "logOnly"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable represents object is read-only and always has the value of 'off'.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable_on POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable = "on"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable_off POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Cpeextpseportpolicingenable = "off"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus represents in a state other than those listed above.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_disabled POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "disabled"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_searching POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "searching"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_deliveringPower POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "deliveringPower"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_fault POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "fault"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_test POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "test"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus_otherFault POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportdetectionstatus = "otherFault"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications represents is reporting the enumeration deliveringPower.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications_class0 POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications = "class0"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications_class1 POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications = "class1"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications_class2 POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications = "class2"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications_class3 POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications = "class3"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications_class4 POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerclassifications = "class4"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs represents only are in use.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs_signal POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs = "signal"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs_spare POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpairs = "spare"
)

// POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority represents telephones ports - should be set to higher priority.
type POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority string

const (
    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority_critical POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority = "critical"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority_high POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority = "high"

    POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority_low POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportpowerpriority = "low"
)

// POWERETHERNETMIB_Pethmainpsetable
// A table of objects that display and control attributes
// of the main power source in a PSE  device.  Ethernet
// switches are one example of boxes that would support
// these objects.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_Pethmainpsetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that display and control the Main power of a PSE. . The
    // type is slice of POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry.
    Pethmainpseentry []POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry
}

func (pethmainpsetable *POWERETHERNETMIB_Pethmainpsetable) GetEntityData() *types.CommonEntityData {
    pethmainpsetable.EntityData.YFilter = pethmainpsetable.YFilter
    pethmainpsetable.EntityData.YangName = "pethMainPseTable"
    pethmainpsetable.EntityData.BundleName = "cisco_ios_xe"
    pethmainpsetable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethmainpsetable.EntityData.SegmentPath = "pethMainPseTable"
    pethmainpsetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethmainpsetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethmainpsetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethmainpsetable.EntityData.Children = make(map[string]types.YChild)
    pethmainpsetable.EntityData.Children["pethMainPseEntry"] = types.YChild{"Pethmainpseentry", nil}
    for i := range pethmainpsetable.Pethmainpseentry {
        pethmainpsetable.EntityData.Children[types.GetSegmentPath(&pethmainpsetable.Pethmainpseentry[i])] = types.YChild{"Pethmainpseentry", &pethmainpsetable.Pethmainpseentry[i]}
    }
    pethmainpsetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pethmainpsetable.EntityData)
}

// POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry
// A set of objects that display and control the Main
// power of a PSE. 
type POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group to
    // which power Ethernet PSE is connected.  Group means (box in the stack,
    // module in a rack) and the value 1 MUST be used for non-modular devices. 
    // Furthermore, the same value MUST be used in this variable,
    // pethPsePortGroupIndex, and pethNotificationControlGroupIndex to refer to a
    // given box in a stack or module in a rack. The type is interface{} with
    // range: 1..2147483647.
    Pethmainpsegroupindex interface{}

    // The nominal power of the PSE expressed in Watts. The type is interface{}
    // with range: 1..65535. Units are Watts.
    Pethmainpsepower interface{}

    // The operational status of the main PSE. The type is Pethmainpseoperstatus.
    Pethmainpseoperstatus interface{}

    // Measured usage power expressed in Watts. The type is interface{} with
    // range: 0..4294967295. Units are Watts.
    Pethmainpseconsumptionpower interface{}

    // The usage threshold expressed in percents for comparing the measured power
    // and initiating an alarm if the threshold is exceeded. The type is
    // interface{} with range: 1..99. Units are %.
    Pethmainpseusagethreshold interface{}
}

func (pethmainpseentry *POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry) GetEntityData() *types.CommonEntityData {
    pethmainpseentry.EntityData.YFilter = pethmainpseentry.YFilter
    pethmainpseentry.EntityData.YangName = "pethMainPseEntry"
    pethmainpseentry.EntityData.BundleName = "cisco_ios_xe"
    pethmainpseentry.EntityData.ParentYangName = "pethMainPseTable"
    pethmainpseentry.EntityData.SegmentPath = "pethMainPseEntry" + "[pethMainPseGroupIndex='" + fmt.Sprintf("%v", pethmainpseentry.Pethmainpsegroupindex) + "']"
    pethmainpseentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethmainpseentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethmainpseentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethmainpseentry.EntityData.Children = make(map[string]types.YChild)
    pethmainpseentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pethmainpseentry.EntityData.Leafs["pethMainPseGroupIndex"] = types.YLeaf{"Pethmainpsegroupindex", pethmainpseentry.Pethmainpsegroupindex}
    pethmainpseentry.EntityData.Leafs["pethMainPsePower"] = types.YLeaf{"Pethmainpsepower", pethmainpseentry.Pethmainpsepower}
    pethmainpseentry.EntityData.Leafs["pethMainPseOperStatus"] = types.YLeaf{"Pethmainpseoperstatus", pethmainpseentry.Pethmainpseoperstatus}
    pethmainpseentry.EntityData.Leafs["pethMainPseConsumptionPower"] = types.YLeaf{"Pethmainpseconsumptionpower", pethmainpseentry.Pethmainpseconsumptionpower}
    pethmainpseentry.EntityData.Leafs["pethMainPseUsageThreshold"] = types.YLeaf{"Pethmainpseusagethreshold", pethmainpseentry.Pethmainpseusagethreshold}
    return &(pethmainpseentry.EntityData)
}

// POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus represents The operational status of the main PSE.
type POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus string

const (
    POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus_on POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus = "on"

    POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus_off POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus = "off"

    POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus_faulty POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpseoperstatus = "faulty"
)

// POWERETHERNETMIB_Pethnotificationcontroltable
// A table of objects that display and control the
// Notification on a PSE  device.
// Values of all read-write objects in this table are
// persistent at restart/reboot.
type POWERETHERNETMIB_Pethnotificationcontroltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A set of objects that control the Notification events. The type is slice of
    // POWERETHERNETMIB_Pethnotificationcontroltable_Pethnotificationcontrolentry.
    Pethnotificationcontrolentry []POWERETHERNETMIB_Pethnotificationcontroltable_Pethnotificationcontrolentry
}

func (pethnotificationcontroltable *POWERETHERNETMIB_Pethnotificationcontroltable) GetEntityData() *types.CommonEntityData {
    pethnotificationcontroltable.EntityData.YFilter = pethnotificationcontroltable.YFilter
    pethnotificationcontroltable.EntityData.YangName = "pethNotificationControlTable"
    pethnotificationcontroltable.EntityData.BundleName = "cisco_ios_xe"
    pethnotificationcontroltable.EntityData.ParentYangName = "POWER-ETHERNET-MIB"
    pethnotificationcontroltable.EntityData.SegmentPath = "pethNotificationControlTable"
    pethnotificationcontroltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethnotificationcontroltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethnotificationcontroltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethnotificationcontroltable.EntityData.Children = make(map[string]types.YChild)
    pethnotificationcontroltable.EntityData.Children["pethNotificationControlEntry"] = types.YChild{"Pethnotificationcontrolentry", nil}
    for i := range pethnotificationcontroltable.Pethnotificationcontrolentry {
        pethnotificationcontroltable.EntityData.Children[types.GetSegmentPath(&pethnotificationcontroltable.Pethnotificationcontrolentry[i])] = types.YChild{"Pethnotificationcontrolentry", &pethnotificationcontroltable.Pethnotificationcontrolentry[i]}
    }
    pethnotificationcontroltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pethnotificationcontroltable.EntityData)
}

// POWERETHERNETMIB_Pethnotificationcontroltable_Pethnotificationcontrolentry
// A set of objects that control the Notification events.
type POWERETHERNETMIB_Pethnotificationcontroltable_Pethnotificationcontrolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This variable uniquely identifies the group. 
    // Group means box in the stack, module in a rack and the value 1 MUST be used
    // for non-modular devices.  Furthermore, the same value MUST be used in this
    // variable, pethPsePortGroupIndex, and pethMainPseGroupIndex to refer to a
    // given box in a stack or module in a rack. . The type is interface{} with
    // range: 1..2147483647.
    Pethnotificationcontrolgroupindex interface{}

    // This object controls, on a per-group basis, whether or not notifications
    // from the agent are enabled.  The value true(1) means that notifications are
    // enabled; the value false(2) means that they are not. The type is bool.
    Pethnotificationcontrolenable interface{}
}

func (pethnotificationcontrolentry *POWERETHERNETMIB_Pethnotificationcontroltable_Pethnotificationcontrolentry) GetEntityData() *types.CommonEntityData {
    pethnotificationcontrolentry.EntityData.YFilter = pethnotificationcontrolentry.YFilter
    pethnotificationcontrolentry.EntityData.YangName = "pethNotificationControlEntry"
    pethnotificationcontrolentry.EntityData.BundleName = "cisco_ios_xe"
    pethnotificationcontrolentry.EntityData.ParentYangName = "pethNotificationControlTable"
    pethnotificationcontrolentry.EntityData.SegmentPath = "pethNotificationControlEntry" + "[pethNotificationControlGroupIndex='" + fmt.Sprintf("%v", pethnotificationcontrolentry.Pethnotificationcontrolgroupindex) + "']"
    pethnotificationcontrolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pethnotificationcontrolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pethnotificationcontrolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pethnotificationcontrolentry.EntityData.Children = make(map[string]types.YChild)
    pethnotificationcontrolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    pethnotificationcontrolentry.EntityData.Leafs["pethNotificationControlGroupIndex"] = types.YLeaf{"Pethnotificationcontrolgroupindex", pethnotificationcontrolentry.Pethnotificationcontrolgroupindex}
    pethnotificationcontrolentry.EntityData.Leafs["pethNotificationControlEnable"] = types.YLeaf{"Pethnotificationcontrolenable", pethnotificationcontrolentry.Pethnotificationcontrolenable}
    return &(pethnotificationcontrolentry.EntityData)
}

