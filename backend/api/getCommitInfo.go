package api

import (
	"errors"
	"net/http"
	"os"
	"path"

	"gihachi/repository-research/api/json"

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

		errJSON := &json.Error{
			ErrorStr: err.Error(),
		}
		return c.JSON(http.StatusNotFound, errJSON)
	}

	parentCommit, _ := commit.Parent(0)

	commitTree, _ := commit.Tree()
	parentCommitTree, _ := parentCommit.Tree()

	treeDiff, _ := parentCommitTree.Patch(commitTree)

	diffJSON := &json.Diff{
		DiffStr: treeDiff.String(),
	}

	return c.JSON(http.StatusOK, diffJSON)
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
