package user

const getUserByUsernameSQL = "SELECT id, username, password, display_name, user_role FROM users WHERE username = $1"
