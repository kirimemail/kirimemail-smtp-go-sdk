package main

import (
	"fmt"
	"log"

	smtpsdk "github.com/kirimemail/kirimemail-smtp-go-sdk"
)

func main() {
	client := smtpsdk.NewClient("your-username", "your-api-token")

	domains, err := client.Domains().List(nil, nil, nil)
	if err != nil {
		log.Fatalf("Failed to list domains: %v", err)
	}

	fmt.Printf("Found %d domains\n", len(domains.Data))

	if len(domains.Data) > 0 {
		domain := domains.Data[0]
		fmt.Printf("Domain: %s, Verified: %v\n", domain.Domain, domain.IsVerified)

		credentials, err := client.Credentials().List(domain.Domain, nil, nil)
		if err != nil {
			log.Printf("Failed to list credentials: %v", err)
		} else {
			fmt.Printf("Found %d credentials\n", len(credentials.Data))
		}

		result, err := client.Validation().ValidateEmail("test@example.com")
		if err != nil {
			log.Printf("Failed to validate email: %v", err)
		} else {
			fmt.Printf("Email validation result: valid=%v\n", result.Data.IsValid)
		}

		quota, err := client.User().GetQuota()
		if err != nil {
			log.Printf("Failed to get quota: %v", err)
		} else {
			fmt.Printf("Quota: %d/%d (%.1f%%)\n", quota.Data.CurrentQuota, quota.Data.MaxQuota, quota.Data.UsagePercentage)
		}
	}
}
