package user

const getUserByUsernameSQL = "SELECT id, username, password, user_role FROM users WHERE username = $1"
