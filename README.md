
# showdownjs 1.0.0

Convert github markdown files to HTML.

---
- #### Image: direktiv/showdownjs 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/showdownjs/issues
- #### URL: https://github.com/direktiv-apps/showdownjs
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About showdownjs

Showdownjs is a library for coverting markdown into HTML. It can not only convert basic markdown but understands Github's 
markdown extensions. This function takes a file 'input.md' and returns it as HTML as base64 in the response. 
Additionally it can store the result in Direktiv variables. 

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: showdownjs
    image: direktiv/showdownjs
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: showdownjs
     type: action
      action:
        files:
        - key: readme.md
          scope: workflow
          as: input.md
        function: get
        input:
          # saves the result in workflow variable 'final.html'
          scope: workflow
          name: final.html
   ```

### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Returns base64 encoded HTML of the provided markdown file.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| html | string| `string` |  | | Base64 encoded HTML |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | Scope of the saved html output within the scope. | `myhtml` |
| scope | string| `string` |  | | Scope of the saved html output. | `instance` |

 
