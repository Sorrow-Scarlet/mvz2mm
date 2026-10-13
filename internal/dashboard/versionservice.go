package dashboard

type VersionService struct{}

var AppVersion = "Dev Build"

func (v *VersionService) GetCurrentVersion() string {
	return AppVersion
}
