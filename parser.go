package main

import (
	"bufio"
	"io"
	"strings"
)

type Clipping struct {
	Title string
	Meta  string
	Body  string
}

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(reader io.Reader) ([]Clipping, error) {
	lineReader := bufio.NewReader(reader)
	clippings := make([]Clipping, 0)
	for {
		clipping, err := p.parseClipping(lineReader)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		clippings = append(clippings, clipping)
	}
	return clippings, nil
}

func (p *Parser) parseClipping(reader *bufio.Reader) (Clipping, error) {
	clipping := Clipping{}

	line, err := reader.ReadString('\n')
	if err != nil {
		return clipping, err
	}
	clipping.Title = strings.TrimSpace(line)

	line, err = reader.ReadString('\n')
	if err != nil {
		return clipping, err
	}
	clipping.Meta = strings.TrimSpace(line)

	line, err = reader.ReadString('\n')
	if err != nil {
		return clipping, err
	}

	var body string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return clipping, err
		}
		if strings.TrimSpace(line) == "==========" {
			break
		}
		body += line
	}
	clipping.Body = strings.TrimSpace(body)

	return clipping, nil
}
