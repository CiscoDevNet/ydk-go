// This MIB module defines Textual Conventions for
// representing targets which have class based policy 
// mappings. A target can be any logical interface 
// or entity to which a class based policy is able to be 
// associated.
package cisco_cbp_target_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cbp_target_tc_mib"))
}

// CcbptTargetType represents               CcbptTargetIdAaaSession Textual Convention.
type CcbptTargetType string

const (
    CcbptTargetType_genIf CcbptTargetType = "genIf"

    CcbptTargetType_atmPvc CcbptTargetType = "atmPvc"

    CcbptTargetType_frDlci CcbptTargetType = "frDlci"

    CcbptTargetType_entity CcbptTargetType = "entity"

    CcbptTargetType_fwZone CcbptTargetType = "fwZone"

    CcbptTargetType_fwZonePair CcbptTargetType = "fwZonePair"

    CcbptTargetType_aaaSession CcbptTargetType = "aaaSession"
)

// CcbptTargetDirection represents             relative to the target.
type CcbptTargetDirection string

const (
    CcbptTargetDirection_undirected CcbptTargetDirection = "undirected"

    CcbptTargetDirection_input CcbptTargetDirection = "input"

    CcbptTargetDirection_output CcbptTargetDirection = "output"

    CcbptTargetDirection_inOut CcbptTargetDirection = "inOut"
)

// CcbptPolicySourceType represents                    Class Based.
type CcbptPolicySourceType string

const (
    CcbptPolicySourceType_ciscoCbQos CcbptPolicySourceType = "ciscoCbQos"

    CcbptPolicySourceType_ciscoCbpBase CcbptPolicySourceType = "ciscoCbpBase"
)

