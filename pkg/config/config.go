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
        Ip: getEnv("hampr:mongo:ip", "localhost:27017"),
        DbName: getEnv("hampr:mongo:db", "hampr"),
    }
    authConfig := &AuthConfig {
        AppId: os.Getenv("HAMPR_APP_ID"),
        AppSecret: os.Getenv("HAMPR_APP_SECRET"),
    }
    serverConfig := &ServerConfig {
        Port: getEnv("hampr:server:port", ":8080"),
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
