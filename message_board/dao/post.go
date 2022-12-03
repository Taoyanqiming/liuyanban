package dao

import "message_board/model"

func InsetPost(post model.Post) error {
	_, err := DB.Exec("INSERT INTO post(id,username,text)"+"value(?,?,?;", post.Id, post.UserName, post.Text)
	return err
}
func ChangePost(post model.Post) error {
	_, err := DB.Exec("DELETE FROM post(text)"+"values(?;", post.Text)
	return err
	_, err1 := DB.Exec("INSERT INTO post(txt) "+"value(?;", post.Text)
	return err1
}
func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	row, err := DB.Query("SELECT id,username,txt,post_time,update_time,comment_num FROM post")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var post model.Post
		err = row.Scan(&post.Id, &post.UserName, &post.Text)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, nil
}
func SelectPostByPostId(postId int) (model.Post, error) {
	var post model.Post
	row := DB.QueryRow("SELECT id,username,txt,post_time,update_time,comment_num FROM post WHERE id =?")
	if row.Err() != nil {
		return post, row.Err()
	}
	err := row.Scan(&post.Id, &post.UserName, &post.Text)
	if err != nil {
		return post, err
	}
	return post, nil
}
