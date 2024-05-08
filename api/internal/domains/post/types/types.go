package types

const (
	ReactionLike = iota
	ReactionDislike
	ReactinHeart
	ReactionFire
	ReactionStone
	ReactionThinking
)

const NumbererOfReactions = 6

type Post struct {
	ID          int        `json:"id" db:"post_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description,omitempty" db:"description"`
	Cover       string     `json:"cover" db:"cover"`
	Content     string     `json:"content,omitempty" db:"content"`
	Reactions   []Reaction `json:"reactions" db:"reactions"`
	Tags        []string   `json:"tags" db:"tags"`
	CreatedAt   int        `json:"created_at" db:"created_at"`
	Lang        string     `json:"lang,omitempty" db:"lang"`

	Tag string `json:"-" db:"tag"`
}

type Reaction struct {
	Type   int `json:"type" db:"type"`
	Number int `json:"number" db:"number"`
}
