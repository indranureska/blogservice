package common

const LIST_OF_USER_SERVICE_PATH = "/user-list"
const FIND_USER_SERVICE_PATH = "/find-user/{usrEmail}"
const CREATE_USER_SERVICE_PATH = "/create-user"
const UPDATE_USER_SERVICE_PATH = "/update-user"
const DELETE_USER_SERVICE_PATH = "/delete-user/{id:[a-zA-Z0-9]*}"
const USER_LOGIN_SERVICE_PATH = "/login"
const USER_LOGOUT_SERVICE_PATH = "/logout"

const BLOG_DB_URI = "mongodb://127.0.0.1:27018"

const BLOG_DB_URI_ENV_KEY = "BLOG_DB_URI"

const DB_OPERATION_TIMEOUT_SECONDS = 5;
