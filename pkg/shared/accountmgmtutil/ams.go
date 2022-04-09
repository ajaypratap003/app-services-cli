package accountmgmtutil

import (
	"context"
	"errors"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/shared/connection"

	amsclient "github.com/redhat-developer/app-services-sdk-go/accountmgmt/apiv1/client"

	"github.com/redhat-developer/app-services-cli/pkg/shared/remote"
)

func CheckTermsAccepted(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (accepted bool, redirectURI string, err error) {
	termsReview, _, err := conn.API().AccountMgmt().
		ApiAuthorizationsV1SelfTermsReviewPost(ctx).
		SelfTermsReview(amsclient.SelfTermsReview{
			EventCode: &spec.TermsAndConditionsEventCode,
			SiteCode:  &spec.TermsAndConditionsSiteCode,
		}).Execute()
	if err != nil {
		return false, "", err
	}

	if !termsReview.GetTermsAvailable() && !termsReview.GetTermsRequired() {
		return true, "", nil
	}

	if !termsReview.HasRedirectUrl() {
		return false, "", errors.New("terms must be signed, but there is no terms URL")
	}

	return false, termsReview.GetRedirectUrl(), nil
}

func GetUserSupportedInstanceTypes(ctx context.Context, spec remote.AmsConfig, conn connection.Connection) (quota []string, err error) {
	orgId, err := GetOrganizationID(ctx, conn)
	if err != nil {
		return nil, err
	}

	quotaCostGet, _, err := conn.API().AccountMgmt().
		ApiAccountsMgmtV1OrganizationsOrgIdQuotaCostGet(ctx, orgId).
		FetchRelatedResources(true).
		Execute()
	if err != nil {
		return nil, err
	}

	var quotas []string
	for _, quota := range quotaCostGet.GetItems() {
		for _, relatedResources := range quota.GetRelatedResources() {
			Product := strings.TrimSpace(relatedResources.GetProduct())
			ResouceName := strings.TrimSpace(relatedResources.GetResourceName())

			if Product == spec.QuotaProductId && ResouceName == spec.ResourceName {
				quotas = append(quotas, QuotaStandardType)
			}

			if Product == spec.TrialQuotaProductId && ResouceName == spec.ResourceName {
				quotas = append(quotas, QuotaTrialType)
			}
		}
	}

	return quotas, err
}

func GetOrganizationID(ctx context.Context, conn connection.Connection) (accountID string, err error) {
	account, _, err := conn.API().AccountMgmt().ApiAccountsMgmtV1CurrentAccountGet(ctx).
		Execute()
	if err != nil {
		return "", err
	}

	return account.Organization.GetId(), nil
}
