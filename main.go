package main

import (
	"fmt"
	"github.com/ladydascalie/sortdir/sortdir"
	"github.com/ghthor/journal/git"
	"os/exec"
	"time"
)


func main() {
	r := Repository{
		Path:sortdir.Pwd(),
	}

	if r.isClean() == true {
		fmt.Println("Repository is clean!")
		return
	}

	branch := createBranch()

	commitAllFiles(branch)
}

type Repository struct {
	Path string
	Status bool
}

func (r *Repository) isClean() bool {
	err := git.IsClean(r.Path)
	if err == nil {
		fmt.Println("Directory is clean!")
		r.Status = true
		return true
	}

	r.Status = false
	return false
}

func makeGitUsableDate() string {
	y := time.Now().Year()
	M := time.Now().Month()
	d := time.Now().Day()
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()

	return fmt.Sprintf("%d_%d_%d_%d_%d_%d", y, M, d, h, m, s)
}

func createBranch() string {
	branchName := fmt.Sprintf("GitFire_%s", makeGitUsableDate())
	cmd := exec.Command("git", "checkout", "-b", branchName)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out));

	return branchName
}

func commitAllFiles(branch string) {
	err := exec.Command("git", "add", ".").Run()
	if err != nil {
		panic(err)
	}

	err = exec.Command("git", "commit", "-m", "Emergency Commit From GitFire!").Run()
	if err != nil {
		panic(err)
	}

	err = exec.Command("git", "push", "-u", "origin", branch).Run()
	if err != nil {
		panic(err)
	}
}
