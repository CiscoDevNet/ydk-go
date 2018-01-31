// The MIB module for defining expressions of MIB objects
// for network management purposes.
// 
// This MIB is an early snapshot of work done by the IETF's
// Distributed Management working group.  After this snapshot
// was taken, the MIB was modified, had new OIDs assigned,
// and then published as RFC 2982.
package expression_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package expression_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:EXPRESSION-MIB EXPRESSION-MIB}", reflect.TypeOf(EXPRESSIONMIB{}))
    ydk.RegisterEntity("EXPRESSION-MIB:EXPRESSION-MIB", reflect.TypeOf(EXPRESSIONMIB{}))
}

// EXPRESSIONMIB
type EXPRESSIONMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Expresource EXPRESSIONMIB_Expresource

    
    Expnames EXPRESSIONMIB_Expnames

    // A table of expression names, for creating and deleting expressions.
    Expnametable EXPRESSIONMIB_Expnametable

    // A table of expression definitions.
    Expexpressiontable EXPRESSIONMIB_Expexpressiontable

    // A table of object definitions for each expExpression.  Wildcarding instance
    // IDs:  It is legal to omit all or part of the instance portion for some or
    // all of the objects in an expression. (See the DESCRIPTION of expObjectID
    // for details.  However, note that if more than one object in the same
    // expression is wildcarded in this way, they all must be objects where that
    // portion of the instance is the same.  In other words, all objects may be in
    // the same SEQUENCE or in different SEQUENCEs but with the same semantic
    // index value (e.g., a value of ifIndex) for the wildcarded portion.
    Expobjecttable EXPRESSIONMIB_Expobjecttable

    // A table of values from evaluated expressions.
    Expvaluetable EXPRESSIONMIB_Expvaluetable
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetFilter() yfilter.YFilter { return eXPRESSIONMIB.YFilter }

func (eXPRESSIONMIB *EXPRESSIONMIB) SetFilter(yf yfilter.YFilter) { eXPRESSIONMIB.YFilter = yf }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetGoName(yname string) string {
    if yname == "expResource" { return "Expresource" }
    if yname == "expNames" { return "Expnames" }
    if yname == "expNameTable" { return "Expnametable" }
    if yname == "expExpressionTable" { return "Expexpressiontable" }
    if yname == "expObjectTable" { return "Expobjecttable" }
    if yname == "expValueTable" { return "Expvaluetable" }
    return ""
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetSegmentPath() string {
    return "EXPRESSION-MIB:EXPRESSION-MIB"
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expResource" {
        return &eXPRESSIONMIB.Expresource
    }
    if childYangName == "expNames" {
        return &eXPRESSIONMIB.Expnames
    }
    if childYangName == "expNameTable" {
        return &eXPRESSIONMIB.Expnametable
    }
    if childYangName == "expExpressionTable" {
        return &eXPRESSIONMIB.Expexpressiontable
    }
    if childYangName == "expObjectTable" {
        return &eXPRESSIONMIB.Expobjecttable
    }
    if childYangName == "expValueTable" {
        return &eXPRESSIONMIB.Expvaluetable
    }
    return nil
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["expResource"] = &eXPRESSIONMIB.Expresource
    children["expNames"] = &eXPRESSIONMIB.Expnames
    children["expNameTable"] = &eXPRESSIONMIB.Expnametable
    children["expExpressionTable"] = &eXPRESSIONMIB.Expexpressiontable
    children["expObjectTable"] = &eXPRESSIONMIB.Expobjecttable
    children["expValueTable"] = &eXPRESSIONMIB.Expvaluetable
    return children
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetBundleName() string { return "cisco_ios_xe" }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetYangName() string { return "EXPRESSION-MIB" }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eXPRESSIONMIB *EXPRESSIONMIB) SetParent(parent types.Entity) { eXPRESSIONMIB.parent = parent }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetParent() types.Entity { return eXPRESSIONMIB.parent }

func (eXPRESSIONMIB *EXPRESSIONMIB) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expresource
type EXPRESSIONMIB_Expresource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The minimum expExpressionDeltaInterval this system will accept.  A system
    // may use the larger values of this minimum to lessen the impact of
    // constantly computing deltas.  The value -1 indicates this system will not
    // accept deltaValue as a value for expObjectSampleType.  Unless explicitly
    // resource limited, a system's value for this object should be 1.  Changing
    // this value will not invalidate an existing setting of expObjectSampleType.
    // The type is interface{} with range: -1..None | 1..600. Units are seconds.
    Expresourcedeltaminimum interface{}

    // The maximum number of dynamic instance entries this system will support for
    // wildcarded delta objects in expressions. These are the entries that
    // maintain state, one for each instance of each deltaValue object for each
    // value of an expression.  A value of 0 indicates no preset limit, that is,
    // the limit is dynamic based on system operation and resources.  Unless
    // explicitly resource limited, a system's value for this object should be 0. 
    // Changing this value will not eliminate or inhibit existing delta wildcard
    // instance objects but will prevent the creation of more such objects. The
    // type is interface{} with range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstancemaximum interface{}

    // The number of currently active instance entries as defined for
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstances interface{}

    // The highest value of expResourceDeltaWildcardInstances that has occurred
    // since initialization of the management system. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstanceshigh interface{}

    // The number of times this system could not evaluate an expression because
    // that would have created a value instance in excess of
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Expresourcedeltawildcardinstanceresourcelacks interface{}
}

func (expresource *EXPRESSIONMIB_Expresource) GetFilter() yfilter.YFilter { return expresource.YFilter }

func (expresource *EXPRESSIONMIB_Expresource) SetFilter(yf yfilter.YFilter) { expresource.YFilter = yf }

func (expresource *EXPRESSIONMIB_Expresource) GetGoName(yname string) string {
    if yname == "expResourceDeltaMinimum" { return "Expresourcedeltaminimum" }
    if yname == "expResourceDeltaWildcardInstanceMaximum" { return "Expresourcedeltawildcardinstancemaximum" }
    if yname == "expResourceDeltaWildcardInstances" { return "Expresourcedeltawildcardinstances" }
    if yname == "expResourceDeltaWildcardInstancesHigh" { return "Expresourcedeltawildcardinstanceshigh" }
    if yname == "expResourceDeltaWildcardInstanceResourceLacks" { return "Expresourcedeltawildcardinstanceresourcelacks" }
    return ""
}

func (expresource *EXPRESSIONMIB_Expresource) GetSegmentPath() string {
    return "expResource"
}

func (expresource *EXPRESSIONMIB_Expresource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expresource *EXPRESSIONMIB_Expresource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expresource *EXPRESSIONMIB_Expresource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expResourceDeltaMinimum"] = expresource.Expresourcedeltaminimum
    leafs["expResourceDeltaWildcardInstanceMaximum"] = expresource.Expresourcedeltawildcardinstancemaximum
    leafs["expResourceDeltaWildcardInstances"] = expresource.Expresourcedeltawildcardinstances
    leafs["expResourceDeltaWildcardInstancesHigh"] = expresource.Expresourcedeltawildcardinstanceshigh
    leafs["expResourceDeltaWildcardInstanceResourceLacks"] = expresource.Expresourcedeltawildcardinstanceresourcelacks
    return leafs
}

func (expresource *EXPRESSIONMIB_Expresource) GetBundleName() string { return "cisco_ios_xe" }

func (expresource *EXPRESSIONMIB_Expresource) GetYangName() string { return "expResource" }

func (expresource *EXPRESSIONMIB_Expresource) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expresource *EXPRESSIONMIB_Expresource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expresource *EXPRESSIONMIB_Expresource) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expresource *EXPRESSIONMIB_Expresource) SetParent(parent types.Entity) { expresource.parent = parent }

func (expresource *EXPRESSIONMIB_Expresource) GetParent() types.Entity { return expresource.parent }

func (expresource *EXPRESSIONMIB_Expresource) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expnames
type EXPRESSIONMIB_Expnames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime the last time an expression was created or deleted
    // or had its name changed using expExpressionName. The type is interface{}
    // with range: 0..4294967295.
    Expnamelastchange interface{}

    // The highest value of ExpressionIndex ever assigned on this system. 
    // Preferrably this value is preserved across system reboots.  A managed
    // system that is unable to store expressions across reboots need not preserve
    // this value across reboots.  If all expression-creating applications
    // cooperate, they may use this to avoid reusing an ExpressionIndex.  To do
    // so, attempt creation of a new entry with this value + 1 as the value of
    // expExpressionIndex.  Although reusing ExpressionIndexes could lead to an
    // application receiving a misunderstood value, it is a matter of local
    // management policy whether to reuse them. The type is interface{} with
    // range: 0..4294967295.
    Expnamehighestindex interface{}
}

func (expnames *EXPRESSIONMIB_Expnames) GetFilter() yfilter.YFilter { return expnames.YFilter }

func (expnames *EXPRESSIONMIB_Expnames) SetFilter(yf yfilter.YFilter) { expnames.YFilter = yf }

func (expnames *EXPRESSIONMIB_Expnames) GetGoName(yname string) string {
    if yname == "expNameLastChange" { return "Expnamelastchange" }
    if yname == "expNameHighestIndex" { return "Expnamehighestindex" }
    return ""
}

func (expnames *EXPRESSIONMIB_Expnames) GetSegmentPath() string {
    return "expNames"
}

func (expnames *EXPRESSIONMIB_Expnames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expnames *EXPRESSIONMIB_Expnames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expnames *EXPRESSIONMIB_Expnames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expNameLastChange"] = expnames.Expnamelastchange
    leafs["expNameHighestIndex"] = expnames.Expnamehighestindex
    return leafs
}

func (expnames *EXPRESSIONMIB_Expnames) GetBundleName() string { return "cisco_ios_xe" }

func (expnames *EXPRESSIONMIB_Expnames) GetYangName() string { return "expNames" }

func (expnames *EXPRESSIONMIB_Expnames) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expnames *EXPRESSIONMIB_Expnames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expnames *EXPRESSIONMIB_Expnames) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expnames *EXPRESSIONMIB_Expnames) SetParent(parent types.Entity) { expnames.parent = parent }

func (expnames *EXPRESSIONMIB_Expnames) GetParent() types.Entity { return expnames.parent }

func (expnames *EXPRESSIONMIB_Expnames) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expnametable
// A table of expression names, for creating and deleting
// expressions.
type EXPRESSIONMIB_Expnametable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single expression.  New expressions can be created
    // using expNameStatus.  To create an expression first create the named entry
    // in this table.  Then use expExpressionIndex to populate expExpressionTable
    // and expObjectTable.  For expression evaluation to succeed all related
    // entries in expNameTable, expExpressionTable, and expObjectTable must be
    // 'active'.  If these conditions are not met the corresponding values in
    // expValue simply are not instantiated.  Deleting an entry deletes all
    // related entries in expExpressionTable and expObjectTable.  Because of the
    // relationships among the multiple tables for an expression (expNameTable,
    // expExpressionTable, expObjectTable, and expValueTable) and the SNMP rules
    // for independence in setting object values, it is necessary to do final
    // error checking when an expression is evaluated, that is, when one of its
    // instances in expValueTable is read.  Earlier checking need not be done and
    // an implementation may not impose any ordering on the creation of objects
    // related to an expression other than to require values for expName and
    // expExpressionIndex before any other related objects can be created.  To
    // maintain security of MIB information, when creating a new row in this
    // table, the managed system must record the security credentials of the
    // requester.  If the subsequent expression includes objects with
    // expObjectSampleType 'deltaValue' the evaluation of that expression takes
    // place under the security credentials of the creator of its expNameEntry.
    // The type is slice of EXPRESSIONMIB_Expnametable_Expnameentry.
    Expnameentry []EXPRESSIONMIB_Expnametable_Expnameentry
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetFilter() yfilter.YFilter { return expnametable.YFilter }

func (expnametable *EXPRESSIONMIB_Expnametable) SetFilter(yf yfilter.YFilter) { expnametable.YFilter = yf }

func (expnametable *EXPRESSIONMIB_Expnametable) GetGoName(yname string) string {
    if yname == "expNameEntry" { return "Expnameentry" }
    return ""
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetSegmentPath() string {
    return "expNameTable"
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expNameEntry" {
        for _, c := range expnametable.Expnameentry {
            if expnametable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EXPRESSIONMIB_Expnametable_Expnameentry{}
        expnametable.Expnameentry = append(expnametable.Expnameentry, child)
        return &expnametable.Expnameentry[len(expnametable.Expnameentry)-1]
    }
    return nil
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expnametable.Expnameentry {
        children[expnametable.Expnameentry[i].GetSegmentPath()] = &expnametable.Expnameentry[i]
    }
    return children
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expnametable *EXPRESSIONMIB_Expnametable) GetBundleName() string { return "cisco_ios_xe" }

func (expnametable *EXPRESSIONMIB_Expnametable) GetYangName() string { return "expNameTable" }

func (expnametable *EXPRESSIONMIB_Expnametable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expnametable *EXPRESSIONMIB_Expnametable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expnametable *EXPRESSIONMIB_Expnametable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expnametable *EXPRESSIONMIB_Expnametable) SetParent(parent types.Entity) { expnametable.parent = parent }

func (expnametable *EXPRESSIONMIB_Expnametable) GetParent() types.Entity { return expnametable.parent }

func (expnametable *EXPRESSIONMIB_Expnametable) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expnametable_Expnameentry
// Information about a single expression.  New expressions
// can be created using expNameStatus.
// 
// To create an expression first create the named entry in this
// table.  Then use expExpressionIndex to populate
// expExpressionTable and expObjectTable.  For expression
// evaluation to succeed all related entries in expNameTable,
// expExpressionTable, and expObjectTable must be 'active'.  If
// these conditions are not met the corresponding values in
// expValue simply are not instantiated.
// 
// Deleting an entry deletes all related entries in
// expExpressionTable and expObjectTable.
// 
// Because of the relationships among the multiple tables
// for an expression (expNameTable, expExpressionTable,
// expObjectTable, and expValueTable) and the SNMP rules
// for independence in setting object values, it is
// necessary to do final error checking when an expression
// is evaluated, that is, when one of its instances in
// expValueTable is read.  Earlier checking need not be
// done and an implementation may not impose any ordering
// on the creation of objects related to an expression other
// than to require values for expName and expExpressionIndex
// before any other related objects can be created.
// 
// To maintain security of MIB information, when creating a new
// row in this table, the managed system must record the
// security credentials of the requester.  If the subsequent
// expression includes objects with expObjectSampleType
// 'deltaValue' the evaluation of that expression takes place
// under the security credentials of the creator of its
// expNameEntry.
type EXPRESSIONMIB_Expnametable_Expnameentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the expression.  Choosing names with
    // useful lexical ordering supports using GetNext or GetBulk to retrieve a
    // useful subset of the table. The type is string with length: 1..64.
    Expname interface{}

    // The numeric identification of the expression.  Applications may select this
    // number in ascending numerical order by using expNameHighestIndex as a hint
    // or may use any other acceptable, unused number.  Once set this value may
    // not be set to a different value. The type is interface{} with range:
    // 1..4294967295.
    Expexpressionindex interface{}

    // The control that allows creation/deletion of entries. The type is
    // RowStatus.
    Expnamestatus interface{}
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetFilter() yfilter.YFilter { return expnameentry.YFilter }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) SetFilter(yf yfilter.YFilter) { expnameentry.YFilter = yf }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetGoName(yname string) string {
    if yname == "expName" { return "Expname" }
    if yname == "expExpressionIndex" { return "Expexpressionindex" }
    if yname == "expNameStatus" { return "Expnamestatus" }
    return ""
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetSegmentPath() string {
    return "expNameEntry" + "[expName='" + fmt.Sprintf("%v", expnameentry.Expname) + "']"
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expName"] = expnameentry.Expname
    leafs["expExpressionIndex"] = expnameentry.Expexpressionindex
    leafs["expNameStatus"] = expnameentry.Expnamestatus
    return leafs
}

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetBundleName() string { return "cisco_ios_xe" }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetYangName() string { return "expNameEntry" }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) SetParent(parent types.Entity) { expnameentry.parent = parent }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetParent() types.Entity { return expnameentry.parent }

func (expnameentry *EXPRESSIONMIB_Expnametable_Expnameentry) GetParentYangName() string { return "expNameTable" }

// EXPRESSIONMIB_Expexpressiontable
// A table of expression definitions.
type EXPRESSIONMIB_Expexpressiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single expression.  An entry appears in this table when
    // an entry is created in expNameTable. Deleting that expNameTable entry
    // automatically deletes this entry and its associated expObjectTable entries.
    // Values of read-write objects in this table may be changed at any time. The
    // type is slice of EXPRESSIONMIB_Expexpressiontable_Expexpressionentry.
    Expexpressionentry []EXPRESSIONMIB_Expexpressiontable_Expexpressionentry
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetFilter() yfilter.YFilter { return expexpressiontable.YFilter }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) SetFilter(yf yfilter.YFilter) { expexpressiontable.YFilter = yf }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetGoName(yname string) string {
    if yname == "expExpressionEntry" { return "Expexpressionentry" }
    return ""
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetSegmentPath() string {
    return "expExpressionTable"
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expExpressionEntry" {
        for _, c := range expexpressiontable.Expexpressionentry {
            if expexpressiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EXPRESSIONMIB_Expexpressiontable_Expexpressionentry{}
        expexpressiontable.Expexpressionentry = append(expexpressiontable.Expexpressionentry, child)
        return &expexpressiontable.Expexpressionentry[len(expexpressiontable.Expexpressionentry)-1]
    }
    return nil
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expexpressiontable.Expexpressionentry {
        children[expexpressiontable.Expexpressionentry[i].GetSegmentPath()] = &expexpressiontable.Expexpressionentry[i]
    }
    return children
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetBundleName() string { return "cisco_ios_xe" }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetYangName() string { return "expExpressionTable" }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) SetParent(parent types.Entity) { expexpressiontable.parent = parent }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetParent() types.Entity { return expexpressiontable.parent }

func (expexpressiontable *EXPRESSIONMIB_Expexpressiontable) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expexpressiontable_Expexpressionentry
// Information about a single expression.  An entry appears
// in this table when an entry is created in expNameTable.
// Deleting that expNameTable entry automatically deletes
// this entry and its associated expObjectTable entries.
// 
// Values of read-write objects in this table may be changed
// at any time.
type EXPRESSIONMIB_Expexpressiontable_Expexpressionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_Expnametable_Expnameentry_Expexpressionindex
    Expexpressionindex interface{}

    // The unique name of the expression, the same as expName.  Use this object to
    // change the expression's name without changing its expExpressionIndex. The
    // type is string with length: 1..64.
    Expexpressionname interface{}

    // The expression to be evaluated.  This object is the same as a DisplayString
    // (RFC 1903) except for its maximum length.  Except for the variable names
    // the expression is in ANSI C syntax.  Only the subset of ANSI C operators
    // and functions listed here is allowed.  Variables are expressed as a dollar
    // sign ('$') and an integer that corresponds to an expObjectIndex.  An
    // example of a valid expression is:          ($1-$5)*100  Expressions may not
    // be recursive, that is although an expression may use the results of another
    // expression, it may not contain any variable that is directly or indirectly
    // a result of its own evaluation.  The only allowed operators are:          (
    // )         - (unary)         + - * / %         & | ^ << >> ~         ! && ||
    // == != > >= < <=  Note the parentheses are included for parenthesizing the
    // expression, not for casting data types.  The only constant types defined
    // are:          int (32-bit signed)         long (64-bit signed)        
    // unsigned int         unsigned long         hexadecimal         character   
    // string         oid  The default type for a positive integer is int unless
    // it is too large in which case it is long.  All but oid are as defined for
    // ANSI C.  Note that a hexadecimal constant may end up as a scalar or an
    // array of 8-bit integers.  A string constant is enclosed in double quotes
    // and may contain back-slashed individual characters as in ANSI C.  An oid
    // constant comprises 32-bit, unsigned integers and at least one period, for
    // example:          0.         .0         1.3.6.1  Integer-typed objects are
    // treated as 32- or 64-bit, signed or unsigned integers, as appropriate.  The
    // results of mixing them are as for ANSI C, including the type of the result.
    // Note that a 32-bit value is thus promoted to 64 bits only in an operation
    // with a 64-bit value.  There is no provision for larger values to handle
    // overflow.  Relative to SNMP data types, a resulting value becomes unsigned
    // when calculating it uses any unsigned value, including a counter.  To force
    // the final value to be of data type counter the expression must explicitly
    // use the counter32() or counter64() function (defined below).  OCTET STRINGS
    // and OBJECT IDENTIFIERs are treated as 1-based arrays of unsigned 8-bit
    // integers and unsigned 32-bit integers, respectively.  IpAddresses are
    // treated as 32-bit, unsigned integers in network byte order, that is, the
    // hex version of 255.0.0.0 is 0xff000000.  Conditional expressions result in
    // a 32-bit, unsigned integer of value 0 for false or 1 for true. When an
    // arbitrary value is used as a boolean 0 is false and non-zero is true. 
    // Rules for the resulting data type from an operation, based on the operator:
    // For << and >> the result is the same as the left hand operand.  For &&, ||,
    // ==, !=, <, <=, >, and >= the result is always Unsigned32.  For unary - the
    // result is always Integer32.  For +, -, *, /, %, &, |, and ^ the result is
    // promoted according to the following rules, in order from most to least
    // preferred:          If left hand and right hand operands are the same      
    // type, use that.          If either side is Counter64, use that.          If
    // either side is IpAddress, use that.          If either side is TimeTicks,
    // use that.          If either side is Counter32, use that.         
    // Otherwise use Unsigned32.  The following rules say what operators apply
    // with what data types.  Any combination not explicitly defined does not
    // work.  For all operators any of the following can be the left hand or right
    // hand operand: Integer32, Counter32, Unsigned32, Counter64.  The operators
    // +, -, *, /, %, <, <=, >, and >= also work with TimeTicks.  The operators &,
    // |, and ^ also work with IpAddress.  The operators << and >> also work with
    // IpAddress but only as the left hand operand.  The + operator performs a
    // concatenation of two OCTET STRINGs or two OBJECT IDENTIFIERs.  The
    // operators &, | perform bitwise operations on OCTET STRINGs.  If the OCTET
    // STRING happens to be a DisplayString the results may be meaningless, but
    // the agent system does not check this as some such systems do not have this
    // information.  The operators << and >> perform bitwise operations on OCTET
    // STRINGs appearing as the left hand operand.  The only functions defined
    // are:          counter32         counter64         arraySection        
    // stringBegins         stringEnds         stringContains         oidBegins   
    // oidEnds         oidContains         sum         exists  The following
    // function definitions indicate their by naming the data type of the
    // parameter in the parameter's position in the parameter list.  The parameter
    // must be of the type indicated and generally may be a constant, a MIB
    // object, a function, or an expression.  counter32(integer) - wrapped around
    // an integer value counter32 forces Counter32 as a data type. 
    // counter64(integer) - similar to counter32 except that the resulting data
    // type is 'counter64'.  arraySection(array, integer, integer) - selects a
    // piece of an array (i.e. part of an OCTET STRING or OBJECT IDENTIFIER).  The
    // integer arguments are in the range 0 to 4,294,967,295.  The first is an
    // initial array index (1-based) and the second is an ending array index.  A
    // value of 0 indicates first or last element, respectively.  If the first
    // element is larger than the array length the result is 0 length.  If the
    // second integer is less than or equal to the first, the result is 0 length. 
    // If the second is larger than the array length it indicates last element. 
    // stringBegins/Ends/Contains(octetString, octetString) - looks for the second
    // string (which can be a string constant) in the first and returns the
    // 1-based index where the match began.  A return value of 0 indicates no
    // match (i.e. boolean false).  oidBegins/Ends/Contains(oid, oid) - looks for
    // the second OID (which can be an OID constant) in the first and returns the
    // the 1-based index where the match began.  A return value of 0 indicates no
    // match (i.e. boolean false).  sum(integerObject*) - sums all availiable
    // values of the wildcarded integer object, resulting in an integer scalar.
    // Must be used with caution as it wraps on overflow with no notification. 
    // exists(anyTypeObject) - verifies the object instance exists. A return value
    // of 0 indicates NoSuchInstance (i.e. boolean false). The type is string with
    // length: 1..1024.
    Expexpression interface{}

    // The type of the expression value.  One and only one of the value objects in
    // expValueTable will be instantiated to match this type.  If the result of
    // the expression can not be made into this type, an invalidOperandType error
    // will occur. The type is Expexpressionvaluetype.
    Expexpressionvaluetype interface{}

    // A comment to explain the use or meaning of the expression. The type is
    // string.
    Expexpressioncomment interface{}

    // Sampling interval for objects in this expression with expObjectSampleType
    // 'deltaValue'.  This object is not instantiated if not applicable.  A value
    // of 0 indicates no automated sampling.  In this case the delta is the
    // difference from the last time the expression was evaluated.  Note that this
    // is subject to unpredictable delta times in the face of retries or multiple
    // managers.  A value greater than zero is the number of seconds between
    // automated samples.  Until the delta interval has expired once the delta for
    // the object is effectively not instantiated and evaluating the expression
    // has results as if the object itself were not instantiated.  Note that delta
    // values potentially consume large amounts of system CPU and memory.  Delta
    // state and processing must continue constantly even if the expression is not
    // being used.  That is, the expression is being evaluated every delta
    // interval, even if no application is reading those values.  For wildcarded
    // objects this can be substantial overhead.  Note that delta intervals,
    // external expression value sampling intervals and delta intervals for
    // expressions within other expressions can have unusual interactions as they
    // are impossible to synchronize accurately.  In general one interval embedded
    // below another must be enough shorter that the higher sample sees relatively
    // smooth, predictable behavior. The type is interface{} with range: 0..86400.
    // Units are seconds.
    Expexpressiondeltainterval interface{}

    // An object prefix to assist an application in determining the instance
    // indexing to use in expValueTable, relieving the application of the need to
    // scan the expObjectTable to determine such a prefix.  See expObjectTable for
    // information on wildcarded objects.  If the expValueInstance portion of the
    // value OID may be treated as a scalar (that is, normally, 0) the value of
    // expExpressionPrefix is zero length, that is, no OID at all.  Otherwise
    // expExpressionPrefix is the value of any wildcarded instance of expObjectID
    // for the expression.  This is sufficient as the remainder, that is, the
    // instance fragment relevant to instancing the values must be the same for
    // all wildcarded objects in the expression. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expexpressionprefix interface{}

    // The number of errors encountered while evaluating this expression.  Note
    // that an object in the expression not being accessible is not considered an
    // error.  It is a legitimate condition that causes the corresponding
    // expression value not to be instantiated. The type is interface{} with
    // range: 0..4294967295.
    Expexpressionerrors interface{}

    // The value of sysUpTime the last time an error caused a failure to evaluate
    // this expression.  This object is not instantiated if there have been no
    // errors. The type is interface{} with range: 0..4294967295.
    Expexpressionerrortime interface{}

    // The 1-based character index into expExpression for where the error
    // occurred.  The value zero indicates irrelevance.  This object is not
    // instantiated if there have been no errors. The type is interface{} with
    // range: -2147483648..2147483647.
    Expexpressionerrorindex interface{}

    // The error that occurred.  In the following explanations the expected timing
    // of the error is in parentheses.  'S' means the error occurs on a Set
    // request.  'E' means the error occurs on the attempt to evaluate the
    // expression either due to Get from expValueTable or in ongoing delta
    // processing.  invalidSyntax           the value sent for expExpression      
    // is not valid Expression MIB                         expression syntax (S)
    // undefinedObjectIndex    an object reference ($n) in                        
    // expExpression does not have a                         matching instance in 
    // expObjectTable (E) unrecognizedOperator    the value sent for expExpression
    // held an unrecognized operator (S) unrecognizedFunction    the value sent
    // for expExpression                         held an unrecognized function    
    // name (S) invalidOperandType      an operand in expExpression is not        
    // the right type for the associated                         operator or
    // result (SE) unmatchedParenthesis    the value sent for expExpression       
    // is not correctly parenthesized (S) tooManyWildcardValues   evaluating the
    // expression exceeded                         the limit set by
    // expResourceDelta                         WildcardInstanceMaximum (E)
    // recursion               through some chain of embedded                     
    // expressions the expression invokes                         itself (E)
    // deltaTooShort           the delta for the next evaluation                  
    // passed before the system could                         evaluate the present
    // sample (E) resourceUnavailable     some resource, typically dynamic        
    // memory, was unavailable (SE) divideByZero            an attempt to divide
    // by zero                         occurred (E)  For the errors that occur
    // when the attempt is made to set expExpression Set request fails with the
    // SNMP error code 'wrongValue'. Such failures refer to the most recent
    // failure to Set expExpression, not to the present value of expExpression
    // which must be either unset or syntactically correct.  Errors that occur
    // during evalutaion for a Get* operation return the SNMP error code 'genErr'
    // except for 'tooManyWildcardValues' and 'resourceUnavailable' which return
    // the SNMP error code 'resourceUnavailable'.  This object is not instantiated
    // if there have been no errors. The type is Expexpressionerror.
    Expexpressionerror interface{}

    // The expValueInstance being evaluated when the error occurred.  A
    // zero-length indicates irrelevance.  This object is not instantiated if
    // there have been no errors. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expexpressioninstance interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..255.
    Expexpressionowner interface{}
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetFilter() yfilter.YFilter { return expexpressionentry.YFilter }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) SetFilter(yf yfilter.YFilter) { expexpressionentry.YFilter = yf }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetGoName(yname string) string {
    if yname == "expExpressionIndex" { return "Expexpressionindex" }
    if yname == "expExpressionName" { return "Expexpressionname" }
    if yname == "expExpression" { return "Expexpression" }
    if yname == "expExpressionValueType" { return "Expexpressionvaluetype" }
    if yname == "expExpressionComment" { return "Expexpressioncomment" }
    if yname == "expExpressionDeltaInterval" { return "Expexpressiondeltainterval" }
    if yname == "expExpressionPrefix" { return "Expexpressionprefix" }
    if yname == "expExpressionErrors" { return "Expexpressionerrors" }
    if yname == "expExpressionErrorTime" { return "Expexpressionerrortime" }
    if yname == "expExpressionErrorIndex" { return "Expexpressionerrorindex" }
    if yname == "expExpressionError" { return "Expexpressionerror" }
    if yname == "expExpressionInstance" { return "Expexpressioninstance" }
    if yname == "expExpressionOwner" { return "Expexpressionowner" }
    return ""
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetSegmentPath() string {
    return "expExpressionEntry" + "[expExpressionIndex='" + fmt.Sprintf("%v", expexpressionentry.Expexpressionindex) + "']"
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionIndex"] = expexpressionentry.Expexpressionindex
    leafs["expExpressionName"] = expexpressionentry.Expexpressionname
    leafs["expExpression"] = expexpressionentry.Expexpression
    leafs["expExpressionValueType"] = expexpressionentry.Expexpressionvaluetype
    leafs["expExpressionComment"] = expexpressionentry.Expexpressioncomment
    leafs["expExpressionDeltaInterval"] = expexpressionentry.Expexpressiondeltainterval
    leafs["expExpressionPrefix"] = expexpressionentry.Expexpressionprefix
    leafs["expExpressionErrors"] = expexpressionentry.Expexpressionerrors
    leafs["expExpressionErrorTime"] = expexpressionentry.Expexpressionerrortime
    leafs["expExpressionErrorIndex"] = expexpressionentry.Expexpressionerrorindex
    leafs["expExpressionError"] = expexpressionentry.Expexpressionerror
    leafs["expExpressionInstance"] = expexpressionentry.Expexpressioninstance
    leafs["expExpressionOwner"] = expexpressionentry.Expexpressionowner
    return leafs
}

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetBundleName() string { return "cisco_ios_xe" }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetYangName() string { return "expExpressionEntry" }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) SetParent(parent types.Entity) { expexpressionentry.parent = parent }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetParent() types.Entity { return expexpressionentry.parent }

func (expexpressionentry *EXPRESSIONMIB_Expexpressiontable_Expexpressionentry) GetParentYangName() string { return "expExpressionTable" }

// EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror represents errors.
type EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror string

const (
    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_invalidSyntax EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "invalidSyntax"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_undefinedObjectIndex EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "undefinedObjectIndex"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_unrecognizedOperator EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "unrecognizedOperator"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_unrecognizedFunction EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "unrecognizedFunction"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_invalidOperandType EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "invalidOperandType"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_unmatchedParenthesis EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "unmatchedParenthesis"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_tooManyWildcardValues EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "tooManyWildcardValues"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_recursion EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "recursion"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_deltaTooShort EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "deltaTooShort"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_resourceUnavailable EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "resourceUnavailable"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror_divideByZero EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionerror = "divideByZero"
)

// EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype represents type, an invalidOperandType error will occur.
type EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype string

const (
    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_counter32 EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "counter32"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_unsignedOrGauge32 EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "unsignedOrGauge32"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_timeTicks EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "timeTicks"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_integer32 EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "integer32"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_ipAddress EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "ipAddress"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_octetString EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "octetString"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_objectId EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "objectId"

    EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype_counter64 EXPRESSIONMIB_Expexpressiontable_Expexpressionentry_Expexpressionvaluetype = "counter64"
)

// EXPRESSIONMIB_Expobjecttable
// A table of object definitions for each expExpression.
// 
// Wildcarding instance IDs:
// 
// It is legal to omit all or part of the instance portion for
// some or all of the objects in an expression. (See the
// DESCRIPTION of expObjectID for details.  However, note that
// if more than one object in the same expression is wildcarded
// in this way, they all must be objects where that portion of
// the instance is the same.  In other words, all objects may
// be in the same SEQUENCE or in different SEQUENCEs but with
// the same semantic index value (e.g., a value of ifIndex)
// for the wildcarded portion.
type EXPRESSIONMIB_Expobjecttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an object.  An application uses expObjectStatus to create
    // entries in this table while in the process of defining an expression. 
    // Values of read-create objects in this table may be changed at any time. The
    // type is slice of EXPRESSIONMIB_Expobjecttable_Expobjectentry.
    Expobjectentry []EXPRESSIONMIB_Expobjecttable_Expobjectentry
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetFilter() yfilter.YFilter { return expobjecttable.YFilter }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) SetFilter(yf yfilter.YFilter) { expobjecttable.YFilter = yf }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetGoName(yname string) string {
    if yname == "expObjectEntry" { return "Expobjectentry" }
    return ""
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetSegmentPath() string {
    return "expObjectTable"
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expObjectEntry" {
        for _, c := range expobjecttable.Expobjectentry {
            if expobjecttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EXPRESSIONMIB_Expobjecttable_Expobjectentry{}
        expobjecttable.Expobjectentry = append(expobjecttable.Expobjectentry, child)
        return &expobjecttable.Expobjectentry[len(expobjecttable.Expobjectentry)-1]
    }
    return nil
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expobjecttable.Expobjectentry {
        children[expobjecttable.Expobjectentry[i].GetSegmentPath()] = &expobjecttable.Expobjectentry[i]
    }
    return children
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetBundleName() string { return "cisco_ios_xe" }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetYangName() string { return "expObjectTable" }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) SetParent(parent types.Entity) { expobjecttable.parent = parent }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetParent() types.Entity { return expobjecttable.parent }

func (expobjecttable *EXPRESSIONMIB_Expobjecttable) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expobjecttable_Expobjectentry
// Information about an object.  An application uses
// expObjectStatus to create entries in this table while
// in the process of defining an expression.
// 
// Values of read-create objects in this table may be
// changed at any time.
type EXPRESSIONMIB_Expobjecttable_Expobjectentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_Expnametable_Expnameentry_Expexpressionindex
    Expexpressionindex interface{}

    // This attribute is a key. Within an expression, a unique, numeric
    // identification for an object.  Prefixed with a dollar sign ('$') this is
    // used to reference the object in the corresponding expExpression. The type
    // is interface{} with range: 1..4294967295.
    Expobjectindex interface{}

    // The OBJECT IDENTIFIER (OID) of this object.  The OID may be fully
    // qualified, meaning it includes a complete instance identifier part (e.g.,
    // ifInOctets.1 or sysUpTime.0), or it may not be fully qualified, meaning it
    // may lack all or part of the instance identifier.  If the expObjectID is not
    // fully qualified, then expObjectWildcard must be set to true(1).   The value
    // of the expression will be multiple values, as if done for a GetNext sweep
    // of the object.  An object here may itself be the result of an expression
    // but recursion is not allowed.  NOTE:  The simplest implementations of this
    // MIB may not allow wildcards. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectid interface{}

    // A true value indicates the expObjecID of this row is a wildcard object.
    // False indicates that expObjectID is fully instanced.  If all
    // expObjectWildcard values for a given expression are FALSE,
    // expExpressionPrefix will reflect a scalar object (ie will be 0.0).  NOTE: 
    // The simplest implementations of this MIB may not allow wildcards. The type
    // is bool.
    Expobjectidwildcard interface{}

    // The method of sampling the selected variable.  An 'absoluteValue' is simply
    // the present value of the object. A 'deltaValue' is the present value minus
    // the previous value, which was sampled expExpressionDeltaInterval seconds
    // ago.  This is intended primarily for use with SNMP counters, which are
    // meaningless as an 'absoluteValue', but may be used with any integer-based
    // value.  When an expression contains both delta and absolute values the
    // absolute values are obtained at the end of the delta period. The type is
    // Expobjectsampletype.
    Expobjectsampletype interface{}

    // The OBJECT IDENTIFIER (OID) of a TimeTicks or TimeStamp object that
    // indicates a discontinuity in the value at expObjectID.  This object is not
    // instantiated if expObject is not 'deltaValue'.  The OID may be for a leaf
    // object (e.g. sysUpTime.0) or may be wildcarded to match expObjectID.  This
    // object supports normal checking for a discontinuity in a counter.  Note
    // that if this object does not point to sysUpTime discontinuity checking must
    // still check sysUpTime for an overall discontinuity.  If the object
    // identified is not accessible no discontinuity check will be made. The type
    // is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectdeltadiscontinuityid interface{}

    // A true value indicates the expObjectDeltaDiscontinuityID of this row is a
    // wildcard object.  False indicates that expObjectDeltaDiscontinuityID is
    // fully instanced.  This object is not instantiated if expObject is not
    // 'deltaValue'.  NOTE:  The simplest implementations of this MIB may not
    // allow wildcards. The type is bool.
    Expobjectdiscontinuityidwildcard interface{}

    // The value 'timeTicks' indicates the expObjectDeltaDiscontinuityID of this
    // row is of syntax TimeTicks.  The value 'timeStamp' indicates that
    // expObjectDeltaDiscontinuityID is of syntax TimeStamp.  This object is not
    // instantiated if expObject is not 'deltaValue'. The type is
    // Expobjectdiscontinuityidtype.
    Expobjectdiscontinuityidtype interface{}

    // The OBJECT IDENTIFIER (OID) of an object that overrides whether the
    // instance of expObjectID is to be considered usable.  If the value of the
    // object at expObjectConditional is 0 or not instantiated, the object at
    // expObjectID is treated as if it is not instantiated.  In other words,
    // expObjectConditional is a filter that controls whether or not to use the
    // value at expObjectID.  The OID may be for a leaf object (e.g.
    // sysObjectID.0) or may be wildcarded to match expObjectID.  If expObject is
    // wildcarded and expObjectID in the same row is not, the wild portion of
    // expObjectConditional must match the wildcarding of the rest of the
    // expression.  If no object in the expression is wildcarded but
    // expObjectConditional is, use the lexically first instance (if any) of
    // expObjectConditional.  If the value of expObjectConditional is 0.0
    // operation is as if the value pointed to by expObjectConditional is a
    // non-zero (true) value.  Note that expObjectConditional can not trivially
    // use an object of syntax TruthValue, since the underlying value is not 0 or
    // 1. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expobjectconditional interface{}

    // A true value indicates the expObjectConditional of this row is a wildcard
    // object. False indicates that expObjectConditional is fully instanced. 
    // NOTE: The simplest implementations of this MIB may not allow wildcards. The
    // type is bool.
    Expobjectconditionalwildcard interface{}

    // The control that allows creation/deletion of entries.  Objects in this
    // table may be changed while expObjectStatus is in any state. The type is
    // RowStatus.
    Expobjectstatus interface{}
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetFilter() yfilter.YFilter { return expobjectentry.YFilter }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) SetFilter(yf yfilter.YFilter) { expobjectentry.YFilter = yf }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetGoName(yname string) string {
    if yname == "expExpressionIndex" { return "Expexpressionindex" }
    if yname == "expObjectIndex" { return "Expobjectindex" }
    if yname == "expObjectID" { return "Expobjectid" }
    if yname == "expObjectIDWildcard" { return "Expobjectidwildcard" }
    if yname == "expObjectSampleType" { return "Expobjectsampletype" }
    if yname == "expObjectDeltaDiscontinuityID" { return "Expobjectdeltadiscontinuityid" }
    if yname == "expObjectDiscontinuityIDWildcard" { return "Expobjectdiscontinuityidwildcard" }
    if yname == "expObjectDiscontinuityIDType" { return "Expobjectdiscontinuityidtype" }
    if yname == "expObjectConditional" { return "Expobjectconditional" }
    if yname == "expObjectConditionalWildcard" { return "Expobjectconditionalwildcard" }
    if yname == "expObjectStatus" { return "Expobjectstatus" }
    return ""
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetSegmentPath() string {
    return "expObjectEntry" + "[expExpressionIndex='" + fmt.Sprintf("%v", expobjectentry.Expexpressionindex) + "']" + "[expObjectIndex='" + fmt.Sprintf("%v", expobjectentry.Expobjectindex) + "']"
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionIndex"] = expobjectentry.Expexpressionindex
    leafs["expObjectIndex"] = expobjectentry.Expobjectindex
    leafs["expObjectID"] = expobjectentry.Expobjectid
    leafs["expObjectIDWildcard"] = expobjectentry.Expobjectidwildcard
    leafs["expObjectSampleType"] = expobjectentry.Expobjectsampletype
    leafs["expObjectDeltaDiscontinuityID"] = expobjectentry.Expobjectdeltadiscontinuityid
    leafs["expObjectDiscontinuityIDWildcard"] = expobjectentry.Expobjectdiscontinuityidwildcard
    leafs["expObjectDiscontinuityIDType"] = expobjectentry.Expobjectdiscontinuityidtype
    leafs["expObjectConditional"] = expobjectentry.Expobjectconditional
    leafs["expObjectConditionalWildcard"] = expobjectentry.Expobjectconditionalwildcard
    leafs["expObjectStatus"] = expobjectentry.Expobjectstatus
    return leafs
}

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetBundleName() string { return "cisco_ios_xe" }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetYangName() string { return "expObjectEntry" }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) SetParent(parent types.Entity) { expobjectentry.parent = parent }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetParent() types.Entity { return expobjectentry.parent }

func (expobjectentry *EXPRESSIONMIB_Expobjecttable_Expobjectentry) GetParentYangName() string { return "expObjectTable" }

// EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype represents 'deltaValue'.
type EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype string

const (
    EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype_timeTicks EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype = "timeTicks"

    EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype_timeStamp EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectdiscontinuityidtype = "timeStamp"
)

// EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype represents period.
type EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype string

const (
    EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype_absoluteValue EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype = "absoluteValue"

    EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype_deltaValue EXPRESSIONMIB_Expobjecttable_Expobjectentry_Expobjectsampletype = "deltaValue"
)

// EXPRESSIONMIB_Expvaluetable
// A table of values from evaluated expressions.
type EXPRESSIONMIB_Expvaluetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single value from an evaluated expression.  For a given instance, only
    // one 'Val' object in the conceptual row will be instantiated, that is, the
    // one with the appropriate type for the value.  For values that contain no
    // objects of  expObjectSampleType 'deltaValue', reading a value from the
    // table causes the evaluation of the expression for that value.  For those
    // that contain a 'deltaValue' the value read is as of the last delta
    // interval.  If in the attempt to evaluate the expression one or more of the
    // necessary objects is not available, the corresponding entry in this table
    // is effectively not instantiated.  To maintain security of MIB information,
    // expression evaluation must take place using security credentials for the
    // implied Gets of the objects in the expression.  For expressions with no
    // deltaValue those security credentials are the ones that came with the Get*
    // for the value.  For expressions with a deltaValue the ongoing expression
    // evaluation is under the security credentials of the creator of the
    // corresponding expNameEntry. The type is slice of
    // EXPRESSIONMIB_Expvaluetable_Expvalueentry.
    Expvalueentry []EXPRESSIONMIB_Expvaluetable_Expvalueentry
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetFilter() yfilter.YFilter { return expvaluetable.YFilter }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) SetFilter(yf yfilter.YFilter) { expvaluetable.YFilter = yf }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetGoName(yname string) string {
    if yname == "expValueEntry" { return "Expvalueentry" }
    return ""
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetSegmentPath() string {
    return "expValueTable"
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "expValueEntry" {
        for _, c := range expvaluetable.Expvalueentry {
            if expvaluetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EXPRESSIONMIB_Expvaluetable_Expvalueentry{}
        expvaluetable.Expvalueentry = append(expvaluetable.Expvalueentry, child)
        return &expvaluetable.Expvalueentry[len(expvaluetable.Expvalueentry)-1]
    }
    return nil
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range expvaluetable.Expvalueentry {
        children[expvaluetable.Expvalueentry[i].GetSegmentPath()] = &expvaluetable.Expvalueentry[i]
    }
    return children
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetBundleName() string { return "cisco_ios_xe" }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetYangName() string { return "expValueTable" }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) SetParent(parent types.Entity) { expvaluetable.parent = parent }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetParent() types.Entity { return expvaluetable.parent }

func (expvaluetable *EXPRESSIONMIB_Expvaluetable) GetParentYangName() string { return "EXPRESSION-MIB" }

// EXPRESSIONMIB_Expvaluetable_Expvalueentry
// A single value from an evaluated expression.  For a given
// instance, only one 'Val' object in the conceptual row will
// be instantiated, that is, the one with the appropriate type
// for the value.  For values that contain no objects of 
// expObjectSampleType 'deltaValue', reading a value from the
// table causes the evaluation of the expression for that
// value.  For those that contain a 'deltaValue' the value
// read is as of the last delta interval.
// 
// If in the attempt to evaluate the expression one or more
// of the necessary objects is not available, the corresponding
// entry in this table is effectively not instantiated.
// 
// To maintain security of MIB information, expression
// evaluation must take place using security credentials for
// the implied Gets of the objects in the expression.  For
// expressions with no deltaValue those security credentials
// are the ones that came with the Get* for the value.  For
// expressions with a deltaValue the ongoing expression
// evaluation is under the security credentials of the
// creator of the corresponding expNameEntry.
type EXPRESSIONMIB_Expvaluetable_Expvalueentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_Expnametable_Expnameentry_Expexpressionindex
    Expexpressionindex interface{}

    // This attribute is a key. The final instance portion of a value's OID
    // according to the wildcarding in instances of expObjectID for the
    // expression.  The prefix of this OID fragment is 0.0, leading to the
    // following behavior.  If there is no wildcarding, the value is 0.0.0.  In
    // other words, there is one value which standing alone would have been a
    // scalar with a 0 at the end of its OID.  If there is wildcarding, the value
    // is 0.0 followed by a value that the wildcard can take, thus defining one
    // value instance for each real, possible value of the wildcard. So, for
    // example, if the wildcard worked out to be an ifIndex, there is an
    // expValueInstance for each applicable ifIndex. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expvalueinstance interface{}

    // The value when expExpressionValueType is 'counter32'. The type is
    // interface{} with range: 0..4294967295.
    Expvaluecounter32Val interface{}

    // The value when expExpressionValueType is 'unsignedOrGauge32' or
    // 'timeTicks'. The type is interface{} with range: 0..4294967295.
    Expvalueunsigned32Val interface{}

    // The value when expExpressionValueType is 'integer32'. The type is
    // interface{} with range: -2147483648..2147483647.
    Expvalueinteger32Val interface{}

    // The value when expExpressionValueType is 'ipAddress'. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Expvalueipaddressval interface{}

    // The value when expExpressionValueType is 'octetString'. The type is string
    // with length: 0..65535.
    Expvalueoctetstringval interface{}

    // The value when expExpressionValueType is 'objectId'. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Expvalueoidval interface{}

    // The value when expExpressionValueType is 'counter64'. The type is
    // interface{} with range: 0..18446744073709551615.
    Expvaluecounter64Val interface{}
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetFilter() yfilter.YFilter { return expvalueentry.YFilter }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) SetFilter(yf yfilter.YFilter) { expvalueentry.YFilter = yf }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetGoName(yname string) string {
    if yname == "expExpressionIndex" { return "Expexpressionindex" }
    if yname == "expValueInstance" { return "Expvalueinstance" }
    if yname == "expValueCounter32Val" { return "Expvaluecounter32Val" }
    if yname == "expValueUnsigned32Val" { return "Expvalueunsigned32Val" }
    if yname == "expValueInteger32Val" { return "Expvalueinteger32Val" }
    if yname == "expValueIpAddressVal" { return "Expvalueipaddressval" }
    if yname == "expValueOctetStringVal" { return "Expvalueoctetstringval" }
    if yname == "expValueOidVal" { return "Expvalueoidval" }
    if yname == "expValueCounter64Val" { return "Expvaluecounter64Val" }
    return ""
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetSegmentPath() string {
    return "expValueEntry" + "[expExpressionIndex='" + fmt.Sprintf("%v", expvalueentry.Expexpressionindex) + "']" + "[expValueInstance='" + fmt.Sprintf("%v", expvalueentry.Expvalueinstance) + "']"
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["expExpressionIndex"] = expvalueentry.Expexpressionindex
    leafs["expValueInstance"] = expvalueentry.Expvalueinstance
    leafs["expValueCounter32Val"] = expvalueentry.Expvaluecounter32Val
    leafs["expValueUnsigned32Val"] = expvalueentry.Expvalueunsigned32Val
    leafs["expValueInteger32Val"] = expvalueentry.Expvalueinteger32Val
    leafs["expValueIpAddressVal"] = expvalueentry.Expvalueipaddressval
    leafs["expValueOctetStringVal"] = expvalueentry.Expvalueoctetstringval
    leafs["expValueOidVal"] = expvalueentry.Expvalueoidval
    leafs["expValueCounter64Val"] = expvalueentry.Expvaluecounter64Val
    return leafs
}

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetBundleName() string { return "cisco_ios_xe" }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetYangName() string { return "expValueEntry" }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) SetParent(parent types.Entity) { expvalueentry.parent = parent }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetParent() types.Entity { return expvalueentry.parent }

func (expvalueentry *EXPRESSIONMIB_Expvaluetable_Expvalueentry) GetParentYangName() string { return "expValueTable" }

