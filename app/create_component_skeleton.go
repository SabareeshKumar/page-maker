package app

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Create directory: 'generated', if it doesn't exist.
func createGenerateDirectory() {
	err := os.Mkdir(generateDirectory, 0755)
	if err == nil {
		log.Printf("Created directory: %s\n", generateDirectory)
		return
	}
	if os.IsExist(err) {
		log.Printf("Directory %s already exists."+
			" Skipping creation...\n", generateDirectory)
		return
	}
	log.Fatal(err)
}

func writeFile(f *os.File, data string) {
	if _, err := f.Write([]byte(data)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
}

func createTemplateFile() {
	fileName := fmt.Sprintf("%s/%s.html", generateDirectory, fileNamePrefix)
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created template: %s\n", fileName)

	writeFile(f, componentTemplateFormat)
	log.Printf("Wrote template content to %s\n", fileName)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func createComponent() {
	createTemplateFile()
	templateUrl := fmt.Sprintf("%s.html", fileNamePrefix)
	replacer := strings.NewReplacer(
		"{selector}", selector,
		"{templateUrl}", templateUrl,
		"{componentName}", componentName,
		"{expansionPanelName}", expansionPanelName)
	componentBody := replacer.Replace(componentBodyFormatWithPanel)
	
	fileName := fmt.Sprintf("%s/%s.dart", generateDirectory, fileNamePrefix)
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created component: %s\n", fileName)
	
	writeFile(f, copyrightHeader)
	log.Printf("Wrote copyright header to %s\n", fileName)
	
	writeFile(f, componentBody)
	log.Printf("Wrote component body to %s\n", fileName)
	
	writeFile(f, copyrightFooter)
	log.Printf("Wrote copyright footer to %s\n", fileName)
	
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func CreateComponentSkeleton() {
	createGenerateDirectory()
	createComponent()
}
