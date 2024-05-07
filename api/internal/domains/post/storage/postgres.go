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

func (store *PostStorage) SelectPost(post_type string, title string, id string, lang string, tags []string, user_agent string, offset uint64) ([]types.Post, error) {
	where := sq.Eq{}
	if title != "" {
		where["title"] = title
	}
	if id != "" {
		where["post_id"] = id
	}
	if lang != "" {
		where["lang"] = lang
	}
	if len(tags) > 0 {
		where["tags"] = tags
	}
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("posts.post_id, title, cover, content, created_at").From("posts").
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

		posts = append(posts, post)
	}

	return posts, nil
}

func (store *PostStorage) InsertPost(langs []types.Post) error {
	tx, err := store.DB.Begin()
	if err != nil {
		return err
	}
	var post_id int
	row := tx.QueryRow("INSERT INTO posts VALUES(DEFAULT, $1, DEFAULT, DEFAULT, 1) RETURNING post_id;", langs[0].Cover)
	if err := row.Scan(&post_id); err != nil {
		return err
	}
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_lang").Columns("post_id", "lang", "title", "content")
	for _, lang := range langs {
		query = query.Values(post_id, lang.Lang, lang.Title, lang.Content)
	}
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return err
	}

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_tag").Columns("post_id", "tag")
	for _, tag := range langs[0].Tags {
		query = query.Values(post_id, tag)
	}
	sql, args, err = query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return err
	}

	query = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Insert("post_reaction").Columns("post_id", "type", "number")
	for i := range types.NumbererOfReactions {
		query = query.Values(post_id, i, 0)
	}
	sql, args, err = query.ToSql()
	if err != nil {
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return err
	}

	return tx.Commit()

}
