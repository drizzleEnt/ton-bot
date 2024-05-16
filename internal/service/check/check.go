package check

import (
	"context"

	desc "github.com/drizzleent/ton-bot/pkg/scraper"
)

func (s *CheckService) CheckGpu(str string) (string, error) {
	client := desc.NewScraperClient(s.conn)
	ctx := context.Background()
	resp, err := client.Scrap(ctx, &desc.ScrapUrl{Url: str})
	if err != nil {
		return "", err
	}
	return resp.Text, nil
}
