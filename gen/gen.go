package gen

import (
	"fmt"
	"xwdr/gen-srv-tpl/tmpl"
	"github.com/xlab/treeprint"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type config struct {
	// service name, e.g example
	Alias string
	// project name, e.g xwdr
	Project string
	// group name e.g xwdr
	Group string
	// github.com/xwdr/service-name
	Dir string
	// $GOPATH/src/github.com/xwdr/service-name
	GoDir string
	// $GOPATH
	GoPath string
	// Files
	Files []file
	// back quote
	Quote string
}

type file struct {
	Path string
	Tmpl string
}

func write(c config, file, tmpl string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func create(c config) error {
	if _, err := os.Stat(c.GoDir); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.GoDir)
	}

	t := treeprint.New()
	nodes := map[string]treeprint.Tree{}
	nodes[c.GoDir] = t

	for _, file := range c.Files {
		f := filepath.Join(c.GoDir, file.Path)
		dir := filepath.Dir(f)

		b, ok := nodes[dir]
		if !ok {
			d, _ := filepath.Rel(c.GoDir, dir)
			b = t.AddBranch(d)
			nodes[dir] = b
		}

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		p := filepath.Base(f)
		b.AddNode(p)
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}

	// print tree
	fmt.Println(t.String())
	return nil
}

// 生成独立的微服务项目
func GenerateIndependent(dir string) {

	var goPath string
	var goDir string

	goPath = os.Getenv("GOPATH")
	useGoModule := os.Getenv("GO111MODULE")

	if len(goPath) == 0 {
		fmt.Println("unknown GOPATH")
		return
	}

	// attempt to split path if not windows
	if runtime.GOOS == "windows" {
		goPath = strings.Split(goPath, ";")[0]
	} else {
		goPath = strings.Split(goPath, ":")[0]
	}
	goDir = filepath.Join(goPath, "src", path.Clean(dir))

	// get the group name
	group := func() string {
		paths := strings.Split(dir, "/")//string(os.PathSeparator))
		if len(paths) >= 2 {
			return paths[len(paths)-2]
		}
		return ""
	}

	c := config{
		Dir:     dir,
		Alias:   "",
		GoDir:   goDir,
		Project: filepath.Base(dir),
		Group:   group(),
		Quote:   "`",
		Files: []file{
			{"main.go", tmpl.MainSRV},
			{"conf/app.ini", tmpl.Config},
			{"routers/web.go", tmpl.InitWeb},
			{"routers/router.go", tmpl.InitRouter},
			{"controllers/user.go", tmpl.Controllers},
			{"controllers/common.go", tmpl.Common},
			{"models/model.go", tmpl.Models},
			{"models/user.go", tmpl.User},
			{"service/service.go", tmpl.InitService},
			{"utils/utils.go", tmpl.InitUtils},
		},
	}

	// set gomodule
	if useGoModule == "on" || useGoModule == "auto" {
		c.Files = append(c.Files, file{"go.mod", tmpl.Module})
	}

	if err := create(c); err != nil {
		fmt.Println(err)
	}
}

