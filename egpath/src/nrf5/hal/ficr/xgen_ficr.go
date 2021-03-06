package ficr

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"nrf5/hal/internal/mmap"
)

type Periph struct {
	_              [4]uint32
	CODEPAGESIZE   CODEPAGESIZE
	CODESIZE       CODESIZE
	_              [4]uint32
	CLENR0         CLENR0
	PPFC           PPFC
	_              uint32
	NUMRAMBLOCK    NUMRAMBLOCK
	SIZERAMBLOCK   [4]SIZERAMBLOCK
	_              [5]uint32
	CONFIGID       CONFIGID
	DEVICEID       [2]DEVICEID
	_              [6]uint32
	ER             [4]ER
	IR             [4]IR
	DEVICEADDRTYPE DEVICEADDRTYPE
	DEVICEADDR     [2]DEVICEADDR
	OVERRIDDEN     OVERRIDDEN
	NRF_1MBIT      [5]NRF_1MBIT
	_              [10]uint32
	BLE_1MBIT      [5]BLE_1MBIT
	INFO_PART      INFO_PART
	INFO_VARIANT   INFO_VARIANT
	INFO_PACKAGE   INFO_PACKAGE
	INFO_RAM       INFO_RAM
	INFO_FLASH     INFO_FLASH
	_              [188]uint32
	TEMP_A         [6]TEMP_A
	TEMP_B         [6]TEMP_B
	TEMP_T         [5]TEMP_T
	_              [2]uint32
	NFC_TAGHEADER  [4]NFC_TAGHEADER
}

func (p *Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FICR = (*Periph)(unsafe.Pointer(uintptr(mmap.FICR_BASE)))

type CODEPAGESIZE_Bits uint32

func (b CODEPAGESIZE_Bits) Field(mask CODEPAGESIZE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CODEPAGESIZE_Bits) J(v int) CODEPAGESIZE_Bits {
	return CODEPAGESIZE_Bits(bits.Make32(v, uint32(mask)))
}

type CODEPAGESIZE struct{ mmio.U32 }

func (r *CODEPAGESIZE) Bits(mask CODEPAGESIZE_Bits) CODEPAGESIZE_Bits {
	return CODEPAGESIZE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CODEPAGESIZE) StoreBits(mask, b CODEPAGESIZE_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CODEPAGESIZE) SetBits(mask CODEPAGESIZE_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CODEPAGESIZE) ClearBits(mask CODEPAGESIZE_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CODEPAGESIZE) Load() CODEPAGESIZE_Bits             { return CODEPAGESIZE_Bits(r.U32.Load()) }
func (r *CODEPAGESIZE) Store(b CODEPAGESIZE_Bits)           { r.U32.Store(uint32(b)) }

func (r *CODEPAGESIZE) AtomicStoreBits(mask, b CODEPAGESIZE_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *CODEPAGESIZE) AtomicSetBits(mask CODEPAGESIZE_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CODEPAGESIZE) AtomicClearBits(mask CODEPAGESIZE_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CODEPAGESIZE_Mask struct{ mmio.UM32 }

func (rm CODEPAGESIZE_Mask) Load() CODEPAGESIZE_Bits   { return CODEPAGESIZE_Bits(rm.UM32.Load()) }
func (rm CODEPAGESIZE_Mask) Store(b CODEPAGESIZE_Bits) { rm.UM32.Store(uint32(b)) }

type CODESIZE_Bits uint32

func (b CODESIZE_Bits) Field(mask CODESIZE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CODESIZE_Bits) J(v int) CODESIZE_Bits {
	return CODESIZE_Bits(bits.Make32(v, uint32(mask)))
}

type CODESIZE struct{ mmio.U32 }

func (r *CODESIZE) Bits(mask CODESIZE_Bits) CODESIZE_Bits {
	return CODESIZE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CODESIZE) StoreBits(mask, b CODESIZE_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CODESIZE) SetBits(mask CODESIZE_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CODESIZE) ClearBits(mask CODESIZE_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CODESIZE) Load() CODESIZE_Bits             { return CODESIZE_Bits(r.U32.Load()) }
func (r *CODESIZE) Store(b CODESIZE_Bits)           { r.U32.Store(uint32(b)) }

func (r *CODESIZE) AtomicStoreBits(mask, b CODESIZE_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *CODESIZE) AtomicSetBits(mask CODESIZE_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CODESIZE) AtomicClearBits(mask CODESIZE_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CODESIZE_Mask struct{ mmio.UM32 }

func (rm CODESIZE_Mask) Load() CODESIZE_Bits   { return CODESIZE_Bits(rm.UM32.Load()) }
func (rm CODESIZE_Mask) Store(b CODESIZE_Bits) { rm.UM32.Store(uint32(b)) }

type CLENR0_Bits uint32

func (b CLENR0_Bits) Field(mask CLENR0_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLENR0_Bits) J(v int) CLENR0_Bits {
	return CLENR0_Bits(bits.Make32(v, uint32(mask)))
}

type CLENR0 struct{ mmio.U32 }

func (r *CLENR0) Bits(mask CLENR0_Bits) CLENR0_Bits { return CLENR0_Bits(r.U32.Bits(uint32(mask))) }
func (r *CLENR0) StoreBits(mask, b CLENR0_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CLENR0) SetBits(mask CLENR0_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CLENR0) ClearBits(mask CLENR0_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CLENR0) Load() CLENR0_Bits                 { return CLENR0_Bits(r.U32.Load()) }
func (r *CLENR0) Store(b CLENR0_Bits)               { r.U32.Store(uint32(b)) }

func (r *CLENR0) AtomicStoreBits(mask, b CLENR0_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CLENR0) AtomicSetBits(mask CLENR0_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CLENR0) AtomicClearBits(mask CLENR0_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CLENR0_Mask struct{ mmio.UM32 }

func (rm CLENR0_Mask) Load() CLENR0_Bits   { return CLENR0_Bits(rm.UM32.Load()) }
func (rm CLENR0_Mask) Store(b CLENR0_Bits) { rm.UM32.Store(uint32(b)) }

type PPFC_Bits uint32

func (b PPFC_Bits) Field(mask PPFC_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PPFC_Bits) J(v int) PPFC_Bits {
	return PPFC_Bits(bits.Make32(v, uint32(mask)))
}

type PPFC struct{ mmio.U32 }

func (r *PPFC) Bits(mask PPFC_Bits) PPFC_Bits { return PPFC_Bits(r.U32.Bits(uint32(mask))) }
func (r *PPFC) StoreBits(mask, b PPFC_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PPFC) SetBits(mask PPFC_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PPFC) ClearBits(mask PPFC_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PPFC) Load() PPFC_Bits               { return PPFC_Bits(r.U32.Load()) }
func (r *PPFC) Store(b PPFC_Bits)             { r.U32.Store(uint32(b)) }

func (r *PPFC) AtomicStoreBits(mask, b PPFC_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PPFC) AtomicSetBits(mask PPFC_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PPFC) AtomicClearBits(mask PPFC_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PPFC_Mask struct{ mmio.UM32 }

func (rm PPFC_Mask) Load() PPFC_Bits   { return PPFC_Bits(rm.UM32.Load()) }
func (rm PPFC_Mask) Store(b PPFC_Bits) { rm.UM32.Store(uint32(b)) }

type NUMRAMBLOCK_Bits uint32

func (b NUMRAMBLOCK_Bits) Field(mask NUMRAMBLOCK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NUMRAMBLOCK_Bits) J(v int) NUMRAMBLOCK_Bits {
	return NUMRAMBLOCK_Bits(bits.Make32(v, uint32(mask)))
}

type NUMRAMBLOCK struct{ mmio.U32 }

func (r *NUMRAMBLOCK) Bits(mask NUMRAMBLOCK_Bits) NUMRAMBLOCK_Bits {
	return NUMRAMBLOCK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *NUMRAMBLOCK) StoreBits(mask, b NUMRAMBLOCK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NUMRAMBLOCK) SetBits(mask NUMRAMBLOCK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *NUMRAMBLOCK) ClearBits(mask NUMRAMBLOCK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *NUMRAMBLOCK) Load() NUMRAMBLOCK_Bits             { return NUMRAMBLOCK_Bits(r.U32.Load()) }
func (r *NUMRAMBLOCK) Store(b NUMRAMBLOCK_Bits)           { r.U32.Store(uint32(b)) }

func (r *NUMRAMBLOCK) AtomicStoreBits(mask, b NUMRAMBLOCK_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *NUMRAMBLOCK) AtomicSetBits(mask NUMRAMBLOCK_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *NUMRAMBLOCK) AtomicClearBits(mask NUMRAMBLOCK_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type NUMRAMBLOCK_Mask struct{ mmio.UM32 }

func (rm NUMRAMBLOCK_Mask) Load() NUMRAMBLOCK_Bits   { return NUMRAMBLOCK_Bits(rm.UM32.Load()) }
func (rm NUMRAMBLOCK_Mask) Store(b NUMRAMBLOCK_Bits) { rm.UM32.Store(uint32(b)) }

type SIZERAMBLOCK_Bits uint32

func (b SIZERAMBLOCK_Bits) Field(mask SIZERAMBLOCK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SIZERAMBLOCK_Bits) J(v int) SIZERAMBLOCK_Bits {
	return SIZERAMBLOCK_Bits(bits.Make32(v, uint32(mask)))
}

type SIZERAMBLOCK struct{ mmio.U32 }

func (r *SIZERAMBLOCK) Bits(mask SIZERAMBLOCK_Bits) SIZERAMBLOCK_Bits {
	return SIZERAMBLOCK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *SIZERAMBLOCK) StoreBits(mask, b SIZERAMBLOCK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SIZERAMBLOCK) SetBits(mask SIZERAMBLOCK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SIZERAMBLOCK) ClearBits(mask SIZERAMBLOCK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SIZERAMBLOCK) Load() SIZERAMBLOCK_Bits             { return SIZERAMBLOCK_Bits(r.U32.Load()) }
func (r *SIZERAMBLOCK) Store(b SIZERAMBLOCK_Bits)           { r.U32.Store(uint32(b)) }

func (r *SIZERAMBLOCK) AtomicStoreBits(mask, b SIZERAMBLOCK_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *SIZERAMBLOCK) AtomicSetBits(mask SIZERAMBLOCK_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SIZERAMBLOCK) AtomicClearBits(mask SIZERAMBLOCK_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type SIZERAMBLOCK_Mask struct{ mmio.UM32 }

func (rm SIZERAMBLOCK_Mask) Load() SIZERAMBLOCK_Bits   { return SIZERAMBLOCK_Bits(rm.UM32.Load()) }
func (rm SIZERAMBLOCK_Mask) Store(b SIZERAMBLOCK_Bits) { rm.UM32.Store(uint32(b)) }

type CONFIGID_Bits uint32

func (b CONFIGID_Bits) Field(mask CONFIGID_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CONFIGID_Bits) J(v int) CONFIGID_Bits {
	return CONFIGID_Bits(bits.Make32(v, uint32(mask)))
}

type CONFIGID struct{ mmio.U32 }

func (r *CONFIGID) Bits(mask CONFIGID_Bits) CONFIGID_Bits {
	return CONFIGID_Bits(r.U32.Bits(uint32(mask)))
}
func (r *CONFIGID) StoreBits(mask, b CONFIGID_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CONFIGID) SetBits(mask CONFIGID_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CONFIGID) ClearBits(mask CONFIGID_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CONFIGID) Load() CONFIGID_Bits             { return CONFIGID_Bits(r.U32.Load()) }
func (r *CONFIGID) Store(b CONFIGID_Bits)           { r.U32.Store(uint32(b)) }

func (r *CONFIGID) AtomicStoreBits(mask, b CONFIGID_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *CONFIGID) AtomicSetBits(mask CONFIGID_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CONFIGID) AtomicClearBits(mask CONFIGID_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type CONFIGID_Mask struct{ mmio.UM32 }

func (rm CONFIGID_Mask) Load() CONFIGID_Bits   { return CONFIGID_Bits(rm.UM32.Load()) }
func (rm CONFIGID_Mask) Store(b CONFIGID_Bits) { rm.UM32.Store(uint32(b)) }

type DEVICEID_Bits uint32

func (b DEVICEID_Bits) Field(mask DEVICEID_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEID_Bits) J(v int) DEVICEID_Bits {
	return DEVICEID_Bits(bits.Make32(v, uint32(mask)))
}

type DEVICEID struct{ mmio.U32 }

func (r *DEVICEID) Bits(mask DEVICEID_Bits) DEVICEID_Bits {
	return DEVICEID_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEVICEID) StoreBits(mask, b DEVICEID_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEVICEID) SetBits(mask DEVICEID_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEVICEID) ClearBits(mask DEVICEID_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEVICEID) Load() DEVICEID_Bits             { return DEVICEID_Bits(r.U32.Load()) }
func (r *DEVICEID) Store(b DEVICEID_Bits)           { r.U32.Store(uint32(b)) }

func (r *DEVICEID) AtomicStoreBits(mask, b DEVICEID_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DEVICEID) AtomicSetBits(mask DEVICEID_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DEVICEID) AtomicClearBits(mask DEVICEID_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DEVICEID_Mask struct{ mmio.UM32 }

func (rm DEVICEID_Mask) Load() DEVICEID_Bits   { return DEVICEID_Bits(rm.UM32.Load()) }
func (rm DEVICEID_Mask) Store(b DEVICEID_Bits) { rm.UM32.Store(uint32(b)) }

type ER_Bits uint32

func (b ER_Bits) Field(mask ER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ER_Bits) J(v int) ER_Bits {
	return ER_Bits(bits.Make32(v, uint32(mask)))
}

type ER struct{ mmio.U32 }

func (r *ER) Bits(mask ER_Bits) ER_Bits { return ER_Bits(r.U32.Bits(uint32(mask))) }
func (r *ER) StoreBits(mask, b ER_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ER) SetBits(mask ER_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *ER) ClearBits(mask ER_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *ER) Load() ER_Bits             { return ER_Bits(r.U32.Load()) }
func (r *ER) Store(b ER_Bits)           { r.U32.Store(uint32(b)) }

func (r *ER) AtomicStoreBits(mask, b ER_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ER) AtomicSetBits(mask ER_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ER) AtomicClearBits(mask ER_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ER_Mask struct{ mmio.UM32 }

func (rm ER_Mask) Load() ER_Bits   { return ER_Bits(rm.UM32.Load()) }
func (rm ER_Mask) Store(b ER_Bits) { rm.UM32.Store(uint32(b)) }

type IR_Bits uint32

func (b IR_Bits) Field(mask IR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IR_Bits) J(v int) IR_Bits {
	return IR_Bits(bits.Make32(v, uint32(mask)))
}

type IR struct{ mmio.U32 }

func (r *IR) Bits(mask IR_Bits) IR_Bits { return IR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IR) StoreBits(mask, b IR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IR) SetBits(mask IR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *IR) ClearBits(mask IR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *IR) Load() IR_Bits             { return IR_Bits(r.U32.Load()) }
func (r *IR) Store(b IR_Bits)           { r.U32.Store(uint32(b)) }

func (r *IR) AtomicStoreBits(mask, b IR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *IR) AtomicSetBits(mask IR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IR) AtomicClearBits(mask IR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type IR_Mask struct{ mmio.UM32 }

func (rm IR_Mask) Load() IR_Bits   { return IR_Bits(rm.UM32.Load()) }
func (rm IR_Mask) Store(b IR_Bits) { rm.UM32.Store(uint32(b)) }

type DEVICEADDRTYPE_Bits uint32

func (b DEVICEADDRTYPE_Bits) Field(mask DEVICEADDRTYPE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEADDRTYPE_Bits) J(v int) DEVICEADDRTYPE_Bits {
	return DEVICEADDRTYPE_Bits(bits.Make32(v, uint32(mask)))
}

type DEVICEADDRTYPE struct{ mmio.U32 }

func (r *DEVICEADDRTYPE) Bits(mask DEVICEADDRTYPE_Bits) DEVICEADDRTYPE_Bits {
	return DEVICEADDRTYPE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEVICEADDRTYPE) StoreBits(mask, b DEVICEADDRTYPE_Bits) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *DEVICEADDRTYPE) SetBits(mask DEVICEADDRTYPE_Bits)   { r.U32.SetBits(uint32(mask)) }
func (r *DEVICEADDRTYPE) ClearBits(mask DEVICEADDRTYPE_Bits) { r.U32.ClearBits(uint32(mask)) }
func (r *DEVICEADDRTYPE) Load() DEVICEADDRTYPE_Bits          { return DEVICEADDRTYPE_Bits(r.U32.Load()) }
func (r *DEVICEADDRTYPE) Store(b DEVICEADDRTYPE_Bits)        { r.U32.Store(uint32(b)) }

func (r *DEVICEADDRTYPE) AtomicStoreBits(mask, b DEVICEADDRTYPE_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DEVICEADDRTYPE) AtomicSetBits(mask DEVICEADDRTYPE_Bits) { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DEVICEADDRTYPE) AtomicClearBits(mask DEVICEADDRTYPE_Bits) {
	r.U32.AtomicClearBits(uint32(mask))
}

type DEVICEADDRTYPE_Mask struct{ mmio.UM32 }

func (rm DEVICEADDRTYPE_Mask) Load() DEVICEADDRTYPE_Bits   { return DEVICEADDRTYPE_Bits(rm.UM32.Load()) }
func (rm DEVICEADDRTYPE_Mask) Store(b DEVICEADDRTYPE_Bits) { rm.UM32.Store(uint32(b)) }

type DEVICEADDR_Bits uint32

func (b DEVICEADDR_Bits) Field(mask DEVICEADDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEVICEADDR_Bits) J(v int) DEVICEADDR_Bits {
	return DEVICEADDR_Bits(bits.Make32(v, uint32(mask)))
}

type DEVICEADDR struct{ mmio.U32 }

func (r *DEVICEADDR) Bits(mask DEVICEADDR_Bits) DEVICEADDR_Bits {
	return DEVICEADDR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEVICEADDR) StoreBits(mask, b DEVICEADDR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEVICEADDR) SetBits(mask DEVICEADDR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEVICEADDR) ClearBits(mask DEVICEADDR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEVICEADDR) Load() DEVICEADDR_Bits             { return DEVICEADDR_Bits(r.U32.Load()) }
func (r *DEVICEADDR) Store(b DEVICEADDR_Bits)           { r.U32.Store(uint32(b)) }

func (r *DEVICEADDR) AtomicStoreBits(mask, b DEVICEADDR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DEVICEADDR) AtomicSetBits(mask DEVICEADDR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DEVICEADDR) AtomicClearBits(mask DEVICEADDR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DEVICEADDR_Mask struct{ mmio.UM32 }

func (rm DEVICEADDR_Mask) Load() DEVICEADDR_Bits   { return DEVICEADDR_Bits(rm.UM32.Load()) }
func (rm DEVICEADDR_Mask) Store(b DEVICEADDR_Bits) { rm.UM32.Store(uint32(b)) }

type OVERRIDDEN_Bits uint32

func (b OVERRIDDEN_Bits) Field(mask OVERRIDDEN_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OVERRIDDEN_Bits) J(v int) OVERRIDDEN_Bits {
	return OVERRIDDEN_Bits(bits.Make32(v, uint32(mask)))
}

type OVERRIDDEN struct{ mmio.U32 }

func (r *OVERRIDDEN) Bits(mask OVERRIDDEN_Bits) OVERRIDDEN_Bits {
	return OVERRIDDEN_Bits(r.U32.Bits(uint32(mask)))
}
func (r *OVERRIDDEN) StoreBits(mask, b OVERRIDDEN_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OVERRIDDEN) SetBits(mask OVERRIDDEN_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *OVERRIDDEN) ClearBits(mask OVERRIDDEN_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *OVERRIDDEN) Load() OVERRIDDEN_Bits             { return OVERRIDDEN_Bits(r.U32.Load()) }
func (r *OVERRIDDEN) Store(b OVERRIDDEN_Bits)           { r.U32.Store(uint32(b)) }

func (r *OVERRIDDEN) AtomicStoreBits(mask, b OVERRIDDEN_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *OVERRIDDEN) AtomicSetBits(mask OVERRIDDEN_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OVERRIDDEN) AtomicClearBits(mask OVERRIDDEN_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OVERRIDDEN_Mask struct{ mmio.UM32 }

func (rm OVERRIDDEN_Mask) Load() OVERRIDDEN_Bits   { return OVERRIDDEN_Bits(rm.UM32.Load()) }
func (rm OVERRIDDEN_Mask) Store(b OVERRIDDEN_Bits) { rm.UM32.Store(uint32(b)) }

func (p *Periph) NRF_1MBIT_OK() OVERRIDDEN_Mask {
	return OVERRIDDEN_Mask{mmio.UM32{&p.OVERRIDDEN.U32, uint32(NRF_1MBIT_OK)}}
}

func (p *Periph) BLE_1MBIT_OK() OVERRIDDEN_Mask {
	return OVERRIDDEN_Mask{mmio.UM32{&p.OVERRIDDEN.U32, uint32(BLE_1MBIT_OK)}}
}

type NRF_1MBIT_Bits uint32

func (b NRF_1MBIT_Bits) Field(mask NRF_1MBIT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NRF_1MBIT_Bits) J(v int) NRF_1MBIT_Bits {
	return NRF_1MBIT_Bits(bits.Make32(v, uint32(mask)))
}

type NRF_1MBIT struct{ mmio.U32 }

func (r *NRF_1MBIT) Bits(mask NRF_1MBIT_Bits) NRF_1MBIT_Bits {
	return NRF_1MBIT_Bits(r.U32.Bits(uint32(mask)))
}
func (r *NRF_1MBIT) StoreBits(mask, b NRF_1MBIT_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NRF_1MBIT) SetBits(mask NRF_1MBIT_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *NRF_1MBIT) ClearBits(mask NRF_1MBIT_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *NRF_1MBIT) Load() NRF_1MBIT_Bits             { return NRF_1MBIT_Bits(r.U32.Load()) }
func (r *NRF_1MBIT) Store(b NRF_1MBIT_Bits)           { r.U32.Store(uint32(b)) }

func (r *NRF_1MBIT) AtomicStoreBits(mask, b NRF_1MBIT_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *NRF_1MBIT) AtomicSetBits(mask NRF_1MBIT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *NRF_1MBIT) AtomicClearBits(mask NRF_1MBIT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type NRF_1MBIT_Mask struct{ mmio.UM32 }

func (rm NRF_1MBIT_Mask) Load() NRF_1MBIT_Bits   { return NRF_1MBIT_Bits(rm.UM32.Load()) }
func (rm NRF_1MBIT_Mask) Store(b NRF_1MBIT_Bits) { rm.UM32.Store(uint32(b)) }

type BLE_1MBIT_Bits uint32

func (b BLE_1MBIT_Bits) Field(mask BLE_1MBIT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BLE_1MBIT_Bits) J(v int) BLE_1MBIT_Bits {
	return BLE_1MBIT_Bits(bits.Make32(v, uint32(mask)))
}

type BLE_1MBIT struct{ mmio.U32 }

func (r *BLE_1MBIT) Bits(mask BLE_1MBIT_Bits) BLE_1MBIT_Bits {
	return BLE_1MBIT_Bits(r.U32.Bits(uint32(mask)))
}
func (r *BLE_1MBIT) StoreBits(mask, b BLE_1MBIT_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BLE_1MBIT) SetBits(mask BLE_1MBIT_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *BLE_1MBIT) ClearBits(mask BLE_1MBIT_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *BLE_1MBIT) Load() BLE_1MBIT_Bits             { return BLE_1MBIT_Bits(r.U32.Load()) }
func (r *BLE_1MBIT) Store(b BLE_1MBIT_Bits)           { r.U32.Store(uint32(b)) }

func (r *BLE_1MBIT) AtomicStoreBits(mask, b BLE_1MBIT_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *BLE_1MBIT) AtomicSetBits(mask BLE_1MBIT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BLE_1MBIT) AtomicClearBits(mask BLE_1MBIT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type BLE_1MBIT_Mask struct{ mmio.UM32 }

func (rm BLE_1MBIT_Mask) Load() BLE_1MBIT_Bits   { return BLE_1MBIT_Bits(rm.UM32.Load()) }
func (rm BLE_1MBIT_Mask) Store(b BLE_1MBIT_Bits) { rm.UM32.Store(uint32(b)) }

type INFO_PART_Bits uint32

func (b INFO_PART_Bits) Field(mask INFO_PART_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_PART_Bits) J(v int) INFO_PART_Bits {
	return INFO_PART_Bits(bits.Make32(v, uint32(mask)))
}

type INFO_PART struct{ mmio.U32 }

func (r *INFO_PART) Bits(mask INFO_PART_Bits) INFO_PART_Bits {
	return INFO_PART_Bits(r.U32.Bits(uint32(mask)))
}
func (r *INFO_PART) StoreBits(mask, b INFO_PART_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INFO_PART) SetBits(mask INFO_PART_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *INFO_PART) ClearBits(mask INFO_PART_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *INFO_PART) Load() INFO_PART_Bits             { return INFO_PART_Bits(r.U32.Load()) }
func (r *INFO_PART) Store(b INFO_PART_Bits)           { r.U32.Store(uint32(b)) }

func (r *INFO_PART) AtomicStoreBits(mask, b INFO_PART_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *INFO_PART) AtomicSetBits(mask INFO_PART_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *INFO_PART) AtomicClearBits(mask INFO_PART_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type INFO_PART_Mask struct{ mmio.UM32 }

func (rm INFO_PART_Mask) Load() INFO_PART_Bits   { return INFO_PART_Bits(rm.UM32.Load()) }
func (rm INFO_PART_Mask) Store(b INFO_PART_Bits) { rm.UM32.Store(uint32(b)) }

type INFO_VARIANT_Bits uint32

func (b INFO_VARIANT_Bits) Field(mask INFO_VARIANT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_VARIANT_Bits) J(v int) INFO_VARIANT_Bits {
	return INFO_VARIANT_Bits(bits.Make32(v, uint32(mask)))
}

type INFO_VARIANT struct{ mmio.U32 }

func (r *INFO_VARIANT) Bits(mask INFO_VARIANT_Bits) INFO_VARIANT_Bits {
	return INFO_VARIANT_Bits(r.U32.Bits(uint32(mask)))
}
func (r *INFO_VARIANT) StoreBits(mask, b INFO_VARIANT_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INFO_VARIANT) SetBits(mask INFO_VARIANT_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *INFO_VARIANT) ClearBits(mask INFO_VARIANT_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *INFO_VARIANT) Load() INFO_VARIANT_Bits             { return INFO_VARIANT_Bits(r.U32.Load()) }
func (r *INFO_VARIANT) Store(b INFO_VARIANT_Bits)           { r.U32.Store(uint32(b)) }

func (r *INFO_VARIANT) AtomicStoreBits(mask, b INFO_VARIANT_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *INFO_VARIANT) AtomicSetBits(mask INFO_VARIANT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *INFO_VARIANT) AtomicClearBits(mask INFO_VARIANT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type INFO_VARIANT_Mask struct{ mmio.UM32 }

func (rm INFO_VARIANT_Mask) Load() INFO_VARIANT_Bits   { return INFO_VARIANT_Bits(rm.UM32.Load()) }
func (rm INFO_VARIANT_Mask) Store(b INFO_VARIANT_Bits) { rm.UM32.Store(uint32(b)) }

type INFO_PACKAGE_Bits uint32

func (b INFO_PACKAGE_Bits) Field(mask INFO_PACKAGE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_PACKAGE_Bits) J(v int) INFO_PACKAGE_Bits {
	return INFO_PACKAGE_Bits(bits.Make32(v, uint32(mask)))
}

type INFO_PACKAGE struct{ mmio.U32 }

func (r *INFO_PACKAGE) Bits(mask INFO_PACKAGE_Bits) INFO_PACKAGE_Bits {
	return INFO_PACKAGE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *INFO_PACKAGE) StoreBits(mask, b INFO_PACKAGE_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INFO_PACKAGE) SetBits(mask INFO_PACKAGE_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *INFO_PACKAGE) ClearBits(mask INFO_PACKAGE_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *INFO_PACKAGE) Load() INFO_PACKAGE_Bits             { return INFO_PACKAGE_Bits(r.U32.Load()) }
func (r *INFO_PACKAGE) Store(b INFO_PACKAGE_Bits)           { r.U32.Store(uint32(b)) }

func (r *INFO_PACKAGE) AtomicStoreBits(mask, b INFO_PACKAGE_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *INFO_PACKAGE) AtomicSetBits(mask INFO_PACKAGE_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *INFO_PACKAGE) AtomicClearBits(mask INFO_PACKAGE_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type INFO_PACKAGE_Mask struct{ mmio.UM32 }

func (rm INFO_PACKAGE_Mask) Load() INFO_PACKAGE_Bits   { return INFO_PACKAGE_Bits(rm.UM32.Load()) }
func (rm INFO_PACKAGE_Mask) Store(b INFO_PACKAGE_Bits) { rm.UM32.Store(uint32(b)) }

type INFO_RAM_Bits uint32

func (b INFO_RAM_Bits) Field(mask INFO_RAM_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_RAM_Bits) J(v int) INFO_RAM_Bits {
	return INFO_RAM_Bits(bits.Make32(v, uint32(mask)))
}

type INFO_RAM struct{ mmio.U32 }

func (r *INFO_RAM) Bits(mask INFO_RAM_Bits) INFO_RAM_Bits {
	return INFO_RAM_Bits(r.U32.Bits(uint32(mask)))
}
func (r *INFO_RAM) StoreBits(mask, b INFO_RAM_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INFO_RAM) SetBits(mask INFO_RAM_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *INFO_RAM) ClearBits(mask INFO_RAM_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *INFO_RAM) Load() INFO_RAM_Bits             { return INFO_RAM_Bits(r.U32.Load()) }
func (r *INFO_RAM) Store(b INFO_RAM_Bits)           { r.U32.Store(uint32(b)) }

func (r *INFO_RAM) AtomicStoreBits(mask, b INFO_RAM_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *INFO_RAM) AtomicSetBits(mask INFO_RAM_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *INFO_RAM) AtomicClearBits(mask INFO_RAM_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type INFO_RAM_Mask struct{ mmio.UM32 }

func (rm INFO_RAM_Mask) Load() INFO_RAM_Bits   { return INFO_RAM_Bits(rm.UM32.Load()) }
func (rm INFO_RAM_Mask) Store(b INFO_RAM_Bits) { rm.UM32.Store(uint32(b)) }

type INFO_FLASH_Bits uint32

func (b INFO_FLASH_Bits) Field(mask INFO_FLASH_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask INFO_FLASH_Bits) J(v int) INFO_FLASH_Bits {
	return INFO_FLASH_Bits(bits.Make32(v, uint32(mask)))
}

type INFO_FLASH struct{ mmio.U32 }

func (r *INFO_FLASH) Bits(mask INFO_FLASH_Bits) INFO_FLASH_Bits {
	return INFO_FLASH_Bits(r.U32.Bits(uint32(mask)))
}
func (r *INFO_FLASH) StoreBits(mask, b INFO_FLASH_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *INFO_FLASH) SetBits(mask INFO_FLASH_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *INFO_FLASH) ClearBits(mask INFO_FLASH_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *INFO_FLASH) Load() INFO_FLASH_Bits             { return INFO_FLASH_Bits(r.U32.Load()) }
func (r *INFO_FLASH) Store(b INFO_FLASH_Bits)           { r.U32.Store(uint32(b)) }

func (r *INFO_FLASH) AtomicStoreBits(mask, b INFO_FLASH_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *INFO_FLASH) AtomicSetBits(mask INFO_FLASH_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *INFO_FLASH) AtomicClearBits(mask INFO_FLASH_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type INFO_FLASH_Mask struct{ mmio.UM32 }

func (rm INFO_FLASH_Mask) Load() INFO_FLASH_Bits   { return INFO_FLASH_Bits(rm.UM32.Load()) }
func (rm INFO_FLASH_Mask) Store(b INFO_FLASH_Bits) { rm.UM32.Store(uint32(b)) }

type TEMP_A_Bits uint32

func (b TEMP_A_Bits) Field(mask TEMP_A_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_A_Bits) J(v int) TEMP_A_Bits {
	return TEMP_A_Bits(bits.Make32(v, uint32(mask)))
}

type TEMP_A struct{ mmio.U32 }

func (r *TEMP_A) Bits(mask TEMP_A_Bits) TEMP_A_Bits { return TEMP_A_Bits(r.U32.Bits(uint32(mask))) }
func (r *TEMP_A) StoreBits(mask, b TEMP_A_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_A) SetBits(mask TEMP_A_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TEMP_A) ClearBits(mask TEMP_A_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TEMP_A) Load() TEMP_A_Bits                 { return TEMP_A_Bits(r.U32.Load()) }
func (r *TEMP_A) Store(b TEMP_A_Bits)               { r.U32.Store(uint32(b)) }

func (r *TEMP_A) AtomicStoreBits(mask, b TEMP_A_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_A) AtomicSetBits(mask TEMP_A_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TEMP_A) AtomicClearBits(mask TEMP_A_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TEMP_A_Mask struct{ mmio.UM32 }

func (rm TEMP_A_Mask) Load() TEMP_A_Bits   { return TEMP_A_Bits(rm.UM32.Load()) }
func (rm TEMP_A_Mask) Store(b TEMP_A_Bits) { rm.UM32.Store(uint32(b)) }

type TEMP_B_Bits uint32

func (b TEMP_B_Bits) Field(mask TEMP_B_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_B_Bits) J(v int) TEMP_B_Bits {
	return TEMP_B_Bits(bits.Make32(v, uint32(mask)))
}

type TEMP_B struct{ mmio.U32 }

func (r *TEMP_B) Bits(mask TEMP_B_Bits) TEMP_B_Bits { return TEMP_B_Bits(r.U32.Bits(uint32(mask))) }
func (r *TEMP_B) StoreBits(mask, b TEMP_B_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_B) SetBits(mask TEMP_B_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TEMP_B) ClearBits(mask TEMP_B_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TEMP_B) Load() TEMP_B_Bits                 { return TEMP_B_Bits(r.U32.Load()) }
func (r *TEMP_B) Store(b TEMP_B_Bits)               { r.U32.Store(uint32(b)) }

func (r *TEMP_B) AtomicStoreBits(mask, b TEMP_B_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_B) AtomicSetBits(mask TEMP_B_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TEMP_B) AtomicClearBits(mask TEMP_B_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TEMP_B_Mask struct{ mmio.UM32 }

func (rm TEMP_B_Mask) Load() TEMP_B_Bits   { return TEMP_B_Bits(rm.UM32.Load()) }
func (rm TEMP_B_Mask) Store(b TEMP_B_Bits) { rm.UM32.Store(uint32(b)) }

type TEMP_T_Bits uint32

func (b TEMP_T_Bits) Field(mask TEMP_T_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TEMP_T_Bits) J(v int) TEMP_T_Bits {
	return TEMP_T_Bits(bits.Make32(v, uint32(mask)))
}

type TEMP_T struct{ mmio.U32 }

func (r *TEMP_T) Bits(mask TEMP_T_Bits) TEMP_T_Bits { return TEMP_T_Bits(r.U32.Bits(uint32(mask))) }
func (r *TEMP_T) StoreBits(mask, b TEMP_T_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_T) SetBits(mask TEMP_T_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TEMP_T) ClearBits(mask TEMP_T_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TEMP_T) Load() TEMP_T_Bits                 { return TEMP_T_Bits(r.U32.Load()) }
func (r *TEMP_T) Store(b TEMP_T_Bits)               { r.U32.Store(uint32(b)) }

func (r *TEMP_T) AtomicStoreBits(mask, b TEMP_T_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TEMP_T) AtomicSetBits(mask TEMP_T_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TEMP_T) AtomicClearBits(mask TEMP_T_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TEMP_T_Mask struct{ mmio.UM32 }

func (rm TEMP_T_Mask) Load() TEMP_T_Bits   { return TEMP_T_Bits(rm.UM32.Load()) }
func (rm TEMP_T_Mask) Store(b TEMP_T_Bits) { rm.UM32.Store(uint32(b)) }

type NFC_TAGHEADER_Bits uint32

func (b NFC_TAGHEADER_Bits) Field(mask NFC_TAGHEADER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NFC_TAGHEADER_Bits) J(v int) NFC_TAGHEADER_Bits {
	return NFC_TAGHEADER_Bits(bits.Make32(v, uint32(mask)))
}

type NFC_TAGHEADER struct{ mmio.U32 }

func (r *NFC_TAGHEADER) Bits(mask NFC_TAGHEADER_Bits) NFC_TAGHEADER_Bits {
	return NFC_TAGHEADER_Bits(r.U32.Bits(uint32(mask)))
}
func (r *NFC_TAGHEADER) StoreBits(mask, b NFC_TAGHEADER_Bits) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *NFC_TAGHEADER) SetBits(mask NFC_TAGHEADER_Bits)   { r.U32.SetBits(uint32(mask)) }
func (r *NFC_TAGHEADER) ClearBits(mask NFC_TAGHEADER_Bits) { r.U32.ClearBits(uint32(mask)) }
func (r *NFC_TAGHEADER) Load() NFC_TAGHEADER_Bits          { return NFC_TAGHEADER_Bits(r.U32.Load()) }
func (r *NFC_TAGHEADER) Store(b NFC_TAGHEADER_Bits)        { r.U32.Store(uint32(b)) }

func (r *NFC_TAGHEADER) AtomicStoreBits(mask, b NFC_TAGHEADER_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *NFC_TAGHEADER) AtomicSetBits(mask NFC_TAGHEADER_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *NFC_TAGHEADER) AtomicClearBits(mask NFC_TAGHEADER_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type NFC_TAGHEADER_Mask struct{ mmio.UM32 }

func (rm NFC_TAGHEADER_Mask) Load() NFC_TAGHEADER_Bits   { return NFC_TAGHEADER_Bits(rm.UM32.Load()) }
func (rm NFC_TAGHEADER_Mask) Store(b NFC_TAGHEADER_Bits) { rm.UM32.Store(uint32(b)) }
