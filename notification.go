package wns

//https://docs.microsoft.com/en-us/uwp/schemas/tiles/tiles-xml-schema-portal

const (
	TypeToast = "toast"
	TypeTile  = "tile"
	TypeBadge = "badge"
)

type NotificationInterface interface {
	GetWnsType() string
	GetXml() (string, error)
}

func NewToast() *baseNotification {
	return newNotification(TypeToast)
}

func NewTile() *baseNotification {
	return newNotification(TypeTile)
}

func NewBadge() *badge {
	badge := &badge{}
	badge.Type = TypeBadge
	return badge
}

func newNotification(notificationType string) *baseNotification {
	baseNotification := &baseNotification{}
	baseNotification.XMLName.Local = notificationType
	baseNotification.Type = notificationType
	return baseNotification
}
