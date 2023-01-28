package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	blogs := []api.Blog{
		{
			Id:    1,
			Title: "Exploring the Capabilities of ChatGPT: The State-of-the-Art Language Model",
			Body: `ChatGPT is a state-of-the-art language model developed by OpenAI. It is based on the transformer architecture, which has proven to be highly effective in natural language processing tasks such as language translation and text summarization.
One of the key advantages of ChatGPT is its ability to generate human-like text. This is due to its large training dataset, which includes a wide range of text from books, articles, and websites. As a result, ChatGPT is able to generate text that is not only grammatically correct, but also coherent and contextually appropriate.
Another advantage of ChatGPT is its ability to generate text in a wide range of styles and formats. For example, it can be used to generate poetry, fiction, news articles, and even code. This makes it a versatile tool for a wide range of applications, such as content creation, language translation, and dialogue systems.
In addition to its natural language generation capabilities, ChatGPT also has the ability to perform other natural language processing tasks such as language understanding, text summarization, and question answering.
Overall, ChatGPT is a powerful language model that has the potential to revolutionize the way we interact with machines and automate a wide range of tasks that require natural language understanding and generation. It will also bring new opportunities to the field of Artificial Intelligence and will be a game changer in industries such as customer service, content creation and language translation.
		`},
		{
			Id:    2,
			Title: "Understanding the Impact of Artificial Intelligence (AI) on Our Lives",
			Body: `Artificial Intelligence, or AI, is a rapidly growing field that is changing the way we live and work. At its core, AI is the simulation of human intelligence in machines that are programmed to think and learn like humans.
One of the key advancements in AI is machine learning, which allows machines to learn from data without being explicitly programmed. This has led to the development of powerful algorithms that can recognize patterns in data, make predictions, and even take actions based on that data.
One of the most well-known applications of AI is in the field of image and speech recognition. AI algorithms are now able to recognize faces, objects, and even voices with a high degree of accuracy. This has led to the development of products such as self-driving cars, virtual assistants, and even smart home devices.
Another important area of AI is natural language processing (NLP), which enables machines to understand and generate human language. This has led to the development of language-based applications such as machine translation, text summarization, and automated customer service.
AI is also being used in a wide range of industries to improve efficiency, accuracy, and decision-making. For example, AI algorithms are being used in healthcare to analyze medical images and assist in diagnostics, in finance to detect fraudulent transactions, and in manufacturing to optimize production processes.
While AI has the potential to bring many benefits, it also raises important ethical questions about how these technologies will be used and the impact they will have on society. It is essential that we continue to research and develop AI in a responsible and ethical manner.
Overall, AI is a rapidly developing field that has the potential to change the way we live and work. It will bring new opportunities and challenges, and it will be important for society to understand and embrace the implications of these technologies.
		`},
		{
			Id:    3,
			Title: "Discovering the Rich Culture and History of Cambodia",
			Body: `Cambodia is a Southeast Asian country known for its rich culture and history. The country is home to the famous Angkor Wat, a UNESCO World Heritage Site and one of the most important archaeological sites in Southeast Asia. The temple complex, built in the 12th century, is a testament to the Khmer Empire's power and architectural achievements.
Cambodia is also known for its vibrant culture and traditional arts. The country has a rich tradition of dance and music, which is showcased in the many festivals and performances held throughout the year. Cambodian traditional dance is characterized by its fluid movements and intricate hand gestures.
The country is also home to a diverse array of ethnic groups, each with their own unique customs and traditions. The Khmer people make up the largest ethnic group in Cambodia and have a rich history and culture that is reflected in the country's art, architecture, and cuisine.
Cambodia's economy is largely based on agriculture, with rice as the main crop. The country also has a growing tourism industry, with visitors coming to see the historic sites, temples, and enjoy its unique culture.
In recent years, Cambodia has made significant progress in terms of economic development and poverty reduction, however, it is still a developing country and faces challenges such as corruption, human rights violations, and political instability.
Despite these challenges, Cambodia remains a fascinating country with a rich culture and history. It's a unique destination that offers visitors a glimpse into the past and an opportunity to experience the beauty of its people and traditions.
Overall, Cambodia is a country with a rich history, culture, and art. It offers a unique and fascinating experience for travelers, and it's an opportunity to explore the ancient temples and monuments, as well as to discover its vibrant culture, customs, and traditions.
		`},
	}
	j, _ := json.Marshal(blogs)
	data := string(j)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
