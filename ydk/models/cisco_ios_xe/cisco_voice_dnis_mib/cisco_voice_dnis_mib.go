// The MIB module provides management support for Dialer
// Number Information Service (DNIS) mapping.  A DNIS
// entry is associated with a Voice XML (VXML) page to
// provide audio play back features. Multiple DNIS
// entries can be grouped together to form a DNIS
// mapping with a unique map name.
// 
// *** ABBREVIATIONS, ACRONYMS, AND SYMBOLS ***
// 
// DNIS - Dialer Number Information Service
// 
// XML  - Extensible Markup Language
// 
// VXML - Voice XML
// 
// URL  - Uniform Resource Locator  
package cisco_voice_dnis_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_voice_dnis_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-VOICE-DNIS-MIB CISCO-VOICE-DNIS-MIB}", reflect.TypeOf(CISCOVOICEDNISMIB{}))
    ydk.RegisterEntity("CISCO-VOICE-DNIS-MIB:CISCO-VOICE-DNIS-MIB", reflect.TypeOf(CISCOVOICEDNISMIB{}))
}

// CISCOVOICEDNISMIB
type CISCOVOICEDNISMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The table contains the map name and a url specifying a file name. The file
    // contains DNIS entries that belong to the DNIS mapping.
    Cvdnismappingtable CISCOVOICEDNISMIB_Cvdnismappingtable

    // The table contains a DNIS name and a url. The url is a pointer to a VXML
    // page for the DNIS name. .
    Cvdnisnodetable CISCOVOICEDNISMIB_Cvdnisnodetable
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetFilter() yfilter.YFilter { return cISCOVOICEDNISMIB.YFilter }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) SetFilter(yf yfilter.YFilter) { cISCOVOICEDNISMIB.YFilter = yf }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetGoName(yname string) string {
    if yname == "cvDnisMappingTable" { return "Cvdnismappingtable" }
    if yname == "cvDnisNodeTable" { return "Cvdnisnodetable" }
    return ""
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetSegmentPath() string {
    return "CISCO-VOICE-DNIS-MIB:CISCO-VOICE-DNIS-MIB"
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvDnisMappingTable" {
        return &cISCOVOICEDNISMIB.Cvdnismappingtable
    }
    if childYangName == "cvDnisNodeTable" {
        return &cISCOVOICEDNISMIB.Cvdnisnodetable
    }
    return nil
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cvDnisMappingTable"] = &cISCOVOICEDNISMIB.Cvdnismappingtable
    children["cvDnisNodeTable"] = &cISCOVOICEDNISMIB.Cvdnisnodetable
    return children
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetYangName() string { return "CISCO-VOICE-DNIS-MIB" }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) SetParent(parent types.Entity) { cISCOVOICEDNISMIB.parent = parent }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetParent() types.Entity { return cISCOVOICEDNISMIB.parent }

func (cISCOVOICEDNISMIB *CISCOVOICEDNISMIB) GetParentYangName() string { return "CISCO-VOICE-DNIS-MIB" }

// CISCOVOICEDNISMIB_Cvdnismappingtable
// The table contains the map name and a url specifying
// a file name. The file contains DNIS entries that belong
// to the DNIS mapping.
type CISCOVOICEDNISMIB_Cvdnismappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single DNIS mapping. There is a unique DNIS map name.
    // New DNIS mapping can be created using cvDnisMappingStatus.  The entry can
    // be created with or without a file location  specified by cvDnisMappingUrl.
    // The mapping file contains DNIS name and VXML page per line. For example, a 
    // cvDnisMappingUrl could be tftp://someserver/dnismap.txt. This file is a
    // text file and the content format is   dnis <dnisname> url <urlname>. An
    // example of the contents of the file itself can be   dnis 18004251234 url
    // http://www.b.com/p/vwelcome.vxml   dnis 18004253421 url
    // http://www.c.com/j/vxmlintf.vxml If a mapping file location is specified,
    // then new rows corresponding to this map name are created and populated in
    // cvDnisNodeTable from the contents of the file. The rows corresponding to
    // this map name in cvDnisNodeTable cannot be created or modified or deleted
    // but can be read.   If a mapping file location is not specified in
    // cvDnisMappingUrl, then individual DNIS entries corresponding to this map
    // name can be created, modified and deleted in cvDnisNodeTable.   Deleting an
    // entry deletes all the related entries in cvDnisNodeTable. . The type is
    // slice of CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry.
    Cvdnismappingentry []CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetFilter() yfilter.YFilter { return cvdnismappingtable.YFilter }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) SetFilter(yf yfilter.YFilter) { cvdnismappingtable.YFilter = yf }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetGoName(yname string) string {
    if yname == "cvDnisMappingEntry" { return "Cvdnismappingentry" }
    return ""
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetSegmentPath() string {
    return "cvDnisMappingTable"
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvDnisMappingEntry" {
        for _, c := range cvdnismappingtable.Cvdnismappingentry {
            if cvdnismappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry{}
        cvdnismappingtable.Cvdnismappingentry = append(cvdnismappingtable.Cvdnismappingentry, child)
        return &cvdnismappingtable.Cvdnismappingentry[len(cvdnismappingtable.Cvdnismappingentry)-1]
    }
    return nil
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvdnismappingtable.Cvdnismappingentry {
        children[cvdnismappingtable.Cvdnismappingentry[i].GetSegmentPath()] = &cvdnismappingtable.Cvdnismappingentry[i]
    }
    return children
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetYangName() string { return "cvDnisMappingTable" }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) SetParent(parent types.Entity) { cvdnismappingtable.parent = parent }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetParent() types.Entity { return cvdnismappingtable.parent }

func (cvdnismappingtable *CISCOVOICEDNISMIB_Cvdnismappingtable) GetParentYangName() string { return "CISCO-VOICE-DNIS-MIB" }

// CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry
// Information about a single DNIS mapping. There is a
// unique DNIS map name. New DNIS mapping can be created
// using cvDnisMappingStatus.
// 
// The entry can be created with or without a file location 
// specified by cvDnisMappingUrl. The mapping file contains
// DNIS name and VXML page per line. For example, a 
// cvDnisMappingUrl could be tftp://someserver/dnismap.txt.
// This file is a text file and the content format is
//   dnis <dnisname> url <urlname>.
// An example of the contents of the file itself can be
//   dnis 18004251234 url http://www.b.com/p/vwelcome.vxml
//   dnis 18004253421 url http://www.c.com/j/vxmlintf.vxml
// If a mapping file location is specified, then new rows
// corresponding to this map name are created and populated
// in cvDnisNodeTable from the contents of the file. The
// rows corresponding to this map name in cvDnisNodeTable
// cannot be created or modified or deleted but can be
// read. 
// 
// If a mapping file location is not specified in
// cvDnisMappingUrl, then individual DNIS entries
// corresponding to this map name can be created, modified
// and deleted in cvDnisNodeTable. 
// 
// Deleting an entry deletes all the related entries in
// cvDnisNodeTable. 
type CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name which uniquely identifies a DNIS mapping.
    // . The type is string with length: 1..32.
    Cvdnismappingname interface{}

    // The url specifies a file location. The file contains individual DNIS
    // entries that belong to the DNIS map  name specified by cvDnisMappingName. 
    // Once a url is created and associated with a map name (the association is
    // complete when the row is made active(1)), it cannot be modified while
    // cvDnisMappingStatus is active. If a different url needs to be associated
    // with the current map name, the row status should be made notInService(2)
    // and this object has to be modified to associate a new url. When a new
    // association is made all the DNIS entries corresponding to the old
    // association will be deleted from the cvDnisNodeTable.  The url is read when
    // the row status is made active(1) or when the row status is active and the
    // object   cvDnisMappingRefresh is explicitly set to refresh(2).   If the url
    // is not accessible then a cvDnisMappingUrlInaccessible notification will be
    // generted. . The type is string.
    Cvdnismappingurl interface{}

    // Whenever there is a need to re-read the contents of the file specified by
    // cvDnisMappingUrl, this object can be set to refresh(2). This will cause the
    // contents of the file to be re-read and correspondingly update the
    // cvDnisNodeTable. After the completion of this operation, the value of this
    // object is reset to idle(1). The only operation allowed on this object is
    // setting it to refresh(2). This can only be done when the current value is
    // idle(1) and the rowstatus is active(1).  idle       - The refreshing
    // process is idle and the user              can modify this object to
    // refresh. refresh    - The refreshing process is currently busy and         
    // the user have to wait till the object              becomes idle to issue
    // new refresh. The type is Cvdnismappingrefresh.
    Cvdnismappingrefresh interface{}

    // ASCII text describing the error on last access of the url specified in
    // cvDnisMappingUrl.  If the url access does not succeed, then this object is
    // populated with an error message indicating the reason for failure. If the
    // url access succeeds, this object is set to null string. The type is string.
    Cvdnismappingurlaccesserror interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. When making the status active(1), if a valid
    // cvDnisMappingUrl is present the contents of the url is downloaded and
    // during that time cvDnisMappingRefresh is set to refresh(2). When
    // cvDnisMappingRefresh is set to refresh(2), only the destroy(6) operation is
    // allowed. The type is RowStatus.
    Cvdnismappingstatus interface{}
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetFilter() yfilter.YFilter { return cvdnismappingentry.YFilter }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) SetFilter(yf yfilter.YFilter) { cvdnismappingentry.YFilter = yf }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetGoName(yname string) string {
    if yname == "cvDnisMappingName" { return "Cvdnismappingname" }
    if yname == "cvDnisMappingUrl" { return "Cvdnismappingurl" }
    if yname == "cvDnisMappingRefresh" { return "Cvdnismappingrefresh" }
    if yname == "cvDnisMappingUrlAccessError" { return "Cvdnismappingurlaccesserror" }
    if yname == "cvDnisMappingStatus" { return "Cvdnismappingstatus" }
    return ""
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetSegmentPath() string {
    return "cvDnisMappingEntry" + "[cvDnisMappingName='" + fmt.Sprintf("%v", cvdnismappingentry.Cvdnismappingname) + "']"
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvDnisMappingName"] = cvdnismappingentry.Cvdnismappingname
    leafs["cvDnisMappingUrl"] = cvdnismappingentry.Cvdnismappingurl
    leafs["cvDnisMappingRefresh"] = cvdnismappingentry.Cvdnismappingrefresh
    leafs["cvDnisMappingUrlAccessError"] = cvdnismappingentry.Cvdnismappingurlaccesserror
    leafs["cvDnisMappingStatus"] = cvdnismappingentry.Cvdnismappingstatus
    return leafs
}

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetYangName() string { return "cvDnisMappingEntry" }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) SetParent(parent types.Entity) { cvdnismappingentry.parent = parent }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetParent() types.Entity { return cvdnismappingentry.parent }

func (cvdnismappingentry *CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry) GetParentYangName() string { return "cvDnisMappingTable" }

// CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh represents              becomes idle to issue new refresh.
type CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh string

const (
    CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh_idle CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh = "idle"

    CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh_refresh CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingrefresh = "refresh"
)

// CISCOVOICEDNISMIB_Cvdnisnodetable
// The table contains a DNIS name and a url. The url is a
// pointer to a VXML page for the DNIS name. 
type CISCOVOICEDNISMIB_Cvdnisnodetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry is a DNIS name and the location of the associated VXML page. New
    // DNIS entries can be created or the existing entries can be modified or
    // deleted only if the corresponding map name (defined in cvDnisMappingTable)
    // does not have any file name provided in the cvDnisMappingUrl object.   If a
    // file name is provided in cvDnisMappingUrl corresponding to this entry's map
    // name, then this row will have read permission only. The type is slice of
    // CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry.
    Cvdnisnodeentry []CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetFilter() yfilter.YFilter { return cvdnisnodetable.YFilter }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) SetFilter(yf yfilter.YFilter) { cvdnisnodetable.YFilter = yf }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetGoName(yname string) string {
    if yname == "cvDnisNodeEntry" { return "Cvdnisnodeentry" }
    return ""
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetSegmentPath() string {
    return "cvDnisNodeTable"
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cvDnisNodeEntry" {
        for _, c := range cvdnisnodetable.Cvdnisnodeentry {
            if cvdnisnodetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry{}
        cvdnisnodetable.Cvdnisnodeentry = append(cvdnisnodetable.Cvdnisnodeentry, child)
        return &cvdnisnodetable.Cvdnisnodeentry[len(cvdnisnodetable.Cvdnisnodeentry)-1]
    }
    return nil
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cvdnisnodetable.Cvdnisnodeentry {
        children[cvdnisnodetable.Cvdnisnodeentry[i].GetSegmentPath()] = &cvdnisnodetable.Cvdnisnodeentry[i]
    }
    return children
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetBundleName() string { return "cisco_ios_xe" }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetYangName() string { return "cvDnisNodeTable" }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) SetParent(parent types.Entity) { cvdnisnodetable.parent = parent }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetParent() types.Entity { return cvdnisnodetable.parent }

func (cvdnisnodetable *CISCOVOICEDNISMIB_Cvdnisnodetable) GetParentYangName() string { return "CISCO-VOICE-DNIS-MIB" }

// CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry
// Each entry is a DNIS name and the location of the
// associated VXML page. New DNIS entries can be created or
// the existing entries can be modified or deleted only if
// the corresponding map name (defined in
// cvDnisMappingTable) does not have any file name provided
// in the cvDnisMappingUrl object. 
// 
// If a file name is provided in cvDnisMappingUrl
// corresponding to this entry's map name, then this row
// will have read permission only.
type CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // cisco_voice_dnis_mib.CISCOVOICEDNISMIB_Cvdnismappingtable_Cvdnismappingentry_Cvdnismappingname
    Cvdnismappingname interface{}

    // This attribute is a key. The individual DNIS name. It is unique within a
    // DNIS mapping. The type is string.
    Cvdnisnumber interface{}

    // The url specifies a VXML page. This page contains voice XML links to play
    // audio data.  This url which is a VXML page is not read immediately when the
    // row is made active(1), but only when a call that requires the use of this
    // DNIS comes through. The type is string.
    Cvdnisnodeurl interface{}

    // This object specifies whether the object in a particular row is modifiable.
    // The object is set to true(1) if the corresponding map name (defined in
    // cvDnisMappingTable) does not have any file name provided in the
    // cvDnisMappingUrl object. Otherwise this object is set to false(2) and the
    // row becomes read only.  . The type is bool.
    Cvdnisnodemodifiable interface{}

    // This object is used to create a new row or modify or delete an existing row
    // in this table. The objects in a row can be modified or deleted while the
    // row status is active(1) and cvDnisNodeModifiable is true(1). The row status
    // cannot be set to notInService(2) or createAndWait(5). . The type is
    // RowStatus.
    Cvdnisnodestatus interface{}
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetFilter() yfilter.YFilter { return cvdnisnodeentry.YFilter }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) SetFilter(yf yfilter.YFilter) { cvdnisnodeentry.YFilter = yf }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetGoName(yname string) string {
    if yname == "cvDnisMappingName" { return "Cvdnismappingname" }
    if yname == "cvDnisNumber" { return "Cvdnisnumber" }
    if yname == "cvDnisNodeUrl" { return "Cvdnisnodeurl" }
    if yname == "cvDnisNodeModifiable" { return "Cvdnisnodemodifiable" }
    if yname == "cvDnisNodeStatus" { return "Cvdnisnodestatus" }
    return ""
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetSegmentPath() string {
    return "cvDnisNodeEntry" + "[cvDnisMappingName='" + fmt.Sprintf("%v", cvdnisnodeentry.Cvdnismappingname) + "']" + "[cvDnisNumber='" + fmt.Sprintf("%v", cvdnisnodeentry.Cvdnisnumber) + "']"
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cvDnisMappingName"] = cvdnisnodeentry.Cvdnismappingname
    leafs["cvDnisNumber"] = cvdnisnodeentry.Cvdnisnumber
    leafs["cvDnisNodeUrl"] = cvdnisnodeentry.Cvdnisnodeurl
    leafs["cvDnisNodeModifiable"] = cvdnisnodeentry.Cvdnisnodemodifiable
    leafs["cvDnisNodeStatus"] = cvdnisnodeentry.Cvdnisnodestatus
    return leafs
}

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetBundleName() string { return "cisco_ios_xe" }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetYangName() string { return "cvDnisNodeEntry" }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) SetParent(parent types.Entity) { cvdnisnodeentry.parent = parent }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetParent() types.Entity { return cvdnisnodeentry.parent }

func (cvdnisnodeentry *CISCOVOICEDNISMIB_Cvdnisnodetable_Cvdnisnodeentry) GetParentYangName() string { return "cvDnisNodeTable" }

