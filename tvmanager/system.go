package tvmanager

import (
	"fmt"
	"github.com/linde12/gowol"
)

func (m *TVManager) PowerOn() error {
	if m.Connected() {
		return fmt.Errorf("TV is already powered on and connected")
	}

	packet, err := gowol.NewMagicPacket(m.cfg.TvMac)
	if err == nil {
		err = packet.Send("255.255.255.255")
	}
	return err
}

func (m *TVManager) PowerOff() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	defer m.Close()
	return m.tv.SystemTurnOff()
}

func (m *TVManager) OpenInfo() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.PressButton("INFO")
}
