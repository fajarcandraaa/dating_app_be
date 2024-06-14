package subscription

type (
	UpdatePackageRequest struct {
		UserId      string `json:"user_code"`
		PackageId   int    `json:"package_id"`
		FeatureCode string `json:"feature_code"`
	}
)
