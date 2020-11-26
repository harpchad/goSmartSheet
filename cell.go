package gosmartsheet

import (
	"encoding/json"
	"strconv"
)

//Cell is a SmartSheet cell
type Cell struct {
	ColumnID     int64      `json:"columnId"`
	Value        *CellValue `json:"value,omitempty"` //TODO: should this be a pointer?
	DisplayValue string     `json:"displayValue,omitempty"`
}

//CellValue represents the possible strongly typed values that could exist in a SS cell
//another good article on it..
//http://attilaolah.eu/2013/11/29/json-decoding-in-go/
type CellValue struct {
	Value json.RawMessage

	StringVal *string
	IntVal    *int
	FloatVal  *float64
}

//CellUpdate to update a cell
type CellUpdate struct {
	ColumnID     int64      `json:"columnId"`
	Value        *CellValue `json:"value,omitempty"`
}

//StringDebug returns a debug string containing each of the underlying values of a Cell
func (c *CellValue) StringDebug() (val string) {
	delim := ""
	if c.StringVal != nil {
		val = "String Val: '" + *(c.StringVal) + "'"
		delim = " "
	}

	if c.IntVal != nil {
		val = val + delim + "Int Val: '" + strconv.Itoa(*c.IntVal) + "'"
		delim = " "
	}

	if c.FloatVal != nil {
		val = val + delim + "Float Val: '" + strconv.FormatFloat(*c.FloatVal, 'f', -1, 64) + "'"
		delim = " "
	}

	return
}

//String returns the underlying value as a string regardless of type
func (c *CellValue) String() (val string) {
	if c.StringVal != nil {
		val = *(c.StringVal)
		return
	}

	if c.IntVal != nil {
		val = strconv.Itoa(*c.IntVal)
		return
	}

	if c.FloatVal != nil {
		val = strconv.FormatFloat(*c.FloatVal, 'f', -1, 64) //-1 will remove unimportant 0s
		return
	}

	val = string(c.Value)
	return
}

//Int will return the Integer representation of the underlying value.  This should only be used if the value is known to be an Int
func (c *CellValue) Int() (val int) {
	if c.IntVal != nil {
		val = (*(c.IntVal))
	}

	return
}

//Float will return the Float representation of the underlying value.  This should only be used if the value is known to be an Float.
func (c *CellValue) Float() (val float64) {
	if c.FloatVal != nil {
		val = (*(c.FloatVal))
	}

	return
}

//SetString will clear all values and set only the string
//This should be used when updating an existing row especially if the type if changing
func (c *CellValue) SetString(v string) {
	c.IntVal = nil
	c.FloatVal = nil
	c.StringVal = &v
}

//SetInt will clear all values and set only the string
//This should be used when updating an existing row especially if the type if changing
func (c *CellValue) SetInt(v int) {
	c.IntVal = &v
	c.FloatVal = nil
	c.StringVal = nil
}

//SetFloat will clear all values and set only the string
//This should be used when updating an existing row especially if the type if changing
func (c *CellValue) SetFloat(v float64) {
	c.IntVal = nil
	c.FloatVal = &v
	c.StringVal = nil
}

//MarshalJSON is a custom marshaller for CellValue
func (c *CellValue) MarshalJSON() ([]byte, error) {
	if c.StringVal != nil {
		return json.Marshal(c.StringVal)
	}

	if c.IntVal != nil {
		return json.Marshal(c.IntVal)
	}

	if c.FloatVal != nil {
		return json.Marshal(c.FloatVal)
	}

	return json.Marshal(c.Value) //default raw message
}

//UnmarshalJSON is a custom unmarshaller for CellValue
func (c *CellValue) UnmarshalJSON(b []byte) (err error) {
	c.StringVal = nil
	c.IntVal = nil
	c.FloatVal = nil

	//errors unmarshalling to the corrsponding types should  not bubble up
	s := ""
	if e := json.Unmarshal(b, &s); e == nil {
		c.StringVal = &s
		return
	}
	var i int
	if e := json.Unmarshal(b, &i); e == nil {
		c.IntVal = &i
		return
	}
	var f float64
	if e := json.Unmarshal(b, &f); e == nil {
		c.FloatVal = &f
	}

	c.Value = json.RawMessage(b) //default to raw message
	return
}
