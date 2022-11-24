package main

import (
	"context"
	"encoding/json"
	"fmt"
	"garage-vault/api/utils"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type GoogleUserInfo struct {
	UserId       *string `json:"userId"`
	ProviderName *string `json:"providerName"`
	ProviderType *string `json:"providerType"`
	Primary      bool    `json:"primary"`
	Sub          *string `json:"sub"`
}

func getUserGoogleGroups(ctx context.Context, email string) []string {
	// Create a new JWT token to authorize server-to-server Google Cloud Storage API calls
	token, err := google.JWTConfigFromJSON(SERVICE_KEY, "")
	if err != nil {
		log.Panicf("failed to create authenticated transport: %v.", err)
	}
	// Create an authenticated HTTP client (expired tokens get refreshed automatically)
	client := token.Client(ctx)

	service, err := admin.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Panicf("failed to get google admin service: %v.", err)
	}
	res, err := service.Groups.List().Domain("garageisep.com").UserKey(email).Do()
	if err != nil {
		log.Panicf("failed to get groups: %v.", err)
	}
	groups := make([]string, len(res.Groups))
	for _, group := range res.Groups {
		fmt.Println(group.Name)
		groups = append(groups, group.Name)
	}
	return groups
}

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPreTokenGen) (interface{}, error) {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var userinfo []GoogleUserInfo
	identitiesStr := event.Request.UserAttributes["identities"]
	if err := json.Unmarshal([]byte(identitiesStr), &userinfo); err != nil || len(userinfo) == 0 {
		if err == nil {
			err = fmt.Errorf("no identities found")
		}
		return nil, err
	}
	fmt.Printf("Connected user: %+v\n", userinfo)

	// Get user groups
	groups := getUserGoogleGroups(ctx, event.UserName)
	fmt.Printf("User groups: %+v\n", groups)

	cognito := cognitoidentityprovider.New(session)
	// Get user AWS groups
	event.Request.GroupConfiguration.GroupsToOverride = []string{}
	for _, group := range groups {
		// Check if group exists
		// TODO: Implement correct behavior
		_, err := cognito.GetGroup(&cognitoidentityprovider.GetGroupInput{
			GroupName:  aws.String(group),
			UserPoolId: aws.String(os.Getenv("USER_POOL_ID")),
		})
		if err != nil {
			fmt.Printf("Group %s does not exist\n", group)
			continue
		}
		event.Request.GroupConfiguration.GroupsToOverride = append(event.Request.GroupConfiguration.GroupsToOverride, group)
	}
	fmt.Printf("User AWS groups: %+v\n", event.Request.GroupConfiguration.GroupsToOverride)
	return nil, nil
}

func main() {
	utils.LambdaStart(HandleRequest)
}
