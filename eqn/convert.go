package eqn

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ConvertUri(filepath string) string {

	 res, err := http.Get(filepath)
	if err != nil {
		fmt.Print(err)
		return "ERROR: url="+filepath+";err="+err.Error()
	}
	if !strings.HasPrefix(res.Status,"200") {
		defer res.Body.Close()
		return "ERROR: http "+res.Status+" error url="+filepath
	}
	bodyRes, err1 := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err1 != nil || len(bodyRes) <1 {
		fmt.Print(err1)
		return "ERROR: filepath="+filepath+";readerr="+err1.Error()
	}

	reader := bytes.NewReader(bodyRes)
	mtef, err2 := Open(reader)
	if err2 != nil {
		fmt.Print(err2)
		return "ERROR: filepath="+filepath+";openerr="+err2.Error()
	}
	latex := mtef.Translate()
	return latex
}
func Convert(filepath string) string {
	buffer, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return "ERROR: filepath="+filepath+";openerr="+err.Error()

	}

	reader := bytes.NewReader(buffer)
	mtef, err := Open(reader)
	if err != nil {

		fmt.Println(err)
		return "ERROR: filepath="+filepath+";openerr="+err.Error()
	}

	latex := mtef.Translate()
	//加密的无法解开； 开通“校园号”即可编辑公式
	if strings.Contains(latex,"校园号") || latex == "$$  $$"{
		return ""
	}
	return latex
}
