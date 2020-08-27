package app

import (
	"log"
	"os"
)

// Create directory: 'generated', if it doesn't exist.
func createGeneratedDirectory() {
	err := os.Mkdir("generated", 0755)
	if err == nil {
		log.Println("Created directory: 'generated'")
		return
	}
	if os.IsExist(err) {
		log.Println("Directory 'generated' already exists.",
			"Skipping creation...")
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

func createComponent() {
	f, err := os.OpenFile("generated/new_component.dart",
		os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created component: new_component.dart")

	writeFile(f, copyrightHeader)
	log.Println("Wrote copyright header to new_component.dart")

	writeFile(f, componentBody)
	log.Println("Wrote component body to new_component.dart")

	writeFile(f, copyrightFooter)
	log.Println("Wrote copyright footer to new_component.dart")

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func CreateComponentSkeleton() {
	createGeneratedDirectory()
	createComponent()
}
