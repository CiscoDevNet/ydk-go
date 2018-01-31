// This MIB module defines textual conventions used by the
// CISCO-DYNAMIC-TEMPLATE-MIB and MIB modules that use and expand
// on dynamic templates.
package cisco_dynamic_template_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_dynamic_template_tc_mib"))
}

// DynamicTemplateTargetType represents error-status of 'inconsistentValue'.
type DynamicTemplateTargetType string

const (
    DynamicTemplateTargetType_other DynamicTemplateTargetType = "other"

    DynamicTemplateTargetType_interface_ DynamicTemplateTargetType = "interface"
)

// DynamicTemplateType represents     life-cycle events.
type DynamicTemplateType string

const (
    DynamicTemplateType_other DynamicTemplateType = "other"

    DynamicTemplateType_derived DynamicTemplateType = "derived"

    DynamicTemplateType_ppp DynamicTemplateType = "ppp"

    DynamicTemplateType_ethernet DynamicTemplateType = "ethernet"

    DynamicTemplateType_ipSubscriber DynamicTemplateType = "ipSubscriber"

    DynamicTemplateType_service DynamicTemplateType = "service"
)

