package queries

const plantHumidityLevelQuery = "INSERT INTO garden_plant_humidity_levels (id, date, currently_on, current_humidity, plant_id) VALUES (DEFAULT, DEFAULT, $1, $2, $3)"
const insertUserQuery = "INSERT INTO garden_users (id, \"user\", crypto, salt, write_permission, read_permission) VALUES (DEFAULT, $1, $2, $3, $4, $5)"
const doesUserExistQuery = "select id from garden_users where \"user\"=$1"
const doesPlantExistQuery = "select min_humidity_levels from garden_plants_info where id=$1"
const removeUserQuery = "DELETE FROM garden_users WHERE \"user\" = $1"
const getSessionDateQuery = "select date from garden_sessions WHERE \"id\"=$1"
const getUserSalt = "SELECT salt from garden_users where \"user\" = $1"
const verifyCredentialsQuery = "select id from garden_users where \"user\"=$1 and \"password\"=$2"
const getSessionIdQuery = "select id from garden_sessions where \"user\"=$1"