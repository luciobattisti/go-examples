package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type Tag struct {
	Key   string
	Value string
}

type TaggingInfo struct {
	TagSet     []Tag
	RoleNames  []string
	PolicyArns []string
}

func ParseJson(fpath string) TaggingInfo {

	data, err := os.ReadFile(fpath)

	if err != nil {
		log.Fatal(err)
	}

	var info TaggingInfo
	json.Unmarshal(data, &info)

	return info

}

func TagPoliciesFromRole(role string, client *iam.Client, tags *[]types.Tag) {

	policies, err := client.ListAttachedRolePolicies(

		context.TODO(),
		&iam.ListAttachedRolePoliciesInput{
			RoleName: aws.String(role),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, value := range policies.AttachedPolicies {
		fmt.Printf("Tagging Policy=%s\n", *value.PolicyName)

		client.TagPolicy(
			context.TODO(),
			&iam.TagPolicyInput{
				PolicyArn: value.PolicyArn,
				Tags:      *tags,
			},
		)
	}
}

func main() {

	// Parse arguments
	infoFpathPtr := flag.String("tagging-file", "intl_tags.json", "JSON file including tagging info")
	tagRolePolicies := flag.Bool("tag-role-policies", false, "If true, tag policies attached to role")

	flag.Parse()

	// Load configuration and credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON tag file
	info := ParseJson(*infoFpathPtr)

	// Create tags variable
	tags := []types.Tag{}

	for _, value := range info.TagSet {
		tags = append(
			tags,
			types.Tag{
				Key:   aws.String(value.Key),
				Value: aws.String(value.Value),
			})
	}

	// Create an IAM client
	client := iam.NewFromConfig(cfg)

	// Tag roles
	for _, role := range info.RoleNames {
		fmt.Printf("Tagging Role=%s\n", role)

		client.TagRole(
			context.TODO(),
			&iam.TagRoleInput{
				RoleName: aws.String(role),
				Tags:     tags,
			})

		if *tagRolePolicies {
			TagPoliciesFromRole(role, client, &tags)
		}

	}

	// Tag policies
	for _, arn := range info.PolicyArns {
		fmt.Printf("Tagging Policy=%s\n", arn)
		client.TagPolicy(
			context.TODO(),
			&iam.TagPolicyInput{
				PolicyArn: aws.String(arn),
				Tags:      tags,
			},
		)
	}
}
