package strfile

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type Flags uint32

// #define STR_RANDOM   0x1 /* randomized pointers */
// #define STR_ORDERED  0x2 /* ordered pointers */
// #define STR_ROTATED  0x4 /* rot-13'd text */
const (
	Random  Flags = 1 << iota // 0b001
	Ordered                   // 0b010
	Rotated                   // 0headerBuffer00
)

const (
	unsignedLongSize byte = 4
	delimSize        byte = 4
)

const strFileHeaderSize byte = (unsignedLongSize*5 + delimSize)

// The format of the header is:
//
// typedef struct {               /* information table */
// #define	VERSION 1
//
//	unsigned long str_version;    /* version number */
//	unsigned long str_numstr;     /* # of strings in the file */
//	unsigned long str_longlen;    /* length of longest string */
//	unsigned long str_shortlen;   /* length of shortest string */
//
// #define  STR_RANDOM  0x1       /* randomized pointers */
// #define  STR_ORDERED 0x2       /* ordered pointers */
// #define  STR_ROTATED 0x4       /* rot-13'd text */
//
//	unsigned long	str_flags;      /* bit field for flags */
//	unsigned char	stuff[4];       /* long aligned space */
//
// #define	str_delim	stuff[0]    /* delimiting character */
// } STRFILE;
//
// All fields are written in network byte order.
type Header struct {
	Delim    [4]byte
	Version  uint32
	NumStr   uint32
	LongLen  uint32
	ShortLen uint32
	Flags    uint32
}

// Permite formatar a saída para a função fmt.Println
func (h Header) String() string {
	return fmt.Sprintf(
		`{
    Header.Delim: %c,
    Header.Version: %d,
    Header.NumStr: %d,
    Header.LongLen: %d,
    Header.ShorLen: %d,
    Header.Flags: %d
}`,
		h.Delim[0],
		h.Version,
		h.NumStr,
		h.LongLen,
		h.ShortLen,
		h.Flags)
}

// Retorna o cabeçalho do arquivo de índice
func GetHeader(strFilePath string) (Header, error) {
	header := Header{}

	strfile, err := os.Open(strFilePath)
	if err != nil {
		return header, err
	}

	defer func() (Header, error) {
		if err = strfile.Close(); err != nil {
			return header, err
		}
		return header, nil
	}()

	reader := bufio.NewReader(strfile)
	headerBuffer := make([]byte, strFileHeaderSize)
	_, err = reader.Read(headerBuffer)
	if err != nil {
		return header, nil
	}

	// Todos os campos estão no formato network byte order
	// que é o mesmo que big endian
	header.Version = binary.BigEndian.Uint32(headerBuffer[0:4])
	header.NumStr = binary.BigEndian.Uint32(headerBuffer[4:8])
	header.LongLen = binary.BigEndian.Uint32(headerBuffer[8:12])
	header.ShortLen = binary.BigEndian.Uint32(headerBuffer[12:16])
	header.Flags = binary.BigEndian.Uint32(headerBuffer[16:20])

	copy(header.Delim[:], headerBuffer[20:24])

	return header, nil
}
