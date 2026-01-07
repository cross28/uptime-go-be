# Uptime

This is a backend to the CrossSystems™ Uptime frontend system. 

I attempted to implement this project using vertical slice architecture, and the repository pattern as much as possible. I wanted to emulate this behavior as an homage from my .NET development.

To get running:
1. First run `go mod tidy`
2. Then run `make run` to start the application. `-env (dev|qa|prod)` and `-port` are optional cmd parameters that can be passed in.