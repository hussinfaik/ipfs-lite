// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.4
// source: ss_ledger.proto

package ss_ledger

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SSLedger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partner            string  `protobuf:"bytes,1,opt,name=partner,proto3" json:"partner,omitempty"`
	LastMpExchange     int64   `protobuf:"varint,2,opt,name=last_mp_exchange,json=lastMpExchange,proto3" json:"last_mp_exchange,omitempty"`
	MpExchangeCount    uint64  `protobuf:"varint,3,opt,name=mp_exchange_count,json=mpExchangeCount,proto3" json:"mp_exchange_count,omitempty"`
	Whitelisted        bool    `protobuf:"varint,4,opt,name=whitelisted,proto3" json:"whitelisted,omitempty"`
	LastWhitelistCheck int64   `protobuf:"varint,5,opt,name=last_whitelist_check,json=lastWhitelistCheck,proto3" json:"last_whitelist_check,omitempty"`
	Invoice            float64 `protobuf:"fixed64,6,opt,name=invoice,proto3" json:"invoice,omitempty"`
	BytesPaid          uint64  `protobuf:"varint,7,opt,name=bytes_paid,json=bytesPaid,proto3" json:"bytes_paid,omitempty"`
	BytesPayRecvd      uint64  `protobuf:"varint,8,opt,name=bytes_pay_recvd,json=bytesPayRecvd,proto3" json:"bytes_pay_recvd,omitempty"`
	Sent               float64 `protobuf:"fixed64,9,opt,name=sent,proto3" json:"sent,omitempty"`
	Recvd              float64 `protobuf:"fixed64,10,opt,name=recvd,proto3" json:"recvd,omitempty"`
	Role               string  `protobuf:"bytes,11,opt,name=role,proto3" json:"role,omitempty"`
	SignedMp           string  `protobuf:"bytes,12,opt,name=signed_mp,json=signedMp,proto3" json:"signed_mp,omitempty"`
	DeviceId           string  `protobuf:"bytes,13,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	BytesSent          uint64  `protobuf:"varint,14,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	BytesRecvd         uint64  `protobuf:"varint,15,opt,name=bytes_recvd,json=bytesRecvd,proto3" json:"bytes_recvd,omitempty"`
	BlocksSent         uint64  `protobuf:"varint,16,opt,name=blocks_sent,json=blocksSent,proto3" json:"blocks_sent,omitempty"`
	BlocksRecvd        uint64  `protobuf:"varint,17,opt,name=blocks_recvd,json=blocksRecvd,proto3" json:"blocks_recvd,omitempty"`
	Metadata           []byte  `protobuf:"bytes,18,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SSLedger) Reset() {
	*x = SSLedger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSLedger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSLedger) ProtoMessage() {}

func (x *SSLedger) ProtoReflect() protoreflect.Message {
	mi := &file_ss_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSLedger.ProtoReflect.Descriptor instead.
func (*SSLedger) Descriptor() ([]byte, []int) {
	return file_ss_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *SSLedger) GetPartner() string {
	if x != nil {
		return x.Partner
	}
	return ""
}

func (x *SSLedger) GetLastMpExchange() int64 {
	if x != nil {
		return x.LastMpExchange
	}
	return 0
}

func (x *SSLedger) GetMpExchangeCount() uint64 {
	if x != nil {
		return x.MpExchangeCount
	}
	return 0
}

func (x *SSLedger) GetWhitelisted() bool {
	if x != nil {
		return x.Whitelisted
	}
	return false
}

func (x *SSLedger) GetLastWhitelistCheck() int64 {
	if x != nil {
		return x.LastWhitelistCheck
	}
	return 0
}

func (x *SSLedger) GetInvoice() float64 {
	if x != nil {
		return x.Invoice
	}
	return 0
}

func (x *SSLedger) GetBytesPaid() uint64 {
	if x != nil {
		return x.BytesPaid
	}
	return 0
}

func (x *SSLedger) GetBytesPayRecvd() uint64 {
	if x != nil {
		return x.BytesPayRecvd
	}
	return 0
}

func (x *SSLedger) GetSent() float64 {
	if x != nil {
		return x.Sent
	}
	return 0
}

func (x *SSLedger) GetRecvd() float64 {
	if x != nil {
		return x.Recvd
	}
	return 0
}

func (x *SSLedger) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SSLedger) GetSignedMp() string {
	if x != nil {
		return x.SignedMp
	}
	return ""
}

func (x *SSLedger) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *SSLedger) GetBytesSent() uint64 {
	if x != nil {
		return x.BytesSent
	}
	return 0
}

func (x *SSLedger) GetBytesRecvd() uint64 {
	if x != nil {
		return x.BytesRecvd
	}
	return 0
}

func (x *SSLedger) GetBlocksSent() uint64 {
	if x != nil {
		return x.BlocksSent
	}
	return 0
}

func (x *SSLedger) GetBlocksRecvd() uint64 {
	if x != nil {
		return x.BlocksRecvd
	}
	return 0
}

func (x *SSLedger) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_ss_ledger_proto protoreflect.FileDescriptor

var file_ss_ledger_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x22, 0xc7, 0x04, 0x0a,
	0x08, 0x53, 0x53, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x70, 0x5f, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x70, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x6d, 0x70, 0x5f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x70, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x68, 0x69,
	0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x70, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x50, 0x61, 0x79, 0x52, 0x65, 0x63, 0x76, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x63, 0x76, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x72, 0x65, 0x63, 0x76, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6d, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x73, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x63, 0x76, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x63, 0x76, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x76, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x63, 0x76, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ss_ledger_proto_rawDescOnce sync.Once
	file_ss_ledger_proto_rawDescData = file_ss_ledger_proto_rawDesc
)

func file_ss_ledger_proto_rawDescGZIP() []byte {
	file_ss_ledger_proto_rawDescOnce.Do(func() {
		file_ss_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_ss_ledger_proto_rawDescData)
	})
	return file_ss_ledger_proto_rawDescData
}

var file_ss_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ss_ledger_proto_goTypes = []interface{}{
	(*SSLedger)(nil), // 0: ss_ledger.SSLedger
}
var file_ss_ledger_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ss_ledger_proto_init() }
func file_ss_ledger_proto_init() {
	if File_ss_ledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ss_ledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSLedger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ss_ledger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ss_ledger_proto_goTypes,
		DependencyIndexes: file_ss_ledger_proto_depIdxs,
		MessageInfos:      file_ss_ledger_proto_msgTypes,
	}.Build()
	File_ss_ledger_proto = out.File
	file_ss_ledger_proto_rawDesc = nil
	file_ss_ledger_proto_goTypes = nil
	file_ss_ledger_proto_depIdxs = nil
}
