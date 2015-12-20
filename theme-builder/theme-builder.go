package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type Package struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Version     string `json:"version"`
	Author      `json:"author"`
	License     string `json:"license"`
	Theme       `json:"theme"`
	Keywords    []string `json:"keywords"`
}

type Author struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Theme struct {
	File         string `json:"file"`
	Dark         bool   `json:"dark"`
	AddModeClass bool   `json:"addModeClass"`
}

func main() {
	var cssDir string = "./css/"
	files, err := ioutil.ReadDir(cssDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		buildTheme(cssDir, f)
	}
}

func buildTheme(cssDir string, f os.FileInfo) {
	var version string = "1.0.1"
	var authorName string = "Enrico D'Angelo"
	var authorUrl string = "https://github.com/enricodangelo"
	var license string = "MIT"
	var themeFile string = "theme.less"
	var addModeClass bool = true
	var themeName string
	var base00, base01, base02, base03, base04, base05, base06, base07, base08, base09, base0a, base0b, base0c, base0d, base0e, base0f string

	fileName := f.Name()
	themeName = strings.TrimSuffix(strings.TrimPrefix(fileName, "base16-"), ".css")
	fmt.Println(themeName)

	file, errFile := os.Open(cssDir + fileName)
	if errFile != nil {
		log.Fatal(errFile)
	}
	defer file.Close()

	var fileReader *bufio.Reader = bufio.NewReader(file)

	for {
		line, errRead := fileReader.ReadString('\n')
		if errRead == io.EOF {
			break
		}
		if errRead != nil {
			log.Fatal(errRead)
		}

		var property string = ""
		if len(line) == 50 {
			property = line[1:7]
		}
		switch property {
		case "base00":
			base00 = line[39:46]
		case "base01":
			base01 = line[39:46]
		case "base02":
			base02 = line[39:46]
		case "base03":
			base03 = line[39:46]
		case "base04":
			base04 = line[39:46]
		case "base05":
			base05 = line[39:46]
		case "base06":
			base06 = line[39:46]
		case "base07":
			base07 = line[39:46]
		case "base08":
			base08 = line[39:46]
		case "base09":
			base09 = line[39:46]
		case "base0A":
			base0a = line[39:46]
		case "base0B":
			base0b = line[39:46]
		case "base0C":
			base0c = line[39:46]
		case "base0D":
			base0d = line[39:46]
		case "base0E":
			base0e = line[39:46]
		case "base0F":
			base0f = line[39:46]
		}
	}

	var dark bool = true
	var dirname string = "../enricodangelo.base16-" + themeName + "-" + getVariation(dark) + "-theme"
	errDir := os.Mkdir(dirname, 0777)
	if errDir != nil {
		log.Fatal(errDir)
	}
	writeToFile(dirname, "LICENSE", makeLicenseFile(authorName))
	writeToFile(dirname, "package.json", makePackageJsonFile(themeName, dark, authorUrl, version, authorName, license, themeFile, addModeClass))
	writeToFile(dirname, "README.md", makeReadmeFile(themeName, dark))
	writeToFile(dirname, "theme.less", makeThemeFile(themeName, dark, base00, base01, base02, base03, base04, base05, base06, base07, base08, base09, base0a, base0b, base0c, base0d, base0e, base0f))
	zipit(dirname, dirname+".zip")

	dark = false
	dirname = "../enricodangelo.base16-" + themeName + "-" + getVariation(dark) + "-theme"
	errDir = os.Mkdir(dirname, 0777)
	if errDir != nil {
		log.Fatal(errDir)
	}
	writeToFile(dirname, "LICENSE", makeLicenseFile(authorName))
	writeToFile(dirname, "package.json", makePackageJsonFile(themeName, dark, authorUrl, version, authorName, license, themeFile, addModeClass))
	writeToFile(dirname, "README.md", makeReadmeFile(themeName, dark))
	writeToFile(dirname, "theme.less", makeThemeFile(themeName, dark, base00, base01, base02, base03, base04, base05, base06, base07, base08, base09, base0a, base0b, base0c, base0d, base0e, base0f))
	zipit(dirname, dirname+".zip")
}

func writeToFile(dir string, filename string, content string) {
	f, errCreate := os.Create(dir + "/" + filename)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, errWrite := w.WriteString(content)
	if errWrite != nil {
		log.Fatal(errWrite)
	}
	w.Flush()
}

func makeThemeFile(themeName string, dark bool, base00 string, base01 string, base02 string, base03 string, base04 string, base05 string, base06 string, base07 string, base08 string, base09 string, base0a string, base0b string, base0c string, base0d string, base0e string, base0f string) string {
	var theme string

	theme += "/* Define variables: Directly from CSS of http://chriskempson.github.io/base16/#" + themeName + " */\n"
	theme += "\n"
	if dark {
		theme += "@base00: " + base00 + ";\n"
		theme += "@base01: " + base01 + ";\n"
		theme += "@base02: " + base02 + ";\n"
		theme += "@base03: " + base03 + ";\n"
		theme += "@base04: " + base04 + ";\n"
		theme += "@base05: " + base05 + ";\n"
		theme += "@base06: " + base06 + ";\n"
		theme += "@base07: " + base07 + ";\n"
		theme += "@base08: " + base08 + ";\n"
		theme += "@base09: " + base09 + ";\n"
		theme += "@base0A: " + base0a + ";\n"
		theme += "@base0B: " + base0b + ";\n"
		theme += "@base0C: " + base0c + ";\n"
		theme += "@base0D: " + base0d + ";\n"
		theme += "@base0E: " + base0e + ";\n"
		theme += "@base0F: " + base0f + ";\n"
	} else {
		theme += "@base00: " + base07 + ";\n"
		theme += "@base01: " + base01 + ";\n"
		theme += "@base02: " + base06 + ";\n"
		theme += "@base03: " + base03 + ";\n"
		theme += "@base04: " + base04 + ";\n"
		theme += "@base05: " + base05 + ";\n"
		theme += "@base06: " + base02 + ";\n"
		theme += "@base07: " + base00 + ";\n"
		theme += "@base08: " + base08 + ";\n"
		theme += "@base09: " + base09 + ";\n"
		theme += "@base0A: " + base0a + ";\n"
		theme += "@base0B: " + base0b + ";\n"
		theme += "@base0C: " + base0c + ";\n"
		theme += "@base0D: " + base0d + ";\n"
		theme += "@base0E: " + base0e + ";\n"
		theme += "@base0F: " + base0f + ";\n"
	}
	theme += "\n"
	theme += "/* Code Styling */\n"
	theme += "\n"
	theme += ".CodeMirror {\n"
	theme += "// Gotta have that sweet sweet line spacing\n"
	theme += "    line-height: 1.2;\n"
	theme += "}\n"
	theme += ".CodeMirror-scroll {\n"
	theme += "    // Background color of main window\n"
	theme += "    background-color: @base00;\n"
	theme += "    // Color of otherwise un-styled text\n"
	theme += "    color: @base06;\n"
	theme += "}\n"
	theme += ".CodeMirror-gutters {\n"
	theme += "    // Gutter to the left\n"
	theme += "    background-color: @base00;\n"
	theme += "}\n"
	theme += ".CodeMirror-linenumber {\n"
	theme += "    // Line numbers in the gutter\n"
	theme += "    color: @base03;\n"
	theme += "}\n"
	theme += ".CodeMirror-selected {\n"
	theme += "    // Highlighed text\n"
	theme += "    background-color: @base02;\n"
	theme += "}\n"
	theme += ".CodeMirror-cursor {\n"
	theme += "    // The cursor\n"
	theme += "    border-left: 1px solid @base06;\n"
	theme += "}\n"
	theme += ".CodeMirror-matchingbracket,\n"
	theme += ".CodeMirror-matchingtag {\n"
	theme += "    // When you click a tag/bracket and the matching one is highlighted\n"
	theme += "    background: @base02 !important;\n"
	theme += "    color: @base06 !important;\n"
	theme += "    border-bottom: 1px solid @base08 !important;\n"
	theme += "}\n"
	theme += ".CodeMirror-foldgutter-open:after {\n"
	theme += "    // Those little arrows in the gutter\n"
	theme += "    color: @base03;\n"
	theme += "}\n"
	theme += ".CodeMirror-foldgutter-folded:after {\n"
	theme += "    // The arrows after you collapse code\n"
	theme += "    color: @base04;\n"
	theme += "}\n"
	theme += ".CodeMirror.over-gutter,\n"
	theme += ".CodeMirror-activeline {\n"
	theme += "    .CodeMirror-foldgutter-open:after {\n"
	theme += "        // Arrows when hovering on the gutter\n"
	theme += "        color: @base06;\n"
	theme += "    }\n"
	theme += "}\n"
	theme += ".CodeMirror-foldmarker {\n"
	theme += "    // The [...] marking collapsed code\n"
	theme += "    border-color: @base06;\n"
	theme += "    background-color: @base06;\n"
	theme += "    color: @base02;\n"
	theme += "}\n"
	theme += ".CodeMirror-searching {\n"
	theme += "    // Ctrl + F results\n"
	theme += "    background-color: @base0A;\n"
	theme += "    // Selected result\n"
	theme += "    &.searching-current-match {\n"
	theme += "        background-color: @base09;\n"
	theme += "    }\n"
	theme += "}\n"
	theme += "\n"
	theme += "/* Non-editor Styling */\n"
	theme += "\n"
	theme += "#image-holder,\n"
	theme += "#not-editor {\n"
	theme += "    background-color: @base00;\n"
	theme += "}\n"
	theme += "#image-holder {\n"
	theme += "    color: @base04;\n"
	theme += "}\n"
	theme += ".view-pane .image-view {\n"
	theme += "    color: @base06;\n"
	theme += "}\n"
	theme += "\n"
	theme += "/* All 'dem swatches */\n"
	theme += "\n"
	theme += ".cm-atom,\n"
	theme += ".cm-string,\n"
	theme += ".cm-string-2,\n"
	theme += ".cm-hr,\n"
	theme += "{\n"
	theme += "    color: @base0B;\n"
	theme += "}\n"
	theme += ".cm-number,\n"
	theme += ".cm-attribute,\n"
	theme += ".cm-plus {\n"
	theme += "    color: @base0A;\n"
	theme += "}\n"
	theme += ".cm-def,\n"
	theme += ".cm-property {\n"
	theme += "    color: @base0C;\n"
	theme += "}\n"
	theme += ".cm-variable,\n"
	theme += ".cm-variable-2,\n"
	theme += ".cm-variable-3 {\n"
	theme += "    color: @base0E;\n"
	theme += "}\n"
	theme += ".cm-operator,\n"
	theme += ".cm-meta,\n"
	theme += ".cm-bracket {\n"
	theme += "    color: @base06;\n"
	theme += "}\n"
	theme += ".cm-comment {\n"
	theme += "    color: @base03;\n"
	theme += "}\n"
	theme += ".cm-error,\n"
	theme += ".cm-minus {\n"
	theme += "    color: @base09;\n"
	theme += "}\n"
	theme += ".cm-header {\n"
	theme += "    color: @base0E;\n"
	theme += "}\n"
	theme += ".cm-link {\n"
	theme += "    color: @base0E;\n"
	theme += "    text-decoration: none;\n"
	theme += "}\n"
	theme += ".cm-rangeinfo {\n"
	theme += "    color: @base0C;\n"
	theme += "}\n"
	theme += ".cm-keyword,\n"
	theme += ".cm-qualifier,\n"
	theme += ".cm-builtin,\n"
	theme += ".cm-tag,\n"
	theme += ".cm-quote {\n"
	theme += "    color: @base08;\n"
	theme += "}\n"
	theme += "\n"
	theme += "/* Active Line Highlight support */\n"
	theme += ".CodeMirror-activeline-background {\n"
	theme += "    background-color: @base02;\n"
	theme += "}\n"
	theme += ".CodeMirror-activeline .CodeMirror-linenumber {\n"
	theme += "    background-color: @base01;\n"
	theme += "}\n"
	theme += ".CodeMirror-focused .CodeMirror-activeline .CodeMirror-gutter-elt {\n"
	theme += "    background: @base01;\n"
	theme += "}\n"
	theme += "\n"
	theme += "/*End of Theme*/\n"
	theme += "\n"

	return theme
}

func makeReadmeFile(themeName string, dark bool) string {
	var readme string
	var variation string = getVariation(dark)

	readme += "Base 16 " + upperInitial(themeName) + " " + upperInitial(variation) + " Theme for Brackets\n"
	readme += "============================\n"
	readme += "\n"
	readme += "Attempting to be as close to [" + upperInitial(themeName) + " " + upperInitial(variation) + "](http://chriskempson.github.io/base16/#" + themeName + ") as possible.\n"
	readme += "\n"
	readme += "Brackets theme adapted from [John Molakvoæ](https://github.com/skjnldsv/default-dark).\n"
	readme += "Colorscheme copied from [Chris Kempson](http://chriskempson.com).\n"

	return readme
}

func makeLicenseFile(authorName string) string {
	var license string

	license += "The MIT License (MIT)\n"
	license += "\n"
	license += "Copyright (c) " + strconv.Itoa(time.Now().Year()) + " " + authorName + "\n"
	license += "\n"
	license += "Permission is hereby granted, free of charge, to any person obtaining a copy\n"
	license += "of this software and associated documentation files (the \"Software\"), to deal\n"
	license += "in the Software without restriction, including without limitation the rights\n"
	license += "to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n"
	license += "copies of the Software, and to permit persons to whom the Software is\n"
	license += "furnished to do so, subject to the following conditions:\n"
	license += "\n"
	license += "The above copyright notice and this permission notice shall be included in all\n"
	license += "copies or substantial portions of the Software.\n"
	license += "\n"
	license += "THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n"
	license += "IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n"
	license += "FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n"
	license += "AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n"
	license += "LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n"
	license += "OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\n"
	license += "SOFTWARE.\n"

	return license
}

func makePackageJsonFile(themeName string, dark bool, authorUrl string, version string, authorName string, license string, themeFile string, addModeClass bool) string {
	var variation string = getVariation(dark)
	var keywords []string = []string{"theme", "base16", themeName, variation}

	packageJson := Package{
		"enricodangelo.base16-" + themeName + "-" + variation + "-theme",
		"Base16 " + upperInitial(themeName) + " " + upperInitial(variation) + " Theme",
		"Completely un-edited theme from Base16 " + upperInitial(themeName) + " " + upperInitial(variation) + ".",
		//authorUrl + "/brackets-themes/base16-" + themeName + "-" + variation + "-theme",
		authorUrl + "/brackets-themes/tree/master/enricodangelo.base16-" + themeName + "-" + variation + "-theme",
		version,
		Author{
			authorName,
			authorUrl,
		},
		license,
		Theme{
			themeFile,
			dark,
			addModeClass,
		},
		keywords,
	}
	jsonString, _ := json.MarshalIndent(packageJson, "", "    ")

	return string(jsonString)
}

func getVariation(dark bool) string {
	var variation string
	if dark {
		variation = "dark"
	} else {
		variation = "light"
	}
	return variation
}

func upperInitial(str string) string {
	if str == "" {
		return str
	}
	r, n := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[n:]
}

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
