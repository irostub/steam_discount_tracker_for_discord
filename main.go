package main

import (
	"flag"
	"steam_discount_tracker_for_discord/config"
	"steam_discount_tracker_for_discord/discord"
	"steam_discount_tracker_for_discord/steam/client"
	"sync"
	"time"
)

type Cache struct {
	Items  Items
	Client *client.Client
}

type Items map[int]client.Items

var lock = &sync.Mutex{}
var cache *Cache
var conf *config.Config

func (cache Cache) updateAndSend() {
	responseItems := cache.Client.Get().Specials.Items
	for _, responseItem := range responseItems {
		if responseItem.DiscountExpiration <= 0 {
			continue
		}
		_, ok := cache.Items[responseItem.ID]
		if !ok {
			cache.Items[responseItem.ID] = responseItem
			var newTrackingItems []client.Items
			newTrackingItems = append(newTrackingItems, responseItem)
			discord.SendWebhook(conf, newTrackingItems)
		}
	}
}

func (cache Cache) clearExpired() {
	for key, item := range cache.Items {
		if item.IsExpired() {
			delete(cache.Items, key)
		}
	}
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		for {
			cache.updateAndSend()
			cache.clearExpired()
			time.Sleep(time.Second * time.Duration(conf.CheckCycle))
		}
	}()
	wg.Wait()
}

func init() {
	cache = &Cache{
		Client: client.GetInstance("KR"),
		Items:  make(Items),
	}

	webhookUrl := flag.String("webhook_url", "", "[Required] enter discord webhook url")
	color := flag.Int("color", 15844367, "discord alert line color, default gold color (color code : 15844367)")
	checkCycle := flag.Int("check_cycle", 30, "capture cycle of steam discount list, default 30 seconds")
	currencySymbol := flag.String("currency_symbol", "₩", "currency symbol character, default '₩'")

	flag.Parse()

	conf = &config.Config{
		WebhookUrl:     *webhookUrl,
		Color:          *color,
		CheckCycle:     *checkCycle,
		CurrencySymbol: *currencySymbol,
	}
}
