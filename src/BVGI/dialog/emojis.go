package dialog

import (
	"math/rand"
	"strings"
)

var smilesMap map[string]string
var smileEmojiReply = []string{"😊", "🙂", "☺", "😀"}

var sadMap map[string]string
var sadEmojiReply = []string{"😞", "😔", "😟"}

func init()  {
	smiles := "😀, 😁, 😃, 😄, 😆, 😉, 😊, 🙂, 😘, 😗, 😙, 😚, 😜, 😝, 😛,haha, hic hic, hichic, hihi, hehe, yep, yes, =)), (y), :v, =]], =], :‑), :), :-], :], :-3, :3, :->, :>, 8-), 8), :-}, :}, :o), :c), :^), =], =), :‑D, :D, 8‑D, 8D, x‑D, xD, X‑D, XD, =D, =3, B^D, :'‑), :'), :-*, :*, :×, ;‑), ;), *-), *), ;‑], ;], ;^), :‑, ;D, :‑P, :P, X‑P, XP, x‑p, xp, :‑p, :p, :‑Þ, :Þ, :‑þ, :þ, :‑b, :b, d:, =p, >:P, O:‑), O:), 0:‑3, 0:3, 0:‑), 0:), 0;^)"
	smilesMap = make(map[string]string)
	for _, v := range strings.SplitN(smiles, ", ", -1) {
		smilesMap[v] = "smile"
	}

	sad := "😏, 😐, 😑, 😒, 😞, 😟, 😡, 😔, 😕, 😣, 😖, 😫, 😩, 😤, 😓, 😪, 😢, 😥, 😭, huhu, huu, :(, :((, =(, =((, :|, :‑(, :(, :‑c, :c, :‑<, :<, :‑[, :[, :-||, >:[, :{, :@, >:(, :'‑(, :'(, D‑':, D:<, D:, D8, D;, D=, DX, :$	"
	sadMap = make(map[string]string)
	for _, v := range strings.SplitN(sad, ", ", -1) {
		sadMap[v] = "sad"
	}
}

func GetEmojiResponse(text string) string {
	if s := strings.Replace(text, " ", "", -1); s == "" { // it's a sticker
		return smileEmojiReply[rand.Intn(len(smileEmojiReply))]
	}

	if emo := smilesMap[text]; emo != "" {
		return smileEmojiReply[rand.Intn(len(smileEmojiReply))]
	}

	if emo := sadMap[text]; emo != "" {
		return sadEmojiReply[rand.Intn(len(sadEmojiReply))]
	}

	return ""
}