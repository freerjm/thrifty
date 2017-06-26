// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package personservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"person"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = person.GoUnusedProtection__
type PersonService interface {  //Service responsible for all interactions with Person objects.

  Get() (r *person.Person, err error)
}

//Service responsible for all interactions with Person objects.
type PersonServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewPersonServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PersonServiceClient {
  return &PersonServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewPersonServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PersonServiceClient {
  return &PersonServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *PersonServiceClient) Get() (r *person.Person, err error) {
  if err = p.sendGet(); err != nil { return }
  return p.recvGet()
}

func (p *PersonServiceClient) sendGet()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("get", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := PersonServiceGetArgs{
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *PersonServiceClient) recvGet() (value *person.Person, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "get" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "get failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "get failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "get failed: invalid message type")
    return
  }
  result := PersonServiceGetResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type PersonServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler PersonService
}

func (p *PersonServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *PersonServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *PersonServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewPersonServiceProcessor(handler PersonService) *PersonServiceProcessor {

  self2 := &PersonServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["get"] = &personServiceProcessorGet{handler:handler}
return self2
}

func (p *PersonServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type personServiceProcessorGet struct {
  handler PersonService
}

func (p *personServiceProcessorGet) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := PersonServiceGetArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := PersonServiceGetResult{}
var retval *person.Person
  var err2 error
  if retval, err2 = p.handler.Get(); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get: " + err2.Error())
    oprot.WriteMessageBegin("get", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("get", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type PersonServiceGetArgs struct {
}

func NewPersonServiceGetArgs() *PersonServiceGetArgs {
  return &PersonServiceGetArgs{}
}

func (p *PersonServiceGetArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PersonServiceGetArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PersonServiceGetArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PersonServiceGetArgs(%+v)", *p)
}

// Attributes:
//  - Success
type PersonServiceGetResult struct {
  Success *person.Person `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewPersonServiceGetResult() *PersonServiceGetResult {
  return &PersonServiceGetResult{}
}

var PersonServiceGetResult_Success_DEFAULT *person.Person
func (p *PersonServiceGetResult) GetSuccess() *person.Person {
  if !p.IsSetSuccess() {
    return PersonServiceGetResult_Success_DEFAULT
  }
return p.Success
}
func (p *PersonServiceGetResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *PersonServiceGetResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *PersonServiceGetResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &person.Person{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *PersonServiceGetResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *PersonServiceGetResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *PersonServiceGetResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("PersonServiceGetResult(%+v)", *p)
}


