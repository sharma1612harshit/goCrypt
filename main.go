package main

import (
	"./cli"
	"./crypt"
	"./logger"
	"./utils"
	"os"
)

// default key for case if you want to use a fixed key
const DefaultKey = "my_default_key"

func main() {
	// check if vim is installed
	err := utils.CheckVim()
	if err != nil {
		logger.Critical(err.Error())
	}

	// parsing the arguments to figure out the actions
	arguments := os.Args
	options, err := utils.ArgParse(arguments)
	if err != nil {
		logger.Help()
		logger.Critical(err.Error())
	}

	if options[0] == "-h" { // display help options
		logger.Help()
	} else if options[0] == "-w" { // write file functions
		filename := options[1]

		cliInput, err := cli.CaptureInputFromEditor()
		if err != nil {
			logger.Critical(err.Error())
		}

		Key := DefaultKey
		if options[2] != "" {
			Key = options[2]
		}

		Encrypted, err := crypt.Encrypt(cliInput, Key)
		if err != nil {
			logger.Critical(err.Error())
		}

		err = utils.WriteToFile(filename, Encrypted)
		if err != nil {
			logger.Warning("saving failed ...")

			Decrypted, decErr := crypt.Decrypt(Encrypted, Key)
			if decErr != nil {
				logger.Critical(err.Error())
			}

			logger.Info(string(Decrypted))

			logger.Critical(err.Error())
		}

		logger.Info(filename + " saved!")
	} else if options[0] == "-r" { // read file functions
		filename := options[1]

		readData, err := utils.ReadFromFile(filename)
		if err != nil {
			logger.Critical(err.Error())
		}

		Key := DefaultKey
		if options[2] != "" {
			Key = options[2]
		}

		Decrypted, err := crypt.Decrypt(readData, Key)
		if err != nil {
			logger.Critical(err.Error())
		}

		logger.Info(string(Decrypted))
	}
}
