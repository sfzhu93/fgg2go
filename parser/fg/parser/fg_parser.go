// Generated from parser/FG.g4 by ANTLR 4.7.

package parser // FG

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 283,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 49, 10, 2, 3, 2, 5, 2, 52, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 5, 2, 71, 10, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 78, 10, 3, 3,
	3, 3, 3, 6, 3, 82, 10, 3, 13, 3, 14, 3, 83, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 103, 10, 5, 12, 5, 14, 5, 106, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 116, 10, 5, 12, 5, 14, 5, 119, 11, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 124, 10, 5, 3, 5, 5, 5, 127, 10, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 141, 10,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 147, 10, 7, 3, 7, 5, 7, 150, 10, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 157, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 163, 10, 9, 3, 10, 3, 10, 3, 10, 7, 10, 168, 10, 10, 12, 10, 14, 10,
	171, 11, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 179, 10,
	12, 12, 12, 14, 12, 182, 11, 12, 3, 13, 3, 13, 5, 13, 186, 10, 13, 3, 14,
	3, 14, 3, 14, 5, 14, 191, 10, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 7, 15, 199, 10, 15, 12, 15, 14, 15, 202, 11, 15, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 211, 10, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 231, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 241, 10, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 5, 20, 248, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 258, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 270, 10, 20, 12, 20, 14, 20,
	273, 11, 20, 3, 21, 3, 21, 3, 21, 7, 21, 278, 10, 21, 12, 21, 14, 21, 281,
	11, 21, 3, 21, 2, 3, 38, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 2, 3, 4, 2, 10, 10, 38, 38, 2, 298, 2,
	42, 3, 2, 2, 2, 4, 81, 3, 2, 2, 2, 6, 85, 3, 2, 2, 2, 8, 126, 3, 2, 2,
	2, 10, 128, 3, 2, 2, 2, 12, 149, 3, 2, 2, 2, 14, 156, 3, 2, 2, 2, 16, 162,
	3, 2, 2, 2, 18, 164, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 175, 3, 2, 2,
	2, 24, 185, 3, 2, 2, 2, 26, 187, 3, 2, 2, 2, 28, 195, 3, 2, 2, 2, 30, 203,
	3, 2, 2, 2, 32, 206, 3, 2, 2, 2, 34, 230, 3, 2, 2, 2, 36, 240, 3, 2, 2,
	2, 38, 257, 3, 2, 2, 2, 40, 274, 3, 2, 2, 2, 42, 43, 7, 20, 2, 2, 43, 44,
	7, 19, 2, 2, 44, 48, 7, 3, 2, 2, 45, 46, 7, 24, 2, 2, 46, 47, 7, 38, 2,
	2, 47, 49, 7, 3, 2, 2, 48, 45, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51,
	3, 2, 2, 2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 53, 3, 2, 2, 2, 53, 54, 7, 17, 2, 2, 54, 55, 7, 19, 2, 2, 55, 56, 7,
	4, 2, 2, 56, 57, 7, 5, 2, 2, 57, 70, 7, 6, 2, 2, 58, 59, 7, 7, 2, 2, 59,
	60, 7, 8, 2, 2, 60, 71, 5, 38, 20, 2, 61, 62, 7, 25, 2, 2, 62, 63, 7, 9,
	2, 2, 63, 64, 7, 26, 2, 2, 64, 65, 7, 4, 2, 2, 65, 66, 7, 10, 2, 2, 66,
	67, 7, 11, 2, 2, 67, 68, 5, 38, 20, 2, 68, 69, 7, 5, 2, 2, 69, 71, 3, 2,
	2, 2, 70, 58, 3, 2, 2, 2, 70, 61, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73,
	7, 12, 2, 2, 73, 74, 7, 2, 2, 3, 74, 3, 3, 2, 2, 2, 75, 78, 5, 6, 4, 2,
	76, 78, 5, 10, 6, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 79, 3,
	2, 2, 2, 79, 80, 7, 3, 2, 2, 80, 82, 3, 2, 2, 2, 81, 77, 3, 2, 2, 2, 82,
	83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 5, 3, 2, 2,
	2, 85, 86, 7, 23, 2, 2, 86, 87, 7, 34, 2, 2, 87, 88, 5, 12, 7, 2, 88, 7,
	3, 2, 2, 2, 89, 90, 7, 21, 2, 2, 90, 127, 5, 38, 20, 2, 91, 92, 5, 34,
	18, 2, 92, 93, 7, 3, 2, 2, 93, 94, 5, 8, 5, 2, 94, 127, 3, 2, 2, 2, 95,
	96, 7, 25, 2, 2, 96, 97, 7, 9, 2, 2, 97, 98, 7, 27, 2, 2, 98, 99, 7, 4,
	2, 2, 99, 104, 9, 2, 2, 2, 100, 103, 7, 11, 2, 2, 101, 103, 5, 38, 20,
	2, 102, 100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 107, 127, 7, 5, 2, 2, 108, 109, 7, 31, 2, 2, 109, 117, 7, 6,
	2, 2, 110, 111, 7, 33, 2, 2, 111, 112, 5, 36, 19, 2, 112, 113, 7, 13, 2,
	2, 113, 114, 5, 8, 5, 2, 114, 116, 3, 2, 2, 2, 115, 110, 3, 2, 2, 2, 116,
	119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 123,
	3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 121, 7, 32, 2, 2, 121, 122, 7, 13,
	2, 2, 122, 124, 5, 8, 5, 2, 123, 120, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 127, 7, 12, 2, 2, 126, 89, 3, 2, 2, 2, 126,
	91, 3, 2, 2, 2, 126, 95, 3, 2, 2, 2, 126, 108, 3, 2, 2, 2, 127, 9, 3, 2,
	2, 2, 128, 129, 7, 17, 2, 2, 129, 130, 7, 4, 2, 2, 130, 131, 5, 30, 16,
	2, 131, 132, 7, 5, 2, 2, 132, 133, 5, 26, 14, 2, 133, 134, 7, 6, 2, 2,
	134, 135, 5, 8, 5, 2, 135, 136, 7, 12, 2, 2, 136, 11, 3, 2, 2, 2, 137,
	138, 7, 22, 2, 2, 138, 140, 7, 6, 2, 2, 139, 141, 5, 18, 10, 2, 140, 139,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 150, 7, 12,
	2, 2, 143, 144, 7, 18, 2, 2, 144, 146, 7, 6, 2, 2, 145, 147, 5, 22, 12,
	2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148,
	150, 7, 12, 2, 2, 149, 137, 3, 2, 2, 2, 149, 143, 3, 2, 2, 2, 150, 13,
	3, 2, 2, 2, 151, 152, 7, 14, 2, 2, 152, 157, 7, 15, 2, 2, 153, 154, 7,
	15, 2, 2, 154, 157, 7, 14, 2, 2, 155, 157, 7, 14, 2, 2, 156, 151, 3, 2,
	2, 2, 156, 153, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 15, 3, 2, 2, 2,
	158, 163, 7, 34, 2, 2, 159, 160, 5, 14, 8, 2, 160, 161, 5, 16, 9, 2, 161,
	163, 3, 2, 2, 2, 162, 158, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 163, 17, 3,
	2, 2, 2, 164, 169, 5, 20, 11, 2, 165, 166, 7, 3, 2, 2, 166, 168, 5, 20,
	11, 2, 167, 165, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2,
	169, 170, 3, 2, 2, 2, 170, 19, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 173,
	7, 34, 2, 2, 173, 174, 5, 16, 9, 2, 174, 21, 3, 2, 2, 2, 175, 180, 5, 24,
	13, 2, 176, 177, 7, 3, 2, 2, 177, 179, 5, 24, 13, 2, 178, 176, 3, 2, 2,
	2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	23, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 186, 5, 26, 14, 2, 184, 186,
	7, 34, 2, 2, 185, 183, 3, 2, 2, 2, 185, 184, 3, 2, 2, 2, 186, 25, 3, 2,
	2, 2, 187, 188, 7, 34, 2, 2, 188, 190, 7, 4, 2, 2, 189, 191, 5, 28, 15,
	2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192,
	193, 7, 5, 2, 2, 193, 194, 7, 34, 2, 2, 194, 27, 3, 2, 2, 2, 195, 200,
	5, 30, 16, 2, 196, 197, 7, 11, 2, 2, 197, 199, 5, 30, 16, 2, 198, 196,
	3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 29, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 7, 34, 2, 2,
	204, 205, 5, 16, 9, 2, 205, 31, 3, 2, 2, 2, 206, 207, 7, 9, 2, 2, 207,
	208, 7, 34, 2, 2, 208, 210, 7, 4, 2, 2, 209, 211, 5, 40, 21, 2, 210, 209,
	3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 7, 5,
	2, 2, 213, 33, 3, 2, 2, 2, 214, 215, 7, 28, 2, 2, 215, 216, 7, 4, 2, 2,
	216, 217, 5, 38, 20, 2, 217, 218, 7, 5, 2, 2, 218, 231, 3, 2, 2, 2, 219,
	220, 5, 38, 20, 2, 220, 221, 7, 15, 2, 2, 221, 222, 5, 38, 20, 2, 222,
	231, 3, 2, 2, 2, 223, 224, 7, 29, 2, 2, 224, 225, 5, 38, 20, 2, 225, 226,
	5, 32, 17, 2, 226, 231, 3, 2, 2, 2, 227, 228, 7, 34, 2, 2, 228, 229, 7,
	16, 2, 2, 229, 231, 5, 38, 20, 2, 230, 214, 3, 2, 2, 2, 230, 219, 3, 2,
	2, 2, 230, 223, 3, 2, 2, 2, 230, 227, 3, 2, 2, 2, 231, 35, 3, 2, 2, 2,
	232, 233, 5, 38, 20, 2, 233, 234, 7, 15, 2, 2, 234, 235, 5, 38, 20, 2,
	235, 241, 3, 2, 2, 2, 236, 237, 7, 34, 2, 2, 237, 238, 7, 16, 2, 2, 238,
	239, 7, 15, 2, 2, 239, 241, 5, 38, 20, 2, 240, 232, 3, 2, 2, 2, 240, 236,
	3, 2, 2, 2, 241, 37, 3, 2, 2, 2, 242, 243, 8, 20, 1, 2, 243, 258, 7, 34,
	2, 2, 244, 245, 7, 34, 2, 2, 245, 247, 7, 6, 2, 2, 246, 248, 5, 40, 21,
	2, 247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	258, 7, 12, 2, 2, 250, 251, 7, 30, 2, 2, 251, 252, 7, 4, 2, 2, 252, 253,
	5, 16, 9, 2, 253, 254, 7, 5, 2, 2, 254, 258, 3, 2, 2, 2, 255, 256, 7, 15,
	2, 2, 256, 258, 5, 38, 20, 3, 257, 242, 3, 2, 2, 2, 257, 244, 3, 2, 2,
	2, 257, 250, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 271, 3, 2, 2, 2, 259,
	260, 12, 7, 2, 2, 260, 261, 7, 9, 2, 2, 261, 270, 7, 34, 2, 2, 262, 263,
	12, 6, 2, 2, 263, 270, 5, 32, 17, 2, 264, 265, 12, 5, 2, 2, 265, 266, 7,
	9, 2, 2, 266, 267, 7, 4, 2, 2, 267, 268, 7, 34, 2, 2, 268, 270, 7, 5, 2,
	2, 269, 259, 3, 2, 2, 2, 269, 262, 3, 2, 2, 2, 269, 264, 3, 2, 2, 2, 270,
	273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 39, 3,
	2, 2, 2, 273, 271, 3, 2, 2, 2, 274, 279, 5, 38, 20, 2, 275, 276, 7, 11,
	2, 2, 276, 278, 5, 38, 20, 2, 277, 275, 3, 2, 2, 2, 278, 281, 3, 2, 2,
	2, 279, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 41, 3, 2, 2, 2, 281,
	279, 3, 2, 2, 2, 30, 48, 51, 70, 77, 83, 102, 104, 117, 123, 126, 140,
	146, 149, 156, 162, 169, 180, 185, 190, 200, 210, 230, 240, 247, 257, 269,
	271, 279,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "')'", "'{'", "'_'", "'='", "'.'", "'\"%#v\"'", "','",
	"'}'", "':'", "'chan'", "'<-'", "':='", "'func'", "'interface'", "'main'",
	"'package'", "'return'", "'struct'", "'type'", "'import'", "'fmt'", "'Printf'",
	"'Sprintf'", "'close'", "'go'", "'make'", "'select'", "'default'", "'case'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "FUNC", "INTERFACE",
	"MAIN", "PACKAGE", "RETURN", "STRUCT", "TYPE", "IMPORT", "FMT", "PRINTF",
	"SPRINTF", "CLOSE", "GORTN", "MAKE", "SELECT", "DEFAULT", "CASE", "NAME",
	"WHITESPACE", "COMMENT", "LINE_COMMENT", "STRING",
}

var ruleNames = []string{
	"program", "decls", "typeDecl", "methBody", "methDecl", "typeLit", "channelTypeLit",
	"type_", "fieldDecls", "fieldDecl", "specs", "spec", "sig", "params", "paramDecl",
	"methCall_", "stmt", "caseGuard", "expr", "exprs",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FGParser struct {
	*antlr.BaseParser
}

func NewFGParser(input antlr.TokenStream) *FGParser {
	this := new(FGParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FG.g4"

	return this
}

// FGParser tokens.
const (
	FGParserEOF          = antlr.TokenEOF
	FGParserT__0         = 1
	FGParserT__1         = 2
	FGParserT__2         = 3
	FGParserT__3         = 4
	FGParserT__4         = 5
	FGParserT__5         = 6
	FGParserT__6         = 7
	FGParserT__7         = 8
	FGParserT__8         = 9
	FGParserT__9         = 10
	FGParserT__10        = 11
	FGParserT__11        = 12
	FGParserT__12        = 13
	FGParserT__13        = 14
	FGParserFUNC         = 15
	FGParserINTERFACE    = 16
	FGParserMAIN         = 17
	FGParserPACKAGE      = 18
	FGParserRETURN       = 19
	FGParserSTRUCT       = 20
	FGParserTYPE         = 21
	FGParserIMPORT       = 22
	FGParserFMT          = 23
	FGParserPRINTF       = 24
	FGParserSPRINTF      = 25
	FGParserCLOSE        = 26
	FGParserGORTN        = 27
	FGParserMAKE         = 28
	FGParserSELECT       = 29
	FGParserDEFAULT      = 30
	FGParserCASE         = 31
	FGParserNAME         = 32
	FGParserWHITESPACE   = 33
	FGParserCOMMENT      = 34
	FGParserLINE_COMMENT = 35
	FGParserSTRING       = 36
)

// FGParser rules.
const (
	FGParserRULE_program        = 0
	FGParserRULE_decls          = 1
	FGParserRULE_typeDecl       = 2
	FGParserRULE_methBody       = 3
	FGParserRULE_methDecl       = 4
	FGParserRULE_typeLit        = 5
	FGParserRULE_channelTypeLit = 6
	FGParserRULE_type_          = 7
	FGParserRULE_fieldDecls     = 8
	FGParserRULE_fieldDecl      = 9
	FGParserRULE_specs          = 10
	FGParserRULE_spec           = 11
	FGParserRULE_sig            = 12
	FGParserRULE_params         = 13
	FGParserRULE_paramDecl      = 14
	FGParserRULE_methCall_      = 15
	FGParserRULE_stmt           = 16
	FGParserRULE_caseGuard      = 17
	FGParserRULE_expr           = 18
	FGParserRULE_exprs          = 19
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(FGParserPACKAGE, 0)
}

func (s *ProgramContext) AllMAIN() []antlr.TerminalNode {
	return s.GetTokens(FGParserMAIN)
}

func (s *ProgramContext) MAIN(i int) antlr.TerminalNode {
	return s.GetToken(FGParserMAIN, i)
}

func (s *ProgramContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FGParserFUNC, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FGParserEOF, 0)
}

func (s *ProgramContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramContext) FMT() antlr.TerminalNode {
	return s.GetToken(FGParserFMT, 0)
}

func (s *ProgramContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(FGParserPRINTF, 0)
}

func (s *ProgramContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(FGParserIMPORT, 0)
}

func (s *ProgramContext) STRING() antlr.TerminalNode {
	return s.GetToken(FGParserSTRING, 0)
}

func (s *ProgramContext) Decls() IDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *FGParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FGParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(FGParserPACKAGE)
	}
	{
		p.SetState(41)
		p.Match(FGParserMAIN)
	}
	{
		p.SetState(42)
		p.Match(FGParserT__0)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGParserIMPORT {
		{
			p.SetState(43)
			p.Match(FGParserIMPORT)
		}
		{
			p.SetState(44)
			p.Match(FGParserSTRING)
		}
		{
			p.SetState(45)
			p.Match(FGParserT__0)
		}

	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(48)
			p.Decls()
		}

	}
	{
		p.SetState(51)
		p.Match(FGParserFUNC)
	}
	{
		p.SetState(52)
		p.Match(FGParserMAIN)
	}
	{
		p.SetState(53)
		p.Match(FGParserT__1)
	}
	{
		p.SetState(54)
		p.Match(FGParserT__2)
	}
	{
		p.SetState(55)
		p.Match(FGParserT__3)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGParserT__4:
		{
			p.SetState(56)
			p.Match(FGParserT__4)
		}
		{
			p.SetState(57)
			p.Match(FGParserT__5)
		}
		{
			p.SetState(58)
			p.expr(0)
		}

	case FGParserFMT:
		{
			p.SetState(59)
			p.Match(FGParserFMT)
		}
		{
			p.SetState(60)
			p.Match(FGParserT__6)
		}
		{
			p.SetState(61)
			p.Match(FGParserPRINTF)
		}
		{
			p.SetState(62)
			p.Match(FGParserT__1)
		}
		{
			p.SetState(63)
			p.Match(FGParserT__7)
		}
		{
			p.SetState(64)
			p.Match(FGParserT__8)
		}
		{
			p.SetState(65)
			p.expr(0)
		}
		{
			p.SetState(66)
			p.Match(FGParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(70)
		p.Match(FGParserT__9)
	}
	{
		p.SetState(71)
		p.Match(FGParserEOF)
	}

	return localctx
}

// IDeclsContext is an interface to support dynamic dispatch.
type IDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclsContext differentiates from other interfaces.
	IsDeclsContext()
}

type DeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclsContext() *DeclsContext {
	var p = new(DeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_decls
	return p
}

func (*DeclsContext) IsDeclsContext() {}

func NewDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclsContext {
	var p = new(DeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_decls

	return p
}

func (s *DeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclsContext) AllTypeDecl() []ITypeDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem())
	var tst = make([]ITypeDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclContext)
		}
	}

	return tst
}

func (s *DeclsContext) TypeDecl(i int) ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclsContext) AllMethDecl() []IMethDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethDeclContext)(nil)).Elem())
	var tst = make([]IMethDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethDeclContext)
		}
	}

	return tst
}

func (s *DeclsContext) MethDecl(i int) IMethDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethDeclContext)
}

func (s *DeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterDecls(s)
	}
}

func (s *DeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitDecls(s)
	}
}

func (p *FGParser) Decls() (localctx IDeclsContext) {
	localctx = NewDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FGParserRULE_decls)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(75)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case FGParserTYPE:
				{
					p.SetState(73)
					p.TypeDecl()
				}

			case FGParserFUNC:
				{
					p.SetState(74)
					p.MethDecl()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(77)
				p.Match(FGParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(FGParserTYPE, 0)
}

func (s *TypeDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *TypeDeclContext) TypeLit() ITypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *FGParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FGParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(FGParserTYPE)
	}
	{
		p.SetState(84)
		p.Match(FGParserNAME)
	}
	{
		p.SetState(85)
		p.TypeLit()
	}

	return localctx
}

// IMethBodyContext is an interface to support dynamic dispatch.
type IMethBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethBodyContext differentiates from other interfaces.
	IsMethBodyContext()
}

type MethBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethBodyContext() *MethBodyContext {
	var p = new(MethBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_methBody
	return p
}

func (*MethBodyContext) IsMethBodyContext() {}

func NewMethBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethBodyContext {
	var p = new(MethBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_methBody

	return p
}

func (s *MethBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MethBodyContext) CopyFrom(ctx *MethBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MethBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethReturnContext struct {
	*MethBodyContext
}

func NewMethReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethReturnContext {
	var p = new(MethReturnContext)

	p.MethBodyContext = NewEmptyMethBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MethBodyContext))

	return p
}

func (s *MethReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(FGParserRETURN, 0)
}

func (s *MethReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MethReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterMethReturn(s)
	}
}

func (s *MethReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitMethReturn(s)
	}
}

type SprintfContext struct {
	*MethBodyContext
}

func NewSprintfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SprintfContext {
	var p = new(SprintfContext)

	p.MethBodyContext = NewEmptyMethBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MethBodyContext))

	return p
}

func (s *SprintfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SprintfContext) FMT() antlr.TerminalNode {
	return s.GetToken(FGParserFMT, 0)
}

func (s *SprintfContext) SPRINTF() antlr.TerminalNode {
	return s.GetToken(FGParserSPRINTF, 0)
}

func (s *SprintfContext) STRING() antlr.TerminalNode {
	return s.GetToken(FGParserSTRING, 0)
}

func (s *SprintfContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SprintfContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SprintfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSprintf(s)
	}
}

func (s *SprintfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSprintf(s)
	}
}

type SequenceContext struct {
	*MethBodyContext
}

func NewSequenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SequenceContext {
	var p = new(SequenceContext)

	p.MethBodyContext = NewEmptyMethBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MethBodyContext))

	return p
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) Stmt() IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SequenceContext) MethBody() IMethBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethBodyContext)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSequence(s)
	}
}

type CaseSelectContext struct {
	*MethBodyContext
}

func NewCaseSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseSelectContext {
	var p = new(CaseSelectContext)

	p.MethBodyContext = NewEmptyMethBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MethBodyContext))

	return p
}

func (s *CaseSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseSelectContext) SELECT() antlr.TerminalNode {
	return s.GetToken(FGParserSELECT, 0)
}

func (s *CaseSelectContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(FGParserCASE)
}

func (s *CaseSelectContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(FGParserCASE, i)
}

func (s *CaseSelectContext) AllCaseGuard() []ICaseGuardContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseGuardContext)(nil)).Elem())
	var tst = make([]ICaseGuardContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseGuardContext)
		}
	}

	return tst
}

func (s *CaseSelectContext) CaseGuard(i int) ICaseGuardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseGuardContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseGuardContext)
}

func (s *CaseSelectContext) AllMethBody() []IMethBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethBodyContext)(nil)).Elem())
	var tst = make([]IMethBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethBodyContext)
		}
	}

	return tst
}

func (s *CaseSelectContext) MethBody(i int) IMethBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethBodyContext)
}

func (s *CaseSelectContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(FGParserDEFAULT, 0)
}

func (s *CaseSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterCaseSelect(s)
	}
}

func (s *CaseSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitCaseSelect(s)
	}
}

func (p *FGParser) MethBody() (localctx IMethBodyContext) {
	localctx = NewMethBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FGParserRULE_methBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGParserRETURN:
		localctx = NewMethReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(FGParserRETURN)
		}
		{
			p.SetState(88)
			p.expr(0)
		}

	case FGParserT__12, FGParserCLOSE, FGParserGORTN, FGParserMAKE, FGParserNAME:
		localctx = NewSequenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Stmt()
		}
		{
			p.SetState(90)
			p.Match(FGParserT__0)
		}
		{
			p.SetState(91)
			p.MethBody()
		}

	case FGParserFMT:
		localctx = NewSprintfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(FGParserFMT)
		}
		{
			p.SetState(94)
			p.Match(FGParserT__6)
		}
		{
			p.SetState(95)
			p.Match(FGParserSPRINTF)
		}
		{
			p.SetState(96)
			p.Match(FGParserT__1)
		}
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FGParserT__7 || _la == FGParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(FGParserT__8-9))|(1<<(FGParserT__12-9))|(1<<(FGParserMAKE-9))|(1<<(FGParserNAME-9)))) != 0 {
			p.SetState(100)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case FGParserT__8:
				{
					p.SetState(98)
					p.Match(FGParserT__8)
				}

			case FGParserT__12, FGParserMAKE, FGParserNAME:
				{
					p.SetState(99)
					p.expr(0)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(FGParserT__2)
		}

	case FGParserSELECT:
		localctx = NewCaseSelectContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(FGParserSELECT)
		}
		{
			p.SetState(107)
			p.Match(FGParserT__3)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FGParserCASE {
			{
				p.SetState(108)
				p.Match(FGParserCASE)
			}
			{
				p.SetState(109)
				p.CaseGuard()
			}
			{
				p.SetState(110)
				p.Match(FGParserT__10)
			}
			{
				p.SetState(111)
				p.MethBody()
			}

			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGParserDEFAULT {
			{
				p.SetState(118)
				p.Match(FGParserDEFAULT)
			}
			{
				p.SetState(119)
				p.Match(FGParserT__10)
			}
			{
				p.SetState(120)
				p.MethBody()
			}

		}
		{
			p.SetState(123)
			p.Match(FGParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMethDeclContext is an interface to support dynamic dispatch.
type IMethDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethDeclContext differentiates from other interfaces.
	IsMethDeclContext()
}

type MethDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethDeclContext() *MethDeclContext {
	var p = new(MethDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_methDecl
	return p
}

func (*MethDeclContext) IsMethDeclContext() {}

func NewMethDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethDeclContext {
	var p = new(MethDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_methDecl

	return p
}

func (s *MethDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FGParserFUNC, 0)
}

func (s *MethDeclContext) ParamDecl() IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *MethDeclContext) Sig() ISigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISigContext)
}

func (s *MethDeclContext) MethBody() IMethBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethBodyContext)
}

func (s *MethDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterMethDecl(s)
	}
}

func (s *MethDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitMethDecl(s)
	}
}

func (p *FGParser) MethDecl() (localctx IMethDeclContext) {
	localctx = NewMethDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FGParserRULE_methDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(FGParserFUNC)
	}
	{
		p.SetState(127)
		p.Match(FGParserT__1)
	}
	{
		p.SetState(128)
		p.ParamDecl()
	}
	{
		p.SetState(129)
		p.Match(FGParserT__2)
	}
	{
		p.SetState(130)
		p.Sig()
	}
	{
		p.SetState(131)
		p.Match(FGParserT__3)
	}
	{
		p.SetState(132)
		p.MethBody()
	}
	{
		p.SetState(133)
		p.Match(FGParserT__9)
	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) CopyFrom(ctx *TypeLitContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructTypeLitContext struct {
	*TypeLitContext
}

func NewStructTypeLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeLitContext {
	var p = new(StructTypeLitContext)

	p.TypeLitContext = NewEmptyTypeLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeLitContext))

	return p
}

func (s *StructTypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeLitContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(FGParserSTRUCT, 0)
}

func (s *StructTypeLitContext) FieldDecls() IFieldDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclsContext)
}

func (s *StructTypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterStructTypeLit(s)
	}
}

func (s *StructTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitStructTypeLit(s)
	}
}

type InterfaceTypeLitContext struct {
	*TypeLitContext
}

func NewInterfaceTypeLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterfaceTypeLitContext {
	var p = new(InterfaceTypeLitContext)

	p.TypeLitContext = NewEmptyTypeLitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeLitContext))

	return p
}

func (s *InterfaceTypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeLitContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(FGParserINTERFACE, 0)
}

func (s *InterfaceTypeLitContext) Specs() ISpecsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecsContext)
}

func (s *InterfaceTypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterInterfaceTypeLit(s)
	}
}

func (s *InterfaceTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitInterfaceTypeLit(s)
	}
}

func (p *FGParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FGParserRULE_typeLit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGParserSTRUCT:
		localctx = NewStructTypeLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(FGParserSTRUCT)
		}
		{
			p.SetState(136)
			p.Match(FGParserT__3)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGParserNAME {
			{
				p.SetState(137)
				p.FieldDecls()
			}

		}
		{
			p.SetState(140)
			p.Match(FGParserT__9)
		}

	case FGParserINTERFACE:
		localctx = NewInterfaceTypeLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(FGParserINTERFACE)
		}
		{
			p.SetState(142)
			p.Match(FGParserT__3)
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGParserNAME {
			{
				p.SetState(143)
				p.Specs()
			}

		}
		{
			p.SetState(146)
			p.Match(FGParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IChannelTypeLitContext is an interface to support dynamic dispatch.
type IChannelTypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChannelTypeLitContext differentiates from other interfaces.
	IsChannelTypeLitContext()
}

type ChannelTypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannelTypeLitContext() *ChannelTypeLitContext {
	var p = new(ChannelTypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_channelTypeLit
	return p
}

func (*ChannelTypeLitContext) IsChannelTypeLitContext() {}

func NewChannelTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelTypeLitContext {
	var p = new(ChannelTypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_channelTypeLit

	return p
}

func (s *ChannelTypeLitContext) GetParser() antlr.Parser { return s.parser }
func (s *ChannelTypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelTypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChannelTypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterChannelTypeLit(s)
	}
}

func (s *ChannelTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitChannelTypeLit(s)
	}
}

func (p *FGParser) ChannelTypeLit() (localctx IChannelTypeLitContext) {
	localctx = NewChannelTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FGParserRULE_channelTypeLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(FGParserT__11)
		}
		{
			p.SetState(150)
			p.Match(FGParserT__12)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(FGParserT__12)
		}
		{
			p.SetState(152)
			p.Match(FGParserT__11)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(FGParserT__11)
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTyName returns the tyName token.
	GetTyName() antlr.Token

	// SetTyName sets the tyName token.
	SetTyName(antlr.Token)

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tyName antlr.Token
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) GetTyName() antlr.Token { return s.tyName }

func (s *Type_Context) SetTyName(v antlr.Token) { s.tyName = v }

func (s *Type_Context) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *Type_Context) ChannelTypeLit() IChannelTypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannelTypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChannelTypeLitContext)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *FGParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FGParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)

			var _m = p.Match(FGParserNAME)

			localctx.(*Type_Context).tyName = _m
		}

	case FGParserT__11, FGParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.ChannelTypeLit()
		}
		{
			p.SetState(158)
			p.Type_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldDeclsContext is an interface to support dynamic dispatch.
type IFieldDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclsContext differentiates from other interfaces.
	IsFieldDeclsContext()
}

type FieldDeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclsContext() *FieldDeclsContext {
	var p = new(FieldDeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_fieldDecls
	return p
}

func (*FieldDeclsContext) IsFieldDeclsContext() {}

func NewFieldDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclsContext {
	var p = new(FieldDeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_fieldDecls

	return p
}

func (s *FieldDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclsContext) AllFieldDecl() []IFieldDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem())
	var tst = make([]IFieldDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDeclContext)
		}
	}

	return tst
}

func (s *FieldDeclsContext) FieldDecl(i int) IFieldDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclContext)
}

func (s *FieldDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterFieldDecls(s)
	}
}

func (s *FieldDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitFieldDecls(s)
	}
}

func (p *FGParser) FieldDecls() (localctx IFieldDeclsContext) {
	localctx = NewFieldDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FGParserRULE_fieldDecls)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.FieldDecl()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGParserT__0 {
		{
			p.SetState(163)
			p.Match(FGParserT__0)
		}
		{
			p.SetState(164)
			p.FieldDecl()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// GetTyp returns the typ rule contexts.
	GetTyp() IType_Context

	// SetTyp sets the typ rule contexts.
	SetTyp(IType_Context)

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	typ    IType_Context
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) GetField() antlr.Token { return s.field }

func (s *FieldDeclContext) SetField(v antlr.Token) { s.field = v }

func (s *FieldDeclContext) GetTyp() IType_Context { return s.typ }

func (s *FieldDeclContext) SetTyp(v IType_Context) { s.typ = v }

func (s *FieldDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *FieldDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (p *FGParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FGParserRULE_fieldDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)

		var _m = p.Match(FGParserNAME)

		localctx.(*FieldDeclContext).field = _m
	}
	{
		p.SetState(171)

		var _x = p.Type_()

		localctx.(*FieldDeclContext).typ = _x
	}

	return localctx
}

// ISpecsContext is an interface to support dynamic dispatch.
type ISpecsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecsContext differentiates from other interfaces.
	IsSpecsContext()
}

type SpecsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecsContext() *SpecsContext {
	var p = new(SpecsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_specs
	return p
}

func (*SpecsContext) IsSpecsContext() {}

func NewSpecsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecsContext {
	var p = new(SpecsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_specs

	return p
}

func (s *SpecsContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecsContext) AllSpec() []ISpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpecContext)(nil)).Elem())
	var tst = make([]ISpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpecContext)
		}
	}

	return tst
}

func (s *SpecsContext) Spec(i int) ISpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *SpecsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSpecs(s)
	}
}

func (s *SpecsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSpecs(s)
	}
}

func (p *FGParser) Specs() (localctx ISpecsContext) {
	localctx = NewSpecsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FGParserRULE_specs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Spec()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGParserT__0 {
		{
			p.SetState(174)
			p.Match(FGParserT__0)
		}
		{
			p.SetState(175)
			p.Spec()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) CopyFrom(ctx *SpecContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InterfaceSpecContext struct {
	*SpecContext
}

func NewInterfaceSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterfaceSpecContext {
	var p = new(InterfaceSpecContext)

	p.SpecContext = NewEmptySpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SpecContext))

	return p
}

func (s *InterfaceSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceSpecContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *InterfaceSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterInterfaceSpec(s)
	}
}

func (s *InterfaceSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitInterfaceSpec(s)
	}
}

type SigSpecContext struct {
	*SpecContext
}

func NewSigSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SigSpecContext {
	var p = new(SigSpecContext)

	p.SpecContext = NewEmptySpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SpecContext))

	return p
}

func (s *SigSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SigSpecContext) Sig() ISigContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISigContext)
}

func (s *SigSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSigSpec(s)
	}
}

func (s *SigSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSigSpec(s)
	}
}

func (p *FGParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FGParserRULE_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSigSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Sig()
		}

	case 2:
		localctx = NewInterfaceSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Match(FGParserNAME)
		}

	}

	return localctx
}

// ISigContext is an interface to support dynamic dispatch.
type ISigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMeth returns the meth token.
	GetMeth() antlr.Token

	// GetRet returns the ret token.
	GetRet() antlr.Token

	// SetMeth sets the meth token.
	SetMeth(antlr.Token)

	// SetRet sets the ret token.
	SetRet(antlr.Token)

	// IsSigContext differentiates from other interfaces.
	IsSigContext()
}

type SigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	meth   antlr.Token
	ret    antlr.Token
}

func NewEmptySigContext() *SigContext {
	var p = new(SigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_sig
	return p
}

func (*SigContext) IsSigContext() {}

func NewSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SigContext {
	var p = new(SigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_sig

	return p
}

func (s *SigContext) GetParser() antlr.Parser { return s.parser }

func (s *SigContext) GetMeth() antlr.Token { return s.meth }

func (s *SigContext) GetRet() antlr.Token { return s.ret }

func (s *SigContext) SetMeth(v antlr.Token) { s.meth = v }

func (s *SigContext) SetRet(v antlr.Token) { s.ret = v }

func (s *SigContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(FGParserNAME)
}

func (s *SigContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(FGParserNAME, i)
}

func (s *SigContext) Params() IParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *SigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSig(s)
	}
}

func (s *SigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSig(s)
	}
}

func (p *FGParser) Sig() (localctx ISigContext) {
	localctx = NewSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FGParserRULE_sig)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)

		var _m = p.Match(FGParserNAME)

		localctx.(*SigContext).meth = _m
	}
	{
		p.SetState(186)
		p.Match(FGParserT__1)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGParserNAME {
		{
			p.SetState(187)
			p.Params()
		}

	}
	{
		p.SetState(190)
		p.Match(FGParserT__2)
	}
	{
		p.SetState(191)

		var _m = p.Match(FGParserNAME)

		localctx.(*SigContext).ret = _m
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllParamDecl() []IParamDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamDeclContext)(nil)).Elem())
	var tst = make([]IParamDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamDeclContext)
		}
	}

	return tst
}

func (s *ParamsContext) ParamDecl(i int) IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *FGParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FGParserRULE_params)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.ParamDecl()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGParserT__8 {
		{
			p.SetState(194)
			p.Match(FGParserT__8)
		}
		{
			p.SetState(195)
			p.ParamDecl()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVari returns the vari token.
	GetVari() antlr.Token

	// SetVari sets the vari token.
	SetVari(antlr.Token)

	// GetTyp returns the typ rule contexts.
	GetTyp() IType_Context

	// SetTyp sets the typ rule contexts.
	SetTyp(IType_Context)

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vari   antlr.Token
	typ    IType_Context
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) GetVari() antlr.Token { return s.vari }

func (s *ParamDeclContext) SetVari(v antlr.Token) { s.vari = v }

func (s *ParamDeclContext) GetTyp() IType_Context { return s.typ }

func (s *ParamDeclContext) SetTyp(v IType_Context) { s.typ = v }

func (s *ParamDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *ParamDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *FGParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FGParserRULE_paramDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)

		var _m = p.Match(FGParserNAME)

		localctx.(*ParamDeclContext).vari = _m
	}
	{
		p.SetState(202)

		var _x = p.Type_()

		localctx.(*ParamDeclContext).typ = _x
	}

	return localctx
}

// IMethCall_Context is an interface to support dynamic dispatch.
type IMethCall_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethCall_Context differentiates from other interfaces.
	IsMethCall_Context()
}

type MethCall_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethCall_Context() *MethCall_Context {
	var p = new(MethCall_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_methCall_
	return p
}

func (*MethCall_Context) IsMethCall_Context() {}

func NewMethCall_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethCall_Context {
	var p = new(MethCall_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_methCall_

	return p
}

func (s *MethCall_Context) GetParser() antlr.Parser { return s.parser }

func (s *MethCall_Context) CopyFrom(ctx *MethCall_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MethCall_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethCall_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethCallPrimeContext struct {
	*MethCall_Context
	args IExprsContext
}

func NewMethCallPrimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethCallPrimeContext {
	var p = new(MethCallPrimeContext)

	p.MethCall_Context = NewEmptyMethCall_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*MethCall_Context))

	return p
}

func (s *MethCallPrimeContext) GetArgs() IExprsContext { return s.args }

func (s *MethCallPrimeContext) SetArgs(v IExprsContext) { s.args = v }

func (s *MethCallPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethCallPrimeContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *MethCallPrimeContext) Exprs() IExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *MethCallPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterMethCallPrime(s)
	}
}

func (s *MethCallPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitMethCallPrime(s)
	}
}

func (p *FGParser) MethCall_() (localctx IMethCall_Context) {
	localctx = NewMethCall_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FGParserRULE_methCall_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewMethCallPrimeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(FGParserT__6)
	}
	{
		p.SetState(205)
		p.Match(FGParserNAME)
	}
	{
		p.SetState(206)
		p.Match(FGParserT__1)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(FGParserT__12-13))|(1<<(FGParserMAKE-13))|(1<<(FGParserNAME-13)))) != 0 {
		{
			p.SetState(207)

			var _x = p.Exprs()

			localctx.(*MethCallPrimeContext).args = _x
		}

	}
	{
		p.SetState(210)
		p.Match(FGParserT__2)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignmentContext struct {
	*StmtContext
	lvalue antlr.Token
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssignmentContext) GetLvalue() antlr.Token { return s.lvalue }

func (s *AssignmentContext) SetLvalue(v antlr.Token) { s.lvalue = v }

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignmentContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitAssignment(s)
	}
}

type ChDispatchContext struct {
	*StmtContext
}

func NewChDispatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChDispatchContext {
	var p = new(ChDispatchContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ChDispatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChDispatchContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ChDispatchContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ChDispatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterChDispatch(s)
	}
}

func (s *ChDispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitChDispatch(s)
	}
}

type GoroutineContext struct {
	*StmtContext
	recv IExprContext
}

func NewGoroutineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GoroutineContext {
	var p = new(GoroutineContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *GoroutineContext) GetRecv() IExprContext { return s.recv }

func (s *GoroutineContext) SetRecv(v IExprContext) { s.recv = v }

func (s *GoroutineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoroutineContext) GORTN() antlr.TerminalNode {
	return s.GetToken(FGParserGORTN, 0)
}

func (s *GoroutineContext) MethCall_() IMethCall_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethCall_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethCall_Context)
}

func (s *GoroutineContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GoroutineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterGoroutine(s)
	}
}

func (s *GoroutineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitGoroutine(s)
	}
}

type ChCloseContext struct {
	*StmtContext
}

func NewChCloseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChCloseContext {
	var p = new(ChCloseContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ChCloseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChCloseContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(FGParserCLOSE, 0)
}

func (s *ChCloseContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ChCloseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterChClose(s)
	}
}

func (s *ChCloseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitChClose(s)
	}
}

func (p *FGParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FGParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewChCloseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(FGParserCLOSE)
		}
		{
			p.SetState(213)
			p.Match(FGParserT__1)
		}
		{
			p.SetState(214)
			p.expr(0)
		}
		{
			p.SetState(215)
			p.Match(FGParserT__2)
		}

	case 2:
		localctx = NewChDispatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.expr(0)
		}
		{
			p.SetState(218)
			p.Match(FGParserT__12)
		}
		{
			p.SetState(219)
			p.expr(0)
		}

	case 3:
		localctx = NewGoroutineContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.Match(FGParserGORTN)
		}
		{
			p.SetState(222)

			var _x = p.expr(0)

			localctx.(*GoroutineContext).recv = _x
		}
		{
			p.SetState(223)
			p.MethCall_()
		}

	case 4:
		localctx = NewAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(225)

			var _m = p.Match(FGParserNAME)

			localctx.(*AssignmentContext).lvalue = _m
		}
		{
			p.SetState(226)
			p.Match(FGParserT__13)
		}
		{
			p.SetState(227)
			p.expr(0)
		}

	}

	return localctx
}

// ICaseGuardContext is an interface to support dynamic dispatch.
type ICaseGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseGuardContext differentiates from other interfaces.
	IsCaseGuardContext()
}

type CaseGuardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseGuardContext() *CaseGuardContext {
	var p = new(CaseGuardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_caseGuard
	return p
}

func (*CaseGuardContext) IsCaseGuardContext() {}

func NewCaseGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseGuardContext {
	var p = new(CaseGuardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_caseGuard

	return p
}

func (s *CaseGuardContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseGuardContext) CopyFrom(ctx *CaseGuardContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CaseGuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseGuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GuardReceiveContext struct {
	*CaseGuardContext
	lvalue antlr.Token
}

func NewGuardReceiveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardReceiveContext {
	var p = new(GuardReceiveContext)

	p.CaseGuardContext = NewEmptyCaseGuardContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CaseGuardContext))

	return p
}

func (s *GuardReceiveContext) GetLvalue() antlr.Token { return s.lvalue }

func (s *GuardReceiveContext) SetLvalue(v antlr.Token) { s.lvalue = v }

func (s *GuardReceiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardReceiveContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardReceiveContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *GuardReceiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterGuardReceive(s)
	}
}

func (s *GuardReceiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitGuardReceive(s)
	}
}

type GuardDispatchContext struct {
	*CaseGuardContext
}

func NewGuardDispatchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardDispatchContext {
	var p = new(GuardDispatchContext)

	p.CaseGuardContext = NewEmptyCaseGuardContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CaseGuardContext))

	return p
}

func (s *GuardDispatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardDispatchContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GuardDispatchContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardDispatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterGuardDispatch(s)
	}
}

func (s *GuardDispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitGuardDispatch(s)
	}
}

func (p *FGParser) CaseGuard() (localctx ICaseGuardContext) {
	localctx = NewCaseGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FGParserRULE_caseGuard)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGuardDispatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.expr(0)
		}
		{
			p.SetState(231)
			p.Match(FGParserT__12)
		}
		{
			p.SetState(232)
			p.expr(0)
		}

	case 2:
		localctx = NewGuardReceiveContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)

			var _m = p.Match(FGParserNAME)

			localctx.(*GuardReceiveContext).lvalue = _m
		}
		{
			p.SetState(235)
			p.Match(FGParserT__13)
		}
		{
			p.SetState(236)
			p.Match(FGParserT__12)
		}
		{
			p.SetState(237)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallContext struct {
	*ExprContext
	recv IExprContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallContext) GetRecv() IExprContext { return s.recv }

func (s *CallContext) SetRecv(v IExprContext) { s.recv = v }

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) MethCall_() IMethCall_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethCall_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethCall_Context)
}

func (s *CallContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitCall(s)
	}
}

type VariableContext struct {
	*ExprContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitVariable(s)
	}
}

type AssertContext struct {
	*ExprContext
}

func NewAssertContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssertContext {
	var p = new(AssertContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssertContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *AssertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterAssert(s)
	}
}

func (s *AssertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitAssert(s)
	}
}

type SelectContext struct {
	*ExprContext
}

func NewSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectContext {
	var p = new(SelectContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelectContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitSelect(s)
	}
}

type ChanRecvContext struct {
	*ExprContext
}

func NewChanRecvContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChanRecvContext {
	var p = new(ChanRecvContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ChanRecvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChanRecvContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ChanRecvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterChanRecv(s)
	}
}

func (s *ChanRecvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitChanRecv(s)
	}
}

type StructLitContext struct {
	*ExprContext
}

func NewStructLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructLitContext {
	var p = new(StructLitContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StructLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructLitContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGParserNAME, 0)
}

func (s *StructLitContext) Exprs() IExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *StructLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterStructLit(s)
	}
}

func (s *StructLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitStructLit(s)
	}
}

type MakeChanContext struct {
	*ExprContext
}

func NewMakeChanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MakeChanContext {
	var p = new(MakeChanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MakeChanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MakeChanContext) MAKE() antlr.TerminalNode {
	return s.GetToken(FGParserMAKE, 0)
}

func (s *MakeChanContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MakeChanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterMakeChan(s)
	}
}

func (s *MakeChanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitMakeChan(s)
	}
}

func (p *FGParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FGParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, FGParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(241)
			p.Match(FGParserNAME)
		}

	case 2:
		localctx = NewStructLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(FGParserNAME)
		}
		{
			p.SetState(243)
			p.Match(FGParserT__3)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(FGParserT__12-13))|(1<<(FGParserMAKE-13))|(1<<(FGParserNAME-13)))) != 0 {
			{
				p.SetState(244)
				p.Exprs()
			}

		}
		{
			p.SetState(247)
			p.Match(FGParserT__9)
		}

	case 3:
		localctx = NewMakeChanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(248)
			p.Match(FGParserMAKE)
		}
		{
			p.SetState(249)
			p.Match(FGParserT__1)
		}
		{
			p.SetState(250)
			p.Type_()
		}
		{
			p.SetState(251)
			p.Match(FGParserT__2)
		}

	case 4:
		localctx = NewChanRecvContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(253)
			p.Match(FGParserT__12)
		}
		{
			p.SetState(254)
			p.expr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FGParserRULE_expr)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(258)
					p.Match(FGParserT__6)
				}
				{
					p.SetState(259)
					p.Match(FGParserNAME)
				}

			case 2:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallContext).recv = _prevctx

				p.PushNewRecursionContext(localctx, _startState, FGParserRULE_expr)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(261)
					p.MethCall_()
				}

			case 3:
				localctx = NewAssertContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FGParserRULE_expr)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(263)
					p.Match(FGParserT__6)
				}
				{
					p.SetState(264)
					p.Match(FGParserT__1)
				}
				{
					p.SetState(265)
					p.Match(FGParserNAME)
				}
				{
					p.SetState(266)
					p.Match(FGParserT__2)
				}

			}

		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IExprsContext is an interface to support dynamic dispatch.
type IExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprsContext differentiates from other interfaces.
	IsExprsContext()
}

type ExprsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprsContext() *ExprsContext {
	var p = new(ExprsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGParserRULE_exprs
	return p
}

func (*ExprsContext) IsExprsContext() {}

func NewExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprsContext {
	var p = new(ExprsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGParserRULE_exprs

	return p
}

func (s *ExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.EnterExprs(s)
	}
}

func (s *ExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGListener); ok {
		listenerT.ExitExprs(s)
	}
}

func (p *FGParser) Exprs() (localctx IExprsContext) {
	localctx = NewExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FGParserRULE_exprs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.expr(0)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGParserT__8 {
		{
			p.SetState(273)
			p.Match(FGParserT__8)
		}
		{
			p.SetState(274)
			p.expr(0)
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *FGParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FGParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
