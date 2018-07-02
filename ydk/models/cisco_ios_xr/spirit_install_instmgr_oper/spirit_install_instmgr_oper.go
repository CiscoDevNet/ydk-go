// This module contains a collection of YANG definitions
// for Cisco IOS-XR spirit-install-instmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   software-install: Install operations info
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package spirit_install_instmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spirit_install_instmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-instmgr-oper software-install}", reflect.TypeOf(SoftwareInstall{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-instmgr-oper:software-install", reflect.TypeOf(SoftwareInstall{}))
}

// IsdErrorEt represents isd error
type IsdErrorEt string

const (
    // ISD ERROR NONE
    IsdErrorEt_none IsdErrorEt = "none"

    // ISD ERROR NOT COMPATIBLE
    IsdErrorEt_not_compatible IsdErrorEt = "not-compatible"

    // ISD ERROR NOT ENOUGH RESOURCE
    IsdErrorEt_not_enough_resource IsdErrorEt = "not-enough-resource"

    // ISD ERROR NOT NSR READY
    IsdErrorEt_not_nsr_ready IsdErrorEt = "not-nsr-ready"

    // ISD ERROR NOT CONNECTED SDR SM
    IsdErrorEt_not_conn_sdrsm IsdErrorEt = "not-conn-sdrsm"

    // ISD ERROR INST CMD INVALID
    IsdErrorEt_cmd_invalid IsdErrorEt = "cmd-invalid"

    // ISD ERROR INST LOAD PREP FAILURE
    IsdErrorEt_load_prep_fail IsdErrorEt = "load-prep-fail"

    // ISD ERROR TIMEOUT
    IsdErrorEt_error_timeout IsdErrorEt = "error-timeout"

    // ISD ERROR NODE DOWN
    IsdErrorEt_err_node_down IsdErrorEt = "err-node-down"

    // ISD ERROR NODE NOT READY
    IsdErrorEt_node_not_ready IsdErrorEt = "node-not-ready"

    // ISD ERROR NODE NEW
    IsdErrorEt_err_node_new IsdErrorEt = "err-node-new"

    // ISD ERROR CARD OIR
    IsdErrorEt_err_card_oir IsdErrorEt = "err-card-oir"

    // ISD ERROR INVALID EVT
    IsdErrorEt_invalid_evt IsdErrorEt = "invalid-evt"

    // ISD ERROR DISCONN FROM CALVADOS
    IsdErrorEt_disconn_from_calv IsdErrorEt = "disconn-from-calv"

    // ISD ERROR GSP DOWN
    IsdErrorEt_gsp_down IsdErrorEt = "gsp-down"

    // ISD ERROR ABORT BY ISM
    IsdErrorEt_abort_by_ism IsdErrorEt = "abort-by-ism"

    // ISD ERROR RPFO
    IsdErrorEt_rpfo IsdErrorEt = "rpfo"

    // ISD ERROR PKG NULL
    IsdErrorEt_pkg_null IsdErrorEt = "pkg-null"

    // ISD ERROR GENERAL
    IsdErrorEt_error_general IsdErrorEt = "error-general"

    // ISD ERROR FSA ERROR
    IsdErrorEt_fsa_error IsdErrorEt = "fsa-error"

    // ISD ERROR POST ISSU
    IsdErrorEt_err_post_issu IsdErrorEt = "err-post-issu"

    // ISD ERROR ISSUDIR RESTART
    IsdErrorEt_err_issu_dir_restart IsdErrorEt = "err-issu-dir-restart"
)

// NodeRoleEt represents node role
type NodeRoleEt string

const (
    // Unknown
    NodeRoleEt_node_unknown NodeRoleEt = "node-unknown"

    // Active
    NodeRoleEt_node_active NodeRoleEt = "node-active"

    // Standby
    NodeRoleEt_node_standby NodeRoleEt = "node-standby"

    // Unusable
    NodeRoleEt_node_unusable NodeRoleEt = "node-unusable"
)

// IsdStateEt represents isd state
type IsdStateEt string

const (
    // ISSU ST NONE
    IsdStateEt_none IsdStateEt = "none"

    // ISSU ST IDLE
    IsdStateEt_idle IsdStateEt = "idle"

    // ISSU ST INIT
    IsdStateEt_init IsdStateEt = "init"

    // ISSU ST INIT DONE
    IsdStateEt_init_done IsdStateEt = "init-done"

    // ISSU ST LOAD PREP
    IsdStateEt_load_prep IsdStateEt = "load-prep"

    // ISSU ST LOAD EXEC
    IsdStateEt_load_exec IsdStateEt = "load-exec"

    // ISSU ST LOAD ISSU GO
    IsdStateEt_load_issu_go IsdStateEt = "load-issu-go"

    // ISSU ST LOAD DONE
    IsdStateEt_load_done IsdStateEt = "load-done"

    // ISSU ST RUN PREP
    IsdStateEt_run_prep IsdStateEt = "run-prep"

    // ISSU ST RUN BIG BANG
    IsdStateEt_big_bang IsdStateEt = "big-bang"

    // ISSU ST RUN DONE
    IsdStateEt_run_done IsdStateEt = "run-done"

    // ISSU ST CLEANUP
    IsdStateEt_cleanup IsdStateEt = "cleanup"

    // ISSU ST CLEANUP DONE
    IsdStateEt_cleanup_done IsdStateEt = "cleanup-done"

    // ISSU ST ABORT
    IsdStateEt_abort IsdStateEt = "abort"

    // ISSU ST ABORT DONE
    IsdStateEt_abort_done IsdStateEt = "abort-done"

    // ISSU ST ABORT CLEANUP
    IsdStateEt_abort_cleanup IsdStateEt = "abort-cleanup"

    // ISSU UNKNOWN STATE
    IsdStateEt_unknown_state IsdStateEt = "unknown-state"
)

// IsdIssuStatusEt represents isd status
type IsdIssuStatusEt string

const (
    // ISSU STATUS OK
    IsdIssuStatusEt_ok IsdIssuStatusEt = "ok"

    // ISSU STATUS PREP DONE
    IsdIssuStatusEt_prep_done IsdIssuStatusEt = "prep-done"

    // ISSU STATUS BIG BANG
    IsdIssuStatusEt_big_bang IsdIssuStatusEt = "big-bang"

    // ISSU STATUS DONE
    IsdIssuStatusEt_done IsdIssuStatusEt = "done"

    // ISSU STATUS ABORT
    IsdIssuStatusEt_abort IsdIssuStatusEt = "abort"

    // ISSU STATUS CMD REJECT
    IsdIssuStatusEt_cmd_reject IsdIssuStatusEt = "cmd-reject"

    // ISSU STATUS UNKNOWN
    IsdIssuStatusEt_unknown IsdIssuStatusEt = "unknown"

    // ISSU STATUS ABORT CLEANUP
    IsdIssuStatusEt_abort_cleanup IsdIssuStatusEt = "abort-cleanup"

    // ISSU STATUS CMD ABORT REJECT
    IsdIssuStatusEt_abort_cmd_reject IsdIssuStatusEt = "abort-cmd-reject"
)

// IssudirNodeStatusEt represents ISSU node status
type IssudirNodeStatusEt string

const (
    // Not ISSU Ready
    IssudirNodeStatusEt_not_issu_ready IssudirNodeStatusEt = "not-issu-ready"

    // ISSU Ready
    IssudirNodeStatusEt_issu_ready IssudirNodeStatusEt = "issu-ready"

    // ISSU Go
    IssudirNodeStatusEt_isus_go IssudirNodeStatusEt = "isus-go"

    // Node Fail
    IssudirNodeStatusEt_node_fail IssudirNodeStatusEt = "node-fail"
)

// IssuNodeRoleEt represents ISSU role
type IssuNodeRoleEt string

const (
    // Unknown
    IssuNodeRoleEt_unknown_role IssuNodeRoleEt = "unknown-role"

    // Primary
    IssuNodeRoleEt_primary_role IssuNodeRoleEt = "primary-role"

    // Secondary
    IssuNodeRoleEt_secondary_role IssuNodeRoleEt = "secondary-role"

    // Tertiary
    IssuNodeRoleEt_tertiary_role IssuNodeRoleEt = "tertiary-role"
)

// CardTypeEt represents card type
type CardTypeEt string

const (
    // Card RP
    CardTypeEt_card_rp CardTypeEt = "card-rp"

    // Card DRP
    CardTypeEt_card_drp CardTypeEt = "card-drp"

    // Card LC
    CardTypeEt_card_lc CardTypeEt = "card-lc"

    // Card SC
    CardTypeEt_card_sc CardTypeEt = "card-sc"

    // Card SP
    CardTypeEt_card_sp CardTypeEt = "card-sp"

    // Card Other
    CardTypeEt_card_other CardTypeEt = "card-other"
)

// SoftwareInstall
// Install operations info
type SoftwareInstall struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show superseded packages.
    Superseded SoftwareInstall_Superseded

    // Show prepared packages ready for activation.
    Prepare SoftwareInstall_Prepare

    // Show active packages installed.
    Active SoftwareInstall_Active

    // Show install version.
    Version SoftwareInstall_Version

    // Show XR inactive packages.
    Inactive SoftwareInstall_Inactive

    // Show current request.
    Request SoftwareInstall_Request

    // ISSU operation.
    Issu SoftwareInstall_Issu

    // Show Committed packages installed.
    Committed SoftwareInstall_Committed

    // Show log file for all operations.
    AllOperationsLog SoftwareInstall_AllOperationsLog

    // Show the list of installed packages.
    Packages SoftwareInstall_Packages

    // Show log file.
    OperationLogs SoftwareInstall_OperationLogs

    // Show packages stored in install software repositories.
    Repository SoftwareInstall_Repository
}

func (softwareInstall *SoftwareInstall) GetEntityData() *types.CommonEntityData {
    softwareInstall.EntityData.YFilter = softwareInstall.YFilter
    softwareInstall.EntityData.YangName = "software-install"
    softwareInstall.EntityData.BundleName = "cisco_ios_xr"
    softwareInstall.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-instmgr-oper"
    softwareInstall.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-instmgr-oper:software-install"
    softwareInstall.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    softwareInstall.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    softwareInstall.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    softwareInstall.EntityData.Children = types.NewOrderedMap()
    softwareInstall.EntityData.Children.Append("superseded", types.YChild{"Superseded", &softwareInstall.Superseded})
    softwareInstall.EntityData.Children.Append("prepare", types.YChild{"Prepare", &softwareInstall.Prepare})
    softwareInstall.EntityData.Children.Append("active", types.YChild{"Active", &softwareInstall.Active})
    softwareInstall.EntityData.Children.Append("version", types.YChild{"Version", &softwareInstall.Version})
    softwareInstall.EntityData.Children.Append("inactive", types.YChild{"Inactive", &softwareInstall.Inactive})
    softwareInstall.EntityData.Children.Append("request", types.YChild{"Request", &softwareInstall.Request})
    softwareInstall.EntityData.Children.Append("issu", types.YChild{"Issu", &softwareInstall.Issu})
    softwareInstall.EntityData.Children.Append("committed", types.YChild{"Committed", &softwareInstall.Committed})
    softwareInstall.EntityData.Children.Append("all-operations-log", types.YChild{"AllOperationsLog", &softwareInstall.AllOperationsLog})
    softwareInstall.EntityData.Children.Append("packages", types.YChild{"Packages", &softwareInstall.Packages})
    softwareInstall.EntityData.Children.Append("operation-logs", types.YChild{"OperationLogs", &softwareInstall.OperationLogs})
    softwareInstall.EntityData.Children.Append("repository", types.YChild{"Repository", &softwareInstall.Repository})
    softwareInstall.EntityData.Leafs = types.NewOrderedMap()

    softwareInstall.EntityData.YListKeys = []string {}

    return &(softwareInstall.EntityData)
}

// SoftwareInstall_Superseded
// Show superseded packages
type SoftwareInstall_Superseded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // superseded package info. The type is slice of
    // SoftwareInstall_Superseded_SupersededPackageInfo.
    SupersededPackageInfo []*SoftwareInstall_Superseded_SupersededPackageInfo
}

func (superseded *SoftwareInstall_Superseded) GetEntityData() *types.CommonEntityData {
    superseded.EntityData.YFilter = superseded.YFilter
    superseded.EntityData.YangName = "superseded"
    superseded.EntityData.BundleName = "cisco_ios_xr"
    superseded.EntityData.ParentYangName = "software-install"
    superseded.EntityData.SegmentPath = "superseded"
    superseded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    superseded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    superseded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    superseded.EntityData.Children = types.NewOrderedMap()
    superseded.EntityData.Children.Append("superseded-package-info", types.YChild{"SupersededPackageInfo", nil})
    for i := range superseded.SupersededPackageInfo {
        superseded.EntityData.Children.Append(types.GetSegmentPath(superseded.SupersededPackageInfo[i]), types.YChild{"SupersededPackageInfo", superseded.SupersededPackageInfo[i]})
    }
    superseded.EntityData.Leafs = types.NewOrderedMap()

    superseded.EntityData.YListKeys = []string {}

    return &(superseded.EntityData)
}

// SoftwareInstall_Superseded_SupersededPackageInfo
// superseded package info
type SoftwareInstall_Superseded_SupersededPackageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ErrorMessage. The type is string.
    ErrorMessage interface{}

    // Location. The type is string.
    Location interface{}

    // NodeType. The type is string.
    NodeType interface{}

    // BootPartitionName. The type is string.
    BootPartitionName interface{}

    // SupersededPackages. The type is string.
    SupersededPackages interface{}
}

func (supersededPackageInfo *SoftwareInstall_Superseded_SupersededPackageInfo) GetEntityData() *types.CommonEntityData {
    supersededPackageInfo.EntityData.YFilter = supersededPackageInfo.YFilter
    supersededPackageInfo.EntityData.YangName = "superseded-package-info"
    supersededPackageInfo.EntityData.BundleName = "cisco_ios_xr"
    supersededPackageInfo.EntityData.ParentYangName = "superseded"
    supersededPackageInfo.EntityData.SegmentPath = "superseded-package-info"
    supersededPackageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supersededPackageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supersededPackageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supersededPackageInfo.EntityData.Children = types.NewOrderedMap()
    supersededPackageInfo.EntityData.Leafs = types.NewOrderedMap()
    supersededPackageInfo.EntityData.Leafs.Append("error-message", types.YLeaf{"ErrorMessage", supersededPackageInfo.ErrorMessage})
    supersededPackageInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", supersededPackageInfo.Location})
    supersededPackageInfo.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", supersededPackageInfo.NodeType})
    supersededPackageInfo.EntityData.Leafs.Append("boot-partition-name", types.YLeaf{"BootPartitionName", supersededPackageInfo.BootPartitionName})
    supersededPackageInfo.EntityData.Leafs.Append("superseded-packages", types.YLeaf{"SupersededPackages", supersededPackageInfo.SupersededPackages})

    supersededPackageInfo.EntityData.YListKeys = []string {}

    return &(supersededPackageInfo.EntityData)
}

// SoftwareInstall_Prepare
// Show prepared packages ready for activation
type SoftwareInstall_Prepare struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NoPrepareDone. The type is string.
    NoPrepareDone interface{}

    // PreparedBootImage. The type is string.
    PreparedBootImage interface{}

    // PreparedBootPartition. The type is string.
    PreparedBootPartition interface{}

    // RestartType. The type is string.
    RestartType interface{}

    // ActivateMessage. The type is string.
    ActivateMessage interface{}

    // PrepareCleanMessage. The type is string.
    PrepareCleanMessage interface{}

    // rpm. The type is slice of SoftwareInstall_Prepare_Rpm.
    Rpm []*SoftwareInstall_Prepare_Rpm

    // package. The type is slice of SoftwareInstall_Prepare_Package.
    Package []*SoftwareInstall_Prepare_Package
}

func (prepare *SoftwareInstall_Prepare) GetEntityData() *types.CommonEntityData {
    prepare.EntityData.YFilter = prepare.YFilter
    prepare.EntityData.YangName = "prepare"
    prepare.EntityData.BundleName = "cisco_ios_xr"
    prepare.EntityData.ParentYangName = "software-install"
    prepare.EntityData.SegmentPath = "prepare"
    prepare.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepare.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepare.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepare.EntityData.Children = types.NewOrderedMap()
    prepare.EntityData.Children.Append("rpm", types.YChild{"Rpm", nil})
    for i := range prepare.Rpm {
        prepare.EntityData.Children.Append(types.GetSegmentPath(prepare.Rpm[i]), types.YChild{"Rpm", prepare.Rpm[i]})
    }
    prepare.EntityData.Children.Append("package", types.YChild{"Package", nil})
    for i := range prepare.Package {
        prepare.EntityData.Children.Append(types.GetSegmentPath(prepare.Package[i]), types.YChild{"Package", prepare.Package[i]})
    }
    prepare.EntityData.Leafs = types.NewOrderedMap()
    prepare.EntityData.Leafs.Append("no-prepare-done", types.YLeaf{"NoPrepareDone", prepare.NoPrepareDone})
    prepare.EntityData.Leafs.Append("prepared-boot-image", types.YLeaf{"PreparedBootImage", prepare.PreparedBootImage})
    prepare.EntityData.Leafs.Append("prepared-boot-partition", types.YLeaf{"PreparedBootPartition", prepare.PreparedBootPartition})
    prepare.EntityData.Leafs.Append("restart-type", types.YLeaf{"RestartType", prepare.RestartType})
    prepare.EntityData.Leafs.Append("activate-message", types.YLeaf{"ActivateMessage", prepare.ActivateMessage})
    prepare.EntityData.Leafs.Append("prepare-clean-message", types.YLeaf{"PrepareCleanMessage", prepare.PrepareCleanMessage})

    prepare.EntityData.YListKeys = []string {}

    return &(prepare.EntityData)
}

// SoftwareInstall_Prepare_Rpm
// rpm
type SoftwareInstall_Prepare_Rpm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // package. The type is string.
    Package interface{}
}

func (rpm *SoftwareInstall_Prepare_Rpm) GetEntityData() *types.CommonEntityData {
    rpm.EntityData.YFilter = rpm.YFilter
    rpm.EntityData.YangName = "rpm"
    rpm.EntityData.BundleName = "cisco_ios_xr"
    rpm.EntityData.ParentYangName = "prepare"
    rpm.EntityData.SegmentPath = "rpm"
    rpm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpm.EntityData.Children = types.NewOrderedMap()
    rpm.EntityData.Leafs = types.NewOrderedMap()
    rpm.EntityData.Leafs.Append("package", types.YLeaf{"Package", rpm.Package})

    rpm.EntityData.YListKeys = []string {}

    return &(rpm.EntityData)
}

// SoftwareInstall_Prepare_Package
// package
type SoftwareInstall_Prepare_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // package. The type is string.
    Package interface{}
}

func (self *SoftwareInstall_Prepare_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "prepare"
    self.EntityData.SegmentPath = "package"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("package", types.YLeaf{"Package", self.Package})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// SoftwareInstall_Active
// Show active packages installed
type SoftwareInstall_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // active package info. The type is slice of
    // SoftwareInstall_Active_ActivePackageInfo.
    ActivePackageInfo []*SoftwareInstall_Active_ActivePackageInfo
}

func (active *SoftwareInstall_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "software-install"
    active.EntityData.SegmentPath = "active"
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("active-package-info", types.YChild{"ActivePackageInfo", nil})
    for i := range active.ActivePackageInfo {
        active.EntityData.Children.Append(types.GetSegmentPath(active.ActivePackageInfo[i]), types.YChild{"ActivePackageInfo", active.ActivePackageInfo[i]})
    }
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// SoftwareInstall_Active_ActivePackageInfo
// active package info
type SoftwareInstall_Active_ActivePackageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ErrorMessage. The type is string.
    ErrorMessage interface{}

    // Location. The type is string.
    Location interface{}

    // NodeType. The type is string.
    NodeType interface{}

    // BootPartitionName. The type is string.
    BootPartitionName interface{}

    // NumberOfActivePackages. The type is interface{} with range: 0..4294967295.
    NumberOfActivePackages interface{}

    // ActivePackages. The type is string.
    ActivePackages interface{}
}

func (activePackageInfo *SoftwareInstall_Active_ActivePackageInfo) GetEntityData() *types.CommonEntityData {
    activePackageInfo.EntityData.YFilter = activePackageInfo.YFilter
    activePackageInfo.EntityData.YangName = "active-package-info"
    activePackageInfo.EntityData.BundleName = "cisco_ios_xr"
    activePackageInfo.EntityData.ParentYangName = "active"
    activePackageInfo.EntityData.SegmentPath = "active-package-info"
    activePackageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activePackageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activePackageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activePackageInfo.EntityData.Children = types.NewOrderedMap()
    activePackageInfo.EntityData.Leafs = types.NewOrderedMap()
    activePackageInfo.EntityData.Leafs.Append("error-message", types.YLeaf{"ErrorMessage", activePackageInfo.ErrorMessage})
    activePackageInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", activePackageInfo.Location})
    activePackageInfo.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", activePackageInfo.NodeType})
    activePackageInfo.EntityData.Leafs.Append("boot-partition-name", types.YLeaf{"BootPartitionName", activePackageInfo.BootPartitionName})
    activePackageInfo.EntityData.Leafs.Append("number-of-active-packages", types.YLeaf{"NumberOfActivePackages", activePackageInfo.NumberOfActivePackages})
    activePackageInfo.EntityData.Leafs.Append("active-packages", types.YLeaf{"ActivePackages", activePackageInfo.ActivePackages})

    activePackageInfo.EntityData.YListKeys = []string {}

    return &(activePackageInfo.EntityData)
}

// SoftwareInstall_Version
// Show install version
type SoftwareInstall_Version struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ImgInfo. The type is string.
    ImgInfo interface{}
}

func (version *SoftwareInstall_Version) GetEntityData() *types.CommonEntityData {
    version.EntityData.YFilter = version.YFilter
    version.EntityData.YangName = "version"
    version.EntityData.BundleName = "cisco_ios_xr"
    version.EntityData.ParentYangName = "software-install"
    version.EntityData.SegmentPath = "version"
    version.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version.EntityData.Children = types.NewOrderedMap()
    version.EntityData.Leafs = types.NewOrderedMap()
    version.EntityData.Leafs.Append("img-info", types.YLeaf{"ImgInfo", version.ImgInfo})

    version.EntityData.YListKeys = []string {}

    return &(version.EntityData)
}

// SoftwareInstall_Inactive
// Show XR inactive packages
type SoftwareInstall_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (inactive *SoftwareInstall_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "software-install"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("log", types.YLeaf{"Log", inactive.Log})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// SoftwareInstall_Request
// Show current request
type SoftwareInstall_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CurrInstOper. The type is string.
    CurrInstOper interface{}
}

func (request *SoftwareInstall_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "software-install"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("curr-inst-oper", types.YLeaf{"CurrInstOper", request.CurrInstOper})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// SoftwareInstall_Issu
// ISSU operation
type SoftwareInstall_Issu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show XR install issu stage.
    Stage SoftwareInstall_Issu_Stage

    // Show XR install issu inventory.
    Inventory SoftwareInstall_Issu_Inventory
}

func (issu *SoftwareInstall_Issu) GetEntityData() *types.CommonEntityData {
    issu.EntityData.YFilter = issu.YFilter
    issu.EntityData.YangName = "issu"
    issu.EntityData.BundleName = "cisco_ios_xr"
    issu.EntityData.ParentYangName = "software-install"
    issu.EntityData.SegmentPath = "issu"
    issu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issu.EntityData.Children = types.NewOrderedMap()
    issu.EntityData.Children.Append("stage", types.YChild{"Stage", &issu.Stage})
    issu.EntityData.Children.Append("inventory", types.YChild{"Inventory", &issu.Inventory})
    issu.EntityData.Leafs = types.NewOrderedMap()

    issu.EntityData.YListKeys = []string {}

    return &(issu.EntityData)
}

// SoftwareInstall_Issu_Stage
// Show XR install issu stage
type SoftwareInstall_Issu_Stage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is IsdStateEt.
    State interface{}

    // ISSU Node Count. The type is interface{} with range:
    // -2147483648..2147483647.
    IssuNodeCnt interface{}

    // ISSU Ready Node Count. The type is interface{} with range:
    // -2147483648..2147483647.
    IssuReadyNodeCnt interface{}

    // Percentage. The type is interface{} with range: -2147483648..2147483647.
    // Units are percentage.
    Percentage interface{}

    // Abort Status. The type is IsdIssuStatusEt.
    IssuStatus interface{}

    // ISSU Error. The type is IsdErrorEt.
    IssuError interface{}
}

func (stage *SoftwareInstall_Issu_Stage) GetEntityData() *types.CommonEntityData {
    stage.EntityData.YFilter = stage.YFilter
    stage.EntityData.YangName = "stage"
    stage.EntityData.BundleName = "cisco_ios_xr"
    stage.EntityData.ParentYangName = "issu"
    stage.EntityData.SegmentPath = "stage"
    stage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stage.EntityData.Children = types.NewOrderedMap()
    stage.EntityData.Leafs = types.NewOrderedMap()
    stage.EntityData.Leafs.Append("state", types.YLeaf{"State", stage.State})
    stage.EntityData.Leafs.Append("issu-node-cnt", types.YLeaf{"IssuNodeCnt", stage.IssuNodeCnt})
    stage.EntityData.Leafs.Append("issu-ready-node-cnt", types.YLeaf{"IssuReadyNodeCnt", stage.IssuReadyNodeCnt})
    stage.EntityData.Leafs.Append("percentage", types.YLeaf{"Percentage", stage.Percentage})
    stage.EntityData.Leafs.Append("issu-status", types.YLeaf{"IssuStatus", stage.IssuStatus})
    stage.EntityData.Leafs.Append("issu-error", types.YLeaf{"IssuError", stage.IssuError})

    stage.EntityData.YListKeys = []string {}

    return &(stage.EntityData)
}

// SoftwareInstall_Issu_Inventory
// Show XR install issu inventory
type SoftwareInstall_Issu_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // invinfo. The type is slice of SoftwareInstall_Issu_Inventory_Invinfo.
    Invinfo []*SoftwareInstall_Issu_Inventory_Invinfo
}

func (inventory *SoftwareInstall_Issu_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "issu"
    inventory.EntityData.SegmentPath = "inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = types.NewOrderedMap()
    inventory.EntityData.Children.Append("invinfo", types.YChild{"Invinfo", nil})
    for i := range inventory.Invinfo {
        inventory.EntityData.Children.Append(types.GetSegmentPath(inventory.Invinfo[i]), types.YChild{"Invinfo", inventory.Invinfo[i]})
    }
    inventory.EntityData.Leafs = types.NewOrderedMap()

    inventory.EntityData.YListKeys = []string {}

    return &(inventory.EntityData)
}

// SoftwareInstall_Issu_Inventory_Invinfo
// invinfo
type SoftwareInstall_Issu_Inventory_Invinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is interface{} with range: -2147483648..2147483647.
    NodeId interface{}

    // Node Type. The type is CardTypeEt.
    NodeType interface{}

    // ISSU Node Role. The type is IssuNodeRoleEt.
    IssuNodeRole interface{}

    // Node State. The type is IssudirNodeStatusEt.
    NodeState interface{}

    // Node role. The type is NodeRoleEt.
    NodeRole interface{}
}

func (invinfo *SoftwareInstall_Issu_Inventory_Invinfo) GetEntityData() *types.CommonEntityData {
    invinfo.EntityData.YFilter = invinfo.YFilter
    invinfo.EntityData.YangName = "invinfo"
    invinfo.EntityData.BundleName = "cisco_ios_xr"
    invinfo.EntityData.ParentYangName = "inventory"
    invinfo.EntityData.SegmentPath = "invinfo"
    invinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invinfo.EntityData.Children = types.NewOrderedMap()
    invinfo.EntityData.Leafs = types.NewOrderedMap()
    invinfo.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", invinfo.NodeId})
    invinfo.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", invinfo.NodeType})
    invinfo.EntityData.Leafs.Append("issu-node-role", types.YLeaf{"IssuNodeRole", invinfo.IssuNodeRole})
    invinfo.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", invinfo.NodeState})
    invinfo.EntityData.Leafs.Append("node-role", types.YLeaf{"NodeRole", invinfo.NodeRole})

    invinfo.EntityData.YListKeys = []string {}

    return &(invinfo.EntityData)
}

// SoftwareInstall_Committed
// Show Committed packages installed
type SoftwareInstall_Committed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // committed package info. The type is slice of
    // SoftwareInstall_Committed_CommittedPackageInfo.
    CommittedPackageInfo []*SoftwareInstall_Committed_CommittedPackageInfo
}

func (committed *SoftwareInstall_Committed) GetEntityData() *types.CommonEntityData {
    committed.EntityData.YFilter = committed.YFilter
    committed.EntityData.YangName = "committed"
    committed.EntityData.BundleName = "cisco_ios_xr"
    committed.EntityData.ParentYangName = "software-install"
    committed.EntityData.SegmentPath = "committed"
    committed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committed.EntityData.Children = types.NewOrderedMap()
    committed.EntityData.Children.Append("committed-package-info", types.YChild{"CommittedPackageInfo", nil})
    for i := range committed.CommittedPackageInfo {
        committed.EntityData.Children.Append(types.GetSegmentPath(committed.CommittedPackageInfo[i]), types.YChild{"CommittedPackageInfo", committed.CommittedPackageInfo[i]})
    }
    committed.EntityData.Leafs = types.NewOrderedMap()

    committed.EntityData.YListKeys = []string {}

    return &(committed.EntityData)
}

// SoftwareInstall_Committed_CommittedPackageInfo
// committed package info
type SoftwareInstall_Committed_CommittedPackageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ErrorMessage. The type is string.
    ErrorMessage interface{}

    // Location. The type is string.
    Location interface{}

    // NodeType. The type is string.
    NodeType interface{}

    // BootPartitionName. The type is string.
    BootPartitionName interface{}

    // NumberOfCommittedPackages. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCommittedPackages interface{}

    // CommittedPackages. The type is string.
    CommittedPackages interface{}
}

func (committedPackageInfo *SoftwareInstall_Committed_CommittedPackageInfo) GetEntityData() *types.CommonEntityData {
    committedPackageInfo.EntityData.YFilter = committedPackageInfo.YFilter
    committedPackageInfo.EntityData.YangName = "committed-package-info"
    committedPackageInfo.EntityData.BundleName = "cisco_ios_xr"
    committedPackageInfo.EntityData.ParentYangName = "committed"
    committedPackageInfo.EntityData.SegmentPath = "committed-package-info"
    committedPackageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedPackageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedPackageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedPackageInfo.EntityData.Children = types.NewOrderedMap()
    committedPackageInfo.EntityData.Leafs = types.NewOrderedMap()
    committedPackageInfo.EntityData.Leafs.Append("error-message", types.YLeaf{"ErrorMessage", committedPackageInfo.ErrorMessage})
    committedPackageInfo.EntityData.Leafs.Append("location", types.YLeaf{"Location", committedPackageInfo.Location})
    committedPackageInfo.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", committedPackageInfo.NodeType})
    committedPackageInfo.EntityData.Leafs.Append("boot-partition-name", types.YLeaf{"BootPartitionName", committedPackageInfo.BootPartitionName})
    committedPackageInfo.EntityData.Leafs.Append("number-of-committed-packages", types.YLeaf{"NumberOfCommittedPackages", committedPackageInfo.NumberOfCommittedPackages})
    committedPackageInfo.EntityData.Leafs.Append("committed-packages", types.YLeaf{"CommittedPackages", committedPackageInfo.CommittedPackages})

    committedPackageInfo.EntityData.YListKeys = []string {}

    return &(committedPackageInfo.EntityData)
}

// SoftwareInstall_AllOperationsLog
// Show log file for all operations
type SoftwareInstall_AllOperationsLog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show summary log file for all operations.
    Summary SoftwareInstall_AllOperationsLog_Summary

    // Show detailed log file for all operations.
    Detail SoftwareInstall_AllOperationsLog_Detail
}

func (allOperationsLog *SoftwareInstall_AllOperationsLog) GetEntityData() *types.CommonEntityData {
    allOperationsLog.EntityData.YFilter = allOperationsLog.YFilter
    allOperationsLog.EntityData.YangName = "all-operations-log"
    allOperationsLog.EntityData.BundleName = "cisco_ios_xr"
    allOperationsLog.EntityData.ParentYangName = "software-install"
    allOperationsLog.EntityData.SegmentPath = "all-operations-log"
    allOperationsLog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOperationsLog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOperationsLog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOperationsLog.EntityData.Children = types.NewOrderedMap()
    allOperationsLog.EntityData.Children.Append("summary", types.YChild{"Summary", &allOperationsLog.Summary})
    allOperationsLog.EntityData.Children.Append("detail", types.YChild{"Detail", &allOperationsLog.Detail})
    allOperationsLog.EntityData.Leafs = types.NewOrderedMap()

    allOperationsLog.EntityData.YListKeys = []string {}

    return &(allOperationsLog.EntityData)
}

// SoftwareInstall_AllOperationsLog_Summary
// Show summary log file for all operations
type SoftwareInstall_AllOperationsLog_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (summary *SoftwareInstall_AllOperationsLog_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "all-operations-log"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("log", types.YLeaf{"Log", summary.Log})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// SoftwareInstall_AllOperationsLog_Detail
// Show detailed log file for all operations
type SoftwareInstall_AllOperationsLog_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (detail *SoftwareInstall_AllOperationsLog_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "all-operations-log"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("log", types.YLeaf{"Log", detail.Log})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// SoftwareInstall_Packages
// Show the list of installed packages
type SoftwareInstall_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show the info for a installed package. The type is slice of
    // SoftwareInstall_Packages_Package.
    Package []*SoftwareInstall_Packages_Package
}

func (packages *SoftwareInstall_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "software-install"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Children.Append("package", types.YChild{"Package", nil})
    for i := range packages.Package {
        packages.EntityData.Children.Append(types.GetSegmentPath(packages.Package[i]), types.YChild{"Package", packages.Package[i]})
    }
    packages.EntityData.Leafs = types.NewOrderedMap()

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// SoftwareInstall_Packages_Package
// Show the info for a installed package
type SoftwareInstall_Packages_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Package name. The type is string.
    PackageName interface{}

    // Show the verbose info for a installed package.
    Verbose SoftwareInstall_Packages_Package_Verbose

    // Show the info for a installed package.
    Brief SoftwareInstall_Packages_Package_Brief

    // Show the deatil info for a installed package.
    Detail SoftwareInstall_Packages_Package_Detail
}

func (self *SoftwareInstall_Packages_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "packages"
    self.EntityData.SegmentPath = "package" + types.AddKeyToken(self.PackageName, "package-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("verbose", types.YChild{"Verbose", &self.Verbose})
    self.EntityData.Children.Append("brief", types.YChild{"Brief", &self.Brief})
    self.EntityData.Children.Append("detail", types.YChild{"Detail", &self.Detail})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("package-name", types.YLeaf{"PackageName", self.PackageName})

    self.EntityData.YListKeys = []string {"PackageName"}

    return &(self.EntityData)
}

// SoftwareInstall_Packages_Package_Verbose
// Show the verbose info for a installed package
type SoftwareInstall_Packages_Package_Verbose struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (verbose *SoftwareInstall_Packages_Package_Verbose) GetEntityData() *types.CommonEntityData {
    verbose.EntityData.YFilter = verbose.YFilter
    verbose.EntityData.YangName = "verbose"
    verbose.EntityData.BundleName = "cisco_ios_xr"
    verbose.EntityData.ParentYangName = "package"
    verbose.EntityData.SegmentPath = "verbose"
    verbose.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verbose.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verbose.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verbose.EntityData.Children = types.NewOrderedMap()
    verbose.EntityData.Leafs = types.NewOrderedMap()
    verbose.EntityData.Leafs.Append("log", types.YLeaf{"Log", verbose.Log})

    verbose.EntityData.YListKeys = []string {}

    return &(verbose.EntityData)
}

// SoftwareInstall_Packages_Package_Brief
// Show the info for a installed package
type SoftwareInstall_Packages_Package_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (brief *SoftwareInstall_Packages_Package_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "package"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Leafs = types.NewOrderedMap()
    brief.EntityData.Leafs.Append("log", types.YLeaf{"Log", brief.Log})

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// SoftwareInstall_Packages_Package_Detail
// Show the deatil info for a installed package
type SoftwareInstall_Packages_Package_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (detail *SoftwareInstall_Packages_Package_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "package"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("log", types.YLeaf{"Log", detail.Log})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// SoftwareInstall_OperationLogs
// Show log file
type SoftwareInstall_OperationLogs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show log file for the specified install ID. The type is slice of
    // SoftwareInstall_OperationLogs_OperationLog.
    OperationLog []*SoftwareInstall_OperationLogs_OperationLog
}

func (operationLogs *SoftwareInstall_OperationLogs) GetEntityData() *types.CommonEntityData {
    operationLogs.EntityData.YFilter = operationLogs.YFilter
    operationLogs.EntityData.YangName = "operation-logs"
    operationLogs.EntityData.BundleName = "cisco_ios_xr"
    operationLogs.EntityData.ParentYangName = "software-install"
    operationLogs.EntityData.SegmentPath = "operation-logs"
    operationLogs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationLogs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationLogs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationLogs.EntityData.Children = types.NewOrderedMap()
    operationLogs.EntityData.Children.Append("operation-log", types.YChild{"OperationLog", nil})
    for i := range operationLogs.OperationLog {
        operationLogs.EntityData.Children.Append(types.GetSegmentPath(operationLogs.OperationLog[i]), types.YChild{"OperationLog", operationLogs.OperationLog[i]})
    }
    operationLogs.EntityData.Leafs = types.NewOrderedMap()

    operationLogs.EntityData.YListKeys = []string {}

    return &(operationLogs.EntityData)
}

// SoftwareInstall_OperationLogs_OperationLog
// Show log file for the specified install ID
type SoftwareInstall_OperationLogs_OperationLog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Log ID number. The type is interface{} with range:
    // 0..4294967295.
    LogId interface{}

    // Show summary log file for the specified install ID.
    Summary SoftwareInstall_OperationLogs_OperationLog_Summary

    // Show detailed log file for the specified install ID.
    Detail SoftwareInstall_OperationLogs_OperationLog_Detail
}

func (operationLog *SoftwareInstall_OperationLogs_OperationLog) GetEntityData() *types.CommonEntityData {
    operationLog.EntityData.YFilter = operationLog.YFilter
    operationLog.EntityData.YangName = "operation-log"
    operationLog.EntityData.BundleName = "cisco_ios_xr"
    operationLog.EntityData.ParentYangName = "operation-logs"
    operationLog.EntityData.SegmentPath = "operation-log" + types.AddKeyToken(operationLog.LogId, "log-id")
    operationLog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationLog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationLog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationLog.EntityData.Children = types.NewOrderedMap()
    operationLog.EntityData.Children.Append("summary", types.YChild{"Summary", &operationLog.Summary})
    operationLog.EntityData.Children.Append("detail", types.YChild{"Detail", &operationLog.Detail})
    operationLog.EntityData.Leafs = types.NewOrderedMap()
    operationLog.EntityData.Leafs.Append("log-id", types.YLeaf{"LogId", operationLog.LogId})

    operationLog.EntityData.YListKeys = []string {"LogId"}

    return &(operationLog.EntityData)
}

// SoftwareInstall_OperationLogs_OperationLog_Summary
// Show summary log file for the specified
// install ID
type SoftwareInstall_OperationLogs_OperationLog_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (summary *SoftwareInstall_OperationLogs_OperationLog_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "operation-log"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("log", types.YLeaf{"Log", summary.Log})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// SoftwareInstall_OperationLogs_OperationLog_Detail
// Show detailed log file for the specified
// install ID
type SoftwareInstall_OperationLogs_OperationLog_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (detail *SoftwareInstall_OperationLogs_OperationLog_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "operation-log"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("log", types.YLeaf{"Log", detail.Log})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// SoftwareInstall_Repository
// Show packages stored in install software
// repositories
type SoftwareInstall_Repository struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show install software repository for XR.
    Xr SoftwareInstall_Repository_Xr

    // Show contents of all install software repositories.
    All SoftwareInstall_Repository_All
}

func (repository *SoftwareInstall_Repository) GetEntityData() *types.CommonEntityData {
    repository.EntityData.YFilter = repository.YFilter
    repository.EntityData.YangName = "repository"
    repository.EntityData.BundleName = "cisco_ios_xr"
    repository.EntityData.ParentYangName = "software-install"
    repository.EntityData.SegmentPath = "repository"
    repository.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    repository.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    repository.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    repository.EntityData.Children = types.NewOrderedMap()
    repository.EntityData.Children.Append("xr", types.YChild{"Xr", &repository.Xr})
    repository.EntityData.Children.Append("all", types.YChild{"All", &repository.All})
    repository.EntityData.Leafs = types.NewOrderedMap()

    repository.EntityData.YListKeys = []string {}

    return &(repository.EntityData)
}

// SoftwareInstall_Repository_Xr
// Show install software repository for XR
type SoftwareInstall_Repository_Xr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (xr *SoftwareInstall_Repository_Xr) GetEntityData() *types.CommonEntityData {
    xr.EntityData.YFilter = xr.YFilter
    xr.EntityData.YangName = "xr"
    xr.EntityData.BundleName = "cisco_ios_xr"
    xr.EntityData.ParentYangName = "repository"
    xr.EntityData.SegmentPath = "xr"
    xr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xr.EntityData.Children = types.NewOrderedMap()
    xr.EntityData.Leafs = types.NewOrderedMap()
    xr.EntityData.Leafs.Append("log", types.YLeaf{"Log", xr.Log})

    xr.EntityData.YListKeys = []string {}

    return &(xr.EntityData)
}

// SoftwareInstall_Repository_All
// Show contents of all install software
// repositories
type SoftwareInstall_Repository_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // log. The type is string.
    Log interface{}
}

func (all *SoftwareInstall_Repository_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "repository"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("log", types.YLeaf{"Log", all.Log})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

