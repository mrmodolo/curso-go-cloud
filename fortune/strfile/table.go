package strfile

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

// strfile reads a file containing groups of lines separated by a line containing
// a single percent '%' sign (or other specified delimiter character) and creates
// a data file which contains a header structure and a table of file offsets for
// each group of lines. This allows random access of the strings.

// off_t
// This is a signed integer type used to represent file sizes.
// In the GNU C Library, this type is no narrower than int.
// If the source is compiled with _FILE_OFFSET_BITS == 64 this
// type is transparently replaced by off64_t.const

// https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/sys_types.h.html#tag_13_67
// off_t
// Used for file sizes.
// blkcnt_t and off_t shall be signed integer types.

const offsetBufferSize int = 4

type offset uint32

func GetTableEntry(strFilePath string, lineNumber uint32) (offset, error) {
	var fileOffset offset = 0

	// A linha zero tem offset zero
	if lineNumber == 0 {
		return fileOffset, nil
	}

	header, err := GetHeader(strFilePath)
	if err != nil {
		return fileOffset, err
	}
	if lineNumber > header.NumStr {
		return fileOffset, fmt.Errorf("o número máximo de linhas no arquivo é %d", header.NumStr)
	}

	strfile, err := os.Open(strFilePath)
	if err != nil {
		return fileOffset, err
	}

	defer func() (offset, error) {
		if err = strfile.Close(); err != nil {
			return fileOffset, err
		}
		return fileOffset, nil
	}()

	reader := bufio.NewReader(strfile)
	offsetBuffer := make([]byte, offsetBufferSize)

	// Descartar o cabeçalho e o deslocamento para a linha indicada
	reader.Discard(int(strFileHeaderSize) + (offsetBufferSize * int(lineNumber)))

	n, err := reader.Read(offsetBuffer)
	if err != nil {
		return fileOffset, err
	}

	if n != offsetBufferSize {
		return fileOffset, fmt.Errorf("não foi possível ober o offset para a linha %d", lineNumber)
	}

	// All fields are written in network byte order.
	fileOffset = offset(binary.BigEndian.Uint32(offsetBuffer[0:4]))

	return fileOffset, nil
}

func GetString(fortuneFile string, strFilePath string, lineNumber uint32) (string, error) {
	quote := ""

	header, err := GetHeader(strFilePath)
	if err != nil {
		return quote, err
	}

	off, err := GetTableEntry(strFilePath, lineNumber)
	if err != nil {
		return quote, err
	}

	strfile, err := os.Open(fortuneFile)
	if err != nil {
		return quote, err
	}
	reader := bufio.NewReader(strfile)
	buffer := make([]byte, 1)
	quotationBuffer := make([]byte, header.LongLen)

	reader.Discard(int(off))

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return "", err
		}
		if n < 1 || buffer[0] == header.Delim[0] {
			break
		}
		quotationBuffer = append(quotationBuffer, buffer[0])
	}

	quote = string(quotationBuffer[0 : len(quotationBuffer)-1])
	return quote, nil
}

func PrintTable(strFilePath string) error {
	strfile, err := os.Open(strFilePath)
	if err != nil {
		return err
	}

	defer func() error {
		if err = strfile.Close(); err != nil {
			return err
		}
		return nil
	}()

	reader := bufio.NewReader(strfile)
	offsetBuffer := make([]byte, offsetBufferSize)

	// Descartar o cabeçalho
	reader.Discard(int(strFileHeaderSize))

	for {
		n, err := reader.Read(offsetBuffer)
		if err != nil {
			return err
		}
		if n < offsetBufferSize {
			break
		}
		// All fields are written in network byte order.
		offset := binary.BigEndian.Uint32(offsetBuffer[0:4])
		fmt.Println(offset)
	}

	return nil
}
