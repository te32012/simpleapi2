package auth

// Интерфейс через который происходит проверка прав для данного токента
type AuthInterface interface {
	HasPermission(s string, permission int) (bool, error)
	// Отладочный метод
	// Используется для загрузки данных в редис
	CreateUserTokentWithPermission(s string, permission int) error
}
