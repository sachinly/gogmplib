package main

import (
	"fmt"
	"gogmplib/pkg/9/gmp"
	"gogmplib/pkg/9/gmp/client"
	"gogmplib/pkg/9/gmp/connections"
	"gogmplib/pkg/utils"
	"strconv"
	"time"
)

func main() {
	// Connect to GVMD
	fmt.Println("Connecting to GVMD...")
	//conn, err := connections.NewUnixConnection("/tmp/openvas-socks/gvmd.sock")
	conn, err := connections.NewTLSConnection("127.0.0.1:9390", true)
	if err != nil {
		panic(err)
	}
	fmt.Println("-> ok")
	defer conn.Close()

	// Instantiate  a new GMP Client
	fmt.Println("Instantiate  a new GMP Client")
	gmpClient := client.New(conn)
	fmt.Println("-> ok")

	// Authenticate
	fmt.Println("Authenticating...")
	auth := &gmp.AuthenticateCommand{}
	auth.Credentials.Username = "admin"
	auth.Credentials.Password = "openvas"
	authResp, err := gmpClient.Authenticate(auth)
	fmt.Println(authResp)
	if err != nil {
		panic(err)
	}
	fmt.Println("-> ok")

	// Get the Default scanner
	fmt.Println("Getting the Default scanner...")
	s := &gmp.GetScannersCommand{}
	s.Filter = `name="OpenVAS Default"`
	getScannersResp, err := gmpClient.GetScanners(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (scanner id: %s)\n", getScannersResp.Scanner[0].ID)

	// Get the configuration named "Full and fast"
	fmt.Println("Getting the configuration named \"Full and fast\"...")
	c := &gmp.GetConfigsCommand{}
	c.Filter = `name="Full and fast"`
	configResp, err := gmpClient.GetConfigs(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (config id: %s)\n", configResp.Config[0].ID)

	// 获取portlist
	pl := &gmp.GetPortListsCommand{}
	pl.Filter = `name="All TCP and Nmap top 100 UDP"`
	getPortlistResp, err := gmpClient.GetPortLists(pl)
	if err != nil {
		panic(err)
	}
	fmt.Println(getPortlistResp.PortList[0].ID)
	// Create a target
	fmt.Println("Creating a new target...")
	ct := &gmp.CreateTargetCommand{}
	ct.Hosts = "New Target" + "_" + utils.GetRandomString(10)
	ct.Name = ct.Hosts
	//ct.PortRange = "1-65535"
	ct.PortList.ID = getPortlistResp.PortList[0].ID
	createTargetResp, err := gmpClient.CreateTarget(ct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (target id: %s)\n", createTargetResp.ID)

	// Create a new task
	fmt.Println("Creating a new task...")
	newTask := &gmp.CreateTaskCommand{}
	newTask.Name = "New Task" + "_" + utils.GetRandomString(10)
	newTask.Config = new(gmp.CreateTaskConfig)
	newTask.Target = new(gmp.CreateTaskTarget)
	newTask.Scanner = new(gmp.CreateTaskScanner)
	newTask.Config.ID = configResp.Config[0].ID
	newTask.Target.ID = createTargetResp.ID
	newTask.Scanner.ID = getScannersResp.Scanner[0].ID
	newTaskResp, err := gmpClient.CreateTask(newTask)
	if err != nil {
		panic(err)
	}

	fmt.Printf("-> ok (task id: %s)\n", newTaskResp.ID)

	// Start the task
	fmt.Println("Start the task")
	st := &gmp.StartTaskCommand{}
	st.TaskID = newTaskResp.ID
	_, err = gmpClient.StartTask(st)
	if err != nil {
		panic(err)
	}

	// Monitoring task progress
	fmt.Println("Monitoring task progress...")
	for {
		gt := &gmp.GetTasksCommand{}
		gt.TaskID = newTaskResp.ID
		getTasksResp, err := gmpClient.GetTasks(gt)
		if err != nil {
			panic(err)
		}
		time.Sleep(10 * time.Second)
		fmt.Printf("Monitoring task progress: %s%%\n", getTasksResp.Task[0].Progress.Value)

		if x, _ := strconv.Atoi(getTasksResp.Task[0].Progress.Value); x >= 100 || getTasksResp.Task[0].Status == "Done" {
			break
		}
	}

	// Get results
	getResults := &gmp.GetResultsCommand{}
	getResults.TaskID = newTaskResp.ID
	getResults.Filter = `min_qod=0 rows=1000`
	results, err := gmpClient.GetResults(getResults)
	if err != nil {
		panic(err)
	}

	// Show results
	for i := 0; i < len(results.Result); i++ {
		fmt.Printf("Result[%d]: %s (score: %s)\n", i, results.Result[i].Name, results.Result[i].Severity)
	}
	fmt.Println("->ok")
	td := &gmp.DeleteTaskCommand{}
	td.TaskID = newTaskResp.ID

	_, err = gmpClient.DeleteTask(td)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("->ok(DeleteTask :%s )", td.TaskID)
	dt := &gmp.DeleteTargetCommand{}
	dt.TargetID = newTask.Target.ID
	_, err = gmpClient.DeleteTarget(dt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("->ok(DeleteTarget :%s )", dt.TargetID)
	a := &gmp.GetAssetsCommand{}
	a.Type = "host"
	a.Filter = `min_qod=0 rows=1000`
	getAssetsResp, err := gmpClient.GetAssetList(a)
	fmt.Printf("->ok(GetAssetList :%d )", len(getAssetsResp.Asset))
	for i := 0; i < len(getAssetsResp.Asset); i++ {
		fmt.Println(getAssetsResp.Asset[i].Name)
		delAsset := &gmp.DeleteAssetCommand{}
		delAsset.AssetID = getAssetsResp.Asset[i].ID
		_, err = gmpClient.DeleteAsset(delAsset)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("->ok(DeleteAsset :%s )", delAsset.AssetID)
	}

}
