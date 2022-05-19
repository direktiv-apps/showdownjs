
# aws-cli 1.0.0

Execute AWS CLI commands from Direktiv.

---
- #### Categories: cloud, aws
- #### Image: direktiv/aws-cli 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/aws-cli/issues
- #### URL: https://github.com/direktiv-apps/aws-cli
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About aws-cli

This service excutes AWS CLI commands. All commands are getting executed in the specified region and return their results as JSON.

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: aws-cli
    image: direktiv/aws-cli
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: req
     type: action
     action:
       function: aws-cli
       secrets: ["awsacess", "awssecret"]
       input:
        access-key: jq(.secrets.awsacess)
        secret-key: jq(.secrets.awssecret)
        region: eu-central-1
        commands:
        - ec2 describe-instances
        - ecr get-login-password
   ```

### Request

The request body includes a list of AWS CLI commands.

#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  AWS CLI response.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "VTQ3U....c2ZaN0FJaldjVnkra2tKV==",
    "success": true
  },
  {
    "result": "exit status 254",
    "success": false
  }
]
```

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
| output | [][PostOKBodyOutputItems](#post-o-k-body-output-items)| `[]*PostOKBodyOutputItems` |  | |  |  |


#### <span id="post-o-k-body-output-items"></span> postOKBodyOutputItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access-key | string| `string` | ✓ | | AWS access key. | `ABCABCABCDABCABCABCD` |
| commands | []string| `[]string` |  | | Array of AWS cli commands. Does NOT include 'aws'. | `["ecr get-login-password","ec2 describe-instances"]` |
| continue | boolean| `bool` |  | | If set to true all commands are getting executed and errors ignored. | `true` |
| region | string| `string` |  | `"us-east-1"`| Region the commands should be executed in. | `eu-central-1` |
| secret-key | string| `string` | ✓ | | AWS secret key. | `Abcd45sa01234+ThIsIsSuPeRsEcReT` |

 
