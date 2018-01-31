// Textual conventions used throughout the various IEEE 802.1 MIB
// modules.
// 
// Unless otherwise indicated, the references in this MIB
// module are to IEEE 802.1Q-2005 as amended by IEEE 802.1ad,
// IEEE 802.1ak, IEEE 802.1ag and IEEE 802.1ah.
// 
// Copyright (C) IEEE.
// This version of this MIB module is part of IEEE802.1Q;
// see the draft itself for full legal notices.
package ieee8021_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ieee8021_tc_mib"))
}

// IEEE8021PriorityCodePoint represents the PCP using the Priority Code Point Decoding Table.
type IEEE8021PriorityCodePoint string

const (
    IEEE8021PriorityCodePoint_codePoint8p0d IEEE8021PriorityCodePoint = "codePoint8p0d"

    IEEE8021PriorityCodePoint_codePoint7p1d IEEE8021PriorityCodePoint = "codePoint7p1d"

    IEEE8021PriorityCodePoint_codePoint6p2d IEEE8021PriorityCodePoint = "codePoint6p2d"

    IEEE8021PriorityCodePoint_codePoint5p3d IEEE8021PriorityCodePoint = "codePoint5p3d"
)

// IEEE8021ServiceSelectorType represents example, lead to an undefined IEEE8021ServiceSelectorValue value.
type IEEE8021ServiceSelectorType string

const (
    IEEE8021ServiceSelectorType_vlanId IEEE8021ServiceSelectorType = "vlanId"

    IEEE8021ServiceSelectorType_isid IEEE8021ServiceSelectorType = "isid"
)

// IEEE8021PortAcceptableFrameTypes represents Acceptable frame types on a port.
type IEEE8021PortAcceptableFrameTypes string

const (
    IEEE8021PortAcceptableFrameTypes_admitAll IEEE8021PortAcceptableFrameTypes = "admitAll"

    IEEE8021PortAcceptableFrameTypes_admitUntaggedAndPriority IEEE8021PortAcceptableFrameTypes = "admitUntaggedAndPriority"

    IEEE8021PortAcceptableFrameTypes_admitTagged IEEE8021PortAcceptableFrameTypes = "admitTagged"
)

// IEEE8021BridgePortType represents     member of an 802.1D bridge.
type IEEE8021BridgePortType string

const (
    IEEE8021BridgePortType_none IEEE8021BridgePortType = "none"

    IEEE8021BridgePortType_customerVlanPort IEEE8021BridgePortType = "customerVlanPort"

    IEEE8021BridgePortType_providerNetworkPort IEEE8021BridgePortType = "providerNetworkPort"

    IEEE8021BridgePortType_customerNetworkPort IEEE8021BridgePortType = "customerNetworkPort"

    IEEE8021BridgePortType_customerEdgePort IEEE8021BridgePortType = "customerEdgePort"

    IEEE8021BridgePortType_customerBackbonePort IEEE8021BridgePortType = "customerBackbonePort"

    IEEE8021BridgePortType_virtualInstancePort IEEE8021BridgePortType = "virtualInstancePort"

    IEEE8021BridgePortType_dBridgePort IEEE8021BridgePortType = "dBridgePort"
)

