// +build l476xx

// Peripheral: SPI_Periph  Serial Peripheral Interface.
// Instances:
//  SPI2  mmap.SPI2_BASE
//  SPI3  mmap.SPI3_BASE
//  SPI1  mmap.SPI1_BASE
// Registers:
//  0x00 32  CR1    Control register 1.
//  0x04 32  CR2    Control register 2.
//  0x08 32  SR     Status register.
//  0x0C 32  DR     Data register.
//  0x10 32  CRCPR  CRC polynomial register.
//  0x14 32  RXCRCR Rx CRC register.
//  0x18 32  TXCRCR Tx CRC register.
// Import:
//  stm32/o/l476xx/mmap
package spi

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	CPHA     CR1_Bits = 0x01 << 0  //+ Clock Phase.
	CPOL     CR1_Bits = 0x01 << 1  //+ Clock Polarity.
	MSTR     CR1_Bits = 0x01 << 2  //+ Master Selection.
	BR       CR1_Bits = 0x07 << 3  //+ BR[2:0] bits (Baud Rate Control).
	SPE      CR1_Bits = 0x01 << 6  //+ SPI Enable.
	LSBFIRST CR1_Bits = 0x01 << 7  //+ Frame Format.
	SSI      CR1_Bits = 0x01 << 8  //+ Internal slave select.
	SSM      CR1_Bits = 0x01 << 9  //+ Software slave management.
	RXONLY   CR1_Bits = 0x01 << 10 //+ Receive only.
	CRCL     CR1_Bits = 0x01 << 11 //+ CRC Length.
	CRCNEXT  CR1_Bits = 0x01 << 12 //+ Transmit CRC next.
	CRCEN    CR1_Bits = 0x01 << 13 //+ Hardware CRC calculation enable.
	BIDIOE   CR1_Bits = 0x01 << 14 //+ Output enable in bidirectional mode.
	BIDIMODE CR1_Bits = 0x01 << 15 //+ Bidirectional data mode enable.
)

const (
	CPHAn     = 0
	CPOLn     = 1
	MSTRn     = 2
	BRn       = 3
	SPEn      = 6
	LSBFIRSTn = 7
	SSIn      = 8
	SSMn      = 9
	RXONLYn   = 10
	CRCLn     = 11
	CRCNEXTn  = 12
	CRCENn    = 13
	BIDIOEn   = 14
	BIDIMODEn = 15
)

const (
	RXDMAEN CR2_Bits = 0x01 << 0  //+ Rx Buffer DMA Enable.
	TXDMAEN CR2_Bits = 0x01 << 1  //+ Tx Buffer DMA Enable.
	SSOE    CR2_Bits = 0x01 << 2  //+ SS Output Enable.
	NSSP    CR2_Bits = 0x01 << 3  //+ NSS pulse management Enable.
	FRF     CR2_Bits = 0x01 << 4  //+ Frame Format Enable.
	ERRIE   CR2_Bits = 0x01 << 5  //+ Error Interrupt Enable.
	RXNEIE  CR2_Bits = 0x01 << 6  //+ RX buffer Not Empty Interrupt Enable.
	TXEIE   CR2_Bits = 0x01 << 7  //+ Tx buffer Empty Interrupt Enable.
	DS      CR2_Bits = 0x0F << 8  //+ DS[3:0] Data Size.
	FRXTH   CR2_Bits = 0x01 << 12 //+ FIFO reception Threshold.
	LDMARX  CR2_Bits = 0x01 << 13 //+ Last DMA transfer for reception.
	LDMATX  CR2_Bits = 0x01 << 14 //+ Last DMA transfer for transmission.
)

const (
	RXDMAENn = 0
	TXDMAENn = 1
	SSOEn    = 2
	NSSPn    = 3
	FRFn     = 4
	ERRIEn   = 5
	RXNEIEn  = 6
	TXEIEn   = 7
	DSn      = 8
	FRXTHn   = 12
	LDMARXn  = 13
	LDMATXn  = 14
)

const (
	RXNE   SR_Bits = 0x01 << 0  //+ Receive buffer Not Empty.
	TXE    SR_Bits = 0x01 << 1  //+ Transmit buffer Empty.
	CHSIDE SR_Bits = 0x01 << 2  //+ Channel side.
	UDR    SR_Bits = 0x01 << 3  //+ Underrun flag.
	CRCERR SR_Bits = 0x01 << 4  //+ CRC Error flag.
	MODF   SR_Bits = 0x01 << 5  //+ Mode fault.
	OVR    SR_Bits = 0x01 << 6  //+ Overrun flag.
	BSY    SR_Bits = 0x01 << 7  //+ Busy flag.
	FRE    SR_Bits = 0x01 << 8  //+ TI frame format error.
	FRLVL  SR_Bits = 0x03 << 9  //+ FIFO Reception Level.
	FTLVL  SR_Bits = 0x03 << 11 //+ FIFO Transmission Level.
)

const (
	RXNEn   = 0
	TXEn    = 1
	CHSIDEn = 2
	UDRn    = 3
	CRCERRn = 4
	MODFn   = 5
	OVRn    = 6
	BSYn    = 7
	FREn    = 8
	FRLVLn  = 9
	FTLVLn  = 11
)

const (
	CRCPOLY CRCPR_Bits = 0xFFFF << 0 //+ CRC polynomial register.
)

const (
	CRCPOLYn = 0
)

const (
	RXCRC RXCRCR_Bits = 0xFFFF << 0 //+ Rx CRC Register.
)

const (
	RXCRCn = 0
)

const (
	TXCRC TXCRCR_Bits = 0xFFFF << 0 //+ Tx CRC Register.
)

const (
	TXCRCn = 0
)
