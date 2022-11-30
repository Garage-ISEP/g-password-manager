package main

import (
	"context"
	"encoding/json"
	"fmt"
	"garage-vault/api/utils"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

var SERVICE_KEY = []byte(`{
  "type": "service_account",
  "project_id": "keen-airlock-369621",
  "private_key_id": "92346835a0e59d25127938681c1ab6389f2e0bea",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDcFV5Ig3FXszWv\nVwNO9GPCJq9eBDVDXnbuN83Pv/29LN/tvdpNgCRd49EUG1Hxx8/laTTW81eDJ2aa\nFCU2rs8jwny3wERADoFUd2FWwNn+HT9e8Xb6fCu+vT4j29lUR4dQwS7CXxh0oVnW\n8XoSnfCynBsNyEWeyx1BL9zu/JzbHFa/aG3aCEt72g5nlXbilMGinIFnZA2NlVtf\n09uFCOos5bw/cRhpY8zYjbWB4OySfLGaIrIQt4AiKAsXxfDscBNVIofZ3PNbLRln\nYyoM+7GSpjgWhMFvCz0wMgFm9TTmSH5enMQ4t3VXsA8sRcOJtN6OVh88VUkMrKcO\n+pl3Q4vvAgMBAAECggEAU/6DwUEJHON0dbrLtHJpMEv7WfJZMaD32urJTaSEfpKF\n1AJFMpbZSkwMG5RhrdkIaxG/7NlqaM+8xMkzeq80tv0fBDII2jmu+kGlkKqBiA45\nhU66jdaeC2nGMYFSbGvwJM3VvrD0cG60DuiSIlDM8B3L5XKtv3DddWeC9mLKqoRt\nFUEZ7XaXwOJhA0dpE6XIne9/k/knltqBlKVOB9qlm3UfFWp6bf6p1oWcPSZJThIG\nLGjlMqFuoFpHuN3D+U4jKSo3/p0ap+ivsjzgYOtchoaBu223d2RNjsCIm48wq/5d\n+B6Axb88jw1C5fcgaROEXW/WMDTPS4zjMimi6ulCYQKBgQDwUfDyAtHe69w3/O4q\n5vhqAurwAUwGBQonPlA0N/KqlqXBcp31H37cqE8a2hAgi0n/7hMhPj4Jjh5h/9sz\njyrOIeqJM9njHAx1lwJvZckTvDM0mQ1Yzj2VgJS7KrHBZRGAn07A5yeFGQvDBbcy\nWh0Kb6krqGjpD2lg155+aDLljQKBgQDqcWpkJqGs4pZcpJ7mgFnaOfSnWSYzAw6J\nLDYLTnjRWpzjZa4WohwLsDBUIMS8Qy64ZwED84cAZ/oFQkZlhA94TP89UdPOKJN7\nS0ZgsFSAGH73MC08cTEr2eJjYkyUGgOXSgF8f92eoI1uy0e3TJ7bB2teboOdHAxN\n4b64ThyCawKBgEvyKUnh+D8RnJOY/A9U1LZz25kjX8obN501RRVrhOXCG/npZd+8\nJ1RGYFFlDmmqeyVFMIh16hcM8a8Ys0Y0/VsNPthDKZ1rFMjogx1/Ni9lb2003RHu\ng2nGq7oGgQxUC1bxgWrW4hde1ee4268u6TSOGxEv9I+KjkvLeEwMU+FVAoGBAKMr\nbr3nuUzno3k1RjbtjBv7jFDRFmoYrYxBThgOdL5ZD6qhAUpVZ6mm5ewXpnXVWHYV\nkrpaS1K4C2aPmwlaoZ28EUVvZzgsO4Frgb8X2qgQ5IVMbZ+MtIltq0g6iUvm0yr/\n4grxv6CA5A8BwpdL56BQCKV/y3CUreeiv6ftDaCjAoGARsR6p7R/6Oo+L1VL0MeQ\n/lEx7SQxHAFOpQ8SQG2IWN4OaOimpj8CJ8cyVwDMiUuThYg/SeMv/4Q6GaLjyINI\nmaVgpC0i+3+5iXW4zP68SlnCdw/CeN93YPMm/iKMa/RO6exFGkIgLn/Oma8GM6BP\nba6ugmtawaaHb/ZsyaYXMx4=\n-----END PRIVATE KEY-----\n",
  "client_email": "testtest@keen-airlock-369621.iam.gserviceaccount.com",
  "client_id": "107555657776634822032",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/testtest%40keen-airlock-369621.iam.gserviceaccount.com"
}`)

type GoogleUserInfo struct {
	UserId       string `json:"userId"`
	ProviderName string `json:"providerName"`
	ProviderType string `json:"providerType"`
	Primary      bool   `json:"primary"`
}

func getUserGoogleGroups(ctx context.Context, email string) []string {
	// Create a new JWT token to authorize server-to-server Google Cloud Storage API calls
	token, err := google.JWTConfigFromJSON(SERVICE_KEY, admin.AdminDirectoryGroupReadonlyScope, admin.AdminDirectoryGroupScope)
	if err != nil {
		log.Panicf("failed to create authenticated transport: %v\n.", err)
	}
	// Create an authenticated HTTP client (expired tokens get refreshed automatically)
	client := token.Client(ctx)

	service, err := admin.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Panicf("failed to get google admin service: %v.\n", err)
	}
	res, err := service.Groups.List().Domain("garageisep.com").UserKey(email).Do()
	if err != nil {
		log.Panicf("failed to get groups: %v.\n", err)
	}
	groups := make([]string, len(res.Groups))
	for _, group := range res.Groups {
		groups = append(groups, group.Name)
	}
	return groups
}

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (interface{}, error) {

	// session := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	fmt.Println("User info request", event.Request.UserAttributes)
	var userinfo []GoogleUserInfo
	identitiesStr := event.Request.UserAttributes["identities"]
	if err := json.Unmarshal([]byte(identitiesStr), &userinfo); err != nil {
		return nil, err
	}
	if len(identitiesStr) == 0 {
		return nil, fmt.Errorf("no identities found")
	}

	fmt.Printf("Connected user: %+v\n", userinfo)

	// Get user groups
	groups := getUserGoogleGroups(ctx, event.Request.UserAttributes["email"])
	fmt.Printf("User groups: %+v\n", groups)
	event.Response.ClaimsOverrideDetails.GroupOverrideDetails.GroupsToOverride = groups

	// cognito := cognitoidentityprovider.New(session)
	// // Get user AWS groups
	// event.Request.GroupConfiguration.GroupsToOverride = []string{}
	// for _, group := range groups {
	// 	// Check if group exists
	// 	// TODO: Implement correct behavior
	// 	_, err := cognito.GetGroup(&cognitoidentityprovider.GetGroupInput{
	// 		GroupName:  aws.String(group),
	// 		UserPoolId: aws.String(os.Getenv("USER_POOL_ID")),
	// 	})
	// 	if err != nil {
	// 		fmt.Printf("Group %s does not exist\n", group)
	// 		continue
	// 	}
	// 	event.Request.GroupConfiguration.GroupsToOverride = append(event.Request.GroupConfiguration.GroupsToOverride, group)
	// }
	return nil, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
