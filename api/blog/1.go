package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sokhuong-uon/go-vercel-serverless/api"
)

func Blog1Handler(w http.ResponseWriter, r *http.Request) {

	blog := api.Blog{
		Title: "Exploring the Capabilities of ChatGPT: The State-of-the-Art Language Model",
		Body: `ChatGPT is a state-of-the-art language model developed by OpenAI. It is based on the transformer architecture, which has proven to be highly effective in natural language processing tasks such as language translation and text summarization.
One of the key advantages of ChatGPT is its ability to generate human-like text. This is due to its large training dataset, which includes a wide range of text from books, articles, and websites. As a result, ChatGPT is able to generate text that is not only grammatically correct, but also coherent and contextually appropriate.
Another advantage of ChatGPT is its ability to generate text in a wide range of styles and formats. For example, it can be used to generate poetry, fiction, news articles, and even code. This makes it a versatile tool for a wide range of applications, such as content creation, language translation, and dialogue systems.
In addition to its natural language generation capabilities, ChatGPT also has the ability to perform other natural language processing tasks such as language understanding, text summarization, and question answering.
Overall, ChatGPT is a powerful language model that has the potential to revolutionize the way we interact with machines and automate a wide range of tasks that require natural language understanding and generation. It will also bring new opportunities to the field of Artificial Intelligence and will be a game changer in industries such as customer service, content creation and language translation.
		`}
	j, _ := json.Marshal(blog)
	data := string(j)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%v", data)
}
