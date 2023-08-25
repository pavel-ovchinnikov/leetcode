package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	qp := make([][]bool, len(s1)+1)
	for i := range qp {
		qp[i] = make([]bool, len(s2)+1)
	}

	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				qp[i][j] = true
			} else if i == 0 {
				qp[i][j] = (qp[i][j-1] && s2[j-1] == s3[i+j-1])
			} else if j == 0 {
				qp[i][j] = (qp[i-1][j] && s1[i-1] == s3[i+j-1])
			} else {
				qp[i][j] = (qp[i-1][j] &&
					s1[i-1] == s3[i+j-1]) ||
					(qp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return qp[len(s1)][len(s2)]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
