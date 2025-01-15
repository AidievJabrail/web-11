package provider

func (p *Provider) CheckUser(username, password string) (string, error) {
	var name string
	err := p.db.QueryRow("SELECT login FROM users WHERE login = $1 AND password = $2", username, password).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (p *Provider) CreateUser(username, password string) error {
	_, err := p.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", username, password)
	return err
}

func (p *Provider) FoundUser(username string) (int, error) {
	var id int
	err := p.db.QueryRow("SELECT id_user FROM users WHERE login = $1", username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}