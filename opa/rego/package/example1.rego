package example

apps := [
    {
        "name": "web",
        "servers": ["web-0", "web-1", "web-1000", "web-1001", "web-dev"]
    },
    {
        "name": "mysql",
        "servers": ["db-0", "db-1000"]
    },
    {
        "name": "mongodb",
        "servers": ["db-dev"]
    }
]

allow {
    apps.name == "web"
}
