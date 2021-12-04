package main 

import (

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"

	// "fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/canvas"

	"strings"

	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/widget"

	// "github.com/knetic/govaluate"

	// "strconv"

	"fmt"
	
	"io/ioutil"
	
	"log"
)

func showGalleryApp(w fyne.Window){
	a:= app.New()
	w:= a.NewWindow("Gallery")
	// w.Resize(fyne.NewSize(500,280))
		root_src :="/config/Desktop/pictures"	
		
			files, err := ioutil.ReadDir(root_src)
			if err != nil {
				log.Fatal(err)
			}
			tabs := container.NewAppTabs()
			
		
			for _, file := range files {
				fmt.Println(file.Name(), file.IsDir())

				if file.IsDir()== false{
					extension:= strings.Split(file.Name(),".")[1];
					if extension =="png"|| extension == "jpeg"|| extension == "jpg"{
						
							image := canvas.NewImageFromFile(root_src+"/"+file.Name());
							image.FillMode = canvas.ImageFillOriginal
							tabs.Append(container.NewTabItem(file.Name(),image));
						
					}

				}
			}
		
		tabs.SetTabLocation(container.TabLocationLeading);
		w.SetContent(tabs);
		w.Resize(fyne.NewSize(700,730));


		
	w.ShowAndRun()

}