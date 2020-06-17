package tvmanager

import "image"

func (m *TVManager) UpdatePointerPos(pos image.Point) error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.pointerSocket.Move(pos.X, pos.Y)
}

func (m *TVManager) ClickPointer() error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.pointerSocket.Input("click", "")
}

func (m *TVManager) PressButton(button string) error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	return m.pointerSocket.Input("button", button)
}
