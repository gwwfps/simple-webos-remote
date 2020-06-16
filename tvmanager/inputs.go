package tvmanager

func (m *TVManager) SwitchInput(id string) error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.tv.TvSwitchInput(id)
}
