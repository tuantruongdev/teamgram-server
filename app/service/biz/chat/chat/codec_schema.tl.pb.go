/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package chat

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	-1542554274: func() mtproto.TLObject { // 0xa40e7d5e
		o := MakeTLChatInviteAlready(nil)
		o.Data2.Constructor = -1542554274
		return o
	},
	-613035609: func() mtproto.TLObject { // 0xdb75d1a7
		o := MakeTLChatInvite(nil)
		o.Data2.Constructor = -613035609
		return o
	},
	-1394351506: func() mtproto.TLObject { // 0xace3e26e
		o := MakeTLChatInvitePeek(nil)
		o.Data2.Constructor = -1394351506
		return o
	},
	1913672182: func() mtproto.TLObject { // 0x721051f6
		o := MakeTLChatInviteImported(nil)
		o.Data2.Constructor = 1913672182
		return o
	},
	476986452: func() mtproto.TLObject { // 0x1c6e3c54
		o := MakeTLRecentChatInviteRequesters(nil)
		o.Data2.Constructor = 476986452
		return o
	},
	1342599716: func() mtproto.TLObject { // 0x50067224
		o := MakeTLUserChatIdList(nil)
		o.Data2.Constructor = 1342599716
		return o
	},

	// Method
	741090770: func() mtproto.TLObject { // 0x2c2c25d2
		return &TLChatGetMutableChat{
			Constructor: 741090770,
		}
	},
	-415173319: func() mtproto.TLObject { // 0xe740f539
		return &TLChatGetChatListByIdList{
			Constructor: -415173319,
		}
	},
	1236736584: func() mtproto.TLObject { // 0x49b71a48
		return &TLChatGetChatBySelfId{
			Constructor: 1236736584,
		}
	},
	-465608273: func() mtproto.TLObject { // 0xe43f61af
		return &TLChatCreateChat2{
			Constructor: -465608273,
		}
	},
	1829891102: func() mtproto.TLObject { // 0x6d11ec1e
		return &TLChatDeleteChat{
			Constructor: 1829891102,
		}
	},
	187109333: func() mtproto.TLObject { // 0xb270fd5
		return &TLChatDeleteChatUser{
			Constructor: 187109333,
		}
	},
	-1782210905: func() mtproto.TLObject { // 0x95c59ea7
		return &TLChatEditChatTitle{
			Constructor: -1782210905,
		}
	},
	1551072376: func() mtproto.TLObject { // 0x5c737c78
		return &TLChatEditChatAbout{
			Constructor: 1551072376,
		}
	},
	1170384488: func() mtproto.TLObject { // 0x45c2a668
		return &TLChatEditChatPhoto{
			Constructor: 1170384488,
		}
	},
	419816940: func() mtproto.TLObject { // 0x1905e5ec
		return &TLChatEditChatAdmin{
			Constructor: 419816940,
		}
	},
	1513399943: func() mtproto.TLObject { // 0x5a34a687
		return &TLChatEditChatDefaultBannedRights{
			Constructor: 1513399943,
		}
	},
	2086511757: func() mtproto.TLObject { // 0x7c5da48d
		return &TLChatAddChatUser{
			Constructor: 2086511757,
		}
	},
	-1570363509: func() mtproto.TLObject { // 0xa266278b
		return &TLChatGetMutableChatByLink{
			Constructor: -1570363509,
		}
	},
	-711644423: func() mtproto.TLObject { // 0xd5952af9
		return &TLChatToggleNoForwards{
			Constructor: -711644423,
		}
	},
	138390239: func() mtproto.TLObject { // 0x83faadf
		return &TLChatMigratedToChannel{
			Constructor: 138390239,
		}
	},
	848700073: func() mtproto.TLObject { // 0x329622a9
		return &TLChatGetChatParticipantIdList{
			Constructor: 848700073,
		}
	},
	792111948: func() mtproto.TLObject { // 0x2f36ab4c
		return &TLChatGetUsersChatIdList{
			Constructor: 792111948,
		}
	},
	-210408312: func() mtproto.TLObject { // 0xf3756c88
		return &TLChatGetMyChatList{
			Constructor: -210408312,
		}
	},
	-976256949: func() mtproto.TLObject { // 0xc5cf804b
		return &TLChatExportChatInvite{
			Constructor: -976256949,
		}
	},
	-756399662: func() mtproto.TLObject { // 0xd2ea41d2
		return &TLChatGetAdminsWithInvites{
			Constructor: -756399662,
		}
	},
	-571854256: func() mtproto.TLObject { // 0xddea3250
		return &TLChatGetExportedChatInvite{
			Constructor: -571854256,
		}
	},
	-1265690378: func() mtproto.TLObject { // 0xb48f18f6
		return &TLChatGetExportedChatInvites{
			Constructor: -1265690378,
		}
	},
	1938289292: func() mtproto.TLObject { // 0x7387f28c
		return &TLChatCheckChatInvite{
			Constructor: 1938289292,
		}
	},
	1491493076: func() mtproto.TLObject { // 0x58e660d4
		return &TLChatImportChatInvite{
			Constructor: 1491493076,
		}
	},
	-1740221057: func() mtproto.TLObject { // 0x9846557f
		return &TLChatGetChatInviteImporters{
			Constructor: -1740221057,
		}
	},
	1445103800: func() mtproto.TLObject { // 0x562288b8
		return &TLChatDeleteExportedChatInvite{
			Constructor: 1445103800,
		}
	},
	-804101527: func() mtproto.TLObject { // 0xd0126269
		return &TLChatDeleteRevokedExportedChatInvites{
			Constructor: -804101527,
		}
	},
	-1348907914: func() mtproto.TLObject { // 0xaf994c76
		return &TLChatEditExportedChatInvite{
			Constructor: -1348907914,
		}
	},
	-992966286: func() mtproto.TLObject { // 0xc4d08972
		return &TLChatSetChatAvailableReactions{
			Constructor: -992966286,
		}
	},
	1023107972: func() mtproto.TLObject { // 0x3cfb6384
		return &TLChatSetHistoryTTL{
			Constructor: 1023107972,
		}
	},
	568333563: func() mtproto.TLObject { // 0x21e014fb
		return &TLChatSearch{
			Constructor: 568333563,
		}
	},
	-19132264: func() mtproto.TLObject { // 0xfedc1098
		return &TLChatGetRecentChatInviteRequesters{
			Constructor: -19132264,
		}
	},
	1051012305: func() mtproto.TLObject { // 0x3ea52cd1
		return &TLChatHideChatJoinRequests{
			Constructor: 1051012305,
		}
	},
	-589742657: func() mtproto.TLObject { // 0xdcd93dbf
		return &TLChatImportChatInvite2{
			Constructor: -589742657,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// ChatInviteExt <--
//  + TL_ChatInviteAlready
//  + TL_ChatInvite
//  + TL_ChatInvitePeek
//

func (m *ChatInviteExt) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_chatInviteAlready:
		t := m.To_ChatInviteAlready()
		xBuf = t.Encode(layer)
	case Predicate_chatInvite:
		t := m.To_ChatInvite()
		xBuf = t.Encode(layer)
	case Predicate_chatInvitePeek:
		t := m.To_ChatInvitePeek()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ChatInviteExt) CalcByteSize(layer int32) int {
	return 0
}

func (m *ChatInviteExt) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xa40e7d5e:
		m2 := MakeTLChatInviteAlready(m)
		m2.Decode(dBuf)
	case 0xdb75d1a7:
		m2 := MakeTLChatInvite(m)
		m2.Decode(dBuf)
	case 0xace3e26e:
		m2 := MakeTLChatInvitePeek(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ChatInviteExt) DebugString() string {
	switch m.PredicateName {
	case Predicate_chatInviteAlready:
		t := m.To_ChatInviteAlready()
		return t.DebugString()
	case Predicate_chatInvite:
		t := m.To_ChatInvite()
		return t.DebugString()
	case Predicate_chatInvitePeek:
		t := m.To_ChatInvitePeek()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ChatInviteAlready
func (m *ChatInviteExt) To_ChatInviteAlready() *TLChatInviteAlready {
	m.PredicateName = Predicate_chatInviteAlready
	return &TLChatInviteAlready{
		Data2: m,
	}
}

// To_ChatInvite
func (m *ChatInviteExt) To_ChatInvite() *TLChatInvite {
	m.PredicateName = Predicate_chatInvite
	return &TLChatInvite{
		Data2: m,
	}
}

// To_ChatInvitePeek
func (m *ChatInviteExt) To_ChatInvitePeek() *TLChatInvitePeek {
	m.PredicateName = Predicate_chatInvitePeek
	return &TLChatInvitePeek{
		Data2: m,
	}
}

// MakeTLChatInviteAlready
func MakeTLChatInviteAlready(data2 *ChatInviteExt) *TLChatInviteAlready {
	if data2 == nil {
		return &TLChatInviteAlready{Data2: &ChatInviteExt{
			PredicateName: Predicate_chatInviteAlready,
		}}
	} else {
		data2.PredicateName = Predicate_chatInviteAlready
		return &TLChatInviteAlready{Data2: data2}
	}
}

func (m *TLChatInviteAlready) To_ChatInviteExt() *ChatInviteExt {
	m.Data2.PredicateName = Predicate_chatInviteAlready
	return m.Data2
}

func (m *TLChatInviteAlready) SetChat(v *mtproto.MutableChat) { m.Data2.Chat = v }
func (m *TLChatInviteAlready) GetChat() *mtproto.MutableChat  { return m.Data2.Chat }

func (m *TLChatInviteAlready) GetPredicateName() string {
	return Predicate_chatInviteAlready
}

func (m *TLChatInviteAlready) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xa40e7d5e: func() []byte {
			x.UInt(0xa40e7d5e)

			x.Bytes(m.GetChat().Encode(layer))
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_chatInviteAlready, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_chatInviteAlready, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLChatInviteAlready) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatInviteAlready) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xa40e7d5e: func() error {

			m0 := &mtproto.MutableChat{}
			m0.Decode(dBuf)
			m.SetChat(m0)

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLChatInviteAlready) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLChatInvite
func MakeTLChatInvite(data2 *ChatInviteExt) *TLChatInvite {
	if data2 == nil {
		return &TLChatInvite{Data2: &ChatInviteExt{
			PredicateName: Predicate_chatInvite,
		}}
	} else {
		data2.PredicateName = Predicate_chatInvite
		return &TLChatInvite{Data2: data2}
	}
}

func (m *TLChatInvite) To_ChatInviteExt() *ChatInviteExt {
	m.Data2.PredicateName = Predicate_chatInvite
	return m.Data2
}

//// flags
func (m *TLChatInvite) SetRequestNeeded(v bool) { m.Data2.RequestNeeded = v }
func (m *TLChatInvite) GetRequestNeeded() bool  { return m.Data2.RequestNeeded }

func (m *TLChatInvite) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChatInvite) GetTitle() string  { return m.Data2.Title }

func (m *TLChatInvite) SetAbout(v *types.StringValue) { m.Data2.About = v }
func (m *TLChatInvite) GetAbout() *types.StringValue  { return m.Data2.About }

func (m *TLChatInvite) SetPhoto(v *mtproto.Photo) { m.Data2.Photo = v }
func (m *TLChatInvite) GetPhoto() *mtproto.Photo  { return m.Data2.Photo }

func (m *TLChatInvite) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChatInvite) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChatInvite) SetParticipants(v []int64) { m.Data2.Participants = v }
func (m *TLChatInvite) GetParticipants() []int64  { return m.Data2.Participants }

func (m *TLChatInvite) GetPredicateName() string {
	return Predicate_chatInvite
}

func (m *TLChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xdb75d1a7: func() []byte {
			x.UInt(0xdb75d1a7)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetRequestNeeded() == true {
					flags |= 1 << 6
				}

				if m.GetAbout() != nil {
					flags |= 1 << 5
				}

				if m.GetParticipants() != nil {
					flags |= 1 << 4
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.String(m.GetTitle())
			if m.GetAbout() != nil {
				x.String(m.GetAbout().Value)
			}

			x.Bytes(m.GetPhoto().Encode(layer))
			x.Int(m.GetParticipantsCount())
			if m.GetParticipants() != nil {
				x.VectorLong(m.GetParticipants())
			}
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_chatInvite, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_chatInvite, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xdb75d1a7: func() error {
			var flags = dBuf.UInt()
			_ = flags
			if (flags & (1 << 6)) != 0 {
				m.SetRequestNeeded(true)
			}
			m.SetTitle(dBuf.String())
			if (flags & (1 << 5)) != 0 {
				m.SetAbout(&types.StringValue{Value: dBuf.String()})
			}

			m5 := &mtproto.Photo{}
			m5.Decode(dBuf)
			m.SetPhoto(m5)

			m.SetParticipantsCount(dBuf.Int())
			if (flags & (1 << 4)) != 0 {
				m.SetParticipants(dBuf.VectorLong())
			}
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLChatInvitePeek
func MakeTLChatInvitePeek(data2 *ChatInviteExt) *TLChatInvitePeek {
	if data2 == nil {
		return &TLChatInvitePeek{Data2: &ChatInviteExt{
			PredicateName: Predicate_chatInvitePeek,
		}}
	} else {
		data2.PredicateName = Predicate_chatInvitePeek
		return &TLChatInvitePeek{Data2: data2}
	}
}

func (m *TLChatInvitePeek) To_ChatInviteExt() *ChatInviteExt {
	m.Data2.PredicateName = Predicate_chatInvitePeek
	return m.Data2
}

func (m *TLChatInvitePeek) SetChat(v *mtproto.MutableChat) { m.Data2.Chat = v }
func (m *TLChatInvitePeek) GetChat() *mtproto.MutableChat  { return m.Data2.Chat }

func (m *TLChatInvitePeek) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLChatInvitePeek) GetExpires() int32  { return m.Data2.Expires }

func (m *TLChatInvitePeek) GetPredicateName() string {
	return Predicate_chatInvitePeek
}

func (m *TLChatInvitePeek) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xace3e26e: func() []byte {
			x.UInt(0xace3e26e)

			x.Bytes(m.GetChat().Encode(layer))
			x.Int(m.GetExpires())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_chatInvitePeek, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_chatInvitePeek, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLChatInvitePeek) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatInvitePeek) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xace3e26e: func() error {

			m0 := &mtproto.MutableChat{}
			m0.Decode(dBuf)
			m.SetChat(m0)

			m.SetExpires(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLChatInvitePeek) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// ChatInviteImported <--
//  + TL_ChatInviteImported
//

func (m *ChatInviteImported) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_chatInviteImported:
		t := m.To_ChatInviteImported()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *ChatInviteImported) CalcByteSize(layer int32) int {
	return 0
}

func (m *ChatInviteImported) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x721051f6:
		m2 := MakeTLChatInviteImported(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *ChatInviteImported) DebugString() string {
	switch m.PredicateName {
	case Predicate_chatInviteImported:
		t := m.To_ChatInviteImported()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_ChatInviteImported
func (m *ChatInviteImported) To_ChatInviteImported() *TLChatInviteImported {
	m.PredicateName = Predicate_chatInviteImported
	return &TLChatInviteImported{
		Data2: m,
	}
}

// MakeTLChatInviteImported
func MakeTLChatInviteImported(data2 *ChatInviteImported) *TLChatInviteImported {
	if data2 == nil {
		return &TLChatInviteImported{Data2: &ChatInviteImported{
			PredicateName: Predicate_chatInviteImported,
		}}
	} else {
		data2.PredicateName = Predicate_chatInviteImported
		return &TLChatInviteImported{Data2: data2}
	}
}

func (m *TLChatInviteImported) To_ChatInviteImported() *ChatInviteImported {
	m.Data2.PredicateName = Predicate_chatInviteImported
	return m.Data2
}

//// flags
func (m *TLChatInviteImported) SetChat(v *mtproto.MutableChat) { m.Data2.Chat = v }
func (m *TLChatInviteImported) GetChat() *mtproto.MutableChat  { return m.Data2.Chat }

func (m *TLChatInviteImported) SetRequesters(v *RecentChatInviteRequesters) { m.Data2.Requesters = v }
func (m *TLChatInviteImported) GetRequesters() *RecentChatInviteRequesters  { return m.Data2.Requesters }

func (m *TLChatInviteImported) GetPredicateName() string {
	return Predicate_chatInviteImported
}

func (m *TLChatInviteImported) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x721051f6: func() []byte {
			x.UInt(0x721051f6)

			// set flags
			var getFlags = func() uint32 {
				var flags uint32 = 0

				if m.GetRequesters() != nil {
					flags |= 1 << 0
				}

				return flags
			}

			// set flags
			var flags = getFlags()
			x.UInt(flags)
			x.Bytes(m.GetChat().Encode(layer))
			if m.GetRequesters() != nil {
				x.Bytes(m.GetRequesters().Encode(layer))
			}

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_chatInviteImported, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_chatInviteImported, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLChatInviteImported) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatInviteImported) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x721051f6: func() error {
			var flags = dBuf.UInt()
			_ = flags

			m1 := &mtproto.MutableChat{}
			m1.Decode(dBuf)
			m.SetChat(m1)

			if (flags & (1 << 0)) != 0 {
				m2 := &RecentChatInviteRequesters{}
				m2.Decode(dBuf)
				m.SetRequesters(m2)
			}
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLChatInviteImported) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// RecentChatInviteRequesters <--
//  + TL_RecentChatInviteRequesters
//

func (m *RecentChatInviteRequesters) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_recentChatInviteRequesters:
		t := m.To_RecentChatInviteRequesters()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *RecentChatInviteRequesters) CalcByteSize(layer int32) int {
	return 0
}

func (m *RecentChatInviteRequesters) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x1c6e3c54:
		m2 := MakeTLRecentChatInviteRequesters(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *RecentChatInviteRequesters) DebugString() string {
	switch m.PredicateName {
	case Predicate_recentChatInviteRequesters:
		t := m.To_RecentChatInviteRequesters()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_RecentChatInviteRequesters
func (m *RecentChatInviteRequesters) To_RecentChatInviteRequesters() *TLRecentChatInviteRequesters {
	m.PredicateName = Predicate_recentChatInviteRequesters
	return &TLRecentChatInviteRequesters{
		Data2: m,
	}
}

// MakeTLRecentChatInviteRequesters
func MakeTLRecentChatInviteRequesters(data2 *RecentChatInviteRequesters) *TLRecentChatInviteRequesters {
	if data2 == nil {
		return &TLRecentChatInviteRequesters{Data2: &RecentChatInviteRequesters{
			PredicateName: Predicate_recentChatInviteRequesters,
		}}
	} else {
		data2.PredicateName = Predicate_recentChatInviteRequesters
		return &TLRecentChatInviteRequesters{Data2: data2}
	}
}

func (m *TLRecentChatInviteRequesters) To_RecentChatInviteRequesters() *RecentChatInviteRequesters {
	m.Data2.PredicateName = Predicate_recentChatInviteRequesters
	return m.Data2
}

func (m *TLRecentChatInviteRequesters) SetRequestsPending(v int32) { m.Data2.RequestsPending = v }
func (m *TLRecentChatInviteRequesters) GetRequestsPending() int32  { return m.Data2.RequestsPending }

func (m *TLRecentChatInviteRequesters) SetRecentRequesters(v []int64) { m.Data2.RecentRequesters = v }
func (m *TLRecentChatInviteRequesters) GetRecentRequesters() []int64  { return m.Data2.RecentRequesters }

func (m *TLRecentChatInviteRequesters) GetPredicateName() string {
	return Predicate_recentChatInviteRequesters
}

func (m *TLRecentChatInviteRequesters) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x1c6e3c54: func() []byte {
			x.UInt(0x1c6e3c54)

			x.Int(m.GetRequestsPending())

			x.VectorLong(m.GetRecentRequesters())

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_recentChatInviteRequesters, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_recentChatInviteRequesters, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLRecentChatInviteRequesters) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLRecentChatInviteRequesters) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x1c6e3c54: func() error {
			m.SetRequestsPending(dBuf.Int())

			m.SetRecentRequesters(dBuf.VectorLong())

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLRecentChatInviteRequesters) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// UserChatIdList <--
//  + TL_UserChatIdList
//

func (m *UserChatIdList) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_userChatIdList:
		t := m.To_UserChatIdList()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *UserChatIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *UserChatIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x50067224:
		m2 := MakeTLUserChatIdList(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *UserChatIdList) DebugString() string {
	switch m.PredicateName {
	case Predicate_userChatIdList:
		t := m.To_UserChatIdList()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_UserChatIdList
func (m *UserChatIdList) To_UserChatIdList() *TLUserChatIdList {
	m.PredicateName = Predicate_userChatIdList
	return &TLUserChatIdList{
		Data2: m,
	}
}

// MakeTLUserChatIdList
func MakeTLUserChatIdList(data2 *UserChatIdList) *TLUserChatIdList {
	if data2 == nil {
		return &TLUserChatIdList{Data2: &UserChatIdList{
			PredicateName: Predicate_userChatIdList,
		}}
	} else {
		data2.PredicateName = Predicate_userChatIdList
		return &TLUserChatIdList{Data2: data2}
	}
}

func (m *TLUserChatIdList) To_UserChatIdList() *UserChatIdList {
	m.Data2.PredicateName = Predicate_userChatIdList
	return m.Data2
}

func (m *TLUserChatIdList) SetUserId(v int64) { m.Data2.UserId = v }
func (m *TLUserChatIdList) GetUserId() int64  { return m.Data2.UserId }

func (m *TLUserChatIdList) SetChatIdList(v []int64) { m.Data2.ChatIdList = v }
func (m *TLUserChatIdList) GetChatIdList() []int64  { return m.Data2.ChatIdList }

func (m *TLUserChatIdList) GetPredicateName() string {
	return Predicate_userChatIdList
}

func (m *TLUserChatIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x50067224: func() []byte {
			x.UInt(0x50067224)

			x.Long(m.GetUserId())

			x.VectorLong(m.GetChatIdList())

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_userChatIdList, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_userChatIdList, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLUserChatIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLUserChatIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x50067224: func() error {
			m.SetUserId(dBuf.Long())

			m.SetChatIdList(dBuf.VectorLong())

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLUserChatIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLChatGetMutableChat
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetMutableChat) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getMutableChat))

	switch uint32(m.Constructor) {
	case 0x2c2c25d2:
		x.UInt(0x2c2c25d2)

		// no flags

		x.Long(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetMutableChat) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetMutableChat) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2c2c25d2:

		// not has flags

		m.ChatId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetMutableChat) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetChatListByIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetChatListByIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getChatListByIdList))

	switch uint32(m.Constructor) {
	case 0xe740f539:
		x.UInt(0xe740f539)

		// no flags

		x.Long(m.GetSelfId())

		x.VectorLong(m.GetIdList())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetChatListByIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetChatListByIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe740f539:

		// not has flags

		m.SelfId = dBuf.Long()

		m.IdList = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetChatListByIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetChatBySelfId
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetChatBySelfId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getChatBySelfId))

	switch uint32(m.Constructor) {
	case 0x49b71a48:
		x.UInt(0x49b71a48)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetChatBySelfId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetChatBySelfId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x49b71a48:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetChatBySelfId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatCreateChat2
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatCreateChat2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_createChat2))

	switch uint32(m.Constructor) {
	case 0xe43f61af:
		x.UInt(0xe43f61af)

		// no flags

		x.Long(m.GetCreatorId())

		x.VectorLong(m.GetUserIdList())

		x.String(m.GetTitle())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatCreateChat2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatCreateChat2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xe43f61af:

		// not has flags

		m.CreatorId = dBuf.Long()

		m.UserIdList = dBuf.VectorLong()

		m.Title = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatCreateChat2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatDeleteChat
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatDeleteChat) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_deleteChat))

	switch uint32(m.Constructor) {
	case 0x6d11ec1e:
		x.UInt(0x6d11ec1e)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetOperatorId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatDeleteChat) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatDeleteChat) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x6d11ec1e:

		// not has flags

		m.ChatId = dBuf.Long()
		m.OperatorId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatDeleteChat) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatDeleteChatUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatDeleteChatUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_deleteChatUser))

	switch uint32(m.Constructor) {
	case 0xb270fd5:
		x.UInt(0xb270fd5)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetOperatorId())
		x.Long(m.GetDeleteUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatDeleteChatUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatDeleteChatUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb270fd5:

		// not has flags

		m.ChatId = dBuf.Long()
		m.OperatorId = dBuf.Long()
		m.DeleteUserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatDeleteChatUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditChatTitle
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditChatTitle) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editChatTitle))

	switch uint32(m.Constructor) {
	case 0x95c59ea7:
		x.UInt(0x95c59ea7)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetEditUserId())
		x.String(m.GetTitle())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditChatTitle) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditChatTitle) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x95c59ea7:

		// not has flags

		m.ChatId = dBuf.Long()
		m.EditUserId = dBuf.Long()
		m.Title = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditChatTitle) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditChatAbout
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditChatAbout) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editChatAbout))

	switch uint32(m.Constructor) {
	case 0x5c737c78:
		x.UInt(0x5c737c78)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetEditUserId())
		x.String(m.GetAbout())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditChatAbout) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditChatAbout) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5c737c78:

		// not has flags

		m.ChatId = dBuf.Long()
		m.EditUserId = dBuf.Long()
		m.About = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditChatAbout) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditChatPhoto
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditChatPhoto) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editChatPhoto))

	switch uint32(m.Constructor) {
	case 0x45c2a668:
		x.UInt(0x45c2a668)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetEditUserId())
		x.Bytes(m.GetChatPhoto().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditChatPhoto) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditChatPhoto) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x45c2a668:

		// not has flags

		m.ChatId = dBuf.Long()
		m.EditUserId = dBuf.Long()

		m3 := &mtproto.Photo{}
		m3.Decode(dBuf)
		m.ChatPhoto = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditChatPhoto) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditChatAdmin
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditChatAdmin) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editChatAdmin))

	switch uint32(m.Constructor) {
	case 0x1905e5ec:
		x.UInt(0x1905e5ec)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetOperatorId())
		x.Long(m.GetEditChatAdminId())
		x.Bytes(m.GetIsAdmin().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditChatAdmin) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditChatAdmin) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x1905e5ec:

		// not has flags

		m.ChatId = dBuf.Long()
		m.OperatorId = dBuf.Long()
		m.EditChatAdminId = dBuf.Long()

		m4 := &mtproto.Bool{}
		m4.Decode(dBuf)
		m.IsAdmin = m4

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditChatAdmin) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditChatDefaultBannedRights
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditChatDefaultBannedRights) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editChatDefaultBannedRights))

	switch uint32(m.Constructor) {
	case 0x5a34a687:
		x.UInt(0x5a34a687)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetOperatorId())
		x.Bytes(m.GetBannedRights().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditChatDefaultBannedRights) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditChatDefaultBannedRights) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x5a34a687:

		// not has flags

		m.ChatId = dBuf.Long()
		m.OperatorId = dBuf.Long()

		m3 := &mtproto.ChatBannedRights{}
		m3.Decode(dBuf)
		m.BannedRights = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditChatDefaultBannedRights) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatAddChatUser
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatAddChatUser) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_addChatUser))

	switch uint32(m.Constructor) {
	case 0x7c5da48d:
		x.UInt(0x7c5da48d)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetInviterId())
		x.Long(m.GetUserId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatAddChatUser) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatAddChatUser) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7c5da48d:

		// not has flags

		m.ChatId = dBuf.Long()
		m.InviterId = dBuf.Long()
		m.UserId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatAddChatUser) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetMutableChatByLink
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetMutableChatByLink) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getMutableChatByLink))

	switch uint32(m.Constructor) {
	case 0xa266278b:
		x.UInt(0xa266278b)

		// no flags

		x.String(m.GetLink())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetMutableChatByLink) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetMutableChatByLink) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa266278b:

		// not has flags

		m.Link = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetMutableChatByLink) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatToggleNoForwards
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatToggleNoForwards) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_toggleNoForwards))

	switch uint32(m.Constructor) {
	case 0xd5952af9:
		x.UInt(0xd5952af9)

		// no flags

		x.Long(m.GetChatId())
		x.Long(m.GetOperatorId())
		x.Bytes(m.GetEnabled().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatToggleNoForwards) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatToggleNoForwards) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd5952af9:

		// not has flags

		m.ChatId = dBuf.Long()
		m.OperatorId = dBuf.Long()

		m3 := &mtproto.Bool{}
		m3.Decode(dBuf)
		m.Enabled = m3

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatToggleNoForwards) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatMigratedToChannel
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatMigratedToChannel) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_migratedToChannel))

	switch uint32(m.Constructor) {
	case 0x83faadf:
		x.UInt(0x83faadf)

		// no flags

		x.Bytes(m.GetChat().Encode(layer))
		x.Long(m.GetId())
		x.Long(m.GetAccessHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatMigratedToChannel) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatMigratedToChannel) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x83faadf:

		// not has flags

		m1 := &mtproto.MutableChat{}
		m1.Decode(dBuf)
		m.Chat = m1

		m.Id = dBuf.Long()
		m.AccessHash = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatMigratedToChannel) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetChatParticipantIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetChatParticipantIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getChatParticipantIdList))

	switch uint32(m.Constructor) {
	case 0x329622a9:
		x.UInt(0x329622a9)

		// no flags

		x.Long(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetChatParticipantIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetChatParticipantIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x329622a9:

		// not has flags

		m.ChatId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetChatParticipantIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetUsersChatIdList
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetUsersChatIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getUsersChatIdList))

	switch uint32(m.Constructor) {
	case 0x2f36ab4c:
		x.UInt(0x2f36ab4c)

		// no flags

		x.VectorLong(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetUsersChatIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetUsersChatIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x2f36ab4c:

		// not has flags

		m.Id = dBuf.VectorLong()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetUsersChatIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetMyChatList
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetMyChatList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getMyChatList))

	switch uint32(m.Constructor) {
	case 0xf3756c88:
		x.UInt(0xf3756c88)

		// no flags

		x.Long(m.GetUserId())
		x.Bytes(m.GetIsCreator().Encode(layer))

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetMyChatList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetMyChatList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf3756c88:

		// not has flags

		m.UserId = dBuf.Long()

		m2 := &mtproto.Bool{}
		m2.Decode(dBuf)
		m.IsCreator = m2

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetMyChatList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatExportChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatExportChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_exportChatInvite))

	switch uint32(m.Constructor) {
	case 0xc5cf804b:
		x.UInt(0xc5cf804b)

		// set flags
		var flags uint32 = 0

		if m.GetLegacyRevokePermanent() == true {
			flags |= 1 << 2
		}
		if m.GetRequestNeeded() == true {
			flags |= 1 << 3
		}
		if m.GetExpireDate() != nil {
			flags |= 1 << 0
		}
		if m.GetUsageLimit() != nil {
			flags |= 1 << 1
		}
		if m.GetTitle() != nil {
			flags |= 1 << 4
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetChatId())
		x.Long(m.GetAdminId())
		if m.GetExpireDate() != nil {
			x.Int(m.GetExpireDate().Value)
		}

		if m.GetUsageLimit() != nil {
			x.Int(m.GetUsageLimit().Value)
		}

		if m.GetTitle() != nil {
			x.String(m.GetTitle().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatExportChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatExportChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc5cf804b:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.ChatId = dBuf.Long()
		m.AdminId = dBuf.Long()
		if (flags & (1 << 2)) != 0 {
			m.LegacyRevokePermanent = true
		}
		if (flags & (1 << 3)) != 0 {
			m.RequestNeeded = true
		}
		if (flags & (1 << 0)) != 0 {
			m.ExpireDate = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 1)) != 0 {
			m.UsageLimit = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 4)) != 0 {
			m.Title = &types.StringValue{Value: dBuf.String()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatExportChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetAdminsWithInvites
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetAdminsWithInvites) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getAdminsWithInvites))

	switch uint32(m.Constructor) {
	case 0xd2ea41d2:
		x.UInt(0xd2ea41d2)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetAdminsWithInvites) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetAdminsWithInvites) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd2ea41d2:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetAdminsWithInvites) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetExportedChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetExportedChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getExportedChatInvite))

	switch uint32(m.Constructor) {
	case 0xddea3250:
		x.UInt(0xddea3250)

		// no flags

		x.Long(m.GetChatId())
		x.String(m.GetLink())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetExportedChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetExportedChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xddea3250:

		// not has flags

		m.ChatId = dBuf.Long()
		m.Link = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetExportedChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetExportedChatInvites
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetExportedChatInvites) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getExportedChatInvites))

	switch uint32(m.Constructor) {
	case 0xb48f18f6:
		x.UInt(0xb48f18f6)

		// set flags
		var flags uint32 = 0

		if m.GetRevoked() == true {
			flags |= 1 << 3
		}
		if m.GetOffsetDate() != nil {
			flags |= 1 << 2
		}
		if m.GetOffsetLink() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetChatId())
		x.Long(m.GetAdminId())
		if m.GetOffsetDate() != nil {
			x.Int(m.GetOffsetDate().Value)
		}

		if m.GetOffsetLink() != nil {
			x.String(m.GetOffsetLink().Value)
		}

		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetExportedChatInvites) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetExportedChatInvites) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xb48f18f6:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.ChatId = dBuf.Long()
		m.AdminId = dBuf.Long()
		if (flags & (1 << 3)) != 0 {
			m.Revoked = true
		}
		if (flags & (1 << 2)) != 0 {
			m.OffsetDate = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 2)) != 0 {
			m.OffsetLink = &types.StringValue{Value: dBuf.String()}
		}

		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetExportedChatInvites) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatCheckChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatCheckChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_checkChatInvite))

	switch uint32(m.Constructor) {
	case 0x7387f28c:
		x.UInt(0x7387f28c)

		// no flags

		x.Long(m.GetSelfId())
		x.String(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatCheckChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatCheckChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x7387f28c:

		// not has flags

		m.SelfId = dBuf.Long()
		m.Hash = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatCheckChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatImportChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatImportChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_importChatInvite))

	switch uint32(m.Constructor) {
	case 0x58e660d4:
		x.UInt(0x58e660d4)

		// no flags

		x.Long(m.GetSelfId())
		x.String(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatImportChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatImportChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x58e660d4:

		// not has flags

		m.SelfId = dBuf.Long()
		m.Hash = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatImportChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetChatInviteImporters
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetChatInviteImporters) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getChatInviteImporters))

	switch uint32(m.Constructor) {
	case 0x9846557f:
		x.UInt(0x9846557f)

		// set flags
		var flags uint32 = 0

		if m.GetRequested() == true {
			flags |= 1 << 0
		}
		if m.GetLink() != nil {
			flags |= 1 << 1
		}
		if m.GetQ() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		if m.GetLink() != nil {
			x.String(m.GetLink().Value)
		}

		if m.GetQ() != nil {
			x.String(m.GetQ().Value)
		}

		x.Int(m.GetOffsetDate())
		x.Long(m.GetOffsetUser())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetChatInviteImporters) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetChatInviteImporters) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9846557f:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Requested = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Link = &types.StringValue{Value: dBuf.String()}
		}

		if (flags & (1 << 2)) != 0 {
			m.Q = &types.StringValue{Value: dBuf.String()}
		}

		m.OffsetDate = dBuf.Int()
		m.OffsetUser = dBuf.Long()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetChatInviteImporters) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatDeleteExportedChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatDeleteExportedChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_deleteExportedChatInvite))

	switch uint32(m.Constructor) {
	case 0x562288b8:
		x.UInt(0x562288b8)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		x.String(m.GetLink())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatDeleteExportedChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatDeleteExportedChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x562288b8:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		m.Link = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatDeleteExportedChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatDeleteRevokedExportedChatInvites
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatDeleteRevokedExportedChatInvites) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_deleteRevokedExportedChatInvites))

	switch uint32(m.Constructor) {
	case 0xd0126269:
		x.UInt(0xd0126269)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		x.Long(m.GetAdminId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatDeleteRevokedExportedChatInvites) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatDeleteRevokedExportedChatInvites) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xd0126269:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		m.AdminId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatDeleteRevokedExportedChatInvites) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatEditExportedChatInvite
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatEditExportedChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_editExportedChatInvite))

	switch uint32(m.Constructor) {
	case 0xaf994c76:
		x.UInt(0xaf994c76)

		// set flags
		var flags uint32 = 0

		if m.GetRevoked() == true {
			flags |= 1 << 2
		}

		if m.GetExpireDate() != nil {
			flags |= 1 << 0
		}
		if m.GetUsageLimit() != nil {
			flags |= 1 << 1
		}
		if m.GetRequestNeeded() != nil {
			flags |= 1 << 3
		}
		if m.GetTitle() != nil {
			flags |= 1 << 4
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		x.String(m.GetLink())
		if m.GetExpireDate() != nil {
			x.Int(m.GetExpireDate().Value)
		}

		if m.GetUsageLimit() != nil {
			x.Int(m.GetUsageLimit().Value)
		}

		if m.GetRequestNeeded() != nil {
			x.Bytes(m.GetRequestNeeded().Encode(layer))
		}

		if m.GetTitle() != nil {
			x.String(m.GetTitle().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatEditExportedChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatEditExportedChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xaf994c76:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		if (flags & (1 << 2)) != 0 {
			m.Revoked = true
		}
		m.Link = dBuf.String()
		if (flags & (1 << 0)) != 0 {
			m.ExpireDate = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 1)) != 0 {
			m.UsageLimit = &types.Int32Value{Value: dBuf.Int()}
		}

		if (flags & (1 << 3)) != 0 {
			m8 := &mtproto.Bool{}
			m8.Decode(dBuf)
			m.RequestNeeded = m8
		}
		if (flags & (1 << 4)) != 0 {
			m.Title = &types.StringValue{Value: dBuf.String()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatEditExportedChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatSetChatAvailableReactions
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatSetChatAvailableReactions) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_setChatAvailableReactions))

	switch uint32(m.Constructor) {
	case 0xc4d08972:
		x.UInt(0xc4d08972)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		x.Int(m.GetAvailableReactionsType())

		x.VectorString(m.GetAvailableReactions())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatSetChatAvailableReactions) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatSetChatAvailableReactions) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xc4d08972:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		m.AvailableReactionsType = dBuf.Int()

		m.AvailableReactions = dBuf.VectorString()

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatSetChatAvailableReactions) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatSetHistoryTTL
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatSetHistoryTTL) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_setHistoryTTL))

	switch uint32(m.Constructor) {
	case 0x3cfb6384:
		x.UInt(0x3cfb6384)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		x.Int(m.GetTtlPeriod())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatSetHistoryTTL) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatSetHistoryTTL) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3cfb6384:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		m.TtlPeriod = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatSetHistoryTTL) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatSearch
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatSearch) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_search))

	switch uint32(m.Constructor) {
	case 0x21e014fb:
		x.UInt(0x21e014fb)

		// no flags

		x.Long(m.GetSelfId())
		x.String(m.GetQ())
		x.Long(m.GetOffset())
		x.Int(m.GetLimit())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatSearch) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatSearch) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x21e014fb:

		// not has flags

		m.SelfId = dBuf.Long()
		m.Q = dBuf.String()
		m.Offset = dBuf.Long()
		m.Limit = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatSearch) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatGetRecentChatInviteRequesters
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatGetRecentChatInviteRequesters) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_getRecentChatInviteRequesters))

	switch uint32(m.Constructor) {
	case 0xfedc1098:
		x.UInt(0xfedc1098)

		// no flags

		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatGetRecentChatInviteRequesters) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatGetRecentChatInviteRequesters) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xfedc1098:

		// not has flags

		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatGetRecentChatInviteRequesters) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatHideChatJoinRequests
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatHideChatJoinRequests) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_hideChatJoinRequests))

	switch uint32(m.Constructor) {
	case 0x3ea52cd1:
		x.UInt(0x3ea52cd1)

		// set flags
		var flags uint32 = 0

		if m.GetApproved() == true {
			flags |= 1 << 0
		}
		if m.GetLink() != nil {
			flags |= 1 << 1
		}
		if m.GetUserId() != nil {
			flags |= 1 << 2
		}

		x.UInt(flags)

		// flags Debug by @benqi
		x.Long(m.GetSelfId())
		x.Long(m.GetChatId())
		if m.GetLink() != nil {
			x.String(m.GetLink().Value)
		}

		if m.GetUserId() != nil {
			x.Long(m.GetUserId().Value)
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatHideChatJoinRequests) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatHideChatJoinRequests) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x3ea52cd1:

		flags := dBuf.UInt()
		_ = flags

		// flags Debug by @benqi
		m.SelfId = dBuf.Long()
		m.ChatId = dBuf.Long()
		if (flags & (1 << 0)) != 0 {
			m.Approved = true
		}
		if (flags & (1 << 1)) != 0 {
			m.Link = &types.StringValue{Value: dBuf.String()}
		}

		if (flags & (1 << 2)) != 0 {
			m.UserId = &types.Int64Value{Value: dBuf.Long()}
		}

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatHideChatJoinRequests) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLChatImportChatInvite2
///////////////////////////////////////////////////////////////////////////////

func (m *TLChatImportChatInvite2) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_chat_importChatInvite2))

	switch uint32(m.Constructor) {
	case 0xdcd93dbf:
		x.UInt(0xdcd93dbf)

		// no flags

		x.Long(m.GetSelfId())
		x.String(m.GetHash())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLChatImportChatInvite2) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLChatImportChatInvite2) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xdcd93dbf:

		// not has flags

		m.SelfId = dBuf.Long()
		m.Hash = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLChatImportChatInvite2) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_MutableChat
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_MutableChat) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_MutableChat) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.MutableChat, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.MutableChat)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_MutableChat) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_MutableChat) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_Long
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorLong(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Long) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_UserChatIdList
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_UserChatIdList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_UserChatIdList) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*UserChatIdList, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(UserChatIdList)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_UserChatIdList) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_UserChatIdList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_ChatAdminWithInvites
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_ChatAdminWithInvites) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_ChatAdminWithInvites) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ChatAdminWithInvites, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ChatAdminWithInvites)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ChatAdminWithInvites) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_ChatAdminWithInvites) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_ExportedChatInvite
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_ExportedChatInvite) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_ExportedChatInvite) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ExportedChatInvite, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ExportedChatInvite)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ExportedChatInvite) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_ExportedChatInvite) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_ChatInviteImporter
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_ChatInviteImporter) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_ChatInviteImporter) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*mtproto.ChatInviteImporter, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(mtproto.ChatInviteImporter)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_ChatInviteImporter) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_ChatInviteImporter) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
