package v1

import routing "github.com/qiangxue/fasthttp-routing"

func (h *AuthHandler) Login(c *routing.Context) error {
	c.SuccessString("application/json", "[LOGIN]: hello net")
	return nil
}

func (h *AuthHandler) Register(c *routing.Context) error {
	c.SuccessString("application/json", "[LOGIN]: hello net")
	return nil
}

func (h *AuthHandler) GetCode(c *routing.Context) error {
	c.SuccessString("application/json", "[LOGIN]: hello net")
	return nil
}

func (h *AuthHandler) Restore(c *routing.Context) error {
	c.SuccessString("application/json", "[LOGIN]: hello net")
	return nil
}
