package structs

type User struct{
    Username string `yaml:"username"`
    Email string`yaml:"email"`
    GpgSign bool`yaml:"gpgsign"`
    SigningKey string`yaml:"signingkey"`
}
