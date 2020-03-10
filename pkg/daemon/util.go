package apd

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

	"github.com/SkycoinProject/skycoin/src/util/logging"
)

var logger = func(moduleName string) *logging.Logger {
	masterLogger := logging.NewMasterLogger()
	return masterLogger.PackageLogger(moduleName)
}

const moduleName = "SPD"

// BroadCast broadcasts a UDP packet containing the public key of the local visor.
// Broadcasts is sent on the local network broadcasts address.
func BroadCast(broadCastIP string, port int, data []byte) error {
	address := fmt.Sprintf("%s:%d", broadCastIP, port)
	bAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		logger(moduleName).Errorf("Couldn't resolve broadcast address: %v", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, bAddr)
	if err != nil {
		return err
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			logger(moduleName).WithError(err)
		}
	}()

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func serialize(packet Packet) ([]byte, error) {
	var buff bytes.Buffer
	decoder := gob.NewEncoder(&buff)
	err := decoder.Encode(packet)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// Deserialize decodes a byte to a packet type
func Deserialize(data []byte) (Packet, error) {
	var packet Packet
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&packet)
	if err != nil {
		return Packet{}, err
	}

	return packet, nil
}

// verifyPacket checks if packet received is sent from local skywire-peering-daemon
func verifyPacket(pubKey string, data []byte) bool {
	packet, err := Deserialize(data)
	if err != nil {
		logger(moduleName).Fatalf("Couldn't serialize packet: %s", err)
	}

	return packet.PublicKey == pubKey
}

// SendPacket sends packet to visor via unix sockets
func SendPacket(socketFile string, packet []byte) error {
	conn, err := net.Dial("unix", socketFile)
	if err != nil {
		return err
	}
	
	defer func() {
		err := conn.Close()
		if err != nil {
			logger("SPD").Error(err)
		}
	}()

	n, err := conn.Write(packet)
	if err != nil {
		return err
	}

	logger("SPD").Infof("Wrote %d bytes", n)
	return nil
}
