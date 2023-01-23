// Generated from parser/FGG.g4 by ANTLR 4.7.

package parser // FGG

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 334,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 56, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 62, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 68, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 73, 10, 4, 12, 4, 14, 4, 76, 11,
	4, 3, 5, 3, 5, 3, 5, 5, 5, 81, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7,
	6, 88, 10, 6, 12, 6, 14, 6, 91, 11, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 5, 8, 102, 10, 8, 3, 8, 5, 8, 105, 10, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 5, 8, 124, 10, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5,
	9, 131, 10, 9, 3, 9, 3, 9, 6, 9, 135, 10, 9, 13, 9, 14, 9, 136, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 157, 10, 11, 12, 11,
	14, 11, 160, 11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 170, 10, 11, 12, 11, 14, 11, 173, 11, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 178, 10, 11, 3, 11, 5, 11, 181, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 199, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 205,
	10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 210, 10, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 222, 10, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 5, 16, 238, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	244, 10, 16, 3, 16, 5, 16, 247, 10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 252,
	10, 17, 12, 17, 14, 17, 255, 11, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 7, 19, 263, 10, 19, 12, 19, 14, 19, 266, 11, 19, 3, 20, 3, 20, 5,
	20, 270, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 276, 10, 21, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 7, 22, 284, 10, 22, 12, 22, 14, 22,
	287, 11, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 297, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 5, 24, 308, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 321, 10, 24, 12, 24, 14, 24, 324,
	11, 24, 3, 25, 3, 25, 3, 25, 7, 25, 329, 10, 25, 12, 25, 14, 25, 332, 11,
	25, 3, 25, 2, 3, 46, 26, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 2, 3, 4, 2, 15, 15, 40, 40,
	2, 351, 2, 55, 3, 2, 2, 2, 4, 67, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 77,
	3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 92, 3, 2, 2, 2, 14, 95, 3, 2, 2, 2,
	16, 134, 3, 2, 2, 2, 18, 138, 3, 2, 2, 2, 20, 180, 3, 2, 2, 2, 22, 198,
	3, 2, 2, 2, 24, 200, 3, 2, 2, 2, 26, 221, 3, 2, 2, 2, 28, 223, 3, 2, 2,
	2, 30, 246, 3, 2, 2, 2, 32, 248, 3, 2, 2, 2, 34, 256, 3, 2, 2, 2, 36, 259,
	3, 2, 2, 2, 38, 269, 3, 2, 2, 2, 40, 271, 3, 2, 2, 2, 42, 280, 3, 2, 2,
	2, 44, 288, 3, 2, 2, 2, 46, 307, 3, 2, 2, 2, 48, 325, 3, 2, 2, 2, 50, 51,
	7, 3, 2, 2, 51, 56, 7, 4, 2, 2, 52, 53, 7, 4, 2, 2, 53, 56, 7, 3, 2, 2,
	54, 56, 7, 3, 2, 2, 55, 50, 3, 2, 2, 2, 55, 52, 3, 2, 2, 2, 55, 54, 3,
	2, 2, 2, 56, 3, 3, 2, 2, 2, 57, 68, 7, 36, 2, 2, 58, 59, 7, 36, 2, 2, 59,
	61, 7, 5, 2, 2, 60, 62, 5, 6, 4, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2,
	2, 62, 63, 3, 2, 2, 2, 63, 68, 7, 6, 2, 2, 64, 65, 5, 2, 2, 2, 65, 66,
	5, 4, 3, 2, 66, 68, 3, 2, 2, 2, 67, 57, 3, 2, 2, 2, 67, 58, 3, 2, 2, 2,
	67, 64, 3, 2, 2, 2, 68, 5, 3, 2, 2, 2, 69, 74, 5, 4, 3, 2, 70, 71, 7, 7,
	2, 2, 71, 73, 5, 4, 3, 2, 72, 70, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72,
	3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 7, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2,
	77, 78, 7, 5, 2, 2, 78, 80, 7, 25, 2, 2, 79, 81, 5, 10, 6, 2, 80, 79, 3,
	2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 7, 6, 2, 2, 83,
	9, 3, 2, 2, 2, 84, 89, 5, 12, 7, 2, 85, 86, 7, 7, 2, 2, 86, 88, 5, 12,
	7, 2, 87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 11, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 36, 2, 2,
	93, 94, 5, 4, 3, 2, 94, 13, 3, 2, 2, 2, 95, 96, 7, 22, 2, 2, 96, 97, 7,
	21, 2, 2, 97, 101, 7, 8, 2, 2, 98, 99, 7, 26, 2, 2, 99, 100, 7, 40, 2,
	2, 100, 102, 7, 8, 2, 2, 101, 98, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102,
	104, 3, 2, 2, 2, 103, 105, 5, 16, 9, 2, 104, 103, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 7, 19, 2, 2, 107, 108, 7, 21,
	2, 2, 108, 109, 7, 9, 2, 2, 109, 110, 7, 10, 2, 2, 110, 123, 7, 11, 2,
	2, 111, 112, 7, 12, 2, 2, 112, 113, 7, 13, 2, 2, 113, 124, 5, 46, 24, 2,
	114, 115, 7, 27, 2, 2, 115, 116, 7, 14, 2, 2, 116, 117, 7, 28, 2, 2, 117,
	118, 7, 9, 2, 2, 118, 119, 7, 15, 2, 2, 119, 120, 7, 7, 2, 2, 120, 121,
	5, 46, 24, 2, 121, 122, 7, 10, 2, 2, 122, 124, 3, 2, 2, 2, 123, 111, 3,
	2, 2, 2, 123, 114, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 7, 16, 2,
	2, 126, 127, 7, 2, 2, 3, 127, 15, 3, 2, 2, 2, 128, 131, 5, 18, 10, 2, 129,
	131, 5, 28, 15, 2, 130, 128, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2, 131, 132,
	3, 2, 2, 2, 132, 133, 7, 8, 2, 2, 133, 135, 3, 2, 2, 2, 134, 130, 3, 2,
	2, 2, 135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2,
	137, 17, 3, 2, 2, 2, 138, 139, 7, 25, 2, 2, 139, 140, 7, 36, 2, 2, 140,
	141, 5, 8, 5, 2, 141, 142, 5, 30, 16, 2, 142, 19, 3, 2, 2, 2, 143, 144,
	7, 23, 2, 2, 144, 181, 5, 46, 24, 2, 145, 146, 5, 22, 12, 2, 146, 147,
	7, 8, 2, 2, 147, 148, 5, 20, 11, 2, 148, 181, 3, 2, 2, 2, 149, 150, 7,
	27, 2, 2, 150, 151, 7, 14, 2, 2, 151, 152, 7, 29, 2, 2, 152, 153, 7, 9,
	2, 2, 153, 158, 9, 2, 2, 2, 154, 157, 7, 7, 2, 2, 155, 157, 5, 46, 24,
	2, 156, 154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158,
	156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 3, 2, 2, 2, 160, 158,
	3, 2, 2, 2, 161, 181, 7, 10, 2, 2, 162, 163, 7, 33, 2, 2, 163, 171, 7,
	11, 2, 2, 164, 165, 7, 35, 2, 2, 165, 166, 5, 26, 14, 2, 166, 167, 7, 17,
	2, 2, 167, 168, 5, 20, 11, 2, 168, 170, 3, 2, 2, 2, 169, 164, 3, 2, 2,
	2, 170, 173, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172,
	177, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 174, 175, 7, 34, 2, 2, 175, 176,
	7, 17, 2, 2, 176, 178, 5, 20, 11, 2, 177, 174, 3, 2, 2, 2, 177, 178, 3,
	2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 7, 16, 2, 2, 180, 143, 3, 2, 2,
	2, 180, 145, 3, 2, 2, 2, 180, 149, 3, 2, 2, 2, 180, 162, 3, 2, 2, 2, 181,
	21, 3, 2, 2, 2, 182, 183, 7, 30, 2, 2, 183, 184, 7, 9, 2, 2, 184, 185,
	5, 46, 24, 2, 185, 186, 7, 10, 2, 2, 186, 199, 3, 2, 2, 2, 187, 188, 5,
	46, 24, 2, 188, 189, 7, 4, 2, 2, 189, 190, 5, 46, 24, 2, 190, 199, 3, 2,
	2, 2, 191, 192, 7, 31, 2, 2, 192, 193, 5, 46, 24, 2, 193, 194, 5, 24, 13,
	2, 194, 199, 3, 2, 2, 2, 195, 196, 7, 36, 2, 2, 196, 197, 7, 18, 2, 2,
	197, 199, 5, 46, 24, 2, 198, 182, 3, 2, 2, 2, 198, 187, 3, 2, 2, 2, 198,
	191, 3, 2, 2, 2, 198, 195, 3, 2, 2, 2, 199, 23, 3, 2, 2, 2, 200, 201, 7,
	14, 2, 2, 201, 202, 7, 36, 2, 2, 202, 204, 7, 5, 2, 2, 203, 205, 5, 6,
	4, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2,
	206, 207, 7, 6, 2, 2, 207, 209, 7, 9, 2, 2, 208, 210, 5, 48, 25, 2, 209,
	208, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212,
	7, 10, 2, 2, 212, 25, 3, 2, 2, 2, 213, 214, 5, 46, 24, 2, 214, 215, 7,
	4, 2, 2, 215, 216, 5, 46, 24, 2, 216, 222, 3, 2, 2, 2, 217, 218, 7, 36,
	2, 2, 218, 219, 7, 18, 2, 2, 219, 220, 7, 4, 2, 2, 220, 222, 5, 46, 24,
	2, 221, 213, 3, 2, 2, 2, 221, 217, 3, 2, 2, 2, 222, 27, 3, 2, 2, 2, 223,
	224, 7, 19, 2, 2, 224, 225, 7, 9, 2, 2, 225, 226, 7, 36, 2, 2, 226, 227,
	7, 36, 2, 2, 227, 228, 5, 8, 5, 2, 228, 229, 7, 10, 2, 2, 229, 230, 5,
	40, 21, 2, 230, 231, 7, 11, 2, 2, 231, 232, 5, 20, 11, 2, 232, 233, 7,
	16, 2, 2, 233, 29, 3, 2, 2, 2, 234, 235, 7, 24, 2, 2, 235, 237, 7, 11,
	2, 2, 236, 238, 5, 32, 17, 2, 237, 236, 3, 2, 2, 2, 237, 238, 3, 2, 2,
	2, 238, 239, 3, 2, 2, 2, 239, 247, 7, 16, 2, 2, 240, 241, 7, 20, 2, 2,
	241, 243, 7, 11, 2, 2, 242, 244, 5, 36, 19, 2, 243, 242, 3, 2, 2, 2, 243,
	244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 247, 7, 16, 2, 2, 246, 234,
	3, 2, 2, 2, 246, 240, 3, 2, 2, 2, 247, 31, 3, 2, 2, 2, 248, 253, 5, 34,
	18, 2, 249, 250, 7, 8, 2, 2, 250, 252, 5, 34, 18, 2, 251, 249, 3, 2, 2,
	2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254,
	33, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 257, 7, 36, 2, 2, 257, 258,
	5, 4, 3, 2, 258, 35, 3, 2, 2, 2, 259, 264, 5, 38, 20, 2, 260, 261, 7, 8,
	2, 2, 261, 263, 5, 38, 20, 2, 262, 260, 3, 2, 2, 2, 263, 266, 3, 2, 2,
	2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 37, 3, 2, 2, 2, 266,
	264, 3, 2, 2, 2, 267, 270, 5, 40, 21, 2, 268, 270, 5, 4, 3, 2, 269, 267,
	3, 2, 2, 2, 269, 268, 3, 2, 2, 2, 270, 39, 3, 2, 2, 2, 271, 272, 7, 36,
	2, 2, 272, 273, 5, 8, 5, 2, 273, 275, 7, 9, 2, 2, 274, 276, 5, 42, 22,
	2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277,
	278, 7, 10, 2, 2, 278, 279, 5, 4, 3, 2, 279, 41, 3, 2, 2, 2, 280, 285,
	5, 44, 23, 2, 281, 282, 7, 7, 2, 2, 282, 284, 5, 44, 23, 2, 283, 281, 3,
	2, 2, 2, 284, 287, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2,
	2, 286, 43, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 288, 289, 7, 36, 2, 2, 289,
	290, 5, 4, 3, 2, 290, 45, 3, 2, 2, 2, 291, 292, 8, 24, 1, 2, 292, 308,
	7, 36, 2, 2, 293, 294, 5, 4, 3, 2, 294, 296, 7, 11, 2, 2, 295, 297, 5,
	48, 25, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2,
	2, 2, 298, 299, 7, 16, 2, 2, 299, 308, 3, 2, 2, 2, 300, 301, 7, 32, 2,
	2, 301, 302, 7, 9, 2, 2, 302, 303, 5, 4, 3, 2, 303, 304, 7, 10, 2, 2, 304,
	308, 3, 2, 2, 2, 305, 306, 7, 4, 2, 2, 306, 308, 5, 46, 24, 3, 307, 291,
	3, 2, 2, 2, 307, 293, 3, 2, 2, 2, 307, 300, 3, 2, 2, 2, 307, 305, 3, 2,
	2, 2, 308, 322, 3, 2, 2, 2, 309, 310, 12, 7, 2, 2, 310, 311, 7, 14, 2,
	2, 311, 321, 7, 36, 2, 2, 312, 313, 12, 6, 2, 2, 313, 321, 5, 24, 13, 2,
	314, 315, 12, 5, 2, 2, 315, 316, 7, 14, 2, 2, 316, 317, 7, 9, 2, 2, 317,
	318, 5, 4, 3, 2, 318, 319, 7, 10, 2, 2, 319, 321, 3, 2, 2, 2, 320, 309,
	3, 2, 2, 2, 320, 312, 3, 2, 2, 2, 320, 314, 3, 2, 2, 2, 321, 324, 3, 2,
	2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 47, 3, 2, 2, 2,
	324, 322, 3, 2, 2, 2, 325, 330, 5, 46, 24, 2, 326, 327, 7, 7, 2, 2, 327,
	329, 5, 46, 24, 2, 328, 326, 3, 2, 2, 2, 329, 332, 3, 2, 2, 2, 330, 328,
	3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 49, 3, 2, 2, 2, 332, 330, 3, 2,
	2, 2, 35, 55, 61, 67, 74, 80, 89, 101, 104, 123, 130, 136, 156, 158, 171,
	177, 180, 198, 204, 209, 221, 237, 243, 246, 253, 264, 269, 275, 285, 296,
	307, 320, 322, 330,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'chan'", "'<-'", "'['", "']'", "','", "';'", "'('", "')'", "'{'",
	"'_'", "'='", "'.'", "'\"%#v\"'", "'}'", "':'", "':='", "'func'", "'interface'",
	"'main'", "'package'", "'return'", "'struct'", "'type'", "'import'", "'fmt'",
	"'Printf'", "'Sprintf'", "'close'", "'go'", "'make'", "'select'", "'default'",
	"'case'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "FUNC",
	"INTERFACE", "MAIN", "PACKAGE", "RETURN", "STRUCT", "TYPE", "IMPORT", "FMT",
	"PRINTF", "SPRINTF", "CLOSE", "GORTN", "MAKE", "SELECT", "DEFAULT", "CASE",
	"NAME", "WHITESPACE", "COMMENT", "LINE_COMMENT", "STRING",
}

var ruleNames = []string{
	"channelTypeLit", "typ", "typs", "typeFormals", "typeFDecls", "typeFDecl",
	"program", "decls", "typeDecl", "methBody", "stmt", "methCall_", "caseGuard",
	"methDecl", "typeLit", "fieldDecls", "fieldDecl", "specs", "spec", "sig",
	"params", "paramDecl", "expr", "exprs",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FGGParser struct {
	*antlr.BaseParser
}

func NewFGGParser(input antlr.TokenStream) *FGGParser {
	this := new(FGGParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FGG.g4"

	return this
}

// FGGParser tokens.
const (
	FGGParserEOF          = antlr.TokenEOF
	FGGParserT__0         = 1
	FGGParserT__1         = 2
	FGGParserT__2         = 3
	FGGParserT__3         = 4
	FGGParserT__4         = 5
	FGGParserT__5         = 6
	FGGParserT__6         = 7
	FGGParserT__7         = 8
	FGGParserT__8         = 9
	FGGParserT__9         = 10
	FGGParserT__10        = 11
	FGGParserT__11        = 12
	FGGParserT__12        = 13
	FGGParserT__13        = 14
	FGGParserT__14        = 15
	FGGParserT__15        = 16
	FGGParserFUNC         = 17
	FGGParserINTERFACE    = 18
	FGGParserMAIN         = 19
	FGGParserPACKAGE      = 20
	FGGParserRETURN       = 21
	FGGParserSTRUCT       = 22
	FGGParserTYPE         = 23
	FGGParserIMPORT       = 24
	FGGParserFMT          = 25
	FGGParserPRINTF       = 26
	FGGParserSPRINTF      = 27
	FGGParserCLOSE        = 28
	FGGParserGORTN        = 29
	FGGParserMAKE         = 30
	FGGParserSELECT       = 31
	FGGParserDEFAULT      = 32
	FGGParserCASE         = 33
	FGGParserNAME         = 34
	FGGParserWHITESPACE   = 35
	FGGParserCOMMENT      = 36
	FGGParserLINE_COMMENT = 37
	FGGParserSTRING       = 38
)

// FGGParser rules.
const (
	FGGParserRULE_channelTypeLit = 0
	FGGParserRULE_typ            = 1
	FGGParserRULE_typs           = 2
	FGGParserRULE_typeFormals    = 3
	FGGParserRULE_typeFDecls     = 4
	FGGParserRULE_typeFDecl      = 5
	FGGParserRULE_program        = 6
	FGGParserRULE_decls          = 7
	FGGParserRULE_typeDecl       = 8
	FGGParserRULE_methBody       = 9
	FGGParserRULE_stmt           = 10
	FGGParserRULE_methCall_      = 11
	FGGParserRULE_caseGuard      = 12
	FGGParserRULE_methDecl       = 13
	FGGParserRULE_typeLit        = 14
	FGGParserRULE_fieldDecls     = 15
	FGGParserRULE_fieldDecl      = 16
	FGGParserRULE_specs          = 17
	FGGParserRULE_spec           = 18
	FGGParserRULE_sig            = 19
	FGGParserRULE_params         = 20
	FGGParserRULE_paramDecl      = 21
	FGGParserRULE_expr           = 22
	FGGParserRULE_exprs          = 23
)

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
	p.RuleIndex = FGGParserRULE_channelTypeLit
	return p
}

func (*ChannelTypeLitContext) IsChannelTypeLitContext() {}

func NewChannelTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelTypeLitContext {
	var p = new(ChannelTypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_channelTypeLit

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterChannelTypeLit(s)
	}
}

func (s *ChannelTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitChannelTypeLit(s)
	}
}

func (p *FGGParser) ChannelTypeLit() (localctx IChannelTypeLitContext) {
	localctx = NewChannelTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FGGParserRULE_channelTypeLit)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(FGGParserT__0)
		}
		{
			p.SetState(49)
			p.Match(FGGParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Match(FGGParserT__1)
		}
		{
			p.SetState(51)
			p.Match(FGGParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.Match(FGGParserT__0)
		}

	}

	return localctx
}

// ITypContext is an interface to support dynamic dispatch.
type ITypContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypContext differentiates from other interfaces.
	IsTypContext()
}

type TypContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypContext() *TypContext {
	var p = new(TypContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_typ
	return p
}

func (*TypContext) IsTypContext() {}

func NewTypContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypContext {
	var p = new(TypContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typ

	return p
}

func (s *TypContext) GetParser() antlr.Parser { return s.parser }

func (s *TypContext) CopyFrom(ctx *TypContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeNameContext struct {
	*TypContext
}

func NewTypeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeNameContext {
	var p = new(TypeNameContext)

	p.TypContext = NewEmptyTypContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypContext))

	return p
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *TypeNameContext) Typs() ITypsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypsContext)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeName(s)
	}
}

type TypeParamContext struct {
	*TypContext
}

func NewTypeParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamContext {
	var p = new(TypeParamContext)

	p.TypContext = NewEmptyTypContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypContext))

	return p
}

func (s *TypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *TypeParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeParam(s)
	}
}

func (s *TypeParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeParam(s)
	}
}

type ChannelTypeContext struct {
	*TypContext
}

func NewChannelTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChannelTypeContext {
	var p = new(ChannelTypeContext)

	p.TypContext = NewEmptyTypContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypContext))

	return p
}

func (s *ChannelTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelTypeContext) ChannelTypeLit() IChannelTypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannelTypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChannelTypeLitContext)
}

func (s *ChannelTypeContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ChannelTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterChannelType(s)
	}
}

func (s *ChannelTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitChannelType(s)
	}
}

func (p *FGGParser) Typ() (localctx ITypContext) {
	localctx = NewTypContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FGGParserRULE_typ)
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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Match(FGGParserNAME)
		}

	case 2:
		localctx = NewTypeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(FGGParserNAME)
		}
		{
			p.SetState(57)
			p.Match(FGGParserT__2)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGGParserT__0 || _la == FGGParserT__1 || _la == FGGParserNAME {
			{
				p.SetState(58)
				p.Typs()
			}

		}
		{
			p.SetState(61)
			p.Match(FGGParserT__3)
		}

	case 3:
		localctx = NewChannelTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.ChannelTypeLit()
		}
		{
			p.SetState(63)
			p.Typ()
		}

	}

	return localctx
}

// ITypsContext is an interface to support dynamic dispatch.
type ITypsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypsContext differentiates from other interfaces.
	IsTypsContext()
}

type TypsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypsContext() *TypsContext {
	var p = new(TypsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_typs
	return p
}

func (*TypsContext) IsTypsContext() {}

func NewTypsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypsContext {
	var p = new(TypsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typs

	return p
}

func (s *TypsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypsContext) AllTyp() []ITypContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypContext)(nil)).Elem())
	var tst = make([]ITypContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypContext)
		}
	}

	return tst
}

func (s *TypsContext) Typ(i int) ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *TypsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTyps(s)
	}
}

func (s *TypsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTyps(s)
	}
}

func (p *FGGParser) Typs() (localctx ITypsContext) {
	localctx = NewTypsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FGGParserRULE_typs)
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
		p.SetState(67)
		p.Typ()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__4 {
		{
			p.SetState(68)
			p.Match(FGGParserT__4)
		}
		{
			p.SetState(69)
			p.Typ()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeFormalsContext is an interface to support dynamic dispatch.
type ITypeFormalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFormalsContext differentiates from other interfaces.
	IsTypeFormalsContext()
}

type TypeFormalsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFormalsContext() *TypeFormalsContext {
	var p = new(TypeFormalsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_typeFormals
	return p
}

func (*TypeFormalsContext) IsTypeFormalsContext() {}

func NewTypeFormalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFormalsContext {
	var p = new(TypeFormalsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typeFormals

	return p
}

func (s *TypeFormalsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFormalsContext) TYPE() antlr.TerminalNode {
	return s.GetToken(FGGParserTYPE, 0)
}

func (s *TypeFormalsContext) TypeFDecls() ITypeFDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFDeclsContext)
}

func (s *TypeFormalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFormalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFormalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeFormals(s)
	}
}

func (s *TypeFormalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeFormals(s)
	}
}

func (p *FGGParser) TypeFormals() (localctx ITypeFormalsContext) {
	localctx = NewTypeFormalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FGGParserRULE_typeFormals)
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
		p.SetState(75)
		p.Match(FGGParserT__2)
	}
	{
		p.SetState(76)
		p.Match(FGGParserTYPE)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGGParserNAME {
		{
			p.SetState(77)
			p.TypeFDecls()
		}

	}
	{
		p.SetState(80)
		p.Match(FGGParserT__3)
	}

	return localctx
}

// ITypeFDeclsContext is an interface to support dynamic dispatch.
type ITypeFDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFDeclsContext differentiates from other interfaces.
	IsTypeFDeclsContext()
}

type TypeFDeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFDeclsContext() *TypeFDeclsContext {
	var p = new(TypeFDeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_typeFDecls
	return p
}

func (*TypeFDeclsContext) IsTypeFDeclsContext() {}

func NewTypeFDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFDeclsContext {
	var p = new(TypeFDeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typeFDecls

	return p
}

func (s *TypeFDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFDeclsContext) AllTypeFDecl() []ITypeFDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeFDeclContext)(nil)).Elem())
	var tst = make([]ITypeFDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeFDeclContext)
		}
	}

	return tst
}

func (s *TypeFDeclsContext) TypeFDecl(i int) ITypeFDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeFDeclContext)
}

func (s *TypeFDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeFDecls(s)
	}
}

func (s *TypeFDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeFDecls(s)
	}
}

func (p *FGGParser) TypeFDecls() (localctx ITypeFDeclsContext) {
	localctx = NewTypeFDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FGGParserRULE_typeFDecls)
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
		p.SetState(82)
		p.TypeFDecl()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__4 {
		{
			p.SetState(83)
			p.Match(FGGParserT__4)
		}
		{
			p.SetState(84)
			p.TypeFDecl()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeFDeclContext is an interface to support dynamic dispatch.
type ITypeFDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeFDeclContext differentiates from other interfaces.
	IsTypeFDeclContext()
}

type TypeFDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeFDeclContext() *TypeFDeclContext {
	var p = new(TypeFDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_typeFDecl
	return p
}

func (*TypeFDeclContext) IsTypeFDeclContext() {}

func NewTypeFDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeFDeclContext {
	var p = new(TypeFDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typeFDecl

	return p
}

func (s *TypeFDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeFDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *TypeFDeclContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *TypeFDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeFDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeFDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeFDecl(s)
	}
}

func (s *TypeFDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeFDecl(s)
	}
}

func (p *FGGParser) TypeFDecl() (localctx ITypeFDeclContext) {
	localctx = NewTypeFDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FGGParserRULE_typeFDecl)

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
		p.SetState(90)
		p.Match(FGGParserNAME)
	}
	{
		p.SetState(91)
		p.Typ()
	}

	return localctx
}

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
	p.RuleIndex = FGGParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(FGGParserPACKAGE, 0)
}

func (s *ProgramContext) AllMAIN() []antlr.TerminalNode {
	return s.GetTokens(FGGParserMAIN)
}

func (s *ProgramContext) MAIN(i int) antlr.TerminalNode {
	return s.GetToken(FGGParserMAIN, i)
}

func (s *ProgramContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FGGParserFUNC, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FGGParserEOF, 0)
}

func (s *ProgramContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramContext) FMT() antlr.TerminalNode {
	return s.GetToken(FGGParserFMT, 0)
}

func (s *ProgramContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(FGGParserPRINTF, 0)
}

func (s *ProgramContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(FGGParserIMPORT, 0)
}

func (s *ProgramContext) STRING() antlr.TerminalNode {
	return s.GetToken(FGGParserSTRING, 0)
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *FGGParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FGGParserRULE_program)
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
		p.SetState(93)
		p.Match(FGGParserPACKAGE)
	}
	{
		p.SetState(94)
		p.Match(FGGParserMAIN)
	}
	{
		p.SetState(95)
		p.Match(FGGParserT__5)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGGParserIMPORT {
		{
			p.SetState(96)
			p.Match(FGGParserIMPORT)
		}
		{
			p.SetState(97)
			p.Match(FGGParserSTRING)
		}
		{
			p.SetState(98)
			p.Match(FGGParserT__5)
		}

	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(101)
			p.Decls()
		}

	}
	{
		p.SetState(104)
		p.Match(FGGParserFUNC)
	}
	{
		p.SetState(105)
		p.Match(FGGParserMAIN)
	}
	{
		p.SetState(106)
		p.Match(FGGParserT__6)
	}
	{
		p.SetState(107)
		p.Match(FGGParserT__7)
	}
	{
		p.SetState(108)
		p.Match(FGGParserT__8)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGGParserT__9:
		{
			p.SetState(109)
			p.Match(FGGParserT__9)
		}
		{
			p.SetState(110)
			p.Match(FGGParserT__10)
		}
		{
			p.SetState(111)
			p.expr(0)
		}

	case FGGParserFMT:
		{
			p.SetState(112)
			p.Match(FGGParserFMT)
		}
		{
			p.SetState(113)
			p.Match(FGGParserT__11)
		}
		{
			p.SetState(114)
			p.Match(FGGParserPRINTF)
		}
		{
			p.SetState(115)
			p.Match(FGGParserT__6)
		}
		{
			p.SetState(116)
			p.Match(FGGParserT__12)
		}
		{
			p.SetState(117)
			p.Match(FGGParserT__4)
		}
		{
			p.SetState(118)
			p.expr(0)
		}
		{
			p.SetState(119)
			p.Match(FGGParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(123)
		p.Match(FGGParserT__13)
	}
	{
		p.SetState(124)
		p.Match(FGGParserEOF)
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
	p.RuleIndex = FGGParserRULE_decls
	return p
}

func (*DeclsContext) IsDeclsContext() {}

func NewDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclsContext {
	var p = new(DeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_decls

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterDecls(s)
	}
}

func (s *DeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitDecls(s)
	}
}

func (p *FGGParser) Decls() (localctx IDeclsContext) {
	localctx = NewDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FGGParserRULE_decls)

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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(128)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case FGGParserTYPE:
				{
					p.SetState(126)
					p.TypeDecl()
				}

			case FGGParserFUNC:
				{
					p.SetState(127)
					p.MethDecl()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(130)
				p.Match(FGGParserT__5)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.RuleIndex = FGGParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(FGGParserTYPE, 0)
}

func (s *TypeDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *TypeDeclContext) TypeFormals() ITypeFormalsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormalsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormalsContext)
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *FGGParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FGGParserRULE_typeDecl)

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
		p.SetState(136)
		p.Match(FGGParserTYPE)
	}
	{
		p.SetState(137)
		p.Match(FGGParserNAME)
	}
	{
		p.SetState(138)
		p.TypeFormals()
	}
	{
		p.SetState(139)
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
	p.RuleIndex = FGGParserRULE_methBody
	return p
}

func (*MethBodyContext) IsMethBodyContext() {}

func NewMethBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethBodyContext {
	var p = new(MethBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_methBody

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
	return s.GetToken(FGGParserRETURN, 0)
}

func (s *MethReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MethReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterMethReturn(s)
	}
}

func (s *MethReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserFMT, 0)
}

func (s *SprintfContext) SPRINTF() antlr.TerminalNode {
	return s.GetToken(FGGParserSPRINTF, 0)
}

func (s *SprintfContext) STRING() antlr.TerminalNode {
	return s.GetToken(FGGParserSTRING, 0)
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSprintf(s)
	}
}

func (s *SprintfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserSELECT, 0)
}

func (s *CaseSelectContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(FGGParserCASE)
}

func (s *CaseSelectContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(FGGParserCASE, i)
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
	return s.GetToken(FGGParserDEFAULT, 0)
}

func (s *CaseSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterCaseSelect(s)
	}
}

func (s *CaseSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitCaseSelect(s)
	}
}

func (p *FGGParser) MethBody() (localctx IMethBodyContext) {
	localctx = NewMethBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FGGParserRULE_methBody)
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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGGParserRETURN:
		localctx = NewMethReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(FGGParserRETURN)
		}
		{
			p.SetState(142)
			p.expr(0)
		}

	case FGGParserT__0, FGGParserT__1, FGGParserCLOSE, FGGParserGORTN, FGGParserMAKE, FGGParserNAME:
		localctx = NewSequenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Stmt()
		}
		{
			p.SetState(144)
			p.Match(FGGParserT__5)
		}
		{
			p.SetState(145)
			p.MethBody()
		}

	case FGGParserFMT:
		localctx = NewSprintfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Match(FGGParserFMT)
		}
		{
			p.SetState(148)
			p.Match(FGGParserT__11)
		}
		{
			p.SetState(149)
			p.Match(FGGParserSPRINTF)
		}
		{
			p.SetState(150)
			p.Match(FGGParserT__6)
		}
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FGGParserT__12 || _la == FGGParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FGGParserT__0)|(1<<FGGParserT__1)|(1<<FGGParserT__4)|(1<<FGGParserMAKE))) != 0) || _la == FGGParserNAME {
			p.SetState(154)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case FGGParserT__4:
				{
					p.SetState(152)
					p.Match(FGGParserT__4)
				}

			case FGGParserT__0, FGGParserT__1, FGGParserMAKE, FGGParserNAME:
				{
					p.SetState(153)
					p.expr(0)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(159)
			p.Match(FGGParserT__7)
		}

	case FGGParserSELECT:
		localctx = NewCaseSelectContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.Match(FGGParserSELECT)
		}
		{
			p.SetState(161)
			p.Match(FGGParserT__8)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FGGParserCASE {
			{
				p.SetState(162)
				p.Match(FGGParserCASE)
			}
			{
				p.SetState(163)
				p.CaseGuard()
			}
			{
				p.SetState(164)
				p.Match(FGGParserT__14)
			}
			{
				p.SetState(165)
				p.MethBody()
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGGParserDEFAULT {
			{
				p.SetState(172)
				p.Match(FGGParserDEFAULT)
			}
			{
				p.SetState(173)
				p.Match(FGGParserT__14)
			}
			{
				p.SetState(174)
				p.MethBody()
			}

		}
		{
			p.SetState(177)
			p.Match(FGGParserT__13)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = FGGParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_stmt

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
	return s.GetToken(FGGParserNAME, 0)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterChDispatch(s)
	}
}

func (s *ChDispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserGORTN, 0)
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterGoroutine(s)
	}
}

func (s *GoroutineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserCLOSE, 0)
}

func (s *ChCloseContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ChCloseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterChClose(s)
	}
}

func (s *ChCloseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitChClose(s)
	}
}

func (p *FGGParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FGGParserRULE_stmt)

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

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewChCloseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(FGGParserCLOSE)
		}
		{
			p.SetState(181)
			p.Match(FGGParserT__6)
		}
		{
			p.SetState(182)
			p.expr(0)
		}
		{
			p.SetState(183)
			p.Match(FGGParserT__7)
		}

	case 2:
		localctx = NewChDispatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.expr(0)
		}
		{
			p.SetState(186)
			p.Match(FGGParserT__1)
		}
		{
			p.SetState(187)
			p.expr(0)
		}

	case 3:
		localctx = NewGoroutineContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.Match(FGGParserGORTN)
		}
		{
			p.SetState(190)

			var _x = p.expr(0)

			localctx.(*GoroutineContext).recv = _x
		}
		{
			p.SetState(191)
			p.MethCall_()
		}

	case 4:
		localctx = NewAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(193)

			var _m = p.Match(FGGParserNAME)

			localctx.(*AssignmentContext).lvalue = _m
		}
		{
			p.SetState(194)
			p.Match(FGGParserT__15)
		}
		{
			p.SetState(195)
			p.expr(0)
		}

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
	p.RuleIndex = FGGParserRULE_methCall_
	return p
}

func (*MethCall_Context) IsMethCall_Context() {}

func NewMethCall_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethCall_Context {
	var p = new(MethCall_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_methCall_

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
	targs ITypsContext
	args  IExprsContext
}

func NewMethCallPrimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethCallPrimeContext {
	var p = new(MethCallPrimeContext)

	p.MethCall_Context = NewEmptyMethCall_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*MethCall_Context))

	return p
}

func (s *MethCallPrimeContext) GetTargs() ITypsContext { return s.targs }

func (s *MethCallPrimeContext) GetArgs() IExprsContext { return s.args }

func (s *MethCallPrimeContext) SetTargs(v ITypsContext) { s.targs = v }

func (s *MethCallPrimeContext) SetArgs(v IExprsContext) { s.args = v }

func (s *MethCallPrimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethCallPrimeContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *MethCallPrimeContext) Typs() ITypsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypsContext)
}

func (s *MethCallPrimeContext) Exprs() IExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *MethCallPrimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterMethCallPrime(s)
	}
}

func (s *MethCallPrimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitMethCallPrime(s)
	}
}

func (p *FGGParser) MethCall_() (localctx IMethCall_Context) {
	localctx = NewMethCall_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FGGParserRULE_methCall_)
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
		p.SetState(198)
		p.Match(FGGParserT__11)
	}
	{
		p.SetState(199)
		p.Match(FGGParserNAME)
	}
	{
		p.SetState(200)
		p.Match(FGGParserT__2)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGGParserT__0 || _la == FGGParserT__1 || _la == FGGParserNAME {
		{
			p.SetState(201)

			var _x = p.Typs()

			localctx.(*MethCallPrimeContext).targs = _x
		}

	}
	{
		p.SetState(204)
		p.Match(FGGParserT__3)
	}
	{
		p.SetState(205)
		p.Match(FGGParserT__6)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FGGParserT__0)|(1<<FGGParserT__1)|(1<<FGGParserMAKE))) != 0) || _la == FGGParserNAME {
		{
			p.SetState(206)

			var _x = p.Exprs()

			localctx.(*MethCallPrimeContext).args = _x
		}

	}
	{
		p.SetState(209)
		p.Match(FGGParserT__7)
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
	p.RuleIndex = FGGParserRULE_caseGuard
	return p
}

func (*CaseGuardContext) IsCaseGuardContext() {}

func NewCaseGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseGuardContext {
	var p = new(CaseGuardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_caseGuard

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
	return s.GetToken(FGGParserNAME, 0)
}

func (s *GuardReceiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterGuardReceive(s)
	}
}

func (s *GuardReceiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterGuardDispatch(s)
	}
}

func (s *GuardDispatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitGuardDispatch(s)
	}
}

func (p *FGGParser) CaseGuard() (localctx ICaseGuardContext) {
	localctx = NewCaseGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FGGParserRULE_caseGuard)

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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGuardDispatchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.expr(0)
		}
		{
			p.SetState(212)
			p.Match(FGGParserT__1)
		}
		{
			p.SetState(213)
			p.expr(0)
		}

	case 2:
		localctx = NewGuardReceiveContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)

			var _m = p.Match(FGGParserNAME)

			localctx.(*GuardReceiveContext).lvalue = _m
		}
		{
			p.SetState(216)
			p.Match(FGGParserT__15)
		}
		{
			p.SetState(217)
			p.Match(FGGParserT__1)
		}
		{
			p.SetState(218)
			p.expr(0)
		}

	}

	return localctx
}

// IMethDeclContext is an interface to support dynamic dispatch.
type IMethDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRecv returns the recv token.
	GetRecv() antlr.Token

	// GetTypn returns the typn token.
	GetTypn() antlr.Token

	// SetRecv sets the recv token.
	SetRecv(antlr.Token)

	// SetTypn sets the typn token.
	SetTypn(antlr.Token)

	// IsMethDeclContext differentiates from other interfaces.
	IsMethDeclContext()
}

type MethDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	recv   antlr.Token
	typn   antlr.Token
}

func NewEmptyMethDeclContext() *MethDeclContext {
	var p = new(MethDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_methDecl
	return p
}

func (*MethDeclContext) IsMethDeclContext() {}

func NewMethDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethDeclContext {
	var p = new(MethDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_methDecl

	return p
}

func (s *MethDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *MethDeclContext) GetRecv() antlr.Token { return s.recv }

func (s *MethDeclContext) GetTypn() antlr.Token { return s.typn }

func (s *MethDeclContext) SetRecv(v antlr.Token) { s.recv = v }

func (s *MethDeclContext) SetTypn(v antlr.Token) { s.typn = v }

func (s *MethDeclContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FGGParserFUNC, 0)
}

func (s *MethDeclContext) TypeFormals() ITypeFormalsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormalsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormalsContext)
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

func (s *MethDeclContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(FGGParserNAME)
}

func (s *MethDeclContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, i)
}

func (s *MethDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterMethDecl(s)
	}
}

func (s *MethDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitMethDecl(s)
	}
}

func (p *FGGParser) MethDecl() (localctx IMethDeclContext) {
	localctx = NewMethDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FGGParserRULE_methDecl)

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
		p.SetState(221)
		p.Match(FGGParserFUNC)
	}
	{
		p.SetState(222)
		p.Match(FGGParserT__6)
	}
	{
		p.SetState(223)

		var _m = p.Match(FGGParserNAME)

		localctx.(*MethDeclContext).recv = _m
	}
	{
		p.SetState(224)

		var _m = p.Match(FGGParserNAME)

		localctx.(*MethDeclContext).typn = _m
	}
	{
		p.SetState(225)
		p.TypeFormals()
	}
	{
		p.SetState(226)
		p.Match(FGGParserT__7)
	}
	{
		p.SetState(227)
		p.Sig()
	}
	{
		p.SetState(228)
		p.Match(FGGParserT__8)
	}
	{
		p.SetState(229)
		p.MethBody()
	}
	{
		p.SetState(230)
		p.Match(FGGParserT__13)
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
	p.RuleIndex = FGGParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_typeLit

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
	return s.GetToken(FGGParserSTRUCT, 0)
}

func (s *StructTypeLitContext) FieldDecls() IFieldDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclsContext)
}

func (s *StructTypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterStructTypeLit(s)
	}
}

func (s *StructTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserINTERFACE, 0)
}

func (s *InterfaceTypeLitContext) Specs() ISpecsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecsContext)
}

func (s *InterfaceTypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterInterfaceTypeLit(s)
	}
}

func (s *InterfaceTypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitInterfaceTypeLit(s)
	}
}

func (p *FGGParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FGGParserRULE_typeLit)
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

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FGGParserSTRUCT:
		localctx = NewStructTypeLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(FGGParserSTRUCT)
		}
		{
			p.SetState(233)
			p.Match(FGGParserT__8)
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGGParserNAME {
			{
				p.SetState(234)
				p.FieldDecls()
			}

		}
		{
			p.SetState(237)
			p.Match(FGGParserT__13)
		}

	case FGGParserINTERFACE:
		localctx = NewInterfaceTypeLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Match(FGGParserINTERFACE)
		}
		{
			p.SetState(239)
			p.Match(FGGParserT__8)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FGGParserT__0 || _la == FGGParserT__1 || _la == FGGParserNAME {
			{
				p.SetState(240)
				p.Specs()
			}

		}
		{
			p.SetState(243)
			p.Match(FGGParserT__13)
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
	p.RuleIndex = FGGParserRULE_fieldDecls
	return p
}

func (*FieldDeclsContext) IsFieldDeclsContext() {}

func NewFieldDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclsContext {
	var p = new(FieldDeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_fieldDecls

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterFieldDecls(s)
	}
}

func (s *FieldDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitFieldDecls(s)
	}
}

func (p *FGGParser) FieldDecls() (localctx IFieldDeclsContext) {
	localctx = NewFieldDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FGGParserRULE_fieldDecls)
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
		p.SetState(246)
		p.FieldDecl()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__5 {
		{
			p.SetState(247)
			p.Match(FGGParserT__5)
		}
		{
			p.SetState(248)
			p.FieldDecl()
		}

		p.SetState(253)
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

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) GetField() antlr.Token { return s.field }

func (s *FieldDeclContext) SetField(v antlr.Token) { s.field = v }

func (s *FieldDeclContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *FieldDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (p *FGGParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FGGParserRULE_fieldDecl)

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
		p.SetState(254)

		var _m = p.Match(FGGParserNAME)

		localctx.(*FieldDeclContext).field = _m
	}
	{
		p.SetState(255)
		p.Typ()
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
	p.RuleIndex = FGGParserRULE_specs
	return p
}

func (*SpecsContext) IsSpecsContext() {}

func NewSpecsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecsContext {
	var p = new(SpecsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_specs

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSpecs(s)
	}
}

func (s *SpecsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitSpecs(s)
	}
}

func (p *FGGParser) Specs() (localctx ISpecsContext) {
	localctx = NewSpecsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FGGParserRULE_specs)
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
		p.SetState(257)
		p.Spec()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__5 {
		{
			p.SetState(258)
			p.Match(FGGParserT__5)
		}
		{
			p.SetState(259)
			p.Spec()
		}

		p.SetState(264)
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
	p.RuleIndex = FGGParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_spec

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

func (s *InterfaceSpecContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *InterfaceSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterInterfaceSpec(s)
	}
}

func (s *InterfaceSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSigSpec(s)
	}
}

func (s *SigSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitSigSpec(s)
	}
}

func (p *FGGParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FGGParserRULE_spec)

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

	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSigSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Sig()
		}

	case 2:
		localctx = NewInterfaceSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Typ()
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

	// SetMeth sets the meth token.
	SetMeth(antlr.Token)

	// IsSigContext differentiates from other interfaces.
	IsSigContext()
}

type SigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	meth   antlr.Token
}

func NewEmptySigContext() *SigContext {
	var p = new(SigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_sig
	return p
}

func (*SigContext) IsSigContext() {}

func NewSigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SigContext {
	var p = new(SigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_sig

	return p
}

func (s *SigContext) GetParser() antlr.Parser { return s.parser }

func (s *SigContext) GetMeth() antlr.Token { return s.meth }

func (s *SigContext) SetMeth(v antlr.Token) { s.meth = v }

func (s *SigContext) TypeFormals() ITypeFormalsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeFormalsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeFormalsContext)
}

func (s *SigContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *SigContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSig(s)
	}
}

func (s *SigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitSig(s)
	}
}

func (p *FGGParser) Sig() (localctx ISigContext) {
	localctx = NewSigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FGGParserRULE_sig)
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
		p.SetState(269)

		var _m = p.Match(FGGParserNAME)

		localctx.(*SigContext).meth = _m
	}
	{
		p.SetState(270)
		p.TypeFormals()
	}
	{
		p.SetState(271)
		p.Match(FGGParserT__6)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FGGParserNAME {
		{
			p.SetState(272)
			p.Params()
		}

	}
	{
		p.SetState(275)
		p.Match(FGGParserT__7)
	}
	{
		p.SetState(276)
		p.Typ()
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
	p.RuleIndex = FGGParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_params

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *FGGParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FGGParserRULE_params)
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
		p.SetState(278)
		p.ParamDecl()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__4 {
		{
			p.SetState(279)
			p.Match(FGGParserT__4)
		}
		{
			p.SetState(280)
			p.ParamDecl()
		}

		p.SetState(285)
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

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vari   antlr.Token
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FGGParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) GetVari() antlr.Token { return s.vari }

func (s *ParamDeclContext) SetVari(v antlr.Token) { s.vari = v }

func (s *ParamDeclContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ParamDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(FGGParserNAME, 0)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterParamDecl(s)
	}
}

func (s *ParamDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitParamDecl(s)
	}
}

func (p *FGGParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FGGParserRULE_paramDecl)

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
		p.SetState(286)

		var _m = p.Match(FGGParserNAME)

		localctx.(*ParamDeclContext).vari = _m
	}
	{
		p.SetState(287)
		p.Typ()
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
	p.RuleIndex = FGGParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_expr

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserNAME, 0)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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

func (s *AssertContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *AssertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterAssert(s)
	}
}

func (s *AssertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserNAME, 0)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterChanRecv(s)
	}
}

func (s *ChanRecvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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

func (s *StructLitContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *StructLitContext) Exprs() IExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *StructLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterStructLit(s)
	}
}

func (s *StructLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
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
	return s.GetToken(FGGParserMAKE, 0)
}

func (s *MakeChanContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *MakeChanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterMakeChan(s)
	}
}

func (s *MakeChanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitMakeChan(s)
	}
}

func (p *FGGParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FGGParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, FGGParserRULE_expr, _p)
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
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(290)
			p.Match(FGGParserNAME)
		}

	case 2:
		localctx = NewStructLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(291)
			p.Typ()
		}
		{
			p.SetState(292)
			p.Match(FGGParserT__8)
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FGGParserT__0)|(1<<FGGParserT__1)|(1<<FGGParserMAKE))) != 0) || _la == FGGParserNAME {
			{
				p.SetState(293)
				p.Exprs()
			}

		}
		{
			p.SetState(296)
			p.Match(FGGParserT__13)
		}

	case 3:
		localctx = NewMakeChanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(298)
			p.Match(FGGParserMAKE)
		}
		{
			p.SetState(299)
			p.Match(FGGParserT__6)
		}
		{
			p.SetState(300)
			p.Typ()
		}
		{
			p.SetState(301)
			p.Match(FGGParserT__7)
		}

	case 4:
		localctx = NewChanRecvContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(303)
			p.Match(FGGParserT__1)
		}
		{
			p.SetState(304)
			p.expr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(318)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FGGParserRULE_expr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(308)
					p.Match(FGGParserT__11)
				}
				{
					p.SetState(309)
					p.Match(FGGParserNAME)
				}

			case 2:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CallContext).recv = _prevctx

				p.PushNewRecursionContext(localctx, _startState, FGGParserRULE_expr)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(311)
					p.MethCall_()
				}

			case 3:
				localctx = NewAssertContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FGGParserRULE_expr)
				p.SetState(312)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(313)
					p.Match(FGGParserT__11)
				}
				{
					p.SetState(314)
					p.Match(FGGParserT__6)
				}
				{
					p.SetState(315)
					p.Typ()
				}
				{
					p.SetState(316)
					p.Match(FGGParserT__7)
				}

			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
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
	p.RuleIndex = FGGParserRULE_exprs
	return p
}

func (*ExprsContext) IsExprsContext() {}

func NewExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprsContext {
	var p = new(ExprsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FGGParserRULE_exprs

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
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.EnterExprs(s)
	}
}

func (s *ExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FGGListener); ok {
		listenerT.ExitExprs(s)
	}
}

func (p *FGGParser) Exprs() (localctx IExprsContext) {
	localctx = NewExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FGGParserRULE_exprs)
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
		p.SetState(323)
		p.expr(0)
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FGGParserT__4 {
		{
			p.SetState(324)
			p.Match(FGGParserT__4)
		}
		{
			p.SetState(325)
			p.expr(0)
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *FGGParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FGGParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
