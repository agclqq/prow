# prow
## intro
prow is a go project scaffolding guide project. It relies directly on the prow-framework.It helps you quickly generate the agreed project structure and basic functionality out of the box.

## useage
### useage 1

```bash
git clone https://github.com/agclqq/prow.git
cd prow
go mod tidy
go run main.go
# Make sure you change it to your git repo address
git remote remove origin
git remote add origin ${your git repo address}
```

### useage 2

```bash
your_project_name=prow && \
mkdir ${your_project_name} && \
cd ${your_project_name} && \
go mod init ${your_project_name} && \
echo 'package main

import (
	"os"

	"github.com/agclqq/prow-framework/prowjob/register"
	"github.com/agclqq/prowjob"
)

func main() {
	pj := prowjob.New()
	register.Register(pj)
	pj.Run("init:project")
	os.Remove("main.go")
}
'> main.go && \
go mod tidy && \
go run main.go
```