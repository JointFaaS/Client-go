package strategy

type price_first_strategy struct {

}

func (s *price_first_strategy) Select(cloudResources []CloudResource) (CloudResource, error)  {
	return nil, nil
}

func (s *price_first_strategy) SelectN(cloudResources []CloudResource, num int32) ([]CloudResource, error)  {
	return make([]CloudResource, 0), nil
}

func NewPriceFirstStrategy() (ResourceStrategy, error) {
	pfs := price_first_strategy{}
	return &pfs, nil
}