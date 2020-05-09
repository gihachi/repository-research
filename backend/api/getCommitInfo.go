package api

import (
	"errors"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GetDiff リポジトリのdiffを取得
func GetDiff(c echo.Context) error {

	repoName := c.QueryParam("repo")
	commitHash := c.QueryParam("commit")

	commit, err := getCommit(repoName, commitHash)
	if err != nil {
		return c.JSON(404, gin.H{
			"message": err.Error(),
		})
	}

	parentCommit, _ := commit.Parent(0)

	commitTree, _ := commit.Tree()
	parentCommitTree, _ := parentCommit.Tree()

	treeDiff, _ := parentCommitTree.Patch(commitTree)

	return c.JSON(200, gin.H{
		"diff": treeDiff.String(),
	})
}

func getCommit(repoName string, commitHash string) (*object.Commit, error) {

	basePath := os.Getenv("WORKING_DIR")
	repoPath := path.Join(basePath, "repository", repoName)

	r, err := git.PlainOpen(repoPath)

	if err != nil {
		return nil, errors.New("repo not found")
	}

	commit, err := r.CommitObject(plumbing.NewHash(commitHash))

	if err != nil {
		return nil, errors.New("commit not found")
	}

	return commit, nil
}

func isExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
