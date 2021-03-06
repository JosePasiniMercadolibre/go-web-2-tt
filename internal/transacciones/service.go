package transacciones

type Service interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error)
	UpdateCodigoMonto(id int, codigoTransaccion string, monto float64) (Transaccion, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Transaccion, error) {
	transacciones, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (s *service) Store(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {

	LastId, err := s.repository.LastId()
	if err != nil {
		return Transaccion{}, err
	}
	LastId++

	transaccion, err := s.repository.Store(LastId, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion)
	if err != nil {
		return Transaccion{}, err
	}
	return transaccion, nil
}

func (s *service) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fechaTransaccion string) (Transaccion, error) {
	return s.repository.Update(id, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion)
}

func (s *service) UpdateCodigoMonto(id int, codigoTransaccion string, monto float64) (Transaccion, error) {
	return s.repository.UpdateCodigoMonto(id, codigoTransaccion, monto)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
