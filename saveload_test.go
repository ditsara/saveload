package saveload_test

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ditsara/saveload"
)

type Message struct {
	ID        *int
	Title     *string
	Body      *string
	CreatedAt *time.Time
}

func NewMessageSL(msg *Message) *SaveLoad {
	sl := NewSaveLoad("messages")
	sl.Int("id", msg.ID)
	sl.String("title", msg.Title)
	sl.String("body", msg.Body)
	sl.Time("created_at", msg.CreatedAt)

	return sl
}

var _ = Describe("Saveload", func() {
	Context("Generate SQL from existing obj", func() {
		It("associates serialized values with DB columns", func() {
			id := 1
			title := "title 1"
			body := "body"
			createdAt, _ := time.Parse(time.RFC3339, "2019-01-01T12:01:01Z")
			msg := Message{ID: &id, Title: &title, Body: &body, CreatedAt: &createdAt}
			msgsl := NewMessageSL(&msg)

			Expect(msgsl.Fields["id"].Val()).To(Equal("1"))
			Expect(msgsl.Fields["title"].Val()).To(Equal("title 1"))
			Expect(msgsl.Fields["body"].Val()).To(Equal("body"))
			Expect(msgsl.Fields["created_at"].Val()).To(Equal("2019-01-01T12:01:01Z"))

			fmt.Println("\n--------------------")
			fmt.Println("-- field values")
			msgsl.Print()
			fmt.Println("-- example SQL")
			msgsl.Save()
			fmt.Println("--------------------")
		})

		It("loads serialized values into new struct", func() {
			// FIXME: I'm not sure what to do about this yet, but the issue is that if
			// a pointer is null, the reference saved in the closure is null, so it won't
			// be able to mutate the struct after it loads from the database
			id := 0
			title := "not the droids"
			body := "you're looking for"
			createdAt := time.Now()
			msg := Message{ID: &id, Title: &title, Body: &body, CreatedAt: &createdAt}
			msgsl := NewMessageSL(&msg)

			// IRL we'll get these keys / values iterating over sql.Row.Columns
			msgsl.Fields["id"].Set("2")
			msgsl.Fields["title"].Set("title 2")
			msgsl.Fields["body"].Set("body 2")
			msgsl.Fields["created_at"].Set("2019-01-01T12:01:01Z")

			fmt.Println("\n--------------------")
			fmt.Println("-- struct after mutation")
			spew.Dump(msg)
			fmt.Println("--------------------")
		})
	})
})
