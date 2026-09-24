package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Hardware telemetry packet
type TelemetryPacket struct {
	DeviceID uint16
	Sensor   float32
	IsActive bool
}

func main() {
	packet := TelemetryPacket{
		DeviceID: 4095,
		Sensor:   98.6,
		IsActive: true,
	}

	buf := new(bytes.Buffer)

	// Write struct to byte buffer in Little Endian format
	err := binary.Write(buf, binary.LittleEndian, packet)
	if err != nil {
		fmt.Println("Binary packing failed:", err)
		return
	}
	fmt.Printf("Raw Bytes Transmitted: % x\n", buf.Bytes())

	// Unpack the bytes back into a struct
	var received TelemetryPacket
	reader := bytes.NewReader(buf.Bytes())
	binary.Read(reader, binary.LittleEndian, &received)
	
	fmt.Printf("Unpacked Data: %+v\n", received)
}