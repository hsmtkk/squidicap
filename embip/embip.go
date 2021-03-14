package embip

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

type templateFormat struct {
	ICAPHostIP string
}

func EmbedIPFile(templateFile, icapHostIP, outputFile string) error {
	bs, err := os.ReadFile(templateFile)
	if err != nil {
		return fmt.Errorf("failed to read template; %s; %w", templateFile, err)
	}
	writer, err := os.OpenFile(outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file; %s; %w", outputFile, err)
	}
	defer writer.Close()
	if err := EmbedIP(string(bs), icapHostIP, writer); err != nil {
		return err
	}
	return nil
}

func EmbedIP(templateText, icapHostIP string, writer io.Writer) error {
	tmpl := template.New("squid.conf")
	tmpl, err := tmpl.Parse(templateText)
	if err != nil {
		return fmt.Errorf("failed to parse template; %w", err)
	}
	tf := templateFormat{
		ICAPHostIP: icapHostIP,
	}
	if err := tmpl.Execute(writer, &tf); err != nil {
		return fmt.Errorf("failed to execute template; %w", err)
	}
	return nil
}
