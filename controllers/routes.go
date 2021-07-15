package controllers

func (s *Server) initializeRoutes() {
	s.Router.Get("/", s.HomeController)
	s.Router.Post("/fake-transactions", s.FakeTransactionController)
	s.Router.Get("/transactions", s.GetAllTransactionsController)
	s.Router.Post("/transactions", s.AddTransactionController)
	s.Router.Get("/transactions/{id}", s.GetTransactionController)
	s.Router.Put("/transactions/{id}", s.EditTransactionController)
	s.Router.Delete("/transactions/{id}", s.DeleteTransactionController)
}
