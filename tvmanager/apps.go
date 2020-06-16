package tvmanager

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/snabb/webostv"
	"net/url"
)

const (
	AppIdYoutube = "youtube.leanback.v4"
	AppIdTwitch  = "tv.twitch.tv.starshot.lg"
	AppIdNetflix = "netflix"
	AppIdPrime   = "amazon"
	AppIdAppleTV = "com.apple.appletv"
)

func (m *TVManager) LaunchApp(id string) error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	_, err := m.tv.ApplicationManagerLaunch(id, nil)
	return err
}

func (m *TVManager) OpenYoutubeURL(u string) error {
	if err := m.checkConnection(); err != nil {
		return err
	}

	parsed, err := url.Parse(u)
	if err != nil {
		return errors.WithMessage(err, "invalid url")
	}

	id := parsed.Query().Get("v")

	_, err = m.tv.ApplicationManagerLaunch(AppIdYoutube, webostv.Payload{"params": map[string]string{"contentTarget": fmt.Sprintf("https://www.youtube.com/tv?v=%s", id)}})
	return err
}
