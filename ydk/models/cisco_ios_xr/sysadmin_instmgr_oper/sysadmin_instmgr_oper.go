// This module contains a collection of YANG
// definitions for Cisco IOS-XR sysadmin-instmgr 
// operational model
package sysadmin_instmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_instmgr_oper"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-instmgr-oper install}", reflect.TypeOf(Install{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-instmgr-oper:install", reflect.TypeOf(Install{}))
}

// Install
type Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Gives sysadmin version information.
    Version Install_Version

    // Calvados inactive package(s) list for this node.
    Inactive Install_Inactive

    // Calvados prepared package(s) list for this node.
    Prepare Install_Prepare

    // Package Name(s) to get info for.
    Package Install_Package

    // Calvados active package(s) list for this node.
    Active Install_Active

    // Calvados superseded package(s) list for this node.
    Superseded Install_Superseded

    // Sysadmin current install operation details.
    Request Install_Request

    // Shows information about the install software repository.
    Repository Install_Repository

    
    Log Install_Log

    // Filename to get info for.
    Which Install_Which

    // Calvados committed package(s) list for this node.
    Committed Install_Committed
}

func (install *Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xr"
    install.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-instmgr-oper"
    install.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install"
    install.EntityData.AbsolutePath = install.EntityData.SegmentPath
    install.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    install.EntityData.Children = types.NewOrderedMap()
    install.EntityData.Children.Append("version", types.YChild{"Version", &install.Version})
    install.EntityData.Children.Append("inactive", types.YChild{"Inactive", &install.Inactive})
    install.EntityData.Children.Append("prepare", types.YChild{"Prepare", &install.Prepare})
    install.EntityData.Children.Append("package", types.YChild{"Package", &install.Package})
    install.EntityData.Children.Append("active", types.YChild{"Active", &install.Active})
    install.EntityData.Children.Append("superseded", types.YChild{"Superseded", &install.Superseded})
    install.EntityData.Children.Append("request", types.YChild{"Request", &install.Request})
    install.EntityData.Children.Append("repository", types.YChild{"Repository", &install.Repository})
    install.EntityData.Children.Append("log", types.YChild{"Log", &install.Log})
    install.EntityData.Children.Append("which", types.YChild{"Which", &install.Which})
    install.EntityData.Children.Append("committed", types.YChild{"Committed", &install.Committed})
    install.EntityData.Leafs = types.NewOrderedMap()

    install.EntityData.YListKeys = []string {}

    return &(install.EntityData)
}

// Install_Version
// Gives sysadmin version information
type Install_Version struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    ImgInfo []interface{}
}

func (version *Install_Version) GetEntityData() *types.CommonEntityData {
    version.EntityData.YFilter = version.YFilter
    version.EntityData.YangName = "version"
    version.EntityData.BundleName = "cisco_ios_xr"
    version.EntityData.ParentYangName = "install"
    version.EntityData.SegmentPath = "version"
    version.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + version.EntityData.SegmentPath
    version.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version.EntityData.Children = types.NewOrderedMap()
    version.EntityData.Leafs = types.NewOrderedMap()
    version.EntityData.Leafs.Append("img_info", types.YLeaf{"ImgInfo", version.ImgInfo})

    version.EntityData.YListKeys = []string {}

    return &(version.EntityData)
}

// Install_Inactive
// Calvados inactive package(s) list for this node
type Install_Inactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiInactiveOutput []interface{}

    // shows summary information of the inactive install software.
    Summary Install_Inactive_Summary
}

func (inactive *Install_Inactive) GetEntityData() *types.CommonEntityData {
    inactive.EntityData.YFilter = inactive.YFilter
    inactive.EntityData.YangName = "inactive"
    inactive.EntityData.BundleName = "cisco_ios_xr"
    inactive.EntityData.ParentYangName = "install"
    inactive.EntityData.SegmentPath = "inactive"
    inactive.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + inactive.EntityData.SegmentPath
    inactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inactive.EntityData.Children = types.NewOrderedMap()
    inactive.EntityData.Children.Append("summary", types.YChild{"Summary", &inactive.Summary})
    inactive.EntityData.Leafs = types.NewOrderedMap()
    inactive.EntityData.Leafs.Append("si_inactive_output", types.YLeaf{"SiInactiveOutput", inactive.SiInactiveOutput})

    inactive.EntityData.YListKeys = []string {}

    return &(inactive.EntityData)
}

// Install_Inactive_Summary
// shows summary information of the inactive install software
type Install_Inactive_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiInactiveSummaryOutput []interface{}
}

func (summary *Install_Inactive_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "inactive"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/inactive/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("si_inactive_summary_output", types.YLeaf{"SiInactiveSummaryOutput", summary.SiInactiveSummaryOutput})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_Prepare
// Calvados prepared package(s) list for this node
type Install_Prepare struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiPrepareOutput []interface{}
}

func (prepare *Install_Prepare) GetEntityData() *types.CommonEntityData {
    prepare.EntityData.YFilter = prepare.YFilter
    prepare.EntityData.YangName = "prepare"
    prepare.EntityData.BundleName = "cisco_ios_xr"
    prepare.EntityData.ParentYangName = "install"
    prepare.EntityData.SegmentPath = "prepare"
    prepare.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + prepare.EntityData.SegmentPath
    prepare.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepare.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepare.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepare.EntityData.Children = types.NewOrderedMap()
    prepare.EntityData.Leafs = types.NewOrderedMap()
    prepare.EntityData.Leafs.Append("si_prepare_output", types.YLeaf{"SiPrepareOutput", prepare.SiPrepareOutput})

    prepare.EntityData.YListKeys = []string {}

    return &(prepare.EntityData)
}

// Install_Package
// Package Name(s) to get info for
type Install_Package struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Install_Package_PkgList.
    PkgList []*Install_Package_PkgList
}

func (self *Install_Package) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "package"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "install"
    self.EntityData.SegmentPath = "package"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("pkg_list", types.YChild{"PkgList", nil})
    for i := range self.PkgList {
        self.EntityData.Children.Append(types.GetSegmentPath(self.PkgList[i]), types.YChild{"PkgList", self.PkgList[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Install_Package_PkgList
type Install_Package_PkgList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    PkgName interface{}

    // The type is slice of string.
    SiPackageOutput []interface{}

    
    Detail Install_Package_PkgList_Detail

    
    Verbose Install_Package_PkgList_Verbose
}

func (pkgList *Install_Package_PkgList) GetEntityData() *types.CommonEntityData {
    pkgList.EntityData.YFilter = pkgList.YFilter
    pkgList.EntityData.YangName = "pkg_list"
    pkgList.EntityData.BundleName = "cisco_ios_xr"
    pkgList.EntityData.ParentYangName = "package"
    pkgList.EntityData.SegmentPath = "pkg_list" + types.AddKeyToken(pkgList.PkgName, "pkg_name")
    pkgList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/package/" + pkgList.EntityData.SegmentPath
    pkgList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pkgList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pkgList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pkgList.EntityData.Children = types.NewOrderedMap()
    pkgList.EntityData.Children.Append("detail", types.YChild{"Detail", &pkgList.Detail})
    pkgList.EntityData.Children.Append("verbose", types.YChild{"Verbose", &pkgList.Verbose})
    pkgList.EntityData.Leafs = types.NewOrderedMap()
    pkgList.EntityData.Leafs.Append("pkg_name", types.YLeaf{"PkgName", pkgList.PkgName})
    pkgList.EntityData.Leafs.Append("si_package_output", types.YLeaf{"SiPackageOutput", pkgList.SiPackageOutput})

    pkgList.EntityData.YListKeys = []string {"PkgName"}

    return &(pkgList.EntityData)
}

// Install_Package_PkgList_Detail
type Install_Package_PkgList_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiPackageDetailOutput []interface{}
}

func (detail *Install_Package_PkgList_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "pkg_list"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/package/pkg_list/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("si_package_detail_output", types.YLeaf{"SiPackageDetailOutput", detail.SiPackageDetailOutput})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Package_PkgList_Verbose
type Install_Package_PkgList_Verbose struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiPackageVerboseOutput []interface{}
}

func (verbose *Install_Package_PkgList_Verbose) GetEntityData() *types.CommonEntityData {
    verbose.EntityData.YFilter = verbose.YFilter
    verbose.EntityData.YangName = "verbose"
    verbose.EntityData.BundleName = "cisco_ios_xr"
    verbose.EntityData.ParentYangName = "pkg_list"
    verbose.EntityData.SegmentPath = "verbose"
    verbose.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/package/pkg_list/" + verbose.EntityData.SegmentPath
    verbose.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verbose.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verbose.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verbose.EntityData.Children = types.NewOrderedMap()
    verbose.EntityData.Leafs = types.NewOrderedMap()
    verbose.EntityData.Leafs.Append("si_package_verbose_output", types.YLeaf{"SiPackageVerboseOutput", verbose.SiPackageVerboseOutput})

    verbose.EntityData.YListKeys = []string {}

    return &(verbose.EntityData)
}

// Install_Active
// Calvados active package(s) list for this node
type Install_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiActiveOutput []interface{}

    // shows summary information of the active install software.
    Summary Install_Active_Summary
}

func (active *Install_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "install"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("summary", types.YChild{"Summary", &active.Summary})
    active.EntityData.Leafs = types.NewOrderedMap()
    active.EntityData.Leafs.Append("si_active_output", types.YLeaf{"SiActiveOutput", active.SiActiveOutput})

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Install_Active_Summary
// shows summary information of the active install software
type Install_Active_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiActiveSummaryOutput []interface{}
}

func (summary *Install_Active_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "active"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/active/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("si_active_summary_output", types.YLeaf{"SiActiveSummaryOutput", summary.SiActiveSummaryOutput})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_Superseded
// Calvados superseded package(s) list for this node
type Install_Superseded struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiSupersededOutput []interface{}

    // shows summary information of the show install superseded.
    Summary Install_Superseded_Summary
}

func (superseded *Install_Superseded) GetEntityData() *types.CommonEntityData {
    superseded.EntityData.YFilter = superseded.YFilter
    superseded.EntityData.YangName = "superseded"
    superseded.EntityData.BundleName = "cisco_ios_xr"
    superseded.EntityData.ParentYangName = "install"
    superseded.EntityData.SegmentPath = "superseded"
    superseded.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + superseded.EntityData.SegmentPath
    superseded.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    superseded.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    superseded.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    superseded.EntityData.Children = types.NewOrderedMap()
    superseded.EntityData.Children.Append("summary", types.YChild{"Summary", &superseded.Summary})
    superseded.EntityData.Leafs = types.NewOrderedMap()
    superseded.EntityData.Leafs.Append("si_superseded_output", types.YLeaf{"SiSupersededOutput", superseded.SiSupersededOutput})

    superseded.EntityData.YListKeys = []string {}

    return &(superseded.EntityData)
}

// Install_Superseded_Summary
// shows summary information of the show install superseded
type Install_Superseded_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiSupersededSummaryOutput []interface{}
}

func (summary *Install_Superseded_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "superseded"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/superseded/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("si_superseded_summary_output", types.YLeaf{"SiSupersededSummaryOutput", summary.SiSupersededSummaryOutput})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Install_Request
// Sysadmin current install operation details
type Install_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    CurrInstOper []interface{}
}

func (request *Install_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "install"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("curr_inst_oper", types.YLeaf{"CurrInstOper", request.CurrInstOper})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Install_Repository
// Shows information about the install software repository.
type Install_Repository struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiRepositoryOutput []interface{}

    // shows contents of all the install software repositories.
    All Install_Repository_All
}

func (repository *Install_Repository) GetEntityData() *types.CommonEntityData {
    repository.EntityData.YFilter = repository.YFilter
    repository.EntityData.YangName = "repository"
    repository.EntityData.BundleName = "cisco_ios_xr"
    repository.EntityData.ParentYangName = "install"
    repository.EntityData.SegmentPath = "repository"
    repository.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + repository.EntityData.SegmentPath
    repository.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    repository.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    repository.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    repository.EntityData.Children = types.NewOrderedMap()
    repository.EntityData.Children.Append("all", types.YChild{"All", &repository.All})
    repository.EntityData.Leafs = types.NewOrderedMap()
    repository.EntityData.Leafs.Append("si_repository_output", types.YLeaf{"SiRepositoryOutput", repository.SiRepositoryOutput})

    repository.EntityData.YListKeys = []string {}

    return &(repository.EntityData)
}

// Install_Repository_All
// shows contents of all the install software repositories
type Install_Repository_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiRepositoryAllOutput []interface{}
}

func (all *Install_Repository_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "repository"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/repository/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("si_repository_all_output", types.YLeaf{"SiRepositoryAllOutput", all.SiRepositoryAllOutput})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Install_Log
type Install_Log struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiLogOutput []interface{}

    // The type is slice of Install_Log_Id.
    Id []*Install_Log_Id

    
    Reverse Install_Log_Reverse

    
    Detail Install_Log_Detail
}

func (log *Install_Log) GetEntityData() *types.CommonEntityData {
    log.EntityData.YFilter = log.YFilter
    log.EntityData.YangName = "log"
    log.EntityData.BundleName = "cisco_ios_xr"
    log.EntityData.ParentYangName = "install"
    log.EntityData.SegmentPath = "log"
    log.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + log.EntityData.SegmentPath
    log.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    log.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    log.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    log.EntityData.Children = types.NewOrderedMap()
    log.EntityData.Children.Append("id", types.YChild{"Id", nil})
    for i := range log.Id {
        log.EntityData.Children.Append(types.GetSegmentPath(log.Id[i]), types.YChild{"Id", log.Id[i]})
    }
    log.EntityData.Children.Append("reverse", types.YChild{"Reverse", &log.Reverse})
    log.EntityData.Children.Append("detail", types.YChild{"Detail", &log.Detail})
    log.EntityData.Leafs = types.NewOrderedMap()
    log.EntityData.Leafs.Append("si_log_output", types.YLeaf{"SiLogOutput", log.SiLogOutput})

    log.EntityData.YListKeys = []string {}

    return &(log.EntityData)
}

// Install_Log_Id
type Install_Log_Id struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Opid interface{}

    // The type is slice of string.
    SiLogIdOutput []interface{}

    
    Detail Install_Log_Id_Detail
}

func (id *Install_Log_Id) GetEntityData() *types.CommonEntityData {
    id.EntityData.YFilter = id.YFilter
    id.EntityData.YangName = "id"
    id.EntityData.BundleName = "cisco_ios_xr"
    id.EntityData.ParentYangName = "log"
    id.EntityData.SegmentPath = "id" + types.AddKeyToken(id.Opid, "opid")
    id.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/log/" + id.EntityData.SegmentPath
    id.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    id.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    id.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    id.EntityData.Children = types.NewOrderedMap()
    id.EntityData.Children.Append("detail", types.YChild{"Detail", &id.Detail})
    id.EntityData.Leafs = types.NewOrderedMap()
    id.EntityData.Leafs.Append("opid", types.YLeaf{"Opid", id.Opid})
    id.EntityData.Leafs.Append("si_log_id_output", types.YLeaf{"SiLogIdOutput", id.SiLogIdOutput})

    id.EntityData.YListKeys = []string {"Opid"}

    return &(id.EntityData)
}

// Install_Log_Id_Detail
type Install_Log_Id_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiLogIdDetailOutput []interface{}
}

func (detail *Install_Log_Id_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "id"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/log/id/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("si_log_id_detail_output", types.YLeaf{"SiLogIdDetailOutput", detail.SiLogIdDetailOutput})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Log_Reverse
type Install_Log_Reverse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiLogReverseOutput []interface{}

    
    Detail Install_Log_Reverse_Detail
}

func (reverse *Install_Log_Reverse) GetEntityData() *types.CommonEntityData {
    reverse.EntityData.YFilter = reverse.YFilter
    reverse.EntityData.YangName = "reverse"
    reverse.EntityData.BundleName = "cisco_ios_xr"
    reverse.EntityData.ParentYangName = "log"
    reverse.EntityData.SegmentPath = "reverse"
    reverse.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/log/" + reverse.EntityData.SegmentPath
    reverse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reverse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reverse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reverse.EntityData.Children = types.NewOrderedMap()
    reverse.EntityData.Children.Append("detail", types.YChild{"Detail", &reverse.Detail})
    reverse.EntityData.Leafs = types.NewOrderedMap()
    reverse.EntityData.Leafs.Append("si_log_reverse_output", types.YLeaf{"SiLogReverseOutput", reverse.SiLogReverseOutput})

    reverse.EntityData.YListKeys = []string {}

    return &(reverse.EntityData)
}

// Install_Log_Reverse_Detail
type Install_Log_Reverse_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiLogReverseDetailOutput []interface{}
}

func (detail *Install_Log_Reverse_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "reverse"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/log/reverse/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("si_log_reverse_detail_output", types.YLeaf{"SiLogReverseDetailOutput", detail.SiLogReverseDetailOutput})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Log_Detail
type Install_Log_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiLogDetailOutput []interface{}
}

func (detail *Install_Log_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "log"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/log/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("si_log_detail_output", types.YLeaf{"SiLogDetailOutput", detail.SiLogDetailOutput})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Which
// Filename to get info for
type Install_Which struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Install_Which_FileList.
    FileList []*Install_Which_FileList
}

func (which *Install_Which) GetEntityData() *types.CommonEntityData {
    which.EntityData.YFilter = which.YFilter
    which.EntityData.YangName = "which"
    which.EntityData.BundleName = "cisco_ios_xr"
    which.EntityData.ParentYangName = "install"
    which.EntityData.SegmentPath = "which"
    which.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + which.EntityData.SegmentPath
    which.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    which.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    which.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    which.EntityData.Children = types.NewOrderedMap()
    which.EntityData.Children.Append("file_list", types.YChild{"FileList", nil})
    for i := range which.FileList {
        which.EntityData.Children.Append(types.GetSegmentPath(which.FileList[i]), types.YChild{"FileList", which.FileList[i]})
    }
    which.EntityData.Leafs = types.NewOrderedMap()

    which.EntityData.YListKeys = []string {}

    return &(which.EntityData)
}

// Install_Which_FileList
type Install_Which_FileList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    FileName interface{}

    // The type is slice of string.
    SiWhichOutput []interface{}

    
    Detail Install_Which_FileList_Detail
}

func (fileList *Install_Which_FileList) GetEntityData() *types.CommonEntityData {
    fileList.EntityData.YFilter = fileList.YFilter
    fileList.EntityData.YangName = "file_list"
    fileList.EntityData.BundleName = "cisco_ios_xr"
    fileList.EntityData.ParentYangName = "which"
    fileList.EntityData.SegmentPath = "file_list" + types.AddKeyToken(fileList.FileName, "file_name")
    fileList.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/which/" + fileList.EntityData.SegmentPath
    fileList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileList.EntityData.Children = types.NewOrderedMap()
    fileList.EntityData.Children.Append("detail", types.YChild{"Detail", &fileList.Detail})
    fileList.EntityData.Leafs = types.NewOrderedMap()
    fileList.EntityData.Leafs.Append("file_name", types.YLeaf{"FileName", fileList.FileName})
    fileList.EntityData.Leafs.Append("si_which_output", types.YLeaf{"SiWhichOutput", fileList.SiWhichOutput})

    fileList.EntityData.YListKeys = []string {"FileName"}

    return &(fileList.EntityData)
}

// Install_Which_FileList_Detail
type Install_Which_FileList_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiWhichDetailOutput []interface{}
}

func (detail *Install_Which_FileList_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "file_list"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/which/file_list/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("si_which_detail_output", types.YLeaf{"SiWhichDetailOutput", detail.SiWhichDetailOutput})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Install_Committed
// Calvados committed package(s) list for this node
type Install_Committed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiCommittedOutput []interface{}

    // shows summary information of the committed install software.
    Summary Install_Committed_Summary
}

func (committed *Install_Committed) GetEntityData() *types.CommonEntityData {
    committed.EntityData.YFilter = committed.YFilter
    committed.EntityData.YangName = "committed"
    committed.EntityData.BundleName = "cisco_ios_xr"
    committed.EntityData.ParentYangName = "install"
    committed.EntityData.SegmentPath = "committed"
    committed.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/" + committed.EntityData.SegmentPath
    committed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committed.EntityData.Children = types.NewOrderedMap()
    committed.EntityData.Children.Append("summary", types.YChild{"Summary", &committed.Summary})
    committed.EntityData.Leafs = types.NewOrderedMap()
    committed.EntityData.Leafs.Append("si_committed_output", types.YLeaf{"SiCommittedOutput", committed.SiCommittedOutput})

    committed.EntityData.YListKeys = []string {}

    return &(committed.EntityData)
}

// Install_Committed_Summary
// shows summary information of the committed install software
type Install_Committed_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    SiCommittedSummaryOutput []interface{}
}

func (summary *Install_Committed_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "committed"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-instmgr-oper:install/committed/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("si_committed_summary_output", types.YLeaf{"SiCommittedSummaryOutput", summary.SiCommittedSummaryOutput})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

