// Package logf comment
// This file was generated by tars2go 1.1.10
// Generated from LogF.tars
package logf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
)

// Log struct
type Log struct {
	servant m.Servant
}

// Logger is the proxy function for the method defined in the tars file, with the context
func (obj *Log) Logger(app string, server string, file string, format string, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(app, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(server, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(file, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(format, 4)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	tarsResp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = obj.servant.TarsInvoke(tarsCtx, 0, "logger", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// LoggerWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Log) LoggerWithContext(tarsCtx context.Context, app string, server string, file string, format string, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(app, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(server, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(file, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(format, 4)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "logger", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// LoggerOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Log) LoggerOneWayWithContext(tarsCtx context.Context, app string, server string, file string, format string, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = buf.WriteString(app, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(server, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(file, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(format, 4)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "logger", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// LoggerbyInfo is the proxy function for the method defined in the tars file, with the context
func (obj *Log) LoggerbyInfo(info *LogInfo, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = info.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	tarsResp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = obj.servant.TarsInvoke(tarsCtx, 0, "loggerbyInfo", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// LoggerbyInfoWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Log) LoggerbyInfoWithContext(tarsCtx context.Context, info *LogInfo, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = info.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 0, "loggerbyInfo", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// LoggerbyInfoOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (obj *Log) LoggerbyInfoOneWayWithContext(tarsCtx context.Context, info *LogInfo, buffer []string, opts ...map[string]string) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = info.WriteBlock(buf, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(buffer)), 0)
	if err != nil {
		return err
	}

	for _, v := range buffer {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	tarsResp := new(requestf.ResponsePacket)
	err = obj.servant.TarsInvoke(tarsCtx, 1, "loggerbyInfo", buf.ToBytes(), statusMap, contextMap, tarsResp)
	if err != nil {
		return err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range tarsResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range tarsResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return nil
}

// SetServant sets servant for the service.
func (obj *Log) SetServant(servant m.Servant) {
	obj.servant = servant
}

// TarsSetTimeout sets the timeout for the servant which is in ms.
func (obj *Log) TarsSetTimeout(timeout int) {
	obj.servant.TarsSetTimeout(timeout)
}

// TarsSetProtocol sets the protocol for the servant.
func (obj *Log) TarsSetProtocol(p m.Protocol) {
	obj.servant.TarsSetProtocol(p)
}

type LogServant interface {
	Logger(app string, server string, file string, format string, buffer []string) (err error)
	LoggerbyInfo(info *LogInfo, buffer []string) (err error)
}
type LogServantWithContext interface {
	Logger(tarsCtx context.Context, app string, server string, file string, format string, buffer []string) (err error)
	LoggerbyInfo(tarsCtx context.Context, info *LogInfo, buffer []string) (err error)
}

// Dispatch is used to call the server side implement for the method defined in the tars file. withContext shows using context or not.
func (obj *Log) Dispatch(tarsCtx context.Context, val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	buf := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "logger":
		var app string
		var server string
		var file string
		var format string
		var buffer []string
		buffer = make([]string, 0)

		if tarsReq.IVersion == basef.TARSVERSION {

			err = readBuf.ReadString(&app, 1, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&server, 2, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&file, 3, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&format, 4, true)
			if err != nil {
				return err
			}

			_, ty, err = readBuf.SkipToNoCheck(5, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = readBuf.ReadInt32(&length, 0, true)
				if err != nil {
					return err
				}

				buffer = make([]string, length)
				for i0, e0 := int32(0), length; i0 < e0; i0++ {

					err = readBuf.ReadString(&buffer[i0], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SimpleList {
				err = fmt.Errorf("not support SimpleList type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}
		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("app", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = readBuf.ReadString(&app, 0, true)
			if err != nil {
				return err
			}

			reqTup.GetBuffer("server", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = readBuf.ReadString(&server, 0, true)
			if err != nil {
				return err
			}

			reqTup.GetBuffer("file", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = readBuf.ReadString(&file, 0, true)
			if err != nil {
				return err
			}

			reqTup.GetBuffer("format", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = readBuf.ReadString(&format, 0, true)
			if err != nil {
				return err
			}

			reqTup.GetBuffer("buffer", &tupBuffer)
			readBuf.Reset(tupBuffer)
			_, ty, err = readBuf.SkipToNoCheck(0, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = readBuf.ReadInt32(&length, 0, true)
				if err != nil {
					return err
				}

				buffer = make([]string, length)
				for i1, e1 := int32(0), length; i1 < e1; i1++ {

					err = readBuf.ReadString(&buffer[i1], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SimpleList {
				err = fmt.Errorf("not support SimpleList type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["app"])
				if err = json.Unmarshal(jsonStr, &app); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["server"])
				if err = json.Unmarshal(jsonStr, &server); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["file"])
				if err = json.Unmarshal(jsonStr, &file); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["format"])
				if err = json.Unmarshal(jsonStr, &format); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["buffer"])
				if err = json.Unmarshal(jsonStr, &buffer); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		if !withContext {
			imp := val.(LogServant)
			err = imp.Logger(app, server, file, format, buffer)
		} else {
			imp := val.(LogServantWithContext)
			err = imp.Logger(tarsCtx, app, server, file, format, buffer)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}
	case "loggerbyInfo":
		var info LogInfo
		var buffer []string
		buffer = make([]string, 0)

		if tarsReq.IVersion == basef.TARSVERSION {

			err = info.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

			_, ty, err = readBuf.SkipToNoCheck(2, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = readBuf.ReadInt32(&length, 0, true)
				if err != nil {
					return err
				}

				buffer = make([]string, length)
				for i2, e2 := int32(0), length; i2 < e2; i2++ {

					err = readBuf.ReadString(&buffer[i2], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SimpleList {
				err = fmt.Errorf("not support SimpleList type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}
		} else if tarsReq.IVersion == basef.TUPVERSION {
			reqTup := tup.NewUniAttribute()
			reqTup.Decode(readBuf)

			var tupBuffer []byte

			reqTup.GetBuffer("info", &tupBuffer)
			readBuf.Reset(tupBuffer)
			err = info.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

			reqTup.GetBuffer("buffer", &tupBuffer)
			readBuf.Reset(tupBuffer)
			_, ty, err = readBuf.SkipToNoCheck(0, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = readBuf.ReadInt32(&length, 0, true)
				if err != nil {
					return err
				}

				buffer = make([]string, length)
				for i3, e3 := int32(0), length; i3 < e3; i3++ {

					err = readBuf.ReadString(&buffer[i3], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SimpleList {
				err = fmt.Errorf("not support SimpleList type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["info"])
				info.ResetDefault()
				if err = json.Unmarshal(jsonStr, &info); err != nil {
					return err
				}
			}
			{
				jsonStr, _ := json.Marshal(jsonData["buffer"])
				if err = json.Unmarshal(jsonStr, &buffer); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", tarsReq.IVersion)
			return err
		}

		if !withContext {
			imp := val.(LogServant)
			err = imp.LoggerbyInfo(&info, buffer)
		} else {
			imp := val.(LogServantWithContext)
			err = imp.LoggerbyInfo(tarsCtx, &info, buffer)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			buf.Reset()

		} else if tarsReq.IVersion == basef.TUPVERSION {
			rspTup := tup.NewUniAttribute()

			buf.Reset()
			err = rspTup.Encode(buf)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(tarsCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(tarsCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}
