package v1

import routing "github.com/qiangxue/fasthttp-routing"

func (h *APIHandler) UpdateProfile(c *routing.Context) error {
	c.SuccessString("application/json", "[UPDATE PROFILE]: hello net")
	return nil
}

func (h *APIHandler) UpdatePassword(c *routing.Context) error {
	c.SuccessString("application/json", "[UPDATE ACCESS STRING]: hello net")
	return nil
}

func (h *APIHandler) All(c *routing.Context) error {
	c.SuccessString("application/json", "[Get ALL]: hello net")
	return nil
}

func (h *APIHandler) GetOne(c *routing.Context) error {
	c.SuccessString("application/json", "[GET ONE]: hello net")
	return nil
}

func (h *APIHandler) Delete(c *routing.Context) error {
	c.SuccessString("application/json", "[DELETE ONE]: hello net")
	return nil
}
