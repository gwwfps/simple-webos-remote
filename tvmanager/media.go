package tvmanager

func (m *TVManager) VolDown() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.AudioVolumeDown()
}

func (m *TVManager) VolUp() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.AudioVolumeUp()
}

func (m *TVManager) Mute() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.AudioSetMute(true)
}

func (m *TVManager) Unmute() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.AudioSetMute(false)
}

func (m *TVManager) PlayMedia() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.MediaControlsPlay()
}

func (m *TVManager) PauseMedia() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.MediaControlsPause()
}
