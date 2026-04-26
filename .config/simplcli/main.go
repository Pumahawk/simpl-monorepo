package main

import "os"

var rootCmd = commandsGroup{
	command{Name: "init", Func: initCmdG.Run},
	command{Name: "code", Func: codeCmdG.Run},
	command{Name: "cluster", Func: clusterCmdG.Run},
}

var initCmdG = commandsGroup{
	InitializationBuildCodeCmd,
}

var codeCmdG = commandsGroup{
	CodeBuildNoTestCmd,
	CodeTestAutomationAutoconfigureClientSecretCmd,
}

var clusterCmdG = commandsGroup{
	ClusterCreateCmd,
	ClusterForwardNodeUpCmd,
	ClusterForwardNodeDownCmd,
	ClusterForwardNodeComposeCmd,
	ClusterRedpandaUpCmd,
	ClusterRedpandaDownCmd,
	ClusterRedpandaComposeCmd,
	ClusterAuthorityInstallOrUpgradeCmd,
	ClusterInitializationEjbcaCmd,
	ClusterDownloadEjbcaPemCmd,
}

func main() {
	os.Exit(int(rootCmd.Run(os.Args[1:]...)))
}
