package train

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func StartPYTraining(clientIPs []string, hostIP string) {
	var trainPhrase string

	trainPhrase += "python3 distribute.py --ps_hosts=" + hostIP + ":2222 --worker_hosts="

	trainPhrase += clientIPs[0] + ":2225"

	for i := 1; i < len(clientIPs); i++ {
		trainPhrase += "," + clientIPs[i] + ":2225"
	}

	trainPhrase += " --job_name=ps --task_index=0"

	cmd := exec.Command("ubuntu1804","-c", "/bin/bash;" +
		"source /root/work/go/venv/bin/activate;" +
		"cd /root/work/go;" + trainPhrase)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Start())
}

func EnterPYTraining(clientIPs []string, hostIP string, index int) {
	var trainPhrase string


	trainPhrase += "python3 distribute.py --ps_hosts=" + hostIP + ":2222 --worker_hosts="

	trainPhrase += clientIPs[0] + ":2225"

	for i := 1; i < len(clientIPs); i++ {
		trainPhrase += "," + clientIPs[i] + ":2225"
	}

	trainPhrase += " --job_name=worker --task_index=" + strconv.Itoa(index)

	//log.Println(trainPhrase)

	cmd := exec.Command("ubuntu1804","-c", "/bin/bash;" +
		"source /root/work/go/venv/bin/activate;" +
		"cd /root/work/go;" + trainPhrase)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Start())
}
