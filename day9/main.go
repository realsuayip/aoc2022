package day9

import (
	"fmt"
	im "image"
	"math"
	"strings"
)

var input = "R 1\nU 1\nD 2\nL 2\nD 1\nU 1\nR 2\nD 2\nR 2\nU 1\nL 1\nU 2\nR 1\nL 1\nR 1\nL 1\nR 1\nU 1\nL 1\nU 1\nR 2\nU 2\nL 1\nR 1\nL 1\nU 2\nR 2\nD 1\nR 2\nU 2\nL 1\nU 2\nD 2\nL 1\nU 2\nR 2\nL 2\nU 1\nD 1\nL 1\nR 1\nL 1\nU 1\nL 1\nU 1\nR 2\nL 1\nR 1\nD 2\nL 2\nR 2\nU 2\nD 1\nL 1\nU 2\nR 2\nD 1\nU 1\nD 2\nU 1\nL 1\nD 2\nU 1\nR 1\nL 1\nU 2\nR 2\nU 1\nD 2\nR 1\nL 2\nD 2\nL 1\nD 2\nL 1\nR 1\nU 1\nR 2\nL 2\nU 2\nL 2\nR 1\nL 1\nD 1\nL 1\nR 1\nL 1\nR 1\nU 2\nD 1\nL 1\nR 1\nL 2\nU 2\nD 1\nU 1\nR 1\nU 1\nL 2\nU 2\nR 2\nU 2\nL 1\nU 1\nR 1\nU 2\nR 2\nU 1\nL 2\nU 2\nL 1\nD 1\nU 1\nR 3\nD 3\nL 2\nD 2\nR 2\nU 3\nL 3\nD 2\nL 1\nD 3\nL 2\nR 3\nD 1\nU 1\nD 3\nR 2\nD 3\nR 3\nD 1\nU 3\nD 1\nR 1\nU 1\nD 1\nR 3\nD 1\nU 3\nL 2\nD 2\nU 2\nL 3\nD 1\nR 3\nU 2\nD 2\nU 2\nR 3\nL 2\nR 2\nU 2\nL 3\nR 3\nU 1\nL 1\nD 1\nR 2\nL 3\nD 1\nU 3\nR 1\nD 3\nL 1\nU 1\nL 3\nU 1\nR 2\nL 3\nD 1\nR 2\nL 1\nU 3\nD 2\nL 1\nU 1\nL 3\nD 1\nU 1\nR 2\nU 1\nR 1\nU 2\nD 2\nU 1\nD 1\nL 1\nD 2\nU 2\nL 3\nD 3\nU 1\nR 2\nU 3\nD 1\nU 1\nR 2\nL 1\nU 2\nD 3\nR 1\nU 2\nR 1\nL 1\nD 1\nU 3\nD 1\nR 1\nL 2\nD 2\nR 2\nL 1\nD 3\nR 3\nL 1\nD 1\nR 2\nU 1\nR 3\nU 1\nL 2\nR 2\nL 2\nD 3\nL 4\nD 3\nL 4\nD 3\nR 4\nU 4\nL 4\nU 4\nR 2\nL 4\nD 2\nU 1\nR 2\nD 2\nR 3\nL 3\nR 4\nU 2\nD 3\nL 4\nR 2\nL 1\nR 3\nL 3\nU 4\nD 2\nL 2\nR 1\nU 1\nR 1\nU 3\nR 2\nD 2\nU 3\nR 2\nU 4\nR 4\nU 3\nD 4\nL 2\nR 4\nD 1\nU 3\nR 1\nD 4\nL 2\nD 2\nR 3\nD 3\nU 1\nL 2\nU 2\nL 4\nU 3\nR 3\nL 2\nR 3\nU 1\nL 1\nD 2\nU 1\nD 1\nU 1\nD 2\nU 3\nD 3\nR 4\nL 4\nD 1\nU 3\nR 4\nU 4\nD 2\nR 1\nU 2\nD 4\nU 1\nR 1\nL 4\nR 3\nU 1\nL 2\nR 3\nL 2\nU 3\nR 2\nD 2\nR 1\nD 2\nU 4\nL 4\nU 2\nR 2\nU 4\nL 3\nR 1\nD 4\nR 2\nD 3\nR 1\nL 3\nD 3\nL 4\nU 2\nR 1\nL 3\nD 4\nL 2\nU 3\nL 1\nU 5\nL 1\nU 3\nD 2\nU 2\nD 3\nR 3\nL 5\nU 1\nD 2\nL 1\nR 5\nD 4\nR 1\nD 2\nR 5\nU 5\nD 1\nR 3\nD 5\nR 5\nU 2\nL 1\nD 1\nR 2\nD 2\nU 5\nD 2\nR 3\nL 4\nD 5\nU 3\nL 4\nR 4\nD 5\nL 2\nR 5\nD 1\nR 2\nL 5\nD 2\nU 3\nD 3\nL 3\nU 3\nR 1\nD 4\nL 1\nU 2\nL 3\nU 2\nR 3\nL 1\nU 4\nL 5\nR 4\nU 4\nR 2\nL 4\nU 2\nR 4\nD 2\nU 2\nD 3\nL 3\nR 1\nU 1\nL 1\nR 5\nU 1\nD 4\nL 4\nR 5\nD 4\nL 3\nD 1\nL 4\nU 3\nR 5\nL 2\nR 4\nL 1\nR 2\nU 3\nD 1\nL 2\nU 2\nR 2\nU 1\nR 2\nU 5\nL 4\nD 1\nU 5\nR 1\nD 1\nU 2\nD 5\nU 5\nL 3\nU 3\nL 1\nD 5\nU 5\nD 1\nL 5\nD 4\nU 4\nL 3\nU 3\nL 6\nD 4\nR 5\nU 4\nR 3\nD 3\nU 3\nD 5\nR 4\nD 5\nR 6\nU 3\nL 3\nR 3\nD 3\nR 5\nL 2\nD 5\nL 3\nU 2\nD 3\nU 1\nD 1\nU 6\nD 4\nR 6\nD 4\nL 4\nD 6\nR 2\nL 3\nR 4\nU 5\nR 5\nD 4\nU 5\nL 3\nR 1\nU 4\nR 1\nD 5\nU 6\nD 6\nU 6\nD 4\nL 6\nR 5\nL 1\nU 6\nL 6\nD 3\nR 3\nD 2\nU 3\nL 5\nD 6\nU 3\nL 3\nD 3\nU 3\nD 4\nR 3\nL 5\nU 5\nD 5\nR 1\nU 4\nD 2\nR 6\nD 2\nL 3\nR 2\nD 1\nU 4\nR 6\nD 6\nU 3\nL 3\nD 1\nR 2\nD 4\nR 5\nL 6\nD 2\nU 6\nL 1\nU 5\nR 4\nU 6\nD 1\nL 4\nU 4\nD 1\nR 6\nL 4\nU 2\nR 3\nU 3\nL 2\nD 6\nR 4\nD 4\nL 6\nR 1\nD 2\nL 4\nR 2\nU 6\nL 1\nR 3\nD 3\nU 4\nL 1\nU 6\nD 1\nL 5\nR 6\nL 1\nD 5\nU 5\nR 4\nU 6\nL 5\nD 4\nR 1\nL 6\nU 7\nL 1\nD 3\nL 5\nR 4\nL 2\nD 4\nU 4\nL 4\nD 2\nR 4\nL 1\nU 1\nR 3\nU 2\nL 4\nR 2\nL 6\nR 1\nD 5\nL 6\nU 6\nD 1\nR 2\nL 2\nR 2\nD 6\nR 7\nD 5\nR 5\nD 1\nR 2\nU 3\nR 7\nU 6\nR 1\nU 2\nL 1\nR 3\nL 2\nR 2\nU 5\nR 1\nL 7\nU 6\nD 1\nL 2\nR 5\nD 5\nR 2\nU 7\nR 5\nU 3\nR 3\nU 1\nR 4\nD 2\nU 7\nD 7\nR 1\nU 4\nR 7\nD 4\nR 7\nL 5\nD 3\nR 2\nL 6\nU 1\nL 5\nR 7\nU 5\nD 5\nL 1\nD 5\nL 3\nU 5\nD 2\nR 7\nL 7\nR 7\nD 4\nR 1\nU 5\nD 3\nU 4\nL 5\nR 6\nU 6\nL 6\nU 3\nD 6\nU 7\nL 1\nU 1\nR 2\nD 4\nU 2\nD 2\nU 4\nR 2\nL 3\nR 5\nU 6\nL 8\nR 7\nD 1\nL 1\nD 3\nL 4\nR 3\nU 7\nD 1\nR 4\nD 5\nR 7\nD 5\nU 2\nD 6\nU 2\nR 3\nD 6\nL 3\nU 5\nL 7\nD 6\nU 5\nD 6\nL 8\nU 2\nR 8\nD 6\nU 8\nR 1\nL 3\nR 7\nL 5\nR 8\nU 1\nD 6\nU 8\nL 4\nR 6\nD 2\nR 2\nD 8\nU 7\nD 7\nL 2\nU 7\nD 3\nL 2\nU 7\nR 3\nU 2\nD 7\nR 1\nL 2\nD 7\nL 2\nR 5\nU 6\nL 6\nR 1\nL 3\nR 5\nD 3\nR 3\nD 6\nU 1\nD 4\nU 4\nL 2\nR 7\nD 5\nL 3\nR 3\nD 2\nU 3\nD 4\nL 3\nU 8\nL 4\nD 1\nL 2\nU 3\nL 8\nR 5\nL 8\nU 1\nD 4\nU 3\nL 6\nR 7\nL 5\nD 5\nL 2\nR 4\nL 6\nD 5\nL 7\nU 1\nL 5\nR 8\nD 6\nR 8\nL 7\nU 5\nD 2\nL 8\nU 2\nR 6\nL 4\nD 6\nL 6\nD 8\nR 1\nL 5\nU 9\nR 6\nL 2\nU 2\nD 6\nL 9\nR 6\nD 5\nR 1\nU 4\nL 2\nR 7\nU 7\nD 3\nR 4\nD 9\nR 2\nL 7\nU 1\nL 6\nR 2\nD 3\nR 9\nD 2\nU 5\nL 3\nR 4\nD 7\nL 8\nU 2\nD 2\nR 1\nL 7\nR 5\nU 1\nL 6\nU 3\nL 1\nU 2\nD 1\nU 5\nL 2\nU 9\nL 5\nD 1\nL 8\nR 2\nL 5\nR 2\nU 1\nD 7\nL 2\nR 7\nU 6\nR 8\nL 7\nR 8\nU 3\nD 6\nU 4\nR 2\nU 2\nL 4\nD 5\nL 4\nR 1\nU 4\nR 7\nD 8\nL 6\nU 6\nD 3\nR 8\nD 4\nR 4\nD 7\nR 9\nD 7\nR 8\nU 4\nD 1\nL 7\nU 9\nR 3\nL 9\nU 2\nD 6\nU 5\nR 5\nU 3\nR 4\nD 7\nL 2\nU 1\nL 1\nU 7\nL 5\nU 5\nR 4\nD 9\nR 9\nD 5\nL 10\nR 6\nD 7\nL 9\nD 1\nL 10\nD 9\nL 7\nR 2\nD 5\nL 2\nU 6\nL 4\nU 2\nL 7\nD 7\nL 2\nR 1\nL 4\nD 2\nR 9\nD 3\nU 4\nD 6\nR 2\nD 7\nL 7\nU 10\nD 6\nL 4\nR 3\nU 9\nR 3\nL 10\nU 8\nD 1\nU 5\nR 2\nU 5\nL 3\nU 7\nR 1\nU 8\nD 4\nL 7\nD 10\nL 8\nU 9\nR 6\nL 8\nD 4\nR 5\nD 9\nU 2\nR 4\nD 4\nU 3\nR 7\nD 1\nL 1\nU 6\nL 2\nU 1\nL 5\nD 1\nL 2\nR 7\nD 8\nR 4\nU 2\nR 10\nU 3\nL 3\nD 1\nL 1\nU 10\nL 10\nU 5\nD 2\nR 6\nL 8\nR 2\nL 5\nU 2\nD 8\nL 1\nD 4\nU 9\nR 6\nD 2\nL 3\nR 4\nD 4\nU 4\nL 2\nR 9\nU 8\nL 10\nD 1\nR 6\nD 2\nR 6\nU 7\nL 7\nU 6\nD 6\nR 3\nD 1\nU 5\nR 3\nD 1\nU 5\nL 9\nR 8\nU 11\nL 11\nU 6\nD 7\nL 7\nR 8\nU 10\nD 10\nL 11\nU 1\nL 5\nR 3\nD 8\nL 8\nD 5\nL 6\nU 7\nD 8\nU 9\nD 11\nU 7\nD 9\nU 11\nL 9\nR 9\nU 1\nL 11\nU 6\nR 1\nU 9\nD 9\nU 1\nR 6\nL 9\nR 11\nD 5\nL 10\nR 2\nD 10\nL 7\nR 2\nL 5\nD 5\nR 9\nL 11\nU 4\nD 2\nU 10\nR 5\nU 10\nL 5\nU 2\nD 6\nU 5\nL 6\nU 8\nR 9\nU 9\nD 8\nR 6\nD 1\nR 2\nL 7\nU 11\nD 8\nL 4\nU 2\nR 3\nD 9\nU 5\nR 7\nL 4\nD 3\nU 11\nL 6\nU 6\nD 5\nU 10\nR 11\nU 5\nL 2\nU 7\nR 11\nD 2\nR 2\nL 10\nD 8\nR 10\nL 10\nU 4\nL 5\nD 7\nL 8\nR 8\nD 4\nL 6\nU 7\nR 9\nU 4\nR 9\nD 4\nR 3\nU 9\nR 6\nD 8\nU 6\nD 2\nU 3\nR 2\nL 9\nU 3\nD 11\nU 10\nD 1\nU 2\nD 7\nR 4\nU 5\nD 12\nL 4\nR 12\nL 3\nR 4\nL 7\nR 2\nD 9\nR 3\nU 2\nL 10\nD 4\nR 9\nL 4\nD 11\nL 3\nU 8\nL 11\nD 12\nU 8\nL 11\nR 4\nD 8\nR 10\nU 1\nD 2\nU 12\nD 6\nR 6\nD 2\nU 5\nR 2\nL 7\nU 4\nR 2\nU 10\nL 8\nD 3\nR 8\nL 4\nR 5\nD 7\nL 12\nU 12\nD 3\nL 10\nR 4\nL 12\nU 3\nL 6\nD 6\nU 5\nL 1\nU 5\nL 12\nR 9\nU 11\nL 11\nR 7\nL 5\nR 8\nU 12\nR 10\nU 7\nL 2\nR 5\nU 2\nD 3\nL 7\nR 5\nU 9\nR 9\nU 10\nD 10\nU 12\nL 2\nR 2\nL 1\nU 6\nR 8\nU 9\nR 1\nL 9\nD 2\nL 1\nR 11\nD 1\nU 4\nL 10\nR 9\nD 12\nR 4\nD 5\nU 6\nR 7\nD 8\nR 5\nD 11\nL 3\nD 11\nU 4\nL 6\nU 4\nL 10\nU 2\nL 4\nU 5\nR 2\nU 4\nR 2\nD 6\nR 13\nU 4\nD 2\nL 12\nD 6\nR 6\nU 5\nD 2\nR 12\nD 2\nR 1\nL 6\nD 6\nU 11\nL 11\nD 3\nL 10\nU 2\nR 13\nD 6\nU 3\nL 3\nR 1\nU 7\nR 11\nD 2\nL 6\nD 5\nL 11\nR 9\nU 13\nR 6\nU 4\nD 6\nU 5\nR 4\nL 3\nR 10\nL 10\nD 5\nU 10\nL 13\nU 3\nD 6\nU 1\nL 11\nR 13\nL 1\nD 7\nR 3\nU 7\nD 8\nR 12\nD 10\nR 10\nL 11\nR 10\nD 1\nL 13\nR 1\nL 3\nR 11\nU 6\nL 1\nD 11\nU 11\nD 7\nL 5\nR 2\nU 7\nL 13\nR 13\nU 12\nR 2\nU 5\nR 3\nU 12\nL 5\nD 13\nL 11\nR 7\nL 2\nD 5\nR 6\nD 5\nL 6\nD 5\nU 13\nL 9\nD 11\nR 9\nD 11\nL 9\nR 11\nU 5\nR 10\nL 7\nR 10\nD 5\nL 9\nU 12\nL 10\nU 9\nL 14\nR 3\nL 4\nD 3\nL 4\nR 4\nD 5\nR 3\nL 10\nR 12\nD 4\nR 13\nU 7\nL 11\nD 2\nL 2\nU 11\nR 4\nD 14\nU 9\nR 9\nL 5\nD 2\nU 6\nD 2\nL 1\nR 3\nD 7\nR 4\nU 9\nD 8\nU 5\nL 12\nU 3\nL 12\nR 14\nU 5\nL 10\nR 1\nD 6\nU 10\nL 2\nU 8\nR 5\nL 9\nR 8\nL 6\nD 8\nU 11\nR 6\nL 7\nD 10\nR 6\nL 9\nU 6\nD 6\nL 3\nR 13\nD 10\nR 11\nL 10\nR 6\nL 9\nD 6\nR 14\nL 2\nU 7\nD 4\nL 11\nU 12\nL 11\nD 7\nL 12\nU 8\nD 9\nU 7\nD 4\nL 9\nR 12\nU 10\nD 12\nU 8\nR 9\nU 4\nD 2\nR 9\nD 1\nL 14\nU 4\nL 13\nR 8\nD 7\nR 6\nL 6\nD 11\nR 11\nD 4\nL 12\nR 4\nD 6\nL 1\nD 10\nR 3\nU 6\nL 6\nR 6\nL 4\nR 1\nU 3\nL 10\nR 3\nU 11\nR 5\nU 12\nD 9\nU 12\nL 5\nR 5\nL 3\nD 11\nU 1\nR 7\nU 4\nD 11\nR 14\nL 1\nU 6\nD 9\nU 4\nR 8\nD 8\nU 13\nD 1\nU 5\nR 9\nU 3\nL 2\nR 11\nU 8\nD 14\nU 4\nR 14\nU 3\nD 11\nU 11\nR 8\nL 4\nR 11\nL 13\nR 13\nD 9\nU 10\nD 12\nU 15\nL 3\nR 9\nL 6\nR 10\nU 13\nL 6\nD 2\nL 10\nD 8\nU 4\nR 9\nD 12\nL 5\nR 13\nU 8\nD 4\nR 10\nD 2\nU 15\nL 13\nU 9\nL 1\nR 3\nU 13\nR 3\nD 7\nL 6\nD 15\nR 12\nU 7\nD 10\nL 5\nU 14\nL 1\nU 3\nD 2\nR 14\nU 9\nR 11\nU 15\nL 9\nU 8\nL 1\nD 13\nR 1\nU 12\nL 11\nU 12\nR 9\nU 10\nL 9\nR 6\nL 6\nU 12\nR 14\nD 10\nL 12\nR 13\nD 8\nL 12\nD 12\nU 3\nL 11\nR 6\nU 15\nD 7\nR 10\nL 1\nU 12\nL 15\nU 12\nL 5\nD 8\nR 8\nL 4\nD 13\nR 3\nD 7\nR 13\nL 9\nR 7\nU 7\nL 16\nD 4\nU 3\nD 1\nU 7\nL 5\nR 1\nD 8\nR 2\nL 8\nR 12\nD 12\nU 14\nR 9\nD 6\nR 4\nU 9\nR 10\nL 10\nD 3\nL 13\nU 1\nD 7\nU 8\nR 14\nL 8\nR 14\nL 5\nU 16\nD 8\nR 6\nU 6\nR 14\nL 16\nR 14\nD 5\nR 14\nD 10\nU 16\nR 1\nD 10\nU 12\nD 12\nR 14\nD 16\nU 2\nL 3\nU 2\nR 9\nU 15\nR 7\nL 16\nU 2\nR 1\nD 11\nU 11\nL 15\nD 8\nL 12\nR 3\nU 1\nR 10\nD 9\nL 15\nU 6\nD 13\nL 16\nR 6\nD 7\nL 8\nD 3\nR 13\nL 2\nD 6\nL 13\nD 16\nU 6\nD 11\nU 15\nD 10\nU 15\nL 11\nR 13\nL 9\nU 2\nD 15\nU 1\nL 8\nR 3\nU 15\nD 2\nR 4\nL 10\nU 15\nR 17\nD 13\nL 15\nU 8\nL 12\nU 13\nL 15\nR 10\nD 11\nL 14\nU 16\nR 4\nD 1\nR 4\nD 17\nU 8\nD 17\nR 9\nU 17\nD 9\nR 13\nD 15\nR 4\nD 17\nR 3\nD 2\nU 15\nL 8\nD 17\nU 1\nL 4\nU 5\nD 17\nU 3\nR 10\nD 10\nR 9\nL 17\nR 4\nD 6\nU 1\nD 12\nU 11\nL 13\nR 3\nU 5\nR 5\nD 7\nR 6\nU 8\nD 5\nR 8\nL 10\nR 4\nU 17\nR 7\nD 16\nU 17\nL 11\nU 13\nR 6\nD 8\nR 15\nL 10\nR 9\nL 6\nR 13\nD 17\nU 7\nR 11\nD 9\nR 13\nL 15\nR 5\nD 3\nL 4\nU 14\nL 17\nU 17\nD 1\nU 3\nL 9\nD 11\nU 14\nL 2\nD 14\nU 9\nR 2\nD 15\nR 16\nD 17\nL 16\nU 8\nD 1\nR 2\nL 6\nU 14\nR 7\nL 15\nD 14\nL 3\nR 5\nU 5\nR 15\nD 15\nU 7\nD 11\nU 13\nR 16\nL 17\nU 11\nR 13\nU 10\nL 7\nU 8\nL 3\nD 3\nL 1\nD 2\nL 2\nD 4\nL 15\nR 14\nU 13\nR 15\nU 15\nL 2\nR 10\nU 4\nD 11\nU 12\nL 1\nR 7\nL 4\nR 12\nD 1\nL 10\nR 8\nL 4\nR 13\nU 6\nR 10\nD 17\nR 14\nD 4\nR 11\nD 13\nL 16\nU 8\nD 1\nR 16\nU 11\nD 15\nU 1\nD 14\nL 9\nD 5\nR 1\nU 2\nL 9\nD 15\nR 18\nD 6\nR 11\nU 7\nR 5\nU 1\nL 13\nD 7\nU 15\nD 9\nR 11\nD 18\nU 11\nD 13\nL 7\nR 18\nU 13\nL 3\nU 18\nR 15\nU 9\nD 6\nU 7\nR 3\nL 6\nD 1\nU 5\nL 16\nR 7\nL 14\nD 6\nL 4\nR 16\nU 3\nL 6\nU 8\nR 2\nU 13\nR 18\nL 10\nD 15\nU 3\nR 12\nU 13\nL 14\nU 15\nR 11\nL 17\nU 2\nR 18\nL 16\nR 6\nD 17\nL 11\nR 16\nD 4\nU 6\nR 18\nD 15\nU 18\nR 5\nD 8\nL 10\nD 7\nL 19\nR 19\nL 10\nD 1\nR 2\nU 5\nD 6\nU 2\nL 11\nR 14\nD 13\nL 15\nU 11\nR 7\nL 5\nU 15\nL 4\nR 3\nL 13\nU 17\nL 3\nU 15\nR 11\nL 17\nU 8\nR 8\nD 10\nL 9\nD 17\nL 14\nR 2\nL 19\nR 3\nD 3\nU 18\nD 15\nL 3\nR 18\nU 5\nD 17\nL 15\nR 7\nD 10\nU 11\nR 8\nU 14\nD 16\nR 10\nU 14\nR 11\nL 7\nR 1\nU 16\nL 6\nD 13\nL 9\nU 7\nR 19\nL 1\nD 11\nL 18\nD 19\nU 15\nR 18\nU 8\nR 4\nD 18\nR 12\nD 18\nR 8\nU 2\nL 2\nU 2\nL 12\nU 18\nR 12\nL 17\nR 2\nU 12\nD 19\nU 12\nL 8\nU 2\nL 12\nU 19\nR 11\nD 9\nR 10\nU 17\nL 19\nD 15\nL 3\nR 13\nU 2\nL 13\nR 6\nD 4\nU 18\nR 11\nU 7\nD 2\nR 4\nL 13\nU 11\nD 10\nU 17\nD 3\nU 10"

func Part1() int {
	visited := make(map[im.Point]struct{}, 0)
	head, tail := im.Pt(0, 0), im.Pt(0, 0)
	for _, motion := range strings.Split(input, "\n") {
		var direction string
		var amount int
		var p im.Point
		_, _ = fmt.Sscanf(motion, "%s %d", &direction, &amount)

		for i := 1; i < amount+1; i++ {
			switch direction {
			case "R":
				{
					p = im.Pt(1, 0)
				}
			case "L":
				{
					p = im.Pt(-1, 0)
				}
			case "U":
				{
					p = im.Pt(0, 1)
				}
			case "D":
				{
					p = im.Pt(0, -1)
				}
			}
			head = head.Add(p)
			diff := head.Sub(tail)
			if math.Abs(float64(diff.Y)) > 1 || math.Abs(float64(diff.X)) > 1 {
				tail = tail.Add(p)
				if (direction == "U" || direction == "D") && tail.X != head.X {
					tail.X = head.X
				}
				if (direction == "L" || direction == "R") && tail.Y != head.Y {
					tail.Y = head.Y
				}
			}
			visited[tail] = struct{}{}
		}
	}
	return len(visited)
}

func Part2() {
	// https://www.reddit.com/r/adventofcode/comments/zgwhh1/2022_day_9_part_2_cant_believe_i_didnt_see_that/
}
