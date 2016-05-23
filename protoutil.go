package iespec

import (
	"github.com/calmh/ipfix"
	"github.com/golang/protobuf/proto"

	"iespec/protomsg"
)

func ConvertFieldListToProtobuf(fieldList []ipfix.InterpretedField) *protomsg.ZFlow {
	pmsg := &protomsg.ZFlow{}

	for _, field := range fieldList {
		switch field.Name {

		case "protocolIdentifier":
			pmsg.ProtocolIdentifier = proto.Int32(int32(field.Value.(uint8)))
		case "zflowPid":
			pmsg.ZflowPid = proto.Int32(int32(field.Value.(uint32)))
		case "zflowParentPid":
			pmsg.ZflowParentPid = proto.Int32(int32(field.Value.(uint32)))
		case "zflowCantHash":
			pmsg.ZflowCantHash = proto.Bool(field.Value.(bool))
		case "zflowParentCantHash":
			pmsg.ZflowParentCantHash = proto.Bool(field.Value.(bool))
		case "destinationTransportPort":
			pmsg.DestinationTransportPort = proto.Int32(int32(field.Value.(uint16)))
		case "octetDeltaCount":
			pmsg.OctetDeltaCount = proto.Int64(int64(field.Value.(uint64)))
		case "packetDeltaCount":
			pmsg.PacketDeltaCount = proto.Int64(int64(field.Value.(uint64)))
		case "flowStartMilliseconds":
			pmsg.FlowStartMilliseconds = proto.Int64(int64(field.Value.(uint64)))
		case "flowEndMilliseconds":
			pmsg.FlowEndMilliseconds = proto.Int64(int64(field.Value.(uint64)))
		case "flowDirection":
			pmsg.FlowDirection = proto.Int64(int64(field.Value.(uint8)))
		// string fields
		case "userName":
			pmsg.UserName = proto.String(field.Value.(string))
		case "zflowAccount":
			pmsg.ZflowAccount = proto.String(field.Value.(string))
		case "zflowImagePath":
			pmsg.ZflowImagePath = proto.String(field.Value.(string))
		case "zflowCommandLine":
			pmsg.ZflowCommandLine = proto.String(field.Value.(string))
		case "zflowImageBaseFileName":
			pmsg.ZflowImageBaseFileName = proto.String(field.Value.(string))
		case "zflowMD5":
			pmsg.ZflowMD5 = proto.String(field.Value.(string))
		case "zflowVerVersionString":
			pmsg.ZflowVerVersionString = proto.String(field.Value.(string))
		case "zflowVerCompanyName":
			pmsg.ZflowVerCompanyName = proto.String(field.Value.(string))
		case "zflowVerFileDescription":
			pmsg.ZflowVerFileDescription = proto.String(field.Value.(string))
		case "zflowVerProductName":
			pmsg.ZflowVerProductName = proto.String(field.Value.(string))
		case "zflowVerInternalName":
			pmsg.ZflowVerInternalName = proto.String(field.Value.(string))
		case "zflowVerLegalCopyright":
			pmsg.ZflowVerLegalCopyright = proto.String(field.Value.(string))
		case "zflowVerLegalTrademark":
			pmsg.ZflowVerLegalTrademark = proto.String(field.Value.(string))
		case "zflowVerOriginalFilename":
			pmsg.ZflowVerOriginalFilename = proto.String(field.Value.(string))
		case "zflowVerProductVersion":
			pmsg.ZflowVerProductVersion = proto.String(field.Value.(string))
		case "flowVerFileTextVersion":
			pmsg.FlowVerFileTextVersion = proto.String(field.Value.(string))
		case "zflowVerProductTextVersion":
			pmsg.ZflowVerProductTextVersion = proto.String(field.Value.(string))
		case "zflowParentImagePath":
			pmsg.ZflowParentImagePath = proto.String(field.Value.(string))
		case "zflowParentImageBaseFileName":
			pmsg.ZflowParentImageBaseFileName = proto.String(field.Value.(string))
		case "zflowParentMD5":
			pmsg.ZflowParentMD5 = proto.String(field.Value.(string))
		case "zflowParentVerVersionString":
			pmsg.ZflowParentVerVersionString = proto.String(field.Value.(string))
		case "zflowParentVerCompanyName":
			pmsg.ZflowParentVerCompanyName = proto.String(field.Value.(string))
		case "zflowParentVerFileDescription":
			pmsg.ZflowParentVerFileDescription = proto.String(field.Value.(string))
		case "zflowParentVerProductName":
			pmsg.ZflowParentVerProductName = proto.String(field.Value.(string))
		case "zflowParentVerInternalName":
			pmsg.ZflowParentVerInternalName = proto.String(field.Value.(string))
		case "zflowParentVerLegalCopyright":
			pmsg.ZflowParentVerLegalCopyright = proto.String(field.Value.(string))
		case "zflowParentVerLegalTrademark":
			pmsg.ZflowParentVerLegalTrademark = proto.String(field.Value.(string))
		case "zflowParentVerOriginalFilename":
			pmsg.ZflowParentVerOriginalFilename = proto.String(field.Value.(string))
		case "zflowParentVerProductVersion":
			pmsg.ZflowParentVerProductVersion = proto.String(field.Value.(string))
		case "zflowParentVerFileTextVersion":
			pmsg.ZflowParentVerFileTextVersion = proto.String(field.Value.(string))
		case "zflowParentVerProductTextVersion":
			pmsg.ZflowParentVerProductTextVersion = proto.String(field.Value.(string))
		case "zflowAgentGuid":
			pmsg.ZflowAgentGuid = proto.String(field.Value.(string))
		case "zflowMachineName":
			pmsg.ZflowMachineName = proto.String(field.Value.(string))
		case "zflowOSName":
			pmsg.ZflowOSName = proto.String(field.Value.(string))
		case "zflowOSVersion":
			pmsg.ZflowOSVersion = proto.String(field.Value.(string))
		case "sourceIPv4Address":
			pmsg.SourceIPv4Address = field.RawValue
		case "sourceTransportPort":
			pmsg.SourceTransportPort = proto.Int32(int32(field.Value.(uint16)))
		case "destinationIPv4Address":
			pmsg.DestinationIPv4Address = field.RawValue
		case "sourceIPv6Address":
			pmsg.SourceIPv6Address = field.RawValue
		case "destinationIPv6Address":
			pmsg.DestinationIPv6Address = field.RawValue

		}
	}

	return pmsg
}
