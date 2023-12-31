// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/richmedia/tts.proto

package richmedia

type TtsRspBody struct {
	RetCode        uint32          `protobuf:"varint,1,opt"`
	SessionId      string          `protobuf:"bytes,2,opt"`
	OutSeq         uint32          `protobuf:"varint,3,opt"`
	VoiceData      []*TtsVoiceItem `protobuf:"bytes,4,rep"`
	Islast         bool            `protobuf:"varint,5,opt"`
	PcmSampleRate  uint32          `protobuf:"varint,6,opt"`
	OpusSampleRate uint32          `protobuf:"varint,7,opt"`
	OpusChannels   uint32          `protobuf:"varint,8,opt"`
	OpusBitRate    uint32          `protobuf:"varint,9,opt"`
	OpusFrameSize  uint32          `protobuf:"varint,10,opt"`
}

type TtsVoiceItem struct {
	Voice []byte `protobuf:"bytes,1,opt"`
	Seq   uint32 `protobuf:"varint,2,opt"`
}