package main

import "os"

var rootCmd = commandsGroup{
	command{Name: "init", Func: initCmdG.Run},
	command{Name: "code", Func: codeCmdG.Run},
	command{Name: "cluster", Func: clusterCmdG.Run},
	command{Name: "config", Func: configCmdG.Run},
	command{Name: "simpl-services", Func: simplServicesCmdG.Run},
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

var configCmdG = commandsGroup{
	ConfigPrepareFilesCmd,
}

var simplServicesCmdG = commandsGroup{
	SimplServicesTier1AuthorityUpCmd,
}

func main() {
	os.Exit(int(rootCmd.Run(os.Args[1:]...)))
}
