// Code generated by thriftrw v1.16.1. DO NOT EDIT.
// @generated

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ThriftTest_TestInsanity_Args represents the arguments for the ThriftTest.testInsanity function.
//
// The arguments for testInsanity are sent and received over the wire as this struct.
type ThriftTest_TestInsanity_Args struct {
	Argument *Insanity `json:"argument,omitempty"`
}

// ToWire translates a ThriftTest_TestInsanity_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestInsanity_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Argument != nil {
		w, err = v.Argument.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestInsanity_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestInsanity_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestInsanity_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestInsanity_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Argument, err = _Insanity_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestInsanity_Args
// struct.
func (v *ThriftTest_TestInsanity_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Argument != nil {
		fields[i] = fmt.Sprintf("Argument: %v", v.Argument)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestInsanity_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestInsanity_Args match the
// provided ThriftTest_TestInsanity_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestInsanity_Args) Equals(rhs *ThriftTest_TestInsanity_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Argument == nil && rhs.Argument == nil) || (v.Argument != nil && rhs.Argument != nil && v.Argument.Equals(rhs.Argument))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestInsanity_Args.
func (v *ThriftTest_TestInsanity_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Argument != nil {
		err = multierr.Append(err, enc.AddObject("argument", v.Argument))
	}
	return err
}

// GetArgument returns the value of Argument if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestInsanity_Args) GetArgument() (o *Insanity) {
	if v != nil && v.Argument != nil {
		return v.Argument
	}

	return
}

// IsSetArgument returns true if Argument is not nil.
func (v *ThriftTest_TestInsanity_Args) IsSetArgument() bool {
	return v != nil && v.Argument != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testInsanity" for this struct.
func (v *ThriftTest_TestInsanity_Args) MethodName() string {
	return "testInsanity"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestInsanity_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestInsanity_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testInsanity
// function.
var ThriftTest_TestInsanity_Helper = struct {
	// Args accepts the parameters of testInsanity in-order and returns
	// the arguments struct for the function.
	Args func(
		argument *Insanity,
	) *ThriftTest_TestInsanity_Args

	// IsException returns true if the given error can be thrown
	// by testInsanity.
	//
	// An error can be thrown by testInsanity only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testInsanity
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// testInsanity into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by testInsanity
	//
	//   value, err := testInsanity(args)
	//   result, err := ThriftTest_TestInsanity_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testInsanity: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(map[UserId]map[Numberz]*Insanity, error) (*ThriftTest_TestInsanity_Result, error)

	// UnwrapResponse takes the result struct for testInsanity
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if testInsanity threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ThriftTest_TestInsanity_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestInsanity_Result) (map[UserId]map[Numberz]*Insanity, error)
}{}

func init() {
	ThriftTest_TestInsanity_Helper.Args = func(
		argument *Insanity,
	) *ThriftTest_TestInsanity_Args {
		return &ThriftTest_TestInsanity_Args{
			Argument: argument,
		}
	}

	ThriftTest_TestInsanity_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ThriftTest_TestInsanity_Helper.WrapResponse = func(success map[UserId]map[Numberz]*Insanity, err error) (*ThriftTest_TestInsanity_Result, error) {
		if err == nil {
			return &ThriftTest_TestInsanity_Result{Success: success}, nil
		}

		return nil, err
	}
	ThriftTest_TestInsanity_Helper.UnwrapResponse = func(result *ThriftTest_TestInsanity_Result) (success map[UserId]map[Numberz]*Insanity, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ThriftTest_TestInsanity_Result represents the result of a ThriftTest.testInsanity function call.
//
// The result of a testInsanity execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ThriftTest_TestInsanity_Result struct {
	// Value returned by testInsanity after a successful execution.
	Success map[UserId]map[Numberz]*Insanity `json:"success,omitempty"`
}

type _Map_Numberz_Insanity_MapItemList map[Numberz]*Insanity

func (m _Map_Numberz_Insanity_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}

		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_Numberz_Insanity_MapItemList) Size() int {
	return len(m)
}

func (_Map_Numberz_Insanity_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_Numberz_Insanity_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_Numberz_Insanity_MapItemList) Close() {}

type _Map_UserId_Map_Numberz_Insanity_MapItemList map[UserId]map[Numberz]*Insanity

func (m _Map_UserId_Map_Numberz_Insanity_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}

		vw, err := wire.NewValueMap(_Map_Numberz_Insanity_MapItemList(v)), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_UserId_Map_Numberz_Insanity_MapItemList) Size() int {
	return len(m)
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) KeyType() wire.Type {
	return wire.TI64
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) ValueType() wire.Type {
	return wire.TMap
}

func (_Map_UserId_Map_Numberz_Insanity_MapItemList) Close() {}

// ToWire translates a ThriftTest_TestInsanity_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestInsanity_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_UserId_Map_Numberz_Insanity_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestInsanity_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_Numberz_Insanity_Read(m wire.MapItemList) (map[Numberz]*Insanity, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}

	if m.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make(map[Numberz]*Insanity, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _Numberz_Read(x.Key)
		if err != nil {
			return err
		}

		v, err := _Insanity_Read(x.Value)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _Map_UserId_Map_Numberz_Insanity_Read(m wire.MapItemList) (map[UserId]map[Numberz]*Insanity, error) {
	if m.KeyType() != wire.TI64 {
		return nil, nil
	}

	if m.ValueType() != wire.TMap {
		return nil, nil
	}

	o := make(map[UserId]map[Numberz]*Insanity, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _UserId_Read(x.Key)
		if err != nil {
			return err
		}

		v, err := _Map_Numberz_Insanity_Read(x.Value.GetMap())
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a ThriftTest_TestInsanity_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestInsanity_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestInsanity_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestInsanity_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_UserId_Map_Numberz_Insanity_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestInsanity_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestInsanity_Result
// struct.
func (v *ThriftTest_TestInsanity_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestInsanity_Result{%v}", strings.Join(fields[:i], ", "))
}

func _Map_Numberz_Insanity_Equals(lhs, rhs map[Numberz]*Insanity) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

func _Map_UserId_Map_Numberz_Insanity_Equals(lhs, rhs map[UserId]map[Numberz]*Insanity) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !_Map_Numberz_Insanity_Equals(lv, rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this ThriftTest_TestInsanity_Result match the
// provided ThriftTest_TestInsanity_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestInsanity_Result) Equals(rhs *ThriftTest_TestInsanity_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Map_UserId_Map_Numberz_Insanity_Equals(v.Success, rhs.Success))) {
		return false
	}

	return true
}

type _Map_Numberz_Insanity_Item_Zapper struct {
	Key   Numberz
	Value *Insanity
}

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Map_Numberz_Insanity_Item_Zapper.
func (v _Map_Numberz_Insanity_Item_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	err = multierr.Append(err, enc.AddObject("key", v.Key))
	err = multierr.Append(err, enc.AddObject("value", v.Value))
	return err
}

type _Map_Numberz_Insanity_Zapper map[Numberz]*Insanity

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Map_Numberz_Insanity_Zapper.
func (m _Map_Numberz_Insanity_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for k, v := range m {
		err = multierr.Append(err, enc.AppendObject(_Map_Numberz_Insanity_Item_Zapper{Key: k, Value: v}))
	}
	return err
}

type _Map_UserId_Map_Numberz_Insanity_Item_Zapper struct {
	Key   UserId
	Value map[Numberz]*Insanity
}

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Map_UserId_Map_Numberz_Insanity_Item_Zapper.
func (v _Map_UserId_Map_Numberz_Insanity_Item_Zapper) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	enc.AddInt64("key", (int64)(v.Key))
	err = multierr.Append(err, enc.AddArray("value", (_Map_Numberz_Insanity_Zapper)(v.Value)))
	return err
}

type _Map_UserId_Map_Numberz_Insanity_Zapper map[UserId]map[Numberz]*Insanity

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Map_UserId_Map_Numberz_Insanity_Zapper.
func (m _Map_UserId_Map_Numberz_Insanity_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for k, v := range m {
		err = multierr.Append(err, enc.AppendObject(_Map_UserId_Map_Numberz_Insanity_Item_Zapper{Key: k, Value: v}))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestInsanity_Result.
func (v *ThriftTest_TestInsanity_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddArray("success", (_Map_UserId_Map_Numberz_Insanity_Zapper)(v.Success)))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestInsanity_Result) GetSuccess() (o map[UserId]map[Numberz]*Insanity) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *ThriftTest_TestInsanity_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testInsanity" for this struct.
func (v *ThriftTest_TestInsanity_Result) MethodName() string {
	return "testInsanity"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestInsanity_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
