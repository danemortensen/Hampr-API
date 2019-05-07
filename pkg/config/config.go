package config

import (
    "os"
)

type MongoConfig struct {
    Ip      string
    DbName  string
}

type AuthConfig struct {
    AppId string
    AppSecret string
}

type ServerConfig struct {
    Port    string
    Auth *AuthConfig
}

type Config struct {
    Mongo   *MongoConfig
    Server  *ServerConfig
}

func NewConfig() *Config {
    mongoConfig := &MongoConfig {
        Ip: getEnv("HAMPR_MONGO_IP", "localhost:27017"),
        DbName: getEnv("HAMPR_MONGO_DB", "hamprdb"),
    }
    authConfig := &AuthConfig {
        AppId: os.Getenv("HAMPR_APP_ID"),
        AppSecret: os.Getenv("HAMPR_APP_SECRET"),
    }
    serverConfig := &ServerConfig {
        Port: getEnv("HAMPR_SERVER_PORT", ":8080"),
        Auth: authConfig,
    }
    return &Config {
        Mongo: mongoConfig,
        Server: serverConfig,
    }
}

func getEnv(envName string, defVal string) string {
    envVal := os.Getenv(envName)
    if envVal != "" {
        return envVal
    }
    return defVal
}
