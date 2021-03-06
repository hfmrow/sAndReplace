// gohAssets.go

/*
	Source file auto-generated on Sat, 24 Apr 2021 04:33:11 using Gotk3 Objects Handler v1.7.8
	©2018-21 hfmrow https://hfmrow.github.io

	This software use gotk3 that is licensed under the ISC License:
	https://github.com/gotk3/gotk3/blob/master/LICENSE

	Copyright ©2018-21 hfmrow - Search And Replace v1.10 github.com/hfmrow/search-and-replace

	This program comes with absolutely no warranty. See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php
*/

/************************************************************************/
/* NOTICE: This file is completely rewritten when updating the assets. */
/**********************************************************************/

package main

import (
	"embed"
	"log"
)

// This functionality does not require explicit encoding of the files, at each
// compilation, the files are inserted into the resulting binary. Thus, updating
// assets is only required when new files are added to be embedded in order to
// create and declare the variables to which the files are linked.
// assetsDeclarationsUseEmbedded: Use native Go 'embed' package to include files
// content at runtime.
func assetsDeclarationsUseEmbedded(embedded ...bool) {
	mainGlade = readEmbedFile("assets/glade/main.glade")
	clipboard = readEmbedFile("assets/images/clipboard.png")
	clipboardRepl = readEmbedFile("assets/images/clipboard-repl.png")
	crossIcon48 = readEmbedFile("assets/images/Cross-icon-48.png")
	folder48 = readEmbedFile("assets/images/folder-48.png")
	linearProgressHorzBlue = readEmbedFile("assets/images/linear-progress-horz-blue.gif")
	logout48 = readEmbedFile("assets/images/logout-48.png")
	mimetypeSourceIconGolang48 = readEmbedFile("assets/images/Mimetype-source-icon-golang-48.png")
	options = readEmbedFile("assets/images/Options.png")
	replace = readEmbedFile("assets/images/replace.png")
	searchAndReplaceTop27 = readEmbedFile("assets/images/search-and-replace-top-27.png")
	searchAndReplaceTop48 = readEmbedFile("assets/images/search-and-replace-top-48.png")
	searchFolder48 = readEmbedFile("assets/images/search-folder-48.png")
	searchIcon48 = readEmbedFile("assets/images/search-icon-48.png")
	tickIcon48 = readEmbedFile("assets/images/Tick-icon-48.png")
}

// readEmbedFile: read 'embed' file system and return []byte data.
func readEmbedFile(filename string) (out []byte) {
	var err error
	out, err = embeddedFiles.ReadFile(filename)
	if err != nil {
		log.Printf("Unable to read embedded file: %s, %v\n", filename, err)
	}
	return
}

//go:embed assets/glade/main.glade
//go:embed assets/images/clipboard.png
//go:embed assets/images/clipboard-repl.png
//go:embed assets/images/Cross-icon-48.png
//go:embed assets/images/folder-48.png
//go:embed assets/images/linear-progress-horz-blue.gif
//go:embed assets/images/logout-48.png
//go:embed assets/images/Mimetype-source-icon-golang-48.png
//go:embed assets/images/Options.png
//go:embed assets/images/replace.png
//go:embed assets/images/search-and-replace-top-27.png
//go:embed assets/images/search-and-replace-top-48.png
//go:embed assets/images/search-folder-48.png
//go:embed assets/images/search-icon-48.png
//go:embed assets/images/Tick-icon-48.png
var embeddedFiles embed.FS
