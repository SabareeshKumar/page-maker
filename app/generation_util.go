package app

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type declarationList []string

func (dl declarationList) Len() int {
	return len(dl)
}

func (dl declarationList) Less(i, j int) bool {
	iFinal := strings.HasPrefix(dl[i], "final")
	jFinal := strings.HasPrefix(dl[j], "final")
	if iFinal != jFinal {
		return iFinal
	}
	return dl[i] < dl[j]
}

func (dl declarationList) Swap(i, j int) {
	dl[i], dl[j] = dl[j], dl[i]
}

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
	sortAndJoin := func(
		sorter func([]string), a []string, sep, suffix string) string {
		sorter(a)
		return strings.Join([]string(a), sep) + suffix
	}
	declarationSorter := func(decl []string) {
		sort.Sort(declarationList(decl))
	}
	importsStr := sortAndJoin(sort.Strings, imports, ";\n", ";")
	directivesStr := sortAndJoin(sort.Strings, directives, ",\n    ", ",")
	declarationsStr := sortAndJoin(
		declarationSorter, declarations, ";\n  ", ";")
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
