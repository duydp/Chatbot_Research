package util

import (
	"strings"
	"regexp"

	"github.com/michlabs/fbbot"
)

func SendTextWithImages(bot *fbbot.Bot, u fbbot.User, text string){
	if strings.Contains(text, "@image") {
		r := regexp.MustCompile("@image\\[[^\\s/$.?#].[^\\s]*\\]")
		
		arr := r.FindAllStringSubmatchIndex(text, -1)
		
		index := 0
		for _, v := range arr {
			if v[0] >= index {
				subtext := text[index:v[0]]	
				if len(subtext) > 0 {
					bot.SendText(u, subtext)	
				}				
				
				subImage := text[v[0]:v[1]]
				subImage = subImage[len("@image["): len(subImage)-1]
				if len(subImage) > 0 {
					bot.SendImage(u, subImage)
				}
				
				index=v[1]
			}
		}

		if index < len(text) {
			bot.SendText(u, text[index:])
		}
	} else {
		bot.SendText(u, text)
	}
	
}
