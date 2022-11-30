package main

import (
	"context"
	"fmt"
	"garage-vault/api/utils"

	"github.com/aws/aws-lambda-go/events"
)

type GoogleUserInfo struct {
	UserId       string `json:"userId"`
	ProviderName string `json:"providerName"`
	ProviderType string `json:"providerType"`
	Primary      bool   `json:"primary"`
}

// func getUserGoogleGroups(ctx context.Context, email string) []string {
// 	// Create a new JWT token to authorize server-to-server Google Cloud Storage API calls
// 	token, err := google.JWTConfigFromJSON(SERVICE_KEY, admin.AdminDirectoryGroupReadonlyScope, admin.AdminDirectoryGroupScope)
// 	if err != nil {
// 		log.Panicf("failed to create authenticated transport: %v\n.", err)
// 	}
// 	// Create an authenticated HTTP client (expired tokens get refreshed automatically)
// 	client := token.Client(ctx)

// 	service, err := admin.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Panicf("failed to get google admin service: %v.\n", err)
// 	}
// 	res, err := service.Groups.List().Domain("garageisep.com").UserKey(email).Do()
// 	if err != nil {
// 		log.Panicf("failed to get groups: %v.\n", err)
// 	}
// 	groups := make([]string, len(res.Groups))
// 	for _, group := range res.Groups {
// 		groups = append(groups, group.Name)
// 	}
// 	return groups
// }

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPreAuthentication) (interface{}, error) {

	// session := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	fmt.Println("User info request", event)
	// var userinfo []GoogleUserInfo
	// identitiesStr := event.Request.UserAttributes["identities"]
	// if err := json.Unmarshal([]byte(identitiesStr), &userinfo); err != nil {
	// 	return nil, err
	// }
	// if len(identitiesStr) == 0 {
	// 	return nil, fmt.Errorf("no identities found")
	// }

	// fmt.Printf("Connected user: %+v\n", userinfo)

	// // Get user groups
	// groups := getUserGoogleGroups(ctx, event.Request.UserAttributes["email"])
	// fmt.Printf("User groups: %+v\n", groups)
	// event.Response.ClaimsOverrideDetails.GroupOverrideDetails.GroupsToOverride = groups

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
