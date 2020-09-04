package git

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// CountCommit : return number of Commit that <email> create
func CountCommit(path string, counter *int, email string) {
	// instantiate a git repo object from path
	repo, err := git.PlainOpen(path)
	if err != nil {
		panic(err)
	}
	// get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	// get the commits history starting from HEAD
	iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}
	// iterate the commits
	err = iterator.ForEach(func(c *object.Commit) error {
		if (string(c.Author.Email) == email) {
			*counter++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}
