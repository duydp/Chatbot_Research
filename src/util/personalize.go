package util

import(
	"strings"

	"github.com/michlabs/fbbot"
)

// Personalize personalizes a text content for an user
// Supported tags: @Gender, @gender, @first_name, @last_name, @full_name
func Personalize(text string, u *fbbot.User) string {
	if strings.Contains(text, "@Gender") || strings.Contains(text, "@gender") {
		switch u.Gender(){
		case "male":
			text = strings.Replace(text, "@Gender", "Anh", -1)
			text = strings.Replace(text, "@gender", "anh", -1)
		case "female":
			text = strings.Replace(text, "@Gender", "Chị", -1)
			text = strings.Replace(text, "@gender", "chị", -1)
		default:
			text = strings.Replace(text, "@Gender", "Anh/Chị", -1)
			text = strings.Replace(text, "@gender", "anh/chị", -1)
		}
	}
	
	if strings.Contains(text, "@first_name") {
		text = strings.Replace(text, "@first_name", u.FirstName(), -1)		
	}

	if strings.Contains(text, "@last_name") {
		text = strings.Replace(text, "@last_name", u.LastName(), -1)		
	}

	if strings.Contains(text, "@full_name") {
		// Minhnt: lay theo name -> fullname
		text = strings.Replace(text, "@full_name", u.Name(), -1)
	}

	return text
}