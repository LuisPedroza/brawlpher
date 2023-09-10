package bs

import (
	"fmt"

	"github.com/LuisPedroza/brawlpher/internal"
)

type BrawlerClient struct {
	c *internal.Client
}

func (c *BrawlerClient) GetBrawlerByID(brawlerID string) (*Brawler, error) {
	var brawler *Brawler
	if err := c.c.GetInto(fmt.Sprintf(brawlerById, brawlerID), &brawler); err != nil {
		return nil, err
	}
	return brawler, nil
}

func (c *BrawlerClient) GetAllBrawlers() (*Brawlers, error) {
	var brawlers *Brawlers
	if err := c.c.GetInto(allBrawlers, &brawlers); err != nil {
		return nil, err
	}
	return brawlers, nil
}
