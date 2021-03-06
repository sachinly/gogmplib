# gogmplib

Library to interact with the Greenbone Vulnerability Manager using the [gmp protocol](https://docs.greenbone.net/API/GMP/gmp-9.0.html) (Greenbone Management Protocol, `version 9.0`)

Here you willl find methods to create tasks, targets, scanners and retrieve results generated by the Openvas scanner.

* This library is compatible with the latest version of GVM: 11.0

# Reference

https://docs.greenbone.net/API/GMP/gmp-9.0.html

https://github.com/filewalkwithme/go-gmp

# Sample Usage

```
import (
	"gogmplib/pkg/9/gmp"
	"gogmplib/pkg/9/gmp/client"
	"gogmplib/pkg/9/gmp/connections"
	"gogmplib/pkg/utils"
)
```

```
// Connect to GVMD
conn, err := connections.NewUnixConnection("/tmp/openvas-socks/gvmd.sock")
if err != nil {
    panic(err)
}
defer conn.Close()

// Instantiate  a new GMP Client
gmpClient := client.New(conn)

// Authenticate
auth := &gmp.AuthenticateCommand{}
auth.Credentials.Username = "openvas"
auth.Credentials.Password = "openvas"
_, err = gmpClient.Authenticate(auth)
if err != nil {
    panic(err)
}

// Create a new task
newTask := &gmp.CreateTaskCommand{}
newTask.Name = "New Task"

newTask.Config = new(gmp.CreateTaskConfig)
newTask.Config.ID = "b9407b88-7b3c-47f3-a684-3605db80e5fd"

newTask.Target = new(gmp.CreateTaskTarget)
newTask.Target.ID = "a2a964f8-daa7-4e1c-ae3a-a72f06d49dbe"

newTask.Scanner = new(gmp.CreateTaskScanner)
newTask.Scanner.ID = "8abd321a-2eb1-4a7a-a368-fc118dc99a85"

newTaskResp, err := gmpClient.CreateTask(newTask)
if err != nil {
    panic(err)
}

// Start the task
st := &gmp.StartTaskCommand{}
st.TaskID = newTaskResp.ID
_, err = gmpClient.StartTask(st)
if err != nil {
    panic(err)
}
```


Confirm that GSA is available at https://127.0.0.1

Next, we will execute the sample application avaible under the `examples` folder:
```
cd examples
go build
./examples
```

The sample application will generate an output like this:

```
Connecting to GVMD...
-> ok
Instantiate  a new GMP Client
-> ok
Authenticating...
-> ok
Getting the Default scanner...
-> ok (scanner id: 08b69003-5fc2-4037-a479-93b440211c73)
Getting the configuration named "Full and fast"...
-> ok (config id: daba56c8-73ec-11df-a475-002264764cea)
Creating a new target...
-> ok (target id: 34e34484-2dfa-4f01-b9fe-2b90e4a38fba)
Creating a new task...
-> ok (task id: d82a368f-acfa-422e-a78a-63727861f5a9)
Start the task

Monitoring task progress...
Monitoring task progress: 1%
Monitoring task progress: 2%
Monitoring task progress: 6%
Monitoring task progress: 8%
Monitoring task progress: 14%
Monitoring task progress: 52%
Monitoring task progress: 80%
Monitoring task progress: 90%
Monitoring task progress: 98%
Monitoring task progress: 100%


Result[0]: CGI Scanning Consolidation (score: 0.0)
...
Result[7]: HTTP Security Headers Detection (score: 0.0)
...
Result[15]: Services (score: 0.0)
...
Result[20]: SSL/TLS: Diffie-Hellman Key Exchange Insufficient DH Group Strength Vulnerability (score: 4.0)
Result[37]: SSL/TLS: Report Vulnerable Cipher Suites for HTTPS (score: 5.0)
Result[38]: SSL/TLS: Report Weak Cipher Suites (score: 4.3)
Result[39]: Traceroute (score: 0.0)
Result[40]: Unknown OS and Service Banner Reporting (score: 0.0)
```


HTML Report:
```
go tool cover -html=coverage.out
```
