package storage

import (
	"github.com/Corray333/portfolio/internal/domains/post/types"
	store "github.com/Corray333/portfolio/internal/storage"
	sq "github.com/Masterminds/squirrel"
)

type PostStorage struct {
	store.Storage
}

func New(global store.Storage) *PostStorage {
	return &PostStorage{
		global,
	}
}

func (store *PostStorage) SelectPosts(post_type string, title string, id string, lang string, tags []string, user_agent string, offset uint64) ([]types.Post, error) {
	where := sq.Eq{}
	if title != "" {
		where["title"] = title
	}
	if id != "" {
		where["posts.post_id"] = id
	}
	if lang != "" {
		where["lang"] = lang
	}
	if len(tags) > 0 {
		where["tags"] = tags
	}
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("posts.post_id, title, description, cover, created_at, tag").From("posts").
		Join("post_lang ON post_lang.post_id = posts.post_id").
		LeftJoin("post_tag ON post_tag.post_id = posts.post_id").
		Where(where).
		Offset(offset * 50).Limit(50)
	strQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := store.DB.Queryx(strQuery, args...)
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}

	var post types.Post
	for rows.Next() {
		err = rows.StructScan(&post)
		if err != nil {
			return nil, err
		}

		if len(posts) > 0 && post.ID == posts[len(posts)-1].ID {
			posts[len(posts)-1].Tags = append(posts[len(posts)-1].Tags, post.Tag)
			continue
		} else {
			post.Tags = []string{post.Tag}
		}

		query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select(`type, number`).From("post_reaction").Where("post_id =?", post.ID)
		strQuery, args, err = query.ToSql()
		if err != nil {
			return nil, err
		}
		reaction_rows, err := store.DB.Queryx(strQuery, args...)
		if err != nil {
			return nil, err
		}
		var reaction types.Reaction
		for reaction_rows.Next() {
			err = reaction_rows.StructScan(&reaction)
			if err != nil {
				return nil, err
			}
			post.Reactions = append(post.Reactions, reaction)
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (store *PostStorage) SelectPost(post_id string, lang string) (*types.Post, error) {
	if lang == "" {
		lang = "eng"
	}
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("posts.post_id, title, description, cover, content, tag, created_at").From("posts").
		Join("post_lang ON post_lang.post_id = posts.post_id").
		LeftJoin("post_tag ON post_tag.post_id = posts.post_id").
		Where("posts.post_id = ?", post_id).Where("lang =?", lang)
	strQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := store.DB.Queryx(strQuery, args...)
	if err != nil {
		return nil, err
	}

	post := types.Post{}
	tempPost := types.Post{}
	for rows.Next() {
		err = rows.StructScan(&tempPost)
		if err != nil {
			return nil, err
		}

		if post.ID == 0 {
			post = tempPost
		}
		post.Tags = append(post.Tags, tempPost.Tag)

	}

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select(`type, number`).From("post_reaction").Where("post_id =?", post.ID)
	strQuery, args, err = query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err = store.DB.Queryx(strQuery, args...)
	if err != nil {
		return nil, err
	}
	var reaction types.Reaction
	for rows.Next() {
		err = rows.StructScan(&reaction)
		if err != nil {
			return nil, err
		}
		post.Reactions = append(post.Reactions, reaction)
	}

	return &post, nil
}

// func (store *PostStorage) SelectPostAllLangs(post_id string) ([]*types.Post, error) {

// 	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("posts.post_id, title, description, cover, content, tag, created_at").From("posts").
// 		Join("post_lang ON post_lang.post_id = posts.post_id").
// 		LeftJoin("post_tag ON post_tag.post_id = posts.post_id").
// 		Where("posts.post_id = ?", post_id)
// 	strQuery, args, err := query.ToSql()
// 	if err != nil {
// 		return nil, err
// 	}

// 	rows, err := store.DB.Queryx(strQuery, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	posts := []types.Post{}
// 	tempPost := types.Post{}
// 	for rows.Next() {
// 		err = rows.StructScan(&tempPost)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if post.ID == 0 {
// 			post = tempPost
// 		}
// 		post.Tags = append(post.Tags, tempPost.Tag)

// 	}

// 	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select(`type, number`).From("post_reaction").Where("post_id =?", post.ID)
// 	strQuery, args, err = query.ToSql()
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err = store.DB.Queryx(strQuery, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var reaction types.Reaction
// 	for rows.Next() {
// 		err = rows.StructScan(&reaction)
// 		if err != nil {
// 			return nil, err
// 		}
// 		post.Reactions = append(post.Reactions, reaction)
// 	}

// 	return &post, nil
// }

func (store *PostStorage) InsertPost(langs []types.Post) (int, error) {
	tx, err := store.DB.Begin()
	if err != nil {
		return 0, err
	}
	var post_id int
	row := tx.QueryRow("INSERT INTO posts VALUES(DEFAULT, $1, DEFAULT, DEFAULT, 1) RETURNING post_id;", langs[0].Cover)
	if err := row.Scan(&post_id); err != nil {
		return 0, err
	}
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_lang").Columns("post_id", "lang", "title", "description", "content")
	for _, lang := range langs {
		query = query.Values(post_id, lang.Lang, lang.Title, lang.Description, lang.Content)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_tag").Columns("post_id", "tag")
	for _, tag := range langs[0].Tags {
		query = query.Values(post_id, tag)
	}
	sql, args, err = query.ToSql()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_reaction").Columns("post_id", "type", "number")
	for i := range types.NumbererOfReactions {
		query = query.Values(post_id, i, 0)
	}
	sql, args, err = query.ToSql()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	return post_id, tx.Commit()

}

func (store *PostStorage) UpdateReaction(post_id string, increment, decriment int) error {
	tx, err := store.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE post_reaction SET number = number + 1 WHERE post_id = $1 AND type = $2", post_id, increment)
	if err != nil {
		return err
	}

	if decriment != 0 {
		_, err = tx.Exec("UPDATE post_reaction SET number = number - 1 WHERE post_id = $1 AND type = $2", post_id, decriment)
		if err != nil {
			return err
		}
	}

	return tx.Commit()

}
