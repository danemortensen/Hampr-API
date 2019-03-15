package config

import (
    "os"
)

type MongoConfig struct {
    Ip      string
    dbName  string
}

type ServerConfig struct {
    Port    string
}

type Config struct {
    Mongo   *MongoConfig
    Server  *ServerConfig
}

func NewConfig() *Config {
    mongoConfig := &MongoConfig {
        Ip: getEnv("hampr:mongo:ip", "localhost:27017"),
        dbName: getEnv("hampr:mongo:db", "hampr"),
    }
    serverConfig := &ServerConfig {
        Port: getEnv("hampr:server:port", ":8080"),
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
