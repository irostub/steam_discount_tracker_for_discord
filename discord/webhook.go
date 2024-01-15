package discord

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"steam_discount_tracker_for_discord/config"
	"steam_discount_tracker_for_discord/steam/client"
	"strconv"
	"time"
)

type Footer struct {
	Text string `json:"text"`
}

type Image struct {
	URL string `json:"url"`
}

type GameEmbed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Color       *int   `json:"color"`
	Footer      Footer `json:"footer"`
	Image       Image  `json:"image"`
}

type Embed struct {
	Username  string      `json:"username"`
	AvatarUrl string      `json:"avatar_url"`
	Embeds    []GameEmbed `json:"embeds"`
}

func SendWebhook(conf *config.Config, items []client.Items) {
	var gameEmbeds []GameEmbed
	colorCode := conf.Color
	for _, item := range items {
		tm := time.Unix(item.DiscountExpiration, 0)
		embed := GameEmbed{
			Title:       item.Name,
			Description: " [-" + strconv.Itoa(item.DiscountPercent) + "%]" + strconv.FormatInt(item.FinalPrice/100, 10) + conf.CurrencySymbol + " ~~" + strconv.FormatInt(item.OriginalPrice/100, 10) + conf.CurrencySymbol + "~~ ",
			URL:         "https://store.steampowered.com/app/" + strconv.Itoa(item.ID),
			Color:       &colorCode,
			Footer: Footer{
				Text: "할인 종료 일 : " + tm.Format("2006-01-02 15:04:05"),
			},
			Image: Image{
				URL: item.HeaderImage,
			},
		}
		gameEmbeds = append(gameEmbeds, embed)
	}

	embed := Embed{
		Username:  "스팀 세일 알람",
		AvatarUrl: "https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Steam_icon_logo.svg/480px-Steam_icon_logo.svg.png",
		Embeds:    gameEmbeds,
	}

	marshal, err := json.Marshal(embed)
	if err != nil {
		log.Fatal(err)
		return
	}
	requestBody := bytes.NewBuffer(marshal)
	resp, err := http.Post(conf.WebhookUrl, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(body)
	log.Println(s)
}
