**! Work in Progress !**

# SaveLoad (name TBD)

Yes, this is uncreatively named.

## Setup Instructions

Make sure you have [dep](https://github.com/golang/dep)
and [retool](https://github.com/twitchtv/retool) installed

- `dep ensure`
- `retool sync`

## Running Tests

- `retool do ginkgo -r `

## Motivation

> Closures == Generics

Can we create a database persistence layer in golang that avoids
reflection and code generation?

Approach:

- Datamapper / Data Access Layer pattern
- Closures to group database column names with values
- Closures to yield pointer to struct fields, so we can set values from DB
 
The user declaratively defines the mapping between struct and DB row
with something like:

```
func NewMessageSL(msg *Message) *SaveLoad {
	sl := NewSaveLoad("messages")
	sl.Int("id", msg.ID)
	sl.String("title", msg.Title)
	sl.String("body", msg.Body)
	sl.Time("created_at", msg.CreatedAt)

	return sl
}
```

The one restriction we may need to accept is that all struct fields must
be pointers, so we can use closures to save a reference to the field of the
struct that we passed in.

# Further Questions

- How should we handle nulls in the database, eg deliberately blank value;
  versus nulls in our struct, eg could be unintentionally missing
