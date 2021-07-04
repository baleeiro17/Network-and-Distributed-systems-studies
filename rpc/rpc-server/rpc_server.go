package rpc_server

// data structures used in RPC server.
type Valor struct {
	N1 int
	N2 int
}

type Handler struct {
}

// Handler allocation.
func new() *Handler {
	h := &Handler{}
	return h
}

// remote procedures function used in RPC server.
func (*Handler) Sum(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 + values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Subtraction(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 - values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Division(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 / values.N2

	// return the value
	*reply = resul

	return nil
}

func (*Handler) Product(values Valor, reply *int) error {

	// implement sum operation
	resul := values.N1 * values.N2

	// return the value
	*reply = resul

	return nil
}
