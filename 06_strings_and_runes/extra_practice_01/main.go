package main

import (
	"bytes"
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrInvalidSignature = errors.New("invalid frame signature: expected 0xAA")
var ErrMalformedFrame = errors.New("malformed frame: corrupted payload or delimiter")
var ErrCorruptedRune = errors.New("corrupted utf-8 rune in stream")

func main() {
	b := FrameBuilder(43, "aaa", '🔥')
	processStream("Валидный фрейм", b)

	CorruptedStreamHandler()

}

func FrameBuilder(id byte, text string, r rune) []byte {
	var b bytes.Buffer
	b.WriteByte(0xAA)
	b.WriteByte(id)
	b.WriteByte(':')
	b.WriteString(text)
	b.WriteByte(0)
	b.WriteRune(r)
	b.WriteByte(0xFF)

	return b.Bytes()
}

func StreamInspector(b []byte) error {
	if len(b) <= 0 {
		return ErrMalformedFrame
	}

	bb := bytes.NewBuffer(b)
	if firstByte, err := bb.ReadByte(); err != nil || firstByte != 0xAA {
		return ErrInvalidSignature
	}

	if _, err := bb.ReadByte(); err != nil {
		return ErrInvalidSignature
	}

	if del, err := bb.ReadByte(); err != nil || del != ':' {
		return ErrMalformedFrame
	}

	if _, err := bb.ReadBytes(0x00); err != nil {
		return ErrInvalidSignature
	}

	r, size, err := bb.ReadRune()
	if err != nil {
		return ErrInvalidSignature
	}

	if r == utf8.RuneError && size == 1 {
		return ErrCorruptedRune
	}

	if endByte, err := bb.ReadByte(); err != nil || endByte != 0xFF {
		return ErrMalformedFrame
	}

	fmt.Printf(
		"Rune is: %c; "+
			"Count of bytes: %d\n"+
			"HEX dump of bytes is: % X\n\n",
		r, size, []byte(string(r)),
	)

	return nil
}

func processStream(name string, b []byte) {
	fmt.Printf("=== Обработка потока: %s ===\n", name)
	err := StreamInspector(b)
	if err == nil {
		fmt.Println("Успех: пакет разобран корректно!\n")
		return
	}

	if errors.Is(err, ErrInvalidSignature) {
		fmt.Println("Ошибка: неверная сигнатура фрейма!\n")
	} else if errors.Is(err, ErrMalformedFrame) {
		fmt.Println("Ошибка: сломались разделители или длина данных!\n")
	} else if errors.Is(err, ErrCorruptedRune) {
		fmt.Println("Ошибка: в потоке обнаружен битый байт UTF-8!\n")
	} else {
		fmt.Printf("Неизвестная ошибка: %v\n\n", err)
	}
}

func CorruptedStreamHandler() {
	b := []byte{0xAA, 200, ':', 'a', 'b', 'c', 0, 0x9F, 255}
	processStream("Битая байт UTF-8", b)
}
