package middleware

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

func Cors(c *routing.Context) error {
	c.Response.Header.Set("Access-Control-Allow-Credentials", "*")
	c.Response.Header.Set("Access-Control-Allow-Headers", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "*")
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Content-Type", "application/json")

	if c.IsOptions() {
		c.SuccessString("application/json", "")
	}

	return c.Next()
}
