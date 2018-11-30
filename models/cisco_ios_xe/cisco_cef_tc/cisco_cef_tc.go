// This MIB module defines Textual Conventions and
// OBJECT-IDENTITIES for use in documents defining
// management information base (MIBs) modules for 
// managing Cisco Express Forwarding (CEF).
package cisco_cef_tc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cef_tc"))
}

// CefFailureReason represents                          detected for CEF
type CefFailureReason string

const (
    CefFailureReason_none CefFailureReason = "none"

    CefFailureReason_mallocFailure CefFailureReason = "mallocFailure"

    CefFailureReason_hwFailure CefFailureReason = "hwFailure"

    CefFailureReason_keepaliveFailure CefFailureReason = "keepaliveFailure"

    CefFailureReason_noMsgBuffer CefFailureReason = "noMsgBuffer"

    CefFailureReason_invalidMsgSize CefFailureReason = "invalidMsgSize"

    CefFailureReason_internalError CefFailureReason = "internalError"
)

// CefCCStatus represents                       
type CefCCStatus string

const (
    CefCCStatus_ccStatusIdle CefCCStatus = "ccStatusIdle"

    CefCCStatus_ccStatusRunning CefCCStatus = "ccStatusRunning"

    CefCCStatus_ccStatusDone CefCCStatus = "ccStatusDone"
)

// CefForwardingElementSpecialType represents                ignored 
type CefForwardingElementSpecialType string

const (
    CefForwardingElementSpecialType_illegal CefForwardingElementSpecialType = "illegal"

    CefForwardingElementSpecialType_punt CefForwardingElementSpecialType = "punt"

    CefForwardingElementSpecialType_drop CefForwardingElementSpecialType = "drop"

    CefForwardingElementSpecialType_discard CefForwardingElementSpecialType = "discard"

    CefForwardingElementSpecialType_null CefForwardingElementSpecialType = "null"

    CefForwardingElementSpecialType_glean CefForwardingElementSpecialType = "glean"

    CefForwardingElementSpecialType_unresolved CefForwardingElementSpecialType = "unresolved"

    CefForwardingElementSpecialType_noRoute CefForwardingElementSpecialType = "noRoute"

    CefForwardingElementSpecialType_none CefForwardingElementSpecialType = "none"
)

// CefPrefixSearchState represents                     match has not been found.
type CefPrefixSearchState string

const (
    CefPrefixSearchState_running CefPrefixSearchState = "running"

    CefPrefixSearchState_matchFound CefPrefixSearchState = "matchFound"

    CefPrefixSearchState_noMatchFound CefPrefixSearchState = "noMatchFound"
)

// CefPathType represents .
type CefPathType string

const (
    CefPathType_receive CefPathType = "receive"

    CefPathType_connectedPrefix CefPathType = "connectedPrefix"

    CefPathType_attachedPrefix CefPathType = "attachedPrefix"

    CefPathType_attachedHost CefPathType = "attachedHost"

    CefPathType_attachedNexthop CefPathType = "attachedNexthop"

    CefPathType_recursiveNexthop CefPathType = "recursiveNexthop"

    CefPathType_adjacencyPrefix CefPathType = "adjacencyPrefix"

    CefPathType_specialPrefix CefPathType = "specialPrefix"

    CefPathType_unknown CefPathType = "unknown"
)

// CefCCType represents                    inconsistencies.
type CefCCType string

const (
    CefCCType_lcDetect CefCCType = "lcDetect"

    CefCCType_scanFibLcRp CefCCType = "scanFibLcRp"

    CefCCType_scanFibRpLc CefCCType = "scanFibRpLc"

    CefCCType_scanRibFib CefCCType = "scanRibFib"

    CefCCType_scanFibRib CefCCType = "scanFibRib"

    CefCCType_scanFibHwSw CefCCType = "scanFibHwSw"

    CefCCType_scanFibSwHw CefCCType = "scanFibSwHw"

    CefCCType_fullScanRibFib CefCCType = "fullScanRibFib"

    CefCCType_fullScanFibRib CefCCType = "fullScanFibRib"

    CefCCType_fullScanFibRpLc CefCCType = "fullScanFibRpLc"

    CefCCType_fullScanFibLcRp CefCCType = "fullScanFibLcRp"

    CefCCType_fullScanFibHwSw CefCCType = "fullScanFibHwSw"

    CefCCType_fullScanFibSwHw CefCCType = "fullScanFibSwHw"
)

// CefOperStatus represents Operational status of CEF.
type CefOperStatus string

const (
    CefOperStatus_up CefOperStatus = "up"

    CefOperStatus_down CefOperStatus = "down"
)

// CefAdjLinkType represents process packets fed through adjacency.
type CefAdjLinkType string

const (
    CefAdjLinkType_ipv4 CefAdjLinkType = "ipv4"

    CefAdjLinkType_ipv6 CefAdjLinkType = "ipv6"

    CefAdjLinkType_mpls CefAdjLinkType = "mpls"

    CefAdjLinkType_raw CefAdjLinkType = "raw"

    CefAdjLinkType_unknown CefAdjLinkType = "unknown"
)

// CefCCAction represents                         on consistency checkers.
type CefCCAction string

const (
    CefCCAction_ccActionStart CefCCAction = "ccActionStart"

    CefCCAction_ccActionAbort CefCCAction = "ccActionAbort"

    CefCCAction_ccActionNone CefCCAction = "ccActionNone"
)

// CefAdminStatus represents upon the success of the admin operation.
type CefAdminStatus string

const (
    CefAdminStatus_enabled CefAdminStatus = "enabled"

    CefAdminStatus_disabled CefAdminStatus = "disabled"
)

// CefIpVersion represents The version of CEF IP forwarding.
type CefIpVersion string

const (
    CefIpVersion_ipv4 CefIpVersion = "ipv4"

    CefIpVersion_ipv6 CefIpVersion = "ipv6"
)

