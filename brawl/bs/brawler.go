package bs

import (
	"fmt"
	"log/slog"

	"github.com/LuisPedroza/brawlpher/api"
	"github.com/LuisPedroza/brawlpher/internal"
)

type BrawlerClient struct {
	c *internal.Client
}

func (c *BrawlerClient) GetBrawlerByID(brawlerID string) (*Brawler, error) {
	var brawler *Brawler
	if err := c.c.GetInto(fmt.Sprintf(brawlerById, brawlerID), &brawler); err != nil {
		c.c.Logger.Debug("GetBrawlerByID",
			slog.String("package", "bs"),
			slog.String("type", "BrawlerClient"),
			slog.String("brawlerID", brawlerID),
			slog.String("error", err.Error()))
		return nil, err
	}
	return brawler, nil
}

func (c *BrawlerClient) GetAllBrawlers(paginationSetters ...api.PaginationOption) (*Brawlers, error) {
	var brawlers *Brawlers
	if err := c.c.GetInto(allBrawlers, &brawlers, paginationSetters...); err != nil {
		c.c.Logger.Debug("GetAllBrawlers",
			slog.String("package", "bs"),
			slog.String("type", "BrawlerClient"),
			slog.String("error", err.Error()))
		return nil, err
	}
	return brawlers, nil
}
