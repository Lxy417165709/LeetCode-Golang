package 未归类

import (
	"bytes"
	"strings"
)

func numUniqueEmails(emails []string) int {
	hasEmailExist := make(map[string]bool)
	for _, email := range emails {
		hasEmailExist[getFormattedEmail(email)] = true
	}
	return len(hasEmailExist)
}

func getFormattedEmail(email string) string {
	parts := strings.Split(email, "@")
	return getFormattedLocalName(parts[0]) + "@" + parts[1]
}

func getFormattedLocalName(sourceLocalName string) string {
	localName := bytes.Buffer{}
	for i := 0; i < len(sourceLocalName); i++ {
		if sourceLocalName[i] == '+' {
			break
		}
		if sourceLocalName[i] != '.' {
			localName.WriteByte(sourceLocalName[i])
		}
	}
	return localName.String()
}


/*
	题目链接: https://leetcode-cn.com/problems/unique-email-addresses/
*/
