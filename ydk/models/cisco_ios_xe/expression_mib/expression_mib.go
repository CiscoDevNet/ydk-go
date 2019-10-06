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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ExpResource EXPRESSIONMIB_ExpResource

    
    ExpNames EXPRESSIONMIB_ExpNames

    // A table of expression names, for creating and deleting expressions.
    ExpNameTable EXPRESSIONMIB_ExpNameTable

    // A table of expression definitions.
    ExpExpressionTable EXPRESSIONMIB_ExpExpressionTable

    // A table of object definitions for each expExpression.  Wildcarding instance
    // IDs:  It is legal to omit all or part of the instance portion for some or
    // all of the objects in an expression. (See the DESCRIPTION of expObjectID
    // for details.  However, note that if more than one object in the same
    // expression is wildcarded in this way, they all must be objects where that
    // portion of the instance is the same.  In other words, all objects may be in
    // the same SEQUENCE or in different SEQUENCEs but with the same semantic
    // index value (e.g., a value of ifIndex) for the wildcarded portion.
    ExpObjectTable EXPRESSIONMIB_ExpObjectTable

    // A table of values from evaluated expressions.
    ExpValueTable EXPRESSIONMIB_ExpValueTable
}

func (eXPRESSIONMIB *EXPRESSIONMIB) GetEntityData() *types.CommonEntityData {
    eXPRESSIONMIB.EntityData.YFilter = eXPRESSIONMIB.YFilter
    eXPRESSIONMIB.EntityData.YangName = "EXPRESSION-MIB"
    eXPRESSIONMIB.EntityData.BundleName = "cisco_ios_xe"
    eXPRESSIONMIB.EntityData.ParentYangName = "EXPRESSION-MIB"
    eXPRESSIONMIB.EntityData.SegmentPath = "EXPRESSION-MIB:EXPRESSION-MIB"
    eXPRESSIONMIB.EntityData.AbsolutePath = eXPRESSIONMIB.EntityData.SegmentPath
    eXPRESSIONMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    eXPRESSIONMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    eXPRESSIONMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    eXPRESSIONMIB.EntityData.Children = types.NewOrderedMap()
    eXPRESSIONMIB.EntityData.Children.Append("expResource", types.YChild{"ExpResource", &eXPRESSIONMIB.ExpResource})
    eXPRESSIONMIB.EntityData.Children.Append("expNames", types.YChild{"ExpNames", &eXPRESSIONMIB.ExpNames})
    eXPRESSIONMIB.EntityData.Children.Append("expNameTable", types.YChild{"ExpNameTable", &eXPRESSIONMIB.ExpNameTable})
    eXPRESSIONMIB.EntityData.Children.Append("expExpressionTable", types.YChild{"ExpExpressionTable", &eXPRESSIONMIB.ExpExpressionTable})
    eXPRESSIONMIB.EntityData.Children.Append("expObjectTable", types.YChild{"ExpObjectTable", &eXPRESSIONMIB.ExpObjectTable})
    eXPRESSIONMIB.EntityData.Children.Append("expValueTable", types.YChild{"ExpValueTable", &eXPRESSIONMIB.ExpValueTable})
    eXPRESSIONMIB.EntityData.Leafs = types.NewOrderedMap()

    eXPRESSIONMIB.EntityData.YListKeys = []string {}

    return &(eXPRESSIONMIB.EntityData)
}

// EXPRESSIONMIB_ExpResource
type EXPRESSIONMIB_ExpResource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The minimum expExpressionDeltaInterval this system will accept.  A system
    // may use the larger values of this minimum to lessen the impact of
    // constantly computing deltas.  The value -1 indicates this system will not
    // accept deltaValue as a value for expObjectSampleType.  Unless explicitly
    // resource limited, a system's value for this object should be 1.  Changing
    // this value will not invalidate an existing setting of expObjectSampleType.
    // The type is interface{} with range: -1..-1 | 1..600. Units are seconds.
    ExpResourceDeltaMinimum interface{}

    // The maximum number of dynamic instance entries this system will support for
    // wildcarded delta objects in expressions. These are the entries that
    // maintain state, one for each instance of each deltaValue object for each
    // value of an expression.  A value of 0 indicates no preset limit, that is,
    // the limit is dynamic based on system operation and resources.  Unless
    // explicitly resource limited, a system's value for this object should be 0. 
    // Changing this value will not eliminate or inhibit existing delta wildcard
    // instance objects but will prevent the creation of more such objects. The
    // type is interface{} with range: 0..4294967295. Units are instances.
    ExpResourceDeltaWildcardInstanceMaximum interface{}

    // The number of currently active instance entries as defined for
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    ExpResourceDeltaWildcardInstances interface{}

    // The highest value of expResourceDeltaWildcardInstances that has occurred
    // since initialization of the management system. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    ExpResourceDeltaWildcardInstancesHigh interface{}

    // The number of times this system could not evaluate an expression because
    // that would have created a value instance in excess of
    // expResourceDeltaWildcardInstanceMaximum. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    ExpResourceDeltaWildcardInstanceResourceLacks interface{}
}

func (expResource *EXPRESSIONMIB_ExpResource) GetEntityData() *types.CommonEntityData {
    expResource.EntityData.YFilter = expResource.YFilter
    expResource.EntityData.YangName = "expResource"
    expResource.EntityData.BundleName = "cisco_ios_xe"
    expResource.EntityData.ParentYangName = "EXPRESSION-MIB"
    expResource.EntityData.SegmentPath = "expResource"
    expResource.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expResource.EntityData.SegmentPath
    expResource.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expResource.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expResource.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expResource.EntityData.Children = types.NewOrderedMap()
    expResource.EntityData.Leafs = types.NewOrderedMap()
    expResource.EntityData.Leafs.Append("expResourceDeltaMinimum", types.YLeaf{"ExpResourceDeltaMinimum", expResource.ExpResourceDeltaMinimum})
    expResource.EntityData.Leafs.Append("expResourceDeltaWildcardInstanceMaximum", types.YLeaf{"ExpResourceDeltaWildcardInstanceMaximum", expResource.ExpResourceDeltaWildcardInstanceMaximum})
    expResource.EntityData.Leafs.Append("expResourceDeltaWildcardInstances", types.YLeaf{"ExpResourceDeltaWildcardInstances", expResource.ExpResourceDeltaWildcardInstances})
    expResource.EntityData.Leafs.Append("expResourceDeltaWildcardInstancesHigh", types.YLeaf{"ExpResourceDeltaWildcardInstancesHigh", expResource.ExpResourceDeltaWildcardInstancesHigh})
    expResource.EntityData.Leafs.Append("expResourceDeltaWildcardInstanceResourceLacks", types.YLeaf{"ExpResourceDeltaWildcardInstanceResourceLacks", expResource.ExpResourceDeltaWildcardInstanceResourceLacks})

    expResource.EntityData.YListKeys = []string {}

    return &(expResource.EntityData)
}

// EXPRESSIONMIB_ExpNames
type EXPRESSIONMIB_ExpNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime the last time an expression was created or deleted
    // or had its name changed using expExpressionName. The type is interface{}
    // with range: 0..4294967295.
    ExpNameLastChange interface{}

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
    ExpNameHighestIndex interface{}
}

func (expNames *EXPRESSIONMIB_ExpNames) GetEntityData() *types.CommonEntityData {
    expNames.EntityData.YFilter = expNames.YFilter
    expNames.EntityData.YangName = "expNames"
    expNames.EntityData.BundleName = "cisco_ios_xe"
    expNames.EntityData.ParentYangName = "EXPRESSION-MIB"
    expNames.EntityData.SegmentPath = "expNames"
    expNames.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expNames.EntityData.SegmentPath
    expNames.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expNames.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expNames.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expNames.EntityData.Children = types.NewOrderedMap()
    expNames.EntityData.Leafs = types.NewOrderedMap()
    expNames.EntityData.Leafs.Append("expNameLastChange", types.YLeaf{"ExpNameLastChange", expNames.ExpNameLastChange})
    expNames.EntityData.Leafs.Append("expNameHighestIndex", types.YLeaf{"ExpNameHighestIndex", expNames.ExpNameHighestIndex})

    expNames.EntityData.YListKeys = []string {}

    return &(expNames.EntityData)
}

// EXPRESSIONMIB_ExpNameTable
// A table of expression names, for creating and deleting
// expressions.
type EXPRESSIONMIB_ExpNameTable struct {
    EntityData types.CommonEntityData
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
    // The type is slice of EXPRESSIONMIB_ExpNameTable_ExpNameEntry.
    ExpNameEntry []*EXPRESSIONMIB_ExpNameTable_ExpNameEntry
}

func (expNameTable *EXPRESSIONMIB_ExpNameTable) GetEntityData() *types.CommonEntityData {
    expNameTable.EntityData.YFilter = expNameTable.YFilter
    expNameTable.EntityData.YangName = "expNameTable"
    expNameTable.EntityData.BundleName = "cisco_ios_xe"
    expNameTable.EntityData.ParentYangName = "EXPRESSION-MIB"
    expNameTable.EntityData.SegmentPath = "expNameTable"
    expNameTable.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expNameTable.EntityData.SegmentPath
    expNameTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expNameTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expNameTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expNameTable.EntityData.Children = types.NewOrderedMap()
    expNameTable.EntityData.Children.Append("expNameEntry", types.YChild{"ExpNameEntry", nil})
    for i := range expNameTable.ExpNameEntry {
        expNameTable.EntityData.Children.Append(types.GetSegmentPath(expNameTable.ExpNameEntry[i]), types.YChild{"ExpNameEntry", expNameTable.ExpNameEntry[i]})
    }
    expNameTable.EntityData.Leafs = types.NewOrderedMap()

    expNameTable.EntityData.YListKeys = []string {}

    return &(expNameTable.EntityData)
}

// EXPRESSIONMIB_ExpNameTable_ExpNameEntry
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
type EXPRESSIONMIB_ExpNameTable_ExpNameEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the expression.  Choosing names with
    // useful lexical ordering supports using GetNext or GetBulk to retrieve a
    // useful subset of the table. The type is string with length: 1..64.
    ExpName interface{}

    // The numeric identification of the expression.  Applications may select this
    // number in ascending numerical order by using expNameHighestIndex as a hint
    // or may use any other acceptable, unused number.  Once set this value may
    // not be set to a different value. The type is interface{} with range:
    // 1..4294967295.
    ExpExpressionIndex interface{}

    // The control that allows creation/deletion of entries. The type is
    // RowStatus.
    ExpNameStatus interface{}
}

func (expNameEntry *EXPRESSIONMIB_ExpNameTable_ExpNameEntry) GetEntityData() *types.CommonEntityData {
    expNameEntry.EntityData.YFilter = expNameEntry.YFilter
    expNameEntry.EntityData.YangName = "expNameEntry"
    expNameEntry.EntityData.BundleName = "cisco_ios_xe"
    expNameEntry.EntityData.ParentYangName = "expNameTable"
    expNameEntry.EntityData.SegmentPath = "expNameEntry" + types.AddKeyToken(expNameEntry.ExpName, "expName")
    expNameEntry.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/expNameTable/" + expNameEntry.EntityData.SegmentPath
    expNameEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expNameEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expNameEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expNameEntry.EntityData.Children = types.NewOrderedMap()
    expNameEntry.EntityData.Leafs = types.NewOrderedMap()
    expNameEntry.EntityData.Leafs.Append("expName", types.YLeaf{"ExpName", expNameEntry.ExpName})
    expNameEntry.EntityData.Leafs.Append("expExpressionIndex", types.YLeaf{"ExpExpressionIndex", expNameEntry.ExpExpressionIndex})
    expNameEntry.EntityData.Leafs.Append("expNameStatus", types.YLeaf{"ExpNameStatus", expNameEntry.ExpNameStatus})

    expNameEntry.EntityData.YListKeys = []string {"ExpName"}

    return &(expNameEntry.EntityData)
}

// EXPRESSIONMIB_ExpExpressionTable
// A table of expression definitions.
type EXPRESSIONMIB_ExpExpressionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single expression.  An entry appears in this table when
    // an entry is created in expNameTable. Deleting that expNameTable entry
    // automatically deletes this entry and its associated expObjectTable entries.
    // Values of read-write objects in this table may be changed at any time. The
    // type is slice of EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry.
    ExpExpressionEntry []*EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry
}

func (expExpressionTable *EXPRESSIONMIB_ExpExpressionTable) GetEntityData() *types.CommonEntityData {
    expExpressionTable.EntityData.YFilter = expExpressionTable.YFilter
    expExpressionTable.EntityData.YangName = "expExpressionTable"
    expExpressionTable.EntityData.BundleName = "cisco_ios_xe"
    expExpressionTable.EntityData.ParentYangName = "EXPRESSION-MIB"
    expExpressionTable.EntityData.SegmentPath = "expExpressionTable"
    expExpressionTable.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expExpressionTable.EntityData.SegmentPath
    expExpressionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expExpressionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expExpressionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expExpressionTable.EntityData.Children = types.NewOrderedMap()
    expExpressionTable.EntityData.Children.Append("expExpressionEntry", types.YChild{"ExpExpressionEntry", nil})
    for i := range expExpressionTable.ExpExpressionEntry {
        expExpressionTable.EntityData.Children.Append(types.GetSegmentPath(expExpressionTable.ExpExpressionEntry[i]), types.YChild{"ExpExpressionEntry", expExpressionTable.ExpExpressionEntry[i]})
    }
    expExpressionTable.EntityData.Leafs = types.NewOrderedMap()

    expExpressionTable.EntityData.YListKeys = []string {}

    return &(expExpressionTable.EntityData)
}

// EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry
// Information about a single expression.  An entry appears
// in this table when an entry is created in expNameTable.
// Deleting that expNameTable entry automatically deletes
// this entry and its associated expObjectTable entries.
// 
// Values of read-write objects in this table may be changed
// at any time.
type EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_ExpNameTable_ExpNameEntry_ExpExpressionIndex
    ExpExpressionIndex interface{}

    // The unique name of the expression, the same as expName.  Use this object to
    // change the expression's name without changing its expExpressionIndex. The
    // type is string with length: 1..64.
    ExpExpressionName interface{}

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
    ExpExpression interface{}

    // The type of the expression value.  One and only one of the value objects in
    // expValueTable will be instantiated to match this type.  If the result of
    // the expression can not be made into this type, an invalidOperandType error
    // will occur. The type is ExpExpressionValueType.
    ExpExpressionValueType interface{}

    // A comment to explain the use or meaning of the expression. The type is
    // string.
    ExpExpressionComment interface{}

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
    ExpExpressionDeltaInterval interface{}

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
    ExpExpressionPrefix interface{}

    // The number of errors encountered while evaluating this expression.  Note
    // that an object in the expression not being accessible is not considered an
    // error.  It is a legitimate condition that causes the corresponding
    // expression value not to be instantiated. The type is interface{} with
    // range: 0..4294967295.
    ExpExpressionErrors interface{}

    // The value of sysUpTime the last time an error caused a failure to evaluate
    // this expression.  This object is not instantiated if there have been no
    // errors. The type is interface{} with range: 0..4294967295.
    ExpExpressionErrorTime interface{}

    // The 1-based character index into expExpression for where the error
    // occurred.  The value zero indicates irrelevance.  This object is not
    // instantiated if there have been no errors. The type is interface{} with
    // range: -2147483648..2147483647.
    ExpExpressionErrorIndex interface{}

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
    // if there have been no errors. The type is ExpExpressionError.
    ExpExpressionError interface{}

    // The expValueInstance being evaluated when the error occurred.  A
    // zero-length indicates irrelevance.  This object is not instantiated if
    // there have been no errors. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    ExpExpressionInstance interface{}

    // The entity that configured this entry and is therefore using the resources
    // assigned to it. The type is string with length: 0..255.
    ExpExpressionOwner interface{}
}

func (expExpressionEntry *EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry) GetEntityData() *types.CommonEntityData {
    expExpressionEntry.EntityData.YFilter = expExpressionEntry.YFilter
    expExpressionEntry.EntityData.YangName = "expExpressionEntry"
    expExpressionEntry.EntityData.BundleName = "cisco_ios_xe"
    expExpressionEntry.EntityData.ParentYangName = "expExpressionTable"
    expExpressionEntry.EntityData.SegmentPath = "expExpressionEntry" + types.AddKeyToken(expExpressionEntry.ExpExpressionIndex, "expExpressionIndex")
    expExpressionEntry.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/expExpressionTable/" + expExpressionEntry.EntityData.SegmentPath
    expExpressionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expExpressionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expExpressionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expExpressionEntry.EntityData.Children = types.NewOrderedMap()
    expExpressionEntry.EntityData.Leafs = types.NewOrderedMap()
    expExpressionEntry.EntityData.Leafs.Append("expExpressionIndex", types.YLeaf{"ExpExpressionIndex", expExpressionEntry.ExpExpressionIndex})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionName", types.YLeaf{"ExpExpressionName", expExpressionEntry.ExpExpressionName})
    expExpressionEntry.EntityData.Leafs.Append("expExpression", types.YLeaf{"ExpExpression", expExpressionEntry.ExpExpression})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionValueType", types.YLeaf{"ExpExpressionValueType", expExpressionEntry.ExpExpressionValueType})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionComment", types.YLeaf{"ExpExpressionComment", expExpressionEntry.ExpExpressionComment})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionDeltaInterval", types.YLeaf{"ExpExpressionDeltaInterval", expExpressionEntry.ExpExpressionDeltaInterval})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionPrefix", types.YLeaf{"ExpExpressionPrefix", expExpressionEntry.ExpExpressionPrefix})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionErrors", types.YLeaf{"ExpExpressionErrors", expExpressionEntry.ExpExpressionErrors})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionErrorTime", types.YLeaf{"ExpExpressionErrorTime", expExpressionEntry.ExpExpressionErrorTime})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionErrorIndex", types.YLeaf{"ExpExpressionErrorIndex", expExpressionEntry.ExpExpressionErrorIndex})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionError", types.YLeaf{"ExpExpressionError", expExpressionEntry.ExpExpressionError})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionInstance", types.YLeaf{"ExpExpressionInstance", expExpressionEntry.ExpExpressionInstance})
    expExpressionEntry.EntityData.Leafs.Append("expExpressionOwner", types.YLeaf{"ExpExpressionOwner", expExpressionEntry.ExpExpressionOwner})

    expExpressionEntry.EntityData.YListKeys = []string {"ExpExpressionIndex"}

    return &(expExpressionEntry.EntityData)
}

// EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError represents errors.
type EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError string

const (
    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_invalidSyntax EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "invalidSyntax"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_undefinedObjectIndex EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "undefinedObjectIndex"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_unrecognizedOperator EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "unrecognizedOperator"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_unrecognizedFunction EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "unrecognizedFunction"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_invalidOperandType EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "invalidOperandType"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_unmatchedParenthesis EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "unmatchedParenthesis"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_tooManyWildcardValues EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "tooManyWildcardValues"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_recursion EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "recursion"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_deltaTooShort EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "deltaTooShort"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_resourceUnavailable EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "resourceUnavailable"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError_divideByZero EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionError = "divideByZero"
)

// EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType represents type, an invalidOperandType error will occur.
type EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType string

const (
    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_counter32 EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "counter32"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_unsignedOrGauge32 EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "unsignedOrGauge32"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_timeTicks EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "timeTicks"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_integer32 EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "integer32"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_ipAddress EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "ipAddress"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_octetString EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "octetString"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_objectId EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "objectId"

    EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType_counter64 EXPRESSIONMIB_ExpExpressionTable_ExpExpressionEntry_ExpExpressionValueType = "counter64"
)

// EXPRESSIONMIB_ExpObjectTable
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
type EXPRESSIONMIB_ExpObjectTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an object.  An application uses expObjectStatus to create
    // entries in this table while in the process of defining an expression. 
    // Values of read-create objects in this table may be changed at any time. The
    // type is slice of EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry.
    ExpObjectEntry []*EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry
}

func (expObjectTable *EXPRESSIONMIB_ExpObjectTable) GetEntityData() *types.CommonEntityData {
    expObjectTable.EntityData.YFilter = expObjectTable.YFilter
    expObjectTable.EntityData.YangName = "expObjectTable"
    expObjectTable.EntityData.BundleName = "cisco_ios_xe"
    expObjectTable.EntityData.ParentYangName = "EXPRESSION-MIB"
    expObjectTable.EntityData.SegmentPath = "expObjectTable"
    expObjectTable.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expObjectTable.EntityData.SegmentPath
    expObjectTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expObjectTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expObjectTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expObjectTable.EntityData.Children = types.NewOrderedMap()
    expObjectTable.EntityData.Children.Append("expObjectEntry", types.YChild{"ExpObjectEntry", nil})
    for i := range expObjectTable.ExpObjectEntry {
        expObjectTable.EntityData.Children.Append(types.GetSegmentPath(expObjectTable.ExpObjectEntry[i]), types.YChild{"ExpObjectEntry", expObjectTable.ExpObjectEntry[i]})
    }
    expObjectTable.EntityData.Leafs = types.NewOrderedMap()

    expObjectTable.EntityData.YListKeys = []string {}

    return &(expObjectTable.EntityData)
}

// EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry
// Information about an object.  An application uses
// expObjectStatus to create entries in this table while
// in the process of defining an expression.
// 
// Values of read-create objects in this table may be
// changed at any time.
type EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_ExpNameTable_ExpNameEntry_ExpExpressionIndex
    ExpExpressionIndex interface{}

    // This attribute is a key. Within an expression, a unique, numeric
    // identification for an object.  Prefixed with a dollar sign ('$') this is
    // used to reference the object in the corresponding expExpression. The type
    // is interface{} with range: 1..4294967295.
    ExpObjectIndex interface{}

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
    ExpObjectID interface{}

    // A true value indicates the expObjecID of this row is a wildcard object.
    // False indicates that expObjectID is fully instanced.  If all
    // expObjectWildcard values for a given expression are FALSE,
    // expExpressionPrefix will reflect a scalar object (ie will be 0.0).  NOTE: 
    // The simplest implementations of this MIB may not allow wildcards. The type
    // is bool.
    ExpObjectIDWildcard interface{}

    // The method of sampling the selected variable.  An 'absoluteValue' is simply
    // the present value of the object. A 'deltaValue' is the present value minus
    // the previous value, which was sampled expExpressionDeltaInterval seconds
    // ago.  This is intended primarily for use with SNMP counters, which are
    // meaningless as an 'absoluteValue', but may be used with any integer-based
    // value.  When an expression contains both delta and absolute values the
    // absolute values are obtained at the end of the delta period. The type is
    // ExpObjectSampleType.
    ExpObjectSampleType interface{}

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
    ExpObjectDeltaDiscontinuityID interface{}

    // A true value indicates the expObjectDeltaDiscontinuityID of this row is a
    // wildcard object.  False indicates that expObjectDeltaDiscontinuityID is
    // fully instanced.  This object is not instantiated if expObject is not
    // 'deltaValue'.  NOTE:  The simplest implementations of this MIB may not
    // allow wildcards. The type is bool.
    ExpObjectDiscontinuityIDWildcard interface{}

    // The value 'timeTicks' indicates the expObjectDeltaDiscontinuityID of this
    // row is of syntax TimeTicks.  The value 'timeStamp' indicates that
    // expObjectDeltaDiscontinuityID is of syntax TimeStamp.  This object is not
    // instantiated if expObject is not 'deltaValue'. The type is
    // ExpObjectDiscontinuityIDType.
    ExpObjectDiscontinuityIDType interface{}

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
    ExpObjectConditional interface{}

    // A true value indicates the expObjectConditional of this row is a wildcard
    // object. False indicates that expObjectConditional is fully instanced. 
    // NOTE: The simplest implementations of this MIB may not allow wildcards. The
    // type is bool.
    ExpObjectConditionalWildcard interface{}

    // The control that allows creation/deletion of entries.  Objects in this
    // table may be changed while expObjectStatus is in any state. The type is
    // RowStatus.
    ExpObjectStatus interface{}
}

func (expObjectEntry *EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry) GetEntityData() *types.CommonEntityData {
    expObjectEntry.EntityData.YFilter = expObjectEntry.YFilter
    expObjectEntry.EntityData.YangName = "expObjectEntry"
    expObjectEntry.EntityData.BundleName = "cisco_ios_xe"
    expObjectEntry.EntityData.ParentYangName = "expObjectTable"
    expObjectEntry.EntityData.SegmentPath = "expObjectEntry" + types.AddKeyToken(expObjectEntry.ExpExpressionIndex, "expExpressionIndex") + types.AddKeyToken(expObjectEntry.ExpObjectIndex, "expObjectIndex")
    expObjectEntry.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/expObjectTable/" + expObjectEntry.EntityData.SegmentPath
    expObjectEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expObjectEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expObjectEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expObjectEntry.EntityData.Children = types.NewOrderedMap()
    expObjectEntry.EntityData.Leafs = types.NewOrderedMap()
    expObjectEntry.EntityData.Leafs.Append("expExpressionIndex", types.YLeaf{"ExpExpressionIndex", expObjectEntry.ExpExpressionIndex})
    expObjectEntry.EntityData.Leafs.Append("expObjectIndex", types.YLeaf{"ExpObjectIndex", expObjectEntry.ExpObjectIndex})
    expObjectEntry.EntityData.Leafs.Append("expObjectID", types.YLeaf{"ExpObjectID", expObjectEntry.ExpObjectID})
    expObjectEntry.EntityData.Leafs.Append("expObjectIDWildcard", types.YLeaf{"ExpObjectIDWildcard", expObjectEntry.ExpObjectIDWildcard})
    expObjectEntry.EntityData.Leafs.Append("expObjectSampleType", types.YLeaf{"ExpObjectSampleType", expObjectEntry.ExpObjectSampleType})
    expObjectEntry.EntityData.Leafs.Append("expObjectDeltaDiscontinuityID", types.YLeaf{"ExpObjectDeltaDiscontinuityID", expObjectEntry.ExpObjectDeltaDiscontinuityID})
    expObjectEntry.EntityData.Leafs.Append("expObjectDiscontinuityIDWildcard", types.YLeaf{"ExpObjectDiscontinuityIDWildcard", expObjectEntry.ExpObjectDiscontinuityIDWildcard})
    expObjectEntry.EntityData.Leafs.Append("expObjectDiscontinuityIDType", types.YLeaf{"ExpObjectDiscontinuityIDType", expObjectEntry.ExpObjectDiscontinuityIDType})
    expObjectEntry.EntityData.Leafs.Append("expObjectConditional", types.YLeaf{"ExpObjectConditional", expObjectEntry.ExpObjectConditional})
    expObjectEntry.EntityData.Leafs.Append("expObjectConditionalWildcard", types.YLeaf{"ExpObjectConditionalWildcard", expObjectEntry.ExpObjectConditionalWildcard})
    expObjectEntry.EntityData.Leafs.Append("expObjectStatus", types.YLeaf{"ExpObjectStatus", expObjectEntry.ExpObjectStatus})

    expObjectEntry.EntityData.YListKeys = []string {"ExpExpressionIndex", "ExpObjectIndex"}

    return &(expObjectEntry.EntityData)
}

// EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType represents 'deltaValue'.
type EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType string

const (
    EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType_timeTicks EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType = "timeTicks"

    EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType_timeStamp EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectDiscontinuityIDType = "timeStamp"
)

// EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType represents period.
type EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType string

const (
    EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType_absoluteValue EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType = "absoluteValue"

    EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType_deltaValue EXPRESSIONMIB_ExpObjectTable_ExpObjectEntry_ExpObjectSampleType = "deltaValue"
)

// EXPRESSIONMIB_ExpValueTable
// A table of values from evaluated expressions.
type EXPRESSIONMIB_ExpValueTable struct {
    EntityData types.CommonEntityData
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
    // EXPRESSIONMIB_ExpValueTable_ExpValueEntry.
    ExpValueEntry []*EXPRESSIONMIB_ExpValueTable_ExpValueEntry
}

func (expValueTable *EXPRESSIONMIB_ExpValueTable) GetEntityData() *types.CommonEntityData {
    expValueTable.EntityData.YFilter = expValueTable.YFilter
    expValueTable.EntityData.YangName = "expValueTable"
    expValueTable.EntityData.BundleName = "cisco_ios_xe"
    expValueTable.EntityData.ParentYangName = "EXPRESSION-MIB"
    expValueTable.EntityData.SegmentPath = "expValueTable"
    expValueTable.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/" + expValueTable.EntityData.SegmentPath
    expValueTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expValueTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expValueTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expValueTable.EntityData.Children = types.NewOrderedMap()
    expValueTable.EntityData.Children.Append("expValueEntry", types.YChild{"ExpValueEntry", nil})
    for i := range expValueTable.ExpValueEntry {
        expValueTable.EntityData.Children.Append(types.GetSegmentPath(expValueTable.ExpValueEntry[i]), types.YChild{"ExpValueEntry", expValueTable.ExpValueEntry[i]})
    }
    expValueTable.EntityData.Leafs = types.NewOrderedMap()

    expValueTable.EntityData.YListKeys = []string {}

    return &(expValueTable.EntityData)
}

// EXPRESSIONMIB_ExpValueTable_ExpValueEntry
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
type EXPRESSIONMIB_ExpValueTable_ExpValueEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // expression_mib.EXPRESSIONMIB_ExpNameTable_ExpNameEntry_ExpExpressionIndex
    ExpExpressionIndex interface{}

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
    ExpValueInstance interface{}

    // The value when expExpressionValueType is 'counter32'. The type is
    // interface{} with range: 0..4294967295.
    ExpValueCounter32Val interface{}

    // The value when expExpressionValueType is 'unsignedOrGauge32' or
    // 'timeTicks'. The type is interface{} with range: 0..4294967295.
    ExpValueUnsigned32Val interface{}

    // The value when expExpressionValueType is 'integer32'. The type is
    // interface{} with range: -2147483648..2147483647.
    ExpValueInteger32Val interface{}

    // The value when expExpressionValueType is 'ipAddress'. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ExpValueIpAddressVal interface{}

    // The value when expExpressionValueType is 'octetString'. The type is string
    // with length: 0..65535.
    ExpValueOctetStringVal interface{}

    // The value when expExpressionValueType is 'objectId'. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    ExpValueOidVal interface{}

    // The value when expExpressionValueType is 'counter64'. The type is
    // interface{} with range: 0..18446744073709551615.
    ExpValueCounter64Val interface{}
}

func (expValueEntry *EXPRESSIONMIB_ExpValueTable_ExpValueEntry) GetEntityData() *types.CommonEntityData {
    expValueEntry.EntityData.YFilter = expValueEntry.YFilter
    expValueEntry.EntityData.YangName = "expValueEntry"
    expValueEntry.EntityData.BundleName = "cisco_ios_xe"
    expValueEntry.EntityData.ParentYangName = "expValueTable"
    expValueEntry.EntityData.SegmentPath = "expValueEntry" + types.AddKeyToken(expValueEntry.ExpExpressionIndex, "expExpressionIndex") + types.AddKeyToken(expValueEntry.ExpValueInstance, "expValueInstance")
    expValueEntry.EntityData.AbsolutePath = "EXPRESSION-MIB:EXPRESSION-MIB/expValueTable/" + expValueEntry.EntityData.SegmentPath
    expValueEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    expValueEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    expValueEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    expValueEntry.EntityData.Children = types.NewOrderedMap()
    expValueEntry.EntityData.Leafs = types.NewOrderedMap()
    expValueEntry.EntityData.Leafs.Append("expExpressionIndex", types.YLeaf{"ExpExpressionIndex", expValueEntry.ExpExpressionIndex})
    expValueEntry.EntityData.Leafs.Append("expValueInstance", types.YLeaf{"ExpValueInstance", expValueEntry.ExpValueInstance})
    expValueEntry.EntityData.Leafs.Append("expValueCounter32Val", types.YLeaf{"ExpValueCounter32Val", expValueEntry.ExpValueCounter32Val})
    expValueEntry.EntityData.Leafs.Append("expValueUnsigned32Val", types.YLeaf{"ExpValueUnsigned32Val", expValueEntry.ExpValueUnsigned32Val})
    expValueEntry.EntityData.Leafs.Append("expValueInteger32Val", types.YLeaf{"ExpValueInteger32Val", expValueEntry.ExpValueInteger32Val})
    expValueEntry.EntityData.Leafs.Append("expValueIpAddressVal", types.YLeaf{"ExpValueIpAddressVal", expValueEntry.ExpValueIpAddressVal})
    expValueEntry.EntityData.Leafs.Append("expValueOctetStringVal", types.YLeaf{"ExpValueOctetStringVal", expValueEntry.ExpValueOctetStringVal})
    expValueEntry.EntityData.Leafs.Append("expValueOidVal", types.YLeaf{"ExpValueOidVal", expValueEntry.ExpValueOidVal})
    expValueEntry.EntityData.Leafs.Append("expValueCounter64Val", types.YLeaf{"ExpValueCounter64Val", expValueEntry.ExpValueCounter64Val})

    expValueEntry.EntityData.YListKeys = []string {"ExpExpressionIndex", "ExpValueInstance"}

    return &(expValueEntry.EntityData)
}

