package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Production struct {
	NonTerminal string
	Rules       [][]string
}

type Parser struct {
	productions map[string][][]string
	input       string
	pos         int
	symbol      string
}

func main() {
	productions := []Production{
		{NonTerminal: "S", Rules: [][]string{{"X", "Y"}}},
		{NonTerminal: "X", Rules: [][]string{{"x", "X"}, {"x"}}},
		{NonTerminal: "Y", Rules: [][]string{{"y", "Y"}, {"y"}}},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan string yang akan diuji: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parser := Parser{
		productions: make(map[string][][]string),
		input:       input,
		pos:         0,
		symbol:      "",
	}

	// Inisialisasi aturan produksi
	for _, production := range productions {
		parser.productions[production.NonTerminal] = production.Rules
	}

	if parser.parseSymbol("S") && parser.pos == len(parser.input) {
		fmt.Println("String diterima oleh grammar.")
	} else {
		fmt.Println("String tidak diterima oleh grammar.")
	}
}

func (p *Parser) parseSymbol(nonTerminal string) bool {
	if p.pos >= len(p.input) {
		return false
	}

	// Simpan posisi saat ini
	savePos := p.pos

	// Cek aturan produksi yang cocok dengan simbol saat ini
	for _, rule := range p.productions[nonTerminal] {
		match := true
		for _, symbol := range rule {
			if isNonTerminal(symbol) {
				if !p.parseSymbol(symbol) {
					match = false
					break
				}
			} else {
				if !p.matchTerminal(symbol) {
					match = false
					break
				}
			}
		}

		// Jika semua simbol dalam aturan produksi cocok
		if match {
			return true
		}

		// Mengembalikan posisi ke awal
		p.pos = savePos
	}

	return false
}

func (p *Parser) matchTerminal(terminal string) bool {
	if p.pos+len(terminal) <= len(p.input) && p.input[p.pos:p.pos+len(terminal)] == terminal {
		p.pos += len(terminal)
		return true
	}
	return false
}

func isNonTerminal(symbol string) bool {
	return symbol[0] >= 'A' && symbol[0] <= 'Z'
}
