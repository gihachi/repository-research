package api

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// GetDiff リポジトリのdiffを取得
func GetDiff(c echo.Context) error {

	basePath := os.Getenv("WORKING_DIR")
	repoName := c.QueryParam("repo")
	repoPath := path.Join(basePath, "repository", repoName)

	existPath := isExistPath(repoPath)

	if !existPath {
		return c.JSON(404, gin.H{
			"error": "repo not found",
		})
	}

	r, err := git.PlainOpen(repoPath)

	if err != nil {
		return c.JSON(404, gin.H{
			"error": "repo not found",
		})
	}

	commitHash := c.QueryParam("commit")
	commit, err := r.CommitObject(plumbing.NewHash(commitHash))

	if err != nil {
		return c.JSON(404, gin.H{
			"error": "commit not found",
		})
	}
	parentCommit, _ := commit.Parent(0)

	commitTree, _ := commit.Tree()
	parentCommitTree, _ := parentCommit.Tree()

	treeDiff, _ := parentCommitTree.Patch(commitTree)

	return c.JSON(200, gin.H{
		"repo": repoName,
		"diff": treeDiff.String(),
	})
}

func isExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
