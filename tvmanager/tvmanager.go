package tvmanager

import (
	"fmt"
	"github.com/gwwfps/simple-webos-remote/config"
	"github.com/rs/zerolog/log"
	"github.com/snabb/webostv"
	"time"
)

type TVManager struct {
	cfg *config.Config
	tv  *webostv.Tv

	Connecting    bool
	ConnectionErr error
}

func New(cfg *config.Config) *TVManager {
	return &TVManager{cfg: cfg}
}

func (m *TVManager) Start() {
	ticker := time.NewTicker(time.Second)
	for {
		m.TryConnect()
		<-ticker.C
	}
}

func (m *TVManager) TryConnect() {
	if m.Connected() {
		return
	}

	m.Connecting = true
	defer func() {
		m.Connecting = false
	}()

	tv, err := webostv.DefaultDialer.Dial(m.cfg.TvAddr)
	if err != nil {
		log.Error().Err(err).Msg("cannot connect to TV")
	} else {
		m.tv = tv
		go tv.MessageHandler()
		err = m.register()
		if err != nil {
			defer tv.Close()
			m.tv = nil
			log.Error().Err(err).Msg("cannot register with TV")
		}
	}
	m.ConnectionErr = err
}

func (m *TVManager) register() error {
	key, err := m.tv.Register(m.cfg.ClientKey)
	if err != nil {
		return err
	}
	m.cfg.ClientKey = key
	m.cfg.Save()
	return nil
}

func (m *TVManager) Close() error {
	defer func() {
		m.tv = nil
	}()

	if m.tv != nil {
		return m.tv.Close()
	}
	return nil
}

func (m *TVManager) Connected() bool {
	return m.tv != nil
}

func (m *TVManager) checkConnection() error {
	if !m.Connected() {
		return fmt.Errorf("not connected to TV")
	}
	return nil
}
