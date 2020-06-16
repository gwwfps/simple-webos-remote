package tvmanager

import (
	"github.com/godbus/dbus/v5"
)

func (m *TVManager) EnableScreen() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	defer conn.Close()

	dc := conn.Object("org.gnome.Mutter.DisplayConfig", "/org/gnome/Mutter/DisplayConfig")
	res := map[string]interface{}{}
	err = dc.Call("GetCurrentState", 0).Store(&res)
	return err

}

func (m *TVManager) DisableScreen() error {
	return nil
}
