package main

import (
	"github.com/CarlosRojas316/mirroring-web-crawler/config"
	"github.com/CarlosRojas316/mirroring-web-crawler/crawler"
	"go.uber.org/zap"

	"testing"
)

func TestHappyPath(t *testing.T) {
	// Load configuration settings
	cfg, _ := config.Load()
	// if err != nil {
	// zap.L().Fatal("failed to start config", zap.Error(err))
	// }

	cfg.StartURL = "https://spaces.w3schools.com/"
	cfg.DirectoryPath = "E:/crawl"

	// Initialize crawler
	c := crawler.NewCrawler(cfg)
	c.Start()

	// Log summary of the crawl
	zap.L().Info("Crawling process exit.")
}
