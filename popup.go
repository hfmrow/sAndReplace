// popup.go

/*
	Source file auto-generated on Fri, 08 Nov 2019 03:45:35 using Gotk3ObjHandler v1.4 ©2018-19 H.F.M
	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2019 H.F.M - Functions & Library Manager github.com/...
	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

package main

import (
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	glts "github.com/hfmrow/genLib/tools"

	gimc "github.com/hfmrow/gtk3Import/misc"
)

/*
* Init Popup
 */

// initPopupTreeView:
func initTreeViewPopupMenu() {
	popupMenu = gimc.PopupMenuNew()
	popupMenu.WithIcons = true
	popupMenu.AddItem("Open _directory", func() {
		getTreeViewFilename()
		openDir(currFilename)
	}, folder48)
	popupMenu.AddItem("Open _file", func() {
		getTreeViewFilename()
		openFile(currFilename)
	}, mimetypeSourceIconGolang48)
	popupMenu.MenuBuild()
}

// Retrieve filename from selected row to "currFilename"
func getTreeViewFilename() {
	var err error
	var value *glib.Value
	var iters []*gtk.TreeIter

	if iters = tvsList.GetSelectedIters(); len(iters) > 0 {

		if value, err = tvsList.ListStore.GetValue(iters[0], 3); err == nil { // Field 3: get full path
			currFilename, err = value.GetString() // Get selected file path
		}
	}
}

// openFile:
func openFile(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		open(filename)
	}
}

// openDir:
func openDir(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		open(filepath.Dir(filename))
	}
}

// open: show file or dir depending on "path".
func open(path string) {
	glib.IdleAdd(func() { // IdleAdd to permit gtk3 working right during goroutine
		// Using goroutine to permit the usage of another thread and freeing the current one.
		go glts.ExecCommand(mainOptions.AppLauncher, path)
	})
}

// popupTextViewPopulateMenu: Append some items to contextmenu of the popup textview
func popupTextViewPopulateMenu(txtView *gtk.TextView, popup *gtk.Widget) {
	// Convert gtk.Widget to gtk.Menu object
	pop := &gtk.Menu{gtk.MenuShell{gtk.Container{*popup}}}
	// create new menuitems
	popMenuTextView = gimc.PopupMenuNew()
	popMenuTextView.WithIcons = true
	popMenuTextView.AddSeparator()
	popMenuTextView.AddItem("Open _directory", func() { openDir(currFilename) }, folder48)
	popMenuTextView.AddItem("Open _file", func() { openFile(currFilename) }, mimetypeSourceIconGolang48)
	// Append them to the existing menu
	popMenuTextView.AppendToExistingMenu(pop)
}
