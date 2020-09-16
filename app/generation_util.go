package app

import (
	"fmt"
	"log"
	"os"
	"sort"
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

func createTemplate(templateBody string) error {
	fileName := fmt.Sprintf("%s/%s.html", generateDirectory, fileNamePrefix)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	log.Printf("Created template: %s\n", fileName)
	writeFile(f, templateBody)
	log.Printf("Wrote template content to %s\n", fileName)
	return f.Close()
}

func createComponent(
	imports, directives, declarations []string, templateBody string) error {
	createGenerateDirectory()
	sortAndJoin := func(a []string, sep, suffix string) string {
		sort.Strings(a)
		return strings.Join(a, sep) + suffix
	}
	importsStr := sortAndJoin(imports, ";\n", ";")
	directivesStr := sortAndJoin(directives, ",\n    ", ",")
	declarationsStr := sortAndJoin(declarations, ";\n  ", ";")
	templateUrl := fmt.Sprintf("%s.html", fileNamePrefix)
	replacer := strings.NewReplacer(
		"{selector}", selector,
		"{templateUrl}", templateUrl,
		"{componentName}", componentName,
		"{imports}", importsStr,
		"{directives}", directivesStr,
		"{declarations}", declarationsStr,
	)
	componentBody := replacer.Replace(componentBodyFormat)
	createGenerateDirectory()
	if err := createTemplate(templateBody); err != nil {
		return err
	}
	fileName := fmt.Sprintf("%s/%s.dart", generateDirectory, fileNamePrefix)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	log.Printf("Created component: %s\n", fileName)
	writeFile(f, copyrightHeader)
	log.Printf("Wrote copyright header to %s\n", fileName)
	writeFile(f, componentBody)
	log.Printf("Wrote component body to %s\n", fileName)
	writeFile(f, copyrightFooter)
	log.Printf("Wrote copyright footer to %s\n", fileName)
	return f.Close()
}
