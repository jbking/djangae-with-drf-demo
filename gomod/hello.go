package hello

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type Poll struct {
	ID int64
	Question  string `datastore:"question"`
	PublishedDate    time.Time `datastore:"pub_date"`
	Choices []Choice
}

type Choice struct {
	Choice string `datastore:"choice"`
	PollId int `datastore:"poll_id"`
	Votes int `datastore:"votes"`
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("polls_poll").Order("-pub_date")
	polls := make([]Poll, 0, 10)
	keys, err := q.GetAll(c, &polls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, _ := range polls {
		polls[i].ID = keys[i].IntID()
		q := datastore.NewQuery("polls_choice").Filter("poll_id =", polls[i].ID)
		choices := make([]Choice, 0, 10)
		if _, err := q.GetAll(c, &choices); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		polls[i].Choices = choices
	}
	 if err := pollTemplate.Execute(w, polls); err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
	 }
}

var pollTemplate = template.Must(template.New("interface").Parse(`
<html>
    <head>
        <title>Go Poll Interface</title>
    </head>
    <body>
        {{range .}}
            <p class="question">{{.Question}}</p>
            <p>{{.PublishedDate}}</p>
            <ol>
            {{range .Choices}}
            <li>{{.Choice}}: {{.Votes}}</li>
            {{end}}
            </ol>
        {{end}}
    </body>
</html>
`))
