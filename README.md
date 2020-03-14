# go-webskeleton
skeleton for go web projects
---
This skeleton project contains a basic go web application, using sqlite3 as database and labstack/echo or gorilla/mux for
writing controller and middleware stuff.
It also contains a login form and a simple menu. You can login with "admin/xyz".
# Installation
1. **`go get -u github.com/devgek/webskeleton`**
2. `cd $GOPATH/src/{repository}/{user}`
3. `mkdir {project}`
4. **`webskeleton bootstrap --repository={repository} --user={user} --project={project} --web=[echo|mux]`**
5. `cd $GOPATH/src/{repository}/{user}/{project}`
6. **`go run main.go`**
