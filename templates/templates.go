package templates

import (
	"cv/cv"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

func GeneratePDF(cv cv.CV, templateN uint) error {
    html, err := renderTemplate(cv, templateN)
    if err != nil {
        return err
    }

    err = os.WriteFile("./out/input.html", []byte(html), 0644)
    if err != nil {
        return err
    }

    cmd := exec.Command("wkhtmltopdf", "--enable-local-file-access", "./out/input.html", "./out/out.pdf")

    err = cmd.Run()
    if err != nil {
        fmt.Println("Error:", err)
    }

    return nil
}

func renderTemplate(cv cv.CV, templateN uint) (string, error) {
    tmpl, err := template.ParseGlob(fmt.Sprintf("./templates/%d/index.html", rune(templateN)))
    if err != nil {
        return "", err
    }

    html := new(strings.Builder)
    err = tmpl.Execute(html, cv)
    if err != nil {
        return "", err;
    }

    return html.String(), nil
}
