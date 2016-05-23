package iespec
 import (
 "github.com/calmh/ipfix"
)
 var MyFields = [...]ipfix.DictionaryEntry{
ipfix.DictionaryEntry{Name: "flowStartMilliseconds", EnterpriseID: 44619, FieldID: 152, Type: ipfix.FieldTypes["unsigned64"],},
ipfix.DictionaryEntry{Name: "zflowVerVersionString", EnterpriseID: 44619, FieldID: 268, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerLegalTrademark", EnterpriseID: 44619, FieldID: 274, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentImagePath", EnterpriseID: 44619, FieldID: 279, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerCompanyName", EnterpriseID: 44619, FieldID: 269, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "sourceIPv4Address", EnterpriseID: 44619, FieldID: 8, Type: ipfix.FieldTypes["ipv4Address"],},
ipfix.DictionaryEntry{Name: "destinationTransportPort", EnterpriseID: 44619, FieldID: 11, Type: ipfix.FieldTypes["unsigned16"],},
ipfix.DictionaryEntry{Name: "flowDirection", EnterpriseID: 44619, FieldID: 61, Type: ipfix.FieldTypes["unsigned8"],},
ipfix.DictionaryEntry{Name: "userName", EnterpriseID: 44619, FieldID: 371, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowAccount", EnterpriseID: 44619, FieldID: 260, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowCommandLine", EnterpriseID: 44619, FieldID: 262, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowAgentGuid", EnterpriseID: 44619, FieldID: 266, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerProductTextVersion", EnterpriseID: 44619, FieldID: 293, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentPid", EnterpriseID: 44619, FieldID: 259, Type: ipfix.FieldTypes["unsigned32"],},
ipfix.DictionaryEntry{Name: "zflowMD5", EnterpriseID: 44619, FieldID: 264, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerProductName", EnterpriseID: 44619, FieldID: 271, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerInternalName", EnterpriseID: 44619, FieldID: 272, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "flowVerFileTextVersion", EnterpriseID: 44619, FieldID: 277, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentCantHash", EnterpriseID: 44619, FieldID: 282, Type: ipfix.FieldTypes["boolean"],},
ipfix.DictionaryEntry{Name: "zflowParentVerLegalTrademark", EnterpriseID: 44619, FieldID: 289, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowOSVersion", EnterpriseID: 44619, FieldID: 295, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerFileDescription", EnterpriseID: 44619, FieldID: 270, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerOriginalFilename", EnterpriseID: 44619, FieldID: 275, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerProductVersion", EnterpriseID: 44619, FieldID: 276, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowVerProductTextVersion", EnterpriseID: 44619, FieldID: 278, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentImageBaseFileName", EnterpriseID: 44619, FieldID: 280, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerLegalCopyright", EnterpriseID: 44619, FieldID: 288, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerProductVersion", EnterpriseID: 44619, FieldID: 291, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowImagePath", EnterpriseID: 44619, FieldID: 261, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowCantHash", EnterpriseID: 44619, FieldID: 265, Type: ipfix.FieldTypes["boolean"],},
ipfix.DictionaryEntry{Name: "zflowMachineName", EnterpriseID: 44619, FieldID: 267, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerFileDescription", EnterpriseID: 44619, FieldID: 285, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "sourceIPv6Address", EnterpriseID: 44619, FieldID: 27, Type: ipfix.FieldTypes["ipv6Address"],},
ipfix.DictionaryEntry{Name: "octetDeltaCount", EnterpriseID: 44619, FieldID: 1, Type: ipfix.FieldTypes["unsigned64"],},
ipfix.DictionaryEntry{Name: "zflowVerLegalCopyright", EnterpriseID: 44619, FieldID: 273, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerProductName", EnterpriseID: 44619, FieldID: 286, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerOriginalFilename", EnterpriseID: 44619, FieldID: 290, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerCompanyName", EnterpriseID: 44619, FieldID: 284, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "protocolIdentifier", EnterpriseID: 44619, FieldID: 4, Type: ipfix.FieldTypes["unsigned8"],},
ipfix.DictionaryEntry{Name: "destinationIPv4Address", EnterpriseID: 44619, FieldID: 12, Type: ipfix.FieldTypes["ipv4Address"],},
ipfix.DictionaryEntry{Name: "destinationIPv6Address", EnterpriseID: 44619, FieldID: 28, Type: ipfix.FieldTypes["ipv6Address"],},
ipfix.DictionaryEntry{Name: "flowEndMilliseconds", EnterpriseID: 44619, FieldID: 153, Type: ipfix.FieldTypes["unsigned64"],},
ipfix.DictionaryEntry{Name: "zflowPid", EnterpriseID: 44619, FieldID: 258, Type: ipfix.FieldTypes["unsigned32"],},
ipfix.DictionaryEntry{Name: "zflowParentMD5", EnterpriseID: 44619, FieldID: 281, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerVersionString", EnterpriseID: 44619, FieldID: 283, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerFileTextVersion", EnterpriseID: 44619, FieldID: 292, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowOSName", EnterpriseID: 44619, FieldID: 294, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "sourceTransportPort", EnterpriseID: 44619, FieldID: 7, Type: ipfix.FieldTypes["unsigned16"],},
ipfix.DictionaryEntry{Name: "zflowImageBaseFileName", EnterpriseID: 44619, FieldID: 263, Type: ipfix.FieldTypes["string"],},
ipfix.DictionaryEntry{Name: "zflowParentVerInternalName", EnterpriseID: 44619, FieldID: 287, Type: ipfix.FieldTypes["string"],},}