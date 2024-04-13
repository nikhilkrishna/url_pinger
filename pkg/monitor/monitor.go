package monitor

import (
	"io"
	"net/http"
	"regexp"
	"time"
	"url_pinger/pkg/config"
	"url_pinger/pkg/database"
    "url_pinger/pkg/logger"
)

type WebsiteMonitor struct {
	Configs   []*config.WebsiteConfig
	SessionID string
}

func NewWebsiteMonitor(configs []*config.WebsiteConfig, sessionID string) *WebsiteMonitor {
	return &WebsiteMonitor{
		Configs:   configs,
		SessionID: sessionID,
	}
}

func (wm *WebsiteMonitor) Start() {
	for i, cfg := range wm.Configs {
		go wm.checkWebsite(cfg, i)
	}
}

func (wm *WebsiteMonitor) checkWebsite(cfg *config.WebsiteConfig, threadId int) {
	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			log := CheckWebsite(cfg, wm.SessionID, threadId)
			// Implement a logging mechanism here instead of fmt.Println
			// Other cases like stop signal can be added
		}
	}
}

func CheckWebsite(cfg *config.WebsiteConfig, sessionId string, threadId int) database.WebsiteLog {
	log := database.WebsiteLog{
		SessionId: sessionId,
		ThreadId:  threadId,
		URL:       cfg.URL,
		Pattern:   cfg.Pattern,
	}

	resp, err := http.Get(cfg.URL)
	if err != nil {
		log.Error = err.Error()
		return log
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error = err.Error()
		return log
	}

	if cfg.Pattern != "" {
		matched, _ := regexp.MatchString(cfg.Pattern, string(body))
		if matched {
			log.Response = "Pattern matched"
		} else {
			log.Response = "Pattern not matched"
		}
	} else {
		log.Response = "No pattern provided"
	}

	return log
}
