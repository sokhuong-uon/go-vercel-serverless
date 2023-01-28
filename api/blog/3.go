package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog3Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		Id:    3,
		Title: "Discovering the Rich Culture and History of Cambodia",
		Body: `Cambodia is a Southeast Asian country known for its rich culture and history. The country is home to the famous Angkor Wat, a UNESCO World Heritage Site and one of the most important archaeological sites in Southeast Asia. The temple complex, built in the 12th century, is a testament to the Khmer Empire's power and architectural achievements.
Cambodia is also known for its vibrant culture and traditional arts. The country has a rich tradition of dance and music, which is showcased in the many festivals and performances held throughout the year. Cambodian traditional dance is characterized by its fluid movements and intricate hand gestures.
The country is also home to a diverse array of ethnic groups, each with their own unique customs and traditions. The Khmer people make up the largest ethnic group in Cambodia and have a rich history and culture that is reflected in the country's art, architecture, and cuisine.
Cambodia's economy is largely based on agriculture, with rice as the main crop. The country also has a growing tourism industry, with visitors coming to see the historic sites, temples, and enjoy its unique culture.
In recent years, Cambodia has made significant progress in terms of economic development and poverty reduction, however, it is still a developing country and faces challenges such as corruption, human rights violations, and political instability.
Despite these challenges, Cambodia remains a fascinating country with a rich culture and history. It's a unique destination that offers visitors a glimpse into the past and an opportunity to experience the beauty of its people and traditions.
Overall, Cambodia is a country with a rich history, culture, and art. It offers a unique and fascinating experience for travelers, and it's an opportunity to explore the ancient temples and monuments, as well as to discover its vibrant culture, customs, and traditions.
		`}
	j, _ := json.Marshal(blog)
	data := string(j)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
