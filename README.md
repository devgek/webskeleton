# go-webskeleton
skeleton for go web projects
---
This skeleton project contains a basic go web application, using sqlite3 as database and labstack/echo or gorilla/mux for
writing controller and middleware stuff.
It also contains a login form and a simple menu. You can login with "admin/xyz".
# Installation
1. **`go get github.com/devgek/webskeleton`**
3. `cd {your-go-modules}`
4. `mkdir {project}`
5. **`webskeleton bootstrap --type={type} --repository={repository} --user={user} --project={project} --title={projecttitle}`**
6. `cd {project}`
7. **`go run main.go serve`**
