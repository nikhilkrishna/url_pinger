package monitor

import (
	"database/sql"
	"fmt"
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
	Logger    logger.Logger
	Db        *sql.DB
}

func NewWebsiteMonitor(configs []*config.WebsiteConfig, sessionID string, logger logger.Logger, db *sql.DB) *WebsiteMonitor {
	return &WebsiteMonitor{
		Configs:   configs,
		SessionID: sessionID,
		Logger:    logger,
		Db:        db,
	}
}

func (wm *WebsiteMonitor) Start() {
	for i, cfg := range wm.Configs {
		go wm.checkWebsite(cfg, i)
	}
}

func (wm *WebsiteMonitor) checkWebsite(cfg *config.WebsiteConfig, threadId int) {
	ticker := time.NewTicker(time.Duration(cfg.Interval) * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log := PingWebsite(cfg, wm.SessionID, threadId)
		wm.Logger.Log(fmt.Sprintf("Session %s, Thread %d: %v", wm.SessionID, threadId, log))

		// Save the log to the database
		if err := database.SaveLog(wm.Db, log); err != nil {
			wm.Logger.Log(fmt.Sprintf(`Error saving log to database: %s`, err.Error()))
		}
	}
}

func PingWebsite(cfg *config.WebsiteConfig, sessionId string, threadId int) database.WebsiteLog {
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

	responseText := string(body)
	if len(responseText) > 500 {
		responseText = responseText[:500] + "..." 
	}

	if cfg.Pattern != "" {
		matched, _ := regexp.MatchString(cfg.Pattern, string(body))
		if matched {
			log.Response = fmt.Sprintf("Pattern matched, %s ", responseText)
		} else {
			log.Response = fmt.Sprintf("Pattern not matched, %s ", responseText)
		}
	} else {
		log.Response = fmt.Sprintf("No pattern provided, %s ", responseText)
	}

	return log
}
