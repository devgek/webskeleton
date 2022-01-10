# go-webskeleton
skeleton for go web projects

This skeleton project contains a basic go web application, using sqlite3 as database and labstack/echo for
writing controller and middleware stuff.
It also contains a login form and a simple menu. You can login with "admin/xyz".

---
## Installation
1. `go get github.com/devgek/webskeleton`

---
## Installation
1. `go get github.com/devgek/webskeleton`
2. `cd {your-go-modules}`
3. `mkdir {project}`
4. `webskeleton bootstrap --type={type} --repository={repository} --user={user} --project={project} --title={projecttitle}
5. `cd {project}`
6. `go run main.go serve --config=_test/config-serve.yaml`
