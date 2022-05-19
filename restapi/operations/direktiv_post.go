package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/runtime/middleware"

	"showdownjs/models"
)

const (
	successKey = "success"
	resultKey  = "result"

	// http related
	statusKey  = "status"
	codeKey    = "code"
	headersKey = "headers"
)

var sm sync.Map

const (
	cmdErr = "io.direktiv.command.error"
	outErr = "io.direktiv.output.error"
	riErr  = "io.direktiv.ri.error"
)

type accParams struct {
	PostParams
	Commands []interface{}
}

type accParamsTemplate struct {
	PostBody
	Commands []interface{}
}

func PostDirektivHandle(params PostParams) middleware.Responder {
	fmt.Printf("params in: %+v", params)
	var resp interface{}

	var (
		err  error
		ret  interface{}
		cont bool
	)

	ri, err := apps.RequestinfoFromRequest(params.HTTPRequest)
	if err != nil {
		return generateError(riErr, err)
	}

	ctx, cancel := context.WithCancel(params.HTTPRequest.Context())
	sm.Store(*params.DirektivActionID, cancel)
	defer sm.Delete(params.DirektivActionID)

	var responses []interface{}

	var paramsCollector []interface{}
	accParams := accParams{
		params,
		nil,
	}

	ret, err = runCommand0(ctx, accParams, ri)
	responses = append(responses, ret)

	cont = convertTemplateToBool("<no value>", accParams, true)

	if err != nil && !cont {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	ret, err = runCommand1(ctx, accParams, ri)
	responses = append(responses, ret)

	cont = convertTemplateToBool("<no value>", accParams, true)

	if err != nil && !cont {
		errName := cmdErr
		return generateError(errName, err)
	}

	paramsCollector = append(paramsCollector, ret)
	accParams.Commands = paramsCollector

	fmt.Printf("object going in output template: %+v\n", responses)

	s, err := templateString(`{
  "html": "{{ index (index . 1) "result" }}"
}
`, responses)
	if err != nil {
		return generateError(outErr, err)
	}
	fmt.Printf("object from output template: %+v\n", s)

	responseBytes := []byte(s)

	err = json.Unmarshal(responseBytes, &resp)
	if err != nil {
		fmt.Printf("error parsing output template: %+v\n", err)
		return generateError(outErr, err)
	}

	return NewPostOK().WithPayload(resp)
}

// exec
func runCommand0(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ir := make(map[string]interface{})
	ir[successKey] = false

	ri.Logger().Infof("executing command")

	at := accParamsTemplate{
		params.Body,
		params.Commands,
	}
	fmt.Printf("object going in command template: %+v\n", at)

	cmd, err := templateString(`{{- $v := "out.md" }}
{{- if and (not (empty .Scope)) (not (empty .Name)) }}
  {{- $v = (list "out" .Scope .Name | join "/") }}
{{- end }}
showdown makehtml --input input.md --output {{ $v }} \
  -c strikethrough \
  -c omitExtraWLInCodeBlocks \
  -c simplifiedAutoLink \
  -c literalMidWordUnderscores \
  -c tables \
  -c tablesHeaderId \
  -c ghCodeBlocks \
  -c tasklists \
  -c disableForced4SpacesIndentedSublists \
  -c simpleLineBreaks \
  -c requireSpaceBeforeHeadingText \
  -c ghCompatibleHeaderId \
  -c ghMentions \
  -c backslashEscapesHTMLTags \
  -c emoji \
  -c splitAdjacentBlockquotes `, at)
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	silent := convertTemplateToBool("<no value>", at, false)
	print := convertTemplateToBool("<no value>", at, true)
	output := ""

	envs := []string{}

	return runCmd(ctx, cmd, envs, output, silent, print, ri)

}

// end commands

// exec
func runCommand1(ctx context.Context,
	params accParams, ri *apps.RequestInfo) (map[string]interface{}, error) {

	ir := make(map[string]interface{})
	ir[successKey] = false

	ri.Logger().Infof("executing command")

	at := accParamsTemplate{
		params.Body,
		params.Commands,
	}
	fmt.Printf("object going in command template: %+v\n", at)

	cmd, err := templateString(`{{- $v := "out.md" }}
{{- if and (not (empty .Scope)) (not (empty .Name)) }}
  {{- $v = (list "out" .Scope .Name | join "/") }}
{{- end }}
base64 -w0 {{ $v }}`, at)
	if err != nil {
		ir[resultKey] = err.Error()
		return ir, err
	}
	cmd = strings.Replace(cmd, "\n", "", -1)

	silent := convertTemplateToBool("<no value>", at, false)
	print := convertTemplateToBool("<no value>", at, true)
	output := ""

	envs := []string{}

	return runCmd(ctx, cmd, envs, output, silent, print, ri)

}

// end commands

func generateError(code string, err error) *PostDefault {

	d := NewPostDefault(0).WithDirektivErrorCode(code).
		WithDirektivErrorMessage(err.Error())

	errString := err.Error()

	errResp := models.Error{
		ErrorCode:    &code,
		ErrorMessage: &errString,
	}

	d.SetPayload(&errResp)

	return d
}

func HandleShutdown() {
	// nothing for generated functions
}
