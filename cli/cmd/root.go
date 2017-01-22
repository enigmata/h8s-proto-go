package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
    Use:   "cli",
    Short: "Command-line interface to Gödel system",
    Long: `Provides a means to interact with the Gödel system from a terminal
session, or enables scripted automation of various aspects
and behaviours of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle.`,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    
    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

    //RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
    if cfgFile != "" { // enable ability to specify config file via flag
        viper.SetConfigFile(cfgFile)
    }
    
    viper.SetConfigName(".cli") // name of config file (without extension)
    viper.AddConfigPath("$HOME")  // adding home directory as first search path
    viper.AutomaticEnv()          // read in environment variables that match
    
    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
